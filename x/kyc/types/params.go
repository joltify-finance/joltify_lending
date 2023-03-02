package types

import (
	sdkmath "cosmossdk.io/math"
	"errors"
	"fmt"

	"github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// Parameter keys
var (
	KeyProjects     = []byte("projectInfo")
	KeyKycSubmitter = []byte("kycsubmitters")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	b := BasicInfo{
		"This is the test info",
		"empty",
		"ABC",
		"ABC123",
		[]byte("reserved"),
		"This is the Test Project 1",
	}
	amount, _ := sdkmath.NewIntFromString("100000000000000000000000")
	acc, err := types.AccAddressFromBech32("jolt10jghunnwjka54yzvaly4pjcxmarkvevzvq8cvl")
	if err != nil {
		panic(err)
	}
	pi := ProjectInfo{
		Index:               1,
		SPVName:             "defaultSPV",
		ProjectOwner:        acc,
		BasicInfo:           &b,
		ProjectLength:       31536000, //1 year
		ProjectTargetAmount: types.NewCoin("ausdc", amount),
		BaseApy:             types.NewDecWithPrec(15, 2),
		PayFreq:             "80",
	}
	return Params{[]*ProjectInfo{&pi}, []types.AccAddress{acc}}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyProjects, &p.ProjectsInfo, validateProjectInfo),
		paramtypes.NewParamSetPair(KeyKycSubmitter, &p.Submitter, validateSubmitter),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	err := validateProjectInfo(p.ProjectsInfo)
	if err != nil {
		return err
	}
	return validateSubmitter(p.Submitter)
}

func validateSubmitter(i interface{}) error {
	submitter, ok := i.([]types.AccAddress)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if len(submitter) == 0 {
		return errors.New("empty submitter")
	}
	return nil
}

func validateProjectInfo(i interface{}) error {
	projects, ok := i.([]*ProjectInfo)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	for _, el := range projects {
		if len(el.SPVName) == 0 {
			return errors.New("spv name cannot be nil")
		}
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
