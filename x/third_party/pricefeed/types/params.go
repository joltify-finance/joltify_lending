package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter keys
var (
	KeyMarkets = []byte("Markets")
)

func GenDefaultMarket() Markets {
	acc, err := types.AccAddressFromBech32("jolt1p3jl6udk43vw0cvc5hjqrpnncsqmsz56wd32z8")
	if err != nil {
		panic(err)
	}

	m := Market{
		MarketID:   "jolt:usd",
		BaseAsset:  "jolt",
		QuoteAsset: "usd",
		Active:     true,
		Oracles:    []types.AccAddress{acc},
	}

	m2 := Market{
		MarketID:   "bnb:usd",
		BaseAsset:  "bnb",
		QuoteAsset: "usd",
		Active:     true,
		Oracles:    []types.AccAddress{acc},
	}

	m3 := Market{
		MarketID:   "usdt:usd",
		BaseAsset:  "usdt",
		QuoteAsset: "usd",
		Active:     true,
		Oracles:    []types.AccAddress{acc},
	}

	m4 := Market{
		MarketID:   "usdc:usd",
		BaseAsset:  "usdc",
		QuoteAsset: "usd",
		Active:     true,
		Oracles:    []types.AccAddress{acc},
	}

	return []Market{m, m2, m3, m4}
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
