package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AllowedPools(goCtx context.Context, req *types.QueryAllowedPoolsRequest) (*types.QueryAllowedPoolsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	ret, err := k.kycKeeper.GetByWallet(ctx, req.WalletAddress)
	if err != nil {
		return nil, err
	}

	var allPools []string
	k.IterateInvestorPools(ctx, func(poolWithInvestors types.PoolWithInvestors) (stop bool) {
		for _, el := range poolWithInvestors.Investors {
			if el == ret.InvestorId {
				allPools = append(allPools, poolWithInvestors.PoolIndex)
				break
			}
		}
		return false
	})

	return &types.QueryAllowedPoolsResponse{PoolsIndex: allPools}, nil
}
