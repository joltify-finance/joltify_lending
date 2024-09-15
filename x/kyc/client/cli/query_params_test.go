package cli_test

import (
	"fmt"
	"testing"

	appconfig "github.com/joltify-finance/joltify_lending/app/config"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/kyc/client/cli"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	kyctypes "github.com/joltify-finance/joltify_lending/x/kyc/types"
)

func TestQueryParameters(t *testing.T) {
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

	err = net.WaitForNextBlock()
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
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			var args []string
			args = append(args, tc.fields...)
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdQueryParams(), args)
			if tc.err != nil {
				require.Equal(t, tc.err.Error(), err.Error())
			} else {
				var resp kyctypes.QueryParamsResponse
				require.NoError(t, err)
				require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
			}
		})
	}
}
