package keeper

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

// Implements SwapHooks interface
var _ types.SPVHooks = Keeper{}

// AfterSPVInterestPaid - call hook if registered to update the pool's incentives
func (k Keeper) AfterSPVInterestPaid(ctx sdk.Context, poolID string, amt sdkmath.Int) {
	if k.hooks != nil {
		k.hooks.AfterSPVInterestPaid(ctx, poolID, amt)
	}
}

func (k Keeper) BeforeNFTBurned(ctx sdk.Context, poolIndex string, walletAddress string, linkednfts []string) error {
	if k.hooks != nil {
		return k.hooks.BeforeNFTBurned(ctx, poolIndex, walletAddress, linkednfts)
	}
	return nil
}
