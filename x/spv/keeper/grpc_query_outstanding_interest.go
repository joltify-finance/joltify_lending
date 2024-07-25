package keeper

import (
	"context"

	coserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) OutstandingInterest(goCtx context.Context, req *types.QueryOutstandingInterestRequest) (*types.QueryOutstandingInterestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	investor, err := sdk.AccAddressFromBech32(req.Wallet)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address %v", req.Wallet)
	}

	depositor, found := k.GetDepositor(ctx, req.PoolIndex, investor)
	if !found {
		return nil, coserrors.Wrapf(types.ErrDepositorNotFound, "depositor not found for pool %v", req.PoolIndex)
	}

	poolInfo, found := k.GetPools(ctx, depositor.PoolIndex)
	if !found {
		panic("should never fail to find the pool")
	}

	if poolInfo.PoolStatus == types.PoolInfo_FROZEN {
		return &types.QueryOutstandingInterestResponse{Amount: "0"}, nil
	}

	reserve := poolInfo.ReserveFactor
	lendNFTs := depositor.LinkedNFT
	totalInterest, err := calculateTotalOutstandingInterest(ctx, lendNFTs, k.NftKeeper, reserve)
	if err != nil {
		return nil, err
	}
	return &types.QueryOutstandingInterestResponse{Amount: totalInterest.String()}, nil
}
