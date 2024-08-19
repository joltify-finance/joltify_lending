package keeper

import (
	"fmt"
	"sync/atomic"
	"time"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	pricefeedservertypes "github.com/joltify-finance/joltify_lending/daemons/server/types/pricefeed"
	"github.com/joltify-finance/joltify_lending/dydx_helper/indexer/indexer_manager"
	"github.com/joltify-finance/joltify_lending/lib"
	libtime "github.com/joltify-finance/joltify_lending/lib/time"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/prices/types"
)

type (
	Keeper struct {
		cdc                            codec.BinaryCodec
		storeKey                       storetypes.StoreKey
		indexPriceCache                *pricefeedservertypes.MarketToExchangePrices
		timeProvider                   libtime.TimeProvider
		indexerEventManager            indexer_manager.IndexerEventManager
		marketToCreatedAt              map[uint32]time.Time
		authorities                    map[string]struct{}
		currencyPairIDCache            *CurrencyPairIDCache
		currencyPairIdCacheInitialized *atomic.Bool
	}
)

var _ types.PricesKeeper = &Keeper{}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	indexPriceCache *pricefeedservertypes.MarketToExchangePrices,
	timeProvider libtime.TimeProvider,
	indexerEventManager indexer_manager.IndexerEventManager,
	authorities []string,
) *Keeper {
	return &Keeper{
		cdc:                            cdc,
		storeKey:                       storeKey,
		indexPriceCache:                indexPriceCache,
		timeProvider:                   timeProvider,
		indexerEventManager:            indexerEventManager,
		marketToCreatedAt:              map[uint32]time.Time{},
		authorities:                    lib.UniqueSliceToSet(authorities),
		currencyPairIDCache:            NewCurrencyPairIDCache(),
		currencyPairIdCacheInitialized: &atomic.Bool{}, // Initialized to false
	}
}

func (k Keeper) GetIndexerEventManager() indexer_manager.IndexerEventManager {
	return k.indexerEventManager
}

func (k Keeper) InitializeForGenesis(ctx sdk.Context) {
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With(log.ModuleKey, fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) HasAuthority(authority string) bool {
	_, ok := k.authorities[authority]
	return ok
}

func (k Keeper) InitializeCurrencyPairIdCache(ctx sdk.Context) {
	alreadyInitialized := k.currencyPairIdCacheInitialized.Swap(true)
	if alreadyInitialized {
		return
	}

	// Load the currency pair IDs for the markets from the x/prices state.
	k.LoadCurrencyPairIDCache(ctx)
}