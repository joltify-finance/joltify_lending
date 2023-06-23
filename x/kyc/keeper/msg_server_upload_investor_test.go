package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/utils"
	"github.com/joltify-finance/joltify_lending/x/kyc/types"
	"github.com/stretchr/testify/require"
)

func generateNAddr(n int) []string {
	addresses := make([]string, n)
	for i := 0; i < n; i++ {
		pk := ed25519.GenPrivKey().PubKey()
		addr := pk.Address().Bytes()
		a := sdk.AccAddress(addr)
		addresses[i] = a.String()
	}
	return addresses
}

func TestSubmitInvestor(t *testing.T) {
	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)

	acc, err := sdk.AccAddressFromBech32("jolt1p3jl6udk43vw0cvc5hjqrpnncsqmsz56wd32z8")
	require.NoError(t, err)
	lapp, k, wctx := setupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(wctx)
	pa := types.Params{Submitter: []sdk.AccAddress{acc}}
	k.SetParams(ctx, pa)

	msg := types.MsgUploadInvestor{Creator: "invalid msg"}
	_, err = lapp.UploadInvestor(wctx, &msg)
	require.Error(t, err)

	// unauthorised submitter
	msg.Creator = "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0"
	_, err = lapp.UploadInvestor(wctx, &msg)
	require.Error(t, err)

	msg.Creator = "jolt1p3jl6udk43vw0cvc5hjqrpnncsqmsz56wd32z8"
	msg.InvestorId = "123"
	ret, err := lapp.UploadInvestor(wctx, &msg)
	require.NoError(t, err)
	require.Len(t, ret.Wallets, 0)

	addresses := generateNAddr(10)
	// invalid investor ID test
	msg.InvestorId = ""
	msg.WalletAddress = addresses
	_, err = lapp.UploadInvestor(wctx, &msg)
	require.Error(t, err)

	msg.InvestorId = "first Investor"
	ret, err = lapp.UploadInvestor(wctx, &msg)
	require.NoError(t, err)
	require.EqualValues(t, ret.Wallets, addresses)
	out, err := k.QueryInvestorWallets(wctx, &types.QueryInvestorWalletsRequest{InvestorId: msg.InvestorId})
	require.NoError(t, err)
	require.EqualValues(t, out.Wallets, msg.WalletAddress)

	msg.Creator = "jolt1p3jl6udk43vw0cvc5hjqrpnncsqmsz56wd32z8"
	msg.InvestorId = "333"
	msg.WalletAddress = msg.WalletAddress[2:3]
	_, err = lapp.UploadInvestor(wctx, &msg)
	require.Error(t, err)

	// exceed the max allowed wallets
	addresses = generateNAddr(1000)
	msg.WalletAddress = addresses
	_, err = lapp.UploadInvestor(wctx, &msg)
	require.Error(t, err)
}

func TestUpdateWallets(t *testing.T) {
	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)

	acc, err := sdk.AccAddressFromBech32("jolt1p3jl6udk43vw0cvc5hjqrpnncsqmsz56wd32z8")
	require.NoError(t, err)
	lapp, k, wctx := setupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(wctx)
	pa := types.Params{Submitter: []sdk.AccAddress{acc}}
	k.SetParams(ctx, pa)

	addresses := generateNAddr(150)
	original := make([]string, 150)
	copy(original, addresses)
	msg := types.MsgUploadInvestor{}
	msg.Creator = "jolt1p3jl6udk43vw0cvc5hjqrpnncsqmsz56wd32z8"
	msg.InvestorId = "123"
	msg.WalletAddress = addresses[:30]
	ret, err := lapp.UploadInvestor(wctx, &msg)
	require.NoError(t, err)
	require.EqualValues(t, ret.Wallets, msg.WalletAddress)

	out, err := k.QueryInvestorWallets(wctx, &types.QueryInvestorWalletsRequest{InvestorId: msg.InvestorId})
	require.NoError(t, err)
	require.EqualValues(t, out.Wallets, msg.WalletAddress)

	// we add more wallets
	msg.WalletAddress = addresses[30:100]
	ret, err = lapp.UploadInvestor(wctx, &msg)
	require.NoError(t, err)
	require.EqualValues(t, addresses[:100], ret.Wallets)

	out, err = k.QueryInvestorWallets(wctx, &types.QueryInvestorWalletsRequest{InvestorId: msg.InvestorId})
	require.NoError(t, err)
	require.EqualValues(t, out.Wallets, addresses[:100])

	// we add 1 more wallets that will pop out the first one
	msg.WalletAddress = []string{addresses[100]}
	ret, err = lapp.UploadInvestor(wctx, &msg)
	require.NoError(t, err)
	require.EqualValues(t, original[1:101], ret.Wallets)

	out, err = k.QueryInvestorWallets(wctx, &types.QueryInvestorWalletsRequest{InvestorId: msg.InvestorId})
	require.NoError(t, err)
	require.EqualValues(t, out.Wallets, addresses[1:101])

	// we add the extra 50 wallet
	msg.WalletAddress = addresses[101:]
	ret, err = lapp.UploadInvestor(wctx, &msg)
	require.NoError(t, err)
	require.EqualValues(t, original[50:150], ret.Wallets)

	out, err = k.QueryInvestorWallets(wctx, &types.QueryInvestorWalletsRequest{InvestorId: msg.InvestorId})
	require.NoError(t, err)
	require.EqualValues(t, out.Wallets, original[50:150])
}
