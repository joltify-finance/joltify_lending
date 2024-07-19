package mint

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/joltify-finance/joltify_lending/x/mint/keeper"
	"github.com/joltify-finance/joltify_lending/x/mint/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	// this line is used by starport scaffolding # handler/msgServer
	return func(ctx context.Context, msg sdk.Msg) (*sdk.Result, error) {
		// this line is used by starport scaffolding # 1
		errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
		return nil, errorsmod.Wrap(errorsmod.ErrUnknownRequest, errMsg)
	}
}
