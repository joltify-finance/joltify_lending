//go:build all || integration_test

package cli_test

import (
	"fmt"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/client/cli"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/types"
	"github.com/stretchr/testify/require"
)

func TestCmdGetLiquidationsConfiguration(t *testing.T) {
	net, _ := networkWithClobPairObjects(t, 2)
	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}

	out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdGetLiquidationsConfiguration(), common)
	require.NoError(t, err)
	var resp types.QueryLiquidationsConfigurationResponse
	require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
	require.NotNil(t, resp.LiquidationsConfig)
	require.Equal(
		t,
		types.LiquidationsConfig_Default,
		resp.LiquidationsConfig,
	)
}
