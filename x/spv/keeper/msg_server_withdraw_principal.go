package keeper

import (
	"context"
	"errors"

	coserrors "cosmossdk.io/errors"
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

	depositor.WithdrawalAmount, err = depositor.WithdrawalAmount.SafeSub(totalWithdraw)
	if err != nil {
		return nil, errors.New("withdraw amount too large")
	}

	if poolInfo.PoolStatus == types.PoolInfo_CLOSED {
		amount := k.cleanupDepositor(ctx, poolInfo, depositor)
		tokenSend := sdk.NewCoin(msg.Token.Denom, amount)
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccount, investor, sdk.NewCoins(tokenSend))
		if err != nil {
			return nil, err
		}
		return &types.MsgWithdrawPrincipalResponse{}, nil
	}

	// if withdraw >= withdrawable

	if depositor.DepositType == types.DepositorInfo_deposit_close {
		k.DelDepositor(ctx, depositor.PoolIndex, depositor.DepositorAddress)
	}

	if depositor.DepositType == types.DepositorInfo_unset {
		poolInfo.BorrowableAmount = poolInfo.BorrowableAmount.SubAmount(totalWithdraw.Amount)
	}

	k.SetDepositor(ctx, depositor)
	k.SetPool(ctx, poolInfo)
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccount, investor, sdk.NewCoins(totalWithdraw))
	if err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeWithdrawPrincipal,
			sdk.NewAttribute(types.AttributeCreator, msg.Creator),
			sdk.NewAttribute(types.AttributeAmount, msg.Token.String()),
		),
	)

	return &types.MsgWithdrawPrincipalResponse{}, nil
}
