package keeper

import (
	"context"
	coserrors "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) WithdrawalPrincipal(goCtx context.Context, req *types.QuerywithdrawalPrincipalRequest) (*types.QuerywithdrawalPrincipalResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	investor, err := sdk.AccAddressFromBech32(req.WalletAddress)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address %v", req.WalletAddress)
	}

	depositor, found := k.GetDepositor(ctx, req.PoolIndex, investor)
	if !found {
		return nil, coserrors.Wrapf(types.ErrDepositorNotFound, "depositor not found for pool %v", req.PoolIndex)
	}

	return &types.QuerywithdrawalPrincipalResponse{Amount: depositor.WithdrawalAmount.String()}, nil
}
