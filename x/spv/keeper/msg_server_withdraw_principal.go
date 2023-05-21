package keeper

import (
	"context"
	"errors"
	"strings"
	"time"

	coserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k Keeper) handlerPoolClose(ctx sdk.Context, poolInfo types.PoolInfo, depositor types.DepositorInfo) (sdk.Coin, error) {
	amount, err := k.cleanupDepositor(ctx, poolInfo, depositor)
	if err != nil {
		return sdk.Coin{}, err
	}
	tokenSend := sdk.NewCoin(poolInfo.TargetAmount.Denom, amount)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeWithdrawPrincipal,
			sdk.NewAttribute(types.AttributeCreator, depositor.DepositorAddress.String()),
			sdk.NewAttribute(types.AttributeAmount, amount.String()),
		),
	)
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccount, depositor.DepositorAddress, sdk.NewCoins(tokenSend))
	return tokenSend, err
}

func (k Keeper) handlerPoolLiquidation(ctx sdk.Context, depositor types.DepositorInfo) (sdk.Coin, error) {
	interest, err := calculateTotalInterest(ctx, depositor.LinkedNFT, k.nftKeeper, true)
	if err != nil {
		return sdk.Coin{}, err
	}

	totalRedeem, err := k.doProcessLiquidationForInvestor(ctx, depositor.LinkedNFT)
	if err != nil {
		return sdk.Coin{}, err
	}

	depositor.TotalPaidLiquidationAmount = depositor.TotalPaidLiquidationAmount.Add(totalRedeem)
	totalWithdraw := depositor.WithdrawalAmount.AddAmount(interest).AddAmount(totalRedeem)
	depositor.WithdrawalAmount = sdk.NewCoin(depositor.WithdrawalAmount.Denom, sdk.ZeroInt())
	k.SetDepositor(ctx, depositor)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeWithdrawPrincipal,
			sdk.NewAttribute(types.AttributeCreator, depositor.DepositorAddress.String()),
			sdk.NewAttribute(types.AttributeAmount, totalWithdraw.String()),
		),
	)
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccount, depositor.DepositorAddress, sdk.NewCoins(totalWithdraw))
	return totalWithdraw, err
}

func (k Keeper) isEmptyPool(ctx sdk.Context, poolInfo types.PoolInfo) bool {
	if !poolInfo.BorrowedAmount.IsZero() || !poolInfo.UsableAmount.IsZero() {
		return false
	}

	if len(poolInfo.ProcessedTransferAccounts) != 0 || len(poolInfo.ProcessedWithdrawAccounts) != 0 {
		return false
	}

	for _, el := range poolInfo.PoolNFTIds {
		if k.nftKeeper.GetTotalSupply(ctx, el) > 0 {
			return false
		}
	}
	return true
}

func (k msgServer) handleDepositClose(ctx sdk.Context, depositor types.DepositorInfo, poolInfo types.PoolInfo) (*types.MsgWithdrawPrincipalResponse, error) {
	depositor.DepositType = types.DepositorInfo_deactive
	amountToSend := depositor.WithdrawalAmount
	interest, err := k.claimInterest(ctx, &depositor)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address %v", depositor.DepositorAddress.String())
	}

	amountToSend = amountToSend.Add(interest)
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccount, depositor.DepositorAddress, sdk.NewCoins(amountToSend))
	if err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeWithdrawPrincipal,
			sdk.NewAttribute(types.AttributeCreator, depositor.DepositorAddress.String()),
			sdk.NewAttribute(types.AttributeAmount, amountToSend.String()),
		),
	)
	depositor.LockedAmount = sdk.NewCoin(depositor.LockedAmount.Denom, sdk.ZeroInt())
	depositor.WithdrawalAmount = sdk.NewCoin(depositor.WithdrawalAmount.Denom, sdk.ZeroInt())

	// burn the nft if it is existed
	for _, el := range depositor.LinkedNFT {
		ids := strings.Split(el, ":")
		_, found := k.nftKeeper.GetNFT(ctx, ids[0], ids[1])
		if !found {
			continue
		}
		err := k.nftKeeper.Burn(ctx, ids[0], ids[1])
		if err != nil {
			return &types.MsgWithdrawPrincipalResponse{}, coserrors.Wrapf(err, "burn nft failed")
		}
	}
	k.SetDepositorHistory(ctx, depositor)
	k.DelDepositor(ctx, depositor)
	poolInfo.ProcessedTransferAccounts = deleteElement(poolInfo.ProcessedTransferAccounts, depositor.DepositorAddress)
	poolInfo.ProcessedWithdrawAccounts = deleteElement(poolInfo.ProcessedWithdrawAccounts, depositor.DepositorAddress)

	if k.isEmptyPool(ctx, poolInfo) {
		totalReturn := poolInfo.EscrowPrincipalAmount.AddAmount(poolInfo.EscrowInterestAmount)
		if !totalReturn.IsZero() {
			err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccount, poolInfo.OwnerAddress, sdk.NewCoins(totalReturn))
			if err != nil {
				return nil, err
			}
		}

		// we handle the leftover of each class
		leftover := k.handleClassLeftover(ctx, poolInfo)
		reserve, found := k.GetReserve(ctx, "ausdc")
		if found {
			reserve = reserve.Add(leftover)
			k.SetReserve(ctx, reserve)
		}

		k.DelPool(ctx, poolInfo.Index)
		k.SetHistoryPool(ctx, poolInfo)

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeWithdrawPrincipal,
			),
		)

	} else {
		k.SetPool(ctx, poolInfo)
	}
	return &types.MsgWithdrawPrincipalResponse{Amount: amountToSend.String()}, nil
}

func (k msgServer) WithdrawPrincipal(goCtx context.Context, msg *types.MsgWithdrawPrincipal) (*types.MsgWithdrawPrincipalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	investor, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address %v", msg.Creator)
	}

	depositor, found := k.GetDepositor(ctx, msg.PoolIndex, investor)
	if !found {
		return nil, coserrors.Wrapf(types.ErrDepositorNotFound, "depositor not found for pool %v", msg.PoolIndex)
	}

	if depositor.WithdrawalAmount.Denom != msg.Token.Denom {
		return nil, coserrors.Wrapf(types.ErrInconsistencyToken, "you can only withdraw %v", depositor.WithdrawalAmount.Denom)
	}

	if msg.Token.IsZero() {
		return nil, errors.New("zero amount to withdraw")
	}

	poolInfo, found := k.GetPools(ctx, msg.PoolIndex)
	if !found {
		return nil, errors.New("pool cannot be found")
	}

	if poolInfo.CurrentPoolTotalBorrowCounter == 0 && ctx.BlockTime().Before(poolInfo.PoolCreatedTime.Add(time.Second*time.Duration(poolInfo.PoolLockedSeconds)+poolInfo.GraceTime)) {
		return nil, types.ErrPoolWithdrawLocked
	}

	totalWithdraw := msg.Token
	if msg.Token.IsGTE(depositor.GetWithdrawalAmount()) {
		totalWithdraw = depositor.GetWithdrawalAmount()
	}

	if poolInfo.CurrentPoolTotalBorrowCounter == 0 && ctx.BlockTime().After(poolInfo.PoolCreatedTime.Add(time.Second*time.Duration(poolInfo.PoolLockedSeconds)+poolInfo.GraceTime)) {
		poolInfo.PoolStatus = types.PoolInfo_FROZEN
	}

	if poolInfo.PoolStatus == types.PoolInfo_Liquidation {
		tokenSend, err := k.handlerPoolLiquidation(ctx, depositor)
		if err != nil {
			return nil, err
		}
		return &types.MsgWithdrawPrincipalResponse{Amount: tokenSend.String()}, nil
	}

	if poolInfo.PoolStatus == types.PoolInfo_FROZEN {
		if depositor.DepositType == types.DepositorInfo_deposit_close {
			ret, err := k.handleDepositClose(ctx, depositor, poolInfo)
			return ret, err
		}

		tokenSend, err := k.handlerPoolClose(ctx, poolInfo, depositor)
		if err != nil {
			return nil, err
		}
		return &types.MsgWithdrawPrincipalResponse{Amount: tokenSend.String()}, nil
	}

	switch depositor.DepositType {
	case types.DepositorInfo_deposit_close:
		ret, err := k.handleDepositClose(ctx, depositor, poolInfo)
		return ret, err

	case types.DepositorInfo_unset, types.DepositorInfo_withdraw_proposal, types.DepositorInfo_processed:
		if depositor.DepositType == types.DepositorInfo_unset {
			poolInfo.UsableAmount = poolInfo.UsableAmount.SubAmount(totalWithdraw.Amount)
		}
		if depositor.DepositType == types.DepositorInfo_processed {
			poolInfo.UsableAmount = poolInfo.UsableAmount.Add(depositor.WithdrawalAmount).SubAmount(totalWithdraw.Amount)
		}
		depositor.DepositType = types.DepositorInfo_unset
		depositor.WithdrawalAmount, err = depositor.WithdrawalAmount.SafeSub(totalWithdraw)
		if err != nil {
			return nil, errors.New("withdraw amount too large")
		}
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccount, investor, sdk.NewCoins(totalWithdraw))
		if err != nil {
			return nil, err
		}

		poolInfo.ProcessedTransferAccounts = deleteElement(poolInfo.ProcessedTransferAccounts, depositor.DepositorAddress)

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeWithdrawPrincipal,
				sdk.NewAttribute(types.AttributeCreator, msg.Creator),
				sdk.NewAttribute(types.AttributeAmount, totalWithdraw.String()),
			),
		)

		k.SetDepositor(ctx, depositor)
		k.SetPool(ctx, poolInfo)
		return &types.MsgWithdrawPrincipalResponse{Amount: totalWithdraw.String()}, nil
	default:
		return &types.MsgWithdrawPrincipalResponse{}, coserrors.Wrapf(types.ErrDeposit, "deposit type is %v", depositor.DepositType)
	}
}
