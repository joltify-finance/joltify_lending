package keeper_test

import (
	tmlog "github.com/tendermint/tendermint/libs/log"
	"math/rand"
	"sort"
	"strings"
	"testing"
	"time"

	keeper2 "github.com/joltify-finance/joltify_lending/x/third_party/cdp/keeper"
	types3 "github.com/joltify-finance/joltify_lending/x/third_party/cdp/types"
	pfkeeper "github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/keeper"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/types"

	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/simulation"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/joltify-finance/joltify_lending/app"
)

const (
	custom = "custom"
)

type QuerierTestSuite struct {
	suite.Suite

	keeper          keeper2.Keeper
	pricefeedKeeper pfkeeper.Keeper
	addrs           []sdk.AccAddress
	app             app.TestApp
	cdps            types3.CDPs
	augmentedCDPs   types3.AugmentedCDPs
	ctx             sdk.Context
	querier         sdk.Querier
	cdc             codec.Codec
	legacyAmino     codec.LegacyAmino
}

func (suite *QuerierTestSuite) SetupTest() {
	tApp := app.NewTestApp(tmlog.TestingLogger(), suite.T().TempDir())
	ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})
	suite.cdc = tApp.AppCodec()
	suite.legacyAmino = *tApp.LegacyAmino()

	cdps := make(types3.CDPs, 100)
	augmentedCDPs := make(types3.AugmentedCDPs, 100)
	_, addrs := app.GeneratePrivKeyAddressPairs(100)
	coins := cs(c("btc", 10000000000), c("xrp", 10000000000))

	authGS := app.NewFundedGenStateWithSameCoins(tApp.AppCodec(), coins, addrs)

	var genAcc []authtypes.GenesisAccount
	for _, el := range addrs {
		b := authtypes.NewBaseAccount(el, nil, 0, 0)
		genAcc = append(genAcc, b)
	}

	tApp.InitializeFromGenesisStates(genAcc, coins,
		authGS,
		NewPricefeedGenStateMulti(suite.cdc),
		NewCDPGenStateHighDebtLimit(suite.cdc),
	)

	suite.ctx = ctx
	suite.app = tApp
	suite.keeper = tApp.GetCDPKeeper()
	suite.pricefeedKeeper = tApp.GetPriceFeedKeeper()

	// Set up markets
	oracle := addrs[9]
	marketParams := types2.Params{
		Markets: types2.Markets{
			types2.Market{MarketID: "xrp-usd", BaseAsset: "xrp", QuoteAsset: "usd", Oracles: []sdk.AccAddress{oracle}, Active: true},
			types2.Market{MarketID: "btc-usd", BaseAsset: "btc", QuoteAsset: "usd", Oracles: []sdk.AccAddress{oracle}, Active: true},
		},
	}
	suite.pricefeedKeeper.SetParams(ctx, marketParams)

	// Set collateral prices for use in collateralization calculations
	_, err := suite.pricefeedKeeper.SetPrice(
		ctx, oracle, "xrp-usd",
		sdk.MustNewDecFromStr("0.75"),
		time.Now().Add(1*time.Hour))
	suite.Nil(err)

	_, err = suite.pricefeedKeeper.SetPrice(
		ctx, oracle, "btc-usd",
		sdk.MustNewDecFromStr("5000"),
		time.Now().Add(1*time.Hour))
	suite.Nil(err)

	for j := 0; j < 100; j++ {
		collateral := "xrp"
		amount := simulation.RandIntBetween(rand.New(rand.NewSource(int64(j))), 2500000000, 9000000000)
		debt := simulation.RandIntBetween(rand.New(rand.NewSource(int64(j))), 50000000, 250000000)
		if j%2 == 0 {
			collateral = "btc"
			amount = simulation.RandIntBetween(rand.New(rand.NewSource(int64(j))), 500000000, 5000000000)
			debt = simulation.RandIntBetween(rand.New(rand.NewSource(int64(j))), 1000000000, 25000000000)
		}
		err = suite.keeper.AddCdp(suite.ctx, addrs[j], c(collateral, int64(amount)), c("usdx", int64(debt)), collateral+"-a")
		suite.NoError(err)
		c, f := suite.keeper.GetCDP(suite.ctx, collateral+"-a", uint64(j+1))
		suite.True(f)
		cdps[j] = c
		aCDP := suite.keeper.LoadAugmentedCDP(suite.ctx, c)
		augmentedCDPs[j] = aCDP
	}

	suite.cdps = cdps
	suite.augmentedCDPs = augmentedCDPs
	suite.querier = keeper2.NewQuerier(suite.keeper, tApp.LegacyAmino())
	suite.addrs = addrs
}

func (suite *QuerierTestSuite) TestQueryCdp() {
	ctx := suite.ctx.WithIsCheckTx(false)
	query := abci.RequestQuery{
		Path: strings.Join([]string{custom, types3.QuerierRoute, types3.QueryGetCdp}, "/"),
		Data: suite.legacyAmino.MustMarshalJSON(types3.NewQueryCdpParams(suite.cdps[0].Owner, suite.cdps[0].Type)),
	}
	bz, err := suite.querier(ctx, []string{types3.QueryGetCdp}, query)
	suite.Nil(err)
	suite.NotNil(bz)

	var c types3.AugmentedCDP
	suite.Nil(suite.legacyAmino.UnmarshalJSON(bz, &c))
	suite.Equal(suite.augmentedCDPs[0], c)

	query = abci.RequestQuery{
		Path: strings.Join([]string{custom, types3.QuerierRoute, types3.QueryGetCdp}, "/"),
		Data: suite.legacyAmino.MustMarshalJSON(types3.NewQueryCdpParams(suite.cdps[0].Owner, "lol-a")),
	}
	_, err = suite.querier(ctx, []string{types3.QueryGetCdp}, query)
	suite.Error(err)

	query = abci.RequestQuery{
		Path: strings.Join([]string{custom, "nonsense"}, "/"),
		Data: []byte("nonsense"),
	}

	_, err = suite.querier(ctx, []string{query.Path}, query)
	suite.Error(err)

	_, err = suite.querier(ctx, []string{types3.QueryGetCdp}, query)
	suite.Error(err)

	query = abci.RequestQuery{
		Path: strings.Join([]string{custom, types3.QuerierRoute, types3.QueryGetCdp}, "/"),
		Data: suite.legacyAmino.MustMarshalJSON(types3.NewQueryCdpParams(suite.cdps[0].Owner, "xrp-a")),
	}
	_, err = suite.querier(ctx, []string{types3.QueryGetCdp}, query)
	suite.Error(err)
}

func (suite *QuerierTestSuite) TestQueryCdpsByCollateralType() {
	ctx := suite.ctx.WithIsCheckTx(false)
	query := abci.RequestQuery{
		Path: strings.Join([]string{custom, types3.QuerierRoute, types3.QueryGetCdpsByCollateralType}, "/"),
		Data: suite.legacyAmino.MustMarshalJSON(types3.NewQueryCdpsByCollateralTypeParams(suite.cdps[0].Type)),
	}
	bz, err := suite.querier(ctx, []string{types3.QueryGetCdpsByCollateralType}, query)
	suite.Nil(err)
	suite.NotNil(bz)

	var c types3.AugmentedCDPs
	suite.Nil(suite.legacyAmino.UnmarshalJSON(bz, &c))
	suite.Equal(50, len(c))

	query = abci.RequestQuery{
		Path: strings.Join([]string{custom, types3.QuerierRoute, types3.QueryGetCdpsByCollateralType}, "/"),
		Data: suite.legacyAmino.MustMarshalJSON(types3.NewQueryCdpsByCollateralTypeParams("lol-a")),
	}
	_, err = suite.querier(ctx, []string{types3.QueryGetCdpsByCollateralType}, query)
	suite.Error(err)
}

func (suite *QuerierTestSuite) TestQueryCdpsByRatio() {
	ratioCountBtc := 0
	ratioCountXrp := 0
	xrpRatio := d("2.0")
	btcRatio := d("2500")
	expectedXrpIds := []int{}
	expectedBtcIds := []int{}
	for _, cdp := range suite.cdps {
		absoluteRatio := suite.keeper.CalculateCollateralToDebtRatio(suite.ctx, cdp.Collateral, cdp.Type, cdp.Principal)
		collateralizationRatio, err := suite.keeper.CalculateCollateralizationRatioFromAbsoluteRatio(suite.ctx, cdp.Type, absoluteRatio, "liquidation")
		suite.Nil(err)
		if cdp.Collateral.Denom == "xrp" {
			if collateralizationRatio.LT(xrpRatio) {
				ratioCountXrp += 1
				expectedXrpIds = append(expectedXrpIds, int(cdp.ID))
			}
		} else {
			if collateralizationRatio.LT(btcRatio) {
				ratioCountBtc += 1
				expectedBtcIds = append(expectedBtcIds, int(cdp.ID))
			}
		}
	}

	ctx := suite.ctx.WithIsCheckTx(false)
	query := abci.RequestQuery{
		Path: strings.Join([]string{custom, types3.QuerierRoute, types3.QueryGetCdpsByCollateralization}, "/"),
		Data: suite.legacyAmino.MustMarshalJSON(types3.NewQueryCdpsByRatioParams("xrp-a", xrpRatio)),
	}
	bz, err := suite.querier(ctx, []string{types3.QueryGetCdpsByCollateralization}, query)
	suite.Nil(err)
	suite.NotNil(bz)

	var c types3.AugmentedCDPs
	actualXrpIds := []int{}
	suite.Nil(suite.legacyAmino.UnmarshalJSON(bz, &c))
	for _, k := range c {
		actualXrpIds = append(actualXrpIds, int(k.ID))
	}
	sort.Ints(actualXrpIds)
	suite.Equal(expectedXrpIds, actualXrpIds)

	query = abci.RequestQuery{
		Path: strings.Join([]string{custom, types3.QuerierRoute, types3.QueryGetCdpsByCollateralization}, "/"),
		Data: suite.legacyAmino.MustMarshalJSON(types3.NewQueryCdpsByRatioParams("btc-a", btcRatio)),
	}
	bz, err = suite.querier(ctx, []string{types3.QueryGetCdpsByCollateralization}, query)
	suite.Nil(err)
	suite.NotNil(bz)

	c = types3.AugmentedCDPs{}
	actualBtcIds := []int{}
	suite.Nil(suite.legacyAmino.UnmarshalJSON(bz, &c))
	for _, k := range c {
		actualBtcIds = append(actualBtcIds, int(k.ID))
	}
	sort.Ints(actualBtcIds)
	suite.Equal(expectedBtcIds, actualBtcIds)

	query = abci.RequestQuery{
		Path: strings.Join([]string{custom, types3.QuerierRoute, types3.QueryGetCdpsByCollateralization}, "/"),
		Data: suite.legacyAmino.MustMarshalJSON(types3.NewQueryCdpsByRatioParams("xrp-a", d("0.003"))),
	}
	bz, err = suite.querier(ctx, []string{types3.QueryGetCdpsByCollateralization}, query)
	suite.Nil(err)
	suite.NotNil(bz)
	c = types3.AugmentedCDPs{}
	suite.Nil(suite.legacyAmino.UnmarshalJSON(bz, &c))
	suite.Equal(0, len(c))
}

func (suite *QuerierTestSuite) TestQueryParams() {
	ctx := suite.ctx.WithIsCheckTx(false)
	bz, err := suite.querier(ctx, []string{types3.QueryGetParams}, abci.RequestQuery{})
	suite.NoError(err)
	suite.NotNil(bz)

	var p types3.Params
	err = suite.legacyAmino.UnmarshalJSON(bz, &p)
	suite.NoError(err)

	cdpGS := NewCDPGenStateHighDebtLimit(suite.app.AppCodec())
	gs := types3.GenesisState{}
	err = suite.legacyAmino.UnmarshalJSON(cdpGS["cdp"], &gs)
	suite.NoError(err)

	suite.Equal(gs.Params, p)
}

func (suite *QuerierTestSuite) TestQueryDeposits() {
	ctx := suite.ctx.WithIsCheckTx(false)
	query := abci.RequestQuery{
		Path: strings.Join([]string{custom, types3.QuerierRoute, types3.QueryGetCdpDeposits}, "/"),
		Data: suite.legacyAmino.MustMarshalJSON(types3.NewQueryCdpDeposits(suite.cdps[0].Owner, suite.cdps[0].Type)),
	}

	bz, err := suite.querier(ctx, []string{types3.QueryGetCdpDeposits}, query)
	suite.Nil(err)
	suite.NotNil(bz)

	deposits := suite.keeper.GetDeposits(ctx, suite.cdps[0].ID)

	var d types3.Deposits
	suite.Nil(suite.legacyAmino.UnmarshalJSON(bz, &d))
	suite.Equal(deposits, d)
}

func (suite *QuerierTestSuite) TestQueryAccounts() {
	bz, err := suite.querier(suite.ctx, []string{types3.QueryGetAccounts}, abci.RequestQuery{})
	suite.Require().NoError(err)
	suite.Require().NotNil(bz)

	var accounts []authtypes.ModuleAccount
	suite.Require().Nil(suite.legacyAmino.UnmarshalJSON(bz, &accounts))
	suite.Require().Equal(2, len(accounts))

	findByName := func(name string) bool {
		for _, account := range accounts {
			if account.GetName() == name {
				return true
			}
		}
		return false
	}

	suite.Require().True(findByName("cdp"))
	suite.Require().True(findByName("liquidator"))
}

func (suite *QuerierTestSuite) TestFindIntersection() {
	a := types3.CDPs{suite.cdps[0], suite.cdps[1], suite.cdps[2], suite.cdps[3], suite.cdps[4]}
	b := types3.CDPs{suite.cdps[3], suite.cdps[4], suite.cdps[5], suite.cdps[6], suite.cdps[7]}
	expectedIntersection1 := types3.CDPs{suite.cdps[3], suite.cdps[4]}

	intersection1 := keeper2.FindIntersection(a, b)
	suite.Require().Equal(expectedIntersection1, intersection1)

	c := types3.CDPs{suite.cdps[0], suite.cdps[1], suite.cdps[2], suite.cdps[3], suite.cdps[4]}
	d := types3.CDPs{suite.cdps[5], suite.cdps[6], suite.cdps[7], suite.cdps[8], suite.cdps[9]}
	expectedIntersection2 := types3.CDPs{}

	intersection2 := keeper2.FindIntersection(c, d)
	suite.Require().Equal(expectedIntersection2, intersection2)

	e := types3.CDPs{suite.cdps[0]}
	f := types3.CDPs{}
	expectedIntersection3 := types3.CDPs{}

	intersection3 := keeper2.FindIntersection(e, f)
	suite.Require().Equal(expectedIntersection3, intersection3)
}

func (suite *QuerierTestSuite) TestFilterCDPs() {
	paramsType := types3.NewQueryCdpsParams(1, 100, "btc-a", sdk.AccAddress{}, 0, sdk.ZeroDec())
	filteredCDPs1, err := keeper2.FilterCDPs(suite.ctx, suite.keeper, paramsType)
	suite.Require().NoError(err)
	suite.Require().Equal(50, len(filteredCDPs1))

	paramsOwner := types3.NewQueryCdpsParams(1, 100, "", suite.cdps[10].Owner, 0, sdk.ZeroDec())
	filteredCDPs2, err := keeper2.FilterCDPs(suite.ctx, suite.keeper, paramsOwner)
	suite.Require().NoError(err)
	suite.Require().Equal(1, len(filteredCDPs2))
	suite.Require().Equal(suite.cdps[10].Owner, filteredCDPs2[0].Owner)

	paramsID := types3.NewQueryCdpsParams(1, 100, "", sdk.AccAddress{}, 68, sdk.ZeroDec())
	filteredCDPs3, err := keeper2.FilterCDPs(suite.ctx, suite.keeper, paramsID)
	suite.Require().NoError(err)
	suite.Require().Equal(1, len(filteredCDPs3))
	suite.Require().Equal(suite.cdps[68-1].ID, filteredCDPs3[0].ID)

	ratioCountBtc := 0
	btcRatio := d("2500")
	for _, cdp := range suite.cdps {
		if cdp.Type == "btc-a" {
			absoluteRatio := suite.keeper.CalculateCollateralToDebtRatio(suite.ctx, cdp.Collateral, cdp.Type, cdp.Principal)
			collateralizationRatio, _ := suite.keeper.CalculateCollateralizationRatioFromAbsoluteRatio(suite.ctx, cdp.Type, absoluteRatio, "liquidation")
			if collateralizationRatio.LT(btcRatio) {
				ratioCountBtc += 1
			}
		}
	}
	paramsTypeAndRatio := types3.NewQueryCdpsParams(1, 100, "btc-a", sdk.AccAddress{}, 0, sdk.NewDec(2500))
	filteredCDPs4, err := keeper2.FilterCDPs(suite.ctx, suite.keeper, paramsTypeAndRatio)
	suite.Require().NoError(err)
	suite.Require().Equal(ratioCountBtc, len(filteredCDPs4))
}

func (suite *QuerierTestSuite) TestQueryCdps() {
	ctx := suite.ctx.WithIsCheckTx(false)
	params := types3.NewQueryCdpsParams(1, 100, "btc-a", sdk.AccAddress{}, 0, sdk.ZeroDec())

	query := abci.RequestQuery{
		Path: strings.Join([]string{custom, types3.QuerierRoute, types3.QueryGetCdps}, "/"),
		Data: suite.legacyAmino.MustMarshalJSON(params),
	}

	bz, err := suite.querier(ctx, []string{types3.QueryGetCdps}, query)
	suite.Nil(err)
	suite.NotNil(bz)

	output := types3.AugmentedCDPs{}
	suite.Nil(suite.legacyAmino.UnmarshalJSON(bz, &output))
	suite.Equal(50, len(output))
}

func (suite *QuerierTestSuite) TestQueryTotalPrincipal() {
	ctx := suite.ctx.WithIsCheckTx(false)
	params := types3.NewQueryGetTotalPrincipalParams("btc-a")

	query := abci.RequestQuery{
		Path: strings.Join([]string{custom, types3.QuerierRoute, types3.QueryGetTotalPrincipal}, "/"),
		Data: suite.legacyAmino.MustMarshalJSON(params),
	}

	bz, err := suite.querier(ctx, []string{types3.QueryGetTotalPrincipal}, query)
	suite.Nil(err)
	suite.NotNil(bz)

	var output types3.TotalPrincipals
	suite.Nil(suite.legacyAmino.UnmarshalJSON(bz, &output))
	suite.Equal(1, len(output))
	suite.Equal("btc-a", output[0].CollateralType)
}

func (suite *QuerierTestSuite) TestQueryTotalPrincipalAll() {
	ctx := suite.ctx.WithIsCheckTx(false)
	params := types3.NewQueryGetTotalPrincipalParams("")

	query := abci.RequestQuery{
		Path: strings.Join([]string{custom, types3.QuerierRoute, types3.QueryGetTotalPrincipal}, "/"),
		Data: suite.legacyAmino.MustMarshalJSON(params),
	}

	bz, err := suite.querier(ctx, []string{types3.QueryGetTotalPrincipal}, query)
	suite.Nil(err)
	suite.NotNil(bz)

	var output types3.TotalPrincipals
	suite.Nil(suite.legacyAmino.UnmarshalJSON(bz, &output))

	var outputTypes []string
	for _, c := range output {
		outputTypes = append(outputTypes, c.CollateralType)
	}

	suite.Greater(len(output), 0)
	suite.Subset(outputTypes, []string{"btc-a", "xrp-a"})
}

func (suite *QuerierTestSuite) TestQueryTotalCollateral() {
	ctx := suite.ctx.WithIsCheckTx(false)
	params := types3.NewQueryGetTotalCollateralParams("btc-a")

	query := abci.RequestQuery{
		Path: strings.Join([]string{custom, types3.QuerierRoute, types3.QueryGetTotalCollateral}, "/"),
		Data: suite.legacyAmino.MustMarshalJSON(params),
	}

	bz, err := suite.querier(ctx, []string{types3.QueryGetTotalCollateral}, query)
	suite.Nil(err)
	suite.NotNil(bz)

	var output types3.TotalCollaterals
	suite.Nil(suite.legacyAmino.UnmarshalJSON(bz, &output))
	suite.Equal(1, len(output))
	suite.Equal("btc-a", output[0].CollateralType)
}

func (suite *QuerierTestSuite) TestQueryTotalCollateralAll() {
	ctx := suite.ctx.WithIsCheckTx(false)
	params := types3.NewQueryGetTotalCollateralParams("")

	query := abci.RequestQuery{
		Path: strings.Join([]string{custom, types3.QuerierRoute, types3.QueryGetTotalCollateral}, "/"),
		Data: suite.legacyAmino.MustMarshalJSON(params),
	}

	bz, err := suite.querier(ctx, []string{types3.QueryGetTotalCollateral}, query)
	suite.Nil(err)
	suite.NotNil(bz)

	var output types3.TotalCollaterals
	suite.Nil(suite.legacyAmino.UnmarshalJSON(bz, &output))

	var outputTypes []string
	for _, c := range output {
		outputTypes = append(outputTypes, c.CollateralType)
	}

	suite.Greater(len(output), 0)
	suite.Subset(outputTypes, []string{"btc-a", "xrp-a"})
}

func TestQuerierTestSuite(t *testing.T) {
	suite.Run(t, new(QuerierTestSuite))
}
