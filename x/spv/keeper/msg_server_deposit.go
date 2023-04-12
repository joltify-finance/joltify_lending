package keeper

import (
	"context"
	sdkmath "cosmossdk.io/math"
	"time"

	coserrors "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	kyctypes "github.com/joltify-finance/joltify_lending/x/kyc/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

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

	if poolInfo.PoolStatus != types.PoolInfo_ACTIVE {
		return nil, coserrors.Wrapf(types.ErrPoolNotActive, "pool is not active")
	}

	if poolInfo.CurrentPoolTotalBorrowCounter == 0 && ctx.BlockTime().After(poolInfo.PoolCreatedTime.Add(time.Second*time.Duration(poolInfo.PoolLockedSeconds))) {
		return nil, types.ErrPoolNotAcceptNewFund
	}

	if msg.Token.GetDenom() != poolInfo.TargetAmount.Denom {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidCoins, "we only accept %v", poolInfo.TargetAmount.Denom)
	}

	req := kyctypes.QueryByWalletRequest{
		Wallet: msg.Creator,
	}

	resp, err := k.kycKeeper.QueryByWallet(goCtx, &req)
	if err != nil || resp.Investor == nil {
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

	// now we transfer the token from the investor to the pool.
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, investor, types.ModuleAccount, sdk.NewCoins(msg.Token))
	if err != nil {
		return nil, coserrors.Wrap(sdkerrors.ErrInvalidRequest, "fail to transfer the toekn to the pool")
	}

	// now we update the users deposit data
	previousDepositor, found := k.GetDepositor(ctx, poolInfo.Index, investor)
	if !found {
		depositor := types.DepositorInfo{InvestorId: resp.Investor.InvestorId, DepositorAddress: investor, PoolIndex: msg.PoolIndex, LockedAmount: sdk.NewCoin(msg.Token.Denom, sdkmath.ZeroInt()), WithdrawalAmount: msg.Token, LinkedNFT: []string{}, DepositType: types.DepositorInfo_unset, PendingInterest: sdk.NewCoin(msg.Token.Denom, sdk.ZeroInt())}
		k.SetDepositor(ctx, depositor)

	} else {
		if (previousDepositor.DepositType == types.DepositorInfo_unset) || (previousDepositor.DepositType == types.DepositorInfo_processed) {
			if previousDepositor.DepositType == types.DepositorInfo_processed {
				poolInfo.UsableAmount = poolInfo.UsableAmount.Add(previousDepositor.WithdrawalAmount)
				previousDepositor.DepositType = types.DepositorInfo_unset
			}
			previousDepositor.WithdrawalAmount = previousDepositor.WithdrawalAmount.Add(msg.Token)
			k.SetDepositor(ctx, previousDepositor)
		} else {
			return nil, coserrors.Wrapf(types.ErrDeposit, "you are not allow to deposit as %v", previousDepositor.DepositType)
		}
	}

	// now we update borrowable
	poolInfo.UsableAmount = poolInfo.UsableAmount.Add(msg.Token)
	k.SetPool(ctx, poolInfo)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeDeposit,
			sdk.NewAttribute(types.AttributeCreator, msg.Creator),
		),
	)

	return &types.MsgDepositResponse{}, nil
}
