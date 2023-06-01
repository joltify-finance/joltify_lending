package cli_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/joltify-finance/joltify_lending/x/kyc/types"

	app2 "github.com/joltify-finance/joltify_lending/app"

	"github.com/stretchr/testify/assert"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/joltify-finance/joltify_lending/testutil/network"
	"github.com/joltify-finance/joltify_lending/x/kyc/client/cli"
)

func networkPrepare(t *testing.T, maxValidator uint32, v *keyring.Record) *network.Network {
	t.Helper()
	cfg := network.DefaultConfig()
	cfg.MinGasPrices = "0stake"
	cfg.BondedTokens = sdk.NewInt(10000000000000000)
	cfg.StakingTokens = sdk.NewInt(100000000000000000)
	state := types.GenesisState{}
	stateStaking := stakingtypes.GenesisState{}
	stateBank := banktypes.GenesisState{}
	stateAuth := authtypes.GenesisState{}

	addr, err := v.GetAddress()
	if err != nil {
		panic(err)
	}
	pk, err := v.GetPubKey()
	if err != nil {
		panic(err)
	}

	acc := authtypes.NewBaseAccount(addr, pk, 10, 0)
	//balanceItem := banktypes.Balance{
	//	Address: acc.GetAddress().String(),
	//	Coins:   sdk.NewCoins(sdk.NewCoin("stake", cfg.BondedTokens)),
	//}
	genAccs := []authtypes.GenesisAccount{acc}
	// balances := []banktypes.Balance{balanceItem}

	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[stakingtypes.ModuleName], &stateStaking))
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[banktypes.ModuleName], &stateBank))
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[authtypes.ModuleName], &stateAuth))

	state.Params.Submitter = []sdk.AccAddress{addr}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)

	// stateAuth.

	authGenesis := authtypes.NewGenesisState(authtypes.DefaultParams(), genAccs)

	stateBank.Balances = []banktypes.Balance{{Address: addr.String(), Coins: sdk.Coins{sdk.NewCoin("stake", sdk.NewInt(100000))}}}
	bankBuf, err := cfg.Codec.MarshalJSON(&stateBank)
	require.NoError(t, err)

	cfg.GenesisState[banktypes.ModuleName] = bankBuf
	cfg.GenesisState[types.ModuleName] = buf
	cfg.GenesisState[authtypes.ModuleName] = cfg.Codec.MustMarshalJSON(authGenesis)

	var stateVault stakingtypes.GenesisState
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[stakingtypes.ModuleName], &stateVault))
	stateVault.Params.MaxValidators = maxValidator
	buf, err = cfg.Codec.MarshalJSON(&stateVault)
	require.NoError(t, err)
	cfg.GenesisState[stakingtypes.ModuleName] = buf
	nb := network.New(t, cfg)
	return nb
}

func getCodec() codec.Codec {
	registry := codectypes.NewInterfaceRegistry()
	cryptocodec.RegisterInterfaces(registry)
	return codec.NewProtoCodec(registry)
}

func TestUploadInvestor(t *testing.T) {
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
			err:    nil,
			desc:   "invalid",
			code:   1,
			fields: []string{"1", "jolt1xdpg5l3pxpyhxqg4ey4krq2pf9d3sphmmuuugg"},
			args: []string{
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
			},
		},

		{
			id:     "1",
			desc:   "valid",
			fields: []string{"1", "jolt1xdpg5l3pxpyhxqg4ey4krq2pf9d3sphmmuuugg"},
			args: []string{
				fmt.Sprintf("--%s=%s", flags.FlagFrom, addr.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
			},
		},

		{
			id:     "2",
			desc:   "invalid address",
			fields: []string{"1", "jolt1xdpg5l3pxpyhxqg4ey4krq2pf9d3sphmmuuugg,abc"},
			code:   1,
			args: []string{
				fmt.Sprintf("--%s=%s", flags.FlagFrom, addr.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
			},
		},

		{
			id:     "3",
			desc:   "multipu address",
			fields: []string{"2", "jolt15wtdzw37x4g0fcehvp8twekhdanwrxapnn8ntt,jolt15wtdzw37x4g0fcehvp8twekhdanwrxapnn8ntt"},
			code:   0,
			args: []string{
				fmt.Sprintf("--%s=%s", flags.FlagFrom, addr.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
			},
		},

		{
			id:     "4",
			desc:   "cannot map multiple wallet addresses to different investors",
			fields: []string{"4", "jolt1xdpg5l3pxpyhxqg4ey4krq2pf9d3sphmmuuugg"},
			code:   1,
			args: []string{
				fmt.Sprintf("--%s=%s", flags.FlagFrom, addr.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
			},
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			var args []string
			if tc.id == "4" {
				argsBefore := []string{"1", "jolt1xdpg5l3pxpyhxqg4ey4krq2pf9d3sphmmuuugg"}
				argsBefore = append(argsBefore, tc.args...)
				_, err = clitestutil.ExecTestCLICmd(ctx, cli.CmdUploadInvestor(), argsBefore)
				require.NoError(t, err)
			}
			args = append(tc.fields, tc.args...)
			out, errOut := clitestutil.ExecTestCLICmd(ctx, cli.CmdUploadInvestor(), args)
			if tc.err != nil {
				require.Equal(t, tc.err.Error(), errOut.Error())
			} else {
				var resp sdk.TxResponse
				require.NoError(t, err)
				require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.Equal(t, tc.code, resp.Code)
			}
		})
	}
}
