package keeper_test

import (
	"fmt"
	"strconv"
	"testing"

	tmlog "github.com/cometbft/cometbft/libs/log"

	auctionkeeper "github.com/joltify-finance/joltify_lending/x/third_party/auction/keeper"
	"github.com/joltify-finance/joltify_lending/x/third_party/jolt/keeper"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtime "github.com/cometbft/cometbft/types/time"

	"github.com/joltify-finance/joltify_lending/app"
)

// Test suite used for all keeper tests
type KeeperTestSuite struct {
	suite.Suite
	keeper        keeper.Keeper
	auctionKeeper auctionkeeper.Keeper
	app           app.TestApp
	ctx           context.Context
	addrs         []sdk.AccAddress
}

// The default state used by each test
func (suite *KeeperTestSuite) SetupTest() {
	config := sdk.GetConfig()
	app.SetBech32AddressPrefixes(config)

	tApp := app.NewTestApp(tmlog.TestingLogger(), suite.T().TempDir())
	ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})
	tApp.InitializeFromGenesisStates(nil, nil)
	_, addrs := app.GeneratePrivKeyAddressPairs(1)
	k := tApp.GetJoltKeeper()
	suite.app = tApp
	suite.ctx = ctx
	suite.keeper = k
	suite.addrs = addrs
}

func (suite *KeeperTestSuite) TestGetSetDeleteDeposit() {
	dep := types2.NewDeposit(sdk.AccAddress("test"), sdk.NewCoins(sdk.NewCoin("bnb", sdkmath.NewInt(100))),
		types2.SupplyInterestFactors{types2.NewSupplyInterestFactor("", sdkmath.LegacyMustNewDecFromStr("0"))})

	_, f := suite.keeper.GetDeposit(suite.ctx, sdk.AccAddress("test"))
	suite.Require().False(f)

	suite.keeper.SetDeposit(suite.ctx, dep)

	testDeposit, f := suite.keeper.GetDeposit(suite.ctx, sdk.AccAddress("test"))
	suite.Require().True(f)
	suite.Require().Equal(dep, testDeposit)

	suite.Require().NotPanics(func() { suite.keeper.DeleteDeposit(suite.ctx, dep) })

	_, f = suite.keeper.GetDeposit(suite.ctx, sdk.AccAddress("test"))
	suite.Require().False(f)
}

func (suite *KeeperTestSuite) TestIterateDeposits() {
	for i := 0; i < 5; i++ {
		dep := types2.NewDeposit(sdk.AccAddress("test"+fmt.Sprint(i)), sdk.NewCoins(sdk.NewCoin("bnb", sdkmath.NewInt(100))), types2.SupplyInterestFactors{})
		suite.Require().NotPanics(func() { suite.keeper.SetDeposit(suite.ctx, dep) })
	}
	var deposits []types2.Deposit
	suite.keeper.IterateDeposits(suite.ctx, func(d types2.Deposit) bool {
		deposits = append(deposits, d)
		return false
	})
	suite.Require().Equal(5, len(deposits))
}

func (suite *KeeperTestSuite) TestGetSetDeleteInterestRateModel() {
	denom := "test"
	model := types2.NewInterestRateModel(sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyMustNewDecFromStr("2"), sdkmath.LegacyMustNewDecFromStr("0.8"), sdkmath.LegacyMustNewDecFromStr("10"))
	borrowLimit := types2.NewBorrowLimit(false, sdkmath.LegacyMustNewDecFromStr("0.2"), sdkmath.LegacyMustNewDecFromStr("0.5"))
	moneyMarket := types2.NewMoneyMarket(denom, borrowLimit, denom+":usd", sdkmath.NewInt(1000000), model, sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyZeroDec())

	_, f := suite.keeper.GetMoneyMarket(suite.ctx, denom)
	suite.Require().False(f)

	suite.keeper.SetMoneyMarket(suite.ctx, denom, moneyMarket)

	testMoneyMarket, f := suite.keeper.GetMoneyMarket(suite.ctx, denom)
	suite.Require().True(f)
	suite.Require().Equal(moneyMarket, testMoneyMarket)

	suite.Require().NotPanics(func() { suite.keeper.DeleteMoneyMarket(suite.ctx, denom) })

	_, f = suite.keeper.GetMoneyMarket(suite.ctx, denom)
	suite.Require().False(f)
}

func (suite *KeeperTestSuite) TestIterateInterestRateModels() {
	testDenom := "test"
	var setMMs types2.MoneyMarkets
	var setDenoms []string

	var cleanDenoms []string
	suite.keeper.IterateMoneyMarkets(suite.ctx, func(denom string, i types2.MoneyMarket) bool {
		cleanDenoms = append(cleanDenoms, denom)
		return false
	})

	for _, el := range cleanDenoms {
		suite.keeper.DeleteMoneyMarket(suite.ctx, el)
	}

	for i := 0; i < 5; i++ {
		// Initialize a new money market
		denom := testDenom + strconv.Itoa(i)
		model := types2.NewInterestRateModel(sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyMustNewDecFromStr("2"), sdkmath.LegacyMustNewDecFromStr("0.8"), sdkmath.LegacyMustNewDecFromStr("10"))
		borrowLimit := types2.NewBorrowLimit(false, sdkmath.LegacyMustNewDecFromStr("0.2"), sdkmath.LegacyMustNewDecFromStr("0.5"))
		moneyMarket := types2.NewMoneyMarket(denom, borrowLimit, denom+":usd", sdkmath.NewInt(1000000), model, sdkmath.LegacyMustNewDecFromStr("0.05"), sdkmath.LegacyZeroDec())

		// Store money market in the module's store
		suite.Require().NotPanics(func() { suite.keeper.SetMoneyMarket(suite.ctx, denom, moneyMarket) })

		// Save the denom and model
		setDenoms = append(setDenoms, denom)
		setMMs = append(setMMs, moneyMarket)
	}

	var seenMMs types2.MoneyMarkets
	var seenDenoms []string
	suite.keeper.IterateMoneyMarkets(suite.ctx, func(denom string, i types2.MoneyMarket) bool {
		seenDenoms = append(seenDenoms, denom)
		seenMMs = append(seenMMs, i)
		return false
	})

	suite.Require().Equal(setMMs, seenMMs)
	suite.Require().Equal(setDenoms, seenDenoms)
}

func (suite *KeeperTestSuite) TestGetSetBorrowedCoins() {
	suite.keeper.SetBorrowedCoins(suite.ctx, sdk.Coins{c("ujolt", 123)})

	coins, found := suite.keeper.GetBorrowedCoins(suite.ctx)
	suite.Require().True(found)
	suite.Require().Len(coins, 1)
	suite.Require().Equal(coins, cs(c("ujolt", 123)))
}

func (suite *KeeperTestSuite) TestGetSetBorrowedCoins_Empty() {
	coins, found := suite.keeper.GetBorrowedCoins(suite.ctx)
	suite.Require().False(found)
	suite.Require().Empty(coins)

	// None set and setting empty coins should both be the same
	suite.keeper.SetBorrowedCoins(suite.ctx, sdk.Coins{})

	coins, found = suite.keeper.GetBorrowedCoins(suite.ctx)
	suite.Require().False(found)
	suite.Require().Empty(coins)
}

func (suite *KeeperTestSuite) getAccountCoins(acc authtypes.AccountI) sdk.Coins {
	bk := suite.app.GetBankKeeper()
	return bk.GetAllBalances(suite.ctx, acc.GetAddress())
}

func (suite *KeeperTestSuite) getAccount(addr sdk.AccAddress) authtypes.AccountI {
	ak := suite.app.GetAccountKeeper()
	return ak.GetAccount(suite.ctx, addr)
}

func (suite *KeeperTestSuite) getAccountAtCtx(addr sdk.AccAddress, ctx context.Context) authtypes.AccountI {
	ak := suite.app.GetAccountKeeper()
	return ak.GetAccount(ctx, addr)
}

func (suite *KeeperTestSuite) getModuleAccount(name string) authtypes.ModuleAccountI {
	ak := suite.app.GetAccountKeeper()
	return ak.GetModuleAccount(suite.ctx, name)
}

func (suite *KeeperTestSuite) getModuleAccountAtCtx(name string, ctx context.Context) authtypes.ModuleAccountI {
	ak := suite.app.GetAccountKeeper()
	return ak.GetModuleAccount(ctx, name)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
