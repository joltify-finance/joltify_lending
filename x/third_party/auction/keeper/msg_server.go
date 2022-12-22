package keeper

import (
	"context"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/auction/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	keeper Keeper
}

// NewMsgServerImpl returns an implementation of the auction MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types2.MsgServer {
	return &msgServer{keeper: keeper}
}

var _ types2.MsgServer = msgServer{}

func (k msgServer) PlaceBid(goCtx context.Context, msg *types2.MsgPlaceBid) (*types2.MsgPlaceBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	bidder, err := sdk.AccAddressFromBech32(msg.Bidder)
	if err != nil {
		return nil, err
	}

	err = k.keeper.PlaceBid(ctx, msg.AuctionId, bidder, msg.Amount)
	if err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types2.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Bidder),
		),
	)
	return &types2.MsgPlaceBidResponse{}, nil
}
