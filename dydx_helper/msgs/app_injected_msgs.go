package msgs

import (
	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/testutil/dydx/testutil/constants"
	bridgetypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/bridge/types"
	clobtypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/types"
	perptypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/perpetuals/types"
	pricestypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/prices/types"
)

// AppInjectedMsgSamples are msgs that are injected into the block by the proposing validator.
// These messages are reserved for proposing validator's use only.
var AppInjectedMsgSamples = map[string]sdk.Msg{
	// bridge
	"/joltify.third_party.dydxprotocol.bridge.MsgAcknowledgeBridges": &bridgetypes.MsgAcknowledgeBridges{
		Events: []bridgetypes.BridgeEvent{
			{
				Id: 0,
				Coin: sdk.NewCoin(
					"bridge-token",
					sdkmath.NewIntFromUint64(1234),
				),
				Address: constants.Alice_Num0.Owner,
			},
		},
	},
	"/joltify.third_party.dydxprotocol.bridge.MsgAcknowledgeBridgesResponse": nil,

	// clob
	"/joltify.third_party.dydxprotocol.clob.MsgProposedOperations": &clobtypes.MsgProposedOperations{
		OperationsQueue: make([]clobtypes.OperationRaw, 0),
	},
	"/joltify.third_party.dydxprotocol.clob.MsgProposedOperationsResponse": nil,

	// perpetuals
	"/joltify.third_party.dydxprotocol.perpetuals.MsgAddPremiumVotes": &perptypes.MsgAddPremiumVotes{
		Votes: []perptypes.FundingPremium{
			{PerpetualId: 0, PremiumPpm: 1_000},
		},
	},
	"/joltify.third_party.dydxprotocol.perpetuals.MsgAddPremiumVotesResponse": nil,

	// prices
	"/joltify.third_party.dydxprotocol.prices.MsgUpdateMarketPrices": &pricestypes.MsgUpdateMarketPrices{
		MarketPriceUpdates: []*pricestypes.MsgUpdateMarketPrices_MarketPrice{
			pricestypes.NewMarketPriceUpdate(constants.MarketId0, 123_000),
		},
	},
	"/joltify.third_party.dydxprotocol.prices.MsgUpdateMarketPricesResponse": nil,
}
