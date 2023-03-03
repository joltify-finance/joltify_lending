package keeper

import (
	"context"
	coserrors "cosmossdk.io/errors"
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k msgServer) WithdrawPrincipal(goCtx context.Context, msg *types.MsgWithdrawPrincipal) (*types.MsgWithdrawPrincipalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	investor, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address %v", msg.Creator)
	}

	depositor, found := k.GetDepositor(ctx, msg.PoolIndex, investor)
	if !found {
		return nil, coserrors.Wrapf(types.ErrDepositorNotFound, "depositor not found for pool %v", msg.PoolIndex)
	}

	if depositor.WithdrawalAmount.Denom != msg.Token.Denom {
		return nil, coserrors.Wrapf(types.ErrInconsistencyToken, "you can only withdraw %v", depositor.WithdrawalAmount.Denom)
	}

	lendNFTs := depositor.LinkedNFT

	totalBorrowedNow, err := calculateTotalPrinciple(ctx, lendNFTs, k.nftKeeper)
	if err != nil {
		return nil, err
	}

	//can be negative
	deltaAmount := depositor.LockedAmount.Amount.Sub(totalBorrowedNow)
	depositor.LockedAmount = depositor.LockedAmount.SubAmount(deltaAmount)
	depositor.WithdrawalAmount = depositor.WithdrawalAmount.AddAmount(deltaAmount)
	depositor.WithdrawalAmount, err = depositor.WithdrawalAmount.SafeSub(msg.Token)
	if err != nil {
		return nil, errors.New("withdraw amount too large")
	}

	poolInfo, found := k.GetPools(ctx, msg.PoolIndex)
	if !found {
		return nil, errors.New("pool cannot be found")
	}

	poolInfo.BorrowableAmount = poolInfo.BorrowableAmount.SubAmount(msg.Token.Amount)
	k.SetPool(ctx, poolInfo)

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccount, investor, sdk.NewCoins(msg.Token))
	if err != nil {
		return nil, err
	}

	k.SetDepositor(ctx, depositor)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeWithdrawPrincipal,
			sdk.NewAttribute(types.AttributeCreator, msg.Creator),
			sdk.NewAttribute(types.AttributeAmount, msg.Token.String()),
		),
	)

	return &types.MsgWithdrawPrincipalResponse{}, nil
}
