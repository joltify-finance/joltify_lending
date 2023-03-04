package keeper

import (
	"context"
	"time"

	coserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k msgServer) SubmitWithdrawProposal(goCtx context.Context, msg *types.MsgSubmitWithdrawProposal) (*types.MsgSubmitWithdrawProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)



	// TODO: Handling the message

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

	dueTime:=time.Second*time.Duration( poolInfo.ProjectLength)

	if ctx.BlockTime().After(dueT)

		pool


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
