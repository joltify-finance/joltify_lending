package keeper

import (
	"context"

	coserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func addToList(previousList, newElements []string) []string {
	tempStore := make(map[string]string)
	for _, el := range previousList {
		tempStore[el] = el
	}

	for _, el := range newElements {
		tempStore[el] = el
	}

	var combinedList []string
	for _, val := range tempStore {
		combinedList = append(combinedList, val)
	}
	return combinedList
}

func (k msgServer) AddInvestors(goCtx context.Context, msg *types.MsgAddInvestors) (*types.MsgAddInvestorsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	spvAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address %v", msg.Creator)
	}

	pool, found := k.GetPools(ctx, msg.GetPoolIndex())
	if !found {
		return nil, coserrors.Wrapf(types.PoolNotFound, "pool not found with index %v", msg.GetPoolIndex())
	}
	if !pool.OwnerAddress.Equals(spvAddress) {
		return nil, coserrors.Wrap(types.Unauthorized, "unauthorized operations")
	}

	investorAddresses := make([]string, len(msg.GetInvetorID()))
	investorPoolInfo, found := k.GetInvestorToPool(ctx, msg.PoolIndex)
	if found {
		newList := addToList(investorPoolInfo.Investors, investorAddresses)
		investorPoolInfo.Investors = newList
		k.AddInvestorToPool(ctx, &investorPoolInfo)
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeAddInvestors,
				sdk.NewAttribute(types.AttributeCreator, msg.Creator),
			),
		)
		return &types.MsgAddInvestorsResponse{OperationResult: true}, nil
	}

	v := types.PoolWithInvestors{
		PoolIndex: msg.PoolIndex,
		Investors: investorAddresses,
	}

	k.AddInvestorToPool(ctx, &v)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAddInvestors,
			sdk.NewAttribute(types.AttributeCreator, msg.Creator),
		),
	)

	return &types.MsgAddInvestorsResponse{OperationResult: true}, nil
}
