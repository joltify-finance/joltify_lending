package keeper

import (
	"context"
	coserrors "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k msgServer) ClaimInterest(goCtx context.Context, msg *types.MsgClaimInterest) (*types.MsgClaimInterestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	investorAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address %v", msg.Creator)
	}

	depositor, found := k.GetDepositor(ctx, msg.PoolIndex, investorAddress)
	if !found {
		return nil, coserrors.Wrapf(types.ErrDepositorNotFound, "depositor %v not found for pool index %v", msg.Creator, msg.GetPoolIndex())
	}

	if !depositor.DepositorAddress.Equals(investorAddress) {
		return nil, coserrors.Wrap(types.ErrUnauthorized, "not the depositor")
	}

	if depositor.DepositType == types.DepositorInfo_deposit_close {
		return nil, coserrors.Wrapf(types.ErrUnauthorized, "your deposit has been closed")
	}

	claimed, err := k.claimInterest(ctx, &depositor)
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccount, investorAddress, sdk.NewCoins(claimed))
	if err != nil {
		return nil, err
	}

	k.SetDepositor(ctx, depositor)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeClaimInterest,
			sdk.NewAttribute(types.AttributeCreator, msg.Creator),
			sdk.NewAttribute("amount", claimed.String()),
		),
	)

	return &types.MsgClaimInterestResponse{Amount: claimed.String()}, nil
}

func (k Keeper) claimInterest(ctx sdk.Context, depositor *types.DepositorInfo) (sdk.Coin, error) {

	// for each lending NFT this owner has
	totalInterest, err := calculateTotalInterest(ctx, depositor.LinkedNFT, k.nftKeeper, true)
	if err != nil {
		return sdk.Coin{}, err
	}

	claimed := sdk.NewCoin(depositor.LockedAmount.Denom, totalInterest)

	// we add the pending one
	claimed = claimed.Add(depositor.PendingInterest)

	depositor.PendingInterest = sdk.NewCoin(depositor.GetPendingInterest().Denom, sdk.ZeroInt())

	poolInfo, found := k.GetPools(ctx, depositor.PoolIndex)
	if !found {
		panic("pool must be found")
	}

	if poolInfo.EscrowInterestAmount.IsNegative() {
		return sdk.Coin{}, coserrors.Wrapf(types.ErrClaimInterest, "not enough interest to be paid")
	}
	return claimed, nil

}
