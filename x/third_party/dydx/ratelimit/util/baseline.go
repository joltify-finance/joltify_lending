package util

import (
	"math/big"

	"github.com/dydxprotocol/v4-chain/protocol/lib"
	"github.com/joltify-finance/joltify_lending/x/third_party/dydx/ratelimit/types"
)

// GetBaseline returns the current capacity baseline for the given limiter.
// `baseline` formula:
//
//	baseline = max(baseline_minimum, baseline_tvl_ppm * current_tvl)
func GetBaseline(
	currentTvl *big.Int,
	limiter types.Limiter,
) *big.Int {
	return lib.BigMax(
		limiter.BaselineMinimum.BigInt(),
		lib.BigIntMulPpm(
			currentTvl,
			limiter.BaselineTvlPpm,
		),
	)
}
