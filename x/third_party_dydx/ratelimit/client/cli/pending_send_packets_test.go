//go:build all || integration_test

package cli_test

import (
	"fmt"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/ratelimit/client/cli"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/ratelimit/types"
	"github.com/stretchr/testify/require"
	"gotest.tools/v3/assert"
)

func TestPendingSendPackets(t *testing.T) {
	net, ctx := setupNetwork(t)

	out, err := clitestutil.ExecTestCLICmd(ctx,
		cli.CmdPendingSendPackets(),
		[]string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		})

	require.NoError(t, err)
	var resp types.QueryAllPendingSendPacketsResponse
	require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
	assert.Equal(t, 0, len(resp.PendingSendPackets))
}
