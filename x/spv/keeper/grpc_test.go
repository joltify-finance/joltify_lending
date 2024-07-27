package keeper_test

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/utils"
	spvkeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

// Test suite used for all keeper tests
type querySuite struct {
	suite.Suite
	keeper       *spvkeeper.Keeper
	nftKeeper    types.NFTKeeper
	bankKeeper   types.BankKeeper
	app          types.MsgServer
	ctx          context.Context
	investors    []string
	investorPool string
}

func setupPoolForQueryTest(suite *querySuite) {
	// create the first pool apy 7.8%
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 3, PoolName: "hello", Apy: []string{"0.15", "0.12"}, TargetTokenAmount: sdk.Coins{sdk.NewCoin("ausdc", sdkmath.NewInt(3*1e9)), sdk.NewCoin("ausdc", sdkmath.NewInt(3*1e9))}}
	resp, err := suite.app.CreatePool(suite.ctx, &req)
	suite.Require().NoError(err)

	depositorPool := resp.PoolIndex[0]

	suite.investorPool = depositorPool

	_, err = suite.app.ActivePool(suite.ctx, types.NewMsgActivePool("jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", resp.PoolIndex[0]))
	suite.Require().NoError(err)

	_, err = suite.app.ActivePool(suite.ctx, types.NewMsgActivePool("jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", resp.PoolIndex[1]))
	suite.Require().NoError(err)

	req2 := types.MsgAddInvestors{
		Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: resp.PoolIndex[0],
		InvestorID: []string{"2"},
	}
	_, err = suite.app.AddInvestors(suite.ctx, &req2)
	suite.Require().NoError(err)

	creator1 := "jolt166yyvsypvn6cwj2rc8sme4dl6v0g62hn3862kl"
	creator2 := "jolt1kkujrm0lqeu0e5va5f6mmwk87wva0k8cmam8jq"

	suite.investors = []string{creator1, creator2}

	poolInfo, found := suite.keeper.GetPools(suite.ctx, depositorPool)
	suite.Require().True(found)
	poolInfo.PoolTotalBorrowLimit = 1000
	poolInfo.PoolLockedSeconds = 3600
	suite.keeper.SetPool(suite.ctx, poolInfo)
}

// The default state used by each test
func (suite *querySuite) SetupTest() {
	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)
	lapp, k, nftKeeper, bankKeeper, _, wctx := setupMsgServer(suite.T())
	ctx := sdk.UnwrapSDKContext(wctx)
	// create the first pool apy 7.8%

	suite.ctx = ctx
	suite.keeper = k
	suite.app = lapp
	suite.nftKeeper = nftKeeper
	suite.bankKeeper = bankKeeper
}

func TestQuerySuitTestSuite(t *testing.T) {
	suite.Run(t, new(querySuite))
}

func (suite *querySuite) TestAllQuery() {
	setupPoolForQueryTest(suite)
	_, err := suite.keeper.QueryPool(suite.ctx, nil)
	suite.Require().ErrorContains(err, "invalid request")

	_, err = suite.keeper.ListPools(suite.ctx, nil)
	suite.Require().ErrorContains(err, "invalid request")

	respListPool, err := suite.keeper.ListPools(suite.ctx, &types.QueryListPoolsRequest{Pagination: &query.PageRequest{Limit: 1}})

	suite.Require().Len(respListPool.PoolsInfo, 1)

	respListPool, err = suite.keeper.ListPools(suite.ctx, &types.QueryListPoolsRequest{Pagination: &query.PageRequest{}})

	suite.Require().Len(respListPool.PoolsInfo, 2)

	resp, err := suite.keeper.QueryPool(suite.ctx, &types.QueryQueryPoolRequest{PoolIndex: suite.investorPool})
	suite.Require().NoError(err)
	suite.Require().Equal(resp.PoolInfo.PoolName, "hello-junior")

	resp, err = suite.keeper.QueryPool(suite.ctx, &types.QueryQueryPoolRequest{PoolIndex: "invalid"})
	suite.Require().ErrorContains(err, "index cannot be found")

	_, err = suite.keeper.AllowedPools(suite.ctx, nil)
	suite.Require().ErrorContains(err, "invalid request")

	respPool, err := suite.keeper.AllowedPools(suite.ctx, &types.QueryAllowedPoolsRequest{WalletAddress: "invalid address"})
	suite.Require().NoError(err)
	suite.Require().Len(respPool.PoolsIndex, 0)

	respPool, err = suite.keeper.AllowedPools(suite.ctx, &types.QueryAllowedPoolsRequest{WalletAddress: suite.investors[1]})
	suite.Require().NoError(err)
	suite.Require().Lenf(respPool.PoolsIndex, 2, "should be 2")

	respPool, err = suite.keeper.AllowedPools(suite.ctx, &types.QueryAllowedPoolsRequest{WalletAddress: suite.investors[0]})
	suite.Require().NoError(err)
	suite.Require().Lenf(respPool.PoolsIndex, 2, "should be 2")

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * 5))

	_, err = suite.keeper.WithdrawalPrincipal(suite.ctx, &types.QuerywithdrawalPrincipalRequest{PoolIndex: suite.investorPool, WalletAddress: "1234invalid"})
	suite.Require().ErrorContains(err, "invalid address")

	_, err = suite.keeper.WithdrawalPrincipal(suite.ctx, &types.QuerywithdrawalPrincipalRequest{PoolIndex: "invalid pool", WalletAddress: suite.investors[0]})
	suite.Require().ErrorContains(err, "depositor not found for pool")

	respWithdrawal, err := suite.keeper.WithdrawalPrincipal(suite.ctx, &types.QuerywithdrawalPrincipalRequest{PoolIndex: suite.investorPool, WalletAddress: suite.investors[0]})
	suite.Require().ErrorContains(err, "depositor not found for pool")

	poolInfo, found := suite.keeper.GetPools(suite.ctx, suite.investorPool)
	suite.Require().True(found)
	poolInfo.PoolTotalBorrowLimit = 100
	poolInfo.TargetAmount = sdk.NewCoin("ausdc", sdkmath.NewInt(600000))
	suite.keeper.SetPool(suite.ctx, poolInfo)

	// now we deposit some token and it should be enough to borrow
	creator1 := suite.investors[0]
	creator2 := suite.investors[1]
	depositAmount := sdk.NewCoin("ausdc", sdkmath.NewInt(4e5))
	// suite.Require().NoError(err)
	msgDepositUser1 := &types.MsgDeposit{
		Creator:   creator1,
		PoolIndex: suite.investorPool,
		Token:     depositAmount,
	}

	// user two deposit half of the amount of the user 1
	msgDepositUser2 := &types.MsgDeposit{
		Creator:   creator2,
		PoolIndex: suite.investorPool,
		Token:     depositAmount.SubAmount(sdk.NewInt(2e5)),
	}

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser2)
	suite.Require().NoError(err)

	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: suite.investorPool, BorrowAmount: sdk.NewCoin("ausdc", sdkmath.NewIntFromUint64(1.34e5))}

	// now we borrow 1.34e5
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	addr, _ := sdk.AccAddressFromBech32(suite.investors[0])
	depositor, found := suite.keeper.GetDepositor(suite.ctx, suite.investorPool, addr)
	suite.Require().True(found)

	_, err = suite.keeper.Depositor(suite.ctx, &types.QueryDepositorRequest{WalletAddress: "invalid", DepositPoolIndex: suite.investorPool})
	suite.Require().ErrorContains(err, "invalid address")

	_, err = suite.keeper.Depositor(suite.ctx, &types.QueryDepositorRequest{WalletAddress: suite.investors[0], DepositPoolIndex: "invalid"})
	suite.Require().ErrorContains(err, "depositor not found")

	respDepositor, err := suite.keeper.Depositor(suite.ctx, &types.QueryDepositorRequest{WalletAddress: suite.investors[0], DepositPoolIndex: suite.investorPool})

	respWithdrawal, err = suite.keeper.WithdrawalPrincipal(suite.ctx, &types.QuerywithdrawalPrincipalRequest{PoolIndex: suite.investorPool, WalletAddress: suite.investors[0]})

	suite.Require().Equal(respWithdrawal.Amount, depositor.WithdrawalAmount.String())
	suite.Require().Equal(respDepositor.Depositor.WithdrawalAmount.String(), depositor.WithdrawalAmount.String())

	// query pool investor
	_, err = suite.keeper.PoolInvestors(suite.ctx, nil)
	suite.Require().ErrorContains(err, "invalid request")

	_, err = suite.keeper.PoolInvestors(suite.ctx, &types.QueryPoolInvestorsRequest{PoolIndex: "invalid pool"})
	suite.Require().ErrorContains(err, "pool invalid pool does not exist")

	poolInvestors, err := suite.keeper.PoolInvestors(suite.ctx, &types.QueryPoolInvestorsRequest{PoolIndex: suite.investorPool})
	suite.Require().NoError(err)
	suite.Require().Equal(poolInvestors.Investors[0], "2")

	_, err = suite.keeper.TotalReserve(suite.ctx, &types.QueryTotalReserveRequest{})
	suite.Require().NoError(err)
}
