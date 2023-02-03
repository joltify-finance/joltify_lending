package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/kyc/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QueryInvestorWallets(goCtx context.Context, req *types.QueryInvestorWalletsRequest) (*types.QueryInvestorWalletsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(k.storeKey)
	investorStores := prefix.NewStore(store, types.KeyPrefix(types.InvestorToWalletsPrefix))
	b := investorStores.Get(types.KeyPrefix(req.GetInvestorId()))
	if b == nil {
		return nil, status.Errorf(codes.NotFound, "investor id %v", req.InvestorId)
	}
	var investor types.Investor
	k.cdc.MustUnmarshal(b, &investor)
	return &types.QueryInvestorWalletsResponse{Wallets: investor.WalletAddress}, nil
}
