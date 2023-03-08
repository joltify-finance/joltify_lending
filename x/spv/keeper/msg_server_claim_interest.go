package keeper

import (
	"context"
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
		return nil, coserrors.Wrap(types.ErrUnauthorized, "not the depositor")
	}

	if depositor.DepositType == types.DepositorInfo_deposit_close {
		return nil, coserrors.Wrapf(types.ErrUnauthorized, "your deposit has been closed")
	}

	// for each lending NFT this owner has
	totalInterest, err := calculateTotalInterest(ctx, depositor.LinkedNFT, k.nftKeeper, true)
	if err != nil {
		return nil, err
	}

	claimed := sdk.NewCoin(depositor.LockedAmount.Denom, totalInterest)

	// we pay the interesting and the principle
	//if poolInfo.PoolStatus == types.PoolInfo_CLOSING {
	//	lendNFTs := depositor.LinkedNFT
	//
	//	totalBorrowedNow, err := calculateTotalPrinciple(ctx, lendNFTs, k.nftKeeper)
	//	if err != nil {
	//		return nil, err
	//	}
	//	claimed = claimed.AddAmount(totalBorrowedNow)
	//
	//	/// we do the cleanup
	//	k.cleanupDepositor(ctx, depositor)
	//	k.cleanupPool(ctx, poolInfo)
	//}

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

	return &types.MsgClaimInterestResponse{Amount: claimed.String()}, nil
}
