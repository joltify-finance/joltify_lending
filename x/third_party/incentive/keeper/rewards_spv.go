package keeper

import (
	"context"
	"fmt"
	"math"
	"strings"
	"time"

	coserrors "cosmossdk.io/errors"
	types3 "github.com/cosmos/cosmos-sdk/codec/types"

	types2 "github.com/joltify-finance/joltify_lending/x/spv/types"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"
)

// AccumulateSPVRewards calculates new rewards to distribute this block and updates the global indexes to reflect this.
// The provided rewardPeriod must be valid to avoid panics in calculating time durations.
func (k Keeper) AccumulateSPVRewards(ctx context.Context, rewardPeriod types.MultiRewardPeriod) {
	previousAccrualTime, found := k.GetSPVRewardAccrualTime(ctx, rewardPeriod.CollateralType)
	if !found {
		previousAccrualTime = ctx.BlockTime()
	}

	spvAccRewardTokens, found := k.GetSPVReward(ctx, rewardPeriod.CollateralType)
	if !found {
		spvAccRewardTokens = types.SPVRewardAccTokens{}
	}

	timeDelta := getTimeElapsedWithinLimits(previousAccrualTime, ctx.BlockTime(), rewardPeriod.Start, rewardPeriod.End)
	tokenIncreased := calculateTokenIncrease(timeDelta, rewardPeriod.RewardsPerSecond)
	spvAccRewardTokens.PaymentAmount = spvAccRewardTokens.PaymentAmount.Add(tokenIncreased...)

	updatedTime := minTime(ctx.BlockTime(), rewardPeriod.End)

	k.SetSPVRewardAccrualTime(ctx, rewardPeriod.CollateralType, updatedTime)
	k.SetSPVReward(ctx, rewardPeriod.CollateralType, spvAccRewardTokens)
}

// minTime returns the earliest of two times.
func minTime(t1, t2 time.Time) time.Time {
	if t2.Before(t1) {
		return t2
	}
	return t1
}

// maxTime returns the latest of two times.
func maxTime(t1, t2 time.Time) time.Time {
	if t2.After(t1) {
		return t2
	}
	return t1
}

// getTimeElapsedWithinLimits returns the duration between start and end times, capped by min and max times.
// If the start and end range is outside the min to max time range then zero duration is returned.
func getTimeElapsedWithinLimits(start, end, limitMin, limitMax time.Time) time.Duration {
	if start.After(end) {
		panic(fmt.Sprintf("start time (%s) cannot be after end time (%s)", start, end))
	}
	if limitMin.After(limitMax) {
		panic(fmt.Sprintf("minimum limit time (%s) cannot be after maximum limit time (%s)", limitMin, limitMax))
	}
	if start.After(limitMax) || end.Before(limitMin) {
		// no intersection between the start-end and limitMin-limitMax time ranges
		return 0
	}
	return minTime(end, limitMax).Sub(maxTime(start, limitMin))
}

func calculateTokenIncrease(timeDelta time.Duration, rewardPerSecond sdk.Coins) sdk.Coins {
	durationSeconds := int64(math.RoundToEven(timeDelta.Seconds()))
	if durationSeconds <= 0 {
		// If the duration is zero, there will be no increment.
		// So return an empty increment instead of one full of zeros.
		return sdk.Coins{}
	}
	var increasement sdk.Coins
	rewardPerSecond.Sort()
	for _, el := range rewardPerSecond {
		amt := el.Amount
		amt = amt.Mul(sdk.NewInt(durationSeconds))
		c := sdk.NewCoin(el.Denom, amt)
		increasement = increasement.Add(c)
	}
	return increasement
}

// AfterSPVInterestPaid is called after the interest is paid to the pool
func (k Keeper) AfterSPVInterestPaid(ctx context.Context, poolID string, interestPaid sdkmath.Int) {
	poolInfo, ok := k.spvKeeper.GetPools(ctx, poolID)
	if !ok {
		ctx.Logger().Error("pool not found", "poolID", poolID)
		return
	}

	rewards, ok := k.GetSPVReward(ctx, poolID)
	if !ok {
		ctx.Logger().Info("No rewards for the pool", "poolID", poolID)
		return
	}

	totalRewards := rewards.PaymentAmount.Sort()

	reserve := poolInfo.ReserveFactor
	reserveAmt := sdk.NewDecFromInt(interestPaid).Mul(reserve).TruncateInt()
	paymentToInvestor := interestPaid.Sub(reserveAmt)
	allNFTs := poolInfo.PoolNFTIds

	leftTotalRewards := sdk.NewCoins(totalRewards...)
	if len(allNFTs) == 0 {
		ctx.Logger().Error("No NFTs in the pool")
		return
	}
	for _, classID := range allNFTs[1:] {
		class, found := k.NftKeeper.GetClass(ctx, classID)
		if !found {
			errmsg := fmt.Sprintf("fail to find the linked class %s", classID)
			panic(errmsg)
		}

		var borrowInterest types2.BorrowInterest
		var err error
		err = proto.Unmarshal(class.Data.Value, &borrowInterest)
		if err != nil {
			panic(err)
		}
		// if we do not have the payment info, we skip
		if len(borrowInterest.Payments) == 0 {
			continue
		}

		lastBorrow := borrowInterest.BorrowDetails[len(borrowInterest.BorrowDetails)-1].BorrowedAmount
		lastpayment := borrowInterest.Payments[len(borrowInterest.Payments)-1]
		ratio := sdk.NewDecFromInt(lastpayment.PaymentAmount.Amount).Quo(sdk.NewDecFromInt(paymentToInvestor))

		var incentiveCoins sdk.Coins
		for _, eachCoin := range totalRewards {
			amt := sdk.NewDecFromInt(eachCoin.Amount).Mul(ratio).TruncateInt()
			incentiveCoins = incentiveCoins.Add(sdk.NewCoin(eachCoin.Denom, amt))
		}

		thisIncentivePayment := types2.IncentivePaymentItem{
			PaymentAmount:  incentiveCoins,
			PaymentTime:    lastpayment.PaymentTime,
			BorrowedAmount: lastBorrow,
		}

		if borrowInterest.IncentivePayments == nil {
			borrowInterest.IncentivePayments = []*types2.IncentivePaymentItem{&thisIncentivePayment}
		} else {
			borrowInterest.IncentivePayments = append(borrowInterest.IncentivePayments, &thisIncentivePayment)
		}

		data, err := types3.NewAnyWithValue(&borrowInterest)
		if err != nil {
			panic("pack class any data failed")
		}
		class.Data = data

		err = k.NftKeeper.UpdateClass(ctx, class)
		if err != nil {
			panic("fail to update the class with err" + err.Error())
		}

		if leftTotalRewards.IsAllGT(incentiveCoins) {
			leftTotalRewards = leftTotalRewards.Sub(incentiveCoins...)
		} else {
			leftTotalRewards = incentiveCoins
			break
		}
	}

	class, found := k.NftKeeper.GetClass(ctx, allNFTs[0])
	if !found {
		errmsg := fmt.Sprintf("fail to find the linked class %s", allNFTs[0])
		panic(errmsg)
	}

	var borrowInterest types2.BorrowInterest
	var err error
	err = proto.Unmarshal(class.Data.Value, &borrowInterest)
	if err != nil {
		panic(err)
	}

	if len(borrowInterest.Payments) == 0 {
		return
	}

	lastBorrow := borrowInterest.BorrowDetails[len(borrowInterest.BorrowDetails)-1].BorrowedAmount
	lastpayment := borrowInterest.Payments[len(borrowInterest.Payments)-1]

	thisIncentivePayment := types2.IncentivePaymentItem{
		PaymentAmount:  leftTotalRewards,
		PaymentTime:    lastpayment.PaymentTime,
		BorrowedAmount: lastBorrow,
	}

	if borrowInterest.IncentivePayments == nil {
		borrowInterest.IncentivePayments = []*types2.IncentivePaymentItem{&thisIncentivePayment}
	} else {
		borrowInterest.IncentivePayments = append(borrowInterest.IncentivePayments, &thisIncentivePayment)
	}

	data, err := types3.NewAnyWithValue(&borrowInterest)
	if err != nil {
		panic("pack class any data failed")
	}
	class.Data = data

	err = k.NftKeeper.UpdateClass(ctx, class)
	if err != nil {
		panic("fail to update the class with err" + err.Error())
	}

	// now we set the remaining rewards to the first class
	k.SetSPVReward(ctx, poolID, types.SPVRewardAccTokens{})

	return
}

func (k Keeper) BeforeNFTBurn(ctx context.Context, poolIndex, incestorAddr string, nfts []string) error {
	amt, err := CalculateTotalIncentives(ctx, nfts, k.NftKeeper, true)
	if err != nil {
		ctx.Logger().Error("fail to calculate the total incentives", "error", err)
		return err
	}

	k.SetSPVInvestorReward(ctx, poolIndex, incestorAddr, amt)

	return nil
}

func CalculateTotalIncentives(ctx context.Context, lendNFTs []string, nftKeeper types.NFTKeeper, updateNFT bool) (sdk.Coins, error) {
	allTotalIncentives := sdk.NewCoins()
	for _, el := range lendNFTs {
		ids := strings.Split(el, ":")
		thisNFT, found := nftKeeper.GetNFT(ctx, ids[0], ids[1])
		if !found {
			return sdk.Coins{}, coserrors.Wrapf(types.ErrInvalidNFT, "the given nft %v cannot ben found in storage", ids[1])
		}
		var investorInterestData types2.NftInfo
		err := proto.Unmarshal(thisNFT.Data.Value, &investorInterestData)
		if err != nil {
			panic(err)
		}

		borrowClass, found := nftKeeper.GetClass(ctx, ids[0])
		if !found {
			panic("it should never fail to find the class")
		}

		var borrowClassInfo types2.BorrowInterest
		err = proto.Unmarshal(borrowClass.Data.Value, &borrowClassInfo)
		if err != nil {
			panic(err)
		}

		allIncentivePayments := borrowClassInfo.IncentivePayments
		lastPaymentSet := false

		// no new interest payment
		if len(allIncentivePayments) <= int(investorInterestData.IncentivePaymentOffset) {
			return sdk.NewCoins(), nil
		}
		counter := 0
		allNewIncentivePayments := allIncentivePayments[investorInterestData.IncentivePaymentOffset:]
		for _, eachIncentivePayment := range allNewIncentivePayments {
			counter++
			if eachIncentivePayment.PaymentAmount.IsZero() {
				continue
			}
			classBorrowedAmount := eachIncentivePayment.BorrowedAmount
			incentivePaymentAmount := eachIncentivePayment.PaymentAmount
			// todo there may be the case that because of the truncate, the total payment is larger than the interest paid to investors
			// fixme we should NEVER calculate the interest after the pool status is in luquidation as the user ratio is not correct any more

			var incentiveCoins sdk.Coins
			for _, eachCoin := range incentivePaymentAmount {
				incentiveAmt := sdk.NewDecFromInt(eachCoin.Amount).Mul(sdk.NewDecFromInt(investorInterestData.Borrowed.Amount)).Quo(sdk.NewDecFromInt(classBorrowedAmount.Amount)).TruncateInt()
				incentive := sdk.NewCoin(eachCoin.Denom, incentiveAmt)
				incentiveCoins = incentiveCoins.Add(incentive)
			}

			incentiveCoins.Sort()
			allTotalIncentives = allTotalIncentives.Add(incentiveCoins...)
			lastPaymentSet = true
		}
		if updateNFT && lastPaymentSet {
			investorInterestData.IncentivePaymentOffset += uint32(counter)
			data, err := types3.NewAnyWithValue(&investorInterestData)
			if err != nil {
				panic("pack class any data failed")
			}
			thisNFT.Data = data
			err = nftKeeper.Update(ctx, thisNFT)
			if err != nil {
				panic(err)
			}
		}
	}
	return allTotalIncentives, nil
}

// ClaimSPVReward pays out rewards from a claim to a receiver account for RWA incentives.
func (k Keeper) ClaimSPVReward(ctx context.Context, poolIndex string, investorID sdk.AccAddress) (sdk.Coins, error) {
	amt, found := k.GetSPVInvestorReward(ctx, poolIndex, investorID.String())
	if !found {
		ctx.Logger().Debug("No rewards to claim", "poolIndex", poolIndex, "investorID", investorID)
	}
	k.DeleteSPVInvestorReward(ctx, poolIndex, investorID.String())

	depositor, ok := k.spvKeeper.GetDepositor(ctx, poolIndex, investorID)
	if !ok {
		return amt, nil
	}
	newincentives, err := CalculateTotalIncentives(ctx, depositor.LinkedNFT, k.NftKeeper, true)
	if err != nil {
		ctx.Logger().Error("fail to calculate the total incentives", "error", err)
		return sdk.Coins{}, err
	}

	amt.Sort()
	newincentives.Sort()

	total := newincentives.Add(amt...)

	return total, nil
}

func (k Keeper) GetSPVRewards(ctx context.Context, poolIndex string, investorID sdk.AccAddress) (sdk.Coins, error) {
	amt, found := k.GetSPVInvestorReward(ctx, poolIndex, investorID.String())
	if !found {
		ctx.Logger().Debug("No rewards to claim", "poolIndex", poolIndex, "investorID", investorID)
	}

	depositor, ok := k.spvKeeper.GetDepositor(ctx, poolIndex, investorID)
	if !ok {
		return amt, nil
	}
	newincentives, err := CalculateTotalIncentives(ctx, depositor.LinkedNFT, k.NftKeeper, false)
	if err != nil {
		ctx.Logger().Error("fail to calculate the total incentives", "error", err)
		return sdk.Coins{}, err
	}

	amt.Sort()
	newincentives.Sort()

	total := newincentives.Add(amt...)

	return total, nil
}
