package keeper

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const scalingFactor = 1e18

// CalculateInterestFactor calculates the simple interest scaling factor,
// which is equal to: (per-second interest rate * number of seconds elapsed)
// Will return 1.000x, multiply by principal to get new principal with added interest
func CalculateInterestFactor(perSecondInterestRate sdk.Dec, secondsElapsed sdkmath.Int) sdk.Dec {
	scalingFactorUint := sdk.NewUint(uint64(scalingFactor))
	scalingFactorInt := sdk.NewInt(int64(scalingFactor))

	// Convert per-second interest rate to a uint scaled by 1e18
	interestMantissa := sdkmath.NewUintFromBigInt(perSecondInterestRate.MulInt(scalingFactorInt).RoundInt().BigInt())
	// Convert seconds elapsed to uint (*not scaled*)
	secondsElapsedUint := sdkmath.NewUintFromBigInt(secondsElapsed.BigInt())
	// Calculate the interest factor as a uint scaled by 1e18
	interestFactorMantissa := sdkmath.RelativePow(interestMantissa, secondsElapsedUint, scalingFactorUint)

	// Convert interest factor to an unscaled sdk.Dec
	return sdk.NewDecFromBigInt(interestFactorMantissa.BigInt()).QuoInt(scalingFactorInt)
}
