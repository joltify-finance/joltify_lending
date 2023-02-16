package keeper

import (
	"context"

	coserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	nft "github.com/cosmos/cosmos-sdk/x/nft/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k msgServer) processBorrow(ctx sdk.Context, poolInfo *types.PoolInfo, amount sdk.Coin) (*types.PoolInfo, error) {

	//macc := k.accKeeper.GetModuleAccount(ctx, types.ModuleAccount)
	//modAccCoin := k.bankKeeper.GetBalance(ctx, macc.GetAddress(), amount.GetDenom())
	//
	//if modAccCoin.IsLT(amount) {
	//	return nil, types.ErrInsufficientFund
	//}
	if poolInfo.BorrowableAmount.IsLT(amount) {
		return nil, types.ErrInsufficientFund
	}
	utilization := sdk.NewDecFromInt(amount.Amount).QuoTruncate(sdk.NewDecFromInt(poolInfo.BorrowableAmount.Amount))

	// we transfer the fund from the module to the spv
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccount, poolInfo.OwnerAddress, sdk.NewCoins(amount))
	if err != nil {
		return nil, err
	}

	// we add the amount of the tokens that borrowed in the pool and decreases the borrowable
	poolInfo.BorrowedAmount = poolInfo.BorrowedAmount.Add(amount)
	poolInfo.BorrowableAmount = poolInfo.BorrowableAmount.Sub(amount)

	// we update each investor leftover
	investorWallets, found := k.GetPoolDepositedWallets(ctx, poolInfo.Index)
	if found {
		panic("should never happened that investors have depposited the money while the store is empty")
	}
	k.processInvestors(ctx, investorWallets.WalletAddress, poolInfo, utilization)
	return poolInfo, nil
}

func (k msgServer) processInvestors(ctx sdk.Context, investorWallets []sdk.AccAddress, poolInfo *types.PoolInfo, utilization sdk.Dec) {

	var depositors []types.DepositorInfo

	k.IterateDepositors(ctx, poolInfo.Index, func(depositor types.DepositorInfo) (stop bool) {
		locked := sdk.NewDecFromInt(depositor.WithdrawableAmount.Amount).Mul(utilization).TruncateInt()
		depositor.LockedAmount = sdk.NewCoin(depositor.WithdrawableAmount.Denom, locked)
		depositor.WithdrawableAmount = depositor.WithdrawableAmount.SubAmount(locked)
		depositors = append(depositors)

		nft.



		return false
	})

	for _, el := range investorWallets {
		depositor, found := k.GetDepositor(ctx, poolInfo.Index, el)
		if !found {
			panic("shoud never fail to find the depositor")
		}

		depositor.WithdrawableAmount

	}

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
		return nil, coserrors.Wrapf(types.ErrUnauthorized, "caller is not authorized to borrow money", msg.Creator)
	}

	if msg.BorrowAmount.Denom != poolInfo.TotalAmount.Denom {
		return nil, coserrors.Wrap(types.ErrInconsistencyToken, "token to be borrowed is inconsistency")
	}

	k.processBorrow(ctx, &poolInfo, msg.BorrowAmount)

	return &types.MsgBorrowResponse{}, nil
}
