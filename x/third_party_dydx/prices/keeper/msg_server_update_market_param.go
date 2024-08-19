package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/telemetry"
	gometrics "github.com/hashicorp/go-metrics"
	pricefeedmetrics "github.com/joltify-finance/joltify_lending/daemons/pricefeed/metrics"
	"github.com/joltify-finance/joltify_lending/lib"
	"github.com/joltify-finance/joltify_lending/lib/metrics"

	errorsmod "cosmossdk.io/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/prices/types"
)

func (k msgServer) UpdateMarketParam(
	goCtx context.Context,
	msg *types.MsgUpdateMarketParam,
) (
	response *types.MsgUpdateMarketParamResponse,
	err error,
) {
	// Increment the appropriate success/error counter when the function finishes.
	defer func() {
		success := metrics.Success
		if err != nil {
			success = metrics.Error
		}
		telemetry.IncrCounterWithLabels(
			[]string{types.ModuleName, metrics.UpdateMarketParam, success},
			1,
			[]gometrics.Label{pricefeedmetrics.GetLabelForMarketId(msg.MarketParam.Id)},
		)
	}()

	if !k.Keeper.HasAuthority(msg.Authority) {
		return nil, errorsmod.Wrapf(
			govtypes.ErrInvalidSigner,
			"invalid authority %s",
			msg.Authority,
		)
	}

	ctx := lib.UnwrapSDKContext(goCtx, types.ModuleName)

	if _, err = k.Keeper.ModifyMarketParam(ctx, msg.MarketParam); err != nil {
		return nil, err
	}

	return &types.MsgUpdateMarketParamResponse{}, nil
}