package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeySupportCoins = []byte("SupportCoins")
	// TODO: Determine the default value
	DefaultSupportCoins int32 = 0
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	supportCoins int32,
) Params {
	return Params{
		SupportCoins: supportCoins,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultSupportCoins,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeySupportCoins, &p.SupportCoins, validateSupportCoins),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateSupportCoins(p.SupportCoins); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateSupportCoins validates the SupportCoins param
func validateSupportCoins(v interface{}) error {
	supportCoins, ok := v.(int32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = supportCoins

	return nil
}
