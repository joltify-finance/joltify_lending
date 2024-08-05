package testutil

import (
	"context"

	"cosmossdk.io/log"
	sdkmath "cosmossdk.io/math"
	"github.com/joltify-finance/joltify_lending/x/third_party/auction/keeper"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/auction/types"
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/joltify-finance/joltify_lending/app"
)

// Suite implements a test suite for the mint :"wmodule integration tests
type Suite struct {
	suite.Suite

	Keeper        keeper.Keeper
	BankKeeper    bankkeeper.Keeper
	AccountKeeper authkeeper.AccountKeeper
	App           app.TestApp
	Ctx           context.Context
	Addrs         []sdk.AccAddress
	ModAcc        *authtypes.ModuleAccount
}

// SetupTest instantiates a new app, keepers, and sets suite state
func (suite *Suite) SetupTest(numAddrs int) {
	config := sdk.GetConfig()
	app.SetBech32AddressPrefixes(config)
	tg := log.NewTestLogger(suite.T())
	tApp := app.NewTestApp(tg, suite.T().TempDir())

	_, addrs := app.GeneratePrivKeyAddressPairs(numAddrs)

	// Fund liquidator module account
	coins := sdk.NewCoins(
		sdk.NewCoin("token1", sdkmath.NewInt(100)),
		sdk.NewCoin("token2", sdkmath.NewInt(100)),
	)

	ctx := tApp.NewContext(true)

	modName := "jolt"
	modBaseAcc := authtypes.NewBaseAccount(authtypes.NewModuleAddress(modName), nil, 0, 0)
	modAcc := authtypes.NewModuleAccount(modBaseAcc, modName, []string{authtypes.Minter, authtypes.Burner}...)
	suite.ModAcc = modAcc

	authGS := app.NewFundedGenStateWithSameCoinsWithModuleAccount(tApp.AppCodec(), coins, addrs, modAcc)

	params := types2.NewParams(
		types2.DefaultMaxAuctionDuration,
		types2.DefaultForwardBidDuration,
		types2.DefaultReverseBidDuration,
		types2.DefaultIncrement,
		types2.DefaultIncrement,
		types2.DefaultIncrement,
	)

	auctionGs, err := types2.NewGenesisState(types2.DefaultNextAuctionID, params, []types2.GenesisAuction{})
	suite.Require().NoError(err)

	moduleGs := tApp.AppCodec().MustMarshalJSON(auctionGs)
	gs := app.GenesisState{types2.ModuleName: moduleGs}

	//var genAcc []authtypes.GenesisAccount
	//for _, el := range suite.Addrs {
	//	fmt.Printf(">>>>>>>>##%v\n", el.String())
	//	b := authtypes.NewBaseAccount(el, nil, 0, 0)
	//	genAcc = append(genAcc, b)
	//}
	//
	//addr, err := sdk.AccAddressFromBech32("jolt1ze7y9qwdddejmy7jlw4cymqqlt2wh05ypmjzfy")
	//suite.Require().NoError(err)
	//b := authtypes.NewBaseAccount(addr, nil, 0, 0)
	//genAcc = append(genAcc, b)
	//
	tApp.InitializeFromGenesisStates(nil, nil, authGS, gs)

	suite.App = tApp
	suite.Ctx = ctx
	suite.Addrs = addrs
	suite.Keeper = tApp.GetAuctionKeeper()
	suite.BankKeeper = tApp.GetBankKeeper()
	suite.AccountKeeper = tApp.GetAccountKeeper()
}

// AddCoinsToAccount adds coins to an account address
func (suite *Suite) AddCoinsToAccount(addr sdk.AccAddress, coins sdk.Coins) {
	ak := suite.App.GetAccountKeeper()
	acc := ak.NewAccountWithAddress(suite.Ctx, addr)
	ak.SetAccount(suite.Ctx, acc)

	err := fundAccount(suite.BankKeeper, suite.Ctx, acc.GetAddress(), coins)
	suite.Require().NoError(err)
}

func fundAccount(bankKeeper bankkeeper.Keeper, ctx context.Context, addr sdk.AccAddress, amounts sdk.Coins) error {
	if err := bankKeeper.MintCoins(ctx, minttypes.ModuleName, amounts); err != nil {
		return err
	}
	return bankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, addr, amounts)
}

// fundModuleAccount is a utility function that funds a module account by
// minting and sending the coins to the address. This should be used for testing
// purposes only!
//
// TODO: Instead of using the mint module account, which has the
// permission of minting, create a "faucet" account. (@fdymylja)
func fundModuleAccount(bankKeeper bankkeeper.Keeper, ctx context.Context, recipientMod string, amounts sdk.Coins) error {
	if err := bankKeeper.MintCoins(ctx, minttypes.ModuleName, amounts); err != nil {
		return err
	}

	return bankKeeper.SendCoinsFromModuleToModule(ctx, minttypes.ModuleName, recipientMod, amounts)
}

// AddCoinsToNamedModule adds coins to a named module account
func (suite *Suite) AddCoinsToNamedModule(moduleName string, amount sdk.Coins) {
	// Does not use suite.BankKeeper.MintCoins as module account would not have permission to mint
	err := fundModuleAccount(suite.BankKeeper, suite.Ctx, moduleName, amount)
	suite.Require().NoError(err)
}

// NewModuleAccountFromAddr creates a new module account from the provided address with the provided balance
// func (suite *Suite) NewModuleAccount(moduleName string, balance sdk.Coins) authtypes.AccountI {
// 	ak := suite.App.GetAccountKeeper()

// 	modAccAddr := authtypes.NewModuleAddress(moduleName)
// 	acc := ak.NewAccountWithAddress(suite.Ctx, modAccAddr)
// 	ak.SetAccount(suite.Ctx, acc)

// 	err := simapp.fundModuleAccount(suite.BankKeeper, suite.Ctx, moduleName, balance)
// 	suite.Require().NoError(err)

// 	return acc
// }

// CheckAccountBalanceEqual asserts that
func (suite *Suite) CheckAccountBalanceEqual(owner sdk.AccAddress, expectedCoins sdk.Coins) {
	balances := suite.BankKeeper.GetAllBalances(suite.Ctx, owner)
	suite.Equal(expectedCoins, balances)
}

// // CheckModuleAccountBalanceEqual asserts that a named module account balance matches the provided coins
// func (suite *Suite) CheckModuleAccountBalanceEqual(moduleName string, coins sdk.Coins) {
// 	balance := suite.BankKeeper.GetAllBalances(
// 		suite.Ctx,
// 		suite.AccountKeeper.GetModuleAddress(moduleName),
// 	)
// 	suite.Equal(coins, balance, fmt.Sprintf("expected module account balance to equal coins %s, but got %s", coins, balance))
// }
