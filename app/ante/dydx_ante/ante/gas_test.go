package ante_test

import (
	"reflect"
	"testing"

	"cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	testapp "github.com/joltify-finance/joltify_lending/testutil/dydx/testutil/app"
	assets "github.com/joltify-finance/joltify_lending/x/third_party_dydx/assets/types"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/joltify-finance/joltify_lending/app/ante/dydx_ante/ante"
	testante "github.com/joltify-finance/joltify_lending/testutil/dydx/testutil/ante"
	"github.com/joltify-finance/joltify_lending/testutil/dydx/testutil/constants"
	pricestypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/prices/types"

	"github.com/stretchr/testify/require"
)

const freeInfiniteGasMeterType = "*types.freeInfiniteGasMeter"

func TestValidateMsgType_FreeInfiniteGasDecorator(t *testing.T) {
	tests := map[string]struct {
		msgOne sdk.Msg
		msgTwo sdk.Msg

		expectFreeInfiniteGasMeter bool
		expectedErr                error
	}{
		"no freeInfiniteGasMeter: no msg": {
			expectFreeInfiniteGasMeter: false,
		},
		"yes freeInfiniteGasMeter: single msg, MsgPlaceOrder": {
			msgOne: constants.Msg_PlaceOrder,

			expectFreeInfiniteGasMeter: true,
		},
		"yes freeInfiniteGasMeter: single msg, Msg_CancelOrder": {
			msgOne: constants.Msg_CancelOrder,

			expectFreeInfiniteGasMeter: true,
		},
		"yes freeInfiniteGasMeter: single msg, MsgUpdateMarketPrices": {
			msgOne: &pricestypes.MsgUpdateMarketPrices{}, // app-injected.

			expectFreeInfiniteGasMeter: true,
		},
		"no freeInfiniteGasMeter: single msg": {
			msgOne: &testdata.TestMsg{Signers: []string{"meh"}},

			expectFreeInfiniteGasMeter: false,
		},
		"no freeInfiniteGasMeter: multi msg, MsgUpdateMarketPrices": {
			msgOne: &pricestypes.MsgUpdateMarketPrices{}, // app-injected.
			msgTwo: &testdata.TestMsg{Signers: []string{"meh"}},

			expectFreeInfiniteGasMeter: false,
		},
		"no freeInfiniteGasMeter: mult msgs, NO off-chain single msg clob tx": {
			msgOne: &testdata.TestMsg{Signers: []string{"meh"}},
			msgTwo: &testdata.TestMsg{Signers: []string{"meh"}},

			expectFreeInfiniteGasMeter: false,
		},
		"no freeInfiniteGasMeter: mult msgs, MsgCancelOrder with Transfer": {
			msgOne: constants.Msg_CancelOrder,
			msgTwo: constants.Msg_Transfer,

			expectedErr: sdkerrors.ErrInvalidRequest,
		},
		"no freeInfiniteGasMeter: mult msgs, two MsgCancelOrder": {
			msgOne: constants.Msg_CancelOrder,
			msgTwo: constants.Msg_CancelOrder,

			expectedErr: sdkerrors.ErrInvalidRequest,
		},
		"no freeInfiniteGasMeter: mult msgs, MsgPlaceOrder with Transfer": {
			msgOne: constants.Msg_PlaceOrder,
			msgTwo: constants.Msg_Transfer,

			expectedErr: sdkerrors.ErrInvalidRequest,
		},
		"no freeInfiniteGasMeter: mult msgs, two MsgPlaceOrder": {
			msgOne: constants.Msg_PlaceOrder,
			msgTwo: constants.Msg_PlaceOrder,

			expectedErr: sdkerrors.ErrInvalidRequest,
		},
		"no freeInfiniteGasMeter: mult msgs, MsgPlaceOrder and MsgCancelOrder": {
			msgOne: constants.Msg_PlaceOrder,
			msgTwo: constants.Msg_CancelOrder,

			expectedErr: sdkerrors.ErrInvalidRequest,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			suite := testante.SetupTestSuite(t, true)
			suite.TxBuilder = suite.ClientCtx.TxConfig.NewTxBuilder()

			wrappedHandler := ante.NewFreeInfiniteGasDecorator()
			antehandler := sdk.ChainAnteDecorators(wrappedHandler)

			msgs := make([]sdk.Msg, 0)
			if tc.msgOne != nil {
				msgs = append(msgs, tc.msgOne)
			}
			if tc.msgTwo != nil {
				msgs = append(msgs, tc.msgTwo)
			}

			require.NoError(t, suite.TxBuilder.SetMsgs(msgs...))

			// Empty private key, so tx's signature should be empty.
			privs, accNums, accSeqs := []cryptotypes.PrivKey{}, []uint64{}, []uint64{}

			tx, err := suite.CreateTestTx(
				suite.Ctx,
				privs,
				accNums,
				accSeqs,
				suite.Ctx.ChainID(),
				signing.SignMode_SIGN_MODE_DIRECT,
			)
			require.NoError(t, err)

			resultCtx, err := antehandler(suite.Ctx, tx, false)
			require.ErrorIs(t, tc.expectedErr, err)

			meter := resultCtx.GasMeter()

			if !tc.expectFreeInfiniteGasMeter || tc.expectedErr != nil {
				require.NotEqual(t, freeInfiniteGasMeterType, reflect.TypeOf(meter).String())
				require.Equal(t, suite.Ctx, resultCtx)
			} else {
				require.Equal(t, freeInfiniteGasMeterType, reflect.TypeOf(meter).String())
			}
		})
	}
}

func TestSubmitTxnWithGas(t *testing.T) {
	tests := map[string]struct {
		gasFee       sdk.Coins
		responseCode uint32
		logMessage   string
	}{
		"Success - 5 cents usdc gas fee": {
			gasFee:       constants.TestFeeCoins_5Cents,
			responseCode: errors.SuccessABCICode,
		},
		"Success - 5 cents native token gas fee": {
			gasFee:       constants.TestFeeCoins_5Cents_NativeToken,
			responseCode: errors.SuccessABCICode,
		},
		"Failure: 0 gas fee": {
			gasFee:       sdk.Coins{},
			responseCode: sdkerrors.ErrInsufficientFee.ABCICode(),
			logMessage: "insufficient fees; got:  required: 25000000000000000adv4tnt," +
				"25000ibc/65D0BEC6DAD96C7F5043D1E54E54B6BB5D5B3AEC3FF6CEBB75B9E059F3580EA3: insufficient fee",
		},
		"Failure: unsupported gas fee denom": {
			gasFee: sdk.Coins{
				// 1BTC, which is not supported as a gas fee denom, and should be plenty to cover gas.
				sdk.NewCoin(constants.BtcUsd.Denom, sdkmath.NewInt(100_000_000)),
			},
			responseCode: sdkerrors.ErrInsufficientFee.ABCICode(),
			logMessage: "insufficient fees; got: 100000000btc-denom required: 25000000000000000adv4tnt," +
				"25000ibc/65D0BEC6DAD96C7F5043D1E54E54B6BB5D5B3AEC3FF6CEBB75B9E059F3580EA3: insufficient fee",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			msg := &bank.MsgSend{
				FromAddress: constants.BobAccAddress.String(),
				ToAddress:   constants.AliceAccAddress.String(),
				Amount: []sdk.Coin{
					sdk.NewCoin(assets.AssetUsdc.Denom, sdkmath.NewInt(1)),
				},
			}

			tApp := testapp.NewTestAppBuilder(t).Build()
			ctx := tApp.InitChain()

			msgSendCheckTx := testapp.MustMakeCheckTxWithPrivKeySupplier(
				ctx,
				tApp.App,
				testapp.MustMakeCheckTxOptions{
					AccAddressForSigning: constants.BobAccAddress.String(),
					Gas:                  constants.TestGasLimit,
					FeeAmt:               tc.gasFee,
				},
				constants.GetPrivateKeyFromAddress,
				msg,
			)

			checkTx := tApp.CheckTx(msgSendCheckTx)
			// Sanity check that gas was used.
			require.Greater(t, checkTx.GasUsed, int64(0))
			require.Equal(t, tc.responseCode, checkTx.Code)
			if tc.responseCode != errors.SuccessABCICode {
				require.Equal(t, tc.logMessage, checkTx.Log)
			}
		})
	}
}
