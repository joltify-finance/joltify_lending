package types

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"

	"github.com/gogo/protobuf/proto"
	tmrand "github.com/tendermint/tendermint/libs/rand"

	"github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
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
	acc, err := types.AccAddressFromBech32("jolt10jghunnwjka54yzvaly4pjcxmarkvevzvq8cvl")
	if err != nil {
		panic(err)
	}
	var projects Projects
	allProjects := make([]*ProjectInfo, 100)
	projects.Items = allProjects
	for i := 0; i < 100; i++ {
		b := BasicInfo{
			"This is the test info",
			"empty",
			"ABC",
			"ABC123",
			[]byte("reserved"),
			"This is the Test Project 1",
			"example@example.com",
			"example",
			"empty logo url",
			"empty project Brief",
			"empty project description",
		}
		pi := ProjectInfo{
			Index:                        int32(i + 1),
			SPVName:                      strconv.Itoa(i) + ":" + tmrand.NewRand().Str(10),
			ProjectOwner:                 acc,
			BasicInfo:                    &b,
			ProjectLength:                480, // 5 mins
			SeparatePool:                 true,
			BaseApy:                      types.NewDecWithPrec(10, 2),
			PayFreq:                      "120",
			PoolLockedSeconds:            100,
			PoolTotalBorrowLimit:         100,
			MarketId:                     "aud:usd",
			WithdrawRequestWindowSeconds: 30,
		}
		pi.BasicInfo.ProjectName = fmt.Sprintf("this is the project %v", i)
		allProjects[i] = &pi

		// indexHash := crypto.Keccak256Hash([]byte(pi.BasicInfo.ProjectName), acc.Bytes(), []byte("junior"))
	}

	b, err := proto.Marshal(&projects)
	if err != nil {
		panic("invalid parameter")
	}

	data := base64.StdEncoding.EncodeToString(b)
	return Params{data, []types.AccAddress{acc}}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyProjects, &p.ProjectInfo, validateProjectInfo),
		paramtypes.NewParamSetPair(KeyKycSubmitter, &p.Submitter, validateSubmitter),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	err := validateProjectInfo(p.ProjectInfo)
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

	for _, el := range projects.Items {
		if len(el.SPVName) == 0 {
			return errors.New("spv name cannot be nil")
		}
	}
	return nil
}
