package types

import (
	"fmt"
	"time"
)

var (
	DefaultJoltClaims         = JoltLiquidityProviderClaims{}
	DefaultSwapClaims         = SwapClaims{}
	DefaultGenesisRewardState = NewGenesisRewardState(
		AccumulationTimes{},
		MultiRewardIndexes{},
	)
	DefaultSPVGenesisRewardState = NewSPVGenesisRewardState(
		AccumulationTimes{},
		[]SPVRewardAccIndex{},
	)
)

// NewGenesisState returns a new genesis state
func NewGenesisState(
	params Params,
	joltSupplyState, joltBorrowState, swapState GenesisRewardState, spvState SPVGenesisRewardState,
	hc JoltLiquidityProviderClaims, sc SwapClaims,
) GenesisState {
	return GenesisState{
		Params: params,

		JoltSupplyRewardState:       joltSupplyState,
		JoltBorrowRewardState:       joltBorrowState,
		SwapRewardState:             swapState,
		SpvRewardState:              spvState,
		JoltLiquidityProviderClaims: hc,
		SwapClaims:                  sc,
	}
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() GenesisState {
	return GenesisState{
		Params:                      DefaultParams(),
		JoltSupplyRewardState:       DefaultGenesisRewardState,
		JoltBorrowRewardState:       DefaultGenesisRewardState,
		SwapRewardState:             DefaultGenesisRewardState,
		SpvRewardState:              DefaultSPVGenesisRewardState,
		JoltLiquidityProviderClaims: DefaultJoltClaims,
		SwapClaims:                  DefaultSwapClaims,
	}
}

// Validate performs basic validation of genesis data returning an
// error for any failed validation criteria.
func (gs GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}
	if err := gs.JoltSupplyRewardState.Validate(); err != nil {
		return err
	}
	if err := gs.JoltBorrowRewardState.Validate(); err != nil {
		return err
	}
	if err := gs.JoltLiquidityProviderClaims.Validate(); err != nil {
		return err
	}
	return nil
}

// NewGenesisRewardState returns a new GenesisRewardState
func NewGenesisRewardState(accumTimes AccumulationTimes, indexes MultiRewardIndexes) GenesisRewardState {
	return GenesisRewardState{
		AccumulationTimes:  accumTimes,
		MultiRewardIndexes: indexes,
	}
}

// NewSPVGenesisRewardState returns a new GenesisRewardState
func NewSPVGenesisRewardState(accumTimes AccumulationTimes, indexes []SPVRewardAccIndex) SPVGenesisRewardState {
	return SPVGenesisRewardState{
		AccumulationTimes: accumTimes,
		AccRewardIndexs:   indexes,
	}
}

// Validate performs validation of a GenesisRewardState
func (grs GenesisRewardState) Validate() error {
	if err := grs.AccumulationTimes.Validate(); err != nil {
		return err
	}
	return grs.MultiRewardIndexes.Validate()
}

// NewAccumulationTime returns a new GenesisAccumulationTime
func NewAccumulationTime(ctype string, prevTime time.Time) AccumulationTime {
	return AccumulationTime{
		CollateralType:           ctype,
		PreviousAccumulationTime: prevTime,
	}
}

// Validate performs validation of GenesisAccumulationTime
func (gat AccumulationTime) Validate() error {
	if len(gat.CollateralType) == 0 {
		return fmt.Errorf("genesis accumulation time's collateral type must be defined")
	}
	return nil
}

// AccumulationTimes slice of GenesisAccumulationTime
type AccumulationTimes []AccumulationTime

// Validate performs validation of GenesisAccumulationTimes
func (gats AccumulationTimes) Validate() error {
	for _, gat := range gats {
		if err := gat.Validate(); err != nil {
			return err
		}
	}
	return nil
}
