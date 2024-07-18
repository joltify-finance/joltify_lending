package oracle

import (
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/x/third_party/oracle/keeper"
)

type BlockHandler struct {
	k keeper.Keeper
}

func NewBlockHandler(k keeper.Keeper) *BlockHandler {
	return &BlockHandler{
		k: k,
	}
}

func (h *BlockHandler) BeginBlocker(ctx sdk.Context) {
	bandIBCParams := h.k.GetBandIBCParams(ctx)
	// Request oracle prices using band IBC in frequent intervals
	if bandIBCParams.BandIbcEnabled && ctx.BlockHeight()%bandIBCParams.IbcRequestInterval == 0 {
		h.RequestAllBandIBCRates(ctx)
	}

	if ctx.BlockHeight()%100000 == 0 {
		h.k.CleanupHistoricalPriceRecords(ctx)
	}

	// todo: default cleanup interval (1 day)
	if ctx.BlockHeight()%24*60*60 == 0 {
		h.k.CleanUpStaleBandIBCCalldataRecords(ctx)
	}
}

func (h *BlockHandler) RequestAllBandIBCRates(ctx sdk.Context) {
	bandIBCOracleRequests := h.k.GetAllBandIBCOracleRequests(ctx)

	if len(bandIBCOracleRequests) == 0 {
		return
	}

	for _, req := range bandIBCOracleRequests {
		err := h.k.RequestBandIBCOraclePrices(ctx, req)
		if err != nil {
			ctx.Logger().Error(err.Error())
		}
	}
}

func EndBlocker(ctx sdk.Context, block abci.RequestEndBlock) {
}
