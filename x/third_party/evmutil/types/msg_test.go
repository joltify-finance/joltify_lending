package types_test

import (
	"testing"

	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/x/third_party/evmutil/testutil"
	"github.com/joltify-finance/joltify_lending/x/third_party/evmutil/types"
	"github.com/stretchr/testify/require"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestMsgConvertCoinToERC20(t *testing.T) {
	app.SetSDKConfig()

	type errArgs struct {
		expectPass bool
		contains   string
	}

	tests := []struct {
		name          string
		giveInitiator string
		giveReceiver  string
		giveAmount    sdk.Coin
		errArgs       errArgs
	}{
		{
			"valid",
			"jolt17cnnh8076vgv2y2lln6qtrlv83k05f7zch5dqv",
			"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
			sdk.NewCoin("erc20/weth", sdkmath.NewInt(1234)),
			errArgs{
				expectPass: true,
			},
		},
		{
			"invalid - odd length hex address",
			"jolt17cnnh8076vgv2y2lln6qtrlv83k05f7zch5dqv",
			"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc",
			sdk.NewCoin("erc20/weth", sdkmath.NewInt(1234)),
			errArgs{
				expectPass: false,
				contains:   "Receiver is not a valid hex address: invalid address",
			},
		},
		{
			"invalid - zero amount",
			"jolt17cnnh8076vgv2y2lln6qtrlv83k05f7zch5dqv",
			"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
			sdk.NewCoin("erc20/weth", sdkmath.NewInt(0)),
			errArgs{
				expectPass: false,
				contains:   "amount cannot be zero",
			},
		},
		{
			"invalid - negative amount",
			"jolt17cnnh8076vgv2y2lln6qtrlv83k05f7zch5dqv",
			"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
			// Create manually so there is no validation
			sdk.Coin{Denom: "erc20/weth", Amount: sdkmath.NewInt(-1234)},
			errArgs{
				expectPass: false,
				contains:   "negative coin amount",
			},
		},
		{
			"invalid - empty denom",
			"jolt17cnnh8076vgv2y2lln6qtrlv83k05f7zch5dqv",
			"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
			sdk.Coin{Denom: "", Amount: sdkmath.NewInt(-1234)},
			errArgs{
				expectPass: false,
				contains:   "invalid denom",
			},
		},
		{
			"invalid - invalid denom",
			"jolt17cnnh8076vgv2y2lln6qtrlv83k05f7zch5dqv",
			"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
			sdk.Coin{Denom: "h", Amount: sdkmath.NewInt(-1234)},
			errArgs{
				expectPass: false,
				contains:   "invalid denom",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			msg := types.NewMsgConvertCoinToERC20(
				tc.giveInitiator,
				tc.giveReceiver,
				tc.giveAmount,
			)
			err := msg.ValidateBasic()

			if tc.errArgs.expectPass {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.errArgs.contains)
			}
		})
	}
}

func TestMsgConvertERC20ToCoin(t *testing.T) {
	app.SetSDKConfig()

	type errArgs struct {
		expectPass bool
		contains   string
	}

	tests := []struct {
		name         string
		receiver     string
		initiator    string
		contractAddr string
		amount       sdkmath.Int
		errArgs      errArgs
	}{
		{
			"valid",
			"jolt17cnnh8076vgv2y2lln6qtrlv83k05f7zch5dqv",
			"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
			"0x404F9466d758eA33eA84CeBE9E444b06533b369e",
			sdkmath.NewInt(1234),
			errArgs{
				expectPass: true,
			},
		},
		{
			"invalid - odd length hex address",
			"jolt17cnnh8076vgv2y2lln6qtrlv83k05f7zch5dqv",
			"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc",
			"0x404F9466d758eA33eA84CeBE9E444b06533b369e",
			sdkmath.NewInt(1234),
			errArgs{
				expectPass: false,
				contains:   "initiator is not a valid hex address",
			},
		},
		{
			"invalid - zero amount",
			"jolt17cnnh8076vgv2y2lln6qtrlv83k05f7zch5dqv",
			"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
			"0x404F9466d758eA33eA84CeBE9E444b06533b369e",
			sdkmath.NewInt(0),
			errArgs{
				expectPass: false,
				contains:   "amount cannot be zero",
			},
		},
		{
			"invalid - negative amount",
			"jolt17cnnh8076vgv2y2lln6qtrlv83k05f7zch5dqv",
			"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
			"0x404F9466d758eA33eA84CeBE9E444b06533b369e",
			sdkmath.NewInt(-1234),
			errArgs{
				expectPass: false,
				contains:   "amount cannot be zero or less",
			},
		},
		{
			"invalid - invalid contract address",
			"jolt17cnnh8076vgv2y2lln6qtrlv83k05f7zch5dqv",
			"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
			"0x404F9466d758eA33eA84CeBE9E444b06533b369",
			sdkmath.NewInt(1234),
			errArgs{
				expectPass: false,
				contains:   "erc20 contract address is not a valid hex address",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			msg := types.MsgConvertERC20ToCoin{
				Initiator:        tc.initiator,
				Receiver:         tc.receiver,
				JoltERC20Address: tc.contractAddr,
				Amount:           tc.amount,
			}
			err := msg.ValidateBasic()

			if tc.errArgs.expectPass {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.errArgs.contains)
			}
		})
	}
}

func TestConvertCosmosCoinToERC20_ValidateBasic(t *testing.T) {
	validJoltAddr := app.RandomAddress()
	validHexAddr, _ := testutil.RandomEvmAccount()
	invalidAddr := "not-an-address"
	validAmount := sdk.NewInt64Coin("hard", 5e3)

	testCases := []struct {
		name        string
		initiator   string
		receiver    string
		amount      sdk.Coin
		expectedErr string
	}{
		{
			name:        "valid",
			initiator:   validJoltAddr.String(),
			receiver:    validHexAddr.String(),
			amount:      validAmount,
			expectedErr: "",
		},
		{
			name:        "invalid - sending to jolt addr",
			initiator:   validJoltAddr.String(),
			receiver:    app.RandomAddress().String(),
			amount:      validAmount,
			expectedErr: "receiver is not a valid hex address",
		},
		{
			name:        "invalid - invalid initiator",
			initiator:   invalidAddr,
			receiver:    app.RandomAddress().String(),
			amount:      validAmount,
			expectedErr: "invalid initiator address",
		},
		{
			name:        "invalid - invalid receiver",
			initiator:   validJoltAddr.String(),
			receiver:    invalidAddr,
			amount:      validAmount,
			expectedErr: "receiver is not a valid hex address",
		},
		{
			name:        "invalid - invalid amount - nil",
			initiator:   validJoltAddr.String(),
			receiver:    validHexAddr.String(),
			amount:      sdk.Coin{},
			expectedErr: "invalid coins",
		},
		{
			name:        "invalid - invalid amount - zero",
			initiator:   validJoltAddr.String(),
			receiver:    validHexAddr.String(),
			amount:      sdk.NewInt64Coin("magic", 0),
			expectedErr: "invalid coins",
		},
		{
			name:        "invalid - invalid amount - negative",
			initiator:   validJoltAddr.String(),
			receiver:    validHexAddr.String(),
			amount:      sdk.Coin{Denom: "magic", Amount: sdkmath.NewInt(-42)},
			expectedErr: "invalid coins",
		},
		{
			name:        "invalid - invalid amount - invalid denom",
			initiator:   validJoltAddr.String(),
			receiver:    validHexAddr.String(),
			amount:      sdk.Coin{Denom: "", Amount: sdkmath.NewInt(42)},
			expectedErr: "invalid coins",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			msg := types.NewMsgConvertCosmosCoinToERC20(
				tc.initiator,
				tc.receiver,
				tc.amount,
			)
			err := msg.ValidateBasic()

			if tc.expectedErr != "" {
				require.ErrorContains(t, err, tc.expectedErr)
			} else {
				require.NoError(t, err)
				require.Equal(t, "evmutil", msg.Route())
				require.Equal(t, "evmutil_convert_cosmos_coin_to_erc20", msg.Type())
				require.NotPanics(t, func() { _ = msg.GetSignBytes() })
			}
		})
	}
}

func TestConvertCosmosCoinToERC20_GetSigners(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		initiator := app.RandomAddress()
		signers := types.MsgConvertCosmosCoinToERC20{
			Initiator: initiator.String(),
		}.GetSigners()
		require.Len(t, signers, 1)
		require.Equal(t, initiator, signers[0])
	})

	t.Run("panics when depositor is invalid", func(t *testing.T) {
		require.Panics(t, func() {
			types.MsgConvertCosmosCoinToERC20{
				Initiator: "not-an-address",
			}.GetSigners()
		})
	})
}

func TestConvertCosmosCoinFromERC20_ValidateBasic(t *testing.T) {
	validHexAddr := testutil.RandomEvmAddress()
	validJoltAddr := app.RandomAddress()
	invalidAddr := "not-an-address"
	validAmount := sdk.NewInt64Coin("hard", 5e3)

	testCases := []struct {
		name        string
		initiator   string
		receiver    string
		amount      sdk.Coin
		expectedErr string
	}{
		{
			name:        "valid",
			initiator:   validHexAddr.String(),
			receiver:    validJoltAddr.String(),
			amount:      validAmount,
			expectedErr: "",
		},
		{
			name:        "invalid - sending to 0x addr",
			initiator:   validHexAddr.String(),
			receiver:    testutil.RandomEvmAddress().Hex(),
			amount:      validAmount,
			expectedErr: "invalid receiver address",
		},
		{
			name:        "invalid - invalid initiator",
			initiator:   invalidAddr,
			receiver:    app.RandomAddress().String(),
			amount:      validAmount,
			expectedErr: "initiator is not a valid hex address",
		},
		{
			name:        "invalid - invalid receiver",
			initiator:   validHexAddr.String(),
			receiver:    invalidAddr,
			amount:      validAmount,
			expectedErr: "invalid receiver address",
		},
		{
			name:        "invalid - invalid amount - nil",
			initiator:   validHexAddr.String(),
			receiver:    validJoltAddr.String(),
			amount:      sdk.Coin{},
			expectedErr: "invalid coins",
		},
		{
			name:        "invalid - invalid amount - zero",
			initiator:   validHexAddr.String(),
			receiver:    validJoltAddr.String(),
			amount:      sdk.NewInt64Coin("magic", 0),
			expectedErr: "invalid coins",
		},
		{
			name:        "invalid - invalid amount - negative",
			initiator:   validHexAddr.String(),
			receiver:    validJoltAddr.String(),
			amount:      sdk.Coin{Denom: "magic", Amount: sdkmath.NewInt(-42)},
			expectedErr: "invalid coins",
		},
		{
			name:        "invalid - invalid amount - invalid denom",
			initiator:   validHexAddr.String(),
			receiver:    validJoltAddr.String(),
			amount:      sdk.Coin{Denom: "", Amount: sdkmath.NewInt(42)},
			expectedErr: "invalid coins",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			msg := types.NewMsgConvertCosmosCoinFromERC20(
				tc.initiator,
				tc.receiver,
				tc.amount,
			)
			err := msg.ValidateBasic()

			if tc.expectedErr != "" {
				require.ErrorContains(t, err, tc.expectedErr)
			} else {
				require.NoError(t, err)
				require.Equal(t, "evmutil", msg.Route())
				require.Equal(t, "evmutil_convert_cosmos_coin_from_erc20", msg.Type())
				require.NotPanics(t, func() { _ = msg.GetSignBytes() })
			}
		})
	}
}

func TestConvertCosmosCoinFromERC20_GetSigners(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		initiator0x := testutil.RandomEvmAddress()
		initiator := sdk.AccAddress(initiator0x.Bytes())
		signers := types.MsgConvertCosmosCoinFromERC20{
			Initiator: initiator0x.Hex(),
		}.GetSigners()
		require.Len(t, signers, 1)
		require.Equal(t, initiator, signers[0])
	})

	t.Run("panics when depositor is invalid", func(t *testing.T) {
		require.Panics(t, func() {
			types.MsgConvertCosmosCoinFromERC20{
				Initiator: "not-an-address",
			}.GetSigners()
		})
	})
}
