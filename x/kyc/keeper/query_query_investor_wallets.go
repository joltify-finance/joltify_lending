package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"

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
	investorID := req.GetInvestorId()
	ret, err := k.GetInvestorWallets(ctx, investorID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &types.QueryInvestorWalletsResponse{Wallets: ret.WalletAddress}, nil
}

func (k Keeper) GetInvestorWallets(ctx sdk.Context, investorID string) (types.Investor, error) {
	store := ctx.KVStore(k.storeKey)
	investorStores := prefix.NewStore(store, types.KeyPrefix(types.InvestorToWalletsPrefix))
	b := investorStores.Get(types.KeyPrefix(investorID))
	if b == nil {
		return types.Investor{}, status.Errorf(codes.NotFound, "investor id %v", investorID)
	}
	var investor types.Investor
	k.cdc.MustUnmarshal(b, &investor)
	return investor, nil
}
