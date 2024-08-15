package keeper_test

import (
	"testing"

	"github.com/joltify-finance/joltify_lending/lib"

	indexerevents "github.com/joltify-finance/joltify_lending/dydx_helper/indexer/events"
	"github.com/joltify-finance/joltify_lending/dydx_helper/indexer/indexer_manager"
	"github.com/joltify-finance/joltify_lending/mocks"
	clobtest "github.com/joltify-finance/joltify_lending/testutil/dydx/testutil/clob"
	keepertest "github.com/joltify-finance/joltify_lending/testutil/dydx/testutil/keeper"
	perptest "github.com/joltify-finance/joltify_lending/testutil/dydx/testutil/perpetuals"
	pricestest "github.com/joltify-finance/joltify_lending/testutil/dydx/testutil/prices"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/keeper"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/memclob"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/types"
	perptypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/perpetuals/types"
	pricestypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/prices/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateClobPair(t *testing.T) {
	testClobPair1 := *clobtest.GenerateClobPair(
		clobtest.WithId(1),
		clobtest.WithPerpetualId(1),
	)
	testPerp1 := *perptest.GeneratePerpetual(
		perptest.WithId(1),
		perptest.WithMarketId(1),
	)
	testMarket1 := *pricestest.GenerateMarketParamPrice(pricestest.WithId(1))
	testCases := map[string]struct {
		setup             func(t *testing.T, ks keepertest.ClobKeepersTestContext, manager *mocks.IndexerEventManager)
		msg               *types.MsgCreateClobPair
		expectedClobPairs []types.ClobPair
		expectedErr       string
	}{
		"Succeeds: create new clob pair": {
			setup: func(t *testing.T, ks keepertest.ClobKeepersTestContext, mockIndexerEventManager *mocks.IndexerEventManager) {
				keepertest.CreateTestPricesAndPerpetualMarkets(
					t,
					ks.Ctx,
					ks.PerpetualsKeeper,
					ks.PricesKeeper,
					[]perptypes.Perpetual{testPerp1},
					[]pricestypes.MarketParamPrice{testMarket1},
				)
				mockIndexerEventManager.On("AddTxnEvent",
					ks.Ctx,
					indexerevents.SubtypePerpetualMarket,
					indexerevents.PerpetualMarketEventVersion,
					indexer_manager.GetBytes(
						indexerevents.NewPerpetualMarketCreateEvent(
							testClobPair1.MustGetPerpetualId(),
							testClobPair1.GetId(),
							testPerp1.Params.Ticker,
							testPerp1.Params.MarketId,
							testClobPair1.Status,
							testClobPair1.QuantumConversionExponent,
							testPerp1.Params.AtomicResolution,
							testClobPair1.SubticksPerTick,
							testClobPair1.StepBaseQuantums,
							testPerp1.Params.LiquidityTier,
							testPerp1.Params.MarketType,
						),
					),
				).Return()
			},
			msg: &types.MsgCreateClobPair{
				Authority: lib.GovModuleAddress.String(),
				ClobPair:  testClobPair1,
			},
			expectedClobPairs: []types.ClobPair{testClobPair1},
		},
		"Failure: clob pair already exists": {
			setup: func(t *testing.T, ks keepertest.ClobKeepersTestContext, mockIndexerEventManager *mocks.IndexerEventManager) {
				keepertest.CreateTestPricesAndPerpetualMarkets(
					t,
					ks.Ctx,
					ks.PerpetualsKeeper,
					ks.PricesKeeper,
					[]perptypes.Perpetual{testPerp1},
					[]pricestypes.MarketParamPrice{testMarket1},
				)
				// set up mock indexer event manager to accept anything.
				mockIndexerEventManager.On("AddTxnEvent",
					ks.Ctx,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
				).Return()
				keepertest.CreateTestClobPairs(t, ks.Ctx, ks.ClobKeeper, []types.ClobPair{testClobPair1})
			},
			msg: &types.MsgCreateClobPair{
				Authority: lib.GovModuleAddress.String(),
				ClobPair:  testClobPair1,
			},
			expectedClobPairs: []types.ClobPair{testClobPair1},
			expectedErr:       "ClobPair with id already exists",
		},
		"Failure: perpetual already associated with existing clob pair": {
			setup: func(t *testing.T, ks keepertest.ClobKeepersTestContext, mockIndexerEventManager *mocks.IndexerEventManager) {
				keepertest.CreateTestPricesAndPerpetualMarkets(
					t,
					ks.Ctx,
					ks.PerpetualsKeeper,
					ks.PricesKeeper,
					[]perptypes.Perpetual{testPerp1},
					[]pricestypes.MarketParamPrice{testMarket1},
				)
				// set up mock indexer event manager to accept anything.
				mockIndexerEventManager.On("AddTxnEvent",
					ks.Ctx,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
				).Return()
				keepertest.CreateTestClobPairs(t, ks.Ctx, ks.ClobKeeper, []types.ClobPair{testClobPair1})
			},
			msg: &types.MsgCreateClobPair{
				Authority: lib.GovModuleAddress.String(),
				ClobPair:  *clobtest.GenerateClobPair(clobtest.WithId(3), clobtest.WithPerpetualId(1)),
			},
			expectedClobPairs: []types.ClobPair{testClobPair1},
			expectedErr:       "perpetual ID is already associated with an existing CLOB pair",
		},
		"Failure: refers to non-existing perpetual": {
			setup: func(t *testing.T, ks keepertest.ClobKeepersTestContext, mockIndexerEventManager *mocks.IndexerEventManager) {
			},
			msg: &types.MsgCreateClobPair{
				Authority: lib.GovModuleAddress.String(),
				ClobPair:  testClobPair1,
			},
			expectedClobPairs: nil,
			expectedErr:       "has invalid perpetual.: 1: Perpetual does not exist",
		},
		"Failure: invalid authority": {
			setup: func(t *testing.T, ks keepertest.ClobKeepersTestContext, mockIndexerEventManager *mocks.IndexerEventManager) {
			},
			msg: &types.MsgCreateClobPair{
				Authority: "invalid",
				ClobPair:  testClobPair1,
			},
			expectedClobPairs: nil,
			expectedErr:       "invalid authority invalid: expected gov account as only signer for proposal message",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			memClob := memclob.NewMemClobPriceTimePriority(false)
			mockIndexerEventManager := &mocks.IndexerEventManager{}
			ks := keepertest.NewClobKeepersTestContext(t, memClob, &mocks.BankKeeper{}, mockIndexerEventManager)
			tc.setup(t, ks, mockIndexerEventManager)

			msgServer := keeper.NewMsgServerImpl(ks.ClobKeeper)

			_, err := msgServer.CreateClobPair(ks.Ctx, tc.msg)
			if tc.expectedErr != "" {
				require.ErrorContains(t, err, tc.expectedErr)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tc.expectedClobPairs, ks.ClobKeeper.GetAllClobPairs(ks.Ctx))
		})
	}
}
