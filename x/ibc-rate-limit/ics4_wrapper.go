package ibc_rate_limit

import (
	"context"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	porttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	"github.com/cosmos/ibc-go/v7/modules/core/exported"

	"github.com/joltify-finance/joltify_lending/x/ibc-rate-limit/types"
)

var _ porttypes.ICS4Wrapper = &ICS4Wrapper{}

type ICS4Wrapper struct {
	channel       porttypes.ICS4Wrapper
	accountKeeper *authkeeper.AccountKeeper
	bankKeeper    *bankkeeper.Keeper
	quotaKeeper   types.QuotaKeeper
	paramSpace    paramtypes.Subspace
}

func (i *ICS4Wrapper) GetAppVersion(ctx sdk.Context, portID, channelID string) (string, bool) {
	return i.channel.GetAppVersion(ctx, portID, channelID)
}

func NewICS4Middleware(
	channel porttypes.ICS4Wrapper,
	accountKeeper *authkeeper.AccountKeeper,
	bankKeeper *bankkeeper.Keeper, quotaKeeper types.QuotaKeeper, paramSpace paramtypes.Subspace,
) ICS4Wrapper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}
	return ICS4Wrapper{
		channel:       channel,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		paramSpace:    paramSpace,
		quotaKeeper:   quotaKeeper,
	}
}

// SendPacket implements the ICS4 interface and is called when sending packets.
func (i *ICS4Wrapper) SendPacket(ctx sdk.Context, chanCap *capabilitytypes.Capability, sourcePort, sourceChannel string, timeoutHeight clienttypes.Height, timeoutTimestamp uint64, data []byte) (uint64, error) {
	seq, errSend := i.channel.SendPacket(ctx, chanCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data)
	if errSend != nil {
		return seq, errSend
	}

	if sourcePort == "transfer" {
		onlist, errCheck := i.whecherOnWhitelist(ctx, data)
		if errCheck != nil {
			ctx.Logger().Error("fail to check the whitelist", "transferInfo", string(data), "reason", errCheck.Error())
			return seq, nil
		}
		// the sender is on whitelist, so bypass
		if onlist {
			return seq, nil
		}

		errUpdate := i.UpdateQuota(ctx, seq, data)
		if errUpdate != nil {
			if errUpdate.Error() == "quota exceeded" {
				ctx.Logger().Error("quota exceeded", "transferInfo", string(data), "reason", errUpdate.Error())
				return 0, errorsmod.Wrapf(errUpdate, "rate limit SendPacket failed to authorize transfer")
			}
			ctx.Logger().Error("fail to update the quota", "transferInfo", string(data), "reason", errUpdate.Error())
			return seq, nil
		}
	}
	return seq, nil
}

func (i *ICS4Wrapper) WriteAcknowledgement(ctx sdk.Context, chanCap *capabilitytypes.Capability, packet exported.PacketI, ack exported.Acknowledgement) error {
	return i.channel.WriteAcknowledgement(ctx, chanCap, packet, ack)
}

func (i *ICS4Wrapper) GetParams(ctx sdk.Context) (params types.Params) {
	// This was previously done via i.paramSpace.GetParamSet(ctx, &params). That will
	// panic if the params don't exist. This is a workaround to avoid that panic.
	// Params should be refactored to just use a raw kvstore.
	for _, pair := range params.ParamSetPairs() {
		i.paramSpace.GetIfExists(ctx, pair.Key, pair.Value)
	}
	if params.TokenQuota == "" {
		return types.DefaultParams()
	}
	return params
}

func (i *ICS4Wrapper) SetParams(ctx sdk.Context, params types.Params) {
	i.paramSpace.SetParamSet(ctx, &params)
}

func (i ICS4Wrapper) Params(goCtx context.Context,
	req *types.QueryParamsRequest,
) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := i.GetParams(ctx)
	return &types.QueryParamsResponse{Params: params}, nil
}
