package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetQuota(goCtx context.Context, req *types.QueryGetQuotaRequest) (*types.QueryGetQuotaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req.QueryLength > 100 {
		return nil, status.Error(codes.OutOfRange, "query item should small than 100")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	result, found := k.GetQuotaData(ctx)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "Coin quota data is not ready")
	}

	result.History = result.History[:req.QueryLength]

	return &types.QueryGetQuotaResponse{CoinQuotaResponse: result}, nil
}
