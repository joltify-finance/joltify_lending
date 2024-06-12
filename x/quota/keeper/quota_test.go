package keeper_test

import (
	"fmt"
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/joltify-finance/joltify_lending/testutil/keeper"
	"github.com/joltify-finance/joltify_lending/x/quota/types"
	"github.com/stretchr/testify/require"
)

func createNQuota(moduleName string, n int) types.CoinsQuota {
	ht := types.HistoricalAmount{}

	sum := sdk.NewCoins()
	for i := 2; i < n+2; i++ {
		ht.BlockHeight = int64(i)
		tokenName := "test" + strconv.Itoa(i)
		t := sdk.NewCoin(tokenName, sdk.NewInt(int64(i)))
		ht.Amount = ht.Amount.Add(t)
		sum = sum.Add(t)
	}
	mockQuotaCoins := types.CoinsQuota{
		ModuleName: moduleName,
		History:    []*types.HistoricalAmount{&ht},
		CoinsSum:   sum,
	}

	return mockQuotaCoins
}

func TestQuotaGet(t *testing.T) {
	keeper, ctx := keepertest.QuotaKeeper(t)
	items := createNQuota("testmodule1", 10)
	keeper.SetQuotaData(ctx, items)
	targetItems, found := keeper.GetQuotaData(ctx, "testmodule1")
	require.True(t, found)
	require.True(t, targetItems.CoinsSum.IsEqual(items.CoinsSum))
	for _, el := range targetItems.History {
		require.True(t, el.Amount.IsEqual(items.History[0].Amount))
	}
}

func TestQuotaGetAll(t *testing.T) {
	keeper, ctx := keepertest.QuotaKeeper(t)
	items1 := createNQuota("testmodule1", 10)
	items2 := createNQuota("testmodule2", 10)
	keeper.SetQuotaData(ctx, items1)
	keeper.SetQuotaData(ctx, items2)
	ret := keeper.GetAllQuota(ctx)

	require.Len(t, ret, 2)

	var i1, i2 types.CoinsQuota
	if ret[0].ModuleName == "testmodule1" {
		i1 = ret[0]
		i2 = ret[1]
	} else {
		i1 = ret[1]
		i2 = ret[0]
	}

	require.True(t, i1.CoinsSum.IsEqual(items1.CoinsSum))
	for _, el := range i1.History {
		require.True(t, el.Amount.IsEqual(items1.History[0].Amount))
	}

	require.True(t, i2.CoinsSum.IsEqual(items2.CoinsSum))
	for _, el := range i2.History {
		require.True(t, el.Amount.IsEqual(items2.History[0].Amount))
	}
}

// NewParams creates a new Params instance
func testParams() types.Params {
	// the coin list is the amount of USD for the given token, 100jolt means 100 USD value of jolt
	quota, err := sdk.ParseCoinsNormalized("100000ujolt,1000000usdt")
	if err != nil {
		panic(err)
	}

	quotaAcc, err := sdk.ParseCoinsNormalized("10000000ujolt,100000000usdt")
	if err != nil {
		panic(err)
	}

	targets := types.Target{
		ModuleName:    "ibc",
		CoinsSum:      quota,
		HistoryLength: 512,
	}

	targets2 := types.Target{
		ModuleName:    "bridge",
		CoinsSum:      quota,
		HistoryLength: 512,
	}

	targetsAcc := types.Target{
		ModuleName:    "ibc",
		CoinsSum:      quotaAcc,
		HistoryLength: 512,
	}

	targets2Acc := types.Target{
		ModuleName:    "bridge",
		CoinsSum:      quotaAcc,
		HistoryLength: 512,
	}

	return types.Params{Targets: []*types.Target{&targets, &targets2}, PerAccounttargets: []*types.Target{&targetsAcc, &targets2Acc}}
}

func TestUpdateQuota(t *testing.T) {
	keeper, ctx := keepertest.QuotaKeeper(t)
	items1 := createNQuota("testmodule1", 10)
	keeper.SetQuotaData(ctx, items1)

	testcoins := sdk.NewCoins(sdk.NewCoin("testAcc", sdk.NewInt(100)))
	err := keeper.UpdateQuota(ctx, testcoins, "testaddr", 1, "testmodule1")
	require.Error(t, err, "no quota for this module")
	keeper.SetParams(ctx, testParams())

	testcoins = sdk.NewCoins(sdk.NewCoin("testa", sdk.NewInt(100)))
	err = keeper.UpdateQuota(ctx, testcoins, "testaddr", 1, "testmodule1")
	require.Error(t, err, "no quota for this module")

	testcoins = sdk.NewCoins(sdk.NewCoin("testa", sdk.NewInt(100)))
	err = keeper.UpdateQuota(ctx, testcoins, "testaddr", 1, "ibc")
	require.Error(t, err, "quota not found")

	testcoins = sdk.NewCoins(sdk.NewCoin("invalid", sdk.NewInt(1)))
	err = keeper.UpdateQuota(ctx, testcoins, "testaddr", 1, "ibc")
	require.Error(t, err, "some coins cannot be found in target")

	newcoin := sdk.NewCoin("usdt", sdk.NewInt(1))
	testcoins = sdk.NewCoins(newcoin)

	before, found := keeper.GetQuotaData(ctx, "ibc")
	require.True(t, true)

	err = keeper.UpdateQuota(ctx, testcoins, "testaddr", 1, "ibc")
	require.NoError(t, err)

	after, found := keeper.GetQuotaData(ctx, "ibc")
	require.True(t, found)

	delta := after.CoinsSum.Sub(before.CoinsSum...)
	require.True(t, delta.IsEqual(sdk.NewCoins(newcoin)))

	after.History[0].Amount = sdk.NewCoins(newcoin)
	after.History[0].BlockHeight = ctx.BlockHeight()

	// we add some records
	for i := 0; i < 20; i++ {
		newcoin := sdk.NewCoin("usdt", sdk.NewInt(int64(1+i)))
		testcoins = sdk.NewCoins(newcoin)
		ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)
		err = keeper.UpdateQuota(ctx, testcoins, "testaddr", 1, "ibc")
		require.NoError(t, err)
	}

	after, found = keeper.GetQuotaData(ctx, "ibc")
	left := testParams().Targets[0].CoinsSum.Sub(after.CoinsSum...)

	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)
	err = keeper.UpdateQuota(ctx, left, "testaddr", 1, "ibc")
	require.NoError(t, err)

	after, found = keeper.GetQuotaData(ctx, "ibc")
	require.True(t, found)

	newcoin = sdk.NewCoin("usdt", sdk.NewInt(int64(1)))
	testcoins = sdk.NewCoins(newcoin)
	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)
	err = keeper.UpdateQuota(ctx, testcoins, "testaddr", 1, "ibc")
	require.ErrorContains(t, err, types.ErrQuotaExceed.Error())
}

func TestExceedMaxHistoryLength(t *testing.T) {
	keeper, ctx := keepertest.QuotaKeeper(t)
	items1 := createNQuota("testmodule1", 10)
	keeper.SetQuotaData(ctx, items1)

	testcoins := sdk.NewCoins(sdk.NewCoin("testa", sdk.NewInt(100)))
	err := keeper.UpdateQuota(ctx, testcoins, "testaddr", 1, "testmodule1")
	require.Error(t, err, "no quota for this module")
	keeper.SetParams(ctx, testParams())

	testcoins = sdk.NewCoins(sdk.NewCoin("testa", sdk.NewInt(100)))
	err = keeper.UpdateQuota(ctx, testcoins, "testaddr", 1, "testmodule1")
	require.Error(t, err, "no quota for this module")

	testcoins = sdk.NewCoins(sdk.NewCoin("testa", sdk.NewInt(100)))
	err = keeper.UpdateQuota(ctx, testcoins, "testaddr", 1, "ibc")
	require.Error(t, err, "quota not found")

	testcoins = sdk.NewCoins(sdk.NewCoin("invalid", sdk.NewInt(1)))
	err = keeper.UpdateQuota(ctx, testcoins, "testaddr", 1, "ibc")
	require.Error(t, err, "some coins cannot be found in target")

	newcoin := sdk.NewCoin("usdt", sdk.NewInt(1))
	testcoins = sdk.NewCoins(newcoin)

	err = keeper.UpdateQuota(ctx, testcoins, "testaddr", 1, "ibc")
	require.NoError(t, err)

	before, found := keeper.GetQuotaData(ctx, "ibc")
	require.True(t, found)

	firstEle := before.History[0]
	for i := 0; i < types.MAXHISTORY-1; i++ {
		ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)
		before.History = append(before.History, firstEle)
	}
	keeper.SetQuotaData(ctx, before)
	require.NoError(t, err)

	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)
	err = keeper.UpdateQuota(ctx, testcoins, "testaddr", 1, "ibc")
	require.ErrorContains(t, err, types.ErrQuotaExceed.Error())

	before, found = keeper.GetQuotaData(ctx, "ibc")
	require.True(t, found)

	keeper.BlockUpdateQuota(ctx)

	after, found := keeper.GetQuotaData(ctx, "ibc")
	require.True(t, found)

	require.Equal(t, len(before.History), len(after.History)+1)

	// we can put one more record
	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)
	err = keeper.UpdateQuota(ctx, testcoins, "testaddr", 1, "ibc")
	require.NoError(t, err)
}

func TestBlockUpdate(t *testing.T) {
	keeper, ctx := keepertest.QuotaKeeper(t)

	testcoins := sdk.NewCoins(sdk.NewCoin("testa", sdk.NewInt(100)))
	err := keeper.UpdateQuota(ctx, testcoins, "testaddr", 1, "testmodule1")
	require.Error(t, err, "no quota for this module")
	keeper.SetParams(ctx, testParams())

	var i int64
	for i = 0; i < testParams().Targets[0].HistoryLength; i++ {
		newcoin := sdk.NewCoin("usdt", sdk.NewInt(int64(1+i)))
		testcoins = sdk.NewCoins(newcoin)
		ctx = ctx.WithBlockHeight(ctx.BlockHeight() + int64(i+1))
		err = keeper.UpdateQuota(ctx, testcoins, "testaddr", 1, "ibc")
		require.NoError(t, err)
		if i < testParams().Targets[0].HistoryLength/2 {
			newcoin2 := sdk.NewCoin("ujolt", sdk.NewInt(int64(2+i)))
			testcoins2 := sdk.NewCoins(newcoin2)
			err = keeper.UpdateQuota(ctx, testcoins2, "testaddr", 1, "bridge")
			require.NoError(t, err)
		}
	}

	after := keeper.GetAllQuota(ctx)
	require.Len(t, after, 2)

	var bridgemodule, ibcmodule types.CoinsQuota
	if after[0].ModuleName == "bridge" {
		bridgemodule = after[0]
		ibcmodule = after[1]
	} else {
		bridgemodule = after[1]
		ibcmodule = after[0]
	}

	lengthibc := len(ibcmodule.History)
	lengthbridge := len(bridgemodule.History)

	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 512)

	keeper.BlockUpdateQuota(ctx)

	afterIBC, found := keeper.GetQuotaData(ctx, "ibc")
	require.True(t, found)
	afterBridge, found := keeper.GetQuotaData(ctx, "bridge")
	require.True(t, found)

	require.Equal(t, len(afterIBC.History), lengthibc-1)
	require.Equal(t, len(afterBridge.History), lengthbridge-1)

	keeper.BlockUpdateQuota(ctx)

	afterIBC, found = keeper.GetQuotaData(ctx, "ibc")
	require.True(t, found)
	afterBridge, found = keeper.GetQuotaData(ctx, "bridge")
	require.True(t, found)

	require.Equal(t, len(afterIBC.History), lengthibc-2)
	require.Equal(t, len(afterBridge.History), lengthbridge-2)

	for i := 0; i < 254; i++ {
		keeper.BlockUpdateQuota(ctx)
	}

	afterIBC, found = keeper.GetQuotaData(ctx, "ibc")
	require.True(t, found)
	afterBridge, found = keeper.GetQuotaData(ctx, "bridge")
	require.True(t, found)

	require.Equal(t, len(afterIBC.History), 256)
	require.Equal(t, len(afterBridge.History), 0)

	keeper.BlockUpdateQuota(ctx)

	afterIBC, found = keeper.GetQuotaData(ctx, "ibc")
	require.True(t, found)
	afterBridge, found = keeper.GetQuotaData(ctx, "bridge")
	require.True(t, found)

	require.Equal(t, 255, len(afterIBC.History))
	require.Equal(t, 0, len(afterBridge.History))

	for i := 0; i < 300; i++ {
		keeper.BlockUpdateQuota(ctx)
	}

	afterIBC, found = keeper.GetQuotaData(ctx, "ibc")
	require.True(t, found)
	afterBridge, found = keeper.GetQuotaData(ctx, "bridge")
	require.True(t, found)

	a := afterIBC.History[0].BlockHeight
	b := ctx.BlockHeight()
	require.True(t, b-a <= 512)
	require.Equal(t, 1, len(afterIBC.History))
	require.Equal(t, 0, len(afterBridge.History))

	// now we update the time, and the histor is gone
	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)
	keeper.BlockUpdateQuota(ctx)

	afterIBC, found = keeper.GetQuotaData(ctx, "ibc")
	require.True(t, found)
	afterBridge, found = keeper.GetQuotaData(ctx, "bridge")
	require.True(t, found)

	require.Equal(t, 0, len(afterIBC.History))
	require.Equal(t, 0, len(afterBridge.History))
}

func TestRevokehistory(t *testing.T) {
	keeper, ctx := keepertest.QuotaKeeper(t)

	testcoins := sdk.NewCoins(sdk.NewCoin("testa", sdk.NewInt(100)))
	err := keeper.UpdateQuota(ctx, testcoins, "testaddr", 1, "testmodule1")
	require.Error(t, err, "no quota for this module")
	keeper.SetParams(ctx, testParams())

	var i int64
	for i = 0; i < testParams().Targets[0].HistoryLength; i++ {
		newcoin := sdk.NewCoin("usdt", sdk.NewInt(int64(1+i)))
		testcoins = sdk.NewCoins(newcoin)
		ctx = ctx.WithBlockHeight(ctx.BlockHeight() + int64(1))
		err = keeper.UpdateQuota(ctx, testcoins, "testaddr", uint64(i), "ibc")
		require.NoError(t, err)
		if i < testParams().Targets[0].HistoryLength/2 {
			newcoin2 := sdk.NewCoin("ujolt", sdk.NewInt(int64(i)))
			testcoins2 := sdk.NewCoins(newcoin2)
			err = keeper.UpdateQuota(ctx, testcoins2, "testaddr", uint64(i), "bridge")
			require.NoError(t, err)
		}
	}

	afterIBC, found := keeper.GetQuotaData(ctx, "ibc")
	require.True(t, found)
	afterBridge, found := keeper.GetQuotaData(ctx, "bridge")
	require.True(t, found)
	require.Equal(t, len(afterBridge.History), 256)
	require.Equal(t, len(afterIBC.History), 512)
	keeper.RevokeHistory(ctx, "wrong", 1)

	afterIBC, found = keeper.GetQuotaData(ctx, "ibc")
	require.True(t, found)
	afterBridge, found = keeper.GetQuotaData(ctx, "bridge")
	require.True(t, found)
	require.Equal(t, len(afterBridge.History), 256)
	require.Equal(t, len(afterIBC.History), 512)

	before, found := keeper.GetQuotaData(ctx, "ibc")

	totaldeducted := sdk.NewCoins().Add(before.History[1].Amount...)
	totaldeducted = totaldeducted.Add(before.History[10].Amount...)
	totaldeducted = totaldeducted.Add(before.History[13].Amount...)
	totaldeducted = totaldeducted.Add(before.History[21].Amount...)

	keeper.RevokeHistory(ctx, "ibc", 1)
	keeper.RevokeHistory(ctx, "ibc", 10)
	keeper.RevokeHistory(ctx, "ibc", 13)
	keeper.RevokeHistory(ctx, "ibc", 21)

	afterIBC, found = keeper.GetQuotaData(ctx, "ibc")
	require.True(t, found)
	require.Equal(t, len(afterIBC.History), 508)

	delta := before.CoinsSum.Sub(afterIBC.CoinsSum...)

	require.True(t, delta.IsEqual(totaldeducted))

	targets := []uint64{1, 10, 13, 21}
	for _, el := range afterIBC.History {
		for _, el2 := range targets {
			if el.IbcSequence == el2 {
				require.Failf(t, "should not be found", "should not be found")
			}
		}
	}
}

func TestWhiteList(t *testing.T) {
	keeper, ctx := keepertest.QuotaKeeper(t)

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("jolt", "joltpub")

	testcoins := sdk.NewCoins(sdk.NewCoin("testa", sdk.NewInt(100)))
	err := keeper.UpdateQuota(ctx, testcoins, "testaddr", 1, "testmodule1")
	require.Error(t, err, "no quota for this module")
	tParams := testParams()
	w1 := types.WhiteList{"t1", []string{"jolt1gl7gfy5tjf9wlpumprya3fffxmdmlwcyykx8np", "jolt1h8m4p5vlaup3jzxv3k0tkvwamzel3regpsw5j2"}}
	w2 := types.WhiteList{"t2", []string{"jolt1gl7gfy5tjf9wlpumprya3fffxmdmlwcyykx8np"}}
	tParams.Whitelist = []*types.WhiteList{&w1, &w2}
	keeper.SetParams(ctx, tParams)

	found := keeper.WhetherOnwhitelist(ctx, "t1", "jolt1gl7gfy5tjf9wlpumprya3fffxmdmlwcyykx8np")
	require.True(t, found)
	found = keeper.WhetherOnwhitelist(ctx, "t1", "jolt1h8m4p5vlaup3jzxv3k0tkvwamzel3regpsw32")
	require.False(t, found)
	found = keeper.WhetherOnwhitelist(ctx, "t2", "jolt1gl7gfy5tjf9wlpumprya3fffxmdmlwcyykx8np")
	require.True(t, found)
	found = keeper.WhetherOnwhitelist(ctx, "t2", "jolt1h8m4p5vlaup3jzxv3k0tkvwamzel3regpsw32")
	require.False(t, found)
	found = keeper.WhetherOnwhitelist(ctx, "t1", "jolt1h8m4p5vlaup3jzxv3k0tkvwamzel3regpsw5j2")
	require.True(t, found)
}

func TestSubsetof(t *testing.T) {
	co1 := sdk.NewCoins(sdk.NewCoin("ibc65d0bec6dad96c7f5043d1e54e54b6bb5d5b3aec3ff6cebb75b9e059f3580ea3", sdk.NewInt(123)))
	target := sdk.NewCoins(sdk.NewCoin("ibc65D0BEC6DAD96C7F5043D1E54E54B6BB5D5B3AEC3FF6CEBB75B9E059F3580EA3", sdk.NewInt(123)), sdk.NewCoin("ibc9117A26BA81E29FA4F78F57DC2BD90CD3D26848101BA880445F119B22A1E254E", sdk.NewInt(2234)))

	ret := co1.DenomsSubsetOf(target)

	fmt.Printf(">>>>>%v\n", ret)
}
