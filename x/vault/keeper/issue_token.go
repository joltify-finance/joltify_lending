package keeper

import (
	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
)

// SetIssueToken set a specific issueToken in the store from its index
func (k Keeper) SetIssueToken(ctx context.Context, issueToken types.IssueToken) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssueTokenKey))
	b := k.cdc.MustMarshal(&issueToken)
	store.Set(types.KeyPrefix(issueToken.Index), b)
}

// GetIssueToken returns a issueToken from its index
func (k Keeper) GetIssueToken(ctx context.Context, index string) (val types.IssueToken, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssueTokenKey))

	b := store.Get(types.KeyPrefix(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllIssueToken returns all issueToken
func (k Keeper) GetAllIssueToken(ctx context.Context) (list []types.IssueToken) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IssueTokenKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.IssueToken
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
