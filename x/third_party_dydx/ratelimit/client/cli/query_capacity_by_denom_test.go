//go:build all || integration_test

package cli_test

import (
	"fmt"
	"math/big"
	"strconv"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/joltify-finance/joltify_lending/dydx_helper/dtypes"
	assettypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/assets/types"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/ratelimit/client/cli"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/ratelimit/types"
	ratelimitutil "github.com/joltify-finance/joltify_lending/x/third_party_dydx/ratelimit/util"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestQueryCapacityByDenom(t *testing.T) {
	net, ctx := setupNetwork(t)

	out, err := clitestutil.ExecTestCLICmd(ctx,
		cli.CmdQueryCapacityByDenom(),
		[]string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			assettypes.AssetUsdc.Denom,
		})

	require.NoError(t, err)
	var resp types.QueryCapacityByDenomResponse
	require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
	require.Equal(t,
		// LimiterCapacity resulting from default limiter params and 0 TVL.
		[]types.LimiterCapacity{
			{
				Limiter: types.DefaultUsdcHourlyLimter,
				Capacity: dtypes.NewIntFromBigInt(
					ratelimitutil.GetBaseline(
						big.NewInt(0),
						types.DefaultUsdcHourlyLimter,
					),
				),
			},
			{
				Limiter: types.DefaultUsdcDailyLimiter,
				Capacity: dtypes.NewIntFromBigInt(
					ratelimitutil.GetBaseline(
						big.NewInt(0),
						types.DefaultUsdcDailyLimiter,
					),
				),
			},
		},
		resp.LimiterCapacityList)
}