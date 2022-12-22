package rest

import (
	"bytes"
	"fmt"
	"net/http"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

func registerTxRoutes(cliCtx client.Context, r *mux.Router) {
	// r.HandleFunc("/incentive/claim-cdp", postClaimHandlerFn(cliCtx, usdxMintingGenerator)).Methods("POST")
	r.HandleFunc("/incentive/claim-jolt", postClaimHandlerFn(cliCtx, joltGenerator)).Methods("POST")
	// r.HandleFunc("/incentive/claim-delegator", postClaimHandlerFn(cliCtx, delegatorGenerator)).Methods("POST")
	// r.HandleFunc("/incentive/claim-swap", postClaimHandlerFn(cliCtx, swapGenerator)).Methods("POST")
}

func usdxMintingGenerator(req PostClaimReq) (sdk.Msg, error) {
	if len(req.DenomsToClaim) != 1 {
		return nil, fmt.Errorf("must only claim %s denom for usdx minting rewards, got '%s", types2.USDXMintingRewardDenom, req.DenomsToClaim)
	}
	msg := types2.NewMsgClaimUSDXMintingReward(req.Sender.String(), req.DenomsToClaim[0].MultiplierName)
	return &msg, nil
}

func joltGenerator(req PostClaimReq) (sdk.Msg, error) {
	msg := types2.NewMsgClaimJoltReward(req.Sender.String(), req.DenomsToClaim)
	return &msg, nil
}

func delegatorGenerator(req PostClaimReq) (sdk.Msg, error) {
	msg := types2.NewMsgClaimDelegatorReward(req.Sender.String(), req.DenomsToClaim)
	return &msg, nil
}

func swapGenerator(req PostClaimReq) (sdk.Msg, error) {
	msg := types2.NewMsgClaimSwapReward(req.Sender.String(), req.DenomsToClaim)
	return &msg, nil
}

func postClaimHandlerFn(cliCtx client.Context, msgGenerator func(req PostClaimReq) (sdk.Msg, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody PostClaimReq
		if !rest.ReadRESTReq(w, r, cliCtx.LegacyAmino, &requestBody) {
			return
		}

		requestBody.BaseReq = requestBody.BaseReq.Sanitize()
		if !requestBody.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(requestBody.BaseReq.From)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		if !bytes.Equal(fromAddr, requestBody.Sender) {
			rest.WriteErrorResponse(w, http.StatusUnauthorized, fmt.Sprintf("expected: %s, got: %s", fromAddr, requestBody.Sender))
			return
		}

		msg, err := msgGenerator(requestBody)
		if rest.CheckBadRequestError(w, err) {
			return
		}
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(cliCtx, w, requestBody.BaseReq, msg)
	}
}
