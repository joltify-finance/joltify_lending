package ibc_rate_limit

import (
	"encoding/json"

	"github.com/joltify-finance/joltify_lending/x/ibc-rate-limit/types"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"

	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v8/modules/core/05-port/types"
	"github.com/cosmos/ibc-go/v8/modules/core/exported"
)

var _ porttypes.Middleware = &IBCMiddleware{}

type IBCMiddleware struct {
	app            porttypes.IBCModule
	ics4Middleware *ICS4Wrapper
}

func NewIBCMiddleware(app porttypes.IBCModule, ics4 *ICS4Wrapper) IBCMiddleware {
	return IBCMiddleware{
		app:            app,
		ics4Middleware: ics4,
	}
}

// OnChanOpenInit implements the IBCMiddleware interface
func (im *IBCMiddleware) OnChanOpenInit(ctx context.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	channelCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version string,
) (string, error) {
	return im.app.OnChanOpenInit(
		ctx,
		order,
		connectionHops,
		portID,
		channelID,
		channelCap,
		counterparty,
		version,
	)
}

// OnChanOpenTry implements the IBCMiddleware interface
func (im *IBCMiddleware) OnChanOpenTry(
	ctx context.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID,
	channelID string,
	channelCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	counterpartyVersion string,
) (string, error) {
	return im.app.OnChanOpenTry(ctx, order, connectionHops, portID, channelID, channelCap, counterparty, counterpartyVersion)
}

// OnChanOpenAck implements the IBCMiddleware interface
func (im *IBCMiddleware) OnChanOpenAck(
	ctx context.Context,
	portID,
	channelID string,
	counterpartyChannelID string,
	counterpartyVersion string,
) error {
	// Here we can add initial limits when a new channel is open. For now, they can be added manually on the contract
	return im.app.OnChanOpenAck(ctx, portID, channelID, counterpartyChannelID, counterpartyVersion)
}

// OnChanOpenConfirm implements the IBCMiddleware interface
func (im *IBCMiddleware) OnChanOpenConfirm(
	ctx context.Context,
	portID,
	channelID string,
) error {
	// Here we can add initial limits when a new channel is open. For now, they can be added manually on the contract
	return im.app.OnChanOpenConfirm(ctx, portID, channelID)
}

// OnChanCloseInit implements the IBCMiddleware interface
func (im *IBCMiddleware) OnChanCloseInit(
	ctx context.Context,
	portID,
	channelID string,
) error {
	// Here we can remove the limits when a new channel is closed. For now, they can remove them  manually on the contract
	return im.app.OnChanCloseInit(ctx, portID, channelID)
}

// OnChanCloseConfirm implements the IBCMiddleware interface
func (im *IBCMiddleware) OnChanCloseConfirm(
	ctx context.Context,
	portID,
	channelID string,
) error {
	// Here we can remove the limits when a new channel is closed. For now, they can remove them  manually on the contract
	return im.app.OnChanCloseConfirm(ctx, portID, channelID)
}

// OnRecvPacket implements the IBCMiddleware interface
func (im *IBCMiddleware) OnRecvPacket(
	ctx context.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) exported.Acknowledgement {
	// if this returns an Acknowledgement that isn't successful, all state changes are discarded
	return im.app.OnRecvPacket(ctx, packet, relayer)
}

// OnAcknowledgementPacket implements the IBCMiddleware interface
func (im *IBCMiddleware) OnAcknowledgementPacket(
	ctx context.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	var ack channeltypes.Acknowledgement
	if err := json.Unmarshal(acknowledgement, &ack); err != nil {
		if packet.SourcePort == types.PORTTYPE {
			im.ics4Middleware.RevokeQuotaHistory(ctx, packet.GetSequence())
		}
		return errorsmod.Wrapf(errorsmod.ErrUnknownRequest, "cannot unmarshal ICS-20 transfer packet acknowledgement: %v", err)
	}

	_, ok := ack.Response.(*channeltypes.Acknowledgement_Error)
	if ok {
		if packet.SourcePort == types.PORTTYPE {
			im.ics4Middleware.RevokeQuotaHistory(ctx, packet.GetSequence())
		}
	}

	return im.app.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
}

// OnTimeoutPacket implements the IBCMiddleware interface
func (im *IBCMiddleware) OnTimeoutPacket(
	ctx context.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	if packet.SourcePort == types.PORTTYPE {
		im.ics4Middleware.RevokeQuotaHistory(ctx, packet.GetSequence())
	}

	return im.app.OnTimeoutPacket(ctx, packet, relayer)
}

// RevertSentPacket Notifies the contract that a sent packet wasn't properly received
func (im *IBCMiddleware) RevertSentPacket(
	ctx context.Context,
	packet exported.PacketI,
) error {
	if packet.GetSourcePort() == types.PORTTYPE {
		im.ics4Middleware.RevokeQuotaHistory(ctx, packet.GetSequence())
	}
	return nil
}

// SendPacket implements the ICS4 Wrapper interface
func (im *IBCMiddleware) SendPacket(
	ctx context.Context,
	chanCap *capabilitytypes.Capability,
	sourcePort, sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
	data []byte,
) (uint64, error) {
	return im.ics4Middleware.SendPacket(ctx, chanCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data)
}

// WriteAcknowledgement implements the ICS4 Wrapper interface
func (im *IBCMiddleware) WriteAcknowledgement(
	ctx context.Context,
	chanCap *capabilitytypes.Capability,
	packet exported.PacketI,
	ack exported.Acknowledgement,
) error {
	return im.ics4Middleware.WriteAcknowledgement(ctx, chanCap, packet, ack)
}

func (im *IBCMiddleware) GetAppVersion(ctx context.Context, portID, channelID string) (string, bool) {
	return im.ics4Middleware.GetAppVersion(ctx, portID, channelID)
}
