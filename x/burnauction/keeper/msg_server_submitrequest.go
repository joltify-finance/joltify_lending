package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/joltify-finance/joltify_lending/x/burnauction/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Submitrequest(goCtx context.Context, msg *types.MsgSubmitrequest) (*types.MsgSubmitrequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	burnCoins, err := sdk.ParseCoinsNormalized(msg.Tokens)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidCoin, "invalid coins %v", err)
	}

	sender, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(errors.ErrInvalidAddress, "invalid address %v", sender)
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, burnCoins)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrTransfer, "invalid coins %v", err)
	}

	return &types.MsgSubmitrequestResponse{}, nil
}
