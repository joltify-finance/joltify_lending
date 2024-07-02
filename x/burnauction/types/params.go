package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyBurnThreshold     = []byte("auctionBurnThreshold")
	DefaultBurnThreshold = sdk.NewCoins(sdk.NewCoin("ibc/65D0BEC6DAD96C7F5043D1E54E54B6BB5D5B3AEC3FF6CEBB75B9E059F3580EA3", sdk.NewInt(15e6)))
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	burnThreshold sdk.Coins,
) Params {
	return Params{
		BurnThreshold: burnThreshold,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultBurnThreshold,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyBurnThreshold, &p.BurnThreshold, validateBurnThreshold),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateBurnThreshold(p.BurnThreshold); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateBurnThreshold validates the Burnthreshold param
func validateBurnThreshold(v interface{}) error {
	burnThreshold, ok := v.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	if burnThreshold.IsZero() {
		return fmt.Errorf("burn threshold should not be zero")
	}

	if !burnThreshold.IsValid() {
		return fmt.Errorf("invalid support coins: %s", burnThreshold)
	}

	return nil
}
