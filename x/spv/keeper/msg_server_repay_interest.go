package keeper

import (
	"context"
	"time"

	coserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k msgServer) updateInterestData(currentTime time.Time, interestData types.BorrowInterest) {

	var payment sdk.Coin
	latestTimeStamp := interestData.Payments[len(interestData.Payments())-1]
	delta := int32(currentTime.Sub(latestTimeStamp).Minutes())
	if delta > interestData.PayFreq*BASE {
		// we need to may the whole month
		payment = interestData.CyclePayment
	} else {
		r := CalculateInterestRate(interestData.Apy, int(interestData.PayFreq))
		interest := r.Power(uint64(delta))
		interestData.Borrowed.Amount(interest)
	}

}

func (k msgServer) getAllInterestToBePaid(ctx sdk.Context, poolInfo *types.PoolInfo) {

	nftClasseslasses := poolInfo.PoolNFTIds
	// the first element is the pool class, we skip it
	totalPayment := sdk.NewInt(0)
	for _, el := range nftClasseslasses[1:] {
		class, found := k.nftKeeper.GetClass(ctx, el)
		if !found {
			panic(found)
		}
		data := class.GetData().GetCachedValue()
		interestData, ok := data.(types.BorrowInterest)
		if !ok {
			panic("not the borrow interest type")
		}

		k.updateInterestData(interestData)

		currentTime := ctx.BlockTime()

		delta := int(currentTime.Sub(interestData.IssueTime).Seconds())

	}

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
