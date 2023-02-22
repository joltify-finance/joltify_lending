package keeper

import (
	"context"
	"strings"

	coserrors "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k msgServer) ClaimInterest(goCtx context.Context, msg *types.MsgClaimInterest) (*types.MsgClaimInterestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	investorAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address %v", msg.Creator)
	}

	depositor, found := k.GetDepositor(ctx, msg.PoolIndex, investorAddress)
	if !found {
		return nil, coserrors.Wrapf(types.ErrDepositorNotFound, "depositor %v not found for pool index %v", msg.Creator, msg.GetPoolIndex())
	}

	if !depositor.DepositorAddress.Equals(investorAddress) {
		return nil, coserrors.Wrap(types.ErrUnauthorized, "not the depositer")
	}

	poolInfo, found := k.GetPools(ctx, depositor.PoolIndex)
	if !found {
		panic("should never fail to find the pool")
	}

	// for each lending NFT this owner has
	totalInterest, err := calculateTotalInterest(ctx, depositor.LinkedNFT, k.nftKeeper, poolInfo.ReserveFactor, true)
	if err != nil {
		return nil, err
	}

	claimed := sdk.NewCoin(depositor.LockedAmount.Denom, totalInterest)

	// we pay the interesting and the principle
	if poolInfo.PoolStatus == types.PoolInfo_CLOSING {
		lendNFTs := depositor.LinkedNFT

		totalBorrowedNow, err := calculateTotalPrinciple(ctx, lendNFTs, k.nftKeeper)
		if err != nil {
			return nil, err
		}
		claimed = claimed.AddAmount(totalBorrowedNow)

		/// we do the cleanup
		k.cleanupDepositor(ctx, depositor)
		k.cleanupPool(ctx, poolInfo)
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleAccount, investorAddress, sdk.NewCoins(claimed))
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeClaimInterest,
			sdk.NewAttribute(types.AttributeCreator, msg.Creator),
			sdk.NewAttribute("amount", claimed.String()),
		),
	)

	return &types.MsgClaimInterestResponse{}, nil
}

func (k Keeper) cleanupDepositor(ctx sdk.Context, depositor types.DepositorInfo) error {

	for _, el := range depositor.LinkedNFT {
		ids := strings.Split(el, ":")

		acc := k.accKeeper.GetModuleAddress(types.ModuleAccount)

		err := k.nftKeeper.Transfer(ctx, ids[0], ids[1], acc)
		if err != nil {
			return types.ErrTransferNFT
		}

		err = k.nftKeeper.Burn(ctx, ids[0], ids[1])
		if err != nil {
			return types.ErrBurnNFT
		}
	}

	// delete the deposit
	k.DelDepositor(ctx, depositor.PoolIndex, depositor.DepositorAddress)
	// delete the class
	return nil
}

func (k Keeper) cleanupPool(ctx sdk.Context, poolInfo types.PoolInfo) {
	emptyClassCount := 0
	for _, el := range poolInfo.PoolNFTIds[1:] {
		n := k.nftKeeper.GetTotalSupply(ctx, el)
		if n == 0 {
			emptyClassCount++
		}
	}
	if len(poolInfo.PoolNFTIds) == emptyClassCount+1 {
		ctx.Logger().Info("move the pool to history pool", "pool index", poolInfo.Index)
		k.SetHistoryPool(ctx, poolInfo)
		k.DelPool(ctx, poolInfo.Index)
	}
	return
}
