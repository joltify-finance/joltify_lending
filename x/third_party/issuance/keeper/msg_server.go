package keeper

import (
	"context"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/issuance/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	keeper Keeper
}

// NewMsgServerImpl returns an implementation of the issuance MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types2.MsgServer {
	return &msgServer{keeper: keeper}
}

var _ types2.MsgServer = msgServer{}

func (k msgServer) IssueTokens(goCtx context.Context, msg *types2.MsgIssueTokens) (*types2.MsgIssueTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	receiver, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return nil, err
	}

	err = k.keeper.IssueTokens(ctx, msg.Tokens, sender, receiver)
	if err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types2.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
		),
	)
	return &types2.MsgIssueTokensResponse{}, nil
}

func (k msgServer) RedeemTokens(goCtx context.Context, msg *types2.MsgRedeemTokens) (*types2.MsgRedeemTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	err = k.keeper.RedeemTokens(ctx, msg.Tokens, sender)
	if err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types2.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
		),
	)
	return &types2.MsgRedeemTokensResponse{}, nil
}

func (k msgServer) BlockAddress(goCtx context.Context, msg *types2.MsgBlockAddress) (*types2.MsgBlockAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	blockedAddress, err := sdk.AccAddressFromBech32(msg.BlockedAddress)
	if err != nil {
		return nil, err
	}

	err = k.keeper.BlockAddress(ctx, msg.Denom, sender, blockedAddress)
	if err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types2.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
		),
	)
	return &types2.MsgBlockAddressResponse{}, nil
}

func (k msgServer) UnblockAddress(goCtx context.Context, msg *types2.MsgUnblockAddress) (*types2.MsgUnblockAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	blockedAddress, err := sdk.AccAddressFromBech32(msg.BlockedAddress)
	if err != nil {
		return nil, err
	}

	err = k.keeper.UnblockAddress(ctx, msg.Denom, sender, blockedAddress)
	if err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types2.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
		),
	)
	return &types2.MsgUnblockAddressResponse{}, nil
}

func (k msgServer) SetPauseStatus(goCtx context.Context, msg *types2.MsgSetPauseStatus) (*types2.MsgSetPauseStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	err = k.keeper.SetPauseStatus(ctx, sender, msg.Denom, msg.Status)
	if err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types2.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
		),
	)
	return &types2.MsgSetPauseStatusResponse{}, nil
}
