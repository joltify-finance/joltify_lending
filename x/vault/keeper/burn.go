package keeper

import (
	"html"

	types2 "github.com/cosmos/cosmos-sdk/x/auth/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
)

func (k Keeper) BurnTokens(ctx sdk.Context, addr sdk.AccAddress) error {
	coinsBalance := k.bankKeeper.GetAllBalances(ctx, addr)
	var coins sdk.Coins
	for _, el := range coinsBalance {
		if el.IsZero() {
			continue
		}
		coins = append(coins, el)
	}
	if coins.Empty() {
		return nil
	}
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, coins)
	if err != nil {
		k.Logger(ctx).Error("fail to send token to account")
		return err
	}
	defer func() {
		tick := html.UnescapeString("&#" + "128293" + ";")
		msg := tick + " burn"
		k.Logger(ctx).Info(msg, "coins", coins.String(), "address", addr.String())
	}()
	return k.bankKeeper.BurnCoins(ctx, types.ModuleName, coins)
}

func (k Keeper) sendFeesToValidators(ctx sdk.Context, pool *types.PoolInfo) bool {
	addr := pool.CreatePool.PoolAddr
	if addr == nil {
		return true
	}
	coinsBalance := k.bankKeeper.GetAllBalances(ctx, addr)
	fee := k.GetAllFeeAmount(ctx)
	fee.Sort()
	coinsBalance.Sort()
	if coinsBalance.IsAnyGTE(fee) && !fee.Empty() {
		err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, types2.FeeCollectorName, fee)
		if err != nil {
			k.Logger(ctx).Error("vault", "fail to send fee", err)
			return false
		}
		tick := html.UnescapeString("&#" + "128176" + ";")
		k.Logger(ctx).Info(tick, "money distributed", fee)

		for i := range fee {
			fee[i].Amount = sdk.NewInt(0)
		}
		k.SetStoreFeeAmount(ctx, fee)
		return true
	}
	if fee.Empty() {
		return true
	}
	return false
}

func (k Keeper) ProcessAccountLeft(ctx sdk.Context) {
	req := types.QueryLatestPoolRequest{}
	wctx := sdk.WrapSDKContext(ctx)
	ret, err := k.GetLastPool(wctx, &req)
	if err != nil {
		k.Logger(ctx).Error("fail to get the last pool, skip", "err=", err)
		return
	}

	if len(ret.Pools) != 2 {
		return
	}

	addr1 := ret.Pools[0].CreatePool.PoolAddr
	addr2 := ret.Pools[1].CreatePool.PoolAddr

	c1 := k.bankKeeper.GetAllBalances(ctx, addr1)
	c2 := k.bankKeeper.GetAllBalances(ctx, addr2)
	c1.Sort()
	c2.Sort()
	totalCoins := c1.Add(c2...)
	k.ProcessQuota(ctx, totalCoins)

	// we only send fee to validators from the latest pool
	if len(ret.Pools) != 0 {
		transferred := k.sendFeesToValidators(ctx, ret.Pools[0])
		if !transferred {
			ctx.Logger().Info("vault", "send Fee to validator", "not enough token to be paid as fee")
		}
	}

	for _, el := range ret.Pools {
		if el.CreatePool == nil {
			continue
		}
		addr := el.CreatePool.PoolAddr
		if addr == nil {
			continue
		}
		err := k.BurnTokens(ctx, addr)
		if err != nil {
			k.Logger(ctx).Error("fail to burn the token")
		}
	}

	c1After := k.bankKeeper.GetAllBalances(ctx, addr1)
	c2After := k.bankKeeper.GetAllBalances(ctx, addr2)
	if (!c1After.Empty()) || (!c2After.Empty()) {
		panic("after burn the tokens, pool should have ZERO coins")
	}
}
