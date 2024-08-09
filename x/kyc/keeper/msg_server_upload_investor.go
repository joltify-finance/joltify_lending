package keeper

import (
	"context"
	"errors"

	errorsmod "cosmossdk.io/errors"
	types2 "cosmossdk.io/store/types"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/kyc/types"
)

func updateList(in []string, added []string) []string {
	if len(in)+len(added) < types.MaxWalletNum {
		in = append(in, added...)
		return in
	}

	// otherwise we need to pop n wallet and inset the new one
	delta := len(added) + len(in) - types.MaxWalletNum
	walletsNew := in[delta:]
	walletsNew = append(walletsNew, added...)
	in = walletsNew
	return walletsNew
}

func (k Keeper) GetInvestor(rctx context.Context, investorID string) types.Investor {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InvestorToWalletsPrefix))
	var storedInvestor types.Investor
	b := store.Get(types.KeyPrefix(investorID))
	if b == nil {
		return types.Investor{}
	}
	k.cdc.MustUnmarshal(b, &storedInvestor)
	return storedInvestor
}

// SetInvestor set a specific issueToken in the store from its index
func (k Keeper) SetInvestor(rctx context.Context, investor types.Investor) *types.Investor {
	ctx := sdk.UnwrapSDKContext(rctx)
	gasBefore := ctx.GasMeter().GasConsumed()
	defer func() {
		gasAfter := ctx.GasMeter().GasConsumed()
		ctx.GasMeter().RefundGas(gasAfter-gasBefore, "setInvestor")
	}()

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InvestorToWalletsPrefix))

	var storedInvestor types.Investor
	b := store.Get(types.KeyPrefix(investor.InvestorId))
	if b == nil {
		data := k.cdc.MustMarshal(&investor)
		store.Set(types.KeyPrefix(investor.InvestorId), data)
		return &investor
	}
	k.cdc.MustUnmarshal(b, &storedInvestor)

	wallets := updateList(storedInvestor.WalletAddress, investor.WalletAddress)
	storedInvestor.WalletAddress = wallets
	data := k.cdc.MustMarshal(&storedInvestor)
	store.Set(types.KeyPrefix(storedInvestor.InvestorId), data)
	return &storedInvestor
}

func (k msgServer) UploadInvestor(rctx context.Context, msg *types.MsgUploadInvestor) (*types.MsgUploadInvestorResponse, error) {
	ctx := sdk.UnwrapSDKContext(rctx)

	ctx = ctx.WithGasMeter(types2.NewInfiniteGasMeter())
	if len(msg.InvestorId) == 0 {
		return nil, errors.New("invalid investor ID")
	}
	submitter, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	submitters := k.GetSubmitter(ctx)
	found := false
	for _, el := range submitters {
		if el.Equals(submitter) {
			found = true
			break
		}
	}
	if !found {
		return nil, errorsmod.Wrapf(types.ErrUnauthorised, "unauthorised submitter: %v", submitter.String())
	}

	if len(msg.WalletAddress) == 0 {
		return &types.MsgUploadInvestorResponse{
			Wallets: []string{},
		}, nil
	}
	// we check whether all the wallet address are valid
	for _, el := range msg.WalletAddress {
		_, err := sdk.AccAddressFromBech32(el)
		if err != nil {
			return nil, errorsmod.Wrapf(types.ErrInvalidWallets, "invalid wallets: %v", el)
		}

		// the wallet has been already registered
		_, err = k.GetByWallet(ctx, el)
		if err == nil {
			return nil, errorsmod.Wrapf(types.ErrInvalidWallets, "fail to check the wallets against existing data: %v", msg.WalletAddress)
		}
	}

	// we check whether we exceed the max wallet number
	if len(msg.WalletAddress) > types.MaxWalletNum {
		return nil, errorsmod.Wrapf(types.ErrExceedMaxWalletsNum, "wallets submitted %v", len(msg.WalletAddress))
	}

	investor := types.Investor{InvestorId: msg.InvestorId, WalletAddress: msg.WalletAddress}
	ret := k.SetInvestor(ctx, investor)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(types.AttributeWalletsSubmitted, msg.Creator),
		),
	)

	return &types.MsgUploadInvestorResponse{Wallets: ret.GetWalletAddress()}, nil
}
