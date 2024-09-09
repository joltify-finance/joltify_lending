package cli_test

import (
	"fmt"
	"testing"
	"time"

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

func TestListInvestors(t *testing.T) {
	t.SkipNow()
	config := appconfig.MakeEncodingConfig()
	k2 := keyring.NewInMemory(config.Codec)
	_, _, err := k2.NewMnemonic("0",
		keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	assert.Nil(t, err)
	v, err := k2.Key("0")
	assert.Nil(t, err)

	addr, err := v.GetAddress()
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

	_, err = net.WaitForHeightWithTimeout(1, time.Second*10)
	assert.Nil(t, err)

	for _, tc := range []struct {
		desc   string
		id     string
		fields []string
		args   []string
		err    error
		code   uint32
	}{
		{
			id:   "0",
			err:  nil,
			desc: "valid",
			code: 0,
			args: []string{
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
		},
	} {

		fields := []string{"2", "jolt1xdpg5l3pxpyhxqg4ey4krq2pf9d3sphmmuuugg,jolt1ljsg33ad5wjac6vm5htxxujrxrwgpzy8ucl2yq"}
		defaultArgs := []string{
			fmt.Sprintf("--%s=%s", flags.FlagFrom, addr.String()),
			fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
			fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		}
		var args []string
		args = append(args, fields...)
		args = append(args, defaultArgs...)
		_, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdUploadInvestor(), args)
		require.NoError(t, err)
		tc := tc

		err = net.WaitForNextBlock()
		require.NoError(t, err)

		t.Run(tc.desc, func(t *testing.T) {
			var args []string
			args = append(args, tc.fields...)
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListInvestors(), args)
			if tc.err != nil {
				require.Equal(t, tc.err.Error(), err.Error())
			} else {
				var resp kyctypes.ListInvestorsResponse
				require.NoError(t, err)
				require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))

				require.Equal(t, fields[0], resp.Investor[0].InvestorId)
				require.EqualValues(t, "jolt1xdpg5l3pxpyhxqg4ey4krq2pf9d3sphmmuuugg", resp.Investor[0].GetWalletAddress()[0])
				require.EqualValues(t, "jolt1ljsg33ad5wjac6vm5htxxujrxrwgpzy8ucl2yq", resp.Investor[0].GetWalletAddress()[1])
			}
		})
	}
}
