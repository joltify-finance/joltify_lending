package keeper

import (
	"math/big"
	"strings"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

// UpdateIncentive updates the incentive for the pool to the incentive module
func (k Keeper) UpdateIncentive(ctx context.Context, poolInfo types.PoolInfo) {
	poolIndex := poolInfo.Index
	totalBorrowed := poolInfo.BorrowedAmount

	pa := k.GetParams(ctx)

	var conversion math.Int
	pooldemos := strings.Split(poolInfo.BorrowedAmount.Denom, "-")
	for _, market := range pa.Markets {
		if market.GetDenom() == pooldemos[1] {
			c := market.GetConversionFactor()
			conversion = sdk.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(c)), nil))
			break
		}
	}

	for _, el := range pa.Incentives {
		if el.Poolid == poolIndex {
			// as the spy is 1.XXXX, so we need to minus 1
			spy := sdk.MustNewDecFromStr(el.Spy).Sub(sdkmath.LegacyOneDec())
			joltM, err := k.priceFeedKeeper.GetCurrentPrice(ctx, "jolt:usd")
			if err != nil {
				ctx.Logger().Error("cannot get jolt price", "error", err)
				return
			}

			borrowedDec := sdk.NewDecFromInt(totalBorrowed.Amount)
			incentiveJolt := borrowedDec.Mul(spy).Mul(sdk.NewDecFromInt(sdk.NewIntFromUint64(types.JOLTPRECISION))).Quo(sdk.NewDecFromInt(conversion)).Quo(joltM.Price)

			incentiveCoin := sdk.NewCoins(sdk.NewCoin("ujolt", incentiveJolt.TruncateInt()))

			k.incentivekeeper.SetSPVRewardTokens(ctx, poolIndex, incentiveCoin)
			return
		}
		continue
	}
}
