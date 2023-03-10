package types_test

import (
	"errors"
	"fmt"
	"testing"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto"
)

func TestMsgClaim_Validate(t *testing.T) {
	validAddress := sdk.AccAddress(crypto.AddressHash([]byte("KavaTest1"))).String()

	type expectedErr struct {
		wraps error
		pass  bool
	}
	type msgArgs struct {
		sender        string
		denomsToClaim types2.Selections
	}
	tests := []struct {
		name    string
		msgArgs msgArgs
		expect  expectedErr
	}{
		{
			name: "normal multiplier is valid",
			msgArgs: msgArgs{
				sender: validAddress,
				denomsToClaim: types2.Selections{
					{
						Denom:          "jolt",
						MultiplierName: "large",
					},
				},
			},
			expect: expectedErr{
				pass: true,
			},
		},
		{
			name: "empty multiplier name is invalid",
			msgArgs: msgArgs{
				sender: validAddress,
				denomsToClaim: types2.Selections{
					{
						Denom:          "jolt",
						MultiplierName: "",
					},
				},
			},
			expect: expectedErr{
				wraps: types2.ErrInvalidMultiplier,
			},
		},
		{
			name: "empty denoms to claim is not valid",
			msgArgs: msgArgs{
				sender:        validAddress,
				denomsToClaim: types2.Selections{},
			},
			expect: expectedErr{
				wraps: types2.ErrInvalidClaimDenoms,
			},
		},
		{
			name: "nil denoms to claim is not valid",
			msgArgs: msgArgs{
				sender:        validAddress,
				denomsToClaim: nil,
			},
			expect: expectedErr{
				wraps: types2.ErrInvalidClaimDenoms,
			},
		},
		{
			name: "invalid sender",
			msgArgs: msgArgs{
				sender: "",
				denomsToClaim: types2.Selections{
					{
						Denom:          "jolt",
						MultiplierName: "medium",
					},
				},
			},
			expect: expectedErr{
				wraps: sdkerrors.ErrInvalidAddress,
			},
		},
		{
			name: "invalid claim denom",
			msgArgs: msgArgs{
				sender: validAddress,
				denomsToClaim: types2.Selections{
					{
						Denom:          "a denom string that is invalid because it is much too long",
						MultiplierName: "medium",
					},
				},
			},
			expect: expectedErr{
				wraps: types2.ErrInvalidClaimDenoms,
			},
		},
		{
			name: "too many claim denoms",
			msgArgs: msgArgs{
				sender:        validAddress,
				denomsToClaim: tooManySelections(),
			},
			expect: expectedErr{
				wraps: types2.ErrInvalidClaimDenoms,
			},
		},
		{
			name: "duplicated claim denoms",
			msgArgs: msgArgs{
				sender: validAddress,
				denomsToClaim: types2.Selections{
					{
						Denom:          "jolt",
						MultiplierName: "medium",
					},
					{
						Denom:          "jolt",
						MultiplierName: "large",
					},
				},
			},
			expect: expectedErr{
				wraps: types2.ErrInvalidClaimDenoms,
			},
		},
	}

	for _, tc := range tests {
		msgClaimJoltReward := types2.NewMsgClaimJoltReward(tc.msgArgs.sender, tc.msgArgs.denomsToClaim)
		msgClaimDelegatorReward := types2.NewMsgClaimDelegatorReward(tc.msgArgs.sender, tc.msgArgs.denomsToClaim)
		msgClaimSwapReward := types2.NewMsgClaimSwapReward(tc.msgArgs.sender, tc.msgArgs.denomsToClaim)
		msgClaimSavingsReward := types2.NewMsgClaimSavingsReward(tc.msgArgs.sender, tc.msgArgs.denomsToClaim)
		msgs := []sdk.Msg{&msgClaimJoltReward, &msgClaimDelegatorReward, &msgClaimSwapReward, &msgClaimSavingsReward}
		for _, msg := range msgs {
			t.Run(tc.name, func(t *testing.T) {
				err := msg.ValidateBasic()
				if tc.expect.pass {
					require.NoError(t, err)
				} else {
					require.Truef(t, errors.Is(err, tc.expect.wraps), "expected error '%s' was not actual '%s'", tc.expect.wraps, err)
				}
			})
		}
	}
}

func TestMsgClaimUSDXMintingReward_Validate(t *testing.T) {
	validAddress := sdk.AccAddress(crypto.AddressHash([]byte("KavaTest1"))).String()

	type expectedErr struct {
		wraps error
		pass  bool
	}
	type msgArgs struct {
		sender         string
		multiplierName string
	}
	tests := []struct {
		name    string
		msgArgs msgArgs
		expect  expectedErr
	}{
		{
			name: "normal multiplier is valid",
			msgArgs: msgArgs{
				sender:         validAddress,
				multiplierName: "large",
			},
			expect: expectedErr{
				pass: true,
			},
		},
		{
			name: "invalid sender",
			msgArgs: msgArgs{
				sender:         "",
				multiplierName: "medium",
			},
			expect: expectedErr{
				wraps: sdkerrors.ErrInvalidAddress,
			},
		},
		{
			name: "empty multiplier is invalid",
			msgArgs: msgArgs{
				sender:         validAddress,
				multiplierName: "",
			},
			expect: expectedErr{
				wraps: types2.ErrInvalidMultiplier,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			msg := types2.NewMsgClaimUSDXMintingReward(tc.msgArgs.sender, tc.msgArgs.multiplierName)

			err := msg.ValidateBasic()
			if tc.expect.pass {
				require.NoError(t, err)
			} else {
				require.Truef(t, errors.Is(err, tc.expect.wraps), "expected error '%s' was not actual '%s'", tc.expect.wraps, err)
			}
		})
	}
}

func tooManySelections() types2.Selections {
	selections := make(types2.Selections, types2.MaxDenomsToClaim+1)
	for i := range selections {
		selections[i] = types2.Selection{
			Denom:          fmt.Sprintf("denom%d", i),
			MultiplierName: "large",
		}
	}
	return selections
}
