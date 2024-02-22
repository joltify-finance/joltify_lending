package types

import (
	"fmt"

	"github.com/joltify-finance/joltify_lending/client"

	"github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// Parameter keys
var (
	KeyKycSubmitter = []byte("kycsubmitters")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	if client.MAINNETFLAG == "false" {
		acc, err := types.AccAddressFromBech32("jolt15qdefkmwswysgg4qxgqpqr35k3m49pkxu8ygkq")
		if err != nil {
			panic(err)
		}

		return Params{[]types.AccAddress{acc}}
	}
	return Params{}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyKycSubmitter, &p.Submitter, validateSubmitter),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return validateSubmitter(p.Submitter)
}

func validateSubmitter(i interface{}) error {
	_, ok := i.([]types.AccAddress)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}
