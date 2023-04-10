package keeper

import (
	"context"

    "github.com/joltify-finance/joltify_lending/x/spv/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)


func (k msgServer) Liquidate(goCtx context.Context,  msg *types.MsgLiquidate) (*types.MsgLiquidateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    // TODO: Handling the message
    _ = ctx

	return &types.MsgLiquidateResponse{}, nil
}
