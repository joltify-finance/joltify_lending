package types

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"

	"github.com/cosmos/gogoproto/proto"
	"github.com/joltify-finance/joltify_lending/client"

	"github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// Parameter keys
var (
	KeyKycSubmitter = []byte("kycsubmitters")
	KeyProjects     = []byte("projectInfo")
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

		return Params{"", []types.AccAddress{acc}}
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
		paramtypes.NewParamSetPair(KeyProjects, &p.ProjectInfo, validateProjectInfo),
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

func validateProjectInfo(i interface{}) error {
	projectsStr, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	out, err := base64.StdEncoding.DecodeString(projectsStr)
	if err != nil {
		return fmt.Errorf("fail to decode the project base64 string: %v", err)
	}

	var projects Projects
	err = proto.Unmarshal(out, &projects)
	if err != nil {
		return err
	}

	indexCheck := make(map[int32]bool)

	var exist bool
	for _, el := range projects.Items {
		if len(el.SPVName) == 0 {
			return errors.New("spv name cannot be nil")
		}
		_, exist = indexCheck[el.Index]
		if exist {
			return errors.New("the index has been used")
		}
		indexCheck[el.Index] = true

		if el.PoolLockedSeconds < 0 {
			return errors.New("project time related setting cannot be negative")
		}

		if el.WithdrawRequestWindowSeconds < 0 {
			return errors.New("project time related setting cannot be negative")
		}
		if el.GraceTime.Seconds() < 0 {
			return errors.New("project time related setting cannot be negative")
		}

		if el.MinBorrowAmount.IsNegative() {
			return errors.New("min borrow amount cannot be negative")
		}

		freq, err := strconv.ParseInt(el.PayFreq, 10, 64)
		if err != nil {
			return err
		}
		if freq < 0 {
			return errors.New("pay freq cannot be negative")
		}

	}
	return nil
}
