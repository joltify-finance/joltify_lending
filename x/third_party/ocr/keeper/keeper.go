package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/x/third_party/ocr/types"
)

type Keeper struct {
	types.QueryServer

	OcrParams
	OcrConfig
	OcrReporting
	RewardPool
	FeedObservations
	FeedTransmissions
	OcrHooks

	bankKeeper types.BankKeeper

	storeKey  storetypes.StoreKey
	tStoreKey storetypes.StoreKey
	cdc       codec.BinaryCodec
	hooks     types.OcrHooks

	authority string
}

// NewKeeper creates a ocr Keeper
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	tStoreKey storetypes.StoreKey,
	bankKeeper types.BankKeeper,
	authority string,
) Keeper {
	return Keeper{
		cdc:        cdc,
		bankKeeper: bankKeeper,
		storeKey:   storeKey,
		tStoreKey:  tStoreKey,
		authority:  authority,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k *Keeper) getStore(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.storeKey)
}

func (k *Keeper) getTransientStore(ctx sdk.Context) sdk.KVStore {
	return ctx.TransientStore(k.tStoreKey)
}

func (k *Keeper) GetTransientStoreKey() storetypes.StoreKey {
	return k.tStoreKey
}
