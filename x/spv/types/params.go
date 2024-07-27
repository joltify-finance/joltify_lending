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
	KeyIncentive                         = []byte("incentive")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewTestParams creates a new Params instance
func NewTestParams() Params {
	// 100 usdc
	// amt, ok := sdkmath.NewIntFromString("100000000000000000000")
	amt, ok := sdkmath.NewIntFromString("10000000000000000")
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

// NewParams creates a new Params instance
func NewParams() Params {
	// 100 usdc
	// amt, ok := sdkmath.NewIntFromString("100000000000000000000")
	amt, ok := sdkmath.NewIntFromString("10000000000000000")
	if !ok {
		panic("invalid threshold setting")
	}
	us := sdk.NewCoin("ibc/65D0BEC6DAD96C7F5043D1E54E54B6BB5D5B3AEC3FF6CEBB75B9E059F3580EA3", amt)

	market := Moneymarket{
		Denom:            "ibc/65D0BEC6DAD96C7F5043D1E54E54B6BB5D5B3AEC3FF6CEBB75B9E059F3580EA3",
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
		paramtypes.NewParamSetPair(KeyIncentive, &p.Incentives, validateIncentive),
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

func validateIncentive(i interface{}) error {
	co, ok := i.([]Incentive)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	for _, el := range co {
		if el.Poolid == "" {
			return fmt.Errorf("invalid pool id: %s", el.Poolid)
		}
		_, err := sdkmath.LegacyNewDecFromStr(el.Spy)
		if err != nil {
			return fmt.Errorf("invalid spy: %s with err %v", el.Spy, err)
		}
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
