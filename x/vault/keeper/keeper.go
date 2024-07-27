package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
	// this line is used by starport scaffolding # ibc/keeper/import
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeKey     storetypes.StoreKey
		memKey       storetypes.StoreKey
		vaultStaking types.VaultStaking
		bankKeeper   types.BankKeeper
		paramstore   paramtypes.Subspace
		ak           banktypes.AccountKeeper
		// this line is used by starport scaffolding # ibc/keeper/attribute
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	vaultStaking types.VaultStaking,
	bankKeeper types.BankKeeper,
	ps paramtypes.Subspace,
	ak banktypes.AccountKeeper,
	// this line is used by starport scaffolding # ibc/keeper/parameter
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:          cdc,
		storeKey:     storeKey,
		memKey:       memKey,
		vaultStaking: vaultStaking,
		bankKeeper:   bankKeeper,
		paramstore:   ps,
		ak:           ak,
		// this line is used by starport scaffolding # ibc/keeper/return

	}
}

func (k Keeper) Logger(rctx context.Context) log.Logger {
	ctx := sdk.UnwrapSDKContext(rctx)
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
