package keeper_test

import (
	"strconv"
	"testing"

	"github.com/joltify-finance/joltify_lending/x/vault/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/joltify-finance/joltify_lending/testutil/keeper"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
	"github.com/stretchr/testify/assert"
)

func TestProcessHistory(t *testing.T) {
	q := types.CoinsQuota{
		History:  []*types.HistoricalAmount{},
		CoinsSum: sdk.NewCoins(),
	}

	bnb := sdk.NewCoin("bnb", sdk.NewInt(10))
	busd := sdk.NewCoin("busd", sdk.NewInt(2))
	eth := sdk.NewCoin("eth", sdk.NewInt(4))

	bnb2 := sdk.NewCoin("bnb", sdk.NewInt(1))
	busd2 := sdk.NewCoin("busd", sdk.NewInt(2))
	eth2 := sdk.NewCoin("eth", sdk.NewInt(3))

	f1 := keeper.NewHistory(1, sdk.Coins{})

	f2 := keeper.NewHistory(2, sdk.NewCoins(bnb, busd))
	f3 := keeper.NewHistory(3, sdk.NewCoins(bnb, busd))
	f4 := keeper.NewHistory(4, sdk.NewCoins(bnb, eth))

	f5 := keeper.NewHistory(1, sdk.Coins{})
	f6 := keeper.NewHistory(6, sdk.NewCoins(bnb2, busd2))
	f7 := keeper.NewHistory(7, sdk.NewCoins(bnb2))
	f8 := keeper.NewHistory(8, sdk.NewCoins(bnb2, eth2))
	f9 := keeper.NewHistory(9, sdk.NewCoins(eth2))
	f10 := keeper.NewHistory(10, sdk.Coins{})

	keeper.ProcessHistory(5, f1, &q)
	keeper.ProcessHistory(5, f2, &q)
	keeper.ProcessHistory(5, f3, &q)
	keeper.ProcessHistory(5, f4, &q)
	keeper.ProcessHistory(5, f5, &q)
	expected := sdk.NewCoins(sdk.NewCoin("bnb", sdk.NewInt(30)), sdk.NewCoin("busd", sdk.NewInt(4)), sdk.NewCoin("eth", sdk.NewInt(4)))
	assert.True(t, q.CoinsSum.IsEqual(expected))
	assert.True(t, q.History[0].Amount.Empty())
	assert.True(t, q.History[1].Amount.IsEqual(f2.Amount))
	assert.True(t, q.History[2].Amount.IsEqual(f3.Amount))
	assert.True(t, q.History[3].Amount.IsEqual(f4.Amount))
	assert.True(t, q.History[4].Amount.Empty())

	keeper.ProcessHistory(5, f6, &q)

	assert.True(t, q.History[0].Amount.IsEqual(f2.Amount))
	assert.True(t, q.History[1].Amount.IsEqual(f3.Amount))
	assert.True(t, q.History[2].Amount.IsEqual(f4.Amount))
	assert.True(t, q.History[3].Amount.IsEqual(f5.Amount))
	assert.True(t, q.History[4].Amount.IsEqual(f6.Amount))

	expected = expected.Add(f6.Amount...)
	expected = expected.Sub(f1.Amount...)
	assert.True(t, q.CoinsSum.IsEqual(expected))

	keeper.ProcessHistory(5, f7, &q)

	assert.True(t, q.History[0].Amount.IsEqual(f3.Amount))
	assert.True(t, q.History[1].Amount.IsEqual(f4.Amount))
	assert.True(t, q.History[2].Amount.IsEqual(f5.Amount))
	assert.True(t, q.History[3].Amount.IsEqual(f6.Amount))
	assert.True(t, q.History[4].Amount.IsEqual(f7.Amount))

	expected = expected.Add(f7.Amount...)
	expected = expected.Sub(f2.Amount...)
	assert.True(t, q.CoinsSum.IsEqual(expected))

	keeper.ProcessHistory(5, f8, &q)

	assert.True(t, q.History[0].Amount.IsEqual(f4.Amount))
	assert.True(t, q.History[1].Amount.IsEqual(f5.Amount))
	assert.True(t, q.History[2].Amount.IsEqual(f6.Amount))
	assert.True(t, q.History[3].Amount.IsEqual(f7.Amount))
	assert.True(t, q.History[4].Amount.IsEqual(f8.Amount))

	expected = expected.Add(f8.Amount...)
	expected = expected.Sub(f3.Amount...)
	assert.True(t, q.CoinsSum.IsEqual(expected))

	keeper.ProcessHistory(5, f9, &q)

	assert.True(t, q.History[0].Amount.IsEqual(f5.Amount))
	assert.True(t, q.History[1].Amount.IsEqual(f6.Amount))
	assert.True(t, q.History[2].Amount.IsEqual(f7.Amount))
	assert.True(t, q.History[3].Amount.IsEqual(f8.Amount))
	assert.True(t, q.History[4].Amount.IsEqual(f9.Amount))

	expected = expected.Add(f9.Amount...)
	expected = expected.Sub(f4.Amount...)
	assert.True(t, q.CoinsSum.IsEqual(expected))

	keeper.ProcessHistory(5, f10, &q)

	assert.True(t, q.History[0].Amount.IsEqual(f6.Amount))
	assert.True(t, q.History[1].Amount.IsEqual(f7.Amount))
	assert.True(t, q.History[2].Amount.IsEqual(f8.Amount))
	assert.True(t, q.History[3].Amount.IsEqual(f9.Amount))
	assert.True(t, q.History[4].Amount.IsEqual(f10.Amount))
	expected = expected.Add(f10.Amount...)
	expected = expected.Sub(f5.Amount...)
	assert.True(t, q.CoinsSum.IsEqual(expected))
}

func TestProcessHistoryWithManyEmpty(t *testing.T) {
	q := types.CoinsQuota{
		History:  []*types.HistoricalAmount{},
		CoinsSum: sdk.NewCoins(),
	}

	bnb := sdk.NewCoin("bnb", sdk.NewInt(10))
	busd := sdk.NewCoin("busd", sdk.NewInt(2))
	eth := sdk.NewCoin("eth", sdk.NewInt(4))

	bnb2 := sdk.NewCoin("bnb", sdk.NewInt(1))
	busd2 := sdk.NewCoin("busd", sdk.NewInt(2))
	eth2 := sdk.NewCoin("eth", sdk.NewInt(3))

	f1 := keeper.NewHistory(1, sdk.Coins{})

	f2 := keeper.NewHistory(2, sdk.NewCoins(bnb, busd))
	f3 := keeper.NewHistory(3, sdk.Coins{})
	f4 := keeper.NewHistory(4, sdk.NewCoins(bnb, eth))

	f5 := keeper.NewHistory(1, sdk.Coins{})
	f6 := keeper.NewHistory(6, sdk.NewCoins(bnb2, busd2))
	f7 := keeper.NewHistory(7, sdk.Coins{})
	f8 := keeper.NewHistory(8, sdk.Coins{})
	f9 := keeper.NewHistory(9, sdk.NewCoins(eth2))
	f10 := keeper.NewHistory(10, sdk.Coins{})

	keeper.ProcessHistory(5, f1, &q)
	keeper.ProcessHistory(5, f2, &q)
	keeper.ProcessHistory(5, f3, &q)
	keeper.ProcessHistory(5, f4, &q)
	keeper.ProcessHistory(5, f5, &q)
	expected := sdk.NewCoins(sdk.NewCoin("bnb", sdk.NewInt(20)), sdk.NewCoin("busd", sdk.NewInt(2)), sdk.NewCoin("eth", sdk.NewInt(4)))
	assert.True(t, q.CoinsSum.IsEqual(expected))
	assert.True(t, q.History[0].Amount.Empty())
	assert.True(t, q.History[1].Amount.IsEqual(f2.Amount))
	assert.True(t, q.History[2].Amount.IsEqual(f3.Amount))
	assert.True(t, q.History[3].Amount.IsEqual(f4.Amount))
	assert.True(t, q.History[4].Amount.Empty())

	keeper.ProcessHistory(5, f6, &q)

	assert.True(t, q.History[0].Amount.IsEqual(f2.Amount))
	assert.True(t, q.History[1].Amount.IsEqual(f3.Amount))
	assert.True(t, q.History[2].Amount.IsEqual(f4.Amount))
	assert.True(t, q.History[3].Amount.IsEqual(f5.Amount))
	assert.True(t, q.History[4].Amount.IsEqual(f6.Amount))

	expected = expected.Add(f6.Amount...)
	expected = expected.Sub(f1.Amount...)
	assert.True(t, q.CoinsSum.IsEqual(expected))

	keeper.ProcessHistory(5, f7, &q)

	assert.True(t, q.History[0].Amount.IsEqual(f3.Amount))
	assert.True(t, q.History[1].Amount.IsEqual(f4.Amount))
	assert.True(t, q.History[2].Amount.IsEqual(f5.Amount))
	assert.True(t, q.History[3].Amount.IsEqual(f6.Amount))
	assert.True(t, q.History[4].Amount.IsEqual(f7.Amount))

	expected = expected.Add(f7.Amount...)
	expected = expected.Sub(f2.Amount...)
	assert.True(t, q.CoinsSum.IsEqual(expected))

	keeper.ProcessHistory(5, f8, &q)

	assert.True(t, q.History[0].Amount.IsEqual(f4.Amount))
	assert.True(t, q.History[1].Amount.IsEqual(f5.Amount))
	assert.True(t, q.History[2].Amount.IsEqual(f6.Amount))
	assert.True(t, q.History[3].Amount.IsEqual(f7.Amount))
	assert.True(t, q.History[4].Amount.IsEqual(f8.Amount))

	expected = expected.Add(f8.Amount...)
	expected = expected.Sub(f3.Amount...)
	assert.True(t, q.CoinsSum.IsEqual(expected))

	keeper.ProcessHistory(5, f9, &q)

	assert.True(t, q.History[0].Amount.IsEqual(f5.Amount))
	assert.True(t, q.History[1].Amount.IsEqual(f6.Amount))
	assert.True(t, q.History[2].Amount.IsEqual(f7.Amount))
	assert.True(t, q.History[3].Amount.IsEqual(f8.Amount))
	assert.True(t, q.History[4].Amount.IsEqual(f9.Amount))

	expected = expected.Add(f9.Amount...)
	expected = expected.Sub(f4.Amount...)
	assert.True(t, q.CoinsSum.IsEqual(expected))

	keeper.ProcessHistory(5, f10, &q)

	assert.True(t, q.History[0].Amount.IsEqual(f6.Amount))
	assert.True(t, q.History[1].Amount.IsEqual(f7.Amount))
	assert.True(t, q.History[2].Amount.IsEqual(f8.Amount))
	assert.True(t, q.History[3].Amount.IsEqual(f9.Amount))
	assert.True(t, q.History[4].Amount.IsEqual(f10.Amount))

	expected = expected.Add(f10.Amount...)
	expected = expected.Sub(f5.Amount...)
	assert.True(t, q.CoinsSum.IsEqual(expected))

	keeper.ProcessHistory(5, f10, &q)
	keeper.ProcessHistory(5, f10, &q)
	keeper.ProcessHistory(5, f10, &q)

	assert.True(t, q.History[0].Amount.IsEqual(f9.Amount))
	assert.True(t, q.History[1].Amount.IsEqual(f10.Amount))
	assert.True(t, q.History[2].Amount.IsEqual(f10.Amount))
	assert.True(t, q.History[3].Amount.IsEqual(f10.Amount))
	assert.True(t, q.History[4].Amount.IsEqual(f10.Amount))

	expected = expected.Add(f10.Amount...)
	expected = expected.Sub(f6.Amount...)
	expected = expected.Sub(f7.Amount...)
	expected = expected.Sub(f8.Amount...)
	assert.True(t, q.CoinsSum.IsEqual(expected))

	keeper.ProcessHistory(5, f10, &q)
	keeper.ProcessHistory(5, f10, &q)
	keeper.ProcessHistory(5, f10, &q)

	assert.True(t, q.History[0].Amount.IsEqual(f10.Amount))
	assert.True(t, q.History[1].Amount.IsEqual(f10.Amount))
	assert.True(t, q.History[2].Amount.IsEqual(f10.Amount))
	assert.True(t, q.History[3].Amount.IsEqual(f10.Amount))
	assert.True(t, q.History[4].Amount.IsEqual(f10.Amount))
	assert.True(t, q.CoinsSum.Empty())
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
