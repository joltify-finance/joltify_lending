package types

import (
	"errors"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/joltify-finance/joltify_lending/client"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	tokenThreshold      = []byte("tokenThreshold")
	preAccountThreshold = []byte("preaccountThreshold")
	whitelist           = []byte("whitelist")
	banlist             = []byte("banlist")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	// the coin list is the amount of USD for the given token, 100jolt means 100 USD value of jolt

	if client.MAINNETFLAG == "false" {
		loweramount1 := strings.ToLower("300e189117A26BA81E29FA4F78F57DC2BD90CD3D26848101BA880445F119B22A1E254E,10000e18ujolt")
		loweramount2 := strings.ToLower("10e189117A26BA81E29FA4F78F57DC2BD90CD3D26848101BA880445F119B22A1E254E,100e18ujolt")
		quota, err := sdk.ParseCoinsNormalized(loweramount1)
		if err != nil {
			panic(err)
		}

		preAccountQuota, err := sdk.ParseCoinsNormalized(loweramount2)
		if err != nil {
			panic(err)
		}

		// eacho block takes 5 seconds, so we have 3600*24/5=17280 blocks per day
		targets := Target{
			"ibc",
			quota,
			17280,
		}

		perAccountTargets := Target{
			"ibc",
			preAccountQuota,
			17280,
		}
		w := WhiteList{
			"ibc",
			[]string{"jolt1gl7gfy5tjf9wlpumprya3fffxmdmlwcyykx8np"},
		}

		b := BanList{
			"ibc",
			[]string{"jolt1xdp3ralsry3ux4nuraq9qzr8zzc9r9nh0v3y56"},
		}

		return Params{[]*Target{&targets}, []*Target{&perAccountTargets}, []*WhiteList{&w}, []*BanList{&b}}
	}

	loweramount1 := strings.ToLower("10000ujolt")
	loweramount2 := strings.ToLower("100ujolt")
	quota, err := sdk.ParseCoinsNormalized(loweramount1)
	if err != nil {
		panic(err)
	}

	preAccountQuota, err := sdk.ParseCoinsNormalized(loweramount2)
	if err != nil {
		panic(err)
	}

	// eacho block takes 5 seconds, so we have 3600*24/5=17280 blocks per day
	targets := Target{
		"ibc",
		quota,
		17280,
	}

	perAccountTargets := Target{
		"ibc",
		preAccountQuota,
		17280,
	}
	w := WhiteList{
		"ibc",
		nil,
	}

	b := BanList{
		"ibc",
		nil,
	}

	return Params{[]*Target{&targets}, []*Target{&perAccountTargets}, []*WhiteList{&w}, []*BanList{&b}}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(tokenThreshold, &p.Targets, validateQuotaSet),
		paramtypes.NewParamSetPair(preAccountThreshold, &p.PerAccounttargets, validateQuotaSet),
		paramtypes.NewParamSetPair(whitelist, &p.Whitelist, validateWhitelist),
		paramtypes.NewParamSetPair(banlist, &p.Banlist, validateBanlist),
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

// Validate validates the set of params
func validateBanlist(i interface{}) error {
	co, ok := i.([]*BanList)
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

	for _, target := range p.PerAccounttargets {
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

	err := validateBanlist(p.Banlist)
	if err != nil {
		return err
	}

	return validateWhitelist(p.Whitelist)
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
