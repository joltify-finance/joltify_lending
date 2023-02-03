package keeper_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"

	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/utils"
	"github.com/joltify-finance/joltify_lending/x/kyc/types"
	"github.com/stretchr/testify/require"
)

func TestQueryAllInvestors(t *testing.T) {
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
		msg.InvestorId = fmt.Sprintf("%v", i)
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
	retAllInvestor, err := k.ListInvestors(wctx, &types.ListInvestorsRequest{})
	require.NoError(t, err)

	for i := 0; i < 100; i++ {
		inv := retAllInvestor.Investor[i]
		addr := addressMap[inv.InvestorId]
		require.Equal(t, addr, inv.WalletAddress[0])
	}
}
