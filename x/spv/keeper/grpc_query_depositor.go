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

func (k Keeper) Depositor(goCtx context.Context, req *types.QueryDepositorRequest) (*types.QueryDepositorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	investor, err := sdk.AccAddressFromBech32(req.GetWalletAddress())
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address %v", req.GetDepositPoolIndex())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	depositor, ok := k.GetDepositor(ctx, req.DepositPoolIndex, investor)

	if !ok {
		return nil, coserrors.Wrap(sdkerrors.ErrNotFound, "depositor not found")
	}

	return &types.QueryDepositorResponse{Depositor: &depositor}, nil
}
