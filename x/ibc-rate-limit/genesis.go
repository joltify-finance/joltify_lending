package ibc_rate_limit

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/x/ibc-rate-limit/types"
)

// InitGenesis initializes the x/ibc-rate-limit module's state from a provided genesis
func (i *ICS4Wrapper) InitGenesis(ctx context.Context, genState types.GenesisState) {
}

// ExportGenesis returns the x/ibc-rate-limit module's exported genesis.
func (i *ICS4Wrapper) ExportGenesis(ctx context.Context) *types.GenesisState {
	return &types.GenesisState{}
}
