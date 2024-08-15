package keeper_test

import (
	"errors"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/dydx_helper/testutil/constants"
	keepertest "github.com/joltify-finance/joltify_lending/dydx_helper/testutil/keeper"
	"github.com/joltify-finance/joltify_lending/mocks"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/bridge/keeper"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/bridge/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestMsgServerAcknowledgeBridges(t *testing.T) {
	testMsg := constants.MsgAcknowledgeBridges_Ids0_1_Height0

	tests := map[string]struct {
		setupMocks   func(ctx sdk.Context, mck *mocks.BridgeKeeper)
		expectedResp *types.MsgAcknowledgeBridgesResponse
		expectedErr  string
	}{
		"Success": {
			setupMocks: func(ctx sdk.Context, mck *mocks.BridgeKeeper) {
				mck.On("AcknowledgeBridges", mock.Anything, testMsg.Events).Return(nil)
			},
			expectedResp: &types.MsgAcknowledgeBridgesResponse{},
		},
		"Failure: keeper error is propagated": {
			setupMocks: func(ctx sdk.Context, mck *mocks.BridgeKeeper) {
				mck.On("AcknowledgeBridges", mock.Anything, testMsg.Events).Return(
					errors.New("can't acknowledge bridges"),
				)
			},
			expectedErr: "can't acknowledge bridges",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			// Initialize Mocks and Context.
			mockKeeper := &mocks.BridgeKeeper{}
			msgServer := keeper.NewMsgServerImpl(mockKeeper)
			ctx, _, _, _, _, _, _ := keepertest.BridgeKeepers(t)
			tc.setupMocks(ctx, mockKeeper)

			resp, err := msgServer.AcknowledgeBridges(ctx, testMsg)

			// Assert msg server response.
			require.Equal(t, tc.expectedResp, resp)
			if tc.expectedErr != "" {
				require.Equal(t, tc.expectedErr, err.Error())
			}

			// Assert mock expectations.
			result := mockKeeper.AssertExpectations(t)
			require.True(t, result)
		})
	}
}
