package keeper

import (
	"context"
	"time"

	coserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
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

	poolInfo, found := k.GetPools(ctx, msg.PoolIndex)
	if !found {
		return nil, types.ErrPoolNotFound
	}

	if len(poolInfo.PoolNFTIds) == 0 {
		return nil, coserrors.Wrapf(types.ErrUnexpectedEndOfGroupNft, "no borrow can be found")
	}
	// we get the timestamp of the first borrow
	firstBorrowNFT, found := k.nftKeeper.GetClass(ctx, poolInfo.PoolNFTIds[0])
	if !found {
		panic("should never fail")
	}

	var borrowInterest types.BorrowInterest
	err = proto.Unmarshal(firstBorrowNFT.Data.Value, &borrowInterest)
	if err != nil {
		panic(err)
	}

	oneMonthBeforeProjectDueDate := borrowInterest.Payments[0].PaymentTime.Add(time.Duration(poolInfo.ProjectLength-OneMonth) * time.Second)
	twoMonthBeforeProjectDueDate := borrowInterest.Payments[0].PaymentTime.Add(time.Duration(poolInfo.ProjectLength-OneMonth*2) * time.Second)

	currentTime := ctx.BlockTime()
	if currentTime.Before(twoMonthBeforeProjectDueDate) {
		return nil, coserrors.Wrapf(types.ErrTime, "submit the proposal too early")
	}

	if currentTime.After(oneMonthBeforeProjectDueDate) {
		return nil, coserrors.Wrapf(types.ErrTime, "submit the poposal too late")
	}

	totalBorrowedNow, err := calculateTotalPrinciple(ctx, depositor.LinkedNFT, k.nftKeeper)
	if err != nil {
		return nil, err
	}

	//can be negative
	deltaAmount := depositor.LockedAmount.Amount.Sub(totalBorrowedNow)
	depositor.LockedAmount = depositor.LockedAmount.SubAmount(deltaAmount)
	depositor.WithdrawalAmount = depositor.WithdrawalAmount.AddAmount(deltaAmount)

	depositor.WithdrawProposal = true

	poolInfo.WithdrawProposalAmount = poolInfo.WithdrawProposalAmount.Add(depositor.LockedAmount)

	return &types.MsgSubmitWithdrawProposalResponse{}, nil
}
