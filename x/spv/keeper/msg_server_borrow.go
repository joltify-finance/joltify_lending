package keeper

import (
	"context"
	coserrors "cosmossdk.io/errors"
	"fmt"
	types2 "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k msgServer) processBorrow(ctx sdk.Context, poolInfo *types.PoolInfo, nftClass nfttypes.Class, amount sdk.Coin) error {

	//macc := k.accKeeper.GetModuleAccount(ctx, types.ModuleAccount)
	//modAccCoin := k.bankKeeper.GetBalance(ctx, macc.GetAddress(), amount.GetDenom())
	//
	//if modAccCoin.IsLT(amount) {
	//	return nil, types.ErrInsufficientFund
	//}
	if poolInfo.BorrowableAmount.IsLT(amount) {
		return types.ErrInsufficientFund
	}
	utilization := sdk.NewDecFromInt(amount.Amount).QuoTruncate(sdk.NewDecFromInt(poolInfo.BorrowableAmount.Amount))

	// we transfer the fund from the module to the spv
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccount, poolInfo.OwnerAddress, sdk.NewCoins(amount))
	if err != nil {
		return err
	}

	// we add the amount of the tokens that borrowed in the pool and decreases the borrowable
	poolInfo.BorrowedAmount = poolInfo.BorrowedAmount.Add(amount)
	poolInfo.BorrowableAmount = poolInfo.BorrowableAmount.Sub(amount)

	// we update each investor leftover
	// todo YB it seems we do not need to store the deposit wallets
	//investorWallets, found := k.GetPoolDepositedWallets(ctx, poolInfo.Index)
	//if found {
	//	panic("should never happened that investors have depposited the money while the store is empty")
	//}
	k.processInvestors(ctx, poolInfo, utilization, nftClass)
	return nil
}

func (k msgServer) processInvestors(ctx sdk.Context, poolInfo *types.PoolInfo, utilization sdk.Dec, nftClass nfttypes.Class) error {

	// now we update the depositor's withdrawal amount and locked amount
	var depositors []types.DepositorInfo
	k.IterateDepositors(ctx, poolInfo.Index, func(depositor types.DepositorInfo) (stop bool) {
		locked := sdk.NewDecFromInt(depositor.WithdrawableAmount.Amount).Mul(utilization).TruncateInt()
		depositor.LockedAmount = sdk.NewCoin(depositor.WithdrawableAmount.Denom, locked)
		depositor.WithdrawableAmount = depositor.WithdrawableAmount.SubAmount(locked)
		depositors = append(depositors, depositor)
		return false
	})

	nftTemplate := nfttypes.NFT{
		ClassId: nftClass.Id,
		Uri:     nftClass.Uri,
		UriHash: nftClass.UriHash,
	}

	for _, el := range depositors {
		// nft ID is the hash(nft class ID, investorWallet)
		indexHash := crypto.Keccak256Hash([]byte(nftClass.Id), el.DepositorAddress)
		nftTemplate.Id = indexHash.Hex()

		userData := types.NftInfo{Issuer: poolInfo.PoolName, Receiver: el.DepositorAddress.String(), IssueTime: ctx.BlockTime()}
		data, err := types2.NewAnyWithValue(&userData)
		if err != nil {
			panic("should never fail")
		}
		nftTemplate.Data = data
		err = k.nftKeeper.Mint(ctx, nftTemplate, el.DepositorAddress)
		if err != nil {
			return types.ErrMINTNFT
		}
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

	if poolInfo.PoolStatus == types.PoolInfo_CLOSED {
		return nil, coserrors.Wrap(types.ErrPoolClosed, "pool has been closed")
	}

	if !poolInfo.OwnerAddress.Equals(caller) {
		return nil, coserrors.Wrapf(types.ErrUnauthorized, "%v is not authorized to borrow money", msg.Creator)
	}

	if msg.BorrowAmount.Denom != poolInfo.TotalAmount.Denom {
		return nil, coserrors.Wrap(types.ErrInconsistencyToken, "token to be borrowed is inconsistency")
	}

	// create the new nft class for this borrow event
	poolClass, found := k.nftKeeper.GetClass(ctx, poolInfo.Index)
	if !found {
		panic("pool class must have already been set")
	}

	latestSeries := len(poolInfo.PoolNFTIds)

	currentBorrowClass := poolClass
	currentBorrowClass.Id = fmt.Sprintf("%v-%v", currentBorrowClass.Id, latestSeries)

	bi := types.BorrowInterest{
		PoolIndex: poolInfo.Index,
		Apy:       poolInfo.Apy,
		PayFreq:   poolInfo.PayFreq,
		IssueTime: ctx.BlockTime(),
		Borrowed:  msg.BorrowAmount,
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
	k.SetPool(ctx, poolInfo)
	return &types.MsgBorrowResponse{}, nil
}
