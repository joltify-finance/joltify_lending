package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	app2 "github.com/joltify-finance/joltify_lending/app"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32/legacybech32" //nolint
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/joltify-finance/joltify_lending/testutil/network"
	"github.com/joltify-finance/joltify_lending/x/vault/client/cli"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
)

func networkWithCreatePoolObjects(t *testing.T, n int, maxValidator uint32) (*network.Network, []*types.CreatePool) {
	t.Helper()
	cfg := network.DefaultConfig()
	cfg.BondedTokens = sdk.NewInt(10000000000000000)
	cfg.StakingTokens = sdk.NewInt(100000000000000000)
	state := types.GenesisState{}
	stateStaking := stakingtypes.GenesisState{}

	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[stakingtypes.ModuleName], &stateStaking))

	sk := ed25519.GenPrivKey()

	for i := 1; i < n+1; i++ {
		operator := sk.PubKey().Address().Bytes()

		sk := ed25519.GenPrivKey()
		desc := stakingtypes.NewDescription("tester", "testId", "www.test.com", "aaa", "aaa")
		testValidator, err := stakingtypes.NewValidator(operator, sk.PubKey(), desc)
		require.NoError(t, err)

		randPoolSk := ed25519.GenPrivKey()
		poolPubKey, err := legacybech32.MarshalPubKey(legacybech32.AccPK, randPoolSk.PubKey()) //nolint
		require.NoError(t, err)

		pro := types.PoolProposal{
			PoolPubKey: poolPubKey,
			PoolAddr:   sk.PubKey().Address().Bytes(),
			Nodes:      []sdk.AccAddress{operator},
		}
		state.CreatePoolList = append(state.CreatePoolList, &types.CreatePool{BlockHeight: strconv.Itoa(i), Validators: []stakingtypes.Validator{testValidator}, Proposal: []*types.PoolProposal{&pro}})
	}
	state.LatestTwoPool = state.CreatePoolList[:2]
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf

	var stateVault stakingtypes.GenesisState
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[stakingtypes.ModuleName], &stateVault))
	stateVault.Params.MaxValidators = maxValidator
	buf, err = cfg.Codec.MarshalJSON(&stateVault)
	require.NoError(t, err)
	cfg.GenesisState[stakingtypes.ModuleName] = buf
	nb := network.New(t, cfg)
	return nb, state.CreatePoolList
}

func TestShowCreatePool(t *testing.T) {
	app2.SetSDKConfig()
	net, objs := networkWithCreatePoolObjects(t, 2, 3)
	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc string
		id   string
		args []string
		err  error
		obj  *types.CreatePool
	}{
		{
			desc: "found",
			id:   objs[0].BlockHeight,
			args: common,
			obj:  objs[0],
		},
		{
			desc: "not found",
			id:   "not_found",
			args: common,
			err:  status.Error(codes.InvalidArgument, "not found"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{tc.id}
			args = append(args, tc.args...)

			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowCreatePool(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetCreatePoolResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.CreatePool)
				require.Equal(t, tc.obj.Proposal[0].PoolPubKey, resp.CreatePool.GetPoolPubKey())
			}
		})
	}
}

func TestListCreatePoolNotEnoughValidator(t *testing.T) {
	app2.SetSDKConfig()
	net, _ := networkWithCreatePoolObjects(t, 2, 300)
	net.WaitForHeight(10)
	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc string
		id   string
		args []string
		err  error
		obj  *types.CreatePool
	}{
		{
			desc: "not found because not enough validator",
			id:   "not_found",
			args: common,
			err:  status.Error(codes.InvalidArgument, "not found"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{tc.id}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowCreatePool(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetCreatePoolResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.CreatePool)
				require.Equal(t, tc.obj.Proposal[0].PoolPubKey, resp.CreatePool.GetPoolPubKey())
			}
		})
	}
}

func TestListCreatePool(t *testing.T) {
	app2.SetSDKConfig()
	net, objs := networkWithCreatePoolObjects(t, 5, 3)

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
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListCreatePool(), args)
			require.NoError(t, err)
			var resp types.QueryAllCreatePoolResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			for j := i; j < len(objs) && j < i+step; j++ {
				assert.Equal(t, objs[j].Proposal[0].PoolPubKey, resp.CreatePool[j-i].GetPoolPubKey())
			}
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListCreatePool(), args)
			require.NoError(t, err)
			var resp types.QueryAllCreatePoolResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			for j := i; j < len(objs) && j < i+step; j++ {
				assert.Equal(t, objs[j].Proposal[0].GetPoolPubKey(), resp.CreatePool[j-i].GetPoolPubKey())
			}
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListCreatePool(), args)
		require.NoError(t, err)
		var resp types.QueryAllCreatePoolResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.Len(t, resp.CreatePool, 5, "should be only 5 messages")
		require.Equal(t, objs[0].Proposal[0].PoolPubKey, resp.CreatePool[0].GetPoolPubKey())
	})
}
