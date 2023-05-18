package types

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/crypto"
	tmrand "github.com/tendermint/tendermint/libs/rand"

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
	acc, err := types.AccAddressFromBech32("jolt10jghunnwjka54yzvaly4pjcxmarkvevzvq8cvl")
	if err != nil {
		panic(err)
	}
	var out []string
	allProjects := make([]*ProjectInfo, 100)
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

		indexHash := crypto.Keccak256Hash([]byte(pi.BasicInfo.ProjectName), acc.Bytes(), []byte("junior"))
		out = append(out, indexHash.Hex())
	}
	for _, el := range out {
		fmt.Printf("%v,", el)
	}
	fmt.Printf("\n")
	return Params{allProjects, []types.AccAddress{acc}}
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
