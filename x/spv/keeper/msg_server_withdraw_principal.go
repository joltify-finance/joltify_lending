package keeper

import (
	"context"
	"errors"

	coserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k Keeper) handlerPoolClose(ctx sdk.Context, poolInfo types.PoolInfo, depositor types.DepositorInfo) error {
	amount, err := k.cleanupDepositor(ctx, poolInfo, depositor)
	if err != nil {
		return err
	}
	tokenSend := sdk.NewCoin(depositor.LockedAmount.Denom, amount)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeWithdrawPrincipal,
			sdk.NewAttribute(types.AttributeCreator, depositor.DepositorAddress.String()),
			sdk.NewAttribute(types.AttributeAmount, amount.String()),
		),
	)
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccount, depositor.DepositorAddress, sdk.NewCoins(tokenSend))
}

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

	if msg.Token.IsZero() {
		return nil, errors.New("zero amount to withdraw")
	}

	poolInfo, found := k.GetPools(ctx, msg.PoolIndex)
	if !found {
		return nil, errors.New("pool cannot be found")
	}

	totalWithdraw := msg.Token
	if msg.Token.IsGTE(depositor.GetWithdrawalAmount()) {
		totalWithdraw = depositor.GetWithdrawalAmount()
	}

	if poolInfo.PoolStatus == types.PoolInfo_CLOSED {
		err = k.handlerPoolClose(ctx, poolInfo, depositor)
		if err != nil {
			return nil, err
		}
		return &types.MsgWithdrawPrincipalResponse{}, nil
	}

	switch depositor.DepositType {
	case types.DepositorInfo_deposit_close:
		depositor.DepositType = types.DepositorInfo_deactive
		k.SetDepositor(ctx, depositor)
		amountToSend := depositor.WithdrawalAmount
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccount, investor, sdk.NewCoins(amountToSend))
		if err != nil {
			return nil, err
		}
		k.SetPool(ctx, poolInfo)
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeWithdrawPrincipal,
				sdk.NewAttribute(types.AttributeCreator, msg.Creator),
				sdk.NewAttribute(types.AttributeAmount, amountToSend.String()),
			),
		)

		return &types.MsgWithdrawPrincipalResponse{}, nil
	case types.DepositorInfo_unset, types.DepositorInfo_withdraw_proposal, types.DepositorInfo_processed:
		poolInfo.BorrowableAmount = poolInfo.BorrowableAmount.SubAmount(totalWithdraw.Amount)
		depositor.WithdrawalAmount, err = depositor.WithdrawalAmount.SafeSub(totalWithdraw)
		if err != nil {
			return nil, errors.New("withdraw amount too large")
		}
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccount, investor, sdk.NewCoins(totalWithdraw))
		if err != nil {
			return nil, err
		}

		k.SetDepositor(ctx, depositor)
		k.SetPool(ctx, poolInfo)

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeWithdrawPrincipal,
				sdk.NewAttribute(types.AttributeCreator, msg.Creator),
				sdk.NewAttribute(types.AttributeAmount, totalWithdraw.String()),
			),
		)
		return &types.MsgWithdrawPrincipalResponse{}, nil
	default:
		return &types.MsgWithdrawPrincipalResponse{}, coserrors.Wrapf(types.ErrDeposit, "deposit type is %v", depositor.DepositType)
	}
}
