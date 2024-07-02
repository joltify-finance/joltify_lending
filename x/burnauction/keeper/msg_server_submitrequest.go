package keeper

import (
	"context"

	"github.com/joltify-finance/joltify_lending/x/burnauction/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Submitrequest(goCtx context.Context, msg *types.MsgSubmitrequest) (*types.MsgSubmitrequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSubmitrequestResponse{}, nil
}
