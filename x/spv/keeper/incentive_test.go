package keeper_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/joltify-finance/joltify_lending/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/utils"
	spvkeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/stretchr/testify/suite"
)

// Test suite used for all keeper tests
type IncentiveSuite struct {
	suite.Suite
	keeper          *spvkeeper.Keeper
	nftKeeper       types.NFTKeeper
	bankKeeper      types.BankKeeper
	IncentiveKeeper keeper.FakeIncentiveKeeper
	app             types.MsgServer
	ctx             context.Context
}

func TestIncentiveSuite(t *testing.T) {
	suite.Run(t, new(IncentiveSuite))
}

// The default state used by each test
func (suite *IncentiveSuite) SetupTest() {
	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)

	lapp, k, nftKeeper, bankKeeper, _, incentiveKeeper, wctx := setupMsgServerWithIncentiveKeeper(suite.T())
	ctx := sdk.UnwrapSDKContext(wctx)
	suite.ctx = ctx
	suite.keeper = k
	suite.nftKeeper = nftKeeper
	suite.bankKeeper = bankKeeper
	suite.IncentiveKeeper = incentiveKeeper
	suite.app = lapp
}

func (suite *IncentiveSuite) TestUpdateIncentive() {
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 2, PoolName: "hello", Apy: []string{"0.15", "0.15"}, TargetTokenAmount: sdk.Coins{sdk.NewCoin("ausdc", sdk.NewInt(1*1e6)), sdk.NewCoin("ausdc", sdk.NewInt(1e6))}}
	resp, err := suite.app.CreatePool(suite.ctx, &req)
	suite.Require().NoError(err)

	poolInfo, found := suite.keeper.GetPools(suite.ctx, resp.PoolIndex[0])
	suite.Require().True(found)
	poolInfo.CurrentPoolTotalBorrowCounter = 0
	poolInfo.PoolTotalBorrowLimit = 10
	poolInfo.TargetAmount = sdk.NewCoin("ausdc", sdk.NewInt(1*1e6))
	suite.keeper.SetPool(suite.ctx, poolInfo)

	depositorPool := resp.PoolIndex[0]
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

	// now we deposit some token and it should be enough to borrow
	creator1 := "jolt166yyvsypvn6cwj2rc8sme4dl6v0g62hn3862kl"
	depositAmount := sdk.NewCoin("ausdc", sdk.NewInt(4e5))
	msgDepositUser1 := &types.MsgDeposit{
		Creator:   creator1,
		PoolIndex: depositorPool,
		Token:     depositAmount,
	}

	_, err = suite.app.Deposit(suite.ctx, msgDepositUser1)
	suite.Require().NoError(err)

	// now we borrow 2e5
	borrow := &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: depositorPool, BorrowAmount: sdk.NewCoin("ausdc", sdk.NewIntFromUint64(1.34e5))}
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	poolInfo, ok := suite.keeper.GetPools(suite.ctx, depositorPool)
	suite.Require().True(ok)
	fmt.Printf(">>>%v\n", poolInfo.BorrowedAmount)

	// now we test the incentive
	suite.keeper.UpdateIncentive(suite.ctx, poolInfo)
	ret := suite.IncentiveKeeper.GetPoolIncentive()
	suite.Require().Equal(0, len(ret))

	// now we set the incentive
	params := suite.keeper.GetParams(suite.ctx)
	params.Incentives = []types.Incentive{{Poolid: "t1", Spy: "0.1"}}
	suite.keeper.SetParams(suite.ctx, params)
	suite.keeper.UpdateIncentive(suite.ctx, poolInfo)
	ret = suite.IncentiveKeeper.GetPoolIncentive()
	suite.Require().Equal(0, len(ret))

	// now we set the incentive for the given pool
	params = suite.keeper.GetParams(suite.ctx)
	spy := "1.01234"
	params.Incentives = append(params.Incentives, types.Incentive{Poolid: poolInfo.Index, Spy: spy})
	suite.keeper.SetParams(suite.ctx, params)
	suite.keeper.UpdateIncentive(suite.ctx, poolInfo)
	ret = suite.IncentiveKeeper.GetPoolIncentive()

	spvDev := sdk.MustNewDecFromStr(spy)
	jotlM := sdk.MustNewDecFromStr("0.7")
	incentives := sdk.NewDecFromInt(poolInfo.BorrowedAmount.Amount).Mul(spvDev.Sub(sdk.OneDec())).Quo(jotlM).TruncateInt()

	aa, ok := ret[poolInfo.Index]

	fmt.Printf(">>>>%v--%v\n", incentives.String(), aa)
	suite.Require().True(ok)
	suite.Require().True(incentives.Equal(aa.AmountOf("ujolt")))

	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(time.Second * 5))
	// now we borrow more money and the incentive should be updated
	borrow = &types.MsgBorrow{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", PoolIndex: depositorPool, BorrowAmount: sdk.NewCoin("ausdc", sdk.NewIntFromUint64(1.44e5))}
	_, err = suite.app.Borrow(suite.ctx, borrow)
	suite.Require().NoError(err)

	poolInfo, found = suite.keeper.GetPools(suite.ctx, depositorPool)
	suite.Require().True(found)
	suite.keeper.UpdateIncentive(suite.ctx, poolInfo)
	ret = suite.IncentiveKeeper.GetPoolIncentive()
	spvDev = sdk.MustNewDecFromStr(spy)
	jotlM = sdk.MustNewDecFromStr("0.7")
	fmt.Printf(">>>borrowed new %v\n", poolInfo.BorrowedAmount.Amount)
	incentives = sdk.NewDecFromInt(poolInfo.BorrowedAmount.Amount).Mul(spvDev.Sub(sdk.OneDec())).Quo(jotlM).TruncateInt()

	aa, ok = ret[poolInfo.Index]

	fmt.Printf(">>>>%v--%v\n", incentives.String(), aa)
	suite.Require().True(ok)
	suite.Require().True(incentives.Equal(aa.AmountOf("ujolt")))
}
