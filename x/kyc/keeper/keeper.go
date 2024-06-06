package keeper

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/joltify-finance/joltify_lending/x/kyc/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
		authority  sdk.AccAddress
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	authority sdk.AccAddress,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
		authority:  authority,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) SetProject(ctx sdk.Context, p *types.ProjectInfo) (int32, error) {
	var currentNum uint32
	projectStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectInfoPrefix))
	projectNum := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectInfoNum))
	data := projectNum.Get(types.KeyPrefix(types.ProjectInfoNum))
	if data == nil {
		currentNum = 1
	} else {
		currentNum = binary.BigEndian.Uint32(data)
		currentNum += 1
	}
	p.Index = int32(currentNum)
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, currentNum)
	projectNum.Set(types.KeyPrefix(types.ProjectInfoNum), bs)
	previousProject := projectStore.Get(types.KeyPrefix(string(p.Index)))
	if previousProject != nil {
		return 0, errors.New("project already exists")
	}
	projectStore.Set(types.KeyPrefix(string(p.Index)), k.cdc.MustMarshal(p))
	return int32(currentNum), nil
}

func (k Keeper) UpdateProject(ctx sdk.Context, p *types.ProjectInfo) {
	projectStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectInfoPrefix))
	projectStore.Set(types.KeyPrefix(string(p.Index)), k.cdc.MustMarshal(p))
}

func (k Keeper) GetProject(ctx sdk.Context, index int32) (val types.ProjectInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectInfoPrefix))
	bz := store.Get(types.KeyPrefix(string(index)))
	if bz == nil {
		return
	}
	k.cdc.MustUnmarshal(bz, &val)
	return val, true
}

// IteratePool iterates over all deposit objects in the store and performs a callback function
func (k Keeper) IterateProject(ctx sdk.Context, cb func(poolInfo types.ProjectInfo) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectInfoPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var projects types.ProjectInfo
		k.cdc.MustUnmarshal(iterator.Value(), &projects)
		if cb(projects) {
			break
		}
	}
}

func (k Keeper) DeleteProject(ctx sdk.Context, index int32) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectInfoPrefix))
	store.Delete(types.KeyPrefix(string(index)))
}
