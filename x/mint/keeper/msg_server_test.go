package keeper_test

import (
	"context"
	"testing"

	app2 "github.com/joltify-finance/joltify_lending/app"
	tmprototypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/mint/keeper"
	"github.com/joltify-finance/joltify_lending/x/mint/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	app := app2.NewTestApp()
	ctx := app.NewContext(true, tmprototypes.Header{Height: 1, Time: tmtime.Now()})
	k := app.GetMintKeeper()
	return keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)
}
