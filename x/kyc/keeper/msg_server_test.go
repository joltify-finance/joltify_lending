package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/joltify-finance/joltify_lending/testutil/keeper"
	"github.com/joltify-finance/joltify_lending/x/kyc/keeper"
	"github.com/joltify-finance/joltify_lending/x/kyc/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, *keeper.Keeper, context.Context) {
	k, ctx := keepertest.KycKeeper(t)
	return keeper.NewMsgServerImpl(*k), k, sdk.WrapSDKContext(ctx)
}
