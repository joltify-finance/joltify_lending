package keeper

import (
	"context"
	coserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func addToList(previousList, newElements []string) []string {
	combinedList := make([]string, 0, len(previousList)+len(newElements))
	exists := make(map[string]bool)

	for _, el := range previousList {
		if !exists[el] {
			exists[el] = true
			combinedList = append(combinedList, el)
		}
	}

	for _, el := range newElements {
		if !exists[el] {
			exists[el] = true
			combinedList = append(combinedList, el)
		}
	}

	return combinedList
}

func addAddrToList(previousList []sdk.AccAddress, newElement sdk.AccAddress) []sdk.AccAddress {
	combinedList := make([]sdk.AccAddress, 0, len(previousList)+1)
	exists := make(map[string]bool)

	for _, el := range previousList {
		if !exists[el.String()] {
			exists[el.String()] = true
			combinedList = append(combinedList, el)
		}
	}

	if !exists[newElement.String()] {
		combinedList = append(combinedList, newElement)
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
		return nil, coserrors.Wrapf(types.ErrPoolNotFound, "pool not found with index %v", msg.GetPoolIndex())
	}
	if !pool.OwnerAddress.Equals(spvAddress) {
		return nil, coserrors.Wrap(types.ErrUnauthorized, "unauthorized operations")
	}

	allProjects := k.kycKeeper.GetProjects(ctx)
	targetProject := allProjects[pool.LinkedProject-1]

	poolType := "senior"
	if pool.PoolType == types.PoolInfo_SENIOR {
		poolType = "junior"
	}

	indexHashreq := crypto.Keccak256Hash([]byte(targetProject.BasicInfo.ProjectName), spvAddress.Bytes(), []byte(poolType))

	_, found = k.GetPools(ctx, indexHashreq.Hex())
	if !found {
		panic("second pool cannot be found")
	}
	allPools := []string{msg.PoolIndex, indexHashreq.Hex()}
	for _, poolIndex := range allPools {

		investorPoolInfo, found := k.GetInvestorToPool(ctx, poolIndex)
		if found {
			newList := addToList(investorPoolInfo.Investors, msg.InvestorID)
			investorPoolInfo.Investors = newList
			k.AddInvestorToPool(ctx, &investorPoolInfo)

		} else {
			v := types.PoolWithInvestors{
				PoolIndex: poolIndex,
				Investors: msg.GetInvestorID(),
			}
			k.AddInvestorToPool(ctx, &v)
		}
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAddInvestors,
			sdk.NewAttribute(types.AttributeCreator, msg.Creator),
		),
	)

	return &types.MsgAddInvestorsResponse{OperationResult: true}, nil
}
