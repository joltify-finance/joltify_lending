package rest

import (
	"fmt"
	"net/http"
	"strings"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

func registerQueryRoutes(cliCtx client.Context, r *mux.Router) {
	r.HandleFunc(fmt.Sprintf("/%s/parameters", types2.ModuleName), queryParamsHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/unsynced-deposits", types2.ModuleName), queryUnsyncedDepositsHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/deposits", types2.ModuleName), queryDepositsHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/total-deposited", types2.ModuleName), queryTotalDepositedHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/accounts", types2.ModuleName), queryModAccountsHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/unsynced-borrows", types2.ModuleName), queryUnsyncedBorrowsHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/borrows", types2.ModuleName), queryBorrowsHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/total-borrowed", types2.ModuleName), queryTotalBorrowedHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/interest-rate", types2.ModuleName), queryInterestRateHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/reserves", types2.ModuleName), queryReservesHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/interest-factors", types2.ModuleName), queryInterestFactorsHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/liquidate", types2.ModuleName), queryLiquidateHandlerFn(cliCtx)).Methods("GET")
}

func queryParamsHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types2.QuerierRoute, types2.QueryGetParams)

		res, height, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryDepositsHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, page, limit, err := rest.ParseHTTPArgsWithLimit(r, 0)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		// Parse the query height
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		var denom string
		var owner sdk.AccAddress

		if x := r.URL.Query().Get(RestDenom); len(x) != 0 {
			denom = strings.TrimSpace(x)
		}

		if x := r.URL.Query().Get(RestOwner); len(x) != 0 {
			ownerStr := strings.ToLower(strings.TrimSpace(x))
			owner, err = sdk.AccAddressFromBech32(ownerStr)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("cannot parse address from deposit owner %s", ownerStr))
				return
			}
		}

		params := types2.NewQueryDepositsParams(page, limit, denom, owner)

		bz, err := cliCtx.LegacyAmino.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types2.ModuleName, types2.QueryGetDeposits)
		res, height, err := cliCtx.QueryWithData(route, bz)
		cliCtx = cliCtx.WithHeight(height)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryUnsyncedDepositsHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, page, limit, err := rest.ParseHTTPArgsWithLimit(r, 0)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		// Parse the query height
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		var denom string
		var owner sdk.AccAddress

		if x := r.URL.Query().Get(RestDenom); len(x) != 0 {
			denom = strings.TrimSpace(x)
		}

		if x := r.URL.Query().Get(RestOwner); len(x) != 0 {
			ownerStr := strings.ToLower(strings.TrimSpace(x))
			owner, err = sdk.AccAddressFromBech32(ownerStr)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("cannot parse address from deposit owner %s", ownerStr))
				return
			}
		}

		params := types2.NewQueryUnsyncedDepositsParams(page, limit, denom, owner)

		bz, err := cliCtx.LegacyAmino.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types2.ModuleName, types2.QueryGetUnsyncedDeposits)
		res, height, err := cliCtx.QueryWithData(route, bz)
		cliCtx = cliCtx.WithHeight(height)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryTotalDepositedHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := rest.ParseHTTPArgsWithLimit(r, 0)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		// Parse the query height
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		var denom string

		if x := r.URL.Query().Get(RestDenom); len(x) != 0 {
			denom = strings.TrimSpace(x)
		}

		params := types2.NewQueryTotalDepositedParams(denom)

		bz, err := cliCtx.LegacyAmino.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types2.ModuleName, types2.QueryGetTotalDeposited)
		res, height, err := cliCtx.QueryWithData(route, bz)
		cliCtx = cliCtx.WithHeight(height)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryBorrowsHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, page, limit, err := rest.ParseHTTPArgsWithLimit(r, 0)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		// Parse the query height
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		var denom string
		var owner sdk.AccAddress

		if x := r.URL.Query().Get(RestDenom); len(x) != 0 {
			denom = strings.TrimSpace(x)
		}

		if x := r.URL.Query().Get(RestOwner); len(x) != 0 {
			ownerStr := strings.ToLower(strings.TrimSpace(x))
			owner, err = sdk.AccAddressFromBech32(ownerStr)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("cannot parse address from borrow owner %s", ownerStr))
				return
			}
		}

		params := types2.NewQueryBorrowsParams(page, limit, owner, denom)

		bz, err := cliCtx.LegacyAmino.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types2.ModuleName, types2.QueryGetBorrows)
		res, height, err := cliCtx.QueryWithData(route, bz)
		cliCtx = cliCtx.WithHeight(height)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryUnsyncedBorrowsHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, page, limit, err := rest.ParseHTTPArgsWithLimit(r, 0)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		// Parse the query height
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		var denom string
		var owner sdk.AccAddress

		if x := r.URL.Query().Get(RestDenom); len(x) != 0 {
			denom = strings.TrimSpace(x)
		}

		if x := r.URL.Query().Get(RestOwner); len(x) != 0 {
			ownerStr := strings.ToLower(strings.TrimSpace(x))
			owner, err = sdk.AccAddressFromBech32(ownerStr)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("cannot parse address from borrow owner %s", ownerStr))
				return
			}
		}

		params := types2.NewQueryUnsyncedBorrowsParams(page, limit, owner, denom)

		bz, err := cliCtx.LegacyAmino.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types2.ModuleName, types2.QueryGetUnsyncedBorrows)
		res, height, err := cliCtx.QueryWithData(route, bz)
		cliCtx = cliCtx.WithHeight(height)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryTotalBorrowedHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := rest.ParseHTTPArgsWithLimit(r, 0)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		// Parse the query height
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		var denom string

		if x := r.URL.Query().Get(RestDenom); len(x) != 0 {
			denom = strings.TrimSpace(x)
		}

		params := types2.NewQueryTotalBorrowedParams(denom)

		bz, err := cliCtx.LegacyAmino.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types2.ModuleName, types2.QueryGetTotalBorrowed)
		res, height, err := cliCtx.QueryWithData(route, bz)
		cliCtx = cliCtx.WithHeight(height)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryInterestRateHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := rest.ParseHTTPArgsWithLimit(r, 0)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		// Parse the query height
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		var denom string

		if x := r.URL.Query().Get(RestDenom); len(x) != 0 {
			denom = strings.TrimSpace(x)
		}

		params := types2.NewQueryInterestRateParams(denom)

		bz, err := cliCtx.LegacyAmino.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types2.ModuleName, types2.QueryGetInterestRate)
		res, height, err := cliCtx.QueryWithData(route, bz)
		cliCtx = cliCtx.WithHeight(height)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryModAccountsHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, page, limit, err := rest.ParseHTTPArgsWithLimit(r, 0)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		// Parse the query height
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		var name string

		if x := r.URL.Query().Get(RestName); len(x) != 0 {
			name = strings.TrimSpace(x)
		}

		params := types2.NewQueryAccountParams(page, limit, name)

		bz, err := cliCtx.LegacyAmino.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types2.ModuleName, types2.QueryGetModuleAccounts)
		res, height, err := cliCtx.QueryWithData(route, bz)
		cliCtx = cliCtx.WithHeight(height)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryReservesHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := rest.ParseHTTPArgsWithLimit(r, 0)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		// Parse the query height
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		var denom string

		if x := r.URL.Query().Get(RestDenom); len(x) != 0 {
			denom = strings.TrimSpace(x)
		}

		params := types2.NewQueryReservesParams(denom)

		bz, err := cliCtx.LegacyAmino.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types2.ModuleName, types2.QueryGetReserves)
		res, height, err := cliCtx.QueryWithData(route, bz)
		cliCtx = cliCtx.WithHeight(height)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryInterestFactorsHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := rest.ParseHTTPArgsWithLimit(r, 0)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		// Parse the query height
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		var denom string

		if x := r.URL.Query().Get(RestDenom); len(x) != 0 {
			denom = strings.TrimSpace(x)
		}

		params := types2.NewQueryInterestFactorsParams(denom)

		bz, err := cliCtx.LegacyAmino.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types2.ModuleName, types2.QueryGetInterestFactors)
		res, height, err := cliCtx.QueryWithData(route, bz)
		cliCtx = cliCtx.WithHeight(height)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryLiquidateHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, page, limite, err := rest.ParseHTTPArgsWithLimit(r, 0)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		// Parse the query height
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		var borrowStr string
		if x := r.URL.Query().Get(RestBorrow); len(x) != 0 {
			borrowStr = strings.ToLower(strings.TrimSpace(x))
		}

		req := types2.QueryLiquidate{Page: page, Limit: limite, Borrow: borrowStr}

		bz, err := cliCtx.LegacyAmino.MarshalJSON(req)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		route := fmt.Sprintf("custom/%s/%s", types2.ModuleName, types2.QueryGetLiquidate)
		res, height, err := cliCtx.QueryWithData(route, bz)
		cliCtx = cliCtx.WithHeight(height)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, res)
	}
}
