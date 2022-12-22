package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/cdp/types"
)

// Implements StakingHooks interface
var _ types2.CDPHooks = Keeper{}

// AfterCDPCreated - call hook if registered
func (k Keeper) AfterCDPCreated(ctx sdk.Context, cdp types2.CDP) {
	if k.hooks != nil {
		k.hooks.AfterCDPCreated(ctx, cdp)
	}
}

// BeforeCDPModified - call hook if registered
func (k Keeper) BeforeCDPModified(ctx sdk.Context, cdp types2.CDP) {
	if k.hooks != nil {
		k.hooks.BeforeCDPModified(ctx, cdp)
	}
}
