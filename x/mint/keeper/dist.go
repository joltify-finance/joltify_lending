package keeper

import (
	"context"
	"time"

	"github.com/joltify-finance/joltify_lending/client"

	sdkmath "cosmossdk.io/math"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/mint/types"
	incentivetypes "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"
)

const MAXMINT = 100000000

func (k Keeper) FirstDist(ctx context.Context) error {
	if client.MAINNETFLAG == "unittest" {
		return nil
	}

	base := sdkmath.NewInt(1000000)
	firstDropIncentive := sdkmath.NewInt(100000)

	c1 := sdk.NewCoin("ujolt", firstDropIncentive.Mul(base))
	incentiveReceived := sdk.NewCoins(c1)

	firstDropCommunity := sdkmath.NewInt(20000000)
	c2 := sdk.NewCoin("ujolt", firstDropCommunity.Mul(base))

	minttedCoin := incentiveReceived.Add(c2)

	err := k.bankKeeper.MintCoins(ctx, types.ModuleName, minttedCoin)
	if err != nil {
		return err
	}
	err = k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, incentivetypes.ModuleName, incentiveReceived)
	if err != nil {
		return err
	}

	err = k.distributionKeeper.FundCommunityPool(ctx, sdk.NewCoins(c2), k.accountKeeper.GetModuleAddress(types.ModuleName))
	if err != nil {
		panic("fail to fund the community pool" + err.Error())
	}

	return nil
}

// SetDistInfo sets the historical distribution info
func (k Keeper) SetDistInfo(rctx context.Context, h types.HistoricalDistInfo) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FirstDistTime))
	b := k.cdc.MustMarshal(&h)
	store.Set(types.KeyPrefix("history"), b)
}

// GetDistInfo returns a createPool from its index
func (k Keeper) GetDistInfo(rctx context.Context) (h types.HistoricalDistInfo) {
	ctx := sdk.UnwrapSDKContext(rctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FirstDistTime))

	b := store.Get(types.KeyPrefix("history"))
	if b == nil {
		panic("fail to get the history info")
	}

	k.cdc.MustUnmarshal(b, &h)
	return h
}

func (k Keeper) mintCoinsAndDistribute(rctx context.Context, pa types.Params, delta time.Duration) (sdk.Coins, error) {
	ctx := sdk.UnwrapSDKContext(rctx)
	truncatedDelta := int64(delta.Truncate(time.Second).Seconds())
	interestFactor := CalculateInterestFactor(pa.NodeSPY, sdkmath.NewInt(truncatedDelta))

	totalBoned, err := k.stakingKeeper.TotalBondedTokens(rctx)
	if err != nil {
		return sdk.Coins{}, err
	}

	minttedAmt := interestFactor.Sub(sdkmath.LegacyOneDec()).MulInt(sdkmath.NewIntFromBigInt(totalBoned.BigInt()))
	minttedCoins := sdk.NewCoins(sdk.NewCoin("ujolt", minttedAmt.TruncateInt()))

	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, minttedCoins)
	if err != nil {
		return sdk.Coins{}, err
	}

	err = k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, k.feeCollectorName, minttedCoins)
	if err != nil {
		return sdk.Coins{}, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMint,
			sdk.NewAttribute(types.AttributeKeySPY, pa.NodeSPY.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, minttedCoins.String()),
		),
	)

	return minttedCoins, nil
}

func (k Keeper) DoDistribute(rctx context.Context) {
	ctx := sdk.UnwrapSDKContext(rctx)
	h := k.GetDistInfo(ctx)
	pa := k.GetParams(ctx)
	if pa.NodeSPY.IsZero() {
		return
	}

	base := sdkmath.NewInt(1000000)
	maxMint := base.Mul(sdkmath.NewInt(MAXMINT))
	if h.TotalMintCoins.AmountOf("ujolt").GTE(maxMint) {
		return
	}

	previousDistTime := h.PayoutTime
	currentTime := ctx.BlockTime()
	delta := currentTime.Sub(previousDistTime)

	var truncatedDelta time.Duration
	switch pa.Unit {
	case "hour":
		truncatedDelta = delta.Truncate(time.Minute)
		if truncatedDelta < 60 {
			return
		}

	case "minute":
		// we distribute the token every 5 minutes
		truncatedDelta = delta.Truncate(time.Second)
		if truncatedDelta < time.Second*60 {
			return
		}
	default:
		panic("invalid unit")
	}

	minttedAmt, err := k.mintCoinsAndDistribute(ctx, pa, delta)
	if err != nil {
		ctx.Logger().Error("fail to mint the token", "mint", err.Error())
	}
	h.PayoutTime = currentTime
	h.TotalMintCoins = h.TotalMintCoins.Add(minttedAmt.Sort()...)
	k.SetDistInfo(ctx, h)
}
