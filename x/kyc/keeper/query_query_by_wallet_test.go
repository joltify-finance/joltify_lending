package keeper_test

import (
	"strconv"
	"testing"

	types2 "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/utils"
	"github.com/joltify-finance/joltify_lending/x/kyc/types"
	"github.com/stretchr/testify/require"
)

func TestQueryByWallet(t *testing.T) {
	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)
	acc, err := types2.AccAddressFromBech32("jolt1p3jl6udk43vw0cvc5hjqrpnncsqmsz56wd32z8")
	require.NoError(t, err)
	lapp, k, wctx := setupMsgServer(t)
	ctx := types2.UnwrapSDKContext(wctx)
	pa := types.Params{Submitter: []types2.AccAddress{acc}}
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
		ret, err := lapp.UploadInvestor(wctx, &msg)
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
		retInvestor, err := k.QueryByWallet(wctx, &types.QueryByWalletRequest{Wallet: original[i]})
		require.NoError(t, err)
		require.Equal(t, retInvestor.Investor.InvestorId, strconv.Itoa(i))
	}
}

func TestQueryByWalletTwoInvestorIDWithSameWallet(t *testing.T) {
	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)
	acc, err := types2.AccAddressFromBech32("jolt1p3jl6udk43vw0cvc5hjqrpnncsqmsz56wd32z8")
	require.NoError(t, err)
	lapp, k, wctx := setupMsgServer(t)
	ctx := types2.UnwrapSDKContext(wctx)
	pa := types.Params{Submitter: []types2.AccAddress{acc}}
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
		ret, err := lapp.UploadInvestor(wctx, &msg)
		require.NoError(t, err)
		require.EqualValues(t, ret.Wallets, msg.WalletAddress)
		inv := types.Investor{
			InvestorId:    msg.InvestorId,
			WalletAddress: msg.WalletAddress,
		}
		addressMap[msg.InvestorId] = msg.WalletAddress[0]
		investors = append(investors, &inv)
	}

	msg := types.MsgUploadInvestor{}
	msg.Creator = "jolt1p3jl6udk43vw0cvc5hjqrpnncsqmsz56wd32z8"
	msg.InvestorId = strconv.Itoa(501)
	msg.WalletAddress = []string{addresses[99]}
	_, err = lapp.UploadInvestor(wctx, &msg)
	require.Error(t, err)
}
