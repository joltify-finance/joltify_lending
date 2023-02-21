package keeper

import (
	"context"
	types2 "github.com/cosmos/cosmos-sdk/codec/types"
	"time"

	coserrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k msgServer) updateInterestData(currentTime time.Time, interestData *types.BorrowInterest) sdk.Coin {

	var payment sdk.Coin
	latestTimeStamp := interestData.Payments[len(interestData.Payments)-1]
	delta := currentTime.Sub(latestTimeStamp.PaymentTime).Seconds()
	denom := interestData.Payments[0].PaymentAmount.Denom
	if int32(delta) > interestData.PayFreq*BASE {
		// we need to pay the whole month
		monthlyRatio := interestData.MonthlyRatio
		paymentAmount := monthlyRatio.MulInt(interestData.BorrowedLast.Amount).TruncateInt()
		payment = sdk.NewCoin(denom, paymentAmount)
	} else {
		r := CalculateInterestRate(interestData.Apy, int(interestData.PayFreq))
		interest := r.Power(uint64(delta))
		paymentAmount := interest.MulInt(interestData.BorrowedLast.Amount).TruncateInt()
		payment = sdk.NewCoin(denom, paymentAmount)
	}
	interestData.BorrowedLast = interestData.Borrowed

	// since the spv maynot pay the interest at exact next payment circle, we need to adjust it here

	thisPaymentTime := latestTimeStamp.GetPaymentTime().Add(time.Duration(interestData.PayFreq*BASE) * time.Second)

	currentPayment := types.PaymentItem{PaymentTime: thisPaymentTime, PaymentAmount: payment}
	interestData.Payments = append(interestData.Payments, &currentPayment)
	return payment

}

func (k msgServer) getAllInterestToBePaid(ctx sdk.Context, poolInfo *types.PoolInfo) sdkmath.Int {

	nftClasses := poolInfo.PoolNFTIds
	// the first element is the pool class, we skip it
	totalPayment := sdkmath.NewInt(0)
	for _, el := range nftClasses[1:] {
		class, found := k.nftKeeper.GetClass(ctx, el)
		if !found {
			panic(found)
		}
		data := class.GetData().GetCachedValue()
		interestData, ok := data.(types.BorrowInterest)
		if !ok {
			panic("not the borrow interest type")
		}

		thisBorrowInterest := k.updateInterestData(ctx.BlockTime(), &interestData)
		var err error
		class.Data, err = types2.NewAnyWithValue(&interestData)
		if err != nil {
			panic("pack class any data failed")
		}
		k.nftKeeper.SaveClass(ctx, class)
		totalPayment = totalPayment.Add(thisBorrowInterest.Amount)
	}
	return totalPayment
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
		return nil, types.ErrPoolNotActive
	}

	// todo here we allow all the poeple pay for the spv
	//if !spvAddress.Equals(poolInfo.OwnerAddress) {
	//	return nil, coserrors.Wrap(types.ErrUnauthorized, "only spv can pay the interest")
	//}

	if msg.Token.Denom != poolInfo.TotalAmount.Denom {
		return nil, coserrors.Wrapf(types.ErrInconsistencyToken, "pool denom %v and repaly is %v", poolInfo.TotalAmount.Denom, msg.Token.Denom)
	}

	totalAmountDue := k.getAllInterestToBePaid(ctx, &poolInfo)

	if msg.Token.Amount.LT(totalAmountDue) {
		return nil, coserrors.Wrapf(types.ErrInsufficientFund, "the interest is %v while you try to repay %v", totalAmountDue, msg.Token.Amount)
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, spvAddress, types.ModuleAccount, sdk.Coins{sdk.NewCoin(msg.Token.Denom, totalAmountDue)})
	if err != nil {
		return nil, coserrors.Wrapf(err, "fail to transfer the repayment from spv to module")
	}
	// finally, we update the poolinfo
	poolInfo.LastPaymentTime = ctx.BlockTime()
	k.SetPool(ctx, poolInfo)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRepayInterest,
			sdk.NewAttribute(types.AttributeCreator, msg.Creator),
			sdk.NewAttribute("amount", msg.Token.Amount.String()),
		),
	)

	return &types.MsgRepayInterestResponse{}, nil
}
