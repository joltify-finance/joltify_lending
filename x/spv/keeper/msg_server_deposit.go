package keeper

import (
	"context"
	coserrors "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	kyctypes "github.com/joltify-finance/joltify_lending/x/kyc/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k msgServer) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	investor, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address %v", msg.Creator)
	}

	poolInfo, ok := k.GetPools(ctx, msg.GetPoolIndex())
	if !ok {
		return nil, coserrors.Wrapf(sdkerrors.ErrNotFound, "pool cannot be found %v", msg.GetPoolIndex())
	}

	req := kyctypes.QueryInvestorWalletsRequest{
		InvestorId: msg.GetInvestorID(),
	}

	resp, err := k.kycKeeper.QueryInvestorWallets(ctx, &req)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrNotFound, "the investor cannot be found %v", msg.InvestorID)
	}

	found := false
	for _, el := range resp.Wallets {
		if el == msg.Creator {
			found = true
			break
		}
	}

	if !found {
		return nil, coserrors.Wrapf(sdkerrors.ErrNotFound, "the given wallet cannot be found %v", msg.Creator)
	}

	if msg.Token.GetDenom() != poolInfo.GetTotalAmount().Denom {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidCoins, "we only accept %v", poolInfo.GetTotalAmount().Denom)
	}

	// now we transfer the token from the investor to the pool.
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, investor, types.ModuleName, sdk.NewCoins(msg.Token))
	if err != nil {
		return nil, coserrors.Wrap(sdkerrors.ErrInvalidRequest, "fail to transfer the toekn to the pool")
	}

	// now we update the users deposit database
	previousDepositor, found := k.GetDepositor(ctx, investor)
	if !found {
		depositor := types.DepositorInfo{InvestorId: msg.InvestorID, DepositorAddress: investor, PoolIndex: msg.PoolIndex, WithdrawableAmount: msg.Token}
		k.SetDepositor(ctx, depositor)

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeDeposit,
				sdk.NewAttribute(types.AttributeCreator, msg.Creator),
			),
		)
		return &types.MsgDepositResponse{}, nil
	}
	previousDepositor.WithdrawableAmount = previousDepositor.WithdrawableAmount.Add(msg.Token)
	k.SetDepositor(ctx, previousDepositor)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeDeposit,
			sdk.NewAttribute(types.AttributeCreator, msg.Creator),
		),
	)

	return &types.MsgDepositResponse{}, nil
}
