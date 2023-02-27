package keeper

import (
	"context"
	sdkmath "cosmossdk.io/math"
	"fmt"

	coserrors "cosmossdk.io/errors"
	types2 "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k msgServer) processBorrow(ctx sdk.Context, poolInfo *types.PoolInfo, nftClass nfttypes.Class, amount sdk.Coin) error {

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
	poolInfo.BorrowableAmount = poolInfo.BorrowableAmount.Sub(amount)

	// we update each investor leftover
	k.processInvestors(ctx, poolInfo, utilization, amount.Amount, nftClass)
	return nil
}

func (k msgServer) doProcessInvestor(ctx sdk.Context, depositor *types.DepositorInfo, locked, totalBorrow sdkmath.Int, nftTemplate nfttypes.NFT, nftClassId string, poolInfo *types.PoolInfo, errGlobal error) {
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

func (k msgServer) processInvestors(ctx sdk.Context, poolInfo *types.PoolInfo, utilization sdk.Dec, totalBorrow sdkmath.Int, nftClass nfttypes.Class) error {

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

	if poolInfo.PoolStatus != types.PoolInfo_ACTIVE {
		return nil, coserrors.Wrap(types.ErrPoolNotActive, "pool has been closed")
	}

	if !poolInfo.OwnerAddress.Equals(caller) {
		return nil, coserrors.Wrapf(types.ErrUnauthorized, "%v is not authorized to borrow money", msg.Creator)
	}

	if msg.BorrowAmount.Denom != poolInfo.TotalAmount.Denom {
		return nil, coserrors.Wrap(types.ErrInconsistencyToken, "token to be borrowed is inconsistency")
	}

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
	firstPayment := types.PaymentItem{PaymentTime: ctx.BlockTime(), PaymentAmount: sdk.NewCoin(msg.BorrowAmount.Denom, sdk.NewInt(0))}
	bi := types.BorrowInterest{
		PoolIndex:    poolInfo.Index,
		Apy:          poolInfo.Apy,
		PayFreq:      poolInfo.PayFreq,
		IssueTime:    ctx.BlockTime(),
		Borrowed:     msg.BorrowAmount,
		BorrowedLast: msg.BorrowAmount,
		MonthlyRatio: i,
		InterestSPY:  rate,
		Payments:     []*types.PaymentItem{&firstPayment},
	}

	data, err := types2.NewAnyWithValue(&bi)
	if err != nil {
		panic(err)
	}
	currentBorrowClass.Data = data
	k.nftKeeper.SaveClass(ctx, currentBorrowClass)

	// update the borrow series
	poolInfo.PoolNFTIds = append(poolInfo.PoolNFTIds, currentBorrowClass.Id)

	err = k.processBorrow(ctx, &poolInfo, currentBorrowClass, msg.BorrowAmount)
	if err != nil {
		return nil, err
	}

	// we finally update the pool info
	poolInfo.PoolStatus = types.PoolInfo_ACTIVE
	k.SetPool(ctx, poolInfo)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeBorrow,
			sdk.NewAttribute(types.AttributeCreator, msg.Creator),
			sdk.NewAttribute("amount", msg.BorrowAmount.Amount.String()),
		),
	)

	return &types.MsgBorrowResponse{BorrowAmount: msg.BorrowAmount.String()}, nil
}
