package types

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyFirstProvision     = []byte("FirstProvision")
	DefaultFirstProvision = sdk.MustNewDecFromStr("85616438.3562")
	KeyCurrentProvision   = []byte("CurrentProvision")
)

var (
	DefaultHalfCount = uint64(60 * 24 * 365)
	KeyHalfCount     = []byte("HalfCount")
)

var (
	KeyUnit     = []byte("Unit")
	DefaultUnit = "minute"
)

var (
	KeyCommunityProvision     = []byte("CommunityProvision")
	DefaultCommunityProvision = sdk.MustNewDecFromStr("50000000000000")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	firstProvision sdk.Dec,
	currentProvision sdk.Dec,
	unit string,
	communityProvision sdk.Dec,
	halfCount uint64,
) Params {
	return Params{
		FirstProvisions:     firstProvision,
		CurrentProvisions:   currentProvision,
		Unit:                unit,
		CommunityProvisions: communityProvision,
		HalfCount:           halfCount,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	params := NewParams(
		DefaultFirstProvision,
		DefaultFirstProvision,
		DefaultUnit,
		DefaultCommunityProvision,
		DefaultHalfCount,
	)
	return params
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyFirstProvision, &p.FirstProvisions, validateEachProvision),
		paramtypes.NewParamSetPair(KeyUnit, &p.Unit, validateUnit),
		paramtypes.NewParamSetPair(KeyCommunityProvision, &p.CommunityProvisions, validateEachProvision),
		paramtypes.NewParamSetPair(KeyHalfCount, &p.HalfCount, validateHalfCount),
		paramtypes.NewParamSetPair(KeyCurrentProvision, &p.CurrentProvisions, validateEachProvision),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateEachProvision(p.FirstProvisions); err != nil {
		return err
	}
	if err := validateUnit(p.Unit); err != nil {
		return err
	}
	if err := validateHalfCount(p.HalfCount); err != nil {
		return err
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateTotalCount validates the TotalDays param
func validateHalfCount(v interface{}) error {
	halfCount, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	if halfCount < 1 {
		return fmt.Errorf("total Count should not be larger than 1")
	}
	return nil
}

// validateEachProvision validates the CurrentProvision param
func validateEachProvision(v interface{}) error {
	eachProvision, ok := v.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}
	if eachProvision.IsNil() || eachProvision.IsNegative() {
		return fmt.Errorf("provision should not be nil or negtive")
	}
	return nil
}

// validateEachProvision validates the CurrentProvision param
func validateUnit(v interface{}) error {
	unit, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}
	switch unit {
	case "minute", "hour":
		return nil
	default:
		return errors.New("invalid unit")
	}
}
