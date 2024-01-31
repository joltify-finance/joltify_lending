package types

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store keys.
var (
	TokenQuota = []byte("tokenquota")
	WhiteList  = []byte("whitelist")

	_ paramtypes.ParamSet = &Params{}
)

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// default gamm module parameters.
func DefaultParams() Params {
	return Params{
		TokenQuota: "100uatom,1000000ujolt",
		Whitelist:  []string{},
	}
}

// validate params.
func (p Params) Validate() error {
	c, err := sdk.ParseCoinsNormalized(p.TokenQuota)
	if err != nil {
		return err
	}

	if c == nil {
		return errors.New("empty coins")
	}

	for _, addr := range p.Whitelist {
		_, err = sdk.AccAddressFromBech32(addr)
		if err != nil {
			return err
		}
	}
	return nil
}

func validateQuota(i interface{}) error {
	a, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	c, err := sdk.ParseCoinsNormalized(a)
	if err != nil {
		return err
	}
	if c == nil {
		return errors.New("empty coins")
	}
	return nil
}

func validateWhitelist(i interface{}) error {
	addresses, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	for _, addr := range addresses {
		_, err := sdk.AccAddressFromBech32(addr)
		if err != nil {
			return err
		}
	}
	return nil
}

// Implements params.ParamSet.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	var addresses []sdk.AccAddress
	for _, addr := range p.Whitelist {
		accAddr, err := sdk.AccAddressFromBech32(addr)
		if err != nil {
			panic(err)
		}
		addresses = append(addresses, accAddr)
	}

	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(TokenQuota, &p.TokenQuota, validateQuota),
		paramtypes.NewParamSetPair(WhiteList, &addresses, validateWhitelist),
	}
}
