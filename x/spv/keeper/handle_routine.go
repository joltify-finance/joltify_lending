package keeper

import (
	coserrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	types2 "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/gogo/protobuf/proto"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k Keeper) HandleInterest(ctx sdk.Context, poolInfo *types.PoolInfo) error {
	totalAmountDue, err := k.getAllInterestToBePaid(ctx, poolInfo)
	if err != nil {
		ctx.Logger().Info(err.Error())
		return err
	}

	if poolInfo.EscrowInterestAmount.Amount.LT(totalAmountDue) {
		ctx.Logger().Error("insufficient fund to pay the interest %v<%v", poolInfo.EscrowInterestAmount.String(), totalAmountDue.String())
		return types.ErrInsufficientFund
	}

	// finally, we update the poolinfo
	currentTimeTruncated := ctx.BlockTime().Truncate(time.Duration(poolInfo.PayFreq) * time.Second)
	poolInfo.LastPaymentTime = currentTimeTruncated
	if poolInfo.BorrowedAmount.Equal(sdk.ZeroInt()) {
		poolInfo.PoolStatus = types.PoolInfo_CLOSED
	}

	poolInfo.EscrowInterestAmount = poolInfo.EscrowInterestAmount.SubAmount(totalAmountDue)

	k.SetPool(ctx, *poolInfo)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRepayInterest,
			sdk.NewAttribute("amount", totalAmountDue.String()),
		),
	)
	return nil
}

func (k Keeper) updateClassAndBurnNFT(ctx sdk.Context, classID, nftID string, thisBorrow sdkmath.Int) (sdkmath.Int, error) {
	thisClass, found := k.nftKeeper.GetClass(ctx, classID)
	if !found {
		return sdkmath.ZeroInt(), coserrors.Wrapf(types.ErrClassNotFound, "the class cannot be found")
	}

	var borrowClassInfo types.BorrowInterest
	err := proto.Unmarshal(thisClass.Data.Value, &borrowClassInfo)
	if err != nil {
		panic(err)
	}
	borrowed := borrowClassInfo.Borrowed

	thisNFT, found := k.nftKeeper.GetNFT(ctx, classID, nftID)
	if !found {
		return sdkmath.ZeroInt(), coserrors.Wrapf(types.ErrDepositorNotFound, "the given nft %v cannot ben found in storage", nftID)
	}
	var interestData types.NftInfo
	err = proto.Unmarshal(thisNFT.Data.Value, &interestData)
	if err != nil {
		panic(err)
	}
	thisNFTBorrowed := thisBorrow
	if thisNFTBorrowed.IsZero() {
		thisNFTBorrowed = sdk.NewDecFromInt(borrowed.Amount).Mul(interestData.Ratio).TruncateInt()
	}
	borrowClassInfo.Borrowed = borrowClassInfo.Borrowed.SubAmount(thisNFTBorrowed)
	data, err := types2.NewAnyWithValue(&borrowClassInfo)
	if err != nil {
		panic("should never fail")
	}
	thisClass.Data = data
	err = k.nftKeeper.UpdateClass(ctx, thisClass)
	if err != nil {
		return sdkmath.ZeroInt(), coserrors.Wrapf(err, "fail to update the class")
	}
	err = k.nftKeeper.Burn(ctx, classID, nftID)
	if err != nil {
		return sdkmath.ZeroInt(), coserrors.Wrapf(err, "fail to burn the nft")
	}

	return thisNFTBorrowed, nil

}

func (k Keeper) processEachWithdrawReq(ctx sdk.Context, depositor types.DepositorInfo) error {
	totalBorrowed := sdk.ZeroInt()
	// we will handle the first borrow later
	for _, el := range depositor.LinkedNFT[1:] {
		ids := strings.Split(el, ":")
		thisBorrow, err := k.updateClassAndBurnNFT(ctx, ids[0], ids[1], sdkmath.ZeroInt())
		if err != nil {
			panic(err)
		}
		totalBorrowed = totalBorrowed.Add(thisBorrow)
	}
	// since we check whether this investor has borrow at the begin, we skip it here
	firstNFT := depositor.LinkedNFT[0]
	ids := strings.Split(firstNFT, ":")
	_, err := k.updateClassAndBurnNFT(ctx, ids[0], ids[1], depositor.LockedAmount.SubAmount(totalBorrowed).Amount)
	if err != nil {
		panic(err)
	}
	k.DelDepositor(ctx, depositor.PoolIndex, depositor.DepositorAddress)
	return nil
}

func (k Keeper) HandlePartialPrincipalPayment(ctx sdk.Context, poolInfo *types.PoolInfo) {
	var err error
	token := poolInfo.EscrowPrincipalAmount
	if token.IsLT(poolInfo.WithdrawProposalAmount) {
		ctx.Logger().Info("not enough escrow account balance to pay withdrawal proposal amount")
		return
	}

	totalPaidAmount := sdkmath.ZeroInt()
	for _, el := range poolInfo.WithdrawAccounts {
		depositor, found := k.GetDepositor(ctx, poolInfo.Index, el)
		if !found {
			panic("should never fail to find the depositor")
		}

		err = k.processEachWithdrawReq(ctx, depositor)
		if err != nil {
			ctx.Logger().Error("fail to pay partial principal", err.Error())
			panic(err)
		}

		needToBePaid := depositor.LockedAmount
		totalPaidAmount = totalPaidAmount.Add(needToBePaid.Amount)
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccount, depositor.DepositorAddress, sdk.NewCoins(needToBePaid))
		if err != nil {
			ctx.Logger().Error(err.Error(), "fail to transfer the principal to ", depositor.DepositorAddress.String())
			panic(err)
		}
	}
	if !poolInfo.WithdrawProposalAmount.Equal(totalPaidAmount) {
		panic("withdrawble amount is not equal as the total paid!")
	}

	poolInfo.WithdrawProposalAmount = sdk.NewCoin(poolInfo.WithdrawProposalAmount.Denom, sdk.ZeroInt())
	poolInfo.WithdrawAccounts = make([]sdk.AccAddress, 0, 200)
	k.SetPool(ctx, *poolInfo)
	return

}

func (k Keeper) HandlePrincipalPayment(ctx sdk.Context, poolInfo *types.PoolInfo) {
	token := poolInfo.EscrowPrincipalAmount
	if token.Amount.Equal(sdk.ZeroInt()) {
		return
	}
	aboutToPay := sdk.NewCoin(token.Denom, token.Amount)
	if token.IsGTE(poolInfo.BorrowedAmount) {
		aboutToPay = poolInfo.BorrowedAmount
	}
	k.travelThoughPrincipalToBePaid(ctx, poolInfo, aboutToPay)
	// now we query all the borrows

	poolInfo.BorrowedAmount = poolInfo.BorrowedAmount.Sub(aboutToPay)
	poolInfo.BorrowableAmount = poolInfo.BorrowableAmount.Add(aboutToPay)

	// once the pool borrowed is 0, we will deactive the pool
	if poolInfo.BorrowedAmount.Amount.Equal(sdk.ZeroInt()) {
		poolInfo.PoolStatus = types.PoolInfo_CLOSING
	}
	poolInfo.EscrowPrincipalAmount = poolInfo.EscrowPrincipalAmount.Sub(aboutToPay)

	k.SetPool(ctx, *poolInfo)

}
