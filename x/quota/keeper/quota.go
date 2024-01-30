package keeper

import (
	"errors"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/quota/types"
)

// SetQuotaData set a specific quota for the module
func (k Keeper) SetQuotaData(ctx sdk.Context, coinsQuota types.CoinsQuota) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.QuotaKey))
	b := k.cdc.MustMarshal(&coinsQuota)
	store.Set(types.KeyPrefix(coinsQuota.ModuleName), b)
}

// GetQuotaData returns the quota for a module
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

func (k Keeper) WhetherOnBanlist(ctx sdk.Context, moduleName, sender string) bool {
	params := k.GetParams(ctx)
	for _, el := range params.Banlist {
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

// GetPreAccountQuotaData returns the quota for a given account
func (k Keeper) getAccountQuotaData(ctx sdk.Context, moduleName, accAddr string) (val types.AccQuota, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.QuotaAccKey))
	key := types.KeyPrefix(moduleName + accAddr)
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// setAccQuotaData set coin quota for a specific account
func (k Keeper) setAccQuotaData(ctx sdk.Context, moduleName, accAddress string, accQuota types.AccQuota) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.QuotaAccKey))
	b := k.cdc.MustMarshal(&accQuota)
	store.Set(types.KeyPrefix(moduleName+accAddress), b)
}

func ProcessHistory(newItem *types.HistoricalAmount, coinsQuota *types.CoinsQuota) *types.CoinsQuota {
	if len(coinsQuota.History) >= types.MAXHISTORY {
		return nil
	}
	coinsQuota.History = append(coinsQuota.History, newItem)
	coinsQuota.CoinsSum = coinsQuota.CoinsSum.Add(newItem.Amount...)
	return coinsQuota
}

func (k Keeper) updatePerAccountQuota(ctx sdk.Context, accTargets []*types.Target, coins sdk.Coins, sender string, moduleName string) error {
	var targetQuota sdk.Coins
	var targetHistoryLength int64

	for _, el := range accTargets {
		if el.ModuleName == moduleName {
			targetQuota = el.CoinsSum
			targetHistoryLength = el.HistoryLength
			break
		}
	}
	if targetQuota.Empty() {
		return errors.New("no account quota for this module")
	}

	accQuota, found := k.getAccountQuotaData(ctx, moduleName, sender)
	if !found {
		if coins.IsAnyGT(targetQuota) {
			return types.AccErrQuotaExceed
		}
		el := types.AccQuota{
			ctx.BlockHeight(),
			ctx.BlockHeight(),
			coins,
		}
		k.setAccQuotaData(ctx, moduleName, sender, el)
		return nil
	}

	if ctx.BlockHeight()-accQuota.LastUpdateHeight > targetHistoryLength {
		accQuota.CoinsSum = sdk.NewCoins()
		accQuota.BlockHeight = ctx.BlockHeight()
	}

	tempsum := accQuota.CoinsSum.Add(coins...)
	if tempsum.IsAnyGT(targetQuota) {
		return types.AccErrQuotaExceed
	}
	accQuota.CoinsSum = tempsum
	accQuota.LastUpdateHeight = ctx.BlockHeight()
	k.setAccQuotaData(ctx, moduleName, sender, accQuota)
	return nil
}

func (k Keeper) UpdateQuota(ctx sdk.Context, coins sdk.Coins, sender string, ibcSeq uint64, moduleName string) error {
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

	coins = coins.Sort()
	err := k.updatePerAccountQuota(ctx, params.PerAccounttargets, coins, sender, moduleName)
	if err != nil {
		if err.Error() != "no account quota for this module" {
			return err
		}
	}

	currentQuota, found := k.GetQuotaData(ctx, moduleName)
	if !found {
		currentQuota.History = []*types.HistoricalAmount{}
		currentQuota.ModuleName = moduleName
		currentQuota.CoinsSum = sdk.NewCoins()
	}

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
		return types.ErrQuotaExceed
	}

	allGT := targetQuota.IsAllGTE(newQuota.CoinsSum)
	if !allGT {
		return types.ErrQuotaExceed
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
	lengthMap := make(map[string]int64)
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
		if ctx.BlockHeight()-firstEntry.BlockHeight > maxHistoryLength {
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
