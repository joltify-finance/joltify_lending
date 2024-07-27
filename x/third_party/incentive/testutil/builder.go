package testutil

import (
	"time"

	types3 "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/app"
)

const (
	oneYear = time.Hour * 24 * 365
)

// IncentiveGenesisBuilder is a tool for creating an incentive genesis state.
// Helper methods add values onto a default genesis state.
// All methods are immutable and return updated copies of the builder.
type IncentiveGenesisBuilder struct {
	types3.GenesisState
	genesisTime time.Time
}

func NewIncentiveGenesisBuilder() IncentiveGenesisBuilder {
	return IncentiveGenesisBuilder{
		GenesisState: types3.DefaultGenesisState(),
		genesisTime:  time.Time{},
	}
}

func (builder IncentiveGenesisBuilder) Build() types3.GenesisState {
	return builder.GenesisState
}

func (builder IncentiveGenesisBuilder) BuildMarshalled(cdc codec.JSONCodec) app.GenesisState {
	built := builder.Build()

	return app.GenesisState{
		types3.ModuleName: cdc.MustMarshalJSON(&built),
	}
}

func (builder IncentiveGenesisBuilder) WithGenesisTime(time time.Time) IncentiveGenesisBuilder {
	builder.genesisTime = time
	builder.Params.ClaimEnd = time.Add(5 * oneYear)
	return builder
}

// WithInitializedBorrowRewardPeriod sets the genesis time as the previous accumulation time for the specified period.
// This can be helpful in tests. With no prev time set, the first block accrues no rewards as it just sets the prev time to the current.
func (builder IncentiveGenesisBuilder) WithInitializedBorrowRewardPeriod(period types3.MultiRewardPeriod) IncentiveGenesisBuilder {
	builder.Params.JoltBorrowRewardPeriods = []types3.MultiRewardPeriod{}
	builder.Params.JoltBorrowRewardPeriods = append(builder.Params.JoltBorrowRewardPeriods, period)

	accumulationTimeForPeriod := types3.NewAccumulationTime(period.CollateralType, builder.genesisTime)
	builder.JoltBorrowRewardState.AccumulationTimes = append(
		builder.JoltBorrowRewardState.AccumulationTimes,
		accumulationTimeForPeriod,
	)

	// TODO remove to better reflect real states
	builder.JoltBorrowRewardState.MultiRewardIndexes = builder.JoltBorrowRewardState.MultiRewardIndexes.With(
		period.CollateralType,
		newZeroRewardIndexesFromCoins(period.RewardsPerSecond...),
	)

	return builder
}

func (builder IncentiveGenesisBuilder) WithSimpleBorrowRewardPeriod(ctype string, rewardsPerSecond sdk.Coins) IncentiveGenesisBuilder {
	return builder.WithInitializedBorrowRewardPeriod(builder.simpleRewardPeriod(ctype, rewardsPerSecond))
}

// WithInitializedSupplyRewardPeriod sets the genesis time as the previous accumulation time for the specified period.
// This can be helpful in tests. With no prev time set, the first block accrues no rewards as it just sets the prev time to the current.
func (builder IncentiveGenesisBuilder) WithInitializedSupplyRewardPeriod(period types3.MultiRewardPeriod) IncentiveGenesisBuilder {
	builder.Params.JoltSupplyRewardPeriods = append(builder.Params.JoltSupplyRewardPeriods, period)

	accumulationTimeForPeriod := types3.NewAccumulationTime(period.CollateralType, builder.genesisTime)
	builder.JoltSupplyRewardState.AccumulationTimes = append(
		builder.JoltSupplyRewardState.AccumulationTimes,
		accumulationTimeForPeriod,
	)

	// TODO remove to better reflect real states
	builder.JoltSupplyRewardState.MultiRewardIndexes = builder.JoltSupplyRewardState.MultiRewardIndexes.With(
		period.CollateralType,
		newZeroRewardIndexesFromCoins(period.RewardsPerSecond...),
	)

	return builder
}

func (builder IncentiveGenesisBuilder) WithSimpleSupplyRewardPeriod(ctype string, rewardsPerSecond sdk.Coins) IncentiveGenesisBuilder {
	return builder.WithInitializedSupplyRewardPeriod(builder.simpleRewardPeriod(ctype, rewardsPerSecond))
}

func (builder IncentiveGenesisBuilder) WithMultipliers(multipliers types3.MultipliersPerDenoms) IncentiveGenesisBuilder {
	builder.Params.ClaimMultipliers = multipliers

	return builder
}

func (builder IncentiveGenesisBuilder) simpleRewardPeriod(ctype string, rewardsPerSecond sdk.Coins) types3.MultiRewardPeriod {
	return types3.NewMultiRewardPeriod(
		true,
		ctype,
		builder.genesisTime,
		builder.genesisTime.Add(4*oneYear),
		rewardsPerSecond,
	)
}

func newZeroRewardIndexesFromCoins(coins ...sdk.Coin) types3.RewardIndexes {
	var ri types3.RewardIndexes
	for _, coin := range coins {
		ri = ri.With(coin.Denom, sdkmath.LegacyZeroDec())
	}
	return ri
}

// JoltGenesisBuilder is a tool for creating a jolt genesis state.
// Helper methods add values onto a default genesis state.
// All methods are immutable and return updated copies of the builder.
type JoltGenesisBuilder struct {
	types2.GenesisState
	genesisTime time.Time
}

func NewJoltGenesisBuilder() JoltGenesisBuilder {
	return JoltGenesisBuilder{
		GenesisState: types2.DefaultGenesisState(),
	}
}

func (builder JoltGenesisBuilder) Build() types2.GenesisState {
	return builder.GenesisState
}

func (builder JoltGenesisBuilder) BuildMarshalled(cdc codec.JSONCodec) app.GenesisState {
	built := builder.Build()

	return app.GenesisState{
		types2.ModuleName: cdc.MustMarshalJSON(&built),
	}
}

func (builder JoltGenesisBuilder) WithGenesisTime(genTime time.Time) JoltGenesisBuilder {
	builder.genesisTime = genTime
	return builder
}

func (builder JoltGenesisBuilder) WithInitializedMoneyMarket(market types2.MoneyMarket) JoltGenesisBuilder {
	builder.Params.MoneyMarkets = append(builder.Params.MoneyMarkets, market)

	builder.PreviousAccumulationTimes = append(
		builder.PreviousAccumulationTimes,
		types2.NewGenesisAccumulationTime(market.Denom, builder.genesisTime, sdkmath.LegacyOneDec(), sdkmath.LegacyOneDec()),
	)
	return builder
}

func (builder JoltGenesisBuilder) WithMinBorrow(minUSDValue sdkmath.LegacyDec) JoltGenesisBuilder {
	builder.Params.MinimumBorrowUSDValue = minUSDValue
	return builder
}

func NewStandardMoneyMarket(denom string) types2.MoneyMarket {
	return types2.NewMoneyMarket(
		denom,
		types2.NewBorrowLimit(
			false,
			sdk.NewDec(1e15),
			sdkmath.LegacyMustNewDecFromStr("0.6"),
		),
		denom+":usd",
		sdkmath.NewInt(1e6),
		types2.NewInterestRateModel(
			sdkmath.LegacyMustNewDecFromStr("0.05"),
			sdkmath.LegacyMustNewDecFromStr("2"),
			sdkmath.LegacyMustNewDecFromStr("0.8"),
			sdkmath.LegacyMustNewDecFromStr("10"),
		),
		sdkmath.LegacyMustNewDecFromStr("0.05"),
		sdkmath.LegacyZeroDec(),
	)
}
