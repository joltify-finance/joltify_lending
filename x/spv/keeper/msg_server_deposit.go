package keeper

import (
	"context"

	coserrors "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	kyctypes "github.com/joltify-finance/joltify_lending/x/kyc/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k msgServer) checkAcceptNewDeposit(ctx sdk.Context, poolInfo types.PoolInfo, newAmount sdk.Coin) bool {

	acc := k.accKeeper.GetModuleAccount(ctx, types.ModuleAccount)

	amount := k.bankKeeper.GetBalance(ctx, acc.GetAddress(), poolInfo.TotalAmount.GetDenom())
	if amount.Add(newAmount).IsLT(poolInfo.TotalAmount) {
		return false
	}
	return true

}

func (k msgServer) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	investor, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address %v", msg.Creator)
	}

	poolInfo, ok := k.GetPools(ctx, msg.GetPoolIndex())
	if !ok {
		return nil, coserrors.Wrapf(sdkerrors.ErrNotFound, "pool cannot be found %v", msg.GetPoolIndex())
	}

	if poolInfo.PoolStatus == types.PoolInfo_CLOSED {
		return nil, coserrors.Wrapf(types.ErrPoolClosed, "pool has been closed")
	}

	acceptFund := k.checkAcceptNewDeposit(ctx, poolInfo, msg.Token)
	if !acceptFund {
		return nil, types.ErrPoolFull
	}

	req := kyctypes.QueryInvestorWalletsRequest{
		InvestorId: msg.GetInvestorID(),
	}

	resp, err := k.kycKeeper.QueryInvestorWallets(goCtx, &req)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrNotFound, "the investor cannot be found %v", msg.InvestorID)
	}

	found := false
	for _, el := range resp.Wallets {
		if el == msg.Creator {
			found = true
			break
		}
	}

	if !found {
		return nil, coserrors.Wrapf(types.ErrUnauthorized, "the given wallet cannot be found %v", msg.Creator)
	}

	investors, found := k.GetInvestorToPool(ctx, poolInfo.Index)
	if !found {
		return nil, coserrors.Wrap(types.ErrUnauthorized, "the investor cannot be found")
	}

	found = false
	for _, el := range investors.Investors {
		if el == msg.InvestorID {
			found = true
			break
		}
	}
	if !found {
		return nil, coserrors.Wrapf(types.ErrUnauthorized, "the given investor is not allowed to invest %v", msg.InvestorID)
	}

	if msg.Token.GetDenom() != poolInfo.GetTotalAmount().Denom {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidCoins, "we only accept %v", poolInfo.GetTotalAmount().Denom)
	}

	// now we transfer the token from the investor to the pool.
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, investor, types.ModuleAccount, sdk.NewCoins(msg.Token))
	if err != nil {
		return nil, coserrors.Wrap(sdkerrors.ErrInvalidRequest, "fail to transfer the toekn to the pool")
	}

	// now we update the users deposit data
	previousDepositor, found := k.GetDepositor(ctx, poolInfo.Index, investor)
	if !found {
		depositor := types.DepositorInfo{InvestorId: msg.InvestorID, DepositorAddress: investor, PoolIndex: msg.PoolIndex, WithdrawableAmount: msg.Token}
		k.SetDepositor(ctx, depositor)

	} else {
		previousDepositor.WithdrawableAmount = previousDepositor.WithdrawableAmount.Add(msg.Token)
		k.SetDepositor(ctx, previousDepositor)
	}

	wallets, found := k.GetPoolDepositedWallets(ctx, poolInfo.Index)
	if !found {
		depositorWallets := types.PoolDepositedInvestors{PoolIndex: poolInfo.Index, WalletAddress: []sdk.AccAddress{investor}}
		k.SetPoolDepositedWallets(ctx, depositorWallets)
	} else {
		wallets.WalletAddress = addAddrToList(wallets.WalletAddress, investor)
		k.SetPoolDepositedWallets(ctx, wallets)
	}

	// now we update borrowable
	poolInfo.BorrowedAmount = poolInfo.BorrowableAmount.Add(msg.Token)
	k.SetPool(ctx, poolInfo)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeDeposit,
			sdk.NewAttribute(types.AttributeCreator, msg.Creator),
		),
	)

	return &types.MsgDepositResponse{}, nil
}
