package keeper

import (
	"context"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/cdp/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	keeper Keeper
}

// NewMsgServerImpl returns an implementation of the cdp MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types2.MsgServer {
	return &msgServer{keeper: keeper}
}

var _ types2.MsgServer = msgServer{}

func (k msgServer) CreateCDP(goCtx context.Context, msg *types2.MsgCreateCDP) (*types2.MsgCreateCDPResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	err = k.keeper.AddCdp(ctx, sender, msg.Collateral, msg.Principal, msg.CollateralType)
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

	id, _ := k.keeper.GetCdpID(ctx, sender, msg.CollateralType)
	return &types2.MsgCreateCDPResponse{CdpID: id}, nil
}

func (k msgServer) Deposit(goCtx context.Context, msg *types2.MsgDeposit) (*types2.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	depositor, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		return nil, err
	}

	err = k.keeper.DepositCollateral(ctx, owner, depositor, msg.Collateral, msg.CollateralType)
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

	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	depositor, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		return nil, err
	}

	err = k.keeper.WithdrawCollateral(ctx, owner, depositor, msg.Collateral, msg.CollateralType)
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

func (k msgServer) DrawDebt(goCtx context.Context, msg *types2.MsgDrawDebt) (*types2.MsgDrawDebtResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	err = k.keeper.AddPrincipal(ctx, sender, msg.CollateralType, msg.Principal)
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
	return &types2.MsgDrawDebtResponse{}, nil
}

func (k msgServer) RepayDebt(goCtx context.Context, msg *types2.MsgRepayDebt) (*types2.MsgRepayDebtResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	err = k.keeper.RepayPrincipal(ctx, sender, msg.CollateralType, msg.Payment)
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
	return &types2.MsgRepayDebtResponse{}, nil
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

	err = k.keeper.AttemptKeeperLiquidation(ctx, keeper, borrower, msg.CollateralType)
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
