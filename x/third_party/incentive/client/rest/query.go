package rest

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

func registerQueryRoutes(cliCtx client.Context, r *mux.Router) {
	r.HandleFunc(fmt.Sprintf("/%s/rewards", types2.ModuleName), queryRewardsHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/parameters", types2.ModuleName), queryParamsHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/reward-factors", types2.ModuleName), queryRewardFactorsHandlerFn(cliCtx)).Methods("GET")
}

func queryRewardsHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, page, limit, err := rest.ParseHTTPArgsWithLimit(r, 0)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		var owner sdk.AccAddress
		if x := r.URL.Query().Get(types2.RestClaimOwner); len(x) != 0 {
			ownerStr := strings.ToLower(strings.TrimSpace(x))
			owner, err = sdk.AccAddressFromBech32(ownerStr)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("cannot parse address from claim owner %s", ownerStr))
				return
			}
		}

		var rewardType string
		if x := r.URL.Query().Get(types2.RestClaimType); len(x) != 0 {
			rewardType = strings.ToLower(strings.TrimSpace(x))
		}

		var unsynced bool
		if x := r.URL.Query().Get(types2.RestUnsynced); len(x) != 0 {
			unsyncedStr := strings.ToLower(strings.TrimSpace(x))
			unsynced, err = strconv.ParseBool(unsyncedStr)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("cannot parse bool from unsynced flag %s", unsyncedStr))
				return
			}
		}

		params := types2.NewQueryRewardsParams(page, limit, owner, unsynced)
		switch strings.ToLower(rewardType) {
		case "jolt":
			executeJoltRewardsQuery(w, cliCtx, params)
		case "usdx_minting":
			executeUSDXMintingRewardsQuery(w, cliCtx, params)
		case "delegator":
			executeDelegatorRewardsQuery(w, cliCtx, params)
		case "swap":
			executeSwapRewardsQuery(w, cliCtx, params)
		default:
			executeAllRewardQueries(w, cliCtx, params)
		}
	}
}

func queryParamsHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		route := fmt.Sprintf("custom/%s/parameters", types2.QuerierRoute)

		res, height, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryRewardFactorsHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types2.ModuleName, types2.QueryGetRewardFactors)

		res, height, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func executeJoltRewardsQuery(w http.ResponseWriter, cliCtx client.Context, params types2.QueryRewardsParams) {
	bz, err := cliCtx.LegacyAmino.MarshalJSON(params)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("failed to marshal query params: %s", err))
		return
	}

	res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/incentive/%s", types2.QueryGetJoltRewards), bz)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	cliCtx = cliCtx.WithHeight(height)
	rest.PostProcessResponse(w, cliCtx, res)
}

func executeUSDXMintingRewardsQuery(w http.ResponseWriter, cliCtx client.Context, params types2.QueryRewardsParams) {
	bz, err := cliCtx.LegacyAmino.MarshalJSON(params)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("failed to marshal query params: %s", err))
		return
	}

	res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/incentive/%s", types2.QueryGetUSDXMintingRewards), bz)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	cliCtx = cliCtx.WithHeight(height)
	rest.PostProcessResponse(w, cliCtx, res)
}

func executeDelegatorRewardsQuery(w http.ResponseWriter, cliCtx client.Context, params types2.QueryRewardsParams) {
	bz, err := cliCtx.LegacyAmino.MarshalJSON(params)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("failed to marshal query params: %s", err))
		return
	}

	res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/incentive/%s", types2.QueryGetDelegatorRewards), bz)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	cliCtx = cliCtx.WithHeight(height)
	rest.PostProcessResponse(w, cliCtx, res)
}

func executeSwapRewardsQuery(w http.ResponseWriter, cliCtx client.Context, params types2.QueryRewardsParams) {
	bz, err := cliCtx.LegacyAmino.MarshalJSON(params)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("failed to marshal query params: %s", err))
		return
	}

	res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/incentive/%s", types2.QueryGetSwapRewards), bz)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	cliCtx = cliCtx.WithHeight(height)
	rest.PostProcessResponse(w, cliCtx, res)
}

func executeAllRewardQueries(w http.ResponseWriter, cliCtx client.Context, params types2.QueryRewardsParams) {
	paramsBz, err := cliCtx.LegacyAmino.MarshalJSON(params)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("failed to marshal query params: %s", err))
		return
	}
	joltRes, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/incentive/%s", types2.QueryGetJoltRewards), paramsBz)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	var joltClaims types2.JoltLiquidityProviderClaims
	cliCtx.LegacyAmino.MustUnmarshalJSON(joltRes, &joltClaims)

	usdxMintingRes, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/incentive/%s", types2.QueryGetUSDXMintingRewards), paramsBz)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	var usdxMintingClaims types2.USDXMintingClaims
	cliCtx.LegacyAmino.MustUnmarshalJSON(usdxMintingRes, &usdxMintingClaims)

	delegatorRes, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/incentive/%s", types2.QueryGetDelegatorRewards), paramsBz)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	var delegatorClaims types2.DelegatorClaims
	cliCtx.LegacyAmino.MustUnmarshalJSON(delegatorRes, &delegatorClaims)

	swapRes, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/incentive/%s", types2.QueryGetSwapRewards), paramsBz)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	var swapClaims types2.SwapClaims
	cliCtx.LegacyAmino.MustUnmarshalJSON(swapRes, &swapClaims)

	cliCtx = cliCtx.WithHeight(height)

	type rewardResult struct {
		JoltClaims        types2.JoltLiquidityProviderClaims `json:"jolt_claims" yaml:"jolt_claims"`
		UsdxMintingClaims types2.USDXMintingClaims           `json:"usdx_minting_claims" yaml:"usdx_minting_claims"`
		DelegatorClaims   types2.DelegatorClaims             `json:"delegator_claims" yaml:"delegator_claims"`
		SwapClaims        types2.SwapClaims                  `json:"swap_claims" yaml:"swap_claims"`
	}

	res := rewardResult{
		JoltClaims:        joltClaims,
		UsdxMintingClaims: usdxMintingClaims,
		DelegatorClaims:   delegatorClaims,
		SwapClaims:        swapClaims,
	}

	resBz, err := cliCtx.LegacyAmino.MarshalJSON(res)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("failed to marshal result: %s", err))
		return
	}

	rest.PostProcessResponse(w, cliCtx, resBz)
}
