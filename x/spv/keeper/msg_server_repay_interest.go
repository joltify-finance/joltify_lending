package keeper

import (
	"context"

	coserrors "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k msgServer) getAllInterestToBePaid(poolInfo *types.PoolInfo) {

	//nftClasseslasses := poolInfo.PoolNFTIds
	//for _, el := range nftClasses {
	//
	//}

}
func (k msgServer) RepayInterest(goCtx context.Context, msg *types.MsgRepayInterest) (*types.MsgRepayInterestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	spvAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address %v", msg.Creator)
	}

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

	//k.getAllInterestToBePaid(poolInfo)

	// calcuate the interest amount
	i, err := CalculateInterestAmount(poolInfo.Apy, int(poolInfo.PayFreq))
	if err != nil {
		panic(err)
	}

	interestDue := sdk.NewDecFromInt(poolInfo.BorrowedAmount.Amount).Mul(i).TruncateInt()
	if msg.Token.Amount.LT(interestDue) {
		return nil, coserrors.Wrapf(types.ErrInsufficientFund, "the interest is %v while you try to repay %v", interestDue, msg.Token.Amount)
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, spvAddress, types.ModuleAccount, sdk.Coins{sdk.NewCoin(msg.Token.Denom, interestDue)})
	if err != nil {
		return nil, coserrors.Wrapf(err, "fail to transfer the repayment from spv to module")
	}
	// finally, we update the poolinfo
	poolInfo.LastPaymentTime = ctx.BlockTime()
	k.SetPool(ctx, poolInfo)
	return &types.MsgRepayInterestResponse{}, nil
}
