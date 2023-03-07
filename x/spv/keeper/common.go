package keeper

import (
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

func (k Keeper) doBorrow(ctx sdk.Context, poolInfo types.PoolInfo, tokenAmount sdk.Coin) error {
	// create the new nft class for this borrow event
	classID := fmt.Sprintf("nft-%v", poolInfo.Index[2:])
	poolClass, found := k.nftKeeper.GetClass(ctx, classID)
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
	paymentTime := ctx.BlockTime()
	firstPayment := types.PaymentItem{PaymentTime: paymentTime, PaymentAmount: sdk.NewCoin(tokenAmount.Denom, sdk.NewInt(0))}
	bi := types.BorrowInterest{
		PoolIndex:    poolInfo.Index,
		Apy:          poolInfo.Apy,
		PayFreq:      poolInfo.PayFreq,
		IssueTime:    ctx.BlockTime(),
		Borrowed:     tokenAmount,
		BorrowedLast: tokenAmount,
		MonthlyRatio: i,
		InterestSPY:  rate,
		Payments:     []*types.PaymentItem{&firstPayment},
	}

	data, err := types2.NewAnyWithValue(&bi)
	if err != nil {
		panic(err)
	}
	currentBorrowClass.Data = data
	err = k.nftKeeper.SaveClass(ctx, currentBorrowClass)
	if err != nil {
		return err
	}

	// update the borrow series
	poolInfo.PoolNFTIds = append(poolInfo.PoolNFTIds, currentBorrowClass.Id)

	// we start the project
	if len(poolInfo.PoolNFTIds) == 1 {
		poolInfo.ProjectDueTime = ctx.BlockTime().Add(time.Second * time.Duration(poolInfo.ProjectLength))
	}

	err = k.processBorrow(ctx, &poolInfo, currentBorrowClass, tokenAmount)
	if err != nil {
		return err
	}

	// we finally update the pool info
	poolInfo.PoolStatus = types.PoolInfo_ACTIVE
	poolInfo.LastPaymentTime = paymentTime
	k.SetPool(ctx, poolInfo)
	return nil
}

func (k Keeper) processBorrow(ctx sdk.Context, poolInfo *types.PoolInfo, nftClass nfttypes.Class, amount sdk.Coin) error {
	if poolInfo.BorrowableAmount.IsLT(amount) {
		return types.ErrInsufficientFund
	}
	utilization := sdk.NewDecFromInt(amount.Amount).Quo(sdk.NewDecFromInt(poolInfo.BorrowableAmount.Amount))

	// we transfer the fund from the module to the spv
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccount, poolInfo.OwnerAddress, sdk.NewCoins(amount))
	if err != nil {
		return err
	}

	// we add the amount of the tokens that borrowed in the pool and decreases the borrowable
	poolInfo.BorrowedAmount = poolInfo.BorrowedAmount.Add(amount)
	poolInfo.BorrowableAmount, err = poolInfo.BorrowableAmount.SafeSub(amount)
	if err != nil {
		return types.ErrInsufficientFund
	}

	// we update each investor leftover
	k.processInvestors(ctx, poolInfo, utilization, amount.Amount, nftClass)
	return nil
}

func (k Keeper) doProcessInvestor(ctx sdk.Context, depositor *types.DepositorInfo, locked, totalBorrow sdkmath.Int, nftTemplate nfttypes.NFT, nftClassId string, poolInfo *types.PoolInfo, errGlobal error) {
	depositor.LockedAmount = depositor.LockedAmount.Add(sdk.NewCoin(depositor.WithdrawalAmount.Denom, locked))
	depositor.WithdrawalAmount = depositor.WithdrawalAmount.SubAmount(locked)
	borrowRatio := sdk.NewDecFromInt(locked).Quo(sdk.NewDecFromInt(totalBorrow))

	// nft ID is the hash(nft class ID, investorWallet)
	indexHash := crypto.Keccak256Hash([]byte(nftClassId), depositor.DepositorAddress)
	nftTemplate.Id = fmt.Sprintf("invoice-%v", indexHash.String()[2:])

	userData := types.NftInfo{Issuer: poolInfo.PoolName, Receiver: depositor.DepositorAddress.String(), Ratio: borrowRatio, LastPayment: ctx.BlockTime()}
	data, err := types2.NewAnyWithValue(&userData)
	if err != nil {
		panic("should never fail")
	}
	nftTemplate.Data = data
	err = k.nftKeeper.Mint(ctx, nftTemplate, depositor.DepositorAddress)
	if err != nil {
		errGlobal = types.ErrMINTNFT
		return
	}

	classIDAndNFTID := fmt.Sprintf("%v:%v", nftTemplate.ClassId, nftTemplate.Id)
	depositor.LinkedNFT = append(depositor.LinkedNFT, classIDAndNFTID)
	k.SetDepositor(ctx, *depositor)

}

func (k Keeper) processInvestors(ctx sdk.Context, poolInfo *types.PoolInfo, utilization sdk.Dec, totalBorrow sdkmath.Int, nftClass nfttypes.Class) error {

	nftTemplate := nfttypes.NFT{
		ClassId: nftClass.Id,
		Uri:     nftClass.Uri,
		UriHash: nftClass.UriHash,
	}

	// now we update the depositor's withdrawal amount and locked amount
	var errGlobal error
	var firstDepositor *types.DepositorInfo
	totalLocked := sdk.ZeroInt()
	k.IterateDepositors(ctx, poolInfo.Index, func(depositor types.DepositorInfo) (stop bool) {
		if depositor.WithdrawProposal || depositor.TransferRequest {
			// since you have submitted the withdrawal/transfer request, we skipp the borrow from it
			return false
		}

		if firstDepositor == nil {
			firstDepositor = &depositor
			return false
		}
		locked := sdk.NewDecFromInt(depositor.WithdrawalAmount.Amount).Mul(utilization).TruncateInt()
		k.doProcessInvestor(ctx, &depositor, locked, totalBorrow, nftTemplate, nftClass.Id, poolInfo, errGlobal)
		totalLocked = totalLocked.Add(locked)
		return false
	})

	// now we process the last one
	locked := totalBorrow.Sub(totalLocked)
	k.doProcessInvestor(ctx, firstDepositor, locked, totalBorrow, nftTemplate, nftClass.Id, poolInfo, errGlobal)

	return errGlobal
}
