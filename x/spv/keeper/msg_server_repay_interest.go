package keeper

import (
	"context"
	"errors"
	"time"

	types2 "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/gogo/protobuf/proto"

	coserrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k Keeper) updateInterestData(ctx sdk.Context, interestData *types.BorrowInterest, reserve sdk.Dec, firstBorrow bool, exchangeRatio sdk.Dec) (sdk.Coin, error) {
	var payment, paymentToInvestor sdk.Coin
	var thisPaymentTime time.Time
	// as the payment cannot be happened at exact payfreq time, so we need to round down to the latest payment time
	//currentTimeTruncated := ctx.BlockTime().Truncate(time.Duration(interestData.PayFreq) * time.Second)
	currentTime := ctx.BlockTime().Truncate(time.Duration(interestData.PayFreq*BASE) * time.Second)

	latestPaymentTime := interestData.Payments[len(interestData.Payments)-1].PaymentTime
	if firstBorrow {
		if ctx.BlockTime().Before(latestPaymentTime.Add(time.Duration(interestData.PayFreq) * time.Second)) {
			return sdk.Coin{}, errors.New("pay interest too early")
		}
	}
	delta := currentTime.Sub(latestPaymentTime).Seconds()
	denom := interestData.Payments[0].PaymentAmount.Denom
	lastBorrow := interestData.BorrowDetails[len(interestData.BorrowDetails)-1].BorrowedAmount
	if int32(delta) >= interestData.PayFreq*BASE {
		// we need to pay the whole month
		freqRatio := interestData.MonthlyRatio
		paymentAmount := freqRatio.Mul(sdk.NewDecFromInt(lastBorrow.Amount)).TruncateInt()
		if paymentAmount.IsZero() {
			return sdk.Coin{Denom: lastBorrow.Denom, Amount: sdk.ZeroInt()}, nil
		}

		paymentAmountUsd := outboundConvertToUSD(paymentAmount, exchangeRatio)
		reservedAmount := sdk.NewDecFromInt(paymentAmountUsd).Mul(reserve).TruncateInt()
		toInvestors := paymentAmountUsd.Sub(reservedAmount)
		pReserve, found := k.GetReserve(ctx, denom)
		if !found {
			k.SetReserve(ctx, sdk.NewCoin(denom, reservedAmount))
		} else {
			pReserve = pReserve.AddAmount(reservedAmount)
			k.SetReserve(ctx, pReserve)
		}
		paymentToInvestor = sdk.NewCoin(denom, toInvestors)
		payment = sdk.NewCoin(denom, paymentAmountUsd)
		thisPaymentTime = latestPaymentTime.Add(time.Duration(interestData.PayFreq*BASE) * time.Second).Truncate(time.Duration(interestData.PayFreq*BASE) * time.Second)
	} else {
		currentTimeTruncated := ctx.BlockTime().Truncate(time.Duration(interestData.PayFreq) * time.Second)
		if currentTimeTruncated.Before(latestPaymentTime) {
			return sdk.Coin{Denom: lastBorrow.Denom, Amount: sdk.ZeroInt()}, nil
		}
		deltaTruncated := currentTimeTruncated.Sub(latestPaymentTime).Seconds()
		r := CalculateInterestRate(interestData.Apy, int(interestData.PayFreq))
		interest := r.Power(uint64(deltaTruncated)).Sub(sdk.OneDec())

		usdInterest := interest.Mul(exchangeRatio)
		paymentAmountUsd := usdInterest.MulInt(lastBorrow.Amount).TruncateInt()
		reservedAmountUsd := sdk.NewDecFromInt(paymentAmountUsd).Mul(reserve).TruncateInt()
		toInvestors := paymentAmountUsd.Sub(reservedAmountUsd)

		pReserve, found := k.GetReserve(ctx, denom)
		if !found {
			k.SetReserve(ctx, sdk.NewCoin(denom, reservedAmountUsd))
		} else {
			pReserve = pReserve.AddAmount(reservedAmountUsd)
			k.SetReserve(ctx, pReserve)
		}
		paymentToInvestor = sdk.NewCoin(denom, toInvestors)
		payment = sdk.NewCoin(denom, paymentAmountUsd)
		thisPaymentTime = latestPaymentTime.Add(time.Duration(interestData.PayFreq*BASE) * time.Second).Truncate(time.Duration(interestData.PayFreq*BASE) * time.Second)
	}

	// since the spv may not pay the interest at exact next payment circle, we need to adjust it here
	currentPayment := types.PaymentItem{PaymentTime: thisPaymentTime, PaymentAmount: paymentToInvestor, BorrowedAmount: lastBorrow}
	interestData.Payments = append(interestData.Payments, &currentPayment)
	interestData.AccInterest = interestData.AccInterest.Add(paymentToInvestor)
	return payment, nil

}

// getAllinterestToBePaid returns the total interest to be paid for all the borrows in the pool using the
// LOCAL currency
func (k Keeper) getAllInterestToBePaid(ctx sdk.Context, poolInfo *types.PoolInfo) (sdkmath.Int, error) {

	nftClasses := poolInfo.PoolNFTIds
	// the first element is the pool class, we skip it
	totalPayment := sdkmath.NewInt(0)
	firstBorrow := true
	var exchangeRatio sdk.Dec
	if poolInfo.InterestPrepayment == nil || poolInfo.InterestPrepayment.Counter == 0 {
		a, _ := denomConvertToLocalAndUsd(poolInfo.BorrowedAmount.Denom)
		price, err := k.priceFeedKeeper.GetCurrentPrice(ctx, denomConvertToMarketID(a))
		if err != nil {
			panic(err)
		}
		exchangeRatio = price.Price
	} else {
		exchangeRatio = poolInfo.InterestPrepayment.ExchangeRatio
		poolInfo.InterestPrepayment.Counter--
		if poolInfo.InterestPrepayment.Counter == 0 {
			poolInfo.InterestPrepayment = nil
		}
	}
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
		thisBorrowInterest, err := k.updateInterestData(ctx, &borrowInterest, poolInfo.ReserveFactor, firstBorrow, exchangeRatio)
		if err != nil {
			return sdkmath.Int{}, err
		}
		if thisBorrowInterest.Amount.IsZero() {
			continue
		}
		class.Data, err = types2.NewAnyWithValue(&borrowInterest)
		if err != nil {
			panic("pack class any data failed")
		}
		err = k.nftKeeper.UpdateClass(ctx, class)
		if err != nil {
			return sdkmath.Int{}, err
		}
		totalPayment = totalPayment.Add(thisBorrowInterest.Amount)
		firstBorrow = false
	}
	return totalPayment, nil
}

func (k msgServer) calculatePaymentMonth(ctx sdk.Context, poolInfo types.PoolInfo, marketId string, totalPaid sdkmath.Int) (int32, sdkmath.Int, sdkmath.Int, sdk.Dec, error) {

	paymentAmount, err := k.calculateTotalDueInterest(ctx, poolInfo)
	if err != nil {
		return 0, sdkmath.ZeroInt(), sdkmath.ZeroInt(), sdk.ZeroDec(), err
	}
	usdEachMonth, ratio, err := k.outboundConvertToUSDWithMarketID(ctx, marketId, paymentAmount)
	if err != nil {
		return 0, sdkmath.ZeroInt(), sdk.ZeroInt(), sdk.ZeroDec(), err
	}
	counter := totalPaid.Quo(usdEachMonth)
	return int32(counter.Int64()), usdEachMonth.Mul(counter), usdEachMonth, ratio, nil
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
	if poolInfo.PoolStatus == types.PoolInfo_FROZEN || poolInfo.PoolStatus == types.PoolInfo_INACTIVE || poolInfo.PoolStatus == types.PoolInfo_FREEZING {
		return nil, types.ErrPoolNotActive
	}

	if msg.Token.Denom != poolInfo.TargetAmount.Denom {
		return nil, coserrors.Wrapf(types.ErrInconsistencyToken, "pool denom %v and repay is %v", poolInfo.TargetAmount.Denom, msg.Token.Denom)
	}

	if poolInfo.BorrowedAmount.IsZero() || msg.Token.Amount.IsZero() {
		return nil, coserrors.Wrapf(types.ErrInvalidParameter, "borrow amount is zero, no interest to be paid or interest paid is zero")
	}

	if poolInfo.InterestPrepayment != nil {
		return nil, coserrors.Wrapf(types.ErrInvalidParameter, "we have the prepayment interest, not accepting new interest payment")
	}

	if !poolInfo.EscrowInterestAmount.IsNegative() {

		a, _ := denomConvertToLocalAndUsd(poolInfo.BorrowedAmount.Denom)
		marketID := denomConvertToMarketID(a)
		counter, interestReceived, eachMonth, ratio, err := k.calculatePaymentMonth(ctx, poolInfo, marketID, msg.Token.Amount)
		if err != nil {
			return nil, coserrors.Wrapf(err, "calculate payment month failed")
		}

		if counter < 1 {
			return nil, coserrors.Wrapf(types.ErrInsufficientFund, "you must pay at least one interest cycle (%v)", eachMonth)
		}

		poolInfo.EscrowInterestAmount = poolInfo.EscrowInterestAmount.Add(interestReceived)
		prepayment := types.InterestPrepayment{
			Counter:       counter,
			ExchangeRatio: ratio,
		}

		poolInfo.InterestPrepayment = &prepayment

		err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, spvAddress, types.ModuleAccount, sdk.Coins{sdk.NewCoin(msg.Token.Denom, interestReceived)})
		if err != nil {
			return nil, coserrors.Wrapf(err, "fail to transfer the repayment from spv to module")
		}

		k.SetPool(ctx, poolInfo)
		return &types.MsgRepayInterestResponse{}, nil
	}

	//leftover := poolInfo.EscrowInterestAmount.Add(msg.Token.Amount)
	ownInterest := poolInfo.EscrowInterestAmount.Abs()
	leftover := msg.Token.Amount.Sub(ownInterest)
	if leftover.IsNegative() {
		return nil, coserrors.Wrapf(types.ErrInsufficientFund, "you must pay all the outstanding interest")
	}

	a, _ := denomConvertToLocalAndUsd(poolInfo.BorrowedAmount.Denom)
	marketID := denomConvertToMarketID(a)
	counter, interestReceived, eachMonthPayment, ratio, err := k.calculatePaymentMonth(ctx, poolInfo, marketID, leftover)
	if err != nil {
		return nil, coserrors.Wrapf(err, "calculate payment month failed")
	}

	if counter < 1 {
		return nil, coserrors.Wrapf(types.ErrInsufficientFund, "you must pay at least one interest cycle (%v)", eachMonthPayment)
	}
	poolInfo.EscrowInterestAmount = interestReceived
	prepayment := types.InterestPrepayment{
		Counter:       counter,
		ExchangeRatio: ratio,
	}

	poolInfo.InterestPrepayment = &prepayment
	totalGetFromSPV := ownInterest.Add(interestReceived)

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, spvAddress, types.ModuleAccount, sdk.Coins{sdk.NewCoin(msg.Token.Denom, totalGetFromSPV)})
	if err != nil {
		return nil, coserrors.Wrapf(err, "fail to transfer the repayment from spv to module")
	}

	k.SetPool(ctx, poolInfo)
	return &types.MsgRepayInterestResponse{}, nil

}
