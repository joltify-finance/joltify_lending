package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/joltify-finance/joltify_lending/dydx_helper/lib"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/perpetuals/types"
)

func (k msgServer) SetLiquidityTier(
	goCtx context.Context,
	msg *types.MsgSetLiquidityTier,
) (*types.MsgSetLiquidityTierResponse, error) {
	if !k.Keeper.HasAuthority(msg.Authority) {
		return nil, errorsmod.Wrapf(
			govtypes.ErrInvalidSigner,
			"invalid authority %s",
			msg.Authority,
		)
	}

	ctx := lib.UnwrapSDKContext(goCtx, types.ModuleName)

	if _, err := k.Keeper.SetLiquidityTier(
		ctx,
		msg.LiquidityTier.Id,
		msg.LiquidityTier.Name,
		msg.LiquidityTier.InitialMarginPpm,
		msg.LiquidityTier.MaintenanceFractionPpm,
		msg.LiquidityTier.ImpactNotional,
		msg.LiquidityTier.OpenInterestLowerCap,
		msg.LiquidityTier.OpenInterestUpperCap,
	); err != nil {
		return nil, err
	}

	return &types.MsgSetLiquidityTierResponse{}, nil
}
