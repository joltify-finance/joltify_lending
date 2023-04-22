package keeper

import (
	"strings"
	"time"

	coserrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	types2 "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/gogo/protobuf/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k Keeper) HandleInterest(ctx sdk.Context, poolInfo *types.PoolInfo) error {
	totalAmountDue, err := k.getAllInterestToBePaid(ctx, poolInfo)
	if err != nil {
		ctx.Logger().Info(err.Error())
		if err.Error() == "pay interest too early" {
			return err
		}
		panic(err)
	}

	if err != nil {
		panic("should never fail to convert to usd")
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
	currentTimeTruncated := ctx.BlockTime().Truncate(time.Duration(poolInfo.PayFreq) * time.Second)
	poolInfo.LastPaymentTime = currentTimeTruncated
	k.SetPool(ctx, *poolInfo)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRepayInterest,
			sdk.NewAttribute("amount", totalAmountDue.String()),
		),
	)
	return nil
}

// HandleTransfer if the pool have enough withdrawal amount, we can return the full amount of the investors
// otherwise, we can only return the partial of the principal
func (k Keeper) HandleTransfer(ctx sdk.Context, poolInfo *types.PoolInfo) bool {
	var err error
	var depositors []*types.DepositorInfo
	totalLockedAmount := sdkmath.ZeroInt()
	if len(poolInfo.TransferAccounts) == 0 {
		return false
	}
	for _, el := range poolInfo.TransferAccounts {
		d, found := k.GetDepositor(ctx, poolInfo.Index, el)
		if !found {
			panic("should never fail to find the depositor")
		}
		depositors = append(depositors, &d)
		totalLockedAmount = totalLockedAmount.Add(d.LockedAmount.Amount)
	}

	if totalLockedAmount.Equal(poolInfo.WithdrawProposalAmount.Amount) {
		panic("the total locked of each a")
	}
	poolInfo.TransferAccounts = []sdk.AccAddress{}

	a, _ := denomConvertToLocalAndUsd(poolInfo.WithdrawProposalAmount.Denom)
	usdTotalLocked, ratio, err := k.outboundConvertToUSDWithMarketID(ctx, denomConvertToMarketID(a), totalLockedAmount)
	if err != nil {
		ctx.Logger().Error(err.Error(), "outbound convert with market ID fail to convert", poolInfo.WithdrawProposalAmount)
		return false
	}

	// borrowable is larger than the total required, so we can return the money to these investors
	if poolInfo.UsableAmount.Amount.GTE(usdTotalLocked) {
		for _, el := range depositors {
			interest, err := calculateTotalInterest(ctx, el.LinkedNFT, k.nftKeeper, true)
			if err != nil {
				panic(err)
			}

			err = k.processEachWithdrawReq(ctx, *el, true, ratio)
			if err != nil {
				panic(err)
			}

			el.LinkedNFT = []string{}
			el.PendingInterest = el.PendingInterest.AddAmount(interest)

			usdLocked := outboundConvertToUSD(el.LockedAmount.Amount, ratio)
			el.WithdrawalAmount = el.WithdrawalAmount.AddAmount(usdLocked)
			el.LockedAmount = sdk.NewCoin(el.LockedAmount.Denom, sdk.ZeroInt())
			el.DepositType = types.DepositorInfo_deposit_close
			k.SetDepositor(ctx, *el)
		}

		poolInfo.BorrowedAmount = poolInfo.BorrowedAmount.SubAmount(totalLockedAmount)
		err := k.doBorrow(ctx, poolInfo, sdk.NewCoin(poolInfo.UsableAmount.Denom, totalLockedAmount), false, nil, sdk.ZeroInt())
		if err != nil {
			panic(err)
		}
		return true
	}
	// now we process the partial transfer
	totalBorrowableFromPrevious := sdkmath.ZeroInt()
	for i, el := range depositors {
		interest, err := calculateTotalInterest(ctx, el.LinkedNFT, k.nftKeeper, true)
		if err != nil {
			panic(err)
		}

		err = k.processEachWithdrawReq(ctx, *el, true, ratio)
		if err != nil {
			panic(err)
		}
		depositors[i].DepositType = types.DepositorInfo_unset
		el.LinkedNFT = []string{}
		el.PendingInterest = el.PendingInterest.AddAmount(interest)

		usdLocked := outboundConvertToUSD(el.LockedAmount.Amount, ratio)
		el.WithdrawalAmount = el.WithdrawalAmount.AddAmount(usdLocked)
		el.LockedAmount = sdk.NewCoin(el.LockedAmount.Denom, sdk.ZeroInt())
		totalBorrowableFromPrevious = totalBorrowableFromPrevious.Add(el.WithdrawalAmount.Amount)
	}

	needToBorrowedFromPreviousInvestorsUsd := usdTotalLocked.Sub(poolInfo.UsableAmount.Amount)
	// we need to adjust the amount of the borrowed and borrowable for the pool as we borrow again from these investors
	poolInfo.BorrowedAmount = poolInfo.BorrowedAmount.SubAmount(totalLockedAmount)
	poolInfo.UsableAmount = poolInfo.UsableAmount.AddAmount(needToBorrowedFromPreviousInvestorsUsd)
	err = k.doBorrow(ctx, poolInfo, sdk.NewCoin(poolInfo.UsableAmount.Denom, needToBorrowedFromPreviousInvestorsUsd), false, depositors, totalBorrowableFromPrevious)
	if err != nil {
		panic(err)
	}

	// now we process the partial transfer
	for i := range depositors {
		depositors[i].DepositType = types.DepositorInfo_processed
		k.SetDepositor(ctx, *depositors[i])
	}
	err = k.doBorrow(ctx, poolInfo, poolInfo.UsableAmount, false, nil, sdk.ZeroInt())
	if err != nil {
		panic(err)
	}
	return true
}

func (k Keeper) updateClassAndBurnNFT(ctx sdk.Context, classID, nftID string, burnNFT bool, exchangeRatio sdk.Dec) error {
	thisClass, found := k.nftKeeper.GetClass(ctx, classID)
	if !found {
		return coserrors.Wrapf(types.ErrClassNotFound, "the class cannot be found")
	}

	var borrowClassInfo types.BorrowInterest
	err := proto.Unmarshal(thisClass.Data.Value, &borrowClassInfo)
	if err != nil {
		panic(err)
	}
	lastBorrow := borrowClassInfo.BorrowDetails[len(borrowClassInfo.BorrowDetails)-1]

	thisNFT, found := k.nftKeeper.GetNFT(ctx, classID, nftID)
	if !found {
		return coserrors.Wrapf(types.ErrDepositorNotFound, "the given nft %v cannot ben found in storage", nftID)
	}
	var interestData types.NftInfo
	err = proto.Unmarshal(thisNFT.Data.Value, &interestData)
	if err != nil {
		panic(err)
	}
	thisNFTBorrowed := interestData.Borrowed

	newBorrowAmount := lastBorrow.BorrowedAmount.Sub(thisNFTBorrowed)
	newBorrow := types.BorrowDetail{BorrowedAmount: newBorrowAmount, TimeStamp: ctx.BlockTime(), ExchangeRatio: exchangeRatio}
	borrowClassInfo.BorrowDetails = append(borrowClassInfo.BorrowDetails, newBorrow)

	data, err := types2.NewAnyWithValue(&borrowClassInfo)
	if err != nil {
		panic("should never fail")
	}
	thisClass.Data = data
	err = k.nftKeeper.UpdateClass(ctx, thisClass)
	if err != nil {
		return coserrors.Wrapf(err, "fail to update the class")
	}
	if burnNFT {
		err = k.nftKeeper.Burn(ctx, classID, nftID)
		if err != nil {
			return coserrors.Wrapf(err, "fail to burn the nft")
		}
	}
	return nil
}

func (k Keeper) processEachWithdrawReq(ctx sdk.Context, depositor types.DepositorInfo, burnNFT bool, exchangeRatio sdk.Dec) error {
	for _, el := range depositor.LinkedNFT {
		ids := strings.Split(el, ":")
		err := k.updateClassAndBurnNFT(ctx, ids[0], ids[1], burnNFT, exchangeRatio)
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func (k Keeper) HandlePartialPrincipalPayment(ctx sdk.Context, poolInfo *types.PoolInfo, withdrawAccounts []sdk.AccAddress) bool {
	if len(withdrawAccounts) == 0 {
		return true
	}
	token := poolInfo.EscrowPrincipalAmount

	exchangeRatio := poolInfo.PrincipalPaymentExchangeRatio
	usdWithdrawalTotal := exchangeRatio.MulInt(poolInfo.WithdrawProposalAmount.Amount).TruncateInt()

	if token.Amount.LT(usdWithdrawalTotal) {
		ctx.Logger().Info("not enough escrow account balance to pay withdrawal proposal amount")
		if ctx.BlockTime().After(poolInfo.ProjectDueTime.Add(time.Second * time.Duration(poolInfo.WithdrawRequestWindowSeconds))) {
			poolInfo.PoolStatus = types.PoolInfo_Liquidation
		}
		return false
	}

	total := sdk.ZeroInt()
	for _, el := range withdrawAccounts[1:] {
		depositor, found := k.GetDepositor(ctx, poolInfo.Index, el)
		if !found {
			panic("should never fail to find the depositor")
		}

		interest, err := calculateTotalInterest(ctx, depositor.LinkedNFT, k.nftKeeper, true)
		if err != nil {
			panic(err)
		}

		err = k.processEachWithdrawReq(ctx, depositor, false, exchangeRatio)
		if err != nil {
			ctx.Logger().Error("fail to pay partial principal", err.Error())
			panic(err)
		}

		depositor.PendingInterest = depositor.PendingInterest.AddAmount(interest)
		usdWithdrawal := exchangeRatio.MulInt(depositor.LockedAmount.Amount).TruncateInt()
		total = total.Add(usdWithdrawal)
		depositor.WithdrawalAmount = depositor.WithdrawalAmount.AddAmount(usdWithdrawal)
		depositor.LockedAmount = sdk.NewCoin(depositor.LockedAmount.Denom, sdk.ZeroInt())
		depositor.DepositType = types.DepositorInfo_deposit_close
		k.SetDepositor(ctx, depositor)
	}
	// now we process the first investor
	depositor, found := k.GetDepositor(ctx, poolInfo.Index, withdrawAccounts[0])
	if !found {
		panic("should never fail to find the depositor")
	}

	interest, err := calculateTotalInterest(ctx, depositor.LinkedNFT, k.nftKeeper, true)
	if err != nil {
		panic(err)
	}

	err = k.processEachWithdrawReq(ctx, depositor, false, exchangeRatio)
	if err != nil {
		ctx.Logger().Error("fail to pay partial principal", err.Error())
		panic(err)
	}

	depositor.PendingInterest = depositor.PendingInterest.AddAmount(interest)
	usdAmount := usdWithdrawalTotal.Sub(total)
	depositor.WithdrawalAmount = depositor.WithdrawalAmount.AddAmount(usdAmount)
	depositor.LockedAmount = sdk.NewCoin(depositor.LockedAmount.Denom, sdk.ZeroInt())
	depositor.DepositType = types.DepositorInfo_deposit_close
	k.SetDepositor(ctx, depositor)

	poolInfo.BorrowedAmount = poolInfo.BorrowedAmount.Sub(poolInfo.WithdrawProposalAmount)
	returnToSPV := poolInfo.EscrowPrincipalAmount.SubAmount(usdWithdrawalTotal)
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccount, poolInfo.OwnerAddress, sdk.NewCoins(returnToSPV))
	if err != nil {
		ctx.Logger().Error("fail to send the leftover back to spv ", "err=", err.Error())
		panic(err)
	}

	poolInfo.EscrowPrincipalAmount = sdk.NewCoin(poolInfo.EscrowPrincipalAmount.Denom, sdk.ZeroInt())
	poolInfo.WithdrawProposalAmount = sdk.NewCoin(poolInfo.WithdrawProposalAmount.Denom, sdk.ZeroInt())
	poolInfo.WithdrawAccounts = make([]sdk.AccAddress, 0, 200)
	// clear the flag
	poolInfo.PrincipalPaid = false
	return true

}

func (k Keeper) HandlePrincipalPayment(ctx sdk.Context, poolInfo *types.PoolInfo) bool {
	//fixme this means the pool is empty
	if poolInfo.BorrowedAmount.IsZero() && poolInfo.UsableAmount.IsZero() {
		k.SetHistoryPool(ctx, *poolInfo)
		k.DelPool(ctx, poolInfo.Index)
	}

	borrowedUSD := outboundConvertToUSD(poolInfo.BorrowedAmount.Amount, poolInfo.PrincipalPaymentExchangeRatio)

	if poolInfo.EscrowPrincipalAmount.Amount.LT(borrowedUSD) {
		return false
	}

	totalPaid := poolInfo.EscrowPrincipalAmount.SubAmount(borrowedUSD)
	totalPaid = totalPaid.AddAmount(poolInfo.EscrowInterestAmount)
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccount, poolInfo.OwnerAddress, sdk.NewCoins(totalPaid))
	if err != nil {
		ctx.Logger().Error("fail to send the leftover back to spv ", "err=", err.Error())
		panic(err)
	}

	poolInfo.PoolStatus = types.PoolInfo_FROZEN
	poolInfo.EscrowPrincipalAmount = sdk.NewCoin(poolInfo.BorrowedAmount.Denom, sdk.ZeroInt())
	poolInfo.EscrowInterestAmount = sdk.ZeroInt()
	poolInfo.PrincipalPaid = false
	k.SetPool(ctx, *poolInfo)
	return true
}
