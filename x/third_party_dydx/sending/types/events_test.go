package types_test

import (
	"testing"

	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/testutil/dydx/testutil/constants"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/sending/types"
	"github.com/stretchr/testify/require"
)

func TestNewCreateTransferEvent(t *testing.T) {
	sender := constants.Alice_Num1
	receiver := constants.Bob_Num2
	quantums := uint64(100000000)
	assetId := uint32(1)

	event := types.NewCreateTransferEvent(sender, receiver, assetId, quantums)
	require.Equal(t, event.Type, types.EventTypeCreateTransfer)
	require.Equal(t, event.Attributes, []abci.EventAttribute{
		{
			Key:   types.AttributeKeySender,
			Value: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0",
		},
		{
			Key:   types.AttributeKeySenderNumber,
			Value: "1",
		},
		{
			Key:   types.AttributeKeyRecipient,
			Value: "jolt14jux2kfgelja5394dxquqn3wh974psqzv4hzgg",
		},
		{
			Key:   types.AttributeKeyRecipientNumber,
			Value: "2",
		},
		{
			Key:   types.AttributeKeyAssetId,
			Value: "1",
		},
		{
			Key:   types.AttributeKeyQuantums,
			Value: "100000000",
		},
	})
}

func TestNewDepositToSubaccountEvent(t *testing.T) {
	sender := sdk.MustAccAddressFromBech32(constants.Alice_Num1.Owner)
	receiver := constants.Bob_Num2
	quantums := uint64(100000000)
	assetId := uint32(1)

	event := types.NewDepositToSubaccountEvent(sender, receiver, assetId, quantums)
	require.Equal(t, event.Type, types.EventTypeDepositToSubaccount)
	require.Equal(t, event.Attributes, []abci.EventAttribute{
		{
			Key:   types.AttributeKeySender,
			Value: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0",
		},
		{
			Key:   types.AttributeKeyRecipient,
			Value: "jolt14jux2kfgelja5394dxquqn3wh974psqzv4hzgg",
		},
		{
			Key:   types.AttributeKeyRecipientNumber,
			Value: "2",
		},
		{
			Key:   types.AttributeKeyAssetId,
			Value: "1",
		},
		{
			Key:   types.AttributeKeyQuantums,
			Value: "100000000",
		},
	})
}

func TestNewWithdrawFromSubaccountEvent(t *testing.T) {
	sender := constants.Alice_Num1
	receiver := sdk.MustAccAddressFromBech32(constants.Bob_Num2.Owner)
	quantums := uint64(100000000)
	assetId := uint32(1)

	event := types.NewWithdrawFromSubaccountEvent(sender, receiver, assetId, quantums)
	require.Equal(t, event.Type, types.EventTypeWithdrawFromSubaccount)
	require.Equal(t, event.Attributes, []abci.EventAttribute{
		{
			Key:   types.AttributeKeySender,
			Value: "jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0",
		},
		{
			Key:   types.AttributeKeySenderNumber,
			Value: "1",
		},
		{
			Key:   types.AttributeKeyRecipient,
			Value: "jolt14jux2kfgelja5394dxquqn3wh974psqzv4hzgg",
		},
		{
			Key:   types.AttributeKeyAssetId,
			Value: "1",
		},
		{
			Key:   types.AttributeKeyQuantums,
			Value: "100000000",
		},
	})
}
