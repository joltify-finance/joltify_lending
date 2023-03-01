package keeper

import (
	"context"
	"errors"
	"fmt"
	"time"

	types2 "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/gogo/protobuf/proto"

	coserrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k Keeper) updateInterestData(ctx sdk.Context, interestData *types.BorrowInterest, reserve sdk.Dec, firstBorrow bool) (sdk.Coin, error) {
	var payment, paymentToInvestor sdk.Coin
	// as the payment canot be happed at exact payfreq time, so we need to round down to the latest payment time
	currentTimeTruncated := ctx.BlockTime().Truncate(time.Duration(interestData.PayFreq) * time.Second)

	latestPaymentTime := interestData.Payments[len(interestData.Payments)-1].PaymentTime
	if firstBorrow {
		if ctx.BlockTime().Before(latestPaymentTime.Add(time.Duration(interestData.PayFreq) * time.Second)) {
			return sdk.Coin{}, errors.New("pay interest too early")
		}
	}

	delta := currentTimeTruncated.Sub(latestPaymentTime).Seconds()
	denom := interestData.Payments[0].PaymentAmount.Denom
	if int32(delta) >= interestData.PayFreq*BASE {
		// we need to pay the whole month
		monthlyRatio := interestData.MonthlyRatio
		paymentAmount := monthlyRatio.MulInt(interestData.BorrowedLast.Amount).TruncateInt()
		reservedAmount := sdk.NewDecFromInt(paymentAmount).Mul(reserve).TruncateInt()
		toInvestors := paymentAmount.Sub(reservedAmount)
		pReserve, found := k.GetReserve(ctx, denom)
		if !found {
			k.SetReserve(ctx, sdk.NewCoin(denom, reservedAmount))
		} else {
			pReserve = pReserve.AddAmount(reservedAmount)
			k.SetReserve(ctx, pReserve)
		}
		paymentToInvestor = sdk.NewCoin(denom, toInvestors)
		payment = sdk.NewCoin(denom, paymentAmount)
	} else {
		r := CalculateInterestRate(interestData.Apy, int(interestData.PayFreq))
		interest := r.Power(uint64(delta)).Sub(sdk.OneDec())

		paymentAmount := interest.MulInt(interestData.BorrowedLast.Amount).TruncateInt()
		reservedAmount := sdk.NewDecFromInt(paymentAmount).Mul(reserve).TruncateInt()
		toInvestors := paymentAmount.Sub(reservedAmount)

		pReserve, found := k.GetReserve(ctx, denom)
		if !found {
			k.SetReserve(ctx, sdk.NewCoin(denom, reservedAmount))
		} else {
			pReserve = pReserve.AddAmount(reservedAmount)
			k.SetReserve(ctx, pReserve)
		}
		paymentToInvestor = sdk.NewCoin(denom, toInvestors)
		payment = sdk.NewCoin(denom, paymentAmount)
	}
	interestData.BorrowedLast = interestData.Borrowed

	// since the spv may not pay the interest at exact next payment circle, we need to adjust it here
	// thisPaymentTime := latestTimeStamp.GetPaymentTime().Add(time.Duration(interestData.PayFreq*BASE) * time.Second)
	currentPayment := types.PaymentItem{PaymentTime: currentTimeTruncated, PaymentAmount: paymentToInvestor}
	interestData.Payments = append(interestData.Payments, &currentPayment)
	fmt.Printf(">>>total:>>>%v and %v to investor\n", payment.String(), paymentToInvestor.String())
	return payment, nil

}

func (k Keeper) getAllInterestToBePaid(ctx sdk.Context, poolInfo *types.PoolInfo) (sdkmath.Int, error) {

	nftClasses := poolInfo.PoolNFTIds
	// the first element is the pool class, we skip it
	totalPayment := sdkmath.NewInt(0)
	firstBorrow := true
	for _, el := range nftClasses {
		class, found := k.nftKeeper.GetClass(ctx, el)
		if !found {
			panic(found)
		}
		var borrowInterest types.BorrowInterest
		var err error
		err = proto.Unmarshal(class.Data.Value, &borrowInterest)
		if err != nil {
			panic(err)
		}
		thisBorrowInterest, err := k.updateInterestData(ctx, &borrowInterest, poolInfo.ReserveFactor, firstBorrow)
		firstBorrow = false
		if err != nil {
			return sdkmath.Int{}, err
		}
		class.Data, err = types2.NewAnyWithValue(&borrowInterest)
		if err != nil {
			panic("pack class any data failed")
		}
		k.nftKeeper.SaveClass(ctx, class)
		totalPayment = totalPayment.Add(thisBorrowInterest.Amount)
	}
	return totalPayment, nil
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
		return nil, coserrors.Wrapf(types.ErrInconsistencyToken, "pool denom %v and repay is %v", poolInfo.TotalAmount.Denom, msg.Token.Denom)
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, spvAddress, types.ModuleAccount, sdk.Coins{msg.Token})
	if err != nil {
		return nil, coserrors.Wrapf(err, "fail to transfer the repayment from spv to module")
	}

	poolInfo.EscrowInterestAmount = poolInfo.EscrowInterestAmount.Add(msg.Token)
	k.SetPool(ctx, poolInfo)

	return &types.MsgRepayInterestResponse{}, nil

}
