package keeper_test

import (
	"testing"

	keepertest "github.com/joltify-finance/joltify_lending/dydx_helper/testutil/keeper"
	"github.com/stretchr/testify/require"
)

func TestLogger(t *testing.T) {
	pc := keepertest.PerpetualsKeepers(t)
	logger := pc.PerpetualsKeeper.Logger(pc.Ctx)
	require.NotNil(t, logger)
}
