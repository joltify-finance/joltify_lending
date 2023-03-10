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

	if poolInfo.EscrowInterestAmount.Amount.LT(totalAmountDue) {
		ctx.Logger().Error("insufficient fund to pay the interest %v<%v", poolInfo.EscrowInterestAmount.String(), totalAmountDue.String())
		return types.ErrInsufficientFund
	}

	poolInfo.EscrowInterestAmount = poolInfo.EscrowInterestAmount.SubAmount(totalAmountDue)

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
func (k Keeper) HandleTransfer(ctx sdk.Context, poolInfo *types.PoolInfo) {
	var err error
	var depositors []*types.DepositorInfo
	totalLockedAmount := sdkmath.ZeroInt()
	if len(poolInfo.TransferAccounts) == 0 {
		return
	}
	for _, el := range poolInfo.TransferAccounts {
		d, found := k.GetDepositor(ctx, poolInfo.Index, el)
		if !found {
			panic("should never fail to find the depositor")
		}
		depositors = append(depositors, &d)
		totalLockedAmount = totalLockedAmount.Add(d.LockedAmount.Amount)
	}
	poolInfo.TransferAccounts = []sdk.AccAddress{}

	// borrowable is larger than the total required, so we can return the money to these investors
	if poolInfo.BorrowableAmount.Amount.GTE(totalLockedAmount) {
		for _, el := range depositors {
			interest, err := calculateTotalInterest(ctx, el.LinkedNFT, k.nftKeeper, true)
			if err != nil {
				panic(err)
			}

			err = k.processEachWithdrawReq(ctx, *el)
			if err != nil {
				panic(err)
			}

			el.LinkedNFT = []string{}
			el.WithdrawalAmount = el.WithdrawalAmount.Add(el.LockedAmount.AddAmount(interest))
			el.LockedAmount = sdk.NewCoin(el.LockedAmount.Denom, sdk.ZeroInt())
			el.DepositType = types.DepositorInfo_deposit_close
			k.SetDepositor(ctx, *el)
		}

		poolInfo.BorrowedAmount = poolInfo.BorrowedAmount.SubAmount(totalLockedAmount)
		err := k.doBorrow(ctx, poolInfo, sdk.NewCoin(poolInfo.BorrowableAmount.Denom, totalLockedAmount), false, nil)
		if err != nil {
			panic(err)
		}
		k.SetPool(ctx, *poolInfo)
		return
	}
	// now we process the partial transfer
	for i, el := range depositors {

		interest, err := calculateTotalInterest(ctx, el.LinkedNFT, k.nftKeeper, true)
		if err != nil {
			panic(err)
		}

		err = k.processEachWithdrawReq(ctx, *el)
		if err != nil {
			panic(err)
		}
		depositors[i].DepositType = types.DepositorInfo_unset
		el.LinkedNFT = []string{}
		el.WithdrawalAmount = el.WithdrawalAmount.Add(el.LockedAmount.AddAmount(interest))
		el.LockedAmount = sdk.NewCoin(el.LockedAmount.Denom, sdk.ZeroInt())
	}
	borrowedFromPreviousInvestors := totalLockedAmount.Sub(poolInfo.BorrowableAmount.Amount)
	poolInfo.BorrowedAmount = poolInfo.BorrowedAmount.SubAmount(totalLockedAmount)
	poolInfo.BorrowableAmount = poolInfo.BorrowableAmount.AddAmount(borrowedFromPreviousInvestors)
	err = k.doBorrow(ctx, poolInfo, sdk.NewCoin(poolInfo.BorrowableAmount.Denom, borrowedFromPreviousInvestors), false, depositors)
	if err != nil {
		panic(err)
	}

	// now we process the partial transfer
	for i := range depositors {
		depositors[i].DepositType = types.DepositorInfo_processed
		k.SetDepositor(ctx, *depositors[i])
	}

	err = k.doBorrow(ctx, poolInfo, poolInfo.BorrowableAmount, false, nil)
	if err != nil {
		panic(err)
	}

}

func (k Keeper) updateClassAndBurnNFT(ctx sdk.Context, classID, nftID string) error {
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
	newBorrow := types.BorrowDetail{BorrowedAmount: newBorrowAmount, TimeStamp: ctx.BlockTime()}
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
	err = k.nftKeeper.Burn(ctx, classID, nftID)
	if err != nil {
		return coserrors.Wrapf(err, "fail to burn the nft")
	}

	return nil
}

func (k Keeper) processEachWithdrawReq(ctx sdk.Context, depositor types.DepositorInfo) error {
	for _, el := range depositor.LinkedNFT[1:] {
		ids := strings.Split(el, ":")
		err := k.updateClassAndBurnNFT(ctx, ids[0], ids[1])
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func (k Keeper) HandlePartialPrincipalPayment(ctx sdk.Context, poolInfo *types.PoolInfo, withdrawAccounts []sdk.AccAddress) {
	token := poolInfo.EscrowPrincipalAmount
	if token.IsLT(poolInfo.WithdrawProposalAmount) {
		ctx.Logger().Info("not enough escrow account balance to pay withdrawal proposal amount")
		return
	}

	if len(withdrawAccounts) == 0 {
		return
	}

	totalPaidAmount := sdkmath.ZeroInt()
	for _, el := range withdrawAccounts {
		depositor, found := k.GetDepositor(ctx, poolInfo.Index, el)
		if !found {
			panic("should never fail to find the depositor")
		}

		interest, err := calculateTotalInterest(ctx, depositor.LinkedNFT, k.nftKeeper, true)
		if err != nil {
			panic(err)
		}

		err = k.processEachWithdrawReq(ctx, depositor)
		if err != nil {
			ctx.Logger().Error("fail to pay partial principal", err.Error())
			panic(err)
		}

		totalPaidAmount = totalPaidAmount.Add(depositor.LockedAmount.Amount.Add(interest))
		depositor.LinkedNFT = []string{}
		depositor.WithdrawalAmount = depositor.WithdrawalAmount.Add(depositor.LockedAmount)
		depositor.LockedAmount = sdk.NewCoin(depositor.LockedAmount.Denom, sdk.ZeroInt())
		depositor.DepositType = types.DepositorInfo_deposit_close
		k.SetDepositor(ctx, depositor)
	}
	if !poolInfo.WithdrawProposalAmount.Equal(totalPaidAmount) {
		panic("withdrawble amount is not equal as the total paid!")
	}

	if poolInfo.BorrowedAmount.IsLTE(poolInfo.WithdrawProposalAmount) {
		poolInfo.PoolStatus = types.PoolInfo_CLOSING
		ctx.Logger().Info(" the pool", "pool_ID:", poolInfo.Index)
		poolInfo.BorrowedAmount = sdk.NewCoin(poolInfo.BorrowableAmount.Denom, sdkmath.ZeroInt())
	} else {
		poolInfo.BorrowedAmount = poolInfo.BorrowedAmount.Sub(poolInfo.WithdrawProposalAmount)
	}
	poolInfo.WithdrawProposalAmount = sdk.NewCoin(poolInfo.WithdrawProposalAmount.Denom, sdk.ZeroInt())
	poolInfo.WithdrawAccounts = make([]sdk.AccAddress, 0, 200)

	k.SetPool(ctx, *poolInfo)
	return

}

// not supported yet
func (k Keeper) HandlePrincipalPayment(ctx sdk.Context, poolInfo *types.PoolInfo) {
	if poolInfo.BorrowedAmount.IsZero() {
		k.SetHistoryPool(ctx, *poolInfo)
		k.DelPool(ctx, poolInfo.Index)
	}
	token := poolInfo.EscrowPrincipalAmount
	if token.Amount.LT(poolInfo.BorrowedAmount.Amount) {
		ctx.Logger().Error("not enough money to pay the total principal")
		return
	}

	poolInfo.EscrowPrincipalAmount = poolInfo.EscrowPrincipalAmount.Sub(poolInfo.BorrowedAmount)

	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccount, poolInfo.OwnerAddress, sdk.NewCoins(poolInfo.EscrowInterestAmount))
	if err != nil {
		ctx.Logger().Error("fail to send the leftover back to spv ", "err=", err.Error())
		panic(err)
	}

	poolInfo.PoolStatus = types.PoolInfo_CLOSED
	poolInfo.EscrowPrincipalAmount = sdk.NewCoin(poolInfo.BorrowedAmount.Denom, sdk.ZeroInt())
	poolInfo.EscrowInterestAmount = sdk.NewCoin(poolInfo.BorrowedAmount.Denom, sdk.ZeroInt())
	k.SetPool(ctx, *poolInfo)
}
