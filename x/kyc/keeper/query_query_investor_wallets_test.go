package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/utils"
	"github.com/joltify-finance/joltify_lending/x/kyc/types"
	"github.com/stretchr/testify/require"
)

func TestQueryByInvestor(t *testing.T) {
	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)

	acc, err := sdk.AccAddressFromBech32("jolt1p3jl6udk43vw0cvc5hjqrpnncsqmsz56wd32z8")
	require.NoError(t, err)
	app, k, wctx := setupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(wctx)
	pa := types.Params{Submitter: []sdk.AccAddress{acc}}
	k.SetParams(ctx, pa)

	addresses := generateNAddr(100)
	original := make([]string, 100)
	copy(original, addresses)
	addressMap := make(map[string]string)
	investors := make([]*types.Investor, 100)
	for i := 0; i < 100; i++ {
		msg := types.MsgUploadInvestor{}
		msg.Creator = "jolt1p3jl6udk43vw0cvc5hjqrpnncsqmsz56wd32z8"
		msg.InvestorId = strconv.Itoa(i)
		msg.WalletAddress = []string{addresses[i]}
		ret, err := app.UploadInvestor(wctx, &msg)
		require.NoError(t, err)
		require.EqualValues(t, ret.Wallets, msg.WalletAddress)
		inv := types.Investor{
			InvestorId:    msg.InvestorId,
			WalletAddress: msg.WalletAddress,
		}
		addressMap[msg.InvestorId] = msg.WalletAddress[0]
		investors = append(investors, &inv)
	}

	for i := 0; i < 100; i++ {
		id := strconv.Itoa(i)
		retInvestor, err := k.QueryInvestorWallets(wctx, &types.QueryInvestorWalletsRequest{InvestorId: id})
		require.NoError(t, err)
		actualAddr := retInvestor.Wallets[0]
		addr := addressMap[id]
		require.Equal(t, addr, actualAddr)
	}

	addresses2 := generateNAddr(100)
	msg := types.MsgUploadInvestor{}
	msg.Creator = "jolt1p3jl6udk43vw0cvc5hjqrpnncsqmsz56wd32z8"
	msg.InvestorId = strconv.Itoa(105)
	msg.WalletAddress = addresses2
	_, err = app.UploadInvestor(wctx, &msg)
	require.NoError(t, err)

	retInvestor, err := k.QueryInvestorWallets(wctx, &types.QueryInvestorWalletsRequest{InvestorId: msg.InvestorId})
	require.NoError(t, err)
	require.Equal(t, addresses2, retInvestor.Wallets)
}
