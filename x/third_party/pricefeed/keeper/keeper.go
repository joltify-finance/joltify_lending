package keeper

import (
	"fmt"
	"sort"
	"time"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/types"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

//// Keeper struct for pricefeed module
//type Keeper interface {
//
//	// key used to access the stores from Context
//	key storetypes.StoreKey
//// Codec for binary encoding/decoding
//cdc codec.Codec
//// The reference to the Paramstore to get and set pricefeed specific params
//paramSubspace paramtypes.Subspace
//}

// Keeper struct for pricefeed module
type Keeper struct {
	// key used to access the stores from Context
	key storetypes.StoreKey
	// Codec for binary encoding/decoding
	cdc codec.Codec
	// The reference to the Paramstore to get and set pricefeed specific params
	paramSubspace paramtypes.Subspace
}

// NewKeeper returns a new keeper for the pricefeed module.
func NewKeeper(
	cdc codec.Codec, key storetypes.StoreKey, paramstore paramtypes.Subspace,
) Keeper {
	if !paramstore.HasKeyTable() {
		paramstore = paramstore.WithKeyTable(types2.ParamKeyTable())
	}

	return Keeper{
		cdc:           cdc,
		key:           key,
		paramSubspace: paramstore,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types2.ModuleName))
}

// SetPrice updates the posted price for a specific oracle
func (k Keeper) SetPrice(
	ctx sdk.Context,
	oracle sdk.AccAddress,
	marketID string,
	price sdk.Dec,
	expiry time.Time,
) (types2.PostedPrice, error) {
	// If the expiry is less than or equal to the current blockheight, we consider the price valid
	if !expiry.After(ctx.BlockTime()) {
		return types2.PostedPrice{}, types2.ErrExpired
	}

	store := ctx.KVStore(k.key)

	newRawPrice := types2.NewPostedPrice(marketID, oracle, price, expiry)

	// Emit an event containing the oracle's new price
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types2.EventTypeOracleUpdatedPrice,
			sdk.NewAttribute(types2.AttributeMarketID, marketID),
			sdk.NewAttribute(types2.AttributeOracle, oracle.String()),
			sdk.NewAttribute(types2.AttributeMarketPrice, price.String()),
			sdk.NewAttribute(types2.AttributeExpiry, expiry.UTC().String()),
		),
	)

	// Sets the raw price for a single oracle instead of an array of all oracle's raw prices
	store.Set(types2.RawPriceKey(marketID, oracle), k.cdc.MustMarshal(&newRawPrice))
	return newRawPrice, nil
}

// SetCurrentPrices updates the price of an asset to the median of all valid oracle inputs
func (k Keeper) SetCurrentPrices(ctx sdk.Context, marketID string) error {
	_, ok := k.GetMarket(ctx, marketID)
	if !ok {
		return sdkerrors.Wrap(types2.ErrInvalidMarket, marketID)
	}
	// store current price
	validPrevPrice := true
	prevPrice, err := k.GetCurrentPrice(ctx, marketID)
	if err != nil {
		validPrevPrice = false
	}

	prices := k.GetRawPrices(ctx, marketID)

	var notExpiredPrices []types2.CurrentPrice
	// filter out expired prices
	for _, v := range prices {
		if v.Expiry.After(ctx.BlockTime()) {
			notExpiredPrices = append(notExpiredPrices, types2.NewCurrentPrice(v.MarketID, v.Price))
		}
	}

	if len(notExpiredPrices) == 0 {
		// NOTE: The current price stored will continue storing the most recent (expired)
		// price if this is not set.
		// This zero's out the current price stored value for that market and ensures
		// that CDP methods that GetCurrentPrice will return error.
		k.setCurrentPrice(ctx, marketID, types2.CurrentPrice{})
		return types2.ErrNoValidPrice
	}

	medianPrice := k.CalculateMedianPrice(notExpiredPrices)

	// check case that market price was not set in genesis
	if validPrevPrice && !medianPrice.Equal(prevPrice.Price) {
		// only emit event if price has changed
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types2.EventTypeMarketPriceUpdated,
				sdk.NewAttribute(types2.AttributeMarketID, marketID),
				sdk.NewAttribute(types2.AttributeMarketPrice, medianPrice.String()),
			),
		)
	}

	currentPrice := types2.NewCurrentPrice(marketID, medianPrice)
	k.setCurrentPrice(ctx, marketID, currentPrice)

	return nil
}

func (k Keeper) setCurrentPrice(ctx sdk.Context, marketID string, currentPrice types2.CurrentPrice) {
	store := ctx.KVStore(k.key)
	store.Set(types2.CurrentPriceKey(marketID), k.cdc.MustMarshal(&currentPrice))
}

// CalculateMedianPrice calculates the median prices for the input prices.
func (k Keeper) CalculateMedianPrice(prices []types2.CurrentPrice) sdk.Dec {
	l := len(prices)

	if l == 1 {
		// Return immediately if there's only one price
		return prices[0].Price
	}
	// sort the prices
	sort.Slice(prices, func(i, j int) bool {
		return prices[i].Price.LT(prices[j].Price)
	})
	// for even numbers of prices, the median is calculated as the mean of the two middle prices
	if l%2 == 0 {
		median := k.calculateMeanPrice(prices[l/2-1], prices[l/2])
		return median
	}
	// for odd numbers of prices, return the middle element
	return prices[l/2].Price
}

func (k Keeper) calculateMeanPrice(priceA, priceB types2.CurrentPrice) sdk.Dec {
	sum := priceA.Price.Add(priceB.Price)
	mean := sum.Quo(sdk.NewDec(2))
	return mean
}

// GetCurrentPrice fetches the current median price of all oracles for a specific market
func (k Keeper) GetCurrentPrice(ctx sdk.Context, marketID string) (types2.CurrentPrice, error) {
	store := ctx.KVStore(k.key)
	bz := store.Get(types2.CurrentPriceKey(marketID))

	if bz == nil {
		return types2.CurrentPrice{}, types2.ErrNoValidPrice
	}
	var price types2.CurrentPrice
	err := k.cdc.Unmarshal(bz, &price)
	if err != nil {
		return types2.CurrentPrice{}, err
	}
	if price.Price.Equal(sdk.ZeroDec()) {
		return types2.CurrentPrice{}, types2.ErrNoValidPrice
	}
	return price, nil
}

// IterateCurrentPrices iterates over all current price objects in the store and performs a callback function
func (k Keeper) IterateCurrentPrices(ctx sdk.Context, cb func(cp types2.CurrentPrice) (stop bool)) {
	iterator := sdk.KVStorePrefixIterator(ctx.KVStore(k.key), types2.CurrentPricePrefix)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var cp types2.CurrentPrice
		k.cdc.MustUnmarshal(iterator.Value(), &cp)
		if cb(cp) {
			break
		}
	}
}

// GetCurrentPrices returns all current price objects from the store
func (k Keeper) GetCurrentPrices(ctx sdk.Context) types2.CurrentPrices {
	var cps types2.CurrentPrices
	k.IterateCurrentPrices(ctx, func(cp types2.CurrentPrice) (stop bool) {
		cps = append(cps, cp)
		return false
	})
	return cps
}

// GetRawPrices fetches the set of all prices posted by oracles for an asset
func (k Keeper) GetRawPrices(ctx sdk.Context, marketId string) types2.PostedPrices {
	var pps types2.PostedPrices
	k.IterateRawPricesByMarket(ctx, marketId, func(pp types2.PostedPrice) (stop bool) {
		pps = append(pps, pp)
		return false
	})
	return pps
}

// IterateRawPricesByMarket iterates over all raw prices in the store and performs a callback function
func (k Keeper) IterateRawPricesByMarket(ctx sdk.Context, marketId string, cb func(record types2.PostedPrice) (stop bool)) {
	iterator := sdk.KVStorePrefixIterator(ctx.KVStore(k.key), types2.RawPriceIteratorKey(marketId))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var record types2.PostedPrice
		k.cdc.MustUnmarshal(iterator.Value(), &record)
		if cb(record) {
			break
		}
	}
}
