package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TotalReserve(goCtx context.Context, req *types.QueryTotalReserveRequest) (*types.QueryTotalReserveResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var coins sdk.Coins
	k.IterateReserve(ctx, func(coin sdk.Coin) (stop bool) {
		coins = coins.Add(coin)
		return false
	})
	return &types.QueryTotalReserveResponse{Coins: coins.String()}, nil
}
