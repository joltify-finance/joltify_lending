package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

const OneYear = 365 * 24 * 3600
const OneWeek = 7 * 24 * 3600

var one = sdk.NewDec(1)

func apyTospy(r sdk.Dec, seconds uint64) (sdk.Dec, error) {
	// Note: any APY 179 or greater will cause an out-of-bounds error
	root, err := r.ApproxRoot(seconds)
	if err != nil {
		return sdk.ZeroDec(), err
	}
	return root, nil
}

func CalculateInterest(apy sdk.Dec, payFreq int) (sdk.Dec, error) {
	// by default, we set the interest as the payment for the whole year which is 3600*24*365=31536000 seconds
	// the minimal pay frequency is one week
	seconds := OneWeek * payFreq

	monthAPY := apy.Quo(sdk.NewDec(OneYear / int64(seconds)))

	i, err := apyTospy(monthAPY, uint64(seconds))
	if err != nil {
		return sdk.Dec{}, err
	}

	return i, nil
}
