package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier is the module level router for state queries
func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case types.QueryGetParams:
			return queryGetParams(ctx, req, k, legacyQuerierCdc)
		case types.QueryGetJoltRewards:
			return queryGetJoltRewards(ctx, req, k, legacyQuerierCdc)
		case types.QueryGetRewardFactors:
			return queryGetRewardFactors(ctx, req, k, legacyQuerierCdc)
		case types.QueryGetSwapRewards:
			return queryGetSwapRewards(ctx, req, k, legacyQuerierCdc)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query endpoint", types.ModuleName)
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
	var params types.QueryRewardsParams
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	owner := len(params.Owner) > 0

	var hardClaims types.JoltLiquidityProviderClaims
	switch {
	case owner:
		hardClaim, foundHardClaim := k.GetJoltLiquidityProviderClaim(ctx, params.Owner)
		if foundHardClaim {
			hardClaims = append(hardClaims, hardClaim)
		}
	default:
		hardClaims = k.GetAllJoltLiquidityProviderClaims(ctx)
	}

	var paginatedJoltClaims types.JoltLiquidityProviderClaims
	startH, endH := client.Paginate(len(hardClaims), params.Page, params.Limit, 100)
	if startH < 0 || endH < 0 {
		paginatedJoltClaims = types.JoltLiquidityProviderClaims{}
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

func queryGetSwapRewards(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryRewardsParams
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	owner := len(params.Owner) > 0

	var claims types.SwapClaims
	switch {
	case owner:
		claim, found := k.GetSwapClaim(ctx, params.Owner)
		if found {
			claims = append(claims, claim)
		}
	default:
		claims = k.GetAllSwapClaims(ctx)
	}

	var paginatedClaims types.SwapClaims
	startH, endH := client.Paginate(len(claims), params.Page, params.Limit, 100)
	if startH < 0 || endH < 0 {
		paginatedClaims = types.SwapClaims{}
	} else {
		paginatedClaims = claims[startH:endH]
	}

	if !params.Unsynchronized {
		for i, claim := range paginatedClaims {
			syncedClaim, found := k.GetSynchronizedSwapClaim(ctx, claim.Owner)
			if !found {
				panic("previously found claim should still be found")
			}
			paginatedClaims[i] = syncedClaim
		}
	}

	// Marshal claims
	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, paginatedClaims)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func queryGetRewardFactors(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var supplyFactors types.MultiRewardIndexes
	k.IterateJoltSupplyRewardIndexes(ctx, func(denom string, indexes types.RewardIndexes) (stop bool) {
		supplyFactors = supplyFactors.With(denom, indexes)
		return false
	})

	var borrowFactors types.MultiRewardIndexes
	k.IterateJoltBorrowRewardIndexes(ctx, func(denom string, indexes types.RewardIndexes) (stop bool) {
		borrowFactors = borrowFactors.With(denom, indexes)
		return false
	})

	var swapFactors types.MultiRewardIndexes
	k.IterateSwapRewardIndexes(ctx, func(poolID string, indexes types.RewardIndexes) (stop bool) {
		swapFactors = swapFactors.With(poolID, indexes)
		return false
	})

	response := types.NewQueryGetRewardFactorsResponse(
		supplyFactors,
		borrowFactors,
		swapFactors,
	)

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, response)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}
