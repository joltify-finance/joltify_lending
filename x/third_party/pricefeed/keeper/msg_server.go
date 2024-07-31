package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/types"
)

type msgServer struct {
	keeper Keeper
}

// NewMsgServerImpl returns an implementation of the pricefeed MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types2.MsgServer {
	return &msgServer{keeper: keeper}
}

var _ types2.MsgServer = msgServer{}

func (k msgServer) PostPrice(goCtx context.Context, msg *types2.MsgPostPrice) (*types2.MsgPostPriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	_, err = k.keeper.GetOracle(ctx, msg.MarketID, from)
	if err != nil {
		return nil, err
	}

	_, err = k.keeper.SetPrice(ctx, from, msg.MarketID, msg.Price, msg.Expiry)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types2.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.From),
		),
	)

	return &types2.MsgPostPriceResponse{}, nil
}
