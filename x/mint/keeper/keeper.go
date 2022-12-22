package keeper

import (
	"fmt"

	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/joltify-finance/joltify_lending/x/mint/types"
)

type (
	Keeper struct {
		cdc                codec.BinaryCodec
		storeKey           sdk.StoreKey
		paramstore         paramtypes.Subspace
		accountKeeper      types.AccountKeeper
		bankKeeper         types.BankKeeper
		distributionKeeper types.DistributionKeeper
		stakingKeeper      stakingkeeper.Keeper
		feeCollectorName   string
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
	ps paramtypes.Subspace,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	distributionKeeper types.DistributionKeeper,
	stakingKeeper stakingkeeper.Keeper,
	feeCollectorName string,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:                cdc,
		storeKey:           storeKey,
		paramstore:         ps,
		accountKeeper:      accountKeeper,
		bankKeeper:         bankKeeper,
		distributionKeeper: distributionKeeper,
		stakingKeeper:      stakingKeeper,
		feeCollectorName:   feeCollectorName,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
