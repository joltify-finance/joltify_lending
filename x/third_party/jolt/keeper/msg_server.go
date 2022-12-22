package keeper

import (
	"context"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	keeper Keeper
}

// NewMsgServerImpl returns an implementation of the jolt MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types2.MsgServer {
	return &msgServer{keeper: keeper}
}

var _ types2.MsgServer = msgServer{}

func (k msgServer) Deposit(goCtx context.Context, msg *types2.MsgDeposit) (*types2.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	depositor, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		return nil, err
	}

	err = k.keeper.Deposit(ctx, depositor, msg.Amount)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types2.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Depositor),
		),
	)
	return &types2.MsgDepositResponse{}, nil
}

func (k msgServer) Withdraw(goCtx context.Context, msg *types2.MsgWithdraw) (*types2.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	depositor, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		return nil, err
	}

	err = k.keeper.Withdraw(ctx, depositor, msg.Amount)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types2.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Depositor),
		),
	)
	return &types2.MsgWithdrawResponse{}, nil
}

func (k msgServer) Borrow(goCtx context.Context, msg *types2.MsgBorrow) (*types2.MsgBorrowResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	borrower, err := sdk.AccAddressFromBech32(msg.Borrower)
	if err != nil {
		return nil, err
	}

	err = k.keeper.Borrow(ctx, borrower, msg.Amount)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types2.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Borrower),
		),
	)
	return &types2.MsgBorrowResponse{}, nil
}

func (k msgServer) Repay(goCtx context.Context, msg *types2.MsgRepay) (*types2.MsgRepayResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	err = k.keeper.Repay(ctx, sender, owner, msg.Amount)
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
	return &types2.MsgRepayResponse{}, nil
}

func (k msgServer) Liquidate(goCtx context.Context, msg *types2.MsgLiquidate) (*types2.MsgLiquidateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	keeper, err := sdk.AccAddressFromBech32(msg.Keeper)
	if err != nil {
		return nil, err
	}

	borrower, err := sdk.AccAddressFromBech32(msg.Borrower)
	if err != nil {
		return nil, err
	}

	err = k.keeper.AttemptKeeperLiquidation(ctx, keeper, borrower)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types2.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Keeper),
		),
	)
	return &types2.MsgLiquidateResponse{}, nil
}
