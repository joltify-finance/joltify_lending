package keeper

import (
	"errors"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	OneWeek       = 7 * 24 * 3600
	OneYear       = OneWeek * 52
	OneMonth      = OneWeek * 4
	BASE          = 1
	scalingFactor = 1e18
)

func apyTospy(r sdk.Dec, seconds uint64) (sdk.Dec, error) {
	// Note: any APY 179 or greater will cause an out-of-bounds error
	root, err := r.ApproxRoot(seconds)
	if err != nil {
		return sdk.ZeroDec(), err
	}
	return root, nil
}

func CalculateInterestRate(apy sdk.Dec, payFreq int) sdk.Dec {
	// by default, we set the interest as the payment for the whole year which is 3600*24*365=31536000 seconds
	// the minimal pay frequency is one week

	seconds := BASE * payFreq
	splitAPY, err := CalculateInterestAmount(apy, payFreq)
	if err != nil {
		panic(err)
	}
	adjMonthAPY := sdk.OneDec().Add(splitAPY)
	i, err := apyTospy(adjMonthAPY, uint64(seconds))
	if err != nil {
		return sdk.Dec{}
	}

	return i
}

func CalculateInterestAmount(apy sdk.Dec, payFreq int) (sdk.Dec, error) {
	if payFreq == 0 {
		return sdk.Dec{}, errors.New("payFreq cannot be zero")
	}
	seconds := BASE * payFreq
	eachPayFreqAPY := apy.QuoTruncate(sdk.NewDec(OneYear / int64(seconds)))

	return eachPayFreqAPY, nil
}

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
