package keeper

import (
	"context"
	coserrors "cosmossdk.io/errors"
	types2 "github.com/cosmos/cosmos-sdk/codec/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func (k Keeper) travelThoughPrincipalToBePaid(ctx sdk.Context, poolInfo *types.PoolInfo, amountToPay sdk.Coin) {

	nftClasses := poolInfo.PoolNFTIds
	// the first element is the pool class, we skip it
	totalBorrowedAmount := poolInfo.BorrowedAmount
	borrowTimes := len(nftClasses[1:])
	for index, el := range nftClasses[1:] {
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

		if index == borrowTimes {
			borrowInterest.Borrowed = borrowInterest.Borrowed.Sub(amountToPay)
			class.Data, err = types2.NewAnyWithValue(&borrowInterest)
			if err != nil {
				panic("pack class any data failed")
			}
			k.nftKeeper.SaveClass(ctx, class)
			return
		}

		ratioOfThisBorrow := sdk.NewDecFromInt(borrowInterest.Borrowed.Amount).Quo(sdk.NewDecFromInt(totalBorrowedAmount.Amount))
		thisPayAmount := sdk.NewDecFromInt(borrowInterest.Borrowed.Amount).Mul(ratioOfThisBorrow).TruncateInt()
		borrowInterest.Borrowed = borrowInterest.Borrowed.SubAmount(thisPayAmount)
		k.nftKeeper.SaveClass(ctx, class)
	}
}

func (k msgServer) PayPrincipal(goCtx context.Context, msg *types.MsgPayPrincipal) (*types.MsgPayPrincipalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

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

	if msg.Token.IsGTE(poolInfo.BorrowedAmount) {
		return nil, coserrors.Wrapf(sdkerrors.ErrInvalidRequest, "extra principal rejected")
	}

	if msg.Token.Denom != poolInfo.TotalAmount.Denom {
		return nil, coserrors.Wrapf(types.ErrInconsistencyToken, "pool denom %v and repay is %v", poolInfo.TotalAmount.Denom, msg.Token.Denom)
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
