package keeper

import (
	"context"

	coserrors "cosmossdk.io/errors"
	types2 "cosmossdk.io/store/types"
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

func (k msgServer) AddInvestors(goCtx context.Context, msg *types.MsgAddInvestors) (*types.MsgAddInvestorsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	ctx = ctx.WithGasMeter(types2.NewInfiniteGasMeter())

	spvAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address %v", msg.Creator)
	}

	pool, found := k.GetPools(ctx, msg.PoolIndex)
	if !found {
		return nil, coserrors.Wrapf(types.ErrPoolNotFound, "pool not found with index %v", msg.PoolIndex)
	}
	if !pool.OwnerAddress.Equals(spvAddress) {
		return nil, coserrors.Wrap(types.ErrUnauthorized, "unauthorized operations")
	}

	// check if the investor id exist in kyc module
	for _, el := range msg.InvestorID {
		_, err := k.kycKeeper.GetInvestorWallets(ctx, el)
		if err != nil {
			return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "the given investor id %v cannot be found", el)
		}
	}

	targetProject, ok := k.kycKeeper.GetProject(ctx, pool.LinkedProject)
	if !ok {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "the given project %v cannot be found", pool.LinkedProject)
	}

	poolType := types.Senior
	if pool.PoolType == types.PoolInfo_SENIOR {
		poolType = types.Junior
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
				Investors: msg.InvestorID,
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
