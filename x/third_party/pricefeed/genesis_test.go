package pricefeed_test

import (
	"context"
	"testing"

	"cosmossdk.io/log"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/x/third_party/pricefeed"
	"github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/keeper"

	"github.com/joltify-finance/joltify_lending/app"
	"github.com/stretchr/testify/suite"
)

type GenesisTestSuite struct {
	suite.Suite

	tApp   app.TestApp
	ctx    context.Context
	keeper keeper.Keeper
}

func (suite *GenesisTestSuite) SetupTest() {
	suite.tApp = app.NewTestApp(log.NewTestLogger(suite.T()), suite.T().TempDir())
	suite.ctx = suite.tApp.NewContext(true)
	suite.keeper = suite.tApp.GetPriceFeedKeeper()
}

func (suite *GenesisTestSuite) TestValidGenState() {
	suite.NotPanics(func() {
		suite.tApp.InitializeFromGenesisStates(nil, nil,
			NewPricefeedGenStateMulti(),
		)
	})
	_, addrs := app.GeneratePrivKeyAddressPairs(10)

	// Must create a new TestApp or InitChain will panic with index already set
	suite.tApp = app.NewTestApp(log.NewTestLogger(suite.T()), suite.T().TempDir())
	suite.NotPanics(func() {
		suite.tApp.InitializeFromGenesisStates(nil, nil,
			NewPricefeedGenStateWithOracles(addrs),
		)
	})
}

func (suite *GenesisTestSuite) TestInitExportGenState() {
	gs := NewPricefeedGen()

	suite.NotPanics(func() {
		pricefeed.InitGenesis(sdk.UnwrapSDKContext(suite.ctx), suite.keeper, gs)
	})

	exportedGs := pricefeed.ExportGenesis(sdk.UnwrapSDKContext(suite.ctx), suite.keeper)
	suite.NoError(gs.VerboseEqual(exportedGs), "exported genesis should match init genesis")
}

func (suite *GenesisTestSuite) TestParamPricesGenState() {
	gs := NewPricefeedGen()

	suite.NotPanics(func() {
		pricefeed.InitGenesis(sdk.UnwrapSDKContext(suite.ctx), suite.keeper, gs)
	})

	params := suite.keeper.GetParams(suite.ctx)
	suite.NoError(gs.Params.VerboseEqual(params), "params should equal init params")

	pps := suite.keeper.GetRawPrices(suite.ctx, "btc:usd")
	suite.NoError(gs.PostedPrices[0].VerboseEqual(pps[0]), "posted prices should equal init posted prices")
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GenesisTestSuite))
}
