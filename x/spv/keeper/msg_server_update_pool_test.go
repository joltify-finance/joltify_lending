package keeper_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/utils"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMsgSERvUpdatePool(t *testing.T) {

	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)
	app, k, wctx := setupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(wctx)

	// create the first pool apy 7.8%
	req := types.MsgCreatePool{Creator: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0", ProjectIndex: 1, PoolName: "hello", Apy: "7.8", TargetTokenAmount: sdk.NewCoin("ausdc", sdk.NewInt(322))}
	_, err := app.CreatePool(ctx, &req)
	require.NoError(t, err)

	poolsIndex := []string{"0x907ef9f3822d51da01347e81d1d859b95caf42b18ccc8a266044f5900b8c1371", "0x69da3ffa942f73a8bc9751541c8675cf28a4ffb68ef89559502391e06d37804a"}

	p1, ok := k.GetPools(ctx, poolsIndex[0])
	require.True(t, ok)

	p2, ok := k.GetPools(ctx, poolsIndex[1])
	require.True(t, ok)

	bapy := sdk.NewDecWithPrec(12, 2)
	coin := sdk.NewCoin("ausdc", sdk.NewIntFromUint64(100000000))

	interest1 := bapy.MulInt(coin.Amount)

	interest2 := p1.Apy.MulInt(p1.TargetAmount.Amount)
	interest3 := p2.Apy.MulInt(p2.TargetAmount.Amount)

	fmt.Printf(">>>%v>>>%v>>%v\n", p1.PoolName, p1.TargetAmount, p1.Apy)
	fmt.Printf(">>>%v>>>%v>>%v\n", p2.PoolName, p2.TargetAmount, p2.Apy)

	total := interest2.Add(interest3)
	require.True(t, total.Sub(interest1).Abs().LT(sdk.NewDecWithPrec(1, 8)))

	reqUpdate := types.MsgUpdatePool{
		Creator:           p1.OwnerAddress.String(),
		PoolIndex:         poolsIndex[0],
		PoolName:          "updatedpool-1",
		PoolApy:           "0.3",
		TargetTokenAmount: sdk.NewCoin("ausdc", sdk.NewInt(550)),
	}
	_, err = app.UpdatePool(ctx, &reqUpdate)
	require.NoError(t, err)
	p1, ok = k.GetPools(ctx, poolsIndex[0])
	require.True(t, ok)

	p2, ok = k.GetPools(ctx, poolsIndex[1])
	require.True(t, ok)

	fmt.Printf(">>%v>>>>%v>>%v\n", p1.PoolName, p1.TargetAmount, p1.Apy)
	fmt.Printf(">>>%v>>>%v>>%v\n", p2.PoolName, p2.TargetAmount, p2.Apy)

}
