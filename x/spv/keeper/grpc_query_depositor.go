package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"

	coserrors "cosmossdk.io/errors"
	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Depositor(goCtx context.Context, req *types.QueryDepositorRequest) (*types.QueryDepositorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	investor, err := sdk.AccAddressFromBech32(req.GetWalletAddress())
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address %v", req.GetWalletAddress())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	depositor, ok := k.GetDepositor(ctx, req.DepositPoolIndex, investor)

	if !ok {
		return nil, coserrors.Wrap(errorsmod.ErrNotFound, "depositor not found")
	}

	return &types.QueryDepositorResponse{Depositor: &depositor}, nil
}

func (k Keeper) DepositorHistory(goCtx context.Context, req *types.QueryDepositorHistoryRequest) (*types.QueryHistoryDepositorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	depositorStores := prefix.NewStore(store, types.KeyPrefix(types.PoolDepositorHistory))

	var depositorInfos []*types.DepositorInfo

	pageRes, err := query.Paginate(depositorStores, req.Pagination, func(key []byte, value []byte) error {
		var investor types.DepositorInfo
		if err := k.cdc.Unmarshal(value, &investor); err != nil {
			return err
		}

		if investor.DepositorAddress.String() != req.WalletAddress || investor.PoolIndex != req.DepositPoolIndex {
			return nil
		}

		depositorInfos = append(depositorInfos, &investor)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryHistoryDepositorResponse{Depositors: depositorInfos, Pagination: pageRes}, nil
}
