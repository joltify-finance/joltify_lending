package keeper

import (
	coserrors "cosmossdk.io/errors"
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

func (k Keeper) processEachWithdrawReq(ctx sdk.Context, depositor types.DepositorInfo) error {
	// we firstly update the class borrowed amount

	totalBorrowed := sdk.NewInt(0)
	for _, el := range depositor.LinkedNFT {
		ids := strings.Split(el, ":")

		thisClass, found := k.nftKeeper.GetClass(ctx, ids[0])
		if !found {
			return coserrors.Wrapf(types.ErrClassNotFound, "the class cannot be found")
		}

		var borrowClassInfo types.BorrowInterest
		err := proto.Unmarshal(thisClass.Data.Value, &borrowClassInfo)
		if err != nil {
			panic(err)
		}
		borrowed := borrowClassInfo.Borrowed

		thisNFT, found := k.nftKeeper.GetNFT(ctx, ids[0], ids[1])
		if !found {
			return coserrors.Wrapf(types.ErrDepositorNotFound, "the given nft %v cannot ben found in storage", ids[1])
		}
		var interestData types.NftInfo
		err = proto.Unmarshal(thisNFT.Data.Value, &interestData)
		if err != nil {
			panic(err)
		}
		thisNFTBorrowed := sdk.NewDecFromInt(borrowed.Amount).Mul(interestData.Ratio).TruncateInt()
		borrowed = borrowed.SubAmount(thisNFTBorrowed)

	}

}

func (k Keeper) HandlePartialPrincipalPayment(ctx sdk.Context, poolInfo *types.PoolInfo) {

	token := poolInfo.EscrowPrincipalAmount
	if token.IsLT(poolInfo.WithdrawProposalAmount) {
		ctx.Logger().Info("not enough escrow account balance to pay withdrawal proposal amount")
		return
	}

	for _, el := range poolInfo.WithdrawAccounts {
		depositor, found := k.GetDepositor(ctx, poolInfo.Index, el)
		if !found {
			panic("should never fail to find the depositor")
		}

		k.processEachWithdrawReq(depositor)

	}

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
