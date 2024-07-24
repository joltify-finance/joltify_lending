package ante_test

import (
	"strings"
	"testing"

	tmlog "github.com/cometbft/cometbft/libs/log"

	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/app/ante"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtime "github.com/cometbft/cometbft/types/time"
	sdk "github.com/cosmos/cosmos-sdk/types"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func mustParseDecCoins(value string) sdkmath.LegacyDecCoins {
	coins, err := sdk.ParseDecCoins(strings.ReplaceAll(value, ";", ","))
	if err != nil {
		panic(err)
	}

	return coins
}

func TestEvmMinGasFilter(t *testing.T) {
	tApp := app.NewTestApp(tmlog.TestingLogger(), t.TempDir())
	handler := ante.NewEvmMinGasFilter(tApp.GetEVMKeeper())

	ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})
	err := tApp.GetEVMKeeper().SetParams(ctx, evmtypes.Params{
		EvmDenom: "ajolt",
	})
	if err != nil {
		return
	}

	testCases := []struct {
		name                 string
		minGasPrices         sdkmath.LegacyDecCoins
		expectedMinGasPrices sdkmath.LegacyDecCoins
	}{
		{
			"no min gas prices",
			mustParseDecCoins(""),
			mustParseDecCoins(""),
		},
		{
			"zero ujolt gas price",
			mustParseDecCoins("0ujolt"),
			mustParseDecCoins("0ujolt"),
		},
		{
			"non-zero ujolt gas price",
			mustParseDecCoins("0.001ujolt"),
			mustParseDecCoins("0.001ujolt"),
		},
		{
			"zero ujolt gas price, min ajolt price",
			mustParseDecCoins("0ujolt;100000ajolt"),
			mustParseDecCoins("0ujolt"), // ajolt is removed
		},
		{
			"zero ujolt gas price, min ajolt price, other token",
			mustParseDecCoins("0ujolt;100000ajolt;0.001other"),
			mustParseDecCoins("0ujolt;0.001other"), // ajolt is removed
		},
		{
			"non-zero ujolt gas price, min ajolt price",
			mustParseDecCoins("0.25ujolt;100000ajolt;0.001other"),
			mustParseDecCoins("0.25ujolt;0.001other"), // ajolt is removed
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})

			ctx = ctx.WithMinGasPrices(tc.minGasPrices)
			mmd := MockAnteHandler{}

			_, err := handler.AnteHandle(ctx, nil, false, mmd.AnteHandle)
			require.NoError(t, err)
			require.True(t, mmd.WasCalled)

			assert.NoError(t, mmd.CalledCtx.MinGasPrices().Validate())
			assert.Equal(t, tc.expectedMinGasPrices, mmd.CalledCtx.MinGasPrices())
		})
	}
}
