package keeper_test

import (
	"testing"

	keepertest "github.com/joltify-finance/joltify_lending/testutil/dydx/testutil/keeper"
	"github.com/stretchr/testify/require"
)

func TestLogger(t *testing.T) {
	ctx, keeper, _, _, _ := keepertest.PricesKeepers(t)
	logger := keeper.Logger(ctx)
	require.NotNil(t, logger)
}
