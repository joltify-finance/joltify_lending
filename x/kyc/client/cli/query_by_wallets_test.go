package cli_test

import (
	"fmt"
	"testing"

	appconfig "github.com/joltify-finance/joltify_lending/app/config"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	kyctypes "github.com/joltify-finance/joltify_lending/x/kyc/types"
	"github.com/stretchr/testify/assert"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/joltify-finance/joltify_lending/x/kyc/client/cli"
)

func TestQueryByWallets(t *testing.T) {
	appconfig.SetupConfig()
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

	w1 := []string{"jolt1xdpg5l3pxpyhxqg4ey4krq2pf9d3sphmmuuugg", "jolt1ljsg33ad5wjac6vm5htxxujrxrwgpzy8ucl2yq", "jolt1hnk55w58lje99eqqjmffgg8278l22jmzpsa9rj"}
	w2 := []string{"jolt1hnk55w58lje99eqqjmffgg8278l22jmzpsa9rj", "jolt1t0h9w5w9wl0mhdhlfyk0rl8rnr4qudnk3yxe67", "jolt13rmg7wwpnvw4cq0fps4v4ljzvs5v9xz5k8znqa"}
	defaultArgs := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, addr.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
	}
	fields := []string{"2", "jolt1xdpg5l3pxpyhxqg4ey4krq2pf9d3sphmmuuugg,jolt1ljsg33ad5wjac6vm5htxxujrxrwgpzy8ucl2yq,jolt1hnk55w58lje99eqqjmffgg8278l22jmzpsa9rj"}

	fields2 := []string{"3", "jolt15wtdzw37x4g0fcehvp8twekhdanwrxapnn8ntt,jolt1t0h9w5w9wl0mhdhlfyk0rl8rnr4qudnk3yxe67,jolt13rmg7wwpnvw4cq0fps4v4ljzvs5v9xz5k8znqa"}
	args := append(fields, defaultArgs...)
	_, err = clitestutil.ExecTestCLICmd(ctx, cli.CmdUploadInvestor(), args)
	require.NoError(t, err)
	err = net.WaitForNextBlock()
	require.NoError(t, err)

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
			err:    fmt.Errorf("rpc error: code = InvalidArgument desc = rpc error: code = InvalidArgument desc = invalid request wallet address: invalid request"),
			desc:   "invalid, not found",
			code:   1,
			fields: []string{"111"},
			args: []string{
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
		},
		{
			id:     "1",
			err:    fmt.Errorf("rpc error: code = InvalidArgument desc = rpc error: code = InvalidArgument desc = invalid request wallet address: invalid request"),
			desc:   "invalid as multify address",
			code:   1,
			fields: []string{"jolt1xdpg5l3pxpyhxqg4ey4krq2pf9d3sphmmuuugg,jolt1ljsg33ad5wjac6vm5htxxujrxrwgpzy8ucl2yq"},
			args: []string{
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
		},
		{
			id:     "2",
			err:    nil,
			desc:   "valid",
			code:   0,
			fields: []string{w1[0]},
			args: []string{
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
		},

		{
			id:     "3",
			err:    nil,
			desc:   "valid",
			code:   0,
			fields: []string{w1[1]},
			args: []string{
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
		},

		{
			id:     "4",
			err:    nil,
			desc:   "valid",
			code:   0,
			fields: []string{w2[1]},
			args: []string{
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
		},

		{
			id:     "5",
			err:    nil,
			desc:   "valid",
			code:   0,
			fields: []string{w2[2]},
			args: []string{
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			var args []string
			args = append(tc.fields, tc.args...)
			out, err2 := clitestutil.ExecTestCLICmd(ctx, cli.CmdQueryByWallet(), args)
			if tc.err != nil {
				assert.EqualError(t, err2, tc.err.Error())
			}
			if tc.err != nil {
				require.Equal(t, tc.err.Error(), err2.Error())
			} else {
				err = net.WaitForNextBlock()
				require.NoError(t, err)
				var resp kyctypes.QueryByWalletResponse
				require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
				if tc.id == "2" {
					require.EqualValues(t, "2", resp.Investor.InvestorId)
				}

				if tc.id == "3" {
					require.EqualValues(t, "2", resp.Investor.InvestorId)
				}

				if tc.id == "4" {
					require.EqualValues(t, "3", resp.Investor.InvestorId)
				}

				if tc.id == "5" {
					require.EqualValues(t, "3", resp.Investor.InvestorId)
				}

			}
		})
	}
}
