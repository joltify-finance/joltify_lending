package keeper_test

import (
	"context"
	"testing"

	"cosmossdk.io/log"
	"github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/keeper"
	"github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/types"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/app"
)

type KeeperTestSuite struct {
	suite.Suite

	keeper keeper.Keeper
	addrs  []sdk.AccAddress
	ctx    context.Context
}

func (suite *KeeperTestSuite) SetupTest() {
	// suite.tApp = app.NewTestApp(log.NewTestLogger(suite.T()), suite.T().TempDir())
	// suite.ctx = suite.tApp.NewContext(false)
	// suite.keeper = suite.tApp.GetPriceFeedKeeper()

	tApp := app.NewTestApp(log.NewTestLogger(suite.T()), suite.T().TempDir())
	ctx := tApp.Ctx
	tApp.InitializeFromGenesisStates(nil, nil,
		NewPricefeedGenStateMulti(),
	)
	suite.keeper = tApp.GetPriceFeedKeeper()
	suite.ctx = ctx

	_, addrs := app.GeneratePrivKeyAddressPairs(10)
	suite.addrs = addrs
}

func (suite *KeeperTestSuite) TestGetSetOracles() {
	params := suite.keeper.GetParams(suite.ctx)
	acc, err := sdk.AccAddressFromBech32("jolt15qdefkmwswysgg4qxgqpqr35k3m49pkxu8ygkq")
	suite.Equal([]sdk.AccAddress{acc}, params.Markets[0].Oracles)

	params.Markets[0].Oracles = suite.addrs
	suite.NotPanics(func() { suite.keeper.SetParams(suite.ctx, params) })
	params = suite.keeper.GetParams(suite.ctx)
	suite.Equal(suite.addrs, params.Markets[0].Oracles)

	addr, err := suite.keeper.GetOracle(suite.ctx, params.Markets[0].MarketID, suite.addrs[0])
	suite.NoError(err)
	suite.Equal(suite.addrs[0], addr)
}

func (suite *KeeperTestSuite) TestGetAuthorizedAddresses() {
	_, oracles := app.GeneratePrivKeyAddressPairs(5)

	params := types.Params{
		Markets: []types.Market{
			{MarketID: "btc:usd", BaseAsset: "btc", QuoteAsset: "usd", Oracles: oracles[:3], Active: true},
			{MarketID: "xrp:usd", BaseAsset: "xrp", QuoteAsset: "usd", Oracles: oracles[2:], Active: true},
			{MarketID: "xrp:usd:30", BaseAsset: "xrp", QuoteAsset: "usd", Oracles: nil, Active: true},
		},
	}
	suite.keeper.SetParams(suite.ctx, params)

	actualOracles := suite.keeper.GetAuthorizedAddresses(suite.ctx)

	suite.Require().ElementsMatch(oracles, actualOracles)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
