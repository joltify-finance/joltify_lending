package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/keeper"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func (suite *KeeperTestSuite) TestQuerierGetParams() {
	querier := keeper.NewQuerier(suite.keeper, types2.ModuleCdc.LegacyAmino)
	bz, err := querier(suite.ctx, []string{types2.QueryGetParams}, abci.RequestQuery{})
	suite.Require().NoError(err)
	suite.NotNil(bz)

	expParams := types2.Params{
		Markets: []types2.Market{
			{MarketID: "btc:usd", BaseAsset: "btc", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
			{MarketID: "xrp:usd", BaseAsset: "xrp", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
		},
	}
	var p types2.Params
	suite.Nil(types2.ModuleCdc.UnmarshalJSON(bz, &p))
	suite.Require().NoError(expParams.VerboseEqual(p))
}

func (suite *KeeperTestSuite) TestQuerierGetPrice() {
	// Invalid market
	requestParams := types2.NewQueryWithMarketIDParams("invalid")
	data, err := types2.ModuleCdc.LegacyAmino.MarshalJSON(requestParams)
	suite.Require().NoError(err)

	querier := keeper.NewQuerier(suite.keeper, types2.ModuleCdc.LegacyAmino)
	bz, err := querier(suite.ctx, []string{types2.QueryPrice}, abci.RequestQuery{Data: data})
	suite.Require().Error(err)
	suite.Nil(bz)

	// Valid market
	requestParams = types2.NewQueryWithMarketIDParams("btc:usd")
	data, err = types2.ModuleCdc.LegacyAmino.MarshalJSON(requestParams)
	suite.Require().NoError(err)

	bz, err = querier(suite.ctx, []string{types2.QueryPrice}, abci.RequestQuery{Data: data})
	suite.Require().NoError(err)
	suite.NotNil(bz)

	expPrice := types2.CurrentPrice{
		MarketID: "btc:usd",
		Price:    sdk.MustNewDecFromStr("8000.00"),
	}
	var p types2.CurrentPrice
	suite.Nil(types2.ModuleCdc.UnmarshalJSON(bz, &p))
	suite.Require().NoError(expPrice.VerboseEqual(p))
}

func (suite *KeeperTestSuite) TestQuerierGetPrices() {
	querier := keeper.NewQuerier(suite.keeper, types2.ModuleCdc.LegacyAmino)
	bz, err := querier(suite.ctx, []string{types2.QueryPrices}, abci.RequestQuery{})
	suite.Require().NoError(err)
	suite.NotNil(bz)

	expPrices := []types2.CurrentPrice{
		{
			MarketID: "btc:usd",
			Price:    sdk.MustNewDecFromStr("8000.00"),
		},
		{
			MarketID: "xrp:usd",
			Price:    sdk.MustNewDecFromStr("0.25"),
		},
	}
	var p []types2.CurrentPrice
	suite.Nil(types2.ModuleCdc.LegacyAmino.UnmarshalJSON(bz, &p))
	suite.Require().Len(p, 2, "should have 2 prices")

	suite.Equal(expPrices[0], p[0])
	suite.Equal(expPrices[1], p[1])
}

func (suite *KeeperTestSuite) TestQuerierGetRawPrices() {
	requestParams := types2.NewQueryWithMarketIDParams("btc:usd")
	data, err := types2.ModuleCdc.LegacyAmino.MarshalJSON(requestParams)
	suite.Require().NoError(err)

	querier := keeper.NewQuerier(suite.keeper, types2.ModuleCdc.LegacyAmino)
	bz, err := querier(suite.ctx, []string{types2.QueryRawPrices}, abci.RequestQuery{Data: data})
	suite.Require().NoError(err)
	suite.NotNil(bz)

	expPrices := []types2.PostedPrice{
		{
			MarketID: "btc:usd",
			Price:    sdk.MustNewDecFromStr("8000.00"),
		},
		{
			MarketID: "xrp:usd",
			Price:    sdk.MustNewDecFromStr("0.25"),
		},
	}
	var p []types2.PostedPrice
	suite.Nil(types2.ModuleCdc.LegacyAmino.UnmarshalJSON(bz, &p))
	suite.Require().Len(p, 1, "should have 1 raw price for btc:usd")

	// Expire time different so only compare prices
	suite.True(expPrices[0].Price.Equal(p[0].Price), "first posted price should be equal")
}

func (suite *KeeperTestSuite) TestQuerierGetOracles() {
	requestParams := types2.NewQueryWithMarketIDParams("btc:usd")
	data, err := types2.ModuleCdc.LegacyAmino.MarshalJSON(requestParams)
	suite.Require().NoError(err)

	querier := keeper.NewQuerier(suite.keeper, types2.ModuleCdc.LegacyAmino)
	bz, err := querier(suite.ctx, []string{types2.QueryOracles}, abci.RequestQuery{Data: data})
	suite.Require().NoError(err)
	suite.NotNil(bz)

	var p []string
	suite.Nil(types2.ModuleCdc.LegacyAmino.UnmarshalJSON(bz, &p))
	suite.Require().Empty(p, "there should be no oracles")
}

func (suite *KeeperTestSuite) TestQuerierGetMarkets() {
	querier := keeper.NewQuerier(suite.keeper, types2.ModuleCdc.LegacyAmino)
	bz, err := querier(suite.ctx, []string{types2.QueryMarkets}, abci.RequestQuery{})
	suite.Require().NoError(err)
	suite.NotNil(bz)

	expMarkets := []types2.Market{
		types2.NewMarket("btc:usd", "btc", "usd", []sdk.AccAddress(nil), true),
		types2.NewMarket("xrp:usd", "xrp", "usd", []sdk.AccAddress(nil), true),
	}

	var p []types2.Market
	suite.Nil(types2.ModuleCdc.LegacyAmino.UnmarshalJSON(bz, &p))
	suite.Require().Len(p, 2)

	suite.Equal(expMarkets[0], p[0])
}

func (suite *KeeperTestSuite) TestQuerierInvalid() {
	querier := keeper.NewQuerier(suite.keeper, types2.ModuleCdc.LegacyAmino)
	bz, err := querier(suite.ctx, []string{"invalidpath"}, abci.RequestQuery{})
	suite.Require().Error(err)
	suite.Nil(bz)
}
