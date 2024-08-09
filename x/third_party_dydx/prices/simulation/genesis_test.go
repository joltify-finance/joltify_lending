package simulation_test

import (
	"encoding/json"
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	v4module "github.com/joltify-finance/joltify_lending/dydx_helper/module"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	testutil_rand "github.com/joltify-finance/joltify_lending/testutil/dydx/testutil/rand"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/prices/simulation"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/prices/types"
	"github.com/stretchr/testify/require"
)

func TestRandomizedGenState(t *testing.T) {
	cdc := codec.NewProtoCodec(v4module.InterfaceRegistry)

	r := testutil_rand.NewRand()

	simState := module.SimulationState{
		AppParams:    make(simtypes.AppParams),
		Cdc:          cdc,
		Rand:         r,
		NumBonded:    3,
		Accounts:     simtypes.RandomAccounts(r, 3),
		InitialStake: sdkmath.NewInt(1000),
		GenState:     make(map[string]json.RawMessage),
		BondDenom:    sdk.DefaultBondDenom,
	}

	for i := 0; i < 100; i++ {
		simulation.RandomizedGenState(&simState)
		var pricesGenesis types.GenesisState
		simState.Cdc.MustUnmarshalJSON(simState.GenState[types.ModuleName], &pricesGenesis)

		require.True(t, len(pricesGenesis.MarketParams) >= 1)
		require.True(t, len(pricesGenesis.MarketParams) <= 1_000)
		for _, marketParam := range pricesGenesis.MarketParams {
			require.True(t, len(marketParam.Pair) >= 7)
			require.True(t, strings.HasSuffix(marketParam.Pair, "-USD"))

			require.True(t, marketParam.Exponent >= -15)
			require.True(t, marketParam.Exponent <= 15)

			require.True(t, marketParam.MinPriceChangePpm >= 1)
			require.True(t, marketParam.MinPriceChangePpm < 10_000)
		}
	}
}
