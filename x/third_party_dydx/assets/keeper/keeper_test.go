package keeper_test

import (
	"testing"

	keepertest "github.com/joltify-finance/joltify_lending/dydx_helper/testutil/keeper"
	"github.com/stretchr/testify/require"
)

func TestLogger(t *testing.T) {
	ctx, keeper, _, _, _, _ := keepertest.AssetsKeepers(t, true)
	logger := keeper.Logger(ctx)
	require.NotNil(t, logger)
}
