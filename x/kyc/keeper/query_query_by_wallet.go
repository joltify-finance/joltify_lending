package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/kyc/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QueryByWallet(goCtx context.Context, req *types.QueryByWalletRequest) (*types.QueryByWalletResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	_, err := sdk.AccAddressFromBech32(req.Wallet)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request wallet address")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	investorStores := prefix.NewStore(store, types.KeyPrefix(types.InvestorToWalletsPrefix))

	iterator := sdk.KVStorePrefixIterator(investorStores, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var investor types.Investor
		k.cdc.MustUnmarshal(iterator.Value(), &investor)
		for _, el := range investor.WalletAddress {
			if el == req.Wallet {
				return &types.QueryByWalletResponse{Investor: &investor}, nil
			}
		}
	}
	return &types.QueryByWalletResponse{Investor: nil}, nil
}
