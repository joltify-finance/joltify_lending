package keeper

import (
	"context"
	coserrors "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k msgServer) RepayInterest(goCtx context.Context, msg *types.MsgRepayInterest) (*types.MsgRepayInterestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	poolInfo, found := k.GetPools(ctx, msg.GetPoolIndex())
	if !found {
		return nil, coserrors.Wrapf(types.ErrPoolNotFound, "pool %v not found", msg.PoolIndex)
	}
	if poolInfo.PoolStatus == types.PoolInfo_CLOSED || poolInfo.PoolStatus == types.PoolInfo_INACTIVE {
		return nil, types.ErrPoolClosed
	}

	if msg.Token.Denom != poolInfo.TotalAmount.Denom {
		return nil, coserrors.Wrapf(types.ErrInconsistencyToken, "pool denom %v and repaly is %v", poolInfo.TotalAmount.Denom, msg.Token.Denom)
	}

	//finally, we update the poolinfo
	poolInfo.LastPaymentTime = ctx.BlockTime()
	k.SetPool(ctx, poolInfo)
	return &types.MsgRepayInterestResponse{}, nil
}
