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

func checkEligibility(blockTime time.Time, poolInfo types.PoolInfo) error {
	if poolInfo.PoolStatus != types.PoolInfo_ACTIVE {
		if poolInfo.PoolStatus != types.PoolInfo_PooLPayPartially {
			return coserrors.Wrapf(types.ErrPoolNotActive, "pool is not in active status or partially paid status, current: %v", poolInfo.PoolStatus)
		}
	}

	if poolInfo.CurrentPoolTotalBorrowCounter >= poolInfo.PoolTotalBorrowLimit {
		return types.ErrPoolBorrowLimit
	}

	if poolInfo.InterestPrepayment != nil {
		return coserrors.Wrapf(types.ErrInvalidParameter, "we have the prepayment interest, not accepting new interest payment")
	}

	if poolInfo.CurrentPoolTotalBorrowCounter == 0 && poolInfo.PoolCreatedTime.Add(time.Second*time.Duration(poolInfo.PoolLockedSeconds)+poolInfo.GraceTime).Before(blockTime) {
		return types.ErrPoolBorrowExpire
	}

	token := sdk.NewDecFromIntWithPrec(sdkmath.NewInt(1000), 0)
	minToken := sdk.NewCoin(poolInfo.TargetAmount.Denom, token.RoundInt())
	if poolInfo.CurrentPoolTotalBorrowCounter == 0 && poolInfo.UsableAmount.IsLT(minToken) {
		return coserrors.Wrapf(types.ErrInsufficientFund, "pool target is %v and we have %v usable", poolInfo.TargetAmount, poolInfo.UsableAmount)
	}
	return nil
}

func (k msgServer) Borrow(goCtx context.Context, msg *types.MsgBorrow) (*types.MsgBorrowResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	caller, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address %v", msg.Creator)
	}

	poolInfo, found := k.GetPools(ctx, msg.GetPoolIndex())
	if !found {
		return nil, coserrors.Wrapf(sdkerrors.ErrNotFound, "pool cannot be found %v", msg.GetPoolIndex())
	}

	// check that junior pool must meet its target amount before senior pool can borrow
	if poolInfo.PoolType == types.PoolInfo_SENIOR && !poolInfo.SeparatePool {
		juniorPoolIndex := crypto.Keccak256Hash([]byte(poolInfo.ProjectName), poolInfo.OwnerAddress.Bytes(), []byte("junior"))

		juniorInfo, found := k.GetPools(ctx, juniorPoolIndex.Hex())
		if !found {
			return nil, coserrors.Wrapf(sdkerrors.ErrNotFound, "pool cannot be found %v", msg.GetPoolIndex())
		}
		allBorrowed := k.getAllBorrowed(ctx, juniorInfo)
		if juniorInfo.TargetAmount.Amount.Sub(allBorrowed).GT(sdk.NewIntFromUint64(10)) {
			return nil, coserrors.Wrapf(types.ErrPoolNotActive, "junior pool has not met its target amount, cannot borrow from senior pool current Borrowed Junior %v and target is %v", allBorrowed, juniorInfo.TargetAmount.Amount)
		}
	}

	allBorrowed := k.getAllBorrowed(ctx, poolInfo)

	err = checkEligibility(ctx.BlockTime(), poolInfo)
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

	if msg.BorrowAmount.Denom != poolInfo.TargetAmount.Denom {
		return nil, coserrors.Wrap(types.ErrInconsistencyToken, "token to be borrowed is inconsistency")
	}

	if poolInfo.UsableAmount.IsLT(msg.BorrowAmount) {
		return nil, types.ErrInsufficientFund
	}

	k.doBorrow(ctx, &poolInfo, msg.BorrowAmount, true, nil, sdk.ZeroInt(), false)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeBorrow,
			sdk.NewAttribute(types.AttributeCreator, msg.Creator),
			sdk.NewAttribute("amount", msg.BorrowAmount.Amount.String()),
		),
	)

	return &types.MsgBorrowResponse{BorrowAmount: msg.BorrowAmount.String()}, nil
}
