package keeper_test

import (
	"fmt"
	"testing"
	"time"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/joltify-finance/joltify_lending/x/third_party/jolt/keeper"

	incentivetypes "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"

	tmlog "github.com/cometbft/cometbft/libs/log"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/app"
	joltminttypes "github.com/joltify-finance/joltify_lending/x/mint/types"
	"github.com/stretchr/testify/assert"
)

func TestFirstDist(t *testing.T) {
	lg := tmlog.NewNopLogger()
	tApp := app.NewTestApp(lg, t.TempDir())
	k := tApp.GetMintKeeper()
	ctx := tApp.Ctx

	params := joltminttypes.DefaultParams()
	k.SetParams(ctx, params)

	acc := tApp.GetAccountKeeper().GetModuleAddress(incentivetypes.ModuleName)
	balances := tApp.GetBankKeeper().GetBalance(ctx, acc, "ujolt")

	firstDrop, ok := sdkmath.NewIntFromString("100000000000")
	assert.True(t, ok)
	assert.True(t, balances.Amount.Equal(firstDrop))
}

func TestMintCoinsAndDistribute(t *testing.T) {
	lg := tmlog.NewNopLogger()
	tApp := app.NewTestApp(lg, t.TempDir())
	k := tApp.GetMintKeeper()

	ctx := tApp.Ctx

	apy, err := sdkmath.LegacyNewDecFromStr("0.08")
	assert.NoError(t, err)

	adjAPY := sdkmath.LegacyOneDec().Add(apy)

	spy, err := keeper.APYToSPY(adjAPY)
	assert.NoError(t, err)

	params := joltminttypes.DefaultParams()
	t.Logf("we set the spy to %s", spy.String())
	params.NodeSPY = spy

	k.SetParams(ctx, params)
	h := joltminttypes.HistoricalDistInfo{
		PayoutTime: ctx.BlockTime(),
	}
	k.SetDistInfo(ctx, h)

	k.DoDistribute(ctx)

	bk := tApp.GetBankKeeper()
	received := bk.GetBalance(ctx, tApp.GetAccountKeeper().GetModuleAddress(authtypes.FeeCollectorName), "ujolt")
	assert.True(t, received.IsZero())

	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Second * 59))
	k.DoDistribute(ctx)
	received = bk.GetBalance(ctx, tApp.GetAccountKeeper().GetModuleAddress(authtypes.FeeCollectorName), "ujolt")
	assert.True(t, received.IsZero())

	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Second))
	k.DoDistribute(ctx)
	received = bk.GetBalance(ctx, tApp.GetAccountKeeper().GetModuleAddress(authtypes.FeeCollectorName), "ujolt")

	stakingkepper := tApp.GetStakingKeeper()
	totalBounded := stakingkepper.TotalBondedTokens(ctx)
	fmt.Printf("total bonded>>>>%v\n", totalBounded.String())

	yearlyWeGet := apy.MulInt(totalBounded).TruncateInt()
	t.Logf("we get yearly %s", yearlyWeGet.String())

	t.Logf("we have received for one minute %v", received.Amount.String())
	yearlyMinutes := int64(365 * 24 * 60)
	actualReceived := received.Amount.Mul(sdk.NewInt(yearlyMinutes))
	fmt.Printf("gap is %v\n", yearlyWeGet.Sub(actualReceived).Quo(sdk.NewInt(1000000)))
	gap := yearlyWeGet.Sub(actualReceived).Quo(sdk.NewInt(1000000))
	assert.True(t, gap.LT(sdk.NewInt(40000)))

	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Second))
	k.DoDistribute(ctx)
	received2 := bk.GetBalance(ctx, tApp.GetAccountKeeper().GetModuleAddress(authtypes.FeeCollectorName), "ujolt")
	assert.True(t, received.IsEqual(received2))

	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Second * 59))
	k.DoDistribute(ctx)
	received3 := bk.GetBalance(ctx, tApp.GetAccountKeeper().GetModuleAddress(authtypes.FeeCollectorName), "ujolt")
	assert.True(t, received3.IsEqual(received2.Add(received2)))
}
