package keeper

import (
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/mint/types"
)

//func (k Keeper) FirstDist(ctx sdk.Context, pa types.Params) error {
//	newCoins := sdk.NewCoins(sdk.NewCoin("ujolt", pa.CommunityProvisions.TruncateInt()))
//	err := k.bankKeeper.MintCoins(ctx, types.ModuleName, newCoins)
//	if err != nil {
//		return err
//	}
//	err = k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, incentivetypes.ModuleName, newCoins)
//	if err != nil {
//		return err
//	}
//	senderAddr := k.accountKeeper.GetModuleAddress(incentivetypes.JoltIncentiveMacc)
//	ctx.Logger().Info("the incentive account is %v", senderAddr.String())
//
//	return nil
//}

// SetDistInfo sets the historical distribution info
func (k Keeper) SetDistInfo(ctx sdk.Context, h types.HistoricalDistInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FirstDistTime))
	b := k.cdc.MustMarshal(&h)
	store.Set(types.KeyPrefix("history"), b)
}

// GetDistInfo returns a createPool from its index
func (k Keeper) GetDistInfo(ctx sdk.Context) (h types.HistoricalDistInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FirstDistTime))

	b := store.Get(types.KeyPrefix("history"))
	if b == nil {
		panic("fail to get the history info")
	}

	k.cdc.MustUnmarshal(b, &h)
	return h
}

func (k Keeper) mintCoinsAndDistribute(ctx sdk.Context, pa types.Params) error {
	newCoins := sdk.NewCoins(sdk.NewCoin("ujolt", pa.CurrentProvisions.TruncateInt()))
	err := k.bankKeeper.MintCoins(ctx, types.ModuleName, newCoins)
	if err != nil {
		return err
	}

	amountToCommunity := pa.CurrentProvisions.Mul(sdk.MustNewDecFromStr("0.2"))
	communityCoins := sdk.NewCoins(sdk.NewCoin("ujolt", amountToCommunity.TruncateInt()))
	feeCollector := newCoins.Sub(communityCoins)

	err = k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, k.feeCollectorName, feeCollector)
	if err != nil {
		return err
	}

	addr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	err = k.distributionKeeper.FundCommunityPool(ctx, communityCoins, addr)
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMint,
			sdk.NewAttribute(types.AttributeKeyEachProvisions, pa.CurrentProvisions.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, newCoins.String()),
		),
	)

	return nil
}

func (k Keeper) DoDistribute(ctx sdk.Context) {
	h := k.GetDistInfo(ctx)
	pa := k.GetParams(ctx)
	if pa.CurrentProvisions.TruncateInt().IsZero() {
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
		truncatedDelta = delta.Truncate(time.Second)
		if truncatedDelta < time.Second*60 {
			return
		}
	default:
		panic("invalid unit")
	}

	err := k.mintCoinsAndDistribute(ctx, pa)
	if err != nil {
		ctx.Logger().Error("fail to mint the token", "mint", err.Error())
	}
	h.PayoutTime = currentTime
	h.DistributedRound++
	if h.DistributedRound > 0 && h.DistributedRound%pa.HalfCount == 0 {
		pa.CurrentProvisions = pa.CurrentProvisions.QuoTruncate(sdk.MustNewDecFromStr("2"))
	}
	k.SetDistInfo(ctx, h)
	k.SetParams(ctx, pa)

}
