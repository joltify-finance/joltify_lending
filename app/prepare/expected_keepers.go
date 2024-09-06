package prepare

import (
	"time"

	pricestypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/prices/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bridgetypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/bridge/types"
	clobtypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/types"
	perpstypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/perpetuals/types"
)

// PrepareClobKeeper defines the expected CLOB keeper used for `PrepareProposal`.
type PrepareClobKeeper interface {
	GetOperations(ctx sdk.Context) *clobtypes.MsgProposedOperations
}

// PreparePerpetualsKeeper defines the expected Perpetuals keeper used for `PrepareProposal`.
type PreparePerpetualsKeeper interface {
	GetAddPremiumVotes(ctx sdk.Context) *perpstypes.MsgAddPremiumVotes
}

// PrepareBridgeKeeper defines the expected Bridge keeper used for `PrepareProposal`.
type PrepareBridgeKeeper interface {
	GetAcknowledgeBridges(ctx sdk.Context, blockTimestamp time.Time) *bridgetypes.MsgAcknowledgeBridges
}

// PreparePricesKeeper defines the expected Prices keeper used for `PrepareProposal`.
type PreparePricesKeeper interface {
	GetValidMarketPriceUpdates(ctx sdk.Context) *pricestypes.MsgUpdateMarketPrices
}
