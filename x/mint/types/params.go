package types

import (
	"errors"
	"fmt"

	sdkmath "cosmossdk.io/math"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyFirstProvision     = []byte("FirstProvision")
	DefaultFirstProvision = sdkmath.LegacyMustNewDecFromStr("0")
	DefaultNodeSPY        = sdkmath.LegacyMustNewDecFromStr("1.000000002440418609")
	NodeSPY               = []byte("NodeSPY")
)

var (
	KeyUnit     = []byte("Unit")
	DefaultUnit = "minute"
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	firstProvision sdkmath.LegacyDec,
	unit string,
	nodeSPY sdkmath.LegacyDec,
) Params {
	return Params{
		FirstProvisions: firstProvision,
		Unit:            unit,
		NodeSPY:         nodeSPY,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	params := NewParams(
		DefaultFirstProvision,
		DefaultUnit,
		DefaultNodeSPY,
	)
	return params
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyFirstProvision, &p.FirstProvisions, validateEachProvision),
		paramtypes.NewParamSetPair(KeyUnit, &p.Unit, validateUnit),
		paramtypes.NewParamSetPair(NodeSPY, &p.NodeSPY, validateSPY),
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

	if err := validateSPY(p.NodeSPY); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateSPY(v interface{}) error {
	nodeAPY, ok := v.(sdkmath.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}
	if nodeAPY.IsNil() || nodeAPY.IsNegative() {
		return fmt.Errorf("nodeAPY should not be nil or negtive")
	}
	return nil
}

// validateEachProvision validates the CurrentProvision param
func validateEachProvision(v interface{}) error {
	eachProvision, ok := v.(sdkmath.LegacyDec)
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
