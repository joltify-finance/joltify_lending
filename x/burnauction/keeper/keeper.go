package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/burnauction/types"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc           codec.BinaryCodec
		storeKey      storetypes.StoreKey
		memKey        storetypes.StoreKey
		paramstore    paramtypes.Subspace
		accKeeper     types.AccountKeeper
		bankKeeper    types.BankKeeper
		auctionKeeper types.AuctionKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	accKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	auctionKeeper types.AuctionKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		accKeeper:     accKeeper,
		bankKeeper:    bankKeeper,
		auctionKeeper: auctionKeeper,
	}
}

func (k Keeper) Logger(rctx context.Context) log.Logger {
	ctx := sdk.UnwrapSDKContext(rctx)
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
