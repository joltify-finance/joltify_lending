package keeper

import (
	"fmt"
	"time"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/auction/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/tendermint/tendermint/libs/log"
)

type Keeper struct {
	storeKey      storetypes.StoreKey
	cdc           codec.Codec
	paramSubspace paramtypes.Subspace
	bankKeeper    types2.BankKeeper
	accountKeeper types2.AccountKeeper
}

// NewKeeper returns a new auction keeper.
func NewKeeper(cdc codec.Codec, storeKey storetypes.StoreKey, paramstore paramtypes.Subspace,
	bankKeeper types2.BankKeeper, accountKeeper types2.AccountKeeper,
) Keeper {
	if !paramstore.HasKeyTable() {
		paramstore = paramstore.WithKeyTable(types2.ParamKeyTable())
	}

	return Keeper{
		storeKey:      storeKey,
		cdc:           cdc,
		paramSubspace: paramstore,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
	}
}

// MustUnmarshalAuction attempts to decode and return an Auction object from
// raw encoded bytes. It panics on error.
func (k Keeper) MustUnmarshalAuction(bz []byte) types2.Auction {
	auction, err := k.UnmarshalAuction(bz)
	if err != nil {
		panic(fmt.Errorf("failed to decode auction: %w", err))
	}

	return auction
}

// MustMarshalAuction attempts to encode an Auction object and returns the
// raw encoded bytes. It panics on error.
func (k Keeper) MustMarshalAuction(auction types2.Auction) []byte {
	bz, err := k.MarshalAuction(auction)
	if err != nil {
		panic(fmt.Errorf("failed to encode auction: %w", err))
	}

	return bz
}

// MarshalAuction protobuf serializes an Auction interface
func (k Keeper) MarshalAuction(auctionI types2.Auction) ([]byte, error) {
	return k.cdc.MarshalInterface(auctionI)
}

// UnmarshalAuction returns an Auction interface from raw encoded auction
// bytes of a Proto-based Auction type
func (k Keeper) UnmarshalAuction(bz []byte) (types2.Auction, error) {
	var evi types2.Auction
	return evi, k.cdc.UnmarshalInterface(bz, &evi)
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types2.ModuleName))
}

// SetNextAuctionID stores an ID to be used for the next created auction
func (k Keeper) SetNextAuctionID(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types2.NextAuctionIDKey)
	store.Set(types2.NextAuctionIDKey, types2.Uint64ToBytes(id))
}

// GetNextAuctionID reads the next available global ID from store
func (k Keeper) GetNextAuctionID(ctx sdk.Context) (uint64, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types2.NextAuctionIDKey)
	bz := store.Get(types2.NextAuctionIDKey)
	if bz == nil {
		return 0, types2.ErrInvalidInitialAuctionID
	}
	return types2.Uint64FromBytes(bz), nil
}

// IncrementNextAuctionID increments the next auction ID in the store by 1.
func (k Keeper) IncrementNextAuctionID(ctx sdk.Context) error {
	id, err := k.GetNextAuctionID(ctx)
	if err != nil {
		return err
	}
	k.SetNextAuctionID(ctx, id+1)
	return nil
}

// StoreNewAuction stores an auction, adding a new ID
func (k Keeper) StoreNewAuction(ctx sdk.Context, auction types2.Auction) (uint64, error) {
	newAuctionID, err := k.GetNextAuctionID(ctx)
	if err != nil {
		return 0, err
	}

	auction = auction.WithID(newAuctionID)
	k.SetAuction(ctx, auction)

	err = k.IncrementNextAuctionID(ctx)
	if err != nil {
		return 0, err
	}
	return newAuctionID, nil
}

// SetAuction puts the auction into the store, and updates any indexes.
func (k Keeper) SetAuction(ctx sdk.Context, auction types2.Auction) {
	// remove the auction from the byTime index if it is already in there
	existingAuction, found := k.GetAuction(ctx, auction.GetID())
	if found {
		k.removeFromByTimeIndex(ctx, existingAuction.GetEndTime(), existingAuction.GetID())
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types2.AuctionKeyPrefix)

	store.Set(types2.GetAuctionKey(auction.GetID()), k.MustMarshalAuction(auction))
	k.InsertIntoByTimeIndex(ctx, auction.GetEndTime(), auction.GetID())
}

// GetAuction gets an auction from the store.
func (k Keeper) GetAuction(ctx sdk.Context, auctionID uint64) (types2.Auction, bool) {
	var auction types2.Auction

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types2.AuctionKeyPrefix)
	bz := store.Get(types2.GetAuctionKey(auctionID))
	if bz == nil {
		return auction, false
	}

	return k.MustUnmarshalAuction(bz), true
}

// DeleteAuction removes an auction from the store, and any indexes.
func (k Keeper) DeleteAuction(ctx sdk.Context, auctionID uint64) {
	auction, found := k.GetAuction(ctx, auctionID)
	if found {
		k.removeFromByTimeIndex(ctx, auction.GetEndTime(), auctionID)
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types2.AuctionKeyPrefix)
	store.Delete(types2.GetAuctionKey(auctionID))
}

// InsertIntoByTimeIndex adds an auction ID and end time into the byTime index.
func (k Keeper) InsertIntoByTimeIndex(ctx sdk.Context, endTime time.Time, auctionID uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types2.AuctionByTimeKeyPrefix)
	store.Set(types2.GetAuctionByTimeKey(endTime, auctionID), types2.Uint64ToBytes(auctionID))
}

// removeFromByTimeIndex removes an auction ID and end time from the byTime index.
func (k Keeper) removeFromByTimeIndex(ctx sdk.Context, endTime time.Time, auctionID uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types2.AuctionByTimeKeyPrefix)
	store.Delete(types2.GetAuctionByTimeKey(endTime, auctionID))
}

// IterateAuctionsByTime provides an iterator over auctions ordered by auction.EndTime.
// For each auction cb will be called. If cb returns true the iterator will close and stop.
func (k Keeper) IterateAuctionsByTime(ctx sdk.Context, inclusiveCutoffTime time.Time, cb func(auctionID uint64) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types2.AuctionByTimeKeyPrefix)
	iterator := store.Iterator(
		nil, // start at the very start of the prefix store
		sdk.PrefixEndBytes(sdk.FormatTimeBytes(inclusiveCutoffTime)), // include any keys with times equal to inclusiveCutoffTime
	)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {

		auctionID := types2.Uint64FromBytes(iterator.Value())

		if cb(auctionID) {
			break
		}
	}
}

// IterateAuctions provides an iterator over all stored auctions.
// For each auction, cb will be called. If cb returns true, the iterator will close and stop.
func (k Keeper) IterateAuctions(ctx sdk.Context, cb func(auction types2.Auction) (stop bool)) {
	iterator := sdk.KVStorePrefixIterator(ctx.KVStore(k.storeKey), types2.AuctionKeyPrefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		auction := k.MustUnmarshalAuction(iterator.Value())

		if cb(auction) {
			break
		}
	}
}

// GetAllAuctions returns all auctions from the store
func (k Keeper) GetAllAuctions(ctx sdk.Context) (auctions []types2.Auction) {
	k.IterateAuctions(ctx, func(auction types2.Auction) bool {
		auctions = append(auctions, auction)
		return false
	})
	return
}
