package keeper

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/x/third_party/insurance/types"
)

// ExportNextShareDenomId returns the next share denom id
func (k *Keeper) ExportNextShareDenomId(ctx sdk.Context) uint64 {

	var shareDenomId uint64
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.GlobalShareDenomIdPrefixKey)
	if bz == nil {
		shareDenomId = 1
	} else {
		shareDenomId = sdk.BigEndianToUint64(bz)
	}
	return shareDenomId
}

func (k *Keeper) SetNextShareDenomId(ctx sdk.Context, shareDenomId uint64) {

	store := ctx.KVStore(k.storeKey)
	store.Set(types.GlobalShareDenomIdPrefixKey, sdk.Uint64ToBigEndian(shareDenomId))
}

// getNextShareDenomId returns the next share denom id and increase it
func (k *Keeper) getNextShareDenomId(ctx sdk.Context) uint64 {

	shareDenomId := k.ExportNextShareDenomId(ctx)
	k.SetNextShareDenomId(ctx, shareDenomId+1)
	return shareDenomId
}

// MintShareTokens mint share tokens to an address and increase total share variable of insurance fund
func (k *Keeper) MintShareTokens(ctx sdk.Context, fund *types.InsuranceFund, addr sdk.AccAddress, shares sdkmath.Int) (*types.InsuranceFund, error) {

	amount := sdk.Coins{sdk.NewCoin(fund.ShareDenom(), shares)}
	err := k.bankKeeper.MintCoins(ctx, types.ModuleName, amount)
	if err != nil {

		return fund, err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, amount)
	if err != nil {

		return fund, err
	}

	fund.AddTotalShare(shares)
	return fund, nil
}

// BurnShareTokens burn share tokens locked on insurance module
func (k *Keeper) BurnShareTokens(ctx sdk.Context, fund *types.InsuranceFund, shares sdkmath.Int) (*types.InsuranceFund, error) {

	shareAmount := sdk.Coins{sdk.NewCoin(fund.ShareDenom(), shares)}

	err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, shareAmount)
	if err != nil {

		return fund, err
	}

	fund.SubTotalShare(shares)
	return fund, nil
}
