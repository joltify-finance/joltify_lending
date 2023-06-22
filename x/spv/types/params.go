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
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	// 100 usdc
	amt, ok := sdk.NewIntFromString("100000000000000000000")
	if !ok {
		panic("invalid threshold setting")
	}
	return Params{BurnThreshold: sdk.NewCoin(SupportedToken, amt)}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyBurnThreshold, &p.BurnThreshold, validateBurnToken),
	}
}

func validateBurnToken(i interface{}) error {
	co, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if co.Denom != SupportedToken {
		return fmt.Errorf("we only accept ausdc and current is %v", co.Denom)
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
