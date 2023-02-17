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
	return amount.Add(newAmount).IsLT(poolInfo.TotalAmount)
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

	req := kyctypes.QueryByWalletRequest{
		Wallet: msg.Creator,
	}

	resp, err := k.kycKeeper.QueryByWallet(goCtx, &req)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrNotFound, "the investor cannot be found %v", msg.Creator)
	}

	investorsResp, found := k.GetInvestorToPool(ctx, msg.PoolIndex)
	if !found {
		return nil, coserrors.Wrapf(types.ErrPoolNotFound, "the pool %v does not exist", msg.PoolIndex)
	}

	found = false
	for _, el := range investorsResp.Investors {
		if el == resp.Investor.InvestorId {
			found = true
		}
	}

	if !found {
		return nil, coserrors.Wrapf(types.ErrUnauthorized, "the given investor is not allowed to invest %v", msg.Creator)
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
		depositor := types.DepositorInfo{InvestorId: resp.Investor.InvestorId, DepositorAddress: investor, PoolIndex: msg.PoolIndex, WithdrawableAmount: msg.Token}
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
	poolInfo.BorrowableAmount = poolInfo.BorrowableAmount.Add(msg.Token)
	k.SetPool(ctx, poolInfo)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeDeposit,
			sdk.NewAttribute(types.AttributeCreator, msg.Creator),
		),
	)

	return &types.MsgDepositResponse{}, nil
}
