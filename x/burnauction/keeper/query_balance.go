package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/burnauction/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Balance(goCtx context.Context, req *types.QueryBalanceRequest) (*types.QueryBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	addr := k.accKeeper.GetModuleAddress(types.ModuleName)
	balance := k.bankKeeper.GetAllBalances(ctx, addr)

	return &types.QueryBalanceResponse{ModuleAddress: addr.String(), Tokens: balance}, nil
}
