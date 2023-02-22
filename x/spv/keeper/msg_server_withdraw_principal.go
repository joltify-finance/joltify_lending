package keeper

import (
	"context"
	coserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

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

	poolInfo, found := k.GetPools(ctx, msg.GetPoolIndex())
	if !found {
		return nil, coserrors.Wrapf(sdkerrors.ErrNotFound, "pool cannot be found %v", msg.GetPoolIndex())
	}

	_ = poolInfo
	if depositor.WithdrawalAmount.Denom != msg.Token.Denom {
		return nil, coserrors.Wrapf(types.ErrInconsistencyToken, "you can only withdraw %v", depositor.WithdrawalAmount.Denom)
	}

	lendNFTs := depositor.LinkedNFT

	totalBorrowedNow, err := calculateTotalPrinciple(ctx, lendNFTs, k.nftKeeper)
	if err != nil {
		return nil, err
	}

	depositor.LockedAmount = depositor.LockedAmount.SubAmount(totalBorrowedNow)
	depositor.WithdrawalAmount = depositor.WithdrawalAmount.AddAmount(totalBorrowedNow).Sub(msg.Token)

	// todo we need to delete the nft once the
	//if depositor.WithdrawalAmount.Amount.Equal(sdk.ZeroInt()) {
	//	// we burn the nft
	//	for _, el := range lendNFTs {
	//		ids := strings.Split(el, ":")
	//		err = k.nftKeeper.Transfer(ctx, ids[0], ids[1], poolInfo.OwnerAddress)
	//		if err != nil {
	//			return &types.MsgWithdrawPrincipalResponse{}, types.ErrTransferNFT
	//		}
	//
	//		err = k.nftKeeper.Burn(ctx, ids[0], ids[1])
	//		if err != nil {
	//			return &types.MsgWithdrawPrincipalResponse{}, types.ErrBurnNFT
	//		}
	//
	//	}
	//}

	k.SetDepositor(ctx, depositor)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeWithdrawPrincipal,
			sdk.NewAttribute(types.AttributeCreator, msg.Creator),
			sdk.NewAttribute(types.AttributeAmount, msg.Token.String()),
		),
	)

	return &types.MsgWithdrawPrincipalResponse{}, nil
}
