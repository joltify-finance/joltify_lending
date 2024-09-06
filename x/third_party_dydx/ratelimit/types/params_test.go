package types_test

import (
	"testing"
	"time"

	"github.com/joltify-finance/joltify_lending/dydx_helper/dtypes"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/ratelimit/types"
	"github.com/stretchr/testify/require"
)

func TestDefaultUsdcRateLimitParams(t *testing.T) {
	require.Equal(t,
		types.LimitParams{
			Denom: "ibc/65D0BEC6DAD96C7F5043D1E54E54B6BB5D5B3AEC3FF6CEBB75B9E059F3580EA3",
			Limiters: []types.Limiter{
				{
					Period:          3600 * time.Second,
					BaselineMinimum: dtypes.NewInt(1_000_000_000_000),
					BaselineTvlPpm:  10_000,
				},
				{
					Period:          24 * time.Hour,
					BaselineMinimum: dtypes.NewInt(10_000_000_000_000),
					BaselineTvlPpm:  100_000,
				},
			},
		},
		types.DefaultUsdcRateLimitParams(),
	)
}
