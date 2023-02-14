package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/utils"
	spvkeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/stretchr/testify/suite"
)

// Test suite used for all keeper tests
type payPrincipalSuite struct {
	suite.Suite
	keeper       *spvkeeper.Keeper
	nftKeeper    types.NFTKeeper
	app          types.MsgServer
	ctx          sdk.Context
	investors    []string
	investorPool string
}

func setupPools(suite *payPrincipalSuite) {
	// create the first pool apy 7.8%
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 3, PoolName: "hello", Apy: "0.15", TargetTokenAmount: sdk.NewCoin("ausdc", sdk.NewInt(3*1e9))}
	resp, err := suite.app.CreatePool(suite.ctx, &req)
	suite.Require().NoError(err)

	depositorPool := resp.PoolIndex[0]

	suite.investorPool = depositorPool

	_, err = suite.app.ActivePool(suite.ctx, types.NewMsgActivePool("jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", resp.PoolIndex[0]))
	suite.Require().NoError(err)

	_, err = suite.app.ActivePool(suite.ctx, types.NewMsgActivePool("jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", resp.PoolIndex[1]))
	suite.Require().NoError(err)

	req2 := types.MsgAddInvestors{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: resp.PoolIndex[0],
		InvestorID: []string{"2"}}
	_, err = suite.app.AddInvestors(suite.ctx, &req2)
	suite.Require().NoError(err)

	creator1 := "jolt166yyvsypvn6cwj2rc8sme4dl6v0g62hn3862kl"
	creator2 := "jolt1kkujrm0lqeu0e5va5f6mmwk87wva0k8cmam8jq"

	suite.investors = []string{creator1, creator2}

}

// The default state used by each test
func (suite *payPrincipalSuite) SetupTest() {

	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)
	app, k, nftKeeper, wctx := setupMsgServer(suite.T())
	ctx := sdk.UnwrapSDKContext(wctx)
	// create the first pool apy 7.8%

	suite.ctx = ctx
	suite.keeper = k
	suite.app = app
	suite.nftKeeper = nftKeeper

}
func TestPayPrincipalInterest(t *testing.T) {
	suite.Run(t, new(payPrincipalSuite))
}

func (suite *payPrincipalSuite) TestWithExpectedErrors() {
	setupPools(suite)

	req := types.MsgPayPrincipal{
		Creator:   "invalid",
		PoolIndex: suite.investorPool,
		Token:     sdk.NewCoin("abc", sdk.OneInt()),
	}

	_, err := suite.app.PayPrincipal(suite.ctx, &req)
	suite.Require().ErrorContains(err, "invalid address")

	req.Creator = "jolt1p3jl6udk43vw0cvc5hjqrpnncsqmsz56wd32z8"
	req.PoolIndex = "232"

	_, err = suite.app.PayPrincipal(suite.ctx, &req)
	suite.Require().ErrorContains(err, "pool cannot be found")

	req.PoolIndex = suite.investorPool
	_, err = suite.app.PayPrincipal(suite.ctx, &req)
	suite.Require().ErrorContains(err, "invalid token demo, want")

	req.PoolIndex = suite.investorPool
	msgDepositUser1 := &types.MsgDeposit{Creator: suite.investors[1],
		PoolIndex: suite.investorPool,
		Token:     sdk.NewCoin("ausdc", sdk.NewIntFromUint64(2e5))}

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, BorrowAmount: sdk.NewCoin("ausdc", sdk.NewIntFromUint64(1.34e5))}

	borrow.BorrowAmount = sdk.NewCoin(borrow.BorrowAmount.Denom, sdk.NewInt(1.2e5))
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	req.Token = sdk.NewCoin("ausdc", sdk.NewIntFromUint64(211))
	_, err = suite.app.PayPrincipal(suite.ctx, &req)
	suite.Require().ErrorContains(err, "not enough interest to be paid to close the pool")

	suite.app.RepayInterest(suite.ctx, &types.MsgRepayInterest{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, Token: sdk.NewCoin("ausdc", sdk.NewIntFromUint64(1.2e5))})
	req.Token = sdk.NewCoin("ausdc", sdk.NewIntFromUint64(2e5))
	_, err = suite.app.PayPrincipal(suite.ctx, &req)
	suite.Require().NoError(err)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	suite.Require().EqualValues(poolInfo.GetEscrowPrincipalAmount().Amount, req.Token.Amount)

}
