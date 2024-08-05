package keeper_test

import (
	"testing"

	tmlog "cosmossdk.io/log"
	tmprototypes "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtime "github.com/cometbft/cometbft/types/time"
	"github.com/joltify-finance/joltify_lending/app"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/mint/types"
	"github.com/stretchr/testify/require"
)

func TestParamsQuery(t *testing.T) {
	lg := tmlog.NewNopLogger()
	tApp := app.NewTestApp(lg, t.TempDir())
	ctx := tApp.NewContext(true, tmprototypes.Header{Height: 1, Time: tmtime.Now()})

	keeper := tApp.GetMintKeeper()

	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
