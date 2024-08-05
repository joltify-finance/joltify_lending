package testutil

import (
	"errors"
	"fmt"
	"time"

	tmlog "cosmossdk.io/log"

	incentivekeeper "github.com/joltify-finance/joltify_lending/x/third_party/incentive/keeper"
	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"
	hardkeeper "github.com/joltify-finance/joltify_lending/x/third_party/jolt/keeper"
	hardtypes "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"

	abcitypes "github.com/cometbft/cometbft/abci/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/suite"

	"github.com/joltify-finance/joltify_lending/app"
)

var testChainID = "joltifytest_888-1"

type IntegrationTester struct {
	suite.Suite
	App app.TestApp
	Ctx context.Context
}

func (suite *IntegrationTester) SetupSuite() {
	config := sdk.GetConfig()
	app.SetBech32AddressPrefixes(config)
}

func (suite *IntegrationTester) SetApp() {
	suite.App = app.NewTestApp(tmlog.TestingLogger(), suite.T().TempDir())
}

func (suite *IntegrationTester) StartChain(genAccs []authtypes.GenesisAccount, coins sdk.Coins, genesisTime time.Time, genesisStates ...app.GenesisState) {
	suite.App.InitializeFromGenesisStatesWithTimeAndChainID(
		genesisTime,
		testChainID, genAccs, coins,
		genesisStates...,
	)

	suite.Ctx = suite.App.NewContext(false, tmproto.Header{Height: 1, Time: genesisTime, ChainID: testChainID})
	suite.Ctx = suite.Ctx.WithBlockGasMeter(sdk.NewInfiniteGasMeter())
	suite.Ctx = suite.Ctx.WithConsensusParams(app.DefaultConsensusParams)
}

func (suite *IntegrationTester) NextBlockAt(blockTime time.Time) {
	if !sdk.UnwrapSDKContext(suite.ctx).BlockTime().Before(blockTime) {
		panic(fmt.Sprintf("new block time %s must be after current %s", blockTime, sdk.UnwrapSDKContext(suite.ctx).BlockTime()))
	}
	blockHeight := suite.Ctx.BlockHeight() + 1

	_ = suite.App.EndBlocker(suite.Ctx, abcitypes.RequestEndBlock{})

	suite.Ctx = sdk.UnwrapSDKContext(suite.ctx).WithBlockTime(blockTime).WithBlockHeight(blockHeight).WithChainID(testChainID)
	suite.Ctx = suite.Ctx.WithBlockGasMeter(sdk.NewInfiniteGasMeter())
	suite.Ctx = suite.Ctx.WithConsensusParams(app.DefaultConsensusParams)

	_ = suite.App.BeginBlocker(suite.Ctx, abcitypes.RequestBeginBlock{}) // height and time in RequestBeginBlock are ignored by module begin blockers
}

func (suite *IntegrationTester) NextBlockAfter(blockDuration time.Duration) {
	suite.NextBlockAt(sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(blockDuration))
}

func (suite *IntegrationTester) DeliverIncentiveMsg(msg sdk.Msg) error {
	msgServer := incentivekeeper.NewMsgServerImpl(suite.App.GetIncentiveKeeper())

	var err error

	switch msg := msg.(type) {
	case *types.MsgClaimJoltReward:
		_, err = msgServer.ClaimJoltReward(sdk.WrapSDKContext(suite.Ctx), msg)
	case *types.MsgClaimSwapReward:
		_, err = msgServer.ClaimSwapReward(sdk.WrapSDKContext(suite.Ctx), msg)
		//	case *types.MsgClaimUSDXMintingReward:
		//		_, err = msgServer.ClaimUSDXMintingReward(sdk.WrapSDKContext(suite.Ctx), msg)
	// case *types.MsgClaimDelegatorReward:
	//	_, err = msgServer.ClaimDelegatorReward(sdk.WrapSDKContext(suite.Ctx), msg)
	default:
		panic("unhandled incentive msg")
	}

	return err
}

func (suite *IntegrationTester) DeliverMsgCreateValidator(address sdk.ValAddress, selfDelegation sdk.Coin) error {
	msg, err := stakingtypes.NewMsgCreateValidator(
		address,
		ed25519.GenPrivKey().PubKey(),
		selfDelegation,
		stakingtypes.Description{},
		stakingtypes.NewCommissionRates(sdkmath.LegacyZeroDec(), sdkmath.LegacyZeroDec(), sdkmath.LegacyZeroDec()),
		sdkmath.NewInt(1_000_000),
	)
	if err != nil {
		return err
	}

	stk := suite.App.GetStakingKeeper()
	msgServer := stakingkeeper.NewMsgServerImpl(&stk)
	_, err = msgServer.CreateValidator(sdk.WrapSDKContext(suite.Ctx), msg)

	return err
}

func (suite *IntegrationTester) DeliverMsgDelegate(delegator sdk.AccAddress, validator sdk.ValAddress, amount sdk.Coin) error {
	msg := stakingtypes.NewMsgDelegate(
		delegator,
		validator,
		amount,
	)

	stk := suite.App.GetStakingKeeper()
	msgServer := stakingkeeper.NewMsgServerImpl(&stk)
	_, err := msgServer.Delegate(sdk.WrapSDKContext(suite.Ctx), msg)
	return err
}

//
//func (suite *IntegrationTester) DeliverSwapMsgDeposit(depositor sdk.AccAddress, tokenA, tokenB sdk.Coin, slippage sdkmath.LegacyDec) error {
//	msg := swaptypes.NewMsgDeposit(
//		depositor.String(),
//		tokenA,
//		tokenB,
//		slippage,
//		sdk.UnwrapSDKContext(suite.ctx).BlockTime().Add(time.Hour).Unix(), // ensure msg will not fail due to short deadline
//	)
//	msgServer := swapkeeper.NewMsgServerImpl(suite.App.GetSwapKeeper())
//	_, err := msgServer.Deposit(sdk.WrapSDKContext(suite.Ctx), msg)
//
//	return err
//}

func (suite *IntegrationTester) DeliverJoltMsgDeposit(owner sdk.AccAddress, deposit sdk.Coins) error {
	msg := hardtypes.NewMsgDeposit(owner, deposit)
	msgServer := hardkeeper.NewMsgServerImpl(suite.App.GetJoltKeeper())

	_, err := msgServer.Deposit(sdk.WrapSDKContext(suite.Ctx), &msg)
	return err
}

func (suite *IntegrationTester) DeliverJoltMsgBorrow(owner sdk.AccAddress, borrow sdk.Coins) error {
	msg := hardtypes.NewMsgBorrow(owner, borrow)
	msgServer := hardkeeper.NewMsgServerImpl(suite.App.GetJoltKeeper())

	_, err := msgServer.Borrow(sdk.WrapSDKContext(suite.Ctx), &msg)
	return err
}

func (suite *IntegrationTester) DeliverHardMsgRepay(owner sdk.AccAddress, repay sdk.Coins) error {
	msg := hardtypes.NewMsgRepay(owner, owner, repay)
	msgServer := hardkeeper.NewMsgServerImpl(suite.App.GetJoltKeeper())

	_, err := msgServer.Repay(sdk.WrapSDKContext(suite.Ctx), &msg)
	return err
}

func (suite *IntegrationTester) DeliverJoltMsgWithdraw(owner sdk.AccAddress, withdraw sdk.Coins) error {
	msg := hardtypes.NewMsgWithdraw(owner, withdraw)
	msgServer := hardkeeper.NewMsgServerImpl(suite.App.GetJoltKeeper())

	_, err := msgServer.Withdraw(sdk.WrapSDKContext(suite.Ctx), &msg)
	return err
}

func (suite *IntegrationTester) GetAccount(addr sdk.AccAddress) authtypes.AccountI {
	ak := suite.App.GetAccountKeeper()
	return ak.GetAccount(suite.Ctx, addr)
}

func (suite *IntegrationTester) GetModuleAccount(name string) authtypes.ModuleAccountI {
	ak := suite.App.GetAccountKeeper()
	return ak.GetModuleAccount(suite.Ctx, name)
}

func (suite *IntegrationTester) GetBalance(address sdk.AccAddress) sdk.Coins {
	bk := suite.App.GetBankKeeper()
	return bk.GetAllBalances(suite.Ctx, address)
}

func (suite *IntegrationTester) ErrorIs(err, target error) bool {
	return suite.Truef(errors.Is(err, target), "err didn't match: %s, it was: %s", target, err)
}

func (suite *IntegrationTester) BalanceEquals(address sdk.AccAddress, expected sdk.Coins) {
	bk := suite.App.GetBankKeeper()
	suite.Equalf(
		expected,
		bk.GetAllBalances(suite.Ctx, address),
		"expected account balance to equal coins %s, but got %s",
		expected,
		bk.GetAllBalances(suite.Ctx, address),
	)
}

func (suite *IntegrationTester) BalanceInEpsilon(address sdk.AccAddress, expected sdk.Coins, epsilon float64) {
	actual := suite.GetBalance(address)

	allDenoms := expected.Add(actual...)
	for _, coin := range allDenoms {
		suite.InEpsilonf(
			expected.AmountOf(coin.Denom).Int64(),
			actual.AmountOf(coin.Denom).Int64(),
			epsilon,
			"expected balance to be within %f%% of coins %s, but got %s", epsilon*100, expected, actual,
		)
	}
}

func (suite *IntegrationTester) VestingPeriodsEqual(address sdk.AccAddress, expectedPeriods []vestingtypes.Period) {
	acc := suite.App.GetAccountKeeper().GetAccount(suite.Ctx, address)
	suite.Require().NotNil(acc, "expected vesting account not to be nil")
	vacc, ok := acc.(*vestingtypes.PeriodicVestingAccount)
	suite.Require().True(ok, "expected vesting account to be type PeriodicVestingAccount")
	suite.Equal(expectedPeriods, vacc.VestingPeriods)
}

func (suite *IntegrationTester) JoltRewardEquals(owner sdk.AccAddress, expected sdk.Coins) {
	claim, found := suite.App.GetIncentiveKeeper().GetJoltLiquidityProviderClaim(suite.Ctx, owner)
	suite.Require().Truef(found, "expected delegator claim to be found for %s", owner)
	suite.Equalf(expected, claim.Reward, "expected delegator claim reward to be %s, but got %s", expected, claim.Reward)
}
