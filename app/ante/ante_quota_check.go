package ante

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	vaultkeeper "github.com/joltify-finance/joltify_lending/x/vault/keeper"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
)

func getPools(ctx sdk.Context, keeper vaultkeeper.Keeper) []sdk.AccAddress {
	req := types.QueryLatestPoolRequest{}
	wctx := sdk.WrapSDKContext(ctx)

	resp, err := keeper.GetLastPool(wctx, &req)
	if err != nil || resp == nil || len(resp.Pools) != 2 {
		return nil
	}

	pool0 := resp.Pools[0].CreatePool.PoolAddr
	pool1 := resp.Pools[1].CreatePool.PoolAddr

	return []sdk.AccAddress{pool0, pool1}
}

type VaultQuotaDecorator struct {
	vk vaultkeeper.Keeper
}

func NewVaultQuotaDecorate(keeper vaultkeeper.Keeper) VaultQuotaDecorator {
	return VaultQuotaDecorator{
		vk: keeper,
	}
}

// QuotaCheck if quotaCheck return TRUE, it can continue process the tx
func (vd VaultQuotaDecorator) QuotaCheck(ctx sdk.Context, coins sdk.Coins) bool {
	targetQuota := vd.vk.TargetQuota(ctx)
	currentQuota, found := vd.vk.GetQuotaData(ctx)
	if !found {
		ctx.Logger().Error("the quota info cannot be found")
		return true
	}

	tempCurrent := sdk.NewCoins(currentQuota.CoinsSum...)

	coins.Sort()
	tempCurrent.Sort()
	targetQuota.Sort()

	afterTransfer := coins.Add(tempCurrent...)
	return afterTransfer.IsAllLTE(targetQuota)
}

func (vd VaultQuotaDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	msgs := tx.GetMsgs()
	for _, el := range msgs {
		switch msg := el.(type) {
		case *banktypes.MsgSend:
			pools := getPools(ctx, vd.vk)
			if len(pools) < 2 {
				return next(ctx, tx, simulate)
			}
			if msg.ToAddress == pools[0].String() || msg.ToAddress == pools[1].String() {
				if vd.QuotaCheck(ctx, msg.Amount) {
					return next(ctx, tx, simulate)
				}

				return ctx, errorsmod.Wrapf(
					vaultkeeper.ErrSuspend,
					"pool %v has reached the quota target",
					msg.ToAddress,
				)
			}

		default:
			return next(ctx, tx, simulate)
		}
	}
	return next(ctx, tx, simulate)
}
