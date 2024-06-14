package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var KeyMarkets = []byte("Markets")

func GenDefaultMarket() Markets {
	acc, err := types.AccAddressFromBech32("jolt1xdpg5l3pxpyhxqg4ey4krq2pf9d3sphmmuuugg")
	if err != nil {
		panic(err)
	}

	acc2, err := types.AccAddressFromBech32("jolt1a33x0juy5t8a0zgksfz50yluw8jyvy764p9ych")
	if err != nil {
		panic(err)
	}

	m := Market{
		MarketID:   "jolt:usd",
		BaseAsset:  "jolt",
		QuoteAsset: "usd",
		Active:     true,
		Oracles:    []types.AccAddress{acc, acc2},
	}

	m2 := Market{
		MarketID:   "bnb:usd",
		BaseAsset:  "bnb",
		QuoteAsset: "usd",
		Active:     true,
		Oracles:    []types.AccAddress{acc, acc2},
	}

	m3 := Market{
		MarketID:   "usdt:usd",
		BaseAsset:  "usdt",
		QuoteAsset: "usd",
		Active:     true,
		Oracles:    []types.AccAddress{acc, acc2},
	}

	m4 := Market{
		MarketID:   "usdc:usd",
		BaseAsset:  "usdc",
		QuoteAsset: "usd",
		Active:     true,
		Oracles:    []types.AccAddress{acc, acc2},
	}

	m5 := Market{
		MarketID:   "eth:usd",
		BaseAsset:  "eth",
		QuoteAsset: "usd",
		Active:     true,
		Oracles:    []types.AccAddress{acc, acc2},
	}

	m6 := Market{
		MarketID:   "btc:usd",
		BaseAsset:  "btc",
		QuoteAsset: "usd",
		Active:     true,
		Oracles:    []types.AccAddress{acc, acc2},
	}

	m7 := Market{
		MarketID:   "atom:usd",
		BaseAsset:  "atom",
		QuoteAsset: "usd",
		Active:     true,
		Oracles:    []types.AccAddress{acc, acc2},
	}

	m8 := Market{
		MarketID:   "aud:usd",
		BaseAsset:  "aud",
		QuoteAsset: "usd",
		Active:     true,
		Oracles:    []types.AccAddress{acc, acc2},
	}

	m9 := Market{
		MarketID:   "avax:usd",
		BaseAsset:  "avax",
		QuoteAsset: "usd",
		Active:     true,
		Oracles:    []types.AccAddress{acc, acc2},
	}

	m10 := Market{
		MarketID:   "usd:usd",
		BaseAsset:  "aud",
		QuoteAsset: "usd",
		Active:     true,
		Oracles:    []types.AccAddress{acc},
	}

	m11 := Market{
		MarketID:   "matic:usd",
		BaseAsset:  "matic",
		QuoteAsset: "usd",
		Active:     true,
		Oracles:    []types.AccAddress{acc, acc2},
	}

	return []Market{m, m2, m3, m4, m5, m6, m7, m8, m9, m10, m11}
}

// NewParams creates a new AssetParams object
func NewParams(markets []Market) Params {
	return Params{
		Markets: markets,
	}
}

// DefaultParams default params for pricefeed
func DefaultParams() Params {
	return NewParams(GenDefaultMarket())
}

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs
// pairs of pricefeed module's parameters.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMarkets, &p.Markets, validateMarketParams),
	}
}

// Validate ensure that params have valid values
func (p Params) Validate() error {
	return validateMarketParams(p.Markets)
}

func validateMarketParams(i interface{}) error {
	markets, ok := i.(Markets)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return markets.Validate()
}
