package keeper

import (
	"errors"
	"fmt"
	"strings"
	"time"

	coserrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	types2 "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gogo/protobuf/proto"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func calculateTotalInterest(ctx sdk.Context, lendNFTs []string, nftKeeper types.NFTKeeper, updateNFT bool) (sdkmath.Int, error) {
	totalInterestUsd := sdk.NewInt(0)
	for _, el := range lendNFTs {
		ids := strings.Split(el, ":")
		thisNFT, found := nftKeeper.GetNFT(ctx, ids[0], ids[1])
		if !found {
			return sdkmath.Int{}, coserrors.Wrapf(types.ErrDepositorNotFound, "the given nft %v cannot ben found in storage", ids[1])
		}
		var investorInterestData types.NftInfo
		err := proto.Unmarshal(thisNFT.Data.Value, &investorInterestData)
		if err != nil {
			panic(err)
		}

		borrowClass, found := nftKeeper.GetClass(ctx, ids[0])
		if !found {
			panic("it should never fail to find the class")
		}

		var borrowClassInfo types.BorrowInterest
		err = proto.Unmarshal(borrowClass.Data.Value, &borrowClassInfo)
		if err != nil {
			panic(err)
		}

		allPayments := borrowClassInfo.Payments
		latestTimeStamp := time.Time{}
		lastPaymentSet := false

		// no new interest payment
		if len(allPayments) <= int(investorInterestData.PaymentOffset) {
			return sdk.ZeroInt(), nil
		}
		counter := 0
		allNewPayments := allPayments[investorInterestData.PaymentOffset:]
		for _, eachPayment := range allNewPayments {
			counter++
			if eachPayment.PaymentAmount.Amount.IsZero() {
				continue
			}
			classBorrowedAmount := eachPayment.BorrowedAmount
			paymentAmount := eachPayment.PaymentAmount
			// todo there may be the case that because of the truncate, the total payment is larger than the interest paid to investors
			// fixme we should NEVER calculate the interest after the pool status is in luquidation as the user ratio is not correct any more
			interestUsd := sdk.NewDecFromInt(paymentAmount.Amount).Mul(sdk.NewDecFromInt(investorInterestData.Borrowed.Amount)).Quo(sdk.NewDecFromInt(classBorrowedAmount.Amount)).TruncateInt()
			totalInterestUsd = totalInterestUsd.Add(interestUsd)
			latestTimeStamp = eachPayment.PaymentTime
			lastPaymentSet = true
			borrowClassInfo.InterestPaid = borrowClassInfo.InterestPaid.AddAmount(interestUsd)
		}
		if updateNFT && lastPaymentSet {
			investorInterestData.PaymentOffset += uint32(counter)
			investorInterestData.LastPayment = latestTimeStamp
			data, err := types2.NewAnyWithValue(&investorInterestData)
			if err != nil {
				panic("pack class any data failed")
			}
			thisNFT.Data = data
			err = nftKeeper.Update(ctx, thisNFT)
			if err != nil {
				panic(err)
			}

			data, err = types2.NewAnyWithValue(&borrowClassInfo)
			if err != nil {
				panic("pack class any data failed")
			}
			borrowClass.Data = data
			err = nftKeeper.UpdateClass(ctx, borrowClass)
			if err != nil {
				panic(err)
			}

		}
	}
	return totalInterestUsd, nil
}

func calculateTotalOutstandingInterest(ctx sdk.Context, lendNFTs []string, nftKeeper types.NFTKeeper, reserve sdk.Dec) (sdkmath.Int, error) {
	totalInterestUsd := sdk.NewInt(0)
	for _, el := range lendNFTs {
		ids := strings.Split(el, ":")
		thisNFT, found := nftKeeper.GetNFT(ctx, ids[0], ids[1])
		if !found {
			return sdkmath.Int{}, coserrors.Wrapf(types.ErrDepositorNotFound, "the given nft %v cannot ben found in storage", ids[1])
		}
		var interestData types.NftInfo
		err := proto.Unmarshal(thisNFT.Data.Value, &interestData)
		if err != nil {
			panic(err)
		}

		borrowClass, found := nftKeeper.GetClass(ctx, ids[0])
		if !found {
			panic("it should never fail to find the class")
		}

		var borrowClassInfo types.BorrowInterest
		err = proto.Unmarshal(borrowClass.Data.Value, &borrowClassInfo)
		if err != nil {
			panic(err)
		}

		lastBorrow := borrowClassInfo.BorrowDetails[len(borrowClassInfo.BorrowDetails)-1]
		lastPayment := borrowClassInfo.Payments[len(borrowClassInfo.Payments)-1]
		delta := uint64(ctx.BlockTime().Sub(lastPayment.PaymentTime).Seconds())
		factor := CalculateInterestFactor(borrowClassInfo.InterestSPY, sdk.NewIntFromUint64(delta))

		ratio := sdk.NewDecFromInt(interestData.Borrowed.Amount).Quo(sdk.NewDecFromInt(lastBorrow.BorrowedAmount.Amount))
		paymentAmountToInvestor := sdk.NewDecFromInt(lastBorrow.BorrowedAmount.Amount).Mul(sdk.OneDec().Sub(reserve))
		interestLocal := paymentAmountToInvestor.Mul(ratio).Mul(factor.Sub(sdk.OneDec())).TruncateInt()
		interest := outboundConvertToUSD(interestLocal, lastBorrow.ExchangeRatio)
		totalInterestUsd = totalInterestUsd.Add(interest)
	}
	return totalInterestUsd, nil
}

// func (k Keeper) QueryModuleBalance(ctx sdk.Context) {
//	acc := k.accKeeper.GetModuleAddress(types.ModuleAccount)
//
//	coins := k.bankKeeper.GetAllBalances(ctx, acc)
//	fmt.Printf(">>>>>>>>>>>>%v\n", coins)
// }

// tokenamount is the amount of token that to borrow and borrowedfix is the partial of the money we need to borrow
// rather then all the usable money
func (k Keeper) doBorrow(ctx sdk.Context, poolInfo *types.PoolInfo, usdTokenAmount sdk.Coin, needBankTransfer bool, depositors []*types.DepositorInfo, borrowedFix sdkmath.Int, userPoolLastPaymentTime bool) error {
	if usdTokenAmount.IsZero() {
		return nil
	}

	a, _ := denomConvertToLocalAndUsd(poolInfo.BorrowedAmount.Denom)
	localTokenAmount, ratio, err := k.inboundConvertFromUSDWithMarketID(ctx, denomConvertToMarketID(a), usdTokenAmount.Amount)
	if err != nil {
		return err
	}
	localToken := sdk.NewCoin(poolInfo.PoolDenomPrefix+usdTokenAmount.Denom, localTokenAmount)

	// create the new nft class for this borrow event
	classID := fmt.Sprintf("class-%v", poolInfo.Index[2:])
	poolClass, found := k.NftKeeper.GetClass(ctx, classID)
	if !found {
		panic("pool class must have already been set")
	}

	latestSeries := len(poolInfo.PoolNFTIds)

	currentBorrowClass := poolClass
	currentBorrowClass.Id = fmt.Sprintf("%v-%v", currentBorrowClass.Id, latestSeries)

	i, err := CalculateInterestAmount(poolInfo.Apy, int(poolInfo.PayFreq))
	if err != nil {
		panic(err)
	}

	rate := CalculateInterestRate(poolInfo.Apy, int(poolInfo.PayFreq))

	//fixme as we put the borrow time as the pool last payment time, so the SPV should pay the extra fee if he borrows a few days after the last cycle
	var paymentTime time.Time
	if userPoolLastPaymentTime {
		paymentTime = poolInfo.LastPaymentTime
	} else {
		paymentTime = ctx.BlockTime()
	}

	borrowDetails := make([]types.BorrowDetail, 1, 10)
	borrowDetails[0] = types.BorrowDetail{BorrowedAmount: localToken, TimeStamp: paymentTime, ExchangeRatio: ratio}
	firstPayment := types.PaymentItem{PaymentTime: paymentTime, PaymentAmount: sdk.NewCoin(poolInfo.TargetAmount.Denom, sdk.NewInt(0))}
	bi := types.BorrowInterest{
		PoolIndex:     poolInfo.Index,
		Apy:           poolInfo.Apy,
		PayFreq:       poolInfo.PayFreq,
		IssueTime:     ctx.BlockTime(),
		BorrowDetails: borrowDetails,
		MonthlyRatio:  i,
		InterestSPY:   rate,
		Payments:      []*types.PaymentItem{&firstPayment},
		InterestPaid:  sdk.NewCoin(poolInfo.TargetAmount.Denom, sdk.ZeroInt()), // using the usd
		AccInterest:   sdk.NewCoin(poolInfo.TargetAmount.Denom, sdk.ZeroInt()), // using the usd
	}

	data, err := types2.NewAnyWithValue(&bi)
	if err != nil {
		panic(err)
	}
	currentBorrowClass.Data = data
	err = k.NftKeeper.SaveClass(ctx, currentBorrowClass)
	if err != nil {
		return err
	}

	// update the borrow series
	poolInfo.PoolNFTIds = append(poolInfo.PoolNFTIds, currentBorrowClass.Id)

	// we start the project
	if len(poolInfo.PoolNFTIds) == 1 {
		poolInfo.ProjectDueTime = ctx.BlockTime().Add(time.Second * time.Duration(poolInfo.ProjectLength))
		poolInfo.PoolFirstDueTime = poolInfo.ProjectDueTime
	}

	err = k.processBorrow(ctx, poolInfo, currentBorrowClass, usdTokenAmount, localToken, ratio, depositors, borrowedFix)
	if err != nil {
		return err
	}

	// we finally update the pool info
	poolInfo.PoolStatus = types.PoolInfo_ACTIVE
	if !userPoolLastPaymentTime {
		poolInfo.LastPaymentTime = paymentTime
	}
	k.SetPool(ctx, *poolInfo)

	if needBankTransfer {
		// we transfer the fund from the module to the spv
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccount, poolInfo.OwnerAddress, sdk.NewCoins(usdTokenAmount))
		if err != nil {
			return err
		}
	}
	return nil
}

func (k Keeper) processBorrow(ctx sdk.Context, poolInfo *types.PoolInfo, nftClass nfttypes.Class, usdAmount, localToken sdk.Coin, ratio sdk.Dec, depositors []*types.DepositorInfo, borrowableFix sdkmath.Int) error {
	if poolInfo.UsableAmount.IsLT(usdAmount) {
		return types.ErrInsufficientFund
	}
	var borrowable sdkmath.Int
	if !borrowableFix.IsZero() {
		borrowable = borrowableFix
	} else {
		borrowable = poolInfo.UsableAmount.Amount
	}
	utilization := sdk.NewDecFromInt(usdAmount.Amount).Quo(sdk.NewDecFromInt(borrowable))

	var err error
	// we add the amount of the tokens that borrowed in the pool and decreases the borrowable
	poolInfo.BorrowedAmount = poolInfo.BorrowedAmount.Add(localToken)
	poolInfo.UsableAmount, err = poolInfo.UsableAmount.SafeSub(usdAmount)
	if err != nil {
		return types.ErrInsufficientFund
	}

	// we update each investor leftover
	return k.processInvestors(ctx, poolInfo, utilization, usdAmount.Amount, localToken.Amount, ratio, nftClass, depositors)
}

func (k Keeper) doProcessInvestor(ctx sdk.Context, depositor *types.DepositorInfo, lockedUsd, lockedLocal sdkmath.Int, nftTemplate nfttypes.NFT, nftClassId string, poolInfo *types.PoolInfo, useLastPaymentTime bool) error {
	depositor.LockedAmount = depositor.LockedAmount.Add(sdk.NewCoin(poolInfo.BorrowedAmount.Denom, lockedLocal))

	if depositor.WithdrawalAmount.Amount.LT(lockedUsd) {
		if lockedUsd.Sub(depositor.WithdrawalAmount.Amount).GT(sdk.NewIntFromUint64(5)) {
			panic("withdraw amount is small than the locked amount")
		}
		lockedUsd = depositor.WithdrawalAmount.Amount
	}

	depositor.WithdrawalAmount = depositor.WithdrawalAmount.SubAmount(lockedUsd)

	// nft ID is the hash(nft class ID, investorWallet)
	indexHash := crypto.Keccak256Hash([]byte(nftClassId), depositor.DepositorAddress)
	nftTemplate.Id = fmt.Sprintf("invoice-%v", indexHash.String()[2:])

	var paymentTime time.Time
	if useLastPaymentTime {
		paymentTime = poolInfo.LastPaymentTime
	} else {
		paymentTime = ctx.BlockTime()
	}

	lockedCoin := sdk.NewCoin(depositor.LockedAmount.Denom, lockedLocal)
	userData := types.NftInfo{Issuer: poolInfo.PoolName, Receiver: depositor.DepositorAddress.String(), Borrowed: lockedCoin, LastPayment: paymentTime}
	data, err := types2.NewAnyWithValue(&userData)
	if err != nil {
		panic("should never fail")
	}
	nftTemplate.Data = data
	err = k.NftKeeper.Mint(ctx, nftTemplate, depositor.DepositorAddress)
	if err != nil {
		return err
	}

	classIDAndNFTID := fmt.Sprintf("%v:%v", nftTemplate.ClassId, nftTemplate.Id)
	depositor.LinkedNFT = append(depositor.LinkedNFT, classIDAndNFTID)
	k.SetDepositor(ctx, *depositor)
	return nil
}

func (k Keeper) processInvestors(ctx sdk.Context, poolInfo *types.PoolInfo, utilization sdk.Dec, usdBorrowed, localAmount sdkmath.Int, ratio sdk.Dec, nftClass nfttypes.Class, depositors []*types.DepositorInfo) error {
	nftTemplate := nfttypes.NFT{
		ClassId: nftClass.Id,
		Uri:     nftClass.Uri,
		UriHash: nftClass.UriHash,
	}

	// now we update the depositor's withdrawal amount and locked amount
	var firstDepositor *types.DepositorInfo
	totalLocked := sdk.ZeroInt()
	totalLockedLocal := sdk.ZeroInt()
	if depositors != nil {
		for _, depositor := range depositors {

			if depositor.DepositType != types.DepositorInfo_unset {
				// since you have submitted the withdrawal/transfer request, we skipp the borrow from it
				continue
			}

			// this fix the bug that if we have 3 nodes with amount 0, 2,2. if we have 3 usd to be  withdrawal, it will
			// have the error as 2 2 because the first one despositor will not be involved in the borrow.
			// the correct way is to have the first non-zero withdrawal as the first depositor to be borrowed from
			if firstDepositor == nil && !depositor.WithdrawalAmount.IsZero() {
				firstDepositor = depositor
				continue
			}
			lockedUsd := sdk.NewDecFromInt(depositor.WithdrawalAmount.Amount).Mul(utilization).TruncateInt()
			lockedLocal := inboundConvertFromUSD(lockedUsd, ratio)
			if !lockedLocal.IsPositive() {
				continue
			}
			err := k.doProcessInvestor(ctx, depositor, lockedUsd, lockedLocal, nftTemplate, nftClass.Id, poolInfo, true)
			if err != nil {
				return err
			}
			totalLocked = totalLocked.Add(lockedUsd)
			totalLockedLocal = totalLockedLocal.Add(lockedLocal)
			continue
		}
	} else {
		k.IterateDepositors(ctx, poolInfo.Index, func(depositor types.DepositorInfo) (stop bool) {
			if depositor.DepositType != types.DepositorInfo_unset {
				// since you have submitted the withdrawal/transfer request, we skipp the borrow from it
				return false
			}

			// this fix the bug that if we have 3 nodes with amount 0, 2,2. if we have 3 usd to be  withdrawal, it will
			// have the error as 2 2 because the first one despositor will not be involved in the borrow.
			// the correct way is to have the first non-zero withdrawal as the first depositor to be borrowed from
			if firstDepositor == nil && !depositor.WithdrawalAmount.IsZero() {
				firstDepositor = &depositor
				return false
			}
			lockedUsd := sdk.NewDecFromInt(depositor.WithdrawalAmount.Amount).Mul(utilization).TruncateInt()
			lockedLocal := inboundConvertFromUSD(lockedUsd, ratio)
			if !lockedLocal.IsPositive() {
				return false
			}
			err := k.doProcessInvestor(ctx, &depositor, lockedUsd, lockedLocal, nftTemplate, nftClass.Id, poolInfo, true)
			if err != nil {
				ctx.Logger().Error(err.Error(), "error msg:", "failed to process investor")
				return false
			}
			totalLocked = totalLocked.Add(lockedUsd)
			totalLockedLocal = totalLockedLocal.Add(lockedLocal)

			return false
		})
	}
	// now we process the last one
	lockedUsd := usdBorrowed.Sub(totalLocked)
	lockedLocal := localAmount.Sub(totalLockedLocal)

	// we do not need to borrow from this investor
	if !lockedLocal.IsPositive() {
		return nil
	}
	err := k.doProcessInvestor(ctx, firstDepositor, lockedUsd, lockedLocal, nftTemplate, nftClass.Id, poolInfo, true)

	return err
}

func (k Keeper) handleClassLeftover(ctx sdk.Context, poolinfo types.PoolInfo) sdk.Coin {
	nfts := poolinfo.PoolNFTIds
	var err error
	leftover := sdk.NewCoin(poolinfo.TargetAmount.Denom, sdk.ZeroInt())
	for _, el := range nfts {
		class, found := k.NftKeeper.GetClass(ctx, el)
		if !found {
			panic("class not found")
		}
		var borrowInterest types.BorrowInterest
		err = proto.Unmarshal(class.Data.Value, &borrowInterest)
		if err != nil {
			panic(err)
		}
		delta := borrowInterest.AccInterest.Sub(borrowInterest.InterestPaid)
		if delta.IsPositive() {
			leftover = leftover.Add(delta)
		}
	}
	return leftover
}

func (k Keeper) cleanupDepositor(ctx sdk.Context, poolInfo types.PoolInfo, depositor types.DepositorInfo) (sdkmath.Int, error) {
	interest, err := calculateTotalInterest(ctx, depositor.LinkedNFT, k.NftKeeper, true)
	if err != nil {
		panic(err)
	}

	err = k.processEachWithdrawReq(ctx, depositor, true, poolInfo.PrincipalPaymentExchangeRatio)
	if err != nil {
		ctx.Logger().Error("fail to process partial principal", err.Error())
		return sdk.ZeroInt(), err
	}

	exchange := poolInfo.PrincipalPaymentExchangeRatio
	usdLocked := outboundConvertToUSD(depositor.LockedAmount.Amount, exchange)
	totalPaidAmount := usdLocked.Add(depositor.WithdrawalAmount.Amount)
	totalPaidAmount = totalPaidAmount.Add(interest)
	totalPaidAmount = totalPaidAmount.Add(depositor.PendingInterest.Amount)

	poolInfo.BorrowedAmount, err = poolInfo.BorrowedAmount.SafeSub(depositor.LockedAmount)
	if err != nil {
		return sdk.ZeroInt(), err
	}

	// fix the issue 10. since we have not to add the transfer owner withdrawal amount to the pool, we do not need to deducted it here.
	if depositor.DepositType != types.DepositorInfo_processed {
		poolInfo.UsableAmount, err = poolInfo.UsableAmount.SafeSub(depositor.WithdrawalAmount)
		if err != nil {
			return sdk.ZeroInt(), err
		}

	}

	poolInfo.ProcessedTransferAccounts = deleteElement(poolInfo.ProcessedTransferAccounts, depositor.DepositorAddress)

	if k.isEmptyPool(ctx, poolInfo) {
		ctx.Logger().Info("we delete the pool as it is empty")
		// we transfer the leftover back to spv
		totalReturn := poolInfo.EscrowPrincipalAmount.AddAmount(poolInfo.EscrowInterestAmount)
		// we handle the leftover of each class
		leftover := k.handleClassLeftover(ctx, poolInfo)
		reserve, found := k.GetReserve(ctx, "ausdc")
		if found {
			reserve = reserve.Add(leftover)
			k.SetReserve(ctx, reserve)
		}
		err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccount, poolInfo.OwnerAddress, sdk.NewCoins(totalReturn))
		if err != nil {
			return totalPaidAmount, err
		}
		k.DelPool(ctx, poolInfo.Index)
		k.SetHistoryPool(ctx, poolInfo)

	} else {
		k.SetPool(ctx, poolInfo)
	}
	depositor.DepositType = types.DepositorInfo_deactive
	depositor.LinkedNFT = []string{}
	depositor.WithdrawalAmount = sdk.NewCoin(poolInfo.TargetAmount.Denom, sdk.ZeroInt())
	depositor.LockedAmount = sdk.NewCoin(depositor.LockedAmount.Denom, sdk.ZeroInt())
	k.DelDepositor(ctx, depositor)
	k.SetDepositorHistory(ctx, depositor)
	return totalPaidAmount, nil
}

func (k Keeper) doProcessLiquidationForInvestor(ctx sdk.Context, lendNFTs []string) (sdkmath.Int, error) {
	totalRedeem := sdk.NewInt(0)
	for _, el := range lendNFTs {
		ids := strings.Split(el, ":")
		thisNFT, found := k.NftKeeper.GetNFT(ctx, ids[0], ids[1])
		if !found {
			return sdkmath.Int{}, coserrors.Wrapf(types.ErrDepositorNotFound, "the given nft %v cannot ben found in storage", ids[1])
		}
		var investorInterestData types.NftInfo
		err := proto.Unmarshal(thisNFT.Data.Value, &investorInterestData)
		if err != nil {
			panic(err)
		}

		borrowClass, found := k.NftKeeper.GetClass(ctx, ids[0])
		if !found {
			panic("it should never fail to find the class")
		}

		var borrowClassInfo types.BorrowInterest
		err = proto.Unmarshal(borrowClass.Data.Value, &borrowClassInfo)
		if err != nil {
			panic(err)
		}

		allLiquidationPayments := borrowClassInfo.LiquidationItems
		latestTimeStamp := time.Time{}

		if len(allLiquidationPayments) <= int(investorInterestData.LiquidationPaymentOffset) {
			return sdk.ZeroInt(), nil
		}
		counter := 0
		allNewLiquidationPayments := allLiquidationPayments[investorInterestData.LiquidationPaymentOffset:]

		classLastBorrow := borrowClassInfo.BorrowDetails[len(borrowClassInfo.BorrowDetails)-1].BorrowedAmount
		for _, eachLiquidationPayment := range allNewLiquidationPayments {
			counter++
			if eachLiquidationPayment.Amount.Amount.IsZero() {
				continue
			}
			liquidationPaymentAmount := eachLiquidationPayment.Amount
			// todo there may be the case that because of the tucate, the total payment is larger than the interest paid to investors
			// fixme we should NEVER calculate the interest after the pool status is in luquidation as the user ratio is not correct any more
			investorRedeemAmount := sdk.NewDecFromInt(liquidationPaymentAmount.Amount).Mul(sdk.NewDecFromInt(investorInterestData.Borrowed.Amount)).Quo(sdk.NewDecFromInt(classLastBorrow.Amount)).TruncateInt()
			totalRedeem = totalRedeem.Add(investorRedeemAmount)
			latestTimeStamp = eachLiquidationPayment.LiquidationPaymentTime
			borrowClassInfo.TotalPaidLiquidationAmount = borrowClassInfo.TotalPaidLiquidationAmount.Add(investorRedeemAmount)
		}
		investorInterestData.LastPayment = latestTimeStamp
		investorInterestData.LiquidationPaymentOffset += uint32(counter)
		data, err := types2.NewAnyWithValue(&investorInterestData)
		if err != nil {
			panic("pack class any data failed")
		}
		thisNFT.Data = data
		err = k.NftKeeper.Update(ctx, thisNFT)
		if err != nil {
			panic(err)
		}

		data, err = types2.NewAnyWithValue(&borrowClassInfo)
		if err != nil {
			panic("pack class any data failed")
		}
		borrowClass.Data = data
		err = k.NftKeeper.UpdateClass(ctx, borrowClass)
		if err != nil {
			panic(err)
		}
	}

	return totalRedeem, nil
}

func (k Keeper) updateDepositorStatus(ctx sdk.Context, poolInfo *types.PoolInfo) {
	totalAdjust := sdk.NewInt(0)

	k.IterateDepositors(ctx, poolInfo.Index, func(depositor types.DepositorInfo) (stop bool) {
		if depositor.DepositType == types.DepositorInfo_unset {
			if depositor.WithdrawalAmount.Amount.Sub(poolInfo.MinDepositAmount).IsNegative() {
				depositor.DepositType = types.DepositorInfo_unusable
				totalAdjust = totalAdjust.Add(depositor.WithdrawalAmount.Amount)
				k.SetDepositor(ctx, depositor)
			}
		}
		return false
	})

	// it should never be negative, otherwise panic as serious calculation error
	poolInfo.UsableAmount = poolInfo.UsableAmount.SubAmount(totalAdjust)
	k.SetPool(ctx, *poolInfo)
}

func (k Keeper) syncThePayment(ctx sdk.Context, poolInfo *types.PoolInfo) error {

	totalAmountDue, err := k.syncAllInterestToBePaid(ctx, poolInfo)
	if err != nil {
		return err
	}

	if totalAmountDue.IsZero() {
		// no interest to be paid
		k.Logger(ctx).Info("no interest to be paid", "pool index", poolInfo.Index)
		return
	}

	poolInfo.EscrowInterestAmount = poolInfo.EscrowInterestAmount.Sub(totalAmountDue)
	if poolInfo.EscrowInterestAmount.IsNegative() {
		poolInfo.NegativeInterestCounter++
	} else {
		poolInfo.NegativeInterestCounter = 0
	}
	if poolInfo.NegativeInterestCounter > types.MaxLiquidattion {
		poolInfo.PoolStatus = types.PoolInfo_Liquidation
	}

	// finally, we update the poolinfo
	k.SetPool(ctx, *poolInfo)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRepayInterest,
			sdk.NewAttribute("amount", totalAmountDue.String()),
		),
	)
	return nil

}

func (k Keeper) syncAllInterestToBePaid(ctx sdk.Context, poolInfo *types.PoolInfo) (sdkmath.Int, error) {
	nftClasses := poolInfo.PoolNFTIds
	// the first element is the pool class, we skip it
	totalPayment := sdkmath.NewInt(0)
	firstBorrow := true
	var exchangeRatio sdk.Dec

	if poolInfo.PoolStatus == types.PoolInfo_INACTIVE {
		return sdk.ZeroInt(), errors.New("no interest to be paid")
	}

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
		class, found := k.NftKeeper.GetClass(ctx, el)
		if !found {
			panic(found)
		}
		var borrowInterest types.BorrowInterest
		var err error
		err = proto.Unmarshal(class.Data.Value, &borrowInterest)
		if err != nil {
			panic(err)
		}
		thisBorrowInterest, err := k.syncInterestData(ctx, &borrowInterest, poolInfo.ReserveFactor, firstBorrow, exchangeRatio)
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
		err = k.NftKeeper.UpdateClass(ctx, class)
		if err != nil {
			return sdkmath.Int{}, err
		}
		totalPayment = totalPayment.Add(thisBorrowInterest.Amount)
		firstBorrow = false
	}
	return totalPayment, nil
}

func (k Keeper) syncInterestData(ctx sdk.Context, interestData *types.BorrowInterest, reserve sdk.Dec, firstBorrow bool, exchangeRatio sdk.Dec) (sdk.Coin, time.Time, error) {
	var payment, paymentToInvestor sdk.Coin
	var thisPaymentTime time.Time
	// as the payment cannot be happened at exact payfreq time, so we need to round down to the latest payment time
	// currentTimeTruncated := ctx.BlockTime().Truncate(time.Duration(interestData.PayFreq) * time.Second)
	currentTime := ctx.BlockTime().Truncate(time.Duration(interestData.PayFreq*BASE) * time.Second)

	latestPaymentTime := interestData.Payments[len(interestData.Payments)-1].PaymentTime
	if firstBorrow {
		if ctx.BlockTime().Before(latestPaymentTime.Add(time.Duration(interestData.PayFreq) * time.Second)) {
			return sdk.Coin{}, time.Time{}, errors.New("pay interest too early")
		}
	}
	lastBorrow := interestData.BorrowDetails[len(interestData.BorrowDetails)-1].BorrowedAmount
	lastBorrowTime := interestData.BorrowDetails[len(interestData.BorrowDetails)-1].TimeStamp
	delta := lastBorrowTime.Sub(latestPaymentTime)
	denom := interestData.Payments[0].PaymentAmount.Denom
	if int32(delta) >= interestData.PayFreq*BASE {
		// we need to pay the whole month
		freqRatio := interestData.MonthlyRatio
		paymentAmount := freqRatio.Mul(sdk.NewDecFromInt(lastBorrow.Amount)).TruncateInt()
		if paymentAmount.IsZero() {
			return sdk.Coin{Denom: lastBorrow.Denom, Amount: sdk.ZeroInt()}, time.Time{}, nil
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
			return sdk.Coin{Denom: lastBorrow.Denom, Amount: sdk.ZeroInt()}, time.Time{}, nil
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
	return payment, thisPaymentTime, nil
}
