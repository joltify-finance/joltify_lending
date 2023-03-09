package keeper

import (
	"context"
	"errors"

	coserrors "cosmossdk.io/errors"
	types2 "github.com/cosmos/cosmos-sdk/codec/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k Keeper) travelThoughPrincipalToBePaid(ctx sdk.Context, poolInfo *types.PoolInfo, amountToPay sdk.Coin) {

	nftClasses := poolInfo.PoolNFTIds
	if len(nftClasses) == 0 {
		ctx.Logger().Info("do not have any borrow record")
		return
	}
	totalBorrowedAmount := poolInfo.BorrowedAmount
	currentPayout := sdk.ZeroInt()
	for _, el := range nftClasses {
		class, found := k.nftKeeper.GetClass(ctx, el)
		if !found {
			panic(found)
		}
		var borrowClassInfo types.BorrowInterest
		var err error
		err = proto.Unmarshal(class.Data.Value, &borrowClassInfo)
		if err != nil {
			panic(err)
		}
		latestBorrow := borrowClassInfo.BorrowDetails[len(borrowClassInfo.BorrowDetails)-1].BorrowedAmount
		ratioOfThisBorrow := sdk.NewDecFromInt(latestBorrow.Amount).Quo(sdk.NewDecFromInt(totalBorrowedAmount.Amount))
		thisPayAmount := sdk.NewDecFromInt(amountToPay.Amount).Mul(ratioOfThisBorrow).TruncateInt()
		borrow := types.BorrowDetail{BorrowedAmount: latestBorrow.SubAmount(thisPayAmount), TimeStamp: ctx.BlockTime()}
		borrowClassInfo.BorrowDetails = append(borrowClassInfo.BorrowDetails, borrow)
		currentPayout = currentPayout.Add(thisPayAmount)
		class.Data, err = types2.NewAnyWithValue(&borrowClassInfo)
		if err != nil {
			panic("pack class any data failed")
		}

		err = k.nftKeeper.UpdateClass(ctx, class)
		if err != nil {
			panic(err)
		}
	}

	firstClass, found := k.nftKeeper.GetClass(ctx, nftClasses[0])
	if !found {
		panic(found)
	}
	var borrowClassInfo types.BorrowInterest
	var err error
	err = proto.Unmarshal(firstClass.Data.Value, &borrowClassInfo)
	if err != nil {
		panic(err)
	}

	latestBorrow := borrowClassInfo.BorrowDetails[len(borrowClassInfo.BorrowDetails)-1].BorrowedAmount
	borrow := types.BorrowDetail{BorrowedAmount: latestBorrow.Sub(amountToPay.SubAmount(currentPayout)), TimeStamp: ctx.BlockTime()}
	borrowClassInfo.BorrowDetails = append(borrowClassInfo.BorrowDetails, borrow)
	firstClass.Data, err = types2.NewAnyWithValue(&borrowClassInfo)
	if err != nil {
		panic("pack class any data failed")
	}
	err = k.nftKeeper.UpdateClass(ctx, firstClass)
	if err != nil {
		panic(err)
	}

}

func (k msgServer) PayPrincipalPartial(goCtx context.Context, msg *types.MsgPayPrincipal) (*types.MsgPayPrincipalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	return nil, errors.New("not supported")

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address %v", msg.Creator)
	}

	poolInfo, found := k.GetPools(ctx, msg.GetPoolIndex())
	if !found {
		return nil, coserrors.Wrapf(sdkerrors.ErrNotFound, "pool cannot be found %v", msg.GetPoolIndex())
	}

	if msg.Token.Denom != poolInfo.BorrowedAmount.Denom {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid token demo, want %v", poolInfo.BorrowedAmount.Denom)
	}

	poolInfo.EscrowPrincipalAmount = poolInfo.EscrowPrincipalAmount.Add(msg.Token)
	k.SetPool(ctx, poolInfo)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypePayPrincipal,
			sdk.NewAttribute(types.AttributeCreator, msg.Creator),
			sdk.NewAttribute(types.AttributeAmount, msg.Token.String()),
		),
	)

	return &types.MsgPayPrincipalResponse{}, nil
}
