package keeper_test

import (
	"context"
	"strings"
	"testing"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/joltify-finance/joltify_lending/x/burnauction/keeper"
	"github.com/joltify-finance/joltify_lending/x/burnauction/types"

	keepertest "github.com/joltify-finance/joltify_lending/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, types.BankKeeper, context.Context) {
	k, bk, _, ctx := keepertest.BurnauctionKeeper(t)
	return keeper.NewMsgServerImpl(*k), bk, sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	ms, _, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}

func TestSubmitRequest(t *testing.T) {
	ms, bk, ctx := setupMsgServer(t)

	invalidCoin := "abc"
	sender := sdk.AccAddress([]byte("sender"))
	_, err := ms.Submitrequest(ctx, &types.MsgSubmitrequest{
		Creator: sender.String(),
		Tokens:  invalidCoin,
	})

	require.True(t, strings.Contains(err.Error(), "invalid coins"))

	_, err = ms.Submitrequest(ctx, &types.MsgSubmitrequest{
		Creator: "invalid_address",
		Tokens:  "100stake",
	})
	require.True(t, strings.Contains(err.Error(), "invalid address"))

	testCoin := sdk.NewCoins(sdk.NewCoin("uatom", sdk.NewInt(100)), sdk.NewCoin("ustake", sdk.NewInt(120)))

	_, err = ms.Submitrequest(ctx, &types.MsgSubmitrequest{
		Creator: sender.String(),
		Tokens:  testCoin.String(),
	})
	require.NoError(t, err)

	addr := authtypes.NewModuleAddress(types.ModuleName)

	balances := bk.GetAllBalances(sdk.UnwrapSDKContext(ctx), addr)
	require.Equal(t, balances.IsEqual(testCoin), true)
}
