package keeper_test

import (
	"testing"

	"github.com/joltify-finance/joltify_lending/lib"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/joltify-finance/joltify_lending/testutil/dydx/testutil/keeper"
	perptest "github.com/joltify-finance/joltify_lending/testutil/dydx/testutil/perpetuals"
	pricestest "github.com/joltify-finance/joltify_lending/testutil/dydx/testutil/prices"
	perpkeeper "github.com/joltify-finance/joltify_lending/x/third_party_dydx/perpetuals/keeper"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/perpetuals/types"
	priceskeeper "github.com/joltify-finance/joltify_lending/x/third_party_dydx/prices/keeper"
	pricestypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/prices/types"
	"github.com/stretchr/testify/require"
)

func TestCreatePerpetual(t *testing.T) {
	testPerp1 := *perptest.GeneratePerpetual(
		perptest.WithId(1),
		perptest.WithMarketId(1),
	)
	testPerp2 := *perptest.GeneratePerpetual(
		perptest.WithId(2),
		perptest.WithMarketId(1),
	)
	testPerpIsolated := *perptest.GeneratePerpetual(
		perptest.WithId(3),
		perptest.WithMarketId(2),
		perptest.WithMarketType(types.PerpetualMarketType_PERPETUAL_MARKET_TYPE_ISOLATED),
	)
	testMarket1 := *pricestest.GenerateMarketParamPrice(pricestest.WithId(1))
	testMarket2 := *pricestest.GenerateMarketParamPrice(pricestest.WithId(2))
	testCases := map[string]struct {
		setup              func(*testing.T, sdk.Context, *perpkeeper.Keeper, *priceskeeper.Keeper)
		msg                *types.MsgCreatePerpetual
		expectedPerpetuals []types.Perpetual
		expectedErr        string
	}{
		"Succeeds: create new perpetual (id = 1)": {
			setup: func(t *testing.T, ctx sdk.Context, perpKeeper *perpkeeper.Keeper, pricesKeeper *priceskeeper.Keeper) {
				keepertest.CreateTestLiquidityTiers(t, ctx, perpKeeper)
				keepertest.CreateTestPriceMarkets(
					t,
					ctx,
					pricesKeeper,
					[]pricestypes.MarketParamPrice{testMarket1},
				)
			},
			msg: &types.MsgCreatePerpetual{
				Authority: lib.GovModuleAddress.String(),
				Params:    testPerp1.Params,
			},
			expectedPerpetuals: []types.Perpetual{testPerp1},
		},
		"Succeeds: create new perpetual (id = 2), with existing perpetual (id = 1) which use same market id": {
			setup: func(t *testing.T, ctx sdk.Context, perpKeeper *perpkeeper.Keeper, pricesKeeper *priceskeeper.Keeper) {
				keepertest.CreateTestPricesAndPerpetualMarkets(
					t,
					ctx,
					perpKeeper,
					pricesKeeper,
					[]types.Perpetual{testPerp1},
					[]pricestypes.MarketParamPrice{testMarket1},
				)
			},
			msg: &types.MsgCreatePerpetual{
				Authority: lib.GovModuleAddress.String(),
				Params:    testPerp2.Params,
			},
			expectedPerpetuals: []types.Perpetual{testPerp1, testPerp2},
		},
		"Succeeds: create new isolated market perpetual": {
			setup: func(
				t *testing.T,
				ctx sdk.Context,
				perpKeeper *perpkeeper.Keeper,
				pricesKeeper *priceskeeper.Keeper,
			) {
				keepertest.CreateTestLiquidityTiers(t, ctx, perpKeeper)
				keepertest.CreateTestPriceMarkets(
					t,
					ctx,
					pricesKeeper,
					[]pricestypes.MarketParamPrice{testMarket2},
				)
			},
			msg: &types.MsgCreatePerpetual{
				Authority: lib.GovModuleAddress.String(),
				Params:    testPerpIsolated.Params,
			},
			expectedPerpetuals: []types.Perpetual{testPerpIsolated},
		},
		"Failure: new perpetual id already exists in state": {
			setup: func(t *testing.T, ctx sdk.Context, perpKeeper *perpkeeper.Keeper, pricesKeeper *priceskeeper.Keeper) {
				keepertest.CreateTestPricesAndPerpetualMarkets(
					t,
					ctx,
					perpKeeper,
					pricesKeeper,
					[]types.Perpetual{testPerp1},
					[]pricestypes.MarketParamPrice{testMarket1},
				)
			},
			msg: &types.MsgCreatePerpetual{
				Authority: lib.GovModuleAddress.String(),
				Params:    testPerp1.Params,
			},
			expectedPerpetuals: []types.Perpetual{testPerp1},
			expectedErr:        "Perpetual already exists",
		},
		"Failure: refers to non-existing market id": {
			setup: func(t *testing.T, ctx sdk.Context, perpKeeper *perpkeeper.Keeper, pricesKeeper *priceskeeper.Keeper) {
				keepertest.CreateTestLiquidityTiers(t, ctx, perpKeeper)
			},
			msg: &types.MsgCreatePerpetual{
				Authority: lib.GovModuleAddress.String(),
				Params:    testPerp1.Params,
			},
			expectedPerpetuals: nil,
			expectedErr:        "Market price does not exist",
		},
		"Failure: invalid authority": {
			setup: func(t *testing.T, ctx sdk.Context, perpKeeper *perpkeeper.Keeper, pricesKeeper *priceskeeper.Keeper) {
				keepertest.CreateTestPricesAndPerpetualMarkets(
					t,
					ctx,
					perpKeeper,
					pricesKeeper,
					[]types.Perpetual{testPerp1},
					[]pricestypes.MarketParamPrice{testMarket1},
				)
			},
			msg: &types.MsgCreatePerpetual{
				Authority: "invalid",
				Params:    testPerp1.Params,
			},
			expectedPerpetuals: []types.Perpetual{testPerp1},
			expectedErr:        "invalid authority invalid",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			pc := keepertest.PerpetualsKeepers(t)
			tc.setup(t, pc.Ctx, pc.PerpetualsKeeper, pc.PricesKeeper)

			msgServer := perpkeeper.NewMsgServerImpl(pc.PerpetualsKeeper)

			_, err := msgServer.CreatePerpetual(pc.Ctx, tc.msg)
			if tc.expectedErr != "" {
				require.ErrorContains(t, err, tc.expectedErr)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tc.expectedPerpetuals, pc.PerpetualsKeeper.GetAllPerpetuals(pc.Ctx))
		})
	}
}