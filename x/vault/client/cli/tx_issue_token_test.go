package cli_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"

	app2 "github.com/joltify-finance/joltify_lending/app"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32/legacybech32" //nolint
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/joltify-finance/joltify_lending/testutil/network"
	"github.com/joltify-finance/joltify_lending/x/vault/client/cli"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/joltify-finance/joltify_lending/x/vault/types"
)

func preparePool(t *testing.T) (*network.Network, []*types.CreatePool) {
	t.Helper()
	height := []int{7, 10}
	cfg := network.DefaultConfig()
	cfg.MinGasPrices = "0poppy"
	cfg.BondedTokens = sdk.NewInt(10000000000000000)
	cfg.StakingTokens = sdk.NewInt(100000000000000000)
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))
	sk := ed25519.GenPrivKey()
	poolPubKey, err := legacybech32.MarshalPubKey(legacybech32.AccPK, sk.PubKey()) //nolint
	require.NoError(t, err)
	operator := sk.PubKey().Address().Bytes()
	require.NoError(t, err)
	validators := make([]*stakingtypes.Validator, len(height))
	for i, el := range height {
		sk := ed25519.GenPrivKey()
		desc := stakingtypes.NewDescription("tester", "testId", "www.test.com", "aaa", "aaa")
		testValidator, err := stakingtypes.NewValidator(operator, sk.PubKey(), desc)
		require.NoError(t, err)
		validators = append(validators, &testValidator)
		validators[i] = &testValidator
		pro := types.PoolProposal{
			PoolPubKey: poolPubKey,
			Nodes:      []sdk.AccAddress{operator},
			PoolAddr:   sk.PubKey().Address().Bytes(),
		}
		state.CreatePoolList = append(state.CreatePoolList, &types.CreatePool{BlockHeight: strconv.Itoa(el), Validators: []stakingtypes.Validator{testValidator}, Proposal: []*types.PoolProposal{&pro}})
	}
	state.Params.BlockChurnInterval = 3
	state.LatestTwoPool = state.CreatePoolList[:2]

	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf

	var stateVault stakingtypes.GenesisState
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[stakingtypes.ModuleName], &stateVault))
	stateVault.Params.MaxValidators = 3
	buf, err = cfg.Codec.MarshalJSON(&stateVault)
	require.NoError(t, err)
	cfg.GenesisState[stakingtypes.ModuleName] = buf

	net := network.New(t, cfg)
	net.Config.BondDenom = "poppy"
	return net, state.CreatePoolList
}

// this test will fail as it is not from pool owner
func TestCreateIssueTokenFail(t *testing.T) {
	app2.SetSDKConfig()

	net, _ := preparePool(t)
	val := net.Validators[0]
	ctx := val.ClientCtx
	id := "0"

	_, err := net.WaitForHeightWithTimeout(10, time.Minute)
	assert.Nil(t, err)

	fields := []string{"100vvusd", "jolt18mdnq8x9m07dryymlyf8jknagp87yga0hpe7n6"}
	for _, tc := range []struct {
		desc string
		id   string
		args []string
		err  error
		code uint32
	}{
		{
			id:   id,
			desc: "valid issue token",
			args: []string{
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				// fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdk.NewInt(10))).String()),
			},
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{tc.id}

			args = append(args, fields...)
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateIssueToken(), args)
			var resp sdk.TxResponse
			require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
			err = net.WaitForNextBlock()
			require.NoError(t, err)
			ret, err := tx.QueryTx(ctx, resp.TxHash)
			require.NoError(t, err)
			expected := fmt.Sprintf("failed to execute message; message index: 0: creator %v is not in pool addresses set: invalid request", val.Address.String())
			require.Equal(t, expected, ret.RawLog)
			require.Nil(t, err)
			require.NotEqual(t, uint32(0), ret.Code)
		})
	}
}

func networkPrepare(t *testing.T, maxValidator uint32, addr sdk.AccAddress, pk cryptotypes.PubKey) (*network.Network, []*types.CreatePool) {
	t.Helper()
	cfg := network.DefaultConfig()
	cfg.MinGasPrices = "0stake"
	cfg.BondedTokens = sdk.NewInt(10000000000000000)
	cfg.StakingTokens = sdk.NewInt(100000000000000000)
	state := types.GenesisState{}
	stateStaking := stakingtypes.GenesisState{}
	stateBank := banktypes.GenesisState{}

	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[stakingtypes.ModuleName], &stateStaking))
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[banktypes.ModuleName], &stateBank))

	state.Params.BlockChurnInterval = 3
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	stateBank.Balances = []banktypes.Balance{{Address: addr.String(), Coins: sdk.Coins{sdk.NewCoin("stake", sdk.NewInt(100000))}}}
	bankBuf, err := cfg.Codec.MarshalJSON(&stateBank)
	require.NoError(t, err)
	cfg.GenesisState[banktypes.ModuleName] = bankBuf

	var stateVault stakingtypes.GenesisState
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[stakingtypes.ModuleName], &stateVault))
	stateVault.Params.MaxValidators = maxValidator
	buf, err = cfg.Codec.MarshalJSON(&stateVault)
	require.NoError(t, err)
	cfg.GenesisState[stakingtypes.ModuleName] = buf

	// acc := authtypes.NewBaseAccount(addr, pk, 10, 0)
	// genAccs := []authtypes.GenesisAccount{acc}

	// authGenesis := authtypes.NewGenesisState(authtypes.DefaultParams(), genAccs)
	// cfg.GenesisState[authtypes.ModuleName] = cfg.Codec.MustMarshalJSON(authGenesis)

	nb := network.New(t, cfg)
	return nb, state.CreatePoolList
}

func getCodec() codec.Codec {
	registry := codectypes.NewInterfaceRegistry()
	cryptocodec.RegisterInterfaces(registry)
	return codec.NewProtoCodec(registry)
}

// this test will fail as it is not from pool owner
func TestCreateIssue(t *testing.T) {
	app2.SetSDKConfig()
	k2 := keyring.NewInMemory(getCodec())
	_, _, err := k2.NewMnemonic("0",
		keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	assert.Nil(t, err)
	v, err := k2.Key("0")
	assert.Nil(t, err)
	addr, err := v.GetAddress()
	assert.NoError(t, err)
	pk, err := v.GetPubKey()
	assert.NoError(t, err)

	net, _ := networkPrepare(t, 3, addr, pk)

	err = net.WaitForNextBlock()
	assert.NoError(t, err)
	val := net.Validators[0]
	ctx := val.ClientCtx
	key := ctx.Keyring
	info, err := key.List()
	assert.Nil(t, err)

	am, err := k2.ExportPrivKeyArmor("0", "testme")
	assert.Nil(t, err)

	err = key.ImportPrivKey("0", am, "testme")
	assert.Nil(t, err)

	thisInfo, err := key.Key("0")
	assert.Nil(t, err)

	thisInfoPk, err := thisInfo.GetPubKey()
	assert.Nil(t, err)

	pubkey := legacybech32.MustMarshalPubKey(legacybech32.AccPK, thisInfoPk) //nolint
	createPoolFields := []string{pubkey, "4"}

	infoAddr, err := info[0].GetAddress()
	assert.Nil(t, err)

	thisInfoAddr, err := thisInfo.GetAddress()
	assert.Nil(t, err)

	commonArgs := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, infoAddr),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdk.NewInt(10))).String()),
	}

	commonArgs2 := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, thisInfoAddr),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
	}
	var args []string
	args = append(args, createPoolFields...)
	args = append(args, commonArgs...)

	err = net.WaitForNextBlock()
	//_, err = net.WaitForHeightWithTimeout(30, time.Minute)
	assert.Nil(t, err)

	out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateCreatePool(), args)
	assert.Nil(t, err)
	var resp sdk.TxResponse
	require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
	require.Equal(t, uint32(0), resp.Code)

	current, err := net.LatestHeight()
	assert.Nil(t, err)
	_, err = net.WaitForHeightWithTimeout(int64(current)+5, time.Minute)
	assert.Nil(t, err)

	ret, err := tx.QueryTx(ctx, resp.TxHash)
	require.NoError(t, err)

	//_, err = net.WaitForHeightWithTimeout(15, time.Minute)
	assert.Nil(t, err)
	// now we submit the issue token request
	issueTokenfields := []string{"100vvusd", "jolt18mdnq8x9m07dryymlyf8jknagp87yga0hpe7n6"}
	id := "0"
	issueTokenArgs := []string{id}
	issueTokenArgs = append(issueTokenArgs, issueTokenfields...)
	issueTokenArgs = append(issueTokenArgs, commonArgs2...)
	out, err = clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateIssueToken(), issueTokenArgs)
	assert.Nil(t, err)

	current, err = net.LatestHeight()
	assert.Nil(t, err)
	_, err = net.WaitForHeightWithTimeout(int64(current)+5, time.Minute)
	assert.Nil(t, err)

	var response sdk.TxResponse
	err = net.Validators[0].ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &response)
	require.NoError(t, err)

	ret, err = tx.QueryTx(net.Validators[0].ClientCtx, response.TxHash)
	require.NoError(t, err)

	fmt.Printf(">>>>>>>>>>>>>>>%v\n", ret)
	fmt.Printf(">>>>>>>>>>>>>>>%v\n", ret.Code)
	require.Equal(t, uint32(0), ret.Code)
}
