package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

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
	ret, err := k.GetByWallet(ctx, req.Wallet)
	if err != nil {
		return &types.QueryByWalletResponse{}, status.Errorf(codes.NotFound, "wallet %v", req.Wallet)
	}
	return &types.QueryByWalletResponse{Investor: &ret}, nil
}

func (k Keeper) GetByWallet(rctx context.Context, wallet string) (types.Investor, error) {
	ctx := sdk.UnwrapSDKContext(rctx)
	gasBefore := ctx.GasMeter().GasConsumed()
	defer func() {
		gasAfter := ctx.GasMeter().GasConsumed()
		ctx.GasMeter().RefundGas(gasAfter-gasBefore, "GetByWallet")
	}()

	store := ctx.KVStore(k.storeKey)
	investorStores := prefix.NewStore(store, types.KeyPrefix(types.InvestorToWalletsPrefix))
	iterator := storetypes.KVStorePrefixIterator(investorStores, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var investor types.Investor
		k.cdc.MustUnmarshal(iterator.Value(), &investor)
		for _, el := range investor.WalletAddress {
			if el == wallet {
				return investor, nil
			}
		}
	}
	return types.Investor{}, status.Errorf(codes.NotFound, "wallet %v", wallet)
}
