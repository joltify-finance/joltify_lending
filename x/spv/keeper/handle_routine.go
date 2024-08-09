package keeper

import (
	"context"
	"errors"
	"strings"
	"time"

	coserrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	types2 "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/gogo/protobuf/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k Keeper) HandleInterest(rctx context.Context, poolInfo *types.PoolInfo) error {
	ctx := sdk.UnwrapSDKContext(rctx)
	totalAmountDue, poolLatestPaymentTime, err := k.getAllInterestToBePaid(ctx, poolInfo)
	if err != nil {
		return err
	}

	if totalAmountDue.IsZero() {
		// no interest to be paid
		return errors.New("no interest to be paid")
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
	poolInfo.LastPaymentTime = poolLatestPaymentTime
	k.SetPool(ctx, *poolInfo)

	k.AfterSPVInterestPaid(ctx, poolInfo.Index, totalAmountDue)

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
func (k Keeper) HandleTransfer(rctx context.Context, poolInfo *types.PoolInfo) bool {
	ctx := sdk.UnwrapSDKContext(rctx)
	var err error
	var depositors []*types.DepositorInfo
	totalLockedAmount := sdkmath.ZeroInt()
	if len(poolInfo.TransferAccounts) == 0 {
		return false
	}

	// we rule out the depositors that have less than given min deposit amount
	k.updateDepositorStatus(ctx, poolInfo)
	for _, el := range poolInfo.TransferAccounts {
		d, found := k.GetDepositor(ctx, poolInfo.Index, el)
		if !found {
			// TODO this is temporary solution, will be removed in the future
			ctx.Logger().Error("deposit not found", "addresses", el.String())
			continue
		}
		depositors = append(depositors, &d)
		totalLockedAmount = totalLockedAmount.Add(d.LockedAmount.Amount)
	}

	poolInfo.ProcessedTransferAccounts = append(poolInfo.ProcessedTransferAccounts, poolInfo.TransferAccounts...)
	poolInfo.TransferAccounts = make([]sdk.AccAddress, 0, 200)
	poolInfo.TotalTransferOwnershipAmount = sdk.NewCoin(poolInfo.TotalTransferOwnershipAmount.Denom, sdkmath.ZeroInt())

	a, _ := denomConvertToLocalAndUsd(poolInfo.BorrowedAmount.Denom)
	usdTotalLocked, ratio, err := k.outboundConvertToUSDWithMarketID(ctx, denomConvertToMarketID(a), totalLockedAmount)
	if err != nil {
		ctx.Logger().Error(err.Error(), "outbound convert with market ID fail to convert", poolInfo.WithdrawProposalAmount)
		return false
	}

	// borrowable is larger than the total required, so we can return the money to these investors
	if poolInfo.UsableAmount.Amount.GTE(usdTotalLocked) {
		for _, el := range depositors {
			interest, err := calculateTotalInterest(ctx, el.LinkedNFT, k.NftKeeper, true)
			if err != nil {
				panic(err)
			}

			err = k.hooks.BeforeNFTBurned(ctx, el.PoolIndex, el.DepositorAddress.String(), el.LinkedNFT)
			if err != nil {
				ctx.Logger().Error("fail to process the spv incentives before the nft burn", err.Error())
			}

			err = k.processEachWithdrawReq(ctx, *el, true, ratio)
			if err != nil {
				panic(err)
			}

			el.LinkedNFT = []string{}
			el.PendingInterest = el.PendingInterest.AddAmount(interest)

			usdLocked := outboundConvertToUSD(el.LockedAmount.Amount, ratio)
			el.WithdrawalAmount = el.WithdrawalAmount.AddAmount(usdLocked)
			el.LockedAmount = sdk.NewCoin(el.LockedAmount.Denom, sdkmath.ZeroInt())
			el.DepositType = types.DepositorInfo_deposit_close
			k.SetDepositor(ctx, *el)
		}

		poolInfo.BorrowedAmount = poolInfo.BorrowedAmount.SubAmount(totalLockedAmount)
		err := k.doBorrow(ctx, poolInfo, sdk.NewCoin(poolInfo.UsableAmount.Denom, usdTotalLocked), false, nil, sdkmath.ZeroInt(), true)
		if err != nil {
			panic(err)
		}
		return true
	}
	// now we process the partial transfer
	totalBorrowableFromPrevious := sdkmath.ZeroInt()
	for i, el := range depositors {
		interest, err := calculateTotalInterest(ctx, el.LinkedNFT, k.NftKeeper, true)
		if err != nil {
			panic(err)
		}

		err = k.hooks.BeforeNFTBurned(ctx, poolInfo.Index, el.DepositorAddress.String(), el.LinkedNFT)
		if err != nil {
			ctx.Logger().Error("fail to process the spv incentives before the nft burn", err.Error())
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
		el.LockedAmount = sdk.NewCoin(el.LockedAmount.Denom, sdkmath.ZeroInt())
		totalBorrowableFromPrevious = totalBorrowableFromPrevious.Add(el.WithdrawalAmount.Amount)
	}

	needToBorrowedFromPreviousInvestorsUsd := usdTotalLocked.Sub(poolInfo.UsableAmount.Amount)
	// we need to adjust the amount of the borrowed and borrowable for the pool as we borrow again from these investors
	poolInfo.BorrowedAmount = poolInfo.BorrowedAmount.SubAmount(totalLockedAmount)
	poolInfo.UsableAmount = poolInfo.UsableAmount.AddAmount(needToBorrowedFromPreviousInvestorsUsd)
	err = k.doBorrow(ctx, poolInfo, sdk.NewCoin(poolInfo.UsableAmount.Denom, needToBorrowedFromPreviousInvestorsUsd), false, depositors, totalBorrowableFromPrevious, true)
	if err != nil {
		panic(err)
	}

	// now we process the partial transfer
	for i := range depositors {
		depositors[i].DepositType = types.DepositorInfo_processed
		k.SetDepositor(ctx, *depositors[i])
	}
	err = k.doBorrow(ctx, poolInfo, poolInfo.UsableAmount, false, nil, sdkmath.ZeroInt(), true)
	if err != nil {
		panic(err)
	}
	return true
}

func (k Keeper) updateClassAndBurnNFT(rctx context.Context, classID, nftID string, burnNFT bool, exchangeRatio sdkmath.LegacyDec) error {
	ctx := sdk.UnwrapSDKContext(rctx)
	thisClass, found := k.NftKeeper.GetClass(ctx, classID)
	if !found {
		return coserrors.Wrapf(types.ErrClassNotFound, "the class cannot be found")
	}

	var borrowClassInfo types.BorrowInterest
	err := proto.Unmarshal(thisClass.Data.Value, &borrowClassInfo)
	if err != nil {
		panic(err)
	}
	lastBorrow := borrowClassInfo.BorrowDetails[len(borrowClassInfo.BorrowDetails)-1]

	thisNFT, found := k.NftKeeper.GetNFT(ctx, classID, nftID)
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
	err = k.NftKeeper.UpdateClass(ctx, thisClass)
	if err != nil {
		return coserrors.Wrapf(err, "fail to update the class")
	}
	if burnNFT {
		err = k.ArchiveNFT(ctx, classID, nftID)
		if err != nil {
			return coserrors.Wrapf(err, "fail to burn the nft")
		}
	}
	return nil
}

func (k Keeper) processEachWithdrawReq(ctx context.Context, depositor types.DepositorInfo, burnNFT bool, exchangeRatio sdkmath.LegacyDec) error {
	for _, el := range depositor.LinkedNFT {
		ids := strings.Split(el, ":")
		err := k.updateClassAndBurnNFT(ctx, ids[0], ids[1], burnNFT, exchangeRatio)
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func (k Keeper) HandlePartialPrincipalPayment(rctx context.Context, poolInfo *types.PoolInfo, withdrawAccounts []sdk.AccAddress) bool {
	ctx := sdk.UnwrapSDKContext(rctx)

	if len(withdrawAccounts) == 0 {
		return true
	}
	token := poolInfo.EscrowPrincipalAmount

	exchangeRatio := poolInfo.PrincipalWithdrawalRequestPaymentRatio
	if exchangeRatio.IsZero() {
		ctx.Logger().Info("exchange ratio is zero, cannot process partial principal payment")
		return false
	}
	usdWithdrawalTotal := sdk.NewCoin(poolInfo.TargetAmount.Denom, exchangeRatio.MulInt(poolInfo.WithdrawProposalAmount.Amount).TruncateInt())
	if token.IsLT(usdWithdrawalTotal) {
		ctx.Logger().Info("not enough escrow account balance to pay withdrawal proposal amount")
		if ctx.BlockTime().After(poolInfo.ProjectDueTime.Add(time.Second * time.Duration(poolInfo.WithdrawRequestWindowSeconds))) {
			poolInfo.PoolStatus = types.PoolInfo_Liquidation
		}
		return false
	}

	total := sdkmath.ZeroInt()
	for _, el := range withdrawAccounts[1:] {
		depositor, found := k.GetDepositor(ctx, poolInfo.Index, el)
		if !found {
			panic("should never fail to find the depositor")
		}

		interest, err := calculateTotalInterest(ctx, depositor.LinkedNFT, k.NftKeeper, true)
		if err != nil {
			panic(err)
		}

		err = k.hooks.BeforeNFTBurned(ctx, depositor.PoolIndex, depositor.DepositorAddress.String(), depositor.LinkedNFT)
		if err != nil {
			ctx.Logger().Error("fail to process the spv incentives before the nft burn", err.Error())
		}

		err = k.processEachWithdrawReq(ctx, depositor, true, exchangeRatio)
		if err != nil {
			ctx.Logger().Error("fail to pay partial principal", err.Error())
			panic(err)
		}
		depositor.LinkedNFT = []string{}
		depositor.PendingInterest = depositor.PendingInterest.AddAmount(interest)
		usdWithdrawal := sdkmath.LegacyNewDecFromInt(depositor.LockedAmount.Amount).Mul(exchangeRatio).TruncateInt()
		total = total.Add(usdWithdrawal)
		depositor.WithdrawalAmount = depositor.WithdrawalAmount.AddAmount(usdWithdrawal)
		depositor.LockedAmount = sdk.NewCoin(depositor.LockedAmount.Denom, sdkmath.ZeroInt())
		depositor.DepositType = types.DepositorInfo_deposit_close
		k.SetDepositor(ctx, depositor)
	}
	// now we process the first investor
	depositor, found := k.GetDepositor(ctx, poolInfo.Index, withdrawAccounts[0])
	if !found {
		panic("should never fail to find the depositor")
	}

	interest, err := calculateTotalInterest(ctx, depositor.LinkedNFT, k.NftKeeper, true)
	if err != nil {
		panic(err)
	}

	err = k.hooks.BeforeNFTBurned(ctx, depositor.PoolIndex, depositor.DepositorAddress.String(), depositor.LinkedNFT)
	if err != nil {
		ctx.Logger().Error("fail to process the spv incentives before the nft burn", err.Error())
	}
	err = k.processEachWithdrawReq(ctx, depositor, true, exchangeRatio)
	if err != nil {
		ctx.Logger().Error("fail to pay partial principal", err.Error())
		panic(err)
	}

	depositor.LinkedNFT = []string{}

	depositor.PendingInterest = depositor.PendingInterest.AddAmount(interest)
	usdAmount := usdWithdrawalTotal.SubAmount(total)
	depositor.WithdrawalAmount = depositor.WithdrawalAmount.Add(usdAmount)
	depositor.LockedAmount = sdk.NewCoin(depositor.LockedAmount.Denom, sdkmath.ZeroInt())
	depositor.DepositType = types.DepositorInfo_deposit_close
	k.SetDepositor(ctx, depositor)

	poolInfo.BorrowedAmount = poolInfo.BorrowedAmount.Sub(poolInfo.WithdrawProposalAmount)
	returnToSPV := poolInfo.EscrowPrincipalAmount.Sub(usdWithdrawalTotal)
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccount, poolInfo.OwnerAddress, sdk.NewCoins(returnToSPV))
	if err != nil {
		ctx.Logger().Error("fail to send the leftover back to spv ", "err=", err.Error())
		panic(err)
	}

	poolInfo.EscrowPrincipalAmount = sdk.NewCoin(poolInfo.EscrowPrincipalAmount.Denom, sdkmath.ZeroInt())
	poolInfo.WithdrawProposalAmount = sdk.NewCoin(poolInfo.WithdrawProposalAmount.Denom, sdkmath.ZeroInt())

	poolInfo.ProcessedWithdrawAccounts = append(poolInfo.ProcessedWithdrawAccounts, poolInfo.WithdrawAccounts...)
	poolInfo.WithdrawAccounts = make([]sdk.AccAddress, 0, 200)
	// clear the flag
	poolInfo.PrincipalPaid = false
	poolInfo.PoolStatus = types.PoolInfo_ACTIVE
	return true
}

func (k Keeper) HandlePrincipalPayment(rctx context.Context, poolInfo *types.PoolInfo) bool {
	ctx := sdk.UnwrapSDKContext(rctx)
	// fixme this means the pool is empty
	// if poolInfo.BorrowedAmount.IsZero() && poolInfo.UsableAmount.IsZero() {
	if k.isEmptyPool(ctx, *poolInfo) {
		k.ArchivePool(ctx, *poolInfo)
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
	poolInfo.EscrowPrincipalAmount = sdk.NewCoin(poolInfo.BorrowedAmount.Denom, sdkmath.ZeroInt())
	poolInfo.EscrowInterestAmount = sdkmath.ZeroInt()
	poolInfo.PrincipalPaid = false
	k.SetPool(ctx, *poolInfo)
	return true
}
