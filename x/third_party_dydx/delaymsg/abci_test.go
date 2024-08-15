package delaymsg_test

import (
	"testing"
	"time"

	sdktest "github.com/joltify-finance/joltify_lending/dydx_helper/testutil/sdk"
	"github.com/joltify-finance/joltify_lending/mocks"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/delaymsg"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/delaymsg/types"
)

func TestEndBlocker(t *testing.T) {
	k := &mocks.DelayMsgKeeper{}
	ctx := sdktest.NewContextWithBlockHeightAndTime(0, time.Now())
	// When DispatchMessagesForBlock is called, GetBlockMessageIds will be called, so we expect this call.
	// Return an empty list of message IDs and a false value to indicate that there are no messages.
	// In this case, the method should immediately return.
	k.On("GetBlockMessageIds", ctx, uint32(0)).Return(types.BlockMessageIds{}, false).Once()
	delaymsg.EndBlocker(ctx, k)
	k.AssertExpectations(t)
}
