package keeper_test

import (
	"fmt"
	types2 "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/utils"
	spvkeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMsgRepayInterest(t *testing.T) {

	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)
	app, k, _, wctx := setupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(wctx)

	// create the first pool apy 7.8%
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 1, PoolName: "hello", Apy: "7.8", TargetTokenAmount: sdk.NewCoin("ausdc", sdk.NewInt(322))}
	resp, err := app.CreatePool(ctx, &req)
	require.NoError(t, err)

	poolIndex := resp.PoolIndex
	reqRePayInterest := types.MsgRepayInterest{
		Creator:   "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0",
		PoolIndex: poolIndex[0],
		Token:     sdk.NewCoin("invalid", sdk.NewInt(100)),
	}
	_, err = app.RepayInterest(ctx, &reqRePayInterest)
	require.ErrorContains(t, err, "pool denom ausdc and repay is invalid: inconsistency tokens")

	reqRePayInterest.Creator = "invalid address"
	_, err = app.RepayInterest(ctx, &reqRePayInterest)
	require.ErrorContains(t, err, "invalid address")

	reqRePayInterest = types.MsgRepayInterest{
		Creator:   "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0",
		PoolIndex: poolIndex[0],
		Token:     sdk.NewCoin("ausdc", sdk.NewInt(100)),
	}
	_, err = app.RepayInterest(ctx, &reqRePayInterest)
	require.NoError(t, err)

	poolInfo, found := k.GetPools(ctx, poolIndex[0])
	require.True(t, found)
	poolInfo.PoolStatus = types.PoolInfo_CLOSED
	_, err = app.RepayInterest(ctx, &reqRePayInterest)
	require.NoError(t, err, "pool is not active")

	poolInfo.PoolStatus = types.PoolInfo_INACTIVE
	_, err = app.RepayInterest(ctx, &reqRePayInterest)
	require.NoError(t, err, "pool is not active")

}

func mockBorrow(ctx sdk.Context, nftKeeper types.NFTKeeper, poolInfo *types.PoolInfo, borrowAmount sdk.Coin) {

	classID := fmt.Sprintf("nft-%v", poolInfo.Index[2:])
	poolClass, found := nftKeeper.GetClass(ctx, classID)
	if !found {
		panic("pool class must have already been set")
	}

	latestSeries := len(poolInfo.PoolNFTIds)

	currentBorrowClass := poolClass
	currentBorrowClass.Id = fmt.Sprintf("%v-%v", currentBorrowClass.Id, latestSeries)

	i, err := spvkeeper.CalculateInterestAmount(poolInfo.Apy, int(poolInfo.PayFreq))
	if err != nil {
		panic(err)
	}

	rate := spvkeeper.CalculateInterestRate(poolInfo.Apy, int(poolInfo.PayFreq))
	paymentTime := ctx.BlockTime()
	firstPayment := types.PaymentItem{PaymentTime: paymentTime, PaymentAmount: sdk.NewCoin(borrowAmount.Denom, sdk.NewInt(0))}
	bi := types.BorrowInterest{
		PoolIndex:    poolInfo.Index,
		Apy:          poolInfo.Apy,
		PayFreq:      poolInfo.PayFreq,
		IssueTime:    ctx.BlockTime(),
		Borrowed:     borrowAmount,
		BorrowedLast: borrowAmount,
		MonthlyRatio: i,
		InterestSPY:  rate,
		Payments:     []*types.PaymentItem{&firstPayment},
	}

	data, err := types2.NewAnyWithValue(&bi)
	if err != nil {
		panic(err)
	}
	currentBorrowClass.Data = data
	nftKeeper.SaveClass(ctx, currentBorrowClass)
	poolInfo.PoolNFTIds = append(poolInfo.PoolNFTIds, currentBorrowClass.Id)
}

func TestGetAllInterestToBePaid(t *testing.T) {

	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)
	app, k, nftKeeper, wctx := setupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(wctx)

	// create the first pool apy 7.8%
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 1, PoolName: "hello", Apy: "7.8", TargetTokenAmount: sdk.NewCoin("ausdc", sdk.NewInt(322))}
	resp, err := app.CreatePool(ctx, &req)
	require.NoError(t, err)

	poolIndex := resp.PoolIndex[0]
	samplePool, found := k.GetPools(ctx, poolIndex)
	require.True(t, found)

	samplePool.BorrowableAmount = sdk.NewCoin("ausdc", sdk.NewIntFromUint64(8e12))
	samplePool.PoolStatus = types.PoolInfo_ACTIVE
	mockBorrow(ctx, nftKeeper, &samplePool, sdk.NewCoin("ausdc", sdk.NewIntFromUint64(2e8)))
	k.SetPool(ctx, samplePool)

}
