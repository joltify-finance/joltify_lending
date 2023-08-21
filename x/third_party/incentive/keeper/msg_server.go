package keeper

import (
	"context"

	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	keeper Keeper
}

// NewMsgServerImpl returns an implementation of the incentive MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) ClaimJoltReward(goCtx context.Context, msg *types.MsgClaimJoltReward) (*types.MsgClaimJoltRewardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	for _, selection := range msg.DenomsToClaim {
		err := k.keeper.ClaimJoltReward(ctx, sender, sender, selection.Denom, selection.MultiplierName)
		if err != nil {
			return nil, err
		}
	}

	return &types.MsgClaimJoltRewardResponse{}, nil
}

func (k msgServer) ClaimSwapReward(goCtx context.Context, msg *types.MsgClaimSwapReward) (*types.MsgClaimSwapRewardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	for _, selection := range msg.DenomsToClaim {
		err := k.keeper.ClaimSwapReward(ctx, sender, sender, selection.Denom, selection.MultiplierName)
		if err != nil {
			return nil, err
		}
	}

	return &types.MsgClaimSwapRewardResponse{}, nil
}
