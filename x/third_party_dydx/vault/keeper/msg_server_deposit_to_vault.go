package keeper

import (
	"context"

	"github.com/joltify-finance/joltify_lending/dydx_helper/lib"
	"github.com/joltify-finance/joltify_lending/dydx_helper/lib/log"
	"github.com/joltify-finance/joltify_lending/dydx_helper/lib/metrics"
	assettypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/assets/types"
	sendingtypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/sending/types"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/vault/types"
)

// DepositToVault deposits from a subaccount to a vault.
func (k msgServer) DepositToVault(
	goCtx context.Context,
	msg *types.MsgDepositToVault,
) (*types.MsgDepositToVaultResponse, error) {
	ctx := lib.UnwrapSDKContext(goCtx, types.ModuleName)
	quoteQuantums := msg.QuoteQuantums.BigInt()

	// Mint shares for the vault.
	err := k.MintShares(
		ctx,
		*msg.VaultId,
		msg.SubaccountId.Owner,
		quoteQuantums,
	)
	if err != nil {
		return nil, err
	}

	// Add vault to address store.
	k.AddVaultToAddressStore(ctx, *msg.VaultId)

	// Transfer from sender subaccount to vault.
	// Note: Transfer should take place after minting shares for
	// shares calculation to be correct.
	err = k.sendingKeeper.ProcessTransfer(
		ctx,
		&sendingtypes.Transfer{
			Sender:    *msg.SubaccountId,
			Recipient: *msg.VaultId.ToSubaccountId(),
			AssetId:   assettypes.AssetUsdc.Id,
			Amount:    msg.QuoteQuantums.BigInt().Uint64(),
		},
	)
	if err != nil {
		return nil, err
	}

	// Emit metric on vault equity.
	equity, err := k.GetVaultEquity(ctx, *msg.VaultId)
	if err != nil {
		log.ErrorLogWithError(ctx, "Failed to get vault equity", err, "vaultId", *msg.VaultId)
	} else {
		msg.VaultId.SetGaugeWithLabels(
			metrics.VaultEquity,
			float32(equity.Int64()),
		)
	}

	return &types.MsgDepositToVaultResponse{}, nil
}
