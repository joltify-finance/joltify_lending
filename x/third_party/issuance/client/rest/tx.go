package rest

import (
	"fmt"
	"net/http"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/issuance/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
)

func registerTxRoutes(cliCtx client.Context, r *mux.Router) {
	r.HandleFunc(fmt.Sprintf("/%s/issue", types2.ModuleName), postIssueTokensHandlerFn(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/redeem", types2.ModuleName), postRedeemTokensHandlerFn(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/block", types2.ModuleName), postBlockAddressHandlerFn(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/unblock", types2.ModuleName), postUnblockAddressHandlerFn(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/pause", types2.ModuleName), postPauseHandlerFn(cliCtx)).Methods("POST")
}

func postIssueTokensHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody PostIssueReq
		if !rest.ReadRESTReq(w, r, cliCtx.LegacyAmino, &requestBody) {
			return
		}

		baseReq := requestBody.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(requestBody.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types2.NewMsgIssueTokens(
			fromAddr.String(),
			requestBody.Tokens,
			requestBody.Receiver.String(),
		)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(cliCtx, w, baseReq, msg)
	}
}

func postRedeemTokensHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody PostRedeemReq
		if !rest.ReadRESTReq(w, r, cliCtx.LegacyAmino, &requestBody) {
			return
		}

		baseReq := requestBody.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(requestBody.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types2.NewMsgRedeemTokens(
			fromAddr.String(),
			requestBody.Tokens,
		)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(cliCtx, w, baseReq, msg)
	}
}

func postBlockAddressHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody PostBlockAddressReq
		if !rest.ReadRESTReq(w, r, cliCtx.LegacyAmino, &requestBody) {
			return
		}

		baseReq := requestBody.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(requestBody.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types2.NewMsgBlockAddress(
			fromAddr.String(),
			requestBody.Denom,
			requestBody.Address.String(),
		)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(cliCtx, w, baseReq, msg)
	}
}

func postUnblockAddressHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody PostUnblockAddressReq
		if !rest.ReadRESTReq(w, r, cliCtx.LegacyAmino, &requestBody) {
			return
		}

		baseReq := requestBody.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(requestBody.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types2.NewMsgUnblockAddress(
			fromAddr.String(),
			requestBody.Denom,
			requestBody.Address.String(),
		)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(cliCtx, w, baseReq, msg)
	}
}

func postPauseHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody PostPauseReq
		if !rest.ReadRESTReq(w, r, cliCtx.LegacyAmino, &requestBody) {
			return
		}

		baseReq := requestBody.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(requestBody.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types2.NewMsgSetPauseStatus(
			fromAddr.String(),
			requestBody.Denom,
			requestBody.Status,
		)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(cliCtx, w, baseReq, msg)
	}
}
