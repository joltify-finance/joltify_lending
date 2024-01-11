package keeper

import (
	"context"
	"time"

	sdkmath "cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gogo/protobuf/proto"

	coserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k msgServer) getAllBorrowed(ctx sdk.Context, poolInfo types.PoolInfo) sdkmath.Int {
	var err error
	sum := sdk.ZeroInt()
	for _, el := range poolInfo.PoolNFTIds {

		class, found := k.nftKeeper.GetClass(ctx, el)
		if !found {
			panic(found)
		}
		var borrowInterest types.BorrowInterest
		err = proto.Unmarshal(class.Data.Value, &borrowInterest)
		if err != nil {
			panic(err)
		}
		lastBorrow := borrowInterest.BorrowDetails[len(borrowInterest.BorrowDetails)-1]

		amount := lastBorrow.BorrowedAmount
		ratio := lastBorrow.ExchangeRatio
		usdTotal := outboundConvertToUSD(amount.Amount, ratio)
		sum = sum.Add(usdTotal)
	}
	return sum
}

func checkEligibility(blockTime time.Time, poolInfo types.PoolInfo, borrowAmount sdk.Coin) error {
	if poolInfo.PoolStatus != types.PoolInfo_ACTIVE {
		if poolInfo.PoolStatus != types.PoolInfo_PooLPayPartially {
			return coserrors.Wrapf(types.ErrPoolNotActive, "pool is not in active status or partially paid status, current: %v", poolInfo.PoolStatus)
		}
	}

	if poolInfo.CurrentPoolTotalBorrowCounter >= poolInfo.PoolTotalBorrowLimit {
		return types.ErrPoolBorrowLimit
	}

	if poolInfo.CurrentPoolTotalBorrowCounter == 0 && poolInfo.PoolCreatedTime.Add(time.Second*time.Duration(poolInfo.PoolLockedSeconds)+poolInfo.GraceTime).Before(blockTime) {
		return types.ErrPoolBorrowExpire
	}

	if poolInfo.UsableAmount.IsLT(borrowAmount) {
		return types.ErrInsufficientFund
	}

	if poolInfo.CurrentPoolTotalBorrowCounter == 0 && borrowAmount.IsLT(poolInfo.MinBorrowAmount) {
		return coserrors.Wrapf(types.ErrInvalidParameter, "pool minimal borrow is %v and you try to borrow %v", poolInfo.MinBorrowAmount, borrowAmount)
	}
	return nil
}

func (k msgServer) Borrow(goCtx context.Context, msg *types.MsgBorrow) (*types.MsgBorrowResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	ctx = ctx.WithGasMeter(sdk.NewInfiniteGasMeter())

	caller, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address %v", msg.Creator)
	}

	poolInfo, found := k.GetPools(ctx, msg.GetPoolIndex())
	if !found {
		return nil, coserrors.Wrapf(sdkerrors.ErrNotFound, "pool cannot be found %v", msg.GetPoolIndex())
	}

	if msg.BorrowAmount.Denom != poolInfo.TargetAmount.Denom {
		return nil, coserrors.Wrap(types.ErrInconsistencyToken, "token to be borrowed is inconsistency")
	}

	var allBorrowed sdkmath.Int
	// check that junior pool must meet its target amount before senior pool can borrow
	juniorPoolIndex := crypto.Keccak256Hash([]byte(poolInfo.ProjectName), poolInfo.OwnerAddress.Bytes(), []byte("junior"))

	juniorInfo, found := k.GetPools(ctx, juniorPoolIndex.Hex())
	if !found {
		return nil, coserrors.Wrapf(sdkerrors.ErrNotFound, "pool cannot be found %v", msg.GetPoolIndex())
	}
	allBorrowed = k.getAllBorrowed(ctx, juniorInfo)

	if poolInfo.PoolType == types.PoolInfo_SENIOR && !poolInfo.SeparatePool {
		if juniorInfo.TargetAmount.Amount.Sub(allBorrowed).GT(sdk.NewIntFromUint64(10)) {
			return nil, coserrors.Wrapf(types.ErrPoolNotActive, "junior pool has not met its target amount, cannot borrow from senior pool current Borrowed Junior %v and target is %v", allBorrowed, juniorInfo.TargetAmount.Amount)
		}
	}

	err = checkEligibility(ctx.BlockTime(), poolInfo, msg.BorrowAmount)
	if err != nil {
		return nil, err
	}

	if allBorrowed.Add(msg.BorrowAmount.Amount).GT(poolInfo.TargetAmount.Amount) {
		return nil, coserrors.Wrapf(types.ErrPoolFull, "pool reached its borrow limit with current borrowed %v", allBorrowed)
	}

	poolInfo.CurrentPoolTotalBorrowCounter += 1

	if !poolInfo.OwnerAddress.Equals(caller) {
		return nil, coserrors.Wrapf(types.ErrUnauthorized, "%v is not authorized to borrow money", msg.Creator)
	}

	err = k.doBorrow(ctx, &poolInfo, msg.BorrowAmount, true, nil, sdk.ZeroInt(), false)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "borrow failed %v", err)
	}

	// now we need to update the interest prepaid
	if poolInfo.InterestPrepayment != nil {
		currentInterest := poolInfo.EscrowInterestAmount
		a, _ := denomConvertToLocalAndUsd(poolInfo.BorrowedAmount.Denom)
		marketID := denomConvertToMarketID(a)
		counter, interestReceived, _, ratio, err := k.calculatePaymentMonth(ctx, poolInfo, marketID, currentInterest)
		if err != nil {
			return nil, coserrors.Wrapf(err, "calculate payment month failed")
		}

		if counter < 1 {
			// we return the leftover interest to the spv if it cannot be covered one round
			err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccount, poolInfo.OwnerAddress, sdk.Coins{sdk.NewCoin(poolInfo.TargetAmount.Denom, currentInterest)})
			if err != nil {
				return nil, coserrors.Wrapf(err, "fail to transfer the repayment from spv to module")
			}
			poolInfo.EscrowInterestAmount = sdk.ZeroInt()
			poolInfo.InterestPrepayment = nil
			k.SetPool(ctx, poolInfo)
		} else {
			prepayment := types.InterestPrepayment{
				Counter:       counter,
				ExchangeRatio: ratio,
			}
			poolInfo.InterestPrepayment = &prepayment
			needToReturn := poolInfo.EscrowInterestAmount.Sub(interestReceived)
			err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccount, poolInfo.OwnerAddress, sdk.Coins{sdk.NewCoin(poolInfo.TargetAmount.Denom, needToReturn)})
			if err != nil {
				return nil, coserrors.Wrapf(err, "fail to transfer the repayment from spv to module")
			}
			poolInfo.EscrowInterestAmount = interestReceived
			k.SetPool(ctx, poolInfo)
		}
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeBorrow,
			sdk.NewAttribute(types.AttributeCreator, msg.Creator),
			sdk.NewAttribute("amount", msg.BorrowAmount.Amount.String()),
		),
	)

	return &types.MsgBorrowResponse{BorrowAmount: msg.BorrowAmount.String()}, nil
}
