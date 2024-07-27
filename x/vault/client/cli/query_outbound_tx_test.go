package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/joltify-finance/joltify_lending/testutil/network"
	"github.com/joltify-finance/joltify_lending/testutil/nullify"
	"github.com/joltify-finance/joltify_lending/x/vault/client/cli"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func SetupBech32Prefix() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("jolt", "joltpub")
	config.SetBech32PrefixForValidator("joltvaloper", "joltvpub")
	config.SetBech32PrefixForConsensusNode("joltvalcons", "joltcpub")
}

func networkWithOutboundTxObjects(t *testing.T, n int) (*network.Network, []types.OutboundTx) {
	t.Helper()
	SetupBech32Prefix()

	cfg := network.DefaultConfig()
	cfg.BondedTokens = sdkmath.NewInt(10000000000000000)
	cfg.StakingTokens = sdkmath.NewInt(100000000000000000)
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		outboundTx := types.OutboundTx{
			Index: strconv.Itoa(i),
		}
		nullify.Fill(&outboundTx)
		state.OutboundTxList = append(state.OutboundTxList, outboundTx)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.OutboundTxList
}

func TestShowOutboundTx(t *testing.T) {
	net, objs := networkWithOutboundTxObjects(t, 2)
	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc        string
		idRequestID string

		args []string
		err  error
		obj  types.OutboundTx
	}{
		{
			desc:        "found",
			idRequestID: objs[0].Index,

			args: common,
			obj:  objs[0],
		},
		{
			desc:        "not found",
			idRequestID: strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.InvalidArgument, "not found"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idRequestID,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowOutboundTx(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetOutboundTxResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.OutboundTx)
				require.Equal(t,
					tc.obj.String(),
					resp.OutboundTx.String())
			}
		})
	}
}

func TestListOutboundTx(t *testing.T) {
	net, objs := networkWithOutboundTxObjects(t, 5)

	ctx := net.Validators[0].ClientCtx
	request := func(next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
		if next == nil {
			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
		} else {
			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
		}
		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
		if total {
			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
		}
		return args
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(objs); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListOutboundTx(), args)
			require.NoError(t, err)
			var resp types.QueryAllOutboundTxResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.AllOutbound), step)

			var a []string
			for _, el := range objs {
				a = append(a, el.String())
			}

			var b []string
			for _, el := range resp.AllOutbound {
				b = append(b, el.OutboundTx.String())
			}
			require.Subset(t, a, b)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListOutboundTx(), args)
			require.NoError(t, err)
			var resp types.QueryAllOutboundTxResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.AllOutbound), step)

			var a []string
			for _, el := range objs {
				a = append(a, el.String())
			}

			var b []string
			for _, el := range resp.GetAllOutbound() {
				b = append(b, el.OutboundTx.String())
			}

			require.Subset(t, a, b)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListOutboundTx(), args)
		require.NoError(t, err)
		var resp types.QueryAllOutboundTxResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))

		var a []string
		for _, el := range objs {
			a = append(a, el.String())
		}

		var b []string
		for _, el := range resp.AllOutbound {
			b = append(b, el.OutboundTx.String())
		}

		require.ElementsMatch(t, a, b)
	})
}
