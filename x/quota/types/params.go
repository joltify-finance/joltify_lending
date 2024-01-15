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
	tokenThreshold = []byte("tokenThreshold")
	whitelist      = []byte("whitelist")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	// the coin list is the amount of USD for the given token, 100jolt means 100 USD value of jolt
	quota, err := sdk.ParseCoinsNormalized("100ujolt")
	if err != nil {
		panic(err)
	}

	targets := Target{
		"ibc",
		quota,
		40,
	}
	w := WhiteList{
		"ibc",
		nil,
	}

	return Params{[]*Target{&targets}, []*WhiteList{&w}}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(tokenThreshold, &p.Targets, validateQuotaSet),
		paramtypes.NewParamSetPair(whitelist, &p.Whitelist, validateWhitelist),
	}
}

// Validate validates the set of params
func validateWhitelist(i interface{}) error {
	co, ok := i.([]*WhiteList)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	for _, el := range co {
		if el.ModuleName == "" {
			return errors.New("invalid module name")
		}
		for _, addr := range el.AddressList {
			_, err := sdk.AccAddressFromBech32(addr)
			if err != nil {
				return errors.New("invalid address")
			}
		}
	}

	return nil
}

func validateQuotaSet(i interface{}) error {
	co, ok := i.([]*Target)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	for _, target := range co {
		if target.ModuleName == "" {
			return errors.New("invalid quota module name")
		}
		if target.HistoryLength < 1 {
			return errors.New("invalid history length")
		}

		if !isSorted(target.CoinsSum) {
			return errors.New("the tokens are not sorted")
		}

		if target.CoinsSum.IsZero() {
			return errors.New("invalid quota sum")
		}
	}

	return nil
}

func isSorted(coins sdk.Coins) bool {
	for i := 1; i < len(coins); i++ {
		if coins[i-1].Denom > coins[i].Denom {
			return false
		}
	}
	return true
}

// Validate validates the set of params
func (p Params) Validate() error {
	for _, target := range p.Targets {
		if target.CoinsSum.IsZero() {
			return errors.New("invalid quota sum")
		}

		if !isSorted(target.CoinsSum) {
			return errors.New("the token is not sorted")
		}

		if target.ModuleName == "" {
			return errors.New("invalid module name")
		}
		if target.HistoryLength < 1 {
			return errors.New("invalid history length")
		}
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
