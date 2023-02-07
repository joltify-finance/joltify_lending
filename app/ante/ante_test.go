package ante_test

import (
	"os"
	"testing"
	"time"

	"github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/types"

	"github.com/cosmos/cosmos-sdk/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmdb "github.com/tendermint/tm-db"
)

func TestMain(m *testing.M) {
	app.SetSDKConfig()
	os.Exit(m.Run())
}

func TestAppAnteHandler_AuthorizedMempool(t *testing.T) {
	testPrivKeys, testAddresses := app.GeneratePrivKeyAddressPairs(10)
	deputy := testAddresses[2]
	deputyKey := testPrivKeys[2]
	oracles := testAddresses[3:6]
	oraclesKeys := testPrivKeys[3:6]
	manual := testAddresses[6:]
	manualKeys := testPrivKeys[6:]

	encodingConfig := app.MakeEncodingConfig()

	var options app.Options
	tApp := app.TestApp{
		App: *app.NewApp(
			log.NewNopLogger(),
			tmdb.NewMemDB(),
			app.DefaultNodeHome,
			nil,
			encodingConfig,
			options,
		),
	}

	chainID := "jolttest_1-1"
	tApp = tApp.InitializeFromGenesisStatesWithTimeAndChainID(
		time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC),
		chainID, nil, nil,
		app.NewFundedGenStateWithSameCoins(
			tApp.AppCodec(),
			sdk.NewCoins(sdk.NewInt64Coin("ujolt", 1e6)),
			testAddresses,
		),
		newPricefeedGenStateMulti(tApp.AppCodec(), oracles),
	)

	testcases := []struct {
		name       string
		address    sdk.AccAddress
		privKey    cryptotypes.PrivKey
		expectPass bool
	}{
		{
			name:       "oracle",
			address:    oracles[1],
			privKey:    oraclesKeys[1],
			expectPass: true,
		},
		{
			name:       "deputy",
			address:    deputy,
			privKey:    deputyKey,
			expectPass: true,
		},
		{
			name:       "manual",
			address:    manual[1],
			privKey:    manualKeys[1],
			expectPass: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			stdTx, err := helpers.GenTx(
				encodingConfig.TxConfig,
				[]sdk.Msg{
					banktypes.NewMsgSend(
						tc.address,
						testAddresses[0],
						sdk.NewCoins(sdk.NewInt64Coin("ujolt", 1_000_000)),
					),
				},
				sdk.NewCoins(), // no fee
				helpers.DefaultGenTxGas,
				chainID,
				[]uint64{0},
				[]uint64{0}, // fixed sequence numbers will cause tests to fail sig verification if the same address is used twice
				tc.privKey,
			)
			require.NoError(t, err)
			txBytes, err := encodingConfig.TxConfig.TxEncoder()(stdTx)
			require.NoError(t, err)

			res := tApp.CheckTx(
				abci.RequestCheckTx{
					Tx:   txBytes,
					Type: abci.CheckTxType_New,
				},
			)

			if tc.expectPass {
				require.Zero(t, res.Code, res.Log)
			} else {
				require.NotZero(t, res.Code)
			}
		})
	}
}

func newPricefeedGenStateMulti(cdc codec.JSONCodec, oracles []sdk.AccAddress) app.GenesisState {
	pfGenesis := types.GenesisState{
		Params: types.Params{
			Markets: []types.Market{
				{MarketID: "btc:usd", BaseAsset: "btc", QuoteAsset: "usd", Oracles: oracles, Active: true},
			},
		},
	}
	return app.GenesisState{types.ModuleName: cdc.MustMarshalJSON(&pfGenesis)}
}

// we do not disable any msg, so we disable this test.
//func TestAppAnteHandler_RejectMsgsInAuthz(t *testing.T) {
//	testPrivKeys, testAddresses := app.GeneratePrivKeyAddressPairs(10)
//
//	newMsgGrant := func(msgTypeUrl string) *authz.MsgGrant {
//		msg, err := authz.NewMsgGrant(
//			testAddresses[0],
//			testAddresses[1],
//			authz.NewGenericAuthorization(msgTypeUrl),
//			time.Date(9000, 1, 1, 0, 0, 0, 0, time.UTC),
//		)
//		if err != nil {
//			panic(err)
//		}
//		return msg
//	}
//
//	chainID := "jolttest_1-1"
//	encodingConfig := app.MakeEncodingConfig()
//
//	testcases := []struct {
//		name         string
//		msg          sdk.Msg
//		expectedCode uint32
//	}{
//		{
//			name:         "MsgCreateVestingAccount is blocked",
//			msg:          newMsgGrant(sdk.MsgTypeURL(&vestingtypes.MsgCreateVestingAccount{})),
//			expectedCode: sdkerrors.ErrUnauthorized.ABCICode(),
//		},
//	}
//
//	for _, tc := range testcases {
//		t.Run(tc.name, func(t *testing.T) {
//			tApp := app.NewTestApp()
//
//			tApp = tApp.InitializeFromGenesisStatesWithTimeAndChainID(
//				time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC),
//				chainID,
//				app.NewFundedGenStateWithSameCoins(
//					tApp.AppCodec(),
//					sdk.NewCoins(sdk.NewInt64Coin("ujolt", 1e6)),
//					testAddresses,
//				),
//			)
//
//			stdTx, err := helpers.GenTx(
//				encodingConfig.TxConfig,
//				[]sdk.Msg{tc.msg},
//				sdk.NewCoins(), // no fee
//				helpers.DefaultGenTxGas,
//				chainID,
//				[]uint64{0},
//				[]uint64{0},
//				testPrivKeys[0],
//			)
//			require.NoError(t, err)
//			txBytes, err := encodingConfig.TxConfig.TxEncoder()(stdTx)
//			require.NoError(t, err)
//
//			resCheckTx := tApp.CheckTx(
//				abci.RequestCheckTx{
//					Tx:   txBytes,
//					Type: abci.CheckTxType_New,
//				},
//			)
//			require.Equal(t, resCheckTx.Code, tc.expectedCode, resCheckTx.Log)
//
//			resDeliverTx := tApp.DeliverTx(
//				abci.RequestDeliverTx{
//					Tx: txBytes,
//				},
//			)
//			require.Equal(t, resDeliverTx.Code, tc.expectedCode, resDeliverTx.Log)
//		})
//	}
//}
