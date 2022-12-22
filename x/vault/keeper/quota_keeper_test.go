package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/joltify-finance/joltify_lending/testutil/keeper"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
	"github.com/stretchr/testify/assert"
)

func TestQuotaCheck(t *testing.T) {
	app, ctx := keepertest.SetupVaultApp(t)

	params := app.VaultKeeper.GetParams(ctx)

	testValidators, creators := generateNValidators(t, 4)

	p1 := types.PoolProposal{PoolAddr: creators[0], Nodes: []sdk.AccAddress{creators[0], creators[1], creators[2]}}
	p2 := types.PoolProposal{PoolAddr: creators[1], Nodes: []sdk.AccAddress{creators[0], creators[1], creators[2]}}

	position1 := ctx.BlockHeight() - params.BlockChurnInterval + 1
	position2 := ctx.BlockHeight() - params.BlockChurnInterval*2 + 1

	createPool := types.CreatePool{
		BlockHeight: strconv.FormatInt(position1, 10),
		Validators:  testValidators,
		Proposal:    []*types.PoolProposal{&p1, &p1, &p1},
	}
	app.VaultKeeper.SetCreatePool(ctx, createPool)

	createPool = types.CreatePool{
		BlockHeight: strconv.FormatInt(position2, 10),
		Validators:  testValidators,
		Proposal:    []*types.PoolProposal{&p2, &p2, &p2},
	}
	app.VaultKeeper.SetCreatePool(ctx, createPool)
	q := types.CoinsQuota{
		History:  []*types.HistoricalAmount{},
		CoinsSum: sdk.NewCoins(),
	}
	app.VaultKeeper.SetQuotaData(ctx, q)

	coins := params.GetTargetQuota()
	ret := app.VaultKeeper.QuotaCheck(ctx, coins)
	assert.True(t, ret)

	q.CoinsSum = coins
	app.VaultKeeper.SetQuotaData(ctx, q)
	ret = app.VaultKeeper.QuotaCheck(ctx, coins)
	assert.False(t, ret)
	t1 := sdk.NewCoins(sdk.NewCoin("abnb", sdk.NewInt(100)))

	q.CoinsSum = q.CoinsSum.Sub(t1)
	app.VaultKeeper.SetQuotaData(ctx, q)
	ret = app.VaultKeeper.QuotaCheck(ctx, t1)
	assert.True(t, ret)

	tEth := sdk.NewCoins(sdk.NewCoin("aeth", sdk.NewInt(1)))
	ret = app.VaultKeeper.QuotaCheck(ctx, tEth)
	assert.False(t, ret)

	tEthandBnb := sdk.NewCoins(sdk.NewCoin("aeth", sdk.NewInt(1)), t1[0])
	ret = app.VaultKeeper.QuotaCheck(ctx, tEthandBnb)
	assert.False(t, ret)

	t2 := t1.Add(sdk.NewCoin("abnb", sdk.NewInt(1)))
	ret = app.VaultKeeper.QuotaCheck(ctx, t2)
	assert.False(t, ret)
}

func TestProcessEvent(t *testing.T) {
	app, ctx := keepertest.SetupVaultApp(t)

	params := app.VaultKeeper.GetParams(ctx)
	params.BlockChurnInterval = 20
	app.VaultKeeper.SetParams(ctx, params)

	testValidators, creators := generateNValidators(t, 4)

	p1 := types.PoolProposal{PoolAddr: creators[0], Nodes: []sdk.AccAddress{creators[0], creators[1], creators[2]}}
	p2 := types.PoolProposal{PoolAddr: creators[1], Nodes: []sdk.AccAddress{creators[0], creators[1], creators[2]}}

	position1 := ctx.BlockHeight() - params.BlockChurnInterval + 1
	position2 := ctx.BlockHeight() - params.BlockChurnInterval*2 + 1

	createPool := types.CreatePool{
		BlockHeight: strconv.FormatInt(position1, 10),
		Validators:  testValidators,
		Proposal:    []*types.PoolProposal{&p1, &p1, &p1},
	}
	app.VaultKeeper.SetCreatePool(ctx, createPool)

	createPool = types.CreatePool{
		BlockHeight: strconv.FormatInt(position2, 10),
		Validators:  testValidators,
		Proposal:    []*types.PoolProposal{&p2, &p2, &p2},
	}
	app.VaultKeeper.SetCreatePool(ctx, createPool)

	sendToken := sdk.NewCoins(sdk.NewCoin("abnb", sdk.NewInt(100)), sdk.NewCoin("aeth", sdk.NewInt(222)))

	app.VaultKeeper.ProcessQuota(ctx, sendToken)
	quota, found := app.VaultKeeper.GetQuotaData(ctx)
	assert.True(t, found)
	assert.Equal(t, len(quota.History), 1)
	assert.True(t, quota.CoinsSum.IsEqual(sendToken))
}
