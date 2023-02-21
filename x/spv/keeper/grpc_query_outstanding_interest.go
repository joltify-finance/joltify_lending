package keeper

import (
	"context"

    "github.com/joltify-finance/joltify_lending/x/spv/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) OutstandingInterest(goCtx context.Context,  req *types.QueryOutstandingInterestRequest) (*types.QueryOutstandingInterestResponse, error) {
	if req == nil {
        return nil, status.Error(codes.InvalidArgument, "invalid request")
    }

	ctx := sdk.UnwrapSDKContext(goCtx)

    // TODO: Process the query
    _ = ctx

	return &types.QueryOutstandingInterestResponse{}, nil
}
