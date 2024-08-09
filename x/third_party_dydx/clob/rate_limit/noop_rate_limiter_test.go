package rate_limit_test

import (
	"testing"

	testapp "github.com/joltify-finance/joltify_lending/dydx_helper/testutil/app"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/rate_limit"
	"github.com/stretchr/testify/require"
)

func TestNoOpRateLimiter(t *testing.T) {
	tApp := testapp.NewTestAppBuilder(t).Build()
	ctx := tApp.InitChain()
	rl := rate_limit.NewNoOpRateLimiter[int]()
	for i := 0; i < 100_000; i += 1 {
		require.NoError(t, rl.RateLimit(ctx, 42))
	}
	rl.PruneRateLimits(ctx)
}
