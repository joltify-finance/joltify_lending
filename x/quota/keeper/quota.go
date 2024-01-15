package keeper

import (
	"errors"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/quota/types"
)

// SetQuotaData set a specific createPool in the store from its index
func (k Keeper) SetQuotaData(ctx sdk.Context, coinsQuota types.CoinsQuota) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.QuotaKey))
	b := k.cdc.MustMarshal(&coinsQuota)
	store.Set(types.KeyPrefix(coinsQuota.ModuleName), b)
}

// GetQuotaData returns a createPool from its index
func (k Keeper) GetQuotaData(ctx sdk.Context, moduleName string) (val types.CoinsQuota, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.QuotaKey))
	b := store.Get(types.KeyPrefix(moduleName))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) WhetherOnwhitelist(ctx sdk.Context, moduleName, sender string) bool {
	params := k.GetParams(ctx)
	for _, el := range params.Whitelist {
		if el.ModuleName == moduleName {
			for _, el2 := range el.AddressList {
				if el2 == sender {
					return true
				}
			}
		}
	}
	return false
}

// GetAllQuota returns all quota
func (k Keeper) GetAllQuota(ctx sdk.Context) (list []types.CoinsQuota) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.QuotaKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CoinsQuota
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}
	return
}

func ProcessHistory(newItem *types.HistoricalAmount, coinsQuota *types.CoinsQuota) *types.CoinsQuota {
	if len(coinsQuota.History) >= types.MAXHISTORY {
		return nil
	}
	coinsQuota.History = append(coinsQuota.History, newItem)
	coinsQuota.CoinsSum = coinsQuota.CoinsSum.Add(newItem.Amount...)
	return coinsQuota
}

func (k Keeper) UpdateQuota(ctx sdk.Context, coins sdk.Coins, ibcSeq uint64, moduleName string) error {
	var targetQuota sdk.Coins

	params := k.GetParams(ctx)
	for _, el := range params.Targets {
		if el.ModuleName == moduleName {
			targetQuota = el.CoinsSum
			break
		}
	}

	if targetQuota.Empty() {
		return errors.New("no quota for this module")
	}

	currentQuota, found := k.GetQuotaData(ctx, moduleName)
	if !found {
		currentQuota.History = []*types.HistoricalAmount{}
		currentQuota.ModuleName = moduleName
		currentQuota.CoinsSum = sdk.NewCoins()
	}

	coins = coins.Sort()

	ret := coins.DenomsSubsetOf(targetQuota)
	if !ret {
		return errors.New("some coins cannot be found in target")
	}

	newRecord := types.HistoricalAmount{
		Amount:      coins,
		BlockHeight: ctx.BlockHeight(),
		IbcSequence: ibcSeq,
	}
	newQuota := ProcessHistory(&newRecord, &currentQuota)
	if newQuota == nil {
		return errors.New("quota exceeded")
	}

	allGT := targetQuota.IsAllGTE(newQuota.CoinsSum)
	if !allGT {
		return errors.New("quota exceeded")
	}
	k.SetQuotaData(ctx, *newQuota)
	return nil
}

func (k Keeper) RevokeHistory(ctx sdk.Context, moduleName string, seq uint64) {
	currentQuota, found := k.GetQuotaData(ctx, moduleName)
	if !found {
		return
	}

	for i, el := range currentQuota.History {
		if el.IbcSequence == seq {
			currentQuota.History = append(currentQuota.History[:i], currentQuota.History[i+1:]...)
			currentQuota.CoinsSum = currentQuota.CoinsSum.Sub(el.Amount...)
			k.SetQuotaData(ctx, currentQuota)
			return
		}
	}
	ctx.Logger().Error("cannot find the seq in history", "seq", seq)
}

func (k Keeper) BlockUpdateQuota(ctx sdk.Context) {
	lengthMap := make(map[string]uint32)
	params := k.GetParams(ctx)
	for _, el := range params.Targets {
		lengthMap[el.ModuleName] = el.HistoryLength
	}

	allQuota := k.GetAllQuota(ctx)

	for _, eachQuota := range allQuota {
		if len(eachQuota.History) == 0 {
			continue
		}

		maxHistoryLength, ok := lengthMap[eachQuota.ModuleName]
		if !ok {
			ctx.Logger().Error("cannot find history length for module " + eachQuota.ModuleName)
			continue
		}

		firstEntry := eachQuota.History[0]
		if ctx.BlockHeight()-firstEntry.BlockHeight > int64(maxHistoryLength) {
			if len(eachQuota.History) == 1 {
				eachQuota.History = []*types.HistoricalAmount{}
				eachQuota.CoinsSum = eachQuota.CoinsSum.Sub(firstEntry.Amount...)
				k.SetQuotaData(ctx, eachQuota)
				continue
			}
			eachQuota.History = eachQuota.History[1:]
			eachQuota.CoinsSum = eachQuota.CoinsSum.Sub(firstEntry.Amount...)
			k.SetQuotaData(ctx, eachQuota)
		}
	}
}
