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

	if !depositor.DepositorAddress.Equals(investorAddress) {
		return nil, coserrors.Wrap(types.ErrUnauthorized, "not the depositor")
	}

	if depositor.DepositType != types.DepositorInfo_unset {
		return nil, errors.New(" you are in transferring ownership status")
	}

	poolInfo, found := k.GetPools(ctx, msg.PoolIndex)
	if !found {
		return nil, types.ErrPoolNotFound
	}

	if len(poolInfo.PoolNFTIds) == 0 {
		return nil, coserrors.Wrapf(types.ErrUnexpectedEndOfGroupNft, "no borrow can be found")
	}

	dueDate := poolInfo.ProjectDueTime
	oneMonthBeforeProjectDueDate := dueDate.Add(-time.Second * time.Duration(OneMonth))
	twoMonthBeforeProjectDueDate := dueDate.Add(-time.Second * time.Duration(OneMonth*2))

	currentTime := ctx.BlockTime()
	if currentTime.Before(twoMonthBeforeProjectDueDate) {
		return nil, coserrors.Wrapf(types.ErrTime, "submit the proposal too early")
	}

	if currentTime.After(oneMonthBeforeProjectDueDate) {
		return nil, coserrors.Wrapf(types.ErrTime, "submit the proposal too late")
	}

	//totalBorrowedNow, err := calculateTotalPrinciple(ctx, depositor.LinkedNFT, k.nftKeeper)
	//if err != nil {
	//	return nil, err
	//}

	//can be negative, we now sync the locked amount and withdraw amount
	//deltaAmount := depositor.LockedAmount.Amount.Sub(totalBorrowedNow)
	//depositor.LockedAmount = sdk.NewCoin(depositor.LockedAmount.Denom, totalBorrowedNow)
	//depositor.WithdrawalAmount = depositor.WithdrawalAmount.AddAmount(deltaAmount)

	depositor.DepositType = types.DepositorInfo_withdraw_proposal
	poolInfo.WithdrawProposalAmount = poolInfo.WithdrawProposalAmount.Add(depositor.LockedAmount)
	// since we can not borrow from this investor, we deduct the total borrowable amount
	poolInfo.BorrowableAmount = poolInfo.BorrowableAmount.Sub(depositor.WithdrawalAmount)
	poolInfo.WithdrawAccounts = append(poolInfo.WithdrawAccounts, depositor.DepositorAddress)
	k.SetPool(ctx, poolInfo)
	k.SetDepositor(ctx, depositor)

	return &types.MsgSubmitWithdrawProposalResponse{}, nil
}
