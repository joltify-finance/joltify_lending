package keeper

import (
	sdkmath "cosmossdk.io/math"
)

const scalingFactor = 1e18

// CalculateInterestFactor calculates the simple interest scaling factor,
// which is equal to: (per-second interest rate * number of seconds elapsed)
// Will return 1.000x, multiply by principal to get new principal with added interest
func CalculateInterestFactor(perSecondInterestRate sdkmath.LegacyDec, secondsElapsed sdkmath.Int) sdkmath.LegacyDec {
	scalingFactorUint := sdkmath.NewUint(uint64(scalingFactor))
	scalingFactorInt := sdkmath.NewInt(int64(scalingFactor))

	// Convert per-second interest rate to a uint scaled by 1e18
	interestMantissa := sdkmath.NewUintFromBigInt(perSecondInterestRate.MulInt(scalingFactorInt).RoundInt().BigInt())
	// Convert seconds elapsed to uint (*not scaled*)
	secondsElapsedUint := sdkmath.NewUintFromBigInt(secondsElapsed.BigInt())
	// Calculate the interest factor as a uint scaled by 1e18
	interestFactorMantissa := sdkmath.RelativePow(interestMantissa, secondsElapsedUint, scalingFactorUint)

	// Convert interest factor to an unscaled sdkmath.LegacyDec
	return sdkmath.LegacyNewDecFromBigInt(interestFactorMantissa.BigInt()).QuoInt(scalingFactorInt)
}
