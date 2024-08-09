package keeper

import (
	"context"

	"github.com/joltify-finance/joltify_lending/dydx_helper/lib"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/vest/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VestEntry(
	goCtx context.Context,
	req *types.QueryVestEntryRequest,
) (*types.QueryVestEntryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := lib.UnwrapSDKContext(goCtx, types.ModuleName)
	vestEntry, err := k.GetVestEntry(ctx, req.VesterAccount)
	if err != nil {
		return nil, err
	}

	return &types.QueryVestEntryResponse{Entry: vestEntry}, nil
}
