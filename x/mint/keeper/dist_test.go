package keeper_test

import (
	"testing"
	"time"

	tmlog "github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/joltify-finance/joltify_lending/app"
	joltminttypes "github.com/joltify-finance/joltify_lending/x/mint/types"
	"github.com/stretchr/testify/assert"
)

//func TestFirstDist(t *testing.T) {
//	tApp := app.NewTestApp()
//	ctx := tApp.NewContext(true, tmprototypes.Header{Height: 1, Time: tmtime.Now()})
//	k := tApp.GetMintKeeper()
//
//	params := joltminttypes.DefaultParams()
//	k.SetParams(ctx, params)
//	err := k.FirstDist(ctx, params)
//	assert.NoError(t, err)
//	acc := tApp.GetAccountKeeper().GetModuleAddress(incentivetypes.ModuleName)
//	balances := tApp.GetBankKeeper().GetBalance(ctx, acc, "ujolt")
//	assert.True(t, balances.Amount.Equal(params.CommunityProvisions.TruncateInt()))
//}

func TestMintCoinsAndDistribute(t *testing.T) {
	lg := tmlog.NewNopLogger()
	tApp := app.NewTestApp(lg, t.TempDir())
	k := tApp.GetMintKeeper()
	ctx := tApp.Ctx

	previous1 := ctx.BlockTime()
	ctx = tApp.App.BaseApp.NewContext(false, tmproto.Header{Height: 1, ChainID: "oppyChain-1", Time: previous1.Add(time.Minute)})

	previous1 = previous1.Add(time.Minute)

	params := joltminttypes.DefaultParams()
	k.SetParams(ctx, params)
	h := joltminttypes.HistoricalDistInfo{}
	k.SetDistInfo(ctx, h)

	k.DoDistribute(ctx)
	bk := tApp.GetBankKeeper()

	amountToCommunity := params.CurrentProvisions.Mul(sdk.MustNewDecFromStr("0.15")).TruncateInt()

	feeCollector := params.CurrentProvisions.TruncateInt().Sub(amountToCommunity)

	feeCollocter := authtypes.FeeCollectorName
	addr := tApp.GetAccountKeeper().GetModuleAddress(feeCollocter)
	balance := bk.GetBalance(ctx, addr, "ujolt")

	dis := tApp.GetDistrKeeper()
	balanceC := dis.GetFeePool(ctx).CommunityPool
	assert.True(t, balance.Amount.Equal(feeCollector))
	assert.True(t, balanceC.AmountOf("ujolt").Equal(sdk.NewDecFromInt(amountToCommunity)))
	h2 := k.GetDistInfo(ctx)
	assert.Equal(t, uint64(1), h2.DistributedRound)
	ctx = tApp.App.BaseApp.NewContext(false, tmproto.Header{Height: 1, ChainID: "oppyChain-1", Time: previous1.Add(time.Second * time.Duration(59))})
	previous1 = previous1.Add(time.Second * time.Duration(59))

	k.DoDistribute(ctx)
	balance = bk.GetBalance(ctx, addr, "ujolt")
	balanceC = dis.GetFeePool(ctx).CommunityPool
	assert.True(t, balance.Amount.Equal(feeCollector))
	assert.True(t, balanceC.AmountOf("ujolt").Equal(sdk.NewDecFromInt(amountToCommunity)))
	h2 = k.GetDistInfo(ctx)
	assert.Equal(t, uint64(1), h2.DistributedRound)

	ctx = tApp.App.BaseApp.NewContext(false, tmproto.Header{Height: 2, ChainID: "oppyChain-1", Time: previous1.Add(time.Minute)})

	previous1 = previous1.Add(time.Minute)

	k.DoDistribute(ctx)
	balance = bk.GetBalance(ctx, addr, "ujolt")
	balanceC = dis.GetFeePool(ctx).CommunityPool

	assert.True(t, balance.Amount.Equal(feeCollector.Add(feeCollector)))
	assert.True(t, balanceC.AmountOf("ujolt").Equal(sdk.NewDecFromInt(amountToCommunity.Add(amountToCommunity))))

	h2 = k.GetDistInfo(ctx)
	assert.Equal(t, uint64(2), h2.DistributedRound)

	balanceLast := balance
	balanceCLast := balanceC

	for i := 1; i < 60; i++ {
		delta := 60 + i

		ctx = tApp.App.BaseApp.NewContext(false, tmproto.Header{Height: 2, ChainID: "oppyChain-1", Time: previous1.Add(time.Duration(delta))})

		k.DoDistribute(ctx)
		balance = bk.GetBalance(ctx, addr, "ujolt")
		balanceC = dis.GetFeePool(ctx).CommunityPool

		assert.True(t, balance.Amount.Equal(balanceLast.Amount))
		assert.True(t, balanceC.IsEqual(balanceCLast))

		h2 = k.GetDistInfo(ctx)
		assert.Equal(t, uint64(2), h2.DistributedRound)
	}

	ctx = tApp.App.BaseApp.NewContext(false, tmproto.Header{Height: 1, ChainID: "oppyChain-1", Time: previous1.Add(time.Minute)})

	k.DoDistribute(ctx)

	balance = bk.GetBalance(ctx, addr, "ujolt")
	balanceC = dis.GetFeePool(ctx).CommunityPool

	assert.True(t, balance.Amount.Equal(feeCollector.MulRaw(3)))
	assert.True(t, balanceC.AmountOf("ujolt").Equal(sdk.NewDecFromInt(amountToCommunity.MulRaw(3))))

	h2 = k.GetDistInfo(ctx)
	assert.Equal(t, uint64(3), h2.DistributedRound)
}

//func TestMintCoinsAndDistributeFor100Years(t *testing.T) {
//	tApp := app.NewTestApp()
//	ctx := tApp.Ctx
//	k := tApp.GetMintKeeper()
//	params := joltminttypes.DefaultParams()
//	k.SetParams(ctx, params)
//	h := joltminttypes.HistoricalDistInfo{}
//	k.SetDistInfo(ctx, h)
//
//	previous := ctx.BlockTime()
//	bk := tApp.GetBankKeeper()
//	dis := tApp.GetDistrKeeper()
//	feeCollector := authtypes.FeeCollectorName
//	addr := tApp.GetAccountKeeper().GetModuleAddress(feeCollector)
//	minutesInYear := 365 * 24 * 60
//	currentProvision := params.CurrentProvisions
//	var balance, balanceBefore, deltaCoin sdk.Coin
//	var balanceC, balanceBeforeC, deltaCoinsC sdk.DecCoins
//	var h2 joltminttypes.HistoricalDistInfo
//	for i := 1; i < 100*365*24*60; i++ {
//		delta := 60 * i
//		ctx = tApp.App.BaseApp.NewContext(false, tmproto.Header{Height: 1 + int64(i), ChainID: "oppyChain-1", Time: previous.Add(time.Second * time.Duration(delta))})
//
//		balanceBefore = bk.GetBalance(ctx, addr, "ujolt")
//
//		balanceBeforeC = dis.GetFeePool(ctx).CommunityPool
//		k.DoDistribute(ctx)
//
//		balance = bk.GetBalance(ctx, addr, "ujolt")
//		balanceC = dis.GetFeePool(ctx).CommunityPool
//		deltaCoin = balance.Sub(balanceBefore)
//		deltaCoinsC = balanceC.Sub(balanceBeforeC)
//
//		amountToCommunity := currentProvision.Mul(sdk.MustNewDecFromStr("0.2")).TruncateDec()
//		toFeeCollector := currentProvision.TruncateInt().Sub(amountToCommunity.TruncateInt())
//
//		assert.True(t, deltaCoin.Amount.Equal(toFeeCollector))
//		assert.True(t, deltaCoinsC.AmountOf("ujolt").Equal(amountToCommunity))
//
//		deltaCoinsC = balanceC.Sub(balanceBeforeC)
//
//		if currentProvision.TruncateInt().IsZero() {
//			return
//		}
//		h2 = k.GetDistInfo(ctx)
//		assert.Equal(t, uint64(i), h2.DistributedRound)
//		if h2.DistributedRound%uint64(minutesInYear) == 0 {
//			adj := sdk.MustNewDecFromStr("2")
//			currentProvision = currentProvision.Quo(adj)
//		}
//	}
//}

func TestMintCoinsAndDistributeFor3Years(t *testing.T) {
	lg := tmlog.NewNopLogger()
	tApp := app.NewTestApp(lg, t.TempDir())
	ctx := tApp.Ctx
	k := tApp.GetMintKeeper()
	params := joltminttypes.DefaultParams()
	k.SetParams(ctx, params)
	h := joltminttypes.HistoricalDistInfo{}
	k.SetDistInfo(ctx, h)

	previous := ctx.BlockTime()
	bk := tApp.GetBankKeeper()
	dis := tApp.GetDistrKeeper()
	feeCollector := authtypes.FeeCollectorName
	addr := tApp.GetAccountKeeper().GetModuleAddress(feeCollector)
	minutesInYear := 365 * 24 * 60
	currentProvision := params.CurrentProvisions
	var balance, balanceBefore, deltaCoin sdk.Coin
	var balanceC, balanceBeforeC, deltaCoinsC sdk.DecCoins
	var h2 joltminttypes.HistoricalDistInfo
	counter := 0
	for i := 1; i < 100*365*24*60; i++ {
		delta := 60 * i
		ctx = tApp.App.BaseApp.NewContext(false, tmproto.Header{Height: 1 + int64(i), ChainID: "oppyChain-1", Time: previous.Add(time.Second * time.Duration(delta))})

		balanceBefore = bk.GetBalance(ctx, addr, "ujolt")

		balanceBeforeC = dis.GetFeePool(ctx).CommunityPool
		k.DoDistribute(ctx)

		balance = bk.GetBalance(ctx, addr, "ujolt")
		balanceC = dis.GetFeePool(ctx).CommunityPool
		deltaCoin = balance.Sub(balanceBefore)
		deltaCoinsC = balanceC.Sub(balanceBeforeC)

		amountToCommunity := currentProvision.Mul(sdk.MustNewDecFromStr("0.15")).TruncateDec()
		toFeeCollector := currentProvision.TruncateInt().Sub(amountToCommunity.TruncateInt())

		assert.True(t, deltaCoin.Amount.Equal(toFeeCollector))
		assert.True(t, deltaCoinsC.AmountOf("ujolt").Equal(amountToCommunity))

		deltaCoinsC = balanceC.Sub(balanceBeforeC)

		if currentProvision.TruncateInt().IsZero() {
			return
		}
		h2 = k.GetDistInfo(ctx)
		assert.Equal(t, uint64(i), h2.DistributedRound)
		if h2.DistributedRound%uint64(minutesInYear) == 0 {
			adj := sdk.MustNewDecFromStr("2")
			currentProvision = currentProvision.Quo(adj)
			counter++
			if counter == 2 {
				return
			}
		}
	}
}

func TestMintCoinsAndDistributeForAllYears(t *testing.T) {
	lg := tmlog.TestingLogger()
	tApp := app.NewTestApp(lg, t.TempDir())
	ctx := tApp.Ctx
	k := tApp.GetMintKeeper()
	params := joltminttypes.DefaultParams()
	k.SetParams(ctx, params)
	h := joltminttypes.HistoricalDistInfo{}
	k.SetDistInfo(ctx, h)

	previous := ctx.BlockTime()
	bk := tApp.GetBankKeeper()
	dis := tApp.GetDistrKeeper()
	feeCollector := authtypes.FeeCollectorName
	addr := tApp.GetAccountKeeper().GetModuleAddress(feeCollector)
	minutesInYear := 365 * 24 * 60
	currentProvision := params.CurrentProvisions
	var balance, balanceBefore, deltaCoin sdk.Coin
	var balanceC, balanceBeforeC, deltaCoinsC sdk.DecCoins
	var h2 joltminttypes.HistoricalDistInfo
	for i := 0; i < 100*365*24*60; i++ {
		delta := 24 * 365 * i
		ctx = tApp.App.BaseApp.NewContext(false, tmproto.Header{Height: 1 + int64(i), ChainID: "oppyChain-1", Time: previous.Add(time.Hour * time.Duration(delta))})

		balanceBefore = bk.GetBalance(ctx, addr, "ujolt")

		balanceBeforeC = dis.GetFeePool(ctx).CommunityPool
		k.DoDistribute(ctx)

		// get the balance of the fee collector
		balance = bk.GetBalance(ctx, addr, "ujolt")
		balanceC = dis.GetFeePool(ctx).CommunityPool
		deltaCoin = balance.Sub(balanceBefore)
		deltaCoinsC = balanceC.Sub(balanceBeforeC)

		amountToCommunity := currentProvision.Mul(sdk.MustNewDecFromStr("0.15")).TruncateDec()
		toFeeCollector := currentProvision.TruncateInt().Sub(amountToCommunity.TruncateInt())

		assert.True(t, deltaCoin.Amount.Equal(toFeeCollector))
		assert.True(t, deltaCoinsC.AmountOf("ujolt").Equal(amountToCommunity))

		if currentProvision.TruncateInt().IsZero() {
			return
		}
		h2 = k.GetDistInfo(ctx)
		if h2.DistributedRound%uint64(minutesInYear) == 0 {
			adj := sdk.MustNewDecFromStr("2")
			currentProvision = currentProvision.Quo(adj)
		}

		h2.DistributedRound = uint64(minutesInYear - 1)
		k.SetDistInfo(ctx, h2)
	}
}
