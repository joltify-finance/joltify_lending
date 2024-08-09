package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/kyc/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListInvestors(goCtx context.Context, req *types.ListInvestorsRequest) (*types.ListInvestorsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(k.storeKey)
	investorStores := prefix.NewStore(store, types.KeyPrefix(types.InvestorToWalletsPrefix))

	var investors []*types.Investor
	pageRes, err := query.Paginate(investorStores, req.Pagination, func(key []byte, value []byte) error {
		var investor types.Investor
		if err := k.cdc.Unmarshal(value, &investor); err != nil {
			return err
		}
		investors = append(investors, &investor)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.ListInvestorsResponse{Investor: investors, Pagination: pageRes}, nil
}
