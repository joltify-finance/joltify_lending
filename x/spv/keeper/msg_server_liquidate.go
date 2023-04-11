package keeper

import (
	"context"
	coserrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	types2 "github.com/cosmos/cosmos-sdk/codec/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k Keeper) doUpdateLiquidationInfo(ctx sdk.Context, el string, amountFromLiquidator, totalPoolBorrowed sdk.Coin, paidAmount sdkmath.Int) (sdkmath.Int, error) {
	class, found := k.nftKeeper.GetClass(ctx, el)
	if !found {
		panic(found)
	}
	var borrowInterest types.BorrowInterest
	var err error
	err = proto.Unmarshal(class.Data.Value, &borrowInterest)
	if err != nil {
		panic(err)
	}

	lastBorrow := borrowInterest.BorrowDetails[len(borrowInterest.BorrowDetails)-1].BorrowedAmount
	if paidAmount.IsZero() {
		paidAmount = sdk.NewDecFromInt(amountFromLiquidator.Amount).Mul(sdk.NewDecFromInt(lastBorrow.Amount)).Quo(sdk.NewDecFromInt(totalPoolBorrowed.Amount)).TruncateInt()
	}

	if borrowInterest.LiquidationItems == nil {
		borrowInterest.LiquidationItems = make([]*types.LiquidationItem, 1, 10)
		borrowInterest.LiquidationItems[0] = &types.LiquidationItem{
			Amount:                 sdk.NewCoin(amountFromLiquidator.Denom, paidAmount),
			LiquidationPaymentTime: ctx.BlockTime(),
		}
	} else {
		borrowInterest.LiquidationItems = append(borrowInterest.LiquidationItems, &types.LiquidationItem{
			Amount:                 sdk.NewCoin(amountFromLiquidator.Denom, paidAmount),
			LiquidationPaymentTime: ctx.BlockTime(),
		})
	}

	class.Data, err = types2.NewAnyWithValue(&borrowInterest)
	if err != nil {
		panic("pack class any data failed")
	}
	err = k.nftKeeper.UpdateClass(ctx, class)
	if err != nil {
		return sdk.ZeroInt(), err
	}

	return paidAmount, nil

}

func (k Keeper) handleLiquidation(ctx sdk.Context, poolInfo types.PoolInfo, amount sdk.Coin) error {

	nftClasses := poolInfo.PoolNFTIds
	totalBorrowed := poolInfo.BorrowedAmount
	// the first element is the pool class, we skip it
	totalPaid := sdk.ZeroInt()
	for i, el := range nftClasses {
		if i == 0 {
			// we will handle the fist element later
			continue
		}
		amountPaid, err := k.doUpdateLiquidationInfo(ctx, el, amount, totalBorrowed, sdk.ZeroInt())
		if err != nil {
			return err
		}
		totalPaid = totalPaid.Add(amountPaid)
	}

	if len(nftClasses) > 0 {
		paidAmount := amount.Amount.Sub(totalPaid)
		_, err := k.doUpdateLiquidationInfo(ctx, nftClasses[0], amount, totalBorrowed, paidAmount)
		if err != nil {
			return err
		}
	}
	poolInfo.TotalLiquidationAmount = poolInfo.TotalLiquidationAmount.Add(amount.Amount)
	k.SetPool(ctx, poolInfo)

	return nil
}

func (k msgServer) Liquidate(goCtx context.Context, msg *types.MsgLiquidate) (*types.MsgLiquidateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	liquidator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address %v", msg.Creator)
	}

	if msg.Amount.IsZero() {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "the amount cannot be zero")
	}

	poolInfo, found := k.GetPools(ctx, msg.PoolIndex)
	if !found {
		return nil, coserrors.Wrapf(types.ErrPoolNotFound, "pool cannot be found %v", msg.PoolIndex)
	}

	if msg.Amount.Denom != poolInfo.BorrowedAmount.Denom {
		return nil, coserrors.Wrapf(types.ErrInconsistencyToken, "the token is not the same as the borrowed token %v", msg.Amount.Denom)
	}

	if poolInfo.PoolStatus != types.PoolInfo_Liquidation {
		return nil, coserrors.Wrapf(types.ErrPoolNotInLiquidation, "pool is not in liquidation %v", poolInfo.PoolStatus)
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccount, liquidator, sdk.NewCoins(msg.Amount))

	if err != nil {
		return nil, coserrors.Wrapf(types.ErrLiquidation, "liquidation failed %v", err)
	}

	err = k.handleLiquidation(ctx, poolInfo, msg.Amount)
	if err != nil {
		return nil, coserrors.Wrapf(types.ErrLiquidation, "liquidation failed %v", err)
	}

	return &types.MsgLiquidateResponse{}, nil
}
