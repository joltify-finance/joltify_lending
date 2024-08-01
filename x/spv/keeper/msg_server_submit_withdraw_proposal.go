package keeper

import (
	"context"
	"time"

	coserrors "cosmossdk.io/errors"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k msgServer) SubmitWithdrawProposal(goCtx context.Context, msg *types.MsgSubmitWithdrawProposal) (*types.MsgSubmitWithdrawProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	ctx = ctx.WithGasMeter(storetypes.NewInfiniteGasMeter())

	investorAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address %v", msg.Creator)
	}

	depositor, found := k.GetDepositor(ctx, msg.PoolIndex, investorAddress)
	if !found {
		return nil, coserrors.Wrapf(types.ErrDepositorNotFound, "depositor %v not found for pool index %v", msg.Creator, msg.PoolIndex)
	}

	if depositor.DepositType != types.DepositorInfo_unset {
		return nil, coserrors.Wrapf(types.ErrUNEXPECTEDSTATUS, "depositor %v is not in unset status (current status %v)", msg.Creator, depositor.DepositType)
	}

	poolInfo, found := k.GetPools(ctx, msg.PoolIndex)
	if !found {
		return nil, types.ErrPoolNotFound
	}

	if len(poolInfo.PoolNFTIds) == 0 {
		return nil, coserrors.Wrapf(types.ErrUnexpectedEndOfGroupNft, "no borrow can be found")
	}

	if poolInfo.PoolStatus != types.PoolInfo_PooLPayPartially {
		if poolInfo.PoolStatus != types.PoolInfo_ACTIVE {
			return nil, coserrors.Wrapf(types.ErrPoolNotActive, "pool status is %v not in submit withdraw request or active status", poolInfo.PoolStatus)
		}
	}

	dueDate := poolInfo.ProjectDueTime
	firstTimeStampBeforeProjectDueDate := dueDate.Add(-time.Second * time.Duration(poolInfo.WithdrawRequestWindowSeconds*2))
	secondTimeStampBeforeProjectDueDate := dueDate.Add(-time.Second * time.Duration(poolInfo.WithdrawRequestWindowSeconds*3))

	currentTime := ctx.BlockTime()
	if currentTime.Before(secondTimeStampBeforeProjectDueDate) {
		return nil, coserrors.Wrapf(types.ErrTime, "submit the proposal too early with earliest time %v(current: %v)", secondTimeStampBeforeProjectDueDate.Local(), currentTime.Local())
	}

	if currentTime.After(firstTimeStampBeforeProjectDueDate) {
		return nil, coserrors.Wrapf(types.ErrTime, "submit the proposal too late with latest time %v (current: %v)", firstTimeStampBeforeProjectDueDate.Local(), currentTime.Local())
	}

	depositor.DepositType = types.DepositorInfo_withdraw_proposal
	poolInfo.WithdrawProposalAmount = poolInfo.WithdrawProposalAmount.Add(depositor.LockedAmount)
	// since we can not borrow from this investor, we deduct the total borrowable amount
	poolInfo.UsableAmount = poolInfo.UsableAmount.Sub(depositor.WithdrawalAmount)
	poolInfo.WithdrawAccounts = append(poolInfo.WithdrawAccounts, depositor.DepositorAddress)
	poolInfo.PoolStatus = types.PoolInfo_PooLPayPartially
	k.SetPool(ctx, poolInfo)
	k.SetDepositor(ctx, depositor)

	return &types.MsgSubmitWithdrawProposalResponse{}, nil
}
