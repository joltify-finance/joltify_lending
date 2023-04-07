package keeper

import (
	"context"
	"errors"
	"time"

	coserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k msgServer) SubmitWithdrawProposal(goCtx context.Context, msg *types.MsgSubmitWithdrawProposal) (*types.MsgSubmitWithdrawProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	investorAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address %v", msg.Creator)
	}

	depositor, found := k.GetDepositor(ctx, msg.PoolIndex, investorAddress)
	if !found {
		return nil, coserrors.Wrapf(types.ErrDepositorNotFound, "depositor %v not found for pool index %v", msg.Creator, msg.GetPoolIndex())
	}

	if depositor.DepositType != types.DepositorInfo_unset {
		return nil, errors.New("you are in the unexpected status")
	}

	poolInfo, found := k.GetPools(ctx, msg.PoolIndex)
	if !found {
		return nil, types.ErrPoolNotFound
	}

	if len(poolInfo.PoolNFTIds) == 0 {
		return nil, coserrors.Wrapf(types.ErrUnexpectedEndOfGroupNft, "no borrow can be found")
	}

	dueDate := poolInfo.ProjectDueTime
	firstTimeStampBeforeProjectDueDate := dueDate.Add(-time.Second * time.Duration(poolInfo.WithdrawRequestWindowSeconds))
	secondTimeStampBeforeProjectDueDate := dueDate.Add(-time.Second * time.Duration(poolInfo.WithdrawRequestWindowSeconds*2))

	currentTime := ctx.BlockTime()
	if currentTime.Before(secondTimeStampBeforeProjectDueDate) {
		return nil, coserrors.Wrapf(types.ErrTime, "submit the proposal too early with time  %v:%v", secondTimeStampBeforeProjectDueDate.Local(), currentTime.Local())
	}

	if currentTime.After(firstTimeStampBeforeProjectDueDate) {
		return nil, coserrors.Wrapf(types.ErrTime, "submit the proposal too late with time gap %v%v", firstTimeStampBeforeProjectDueDate.Local(), currentTime.Local())
	}

	depositor.DepositType = types.DepositorInfo_withdraw_proposal
	poolInfo.WithdrawProposalAmount = poolInfo.WithdrawProposalAmount.Add(depositor.LockedAmount)
	// since we can not borrow from this investor, we deduct the total borrowable amount
	poolInfo.UsableAmount = poolInfo.UsableAmount.Sub(depositor.WithdrawalAmount)
	poolInfo.WithdrawAccounts = append(poolInfo.WithdrawAccounts, depositor.DepositorAddress)
	k.SetPool(ctx, poolInfo)
	k.SetDepositor(ctx, depositor)

	return &types.MsgSubmitWithdrawProposalResponse{}, nil
}
