package keeper

import (
	"context"
	"fmt"

	coserrors "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

// TransferOwnership will allow the investor to submit the request to transfer/update their ratio in the pool
func (k msgServer) TransferOwnership(goCtx context.Context, msg *types.MsgTransferOwnership) (*types.MsgTransferOwnershipResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	ctx = ctx.WithGasMeter(sdk.NewInfiniteGasMeter())

	caller, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address %v", msg.Creator)
	}

	d, found := k.GetDepositor(ctx, msg.PoolIndex, caller)
	if !found {
		return &types.MsgTransferOwnershipResponse{}, types.ErrDepositorNotFound
	}

	if len(d.LinkedNFT) == 0 {
		return &types.MsgTransferOwnershipResponse{}, coserrors.Wrapf(types.ErrDepositorNotFound, "no borrow nft to transfer")
	}

	if d.DepositType != types.DepositorInfo_unset {
		return &types.MsgTransferOwnershipResponse{}, fmt.Errorf("you have submitted the %v request", d.DepositType)
	}

	poolInfo, found := k.GetPools(ctx, msg.GetPoolIndex())
	if !found {
		return &types.MsgTransferOwnershipResponse{}, types.ErrPoolNotFound
	}

	if poolInfo.PoolStatus != types.PoolInfo_ACTIVE {
		return &types.MsgTransferOwnershipResponse{}, coserrors.Wrapf(types.ErrUNEXPECTEDSTATUS, "pool current status is %v", poolInfo.PoolStatus)
	}

	// dueDate := poolInfo.ProjectDueTime
	// secondTimeStampBeforeProjectDueDate := dueDate.Add(-time.Second * time.Duration(poolInfo.WithdrawRequestWindowSeconds*2))
	// if ctx.BlockTime().After(secondTimeStampBeforeProjectDueDate.Add(-time.Minute)) && ctx.BlockTime().Before(dueDate.Add(time.Minute)) {
	//	return &types.MsgTransferOwnershipResponse{}, coserrors.Wrapf(types.ErrUNEXPECTEDSTATUS, "you can not transfer the nft in the request during the spv principal preparation window")
	//}

	poolInfo.TransferAccounts = append(poolInfo.TransferAccounts, caller)
	poolInfo.TotalTransferOwnershipAmount = poolInfo.TotalTransferOwnershipAmount.Add(d.LockedAmount)
	d.DepositType = types.DepositorInfo_transfer_request
	poolInfo.UsableAmount, err = poolInfo.UsableAmount.SafeSub(d.WithdrawalAmount)
	if err != nil {
		return &types.MsgTransferOwnershipResponse{}, coserrors.Wrapf(err, "fail to update the borrowable")
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, caller, sdk.NewCoins(d.WithdrawalAmount))
	if err != nil {
		return &types.MsgTransferOwnershipResponse{}, coserrors.Wrapf(err, "fail to send the leftover withdrawable to investor")
	}

	d.WithdrawalAmount = d.WithdrawalAmount.SubAmount(d.WithdrawalAmount.Amount)

	k.SetDepositor(ctx, d)
	k.SetPool(ctx, poolInfo)

	return &types.MsgTransferOwnershipResponse{OperationResult: true}, nil
}
