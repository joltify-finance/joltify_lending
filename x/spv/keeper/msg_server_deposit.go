package keeper

import (
	"context"
	"time"

	sdkmath "cosmossdk.io/math"

	coserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k msgServer) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	ctx = ctx.WithGasMeter(sdk.NewInfiniteGasMeter())

	investor, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, coserrors.Wrapf(errorsmod.ErrInvalidAddress, "invalid address %v", msg.Creator)
	}

	poolInfo, ok := k.GetPools(ctx, msg.GetPoolIndex())
	if !ok {
		return nil, coserrors.Wrapf(errorsmod.ErrNotFound, "pool cannot be found %v", msg.GetPoolIndex())
	}

	if poolInfo.PoolStatus != types.PoolInfo_ACTIVE {
		if poolInfo.PoolStatus != types.PoolInfo_PooLPayPartially {
			return nil, coserrors.Wrapf(types.ErrPoolNotActive, "pool is not active or in partially paid status, current: %v", poolInfo.PoolStatus)
		}
	}

	if poolInfo.CurrentPoolTotalBorrowCounter == 0 && ctx.BlockTime().After(poolInfo.PoolCreatedTime.Add(time.Second*time.Duration(poolInfo.PoolLockedSeconds))) {
		return nil, types.ErrPoolNotAcceptNewFund
	}

	if msg.Token.GetDenom() != poolInfo.TargetAmount.Denom {
		return nil, coserrors.Wrapf(errorsmod.ErrInvalidCoins, "we only accept %v", poolInfo.TargetAmount.Denom)
	}

	resp, err := k.kycKeeper.GetByWallet(ctx, msg.Creator)
	if err != nil {
		return nil, coserrors.Wrapf(errorsmod.ErrNotFound, "the investor cannot be found %v", msg.Creator)
	}

	investorsResp, found := k.GetInvestorToPool(ctx, msg.PoolIndex)
	if !found {
		return nil, coserrors.Wrapf(types.ErrPoolNotFound, "the pool %v does not exist", msg.PoolIndex)
	}

	found = false
	for _, el := range investorsResp.Investors {
		if el == resp.InvestorId {
			found = true
		}
	}

	if !found {
		return nil, coserrors.Wrapf(types.ErrUnauthorized, "the given investor is not allowed to invest %v", msg.Creator)
	}

	// now we transfer the token from the investor to the pool.
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, investor, types.ModuleAccount, sdk.NewCoins(msg.Token))
	if err != nil {
		return nil, coserrors.Wrapf(errorsmod.ErrInvalidRequest, "fail to transfer the token to the pool %v", err)
	}

	// now we update the users deposit data
	previousDepositor, found := k.GetDepositor(ctx, poolInfo.Index, investor)
	if !found {
		if msg.Token.Amount.Sub(poolInfo.MinDepositAmount).IsNegative() {
			return nil, coserrors.Wrapf(types.ErrDeposit, "the deposit amount %v is less than the minimum deposit amount %v", msg.Token.Amount.String(), poolInfo.MinDepositAmount.String())
		}
		depositor := types.DepositorInfo{InvestorId: resp.InvestorId, DepositorAddress: investor, PoolIndex: msg.PoolIndex, LockedAmount: sdk.NewCoin(poolInfo.BorrowedAmount.Denom, sdkmath.ZeroInt()), WithdrawalAmount: msg.Token, LinkedNFT: []string{}, DepositType: types.DepositorInfo_unset, PendingInterest: sdk.NewCoin(msg.Token.Denom, sdk.ZeroInt())}
		k.SetDepositor(ctx, depositor)
	} else {
		switch previousDepositor.DepositType {
		case types.DepositorInfo_unusable, types.DepositorInfo_processed:

			poolInfo.UsableAmount = poolInfo.UsableAmount.Add(previousDepositor.WithdrawalAmount)
			previousDepositor.DepositType = types.DepositorInfo_unset
			previousDepositor.WithdrawalAmount = previousDepositor.WithdrawalAmount.Add(msg.Token)

			if previousDepositor.WithdrawalAmount.Amount.Sub(poolInfo.MinDepositAmount).IsNegative() {
				return nil, coserrors.Wrapf(types.ErrDeposit, "the deposit amount %v is less than the minimum deposit amount %v", previousDepositor.WithdrawalAmount.Amount.String(), poolInfo.MinDepositAmount.String())
			}
			k.SetDepositor(ctx, previousDepositor)

		case types.DepositorInfo_unset:

			previousDepositor.WithdrawalAmount = previousDepositor.WithdrawalAmount.Add(msg.Token)

			if previousDepositor.WithdrawalAmount.Amount.Sub(poolInfo.MinDepositAmount).IsNegative() {
				return nil, coserrors.Wrapf(types.ErrDeposit, "the deposit amount %v is less than the minimum deposit amount %v", previousDepositor.WithdrawalAmount.Amount.String(), poolInfo.MinDepositAmount.String())
			}
			k.SetDepositor(ctx, previousDepositor)

		default:
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
