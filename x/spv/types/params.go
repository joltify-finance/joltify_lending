package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var (
	_                paramtypes.ParamSet = (*Params)(nil)
	KeyBurnThreshold                     = []byte("burnthreshold")
	KeyMoneyMarket                       = []byte("moneymarket")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	// 100 usdc
	// amt, ok := sdk.NewIntFromString("100000000000000000000")
	amt, ok := sdk.NewIntFromString("10000000000000000")
	if !ok {
		panic("invalid threshold setting")
	}
	us := sdk.NewCoin("ausdc", amt)

	market := Moneymarket{
		Denom:            "ausdc",
		ConversionFactor: 6,
	}

	return Params{BurnThreshold: sdk.NewCoins(us), Markets: []Moneymarket{market}}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyBurnThreshold, &p.BurnThreshold, validateBurnToken),
		paramtypes.NewParamSetPair(KeyMoneyMarket, &p.Markets, validateMoneyMarket),
	}
}

func validateBurnToken(i interface{}) error {
	co, ok := i.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !co.IsValid() {
		return fmt.Errorf("invalid coins: %s", co)
	}
	return nil
}

func validateMoneyMarket(i interface{}) error {
	co, ok := i.([]Moneymarket)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	for _, el := range co {
		if el.ConversionFactor > 18 {
			return fmt.Errorf("invalid conversion factor: %d", el.ConversionFactor)
		}
	}
	return nil
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateBurnToken(p.BurnThreshold); err != nil {
		return err
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
