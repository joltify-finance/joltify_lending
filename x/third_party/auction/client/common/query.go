package common

import (
	"fmt"
	"strings"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/auction/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
)

const (
	defaultPage  = 1
	defaultLimit = 100
)

// QueryAuctionByID returns an auction from state if present or falls back to searching old blocks
func QueryAuctionByID(cliCtx client.Context, cdc *codec.Codec, queryRoute string, auctionID uint64) (types2.Auction, int64, error) {
	bz, err := cliCtx.LegacyAmino.MarshalJSON(types2.NewQueryAuctionParams(auctionID))
	if err != nil {
		return nil, 0, err
	}

	res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", queryRoute, types2.QueryGetAuction), bz)

	if err == nil {
		var auction types2.Auction
		cliCtx.LegacyAmino.MustUnmarshalJSON(res, &auction)

		return auction, height, nil
	}

	if err != nil && !strings.Contains(err.Error(), "auction not found") {
		return nil, 0, err
	}

	res, height, err = cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", queryRoute, types2.QueryNextAuctionID), nil)
	if err != nil {
		return nil, 0, err
	}

	var nextAuctionID uint64
	cliCtx.LegacyAmino.MustUnmarshalJSON(res, &nextAuctionID)

	if auctionID >= nextAuctionID {
		return nil, 0, sdkerrors.Wrapf(types2.ErrAuctionNotFound, "%d", auctionID)
	}

	events := []string{
		fmt.Sprintf("%s.%s='%s'", sdk.EventTypeMessage, sdk.AttributeKeyAction, "place_bid"),
		fmt.Sprintf("%s.%s='%s'", types2.EventTypeAuctionBid, types2.AttributeKeyAuctionID, []byte(fmt.Sprintf("%d", auctionID))),
	}

	// if the auction is closed, query for previous bid transactions
	// note, will only fetch a maximum of 100 bids, so if an auction had more than that this
	// query may fail to retreive the final state of the auction
	searchResult, err := authtx.QueryTxsByEvents(cliCtx, events, defaultPage, defaultLimit, "")
	if err != nil {
		return nil, 0, err
	}

	maxHeight := int64(0)
	found := false

	for _, info := range searchResult.Txs {
		for _, msg := range info.GetTx().GetMsgs() {
			_, ok := msg.(*types2.MsgPlaceBid)
			if ok {
				found = true
				if info.Height > maxHeight {
					maxHeight = info.Height
				}
			}
		}
	}

	if !found {
		return nil, 0, sdkerrors.Wrapf(types2.ErrAuctionNotFound, "%d", auctionID)
	}

	queryCLIContext := cliCtx.WithHeight(maxHeight)
	res, height, err = queryCLIContext.QueryWithData(fmt.Sprintf("custom/%s/%s", queryRoute, types2.QueryGetAuction), bz)
	if err != nil {
		return nil, 0, err
	}

	// Decode and print results
	var auction types2.Auction
	cliCtx.LegacyAmino.MustUnmarshalJSON(res, &auction)
	return auction, height, nil
}
