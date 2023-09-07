package keeper

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier is the module level router for state queries
func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case types2.QueryGetParams:
			return queryGetParams(ctx, req, k, legacyQuerierCdc)
		case types2.QueryGetJoltRewards:
			return queryGetJoltRewards(ctx, req, k, legacyQuerierCdc)
		case types2.QueryGetRewardFactors:
			return queryGetRewardFactors(ctx, req, k, legacyQuerierCdc)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query endpoint", types2.ModuleName)
		}
	}
}

// query params in the store
func queryGetParams(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	// Get params
	params := k.GetParams(ctx)

	// Encode results
	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func queryGetJoltRewards(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types2.QueryRewardsParams
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	owner := len(params.Owner) > 0

	var hardClaims types2.JoltLiquidityProviderClaims
	switch {
	case owner:
		hardClaim, foundHardClaim := k.GetJoltLiquidityProviderClaim(ctx, params.Owner)
		if foundHardClaim {
			hardClaims = append(hardClaims, hardClaim)
		}
	default:
		hardClaims = k.GetAllJoltLiquidityProviderClaims(ctx)
	}

	var paginatedJoltClaims types2.JoltLiquidityProviderClaims
	startH, endH := client.Paginate(len(hardClaims), params.Page, params.Limit, 100)
	if startH < 0 || endH < 0 {
		paginatedJoltClaims = types2.JoltLiquidityProviderClaims{}
	} else {
		paginatedJoltClaims = hardClaims[startH:endH]
	}

	if !params.Unsynchronized {
		for i, claim := range paginatedJoltClaims {
			paginatedJoltClaims[i] = k.SimulateJoltSynchronization(ctx, claim)
		}
	}

	// Marshal Hard claims
	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, paginatedJoltClaims)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func queryGetRewardFactors(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var supplyFactors types2.MultiRewardIndexes
	k.IterateJoltSupplyRewardIndexes(ctx, func(denom string, indexes types2.RewardIndexes) (stop bool) {
		supplyFactors = supplyFactors.With(denom, indexes)
		return false
	})

	var borrowFactors types2.MultiRewardIndexes
	k.IterateJoltBorrowRewardIndexes(ctx, func(denom string, indexes types2.RewardIndexes) (stop bool) {
		borrowFactors = borrowFactors.With(denom, indexes)
		return false
	})

	response := types2.NewQueryGetRewardFactorsResponse(
		supplyFactors,
		borrowFactors,
	)

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, response)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}
