package cli_test

import (
	"fmt"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	app2 "github.com/joltify-finance/joltify_lending/app"
	kyctypes "github.com/joltify-finance/joltify_lending/x/kyc/types"
	"github.com/stretchr/testify/assert"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/joltify-finance/joltify_lending/x/kyc/client/cli"
)

func TestQueryInvestorWalletsInvestors(t *testing.T) {
	app2.SetSDKConfig()
	k2 := keyring.NewInMemory(getCodec())
	_, _, err := k2.NewMnemonic("0",
		keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	assert.Nil(t, err)
	v, err := k2.Key("0")
	assert.Nil(t, err)

	addr, err := v.GetAddress()
	_ = addr
	assert.NoError(t, err)

	net := networkPrepare(t, 3, v)
	val := net.Validators[0]
	ctx := val.ClientCtx

	key := ctx.Keyring
	assert.Nil(t, err)

	am, err := k2.ExportPrivKeyArmor("0", "testme")
	assert.Nil(t, err)

	err = key.ImportPrivKey("0", am, "testme")
	assert.Nil(t, err)

	_, err = net.WaitForHeight(1)
	assert.Nil(t, err)

	w1 := []string{"jolt1xdpg5l3pxpyhxqg4ey4krq2pf9d3sphmmuuugg", "jolt1ljsg33ad5wjac6vm5htxxujrxrwgpzy8ucl2yq"}
	w2 := []string{"jolt1hnk55w58lje99eqqjmffgg8278l22jmzpsa9rj", "jolt1t0h9w5w9wl0mhdhlfyk0rl8rnr4qudnk3yxe67", "jolt13rmg7wwpnvw4cq0fps4v4ljzvs5v9xz5k8znqa"}
	fields := []string{"2", "jolt1xdpg5l3pxpyhxqg4ey4krq2pf9d3sphmmuuugg,jolt1ljsg33ad5wjac6vm5htxxujrxrwgpzy8ucl2yq"}
	defaultArgs := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, addr.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
	}
	var args []string
	args = append(fields, defaultArgs...)
	_, err = clitestutil.ExecTestCLICmd(ctx, cli.CmdUploadInvestor(), args)
	require.NoError(t, err)
	err = net.WaitForNextBlock()
	require.NoError(t, err)

	fields2 := []string{"3", "jolt1hnk55w58lje99eqqjmffgg8278l22jmzpsa9rj,jolt1t0h9w5w9wl0mhdhlfyk0rl8rnr4qudnk3yxe67,jolt13rmg7wwpnvw4cq0fps4v4ljzvs5v9xz5k8znqa"}
	args = append(fields2, defaultArgs...)
	_, err = clitestutil.ExecTestCLICmd(ctx, cli.CmdUploadInvestor(), args)
	require.NoError(t, err)

	err = net.WaitForNextBlock()
	require.NoError(t, err)

	for _, tc := range []struct {
		desc   string
		id     string
		fields []string
		args   []string
		err    error
		code   uint32
	}{
		{
			id:     "0",
			err:    fmt.Errorf("NotFound desc = investor id 111"),
			desc:   "invalid, not found",
			code:   0,
			fields: []string{"111"},
			args: []string{
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
		},
		{
			id:     "1",
			err:    nil,
			desc:   "valid",
			code:   0,
			fields: []string{"2"},
			args: []string{
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
		},
		{
			id:     "2",
			err:    nil,
			desc:   "valid",
			code:   0,
			fields: []string{"3"},
			args: []string{
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
		},
	} {

		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			var argsQuery []string
			argsQuery = append(argsQuery, tc.fields...)
			argsQuery = append(argsQuery, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdQueryInvestorWallets(), argsQuery)
			if tc.err != nil {
				require.ErrorContains(t, err, tc.err.Error())
			} else {
				var resp kyctypes.QueryInvestorWalletsResponse
				require.NoError(t, err)
				require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
				if tc.id == "1" {
					require.EqualValues(t, w1, resp.Wallets)
				}
				if tc.id == "2" {
					require.EqualValues(t, w2, resp.Wallets)
				}
			}
		})
	}
}
