package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/x/third_party/exchange/types"
)

// GetCampaignTradingRewardPoints fetches the trading reward points for a given account.
func (k *Keeper) GetCampaignTradingRewardPoints(ctx sdk.Context, account sdk.AccAddress) sdk.Dec {
	store := k.getStore(ctx)
	bz := store.Get(types.GetTradingRewardAccountPointsKey(account))
	if bz == nil {
		return sdk.ZeroDec()
	}
	return types.DecBytesToDec(bz)
}

// DeleteAccountCampaignTradingRewardPoints deletes the trading reward points for a given account.
func (k *Keeper) DeleteAccountCampaignTradingRewardPoints(ctx sdk.Context, account sdk.AccAddress) {
	store := k.getStore(ctx)
	store.Delete(types.GetTradingRewardAccountPointsKey(account))
}

// UpdateAccountCampaignTradingRewardPoints applies a point delta to the existing points.
func (k *Keeper) UpdateAccountCampaignTradingRewardPoints(
	ctx sdk.Context,
	account sdk.AccAddress,
	addedPoints sdk.Dec,
) {
	if addedPoints.IsZero() {
		return
	}

	accountPoints := k.GetCampaignTradingRewardPoints(ctx, account)
	accountPoints = accountPoints.Add(addedPoints)
	k.SetAccountCampaignTradingRewardPoints(ctx, account, accountPoints)
}

// SetAccountCampaignTradingRewardPoints sets the trading reward points for a given account.
func (k *Keeper) SetAccountCampaignTradingRewardPoints(ctx sdk.Context, account sdk.AccAddress, points sdk.Dec) {
	store := k.getStore(ctx)

	key := types.GetTradingRewardAccountPointsKey(account)
	bz := types.DecToDecBytes(points)
	store.Set(key, bz)
}

// GetAllTradingRewardCampaignAccountPoints gets the trading reward points for all accounts
func (k *Keeper) GetAllTradingRewardCampaignAccountPoints(ctx sdk.Context) (accountPoints []*types.TradingRewardCampaignAccountPoints) {
	accountPoints = make([]*types.TradingRewardCampaignAccountPoints, 0)

	appendPoints := func(points *types.TradingRewardAccountPoints) (stop bool) {
		accountPoints = append(accountPoints, &types.TradingRewardCampaignAccountPoints{
			Account: points.Account.String(),
			Points:  points.Points,
		})
		return false
	}

	k.IterateAccountCampaignTradingRewardPoints(ctx, appendPoints)
	return accountPoints
}

// GetAllAccountCampaignTradingRewardPointsWithTotalPoints gets the trading reward points for all accounts
func (k *Keeper) GetAllAccountCampaignTradingRewardPointsWithTotalPoints(ctx sdk.Context) (accountPoints []*types.TradingRewardAccountPoints, totalPoints sdk.Dec) {
	accountPoints = make([]*types.TradingRewardAccountPoints, 0)
	totalPoints = sdk.ZeroDec()

	appendPoints := func(points *types.TradingRewardAccountPoints) (stop bool) {
		accountPoints = append(accountPoints, points)
		totalPoints = totalPoints.Add(points.Points)
		return false
	}

	k.IterateAccountCampaignTradingRewardPoints(ctx, appendPoints)
	return accountPoints, totalPoints
}

// IterateAccountCampaignTradingRewardPoints iterates over the trading reward account points
func (k *Keeper) IterateAccountCampaignTradingRewardPoints(
	ctx sdk.Context,
	process func(*types.TradingRewardAccountPoints) (stop bool),
) {
	store := k.getStore(ctx)

	pointsStore := prefix.NewStore(store, types.TradingRewardAccountPointsPrefix)

	iterator := pointsStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		bz := iterator.Value()
		points := types.DecBytesToDec(bz)
		account := sdk.AccAddress(iterator.Key())

		accountPoints := &types.TradingRewardAccountPoints{
			Account: account,
			Points:  points,
		}
		if process(accountPoints) {
			return
		}
	}
}

// GetTotalTradingRewardPoints gets the total trading reward points
func (k *Keeper) GetTotalTradingRewardPoints(
	ctx sdk.Context,
) sdk.Dec {
	store := k.getStore(ctx)
	bz := store.Get(types.TradingRewardCampaignTotalPointsKey)
	if bz == nil {
		return sdk.ZeroDec()
	}
	return types.DecBytesToDec(bz)
}

// IncrementTotalTradingRewardPoints sets the total trading reward points
func (k *Keeper) IncrementTotalTradingRewardPoints(
	ctx sdk.Context,
	points sdk.Dec,
) {
	currPoints := k.GetTotalTradingRewardPoints(ctx)
	newPoints := currPoints.Add(points)
	k.SetTotalTradingRewardPoints(ctx, newPoints)
}

// SetTotalTradingRewardPoints sets the total trading reward points
func (k *Keeper) SetTotalTradingRewardPoints(
	ctx sdk.Context,
	points sdk.Dec,
) {
	store := k.getStore(ctx)
	bz := types.DecToDecBytes(points)
	store.Set(types.TradingRewardCampaignTotalPointsKey, bz)
}

// DeleteTotalTradingRewardPoints deletes the total trading reward points
func (k *Keeper) DeleteTotalTradingRewardPoints(
	ctx sdk.Context,
) {
	store := k.getStore(ctx)
	store.Delete(types.TradingRewardCampaignTotalPointsKey)
}

// PersistTradingRewardPoints persists the trading reward points
func (k *Keeper) PersistTradingRewardPoints(ctx sdk.Context, tradingRewards types.TradingRewardPoints) {
	totalTradingRewardPoints := sdk.ZeroDec()

	for _, account := range tradingRewards.GetSortedAccountKeys() {
		addr, _ := sdk.AccAddressFromBech32(account)
		accountTradingRewardPoints := tradingRewards[account]

		isRegisteredDMM := k.GetIsOptedOutOfRewards(ctx, addr)
		if isRegisteredDMM {
			continue
		}

		k.UpdateAccountCampaignTradingRewardPoints(ctx, addr, accountTradingRewardPoints)
		totalTradingRewardPoints = totalTradingRewardPoints.Add(accountTradingRewardPoints)
	}
	k.IncrementTotalTradingRewardPoints(ctx, totalTradingRewardPoints)
}
