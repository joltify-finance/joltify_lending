package keeper

import (
	coserrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	types2 "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"strings"
	"time"
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
		data := thisNFT.Data.GetCachedValue()
		interestData, ok := data.(types.NftInfo)
		if !ok {
			panic("not the borrow interest type")
		}
		lendsRatio[el] = interestData.Ratio

		borrowClass, found := nftKeeper.GetClass(ctx, ids[0])
		if !found {
			panic("it should never fail to find the class")
		}

		v := borrowClass.GetData().GetCachedValue()
		borrowClassInfo, ok := v.(types.BorrowInterest)
		if !ok {
			panic("not the class type")
		}

		allPayments := borrowClassInfo.Payments
		latestTimeStamp := time.Time{}
		for _, eachPayment := range allPayments {
			if interestData.LastPayment.Sub(eachPayment.PaymentTime) > time.Second {
				continue
			}
			paymentAmount := eachPayment.PaymentAmount
			interest := sdk.NewDecFromInt(paymentAmount.Amount).Mul(sdk.NewDec(1).Sub(reserve)).Mul(interestData.Ratio).TruncateInt()
			totalInterest = totalInterest.Add(interest)
			latestTimeStamp = eachPayment.PaymentTime
		}
		if updateNFT {
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