package types

import (
	stderr "errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	yaml "gopkg.in/yaml.v2"
)

const (
	DefaultBlockChurnInterval = 241920
	DEFAULTPOWER              = 10000
	DEFAULTSTEP               = 1000
	DEFAULTRATIO              = 80
	TARGETQUOTA               = "300000000000000000000abnb,500000000000000000000000ausdt,500000000000000000000000abusd,100000000000000000000aeth,500000000000000000000000ajolt"
)

var (
	KeyBlockChurnInterval = []byte("blockChurnInterval")
	KeyPower              = []byte("power")
	KeyStep               = []byte("step")
	KeyRatio              = []byte("candidateRatio")
	KeyTargetQuota        = []byte("targetQuota")
	KeyHistoryLength      = []byte("quotaHistoryLength")
)
var _ paramtypes.ParamSet = (*Params)(nil)

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyBlockChurnInterval, &p.BlockChurnInterval, validateInteger),
		paramtypes.NewParamSetPair(KeyPower, &p.Power, validateInteger),
		paramtypes.NewParamSetPair(KeyStep, &p.Step, validateInteger),
		paramtypes.NewParamSetPair(KeyRatio, &p.CandidateRatio, validateFloat),
		paramtypes.NewParamSetPair(KeyTargetQuota, &p.TargetQuota, validateArray),
		paramtypes.NewParamSetPair(KeyHistoryLength, &p.HistoryLength, validateInteger32),
	}
}

// NewParams creates a new Params instance
func NewParams(blockChurnInterval, power, step int64, ratio sdk.Dec, targetQuota sdk.Coins, historyLength int32) Params {
	return Params{
		BlockChurnInterval: blockChurnInterval,
		Power:              power,
		Step:               step,
		CandidateRatio:     ratio,
		TargetQuota:        targetQuota,
		HistoryLength:      historyLength,
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	tokensQuota, err := sdk.ParseCoinsNormalized(TARGETQUOTA)
	if err != nil {
		panic("invaild target quota")
	}
	return NewParams(
		DefaultBlockChurnInterval,
		DEFAULTPOWER,
		DEFAULTSTEP,
		sdk.NewDecWithPrec(DEFAULTRATIO, 2),
		tokensQuota,
		14400,
	)
}

// String returns a human readable string representation of the parameters.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validate a set of params
func (p Params) Validate() error {
	if err := validateInteger(p.BlockChurnInterval); err != nil {
		return err
	}

	if err := validateInteger(p.Power); err != nil {
		return err
	}

	if err := validateInteger(p.Step); err != nil {
		return err
	}

	if err := validateInteger(p.CandidateRatio); err != nil {
		return err
	}
	if err := validateArray(p.TargetQuota); err != nil {
		return err
	}

	if err := validateInteger(p.HistoryLength); err != nil {
		return err
	}

	return nil
}

func validateInteger(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < 0 {
		return stderr.New("invalid height")
	}

	return nil
}

func validateInteger32(i interface{}) error {
	v, ok := i.(int32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < 0 {
		return stderr.New("invalid height")
	}

	return nil
}

func validateFloat(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("goal bonded cannot be negative: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("goal bonded too large: %s", v)
	}

	return nil
}

func validateArray(i interface{}) error {
	v, ok := i.(sdk.Coins)
	if !ok {
		return fmt.Errorf("not a valid coins structure")
	}

	if v.Empty() || v.IsValid() {
		return nil
	}

	return fmt.Errorf("invalid coins list")
}
