package keeper

import (
	"strings"
	"time"

	coserrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	types2 "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func calculateTotalInterest(ctx sdk.Context, lendNFTs []string, nftKeeper types.NFTKeeper, reserve sdk.Dec, updateNFT bool) (sdkmath.Int, error) {

	lendsRatio := make(map[string]sdk.Dec)
	totalInterest := sdk.NewInt(0)
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

		lendsRatio[el] = interestData.Ratio

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
		for _, eachPayment := range allPayments {
			// if the latest payment  this spv has is smaller than the spv that paied to all the investor, we claim the interest
			if eachPayment.PaymentTime.Before(interestData.LastPayment) || eachPayment.PaymentTime.Equal(interestData.LastPayment) {
				continue
			}
			if eachPayment.PaymentAmount.Amount.IsZero() {
				continue
			}
			paymentAmount := eachPayment.PaymentAmount
			interest := sdk.NewDecFromInt(paymentAmount.Amount).Mul(interestData.Ratio).TruncateInt()
			totalInterest = totalInterest.Add(interest)
			latestTimeStamp = eachPayment.PaymentTime
			lastPaymentSet = true
		}
		if updateNFT && lastPaymentSet {
			interestData.LastPayment = latestTimeStamp
			data, err := types2.NewAnyWithValue(&interestData)
			if err != nil {
				panic("pack class any data failed")
			}
			thisNFT.Data = data
			nftKeeper.Update(ctx, thisNFT)
		}
	}
	return totalInterest, nil
}

func calculateTotalOutstandingInterest(ctx sdk.Context, lendNFTs []string, nftKeeper types.NFTKeeper, reserve sdk.Dec) (sdkmath.Int, error) {

	lendsRatio := make(map[string]sdk.Dec)
	totalInterest := sdk.NewInt(0)
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

		lendsRatio[el] = interestData.Ratio

		borrowClass, found := nftKeeper.GetClass(ctx, ids[0])
		if !found {
			panic("it should never fail to find the class")
		}

		var borrowClassInfo types.BorrowInterest
		err = proto.Unmarshal(borrowClass.Data.Value, &borrowClassInfo)
		if err != nil {
			panic(err)
		}

		lastPayment := borrowClassInfo.Payments[len(borrowClassInfo.Payments)-1]
		delta := uint64(ctx.BlockTime().Sub(lastPayment.PaymentTime).Seconds())
		factor := CalculateInterestFactor(borrowClassInfo.InterestSPY, sdk.NewIntFromUint64(delta))
		paymentAmountToInvestor := sdk.NewDecFromInt(borrowClassInfo.Borrowed.Amount).Mul(sdk.OneDec().Sub(reserve))
		interest := paymentAmountToInvestor.Mul(factor.Sub(sdk.OneDec())).TruncateInt()
		totalInterest = totalInterest.Add(interest)
	}
	return totalInterest, nil
}

func calculateTotalPrinciple(ctx sdk.Context, lendNFTs []string, nftKeeper types.NFTKeeper) (sdkmath.Int, error) {
	totalBorrowed := sdk.NewInt(0)
	for _, el := range lendNFTs {
		ids := strings.Split(el, ":")

		thisClass, found := nftKeeper.GetClass(ctx, ids[0])
		if !found {
			return sdkmath.Int{}, coserrors.Wrapf(types.ErrClassNotFound, "the class cannot be found")
		}

		var borrowClassInfo types.BorrowInterest
		err := proto.Unmarshal(thisClass.Data.Value, &borrowClassInfo)
		if err != nil {
			panic(err)
		}
		borrowed := borrowClassInfo.Borrowed

		thisNFT, found := nftKeeper.GetNFT(ctx, ids[0], ids[1])
		if !found {
			return sdkmath.Int{}, coserrors.Wrapf(types.ErrDepositorNotFound, "the given nft %v cannot ben found in storage", ids[1])
		}
		var interestData types.NftInfo
		err = proto.Unmarshal(thisNFT.Data.Value, &interestData)
		if err != nil {
			panic(err)
		}

		thisNFTBorrowed := sdk.NewDecFromInt(borrowed.Amount).Mul(interestData.Ratio).TruncateInt()
		totalBorrowed = totalBorrowed.Add(thisNFTBorrowed)
	}
	return totalBorrowed, nil
}
