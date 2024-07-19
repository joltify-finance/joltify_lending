package v5

import (
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// these codes are only for migration and may out of date

// Key Prefixes
var (
	USDXMintingClaimKeyPrefix                     = []byte{0x01} // prefix for keys that store USDX minting claims
	USDXMintingRewardFactorKeyPrefix              = []byte{0x02} // prefix for key that stores USDX minting reward factors
	PreviousUSDXMintingRewardAccrualTimeKeyPrefix = []byte{0x03} // prefix for key that stores the blocktime
	DelegatorClaimKeyPrefix                       = []byte{0x09} // prefix for keys that store delegator claims
	DelegatorRewardIndexesKeyPrefix               = []byte{0x10} // prefix for key that stores delegator reward indexes
	PreviousDelegatorRewardAccrualTimeKeyPrefix   = []byte{0x11} // prefix for key that stores the previous time delegator rewards accrued
	SwapClaimKeyPrefix                            = []byte{0x12} // prefix for keys that store swap claims
	SwapRewardIndexesKeyPrefix                    = []byte{0x13} // prefix for key that stores swap reward indexes
	PreviousSwapRewardAccrualTimeKeyPrefix        = []byte{0x14} // prefix for key that stores the previous time swap rewards accrued
	SavingsClaimKeyPrefix                         = []byte{0x15} // prefix for keys that store savings claims
	SavingsRewardIndexesKeyPrefix                 = []byte{0x16} // prefix for key that stores savings reward indexes
	PreviousSavingsRewardAccrualTimeKeyPrefix     = []byte{0x17} // prefix for key that stores the previous time savings rewards accrued
)

func MigrateStore(ctx context.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	toBeDeleted := [][]byte{USDXMintingClaimKeyPrefix, USDXMintingRewardFactorKeyPrefix, PreviousUSDXMintingRewardAccrualTimeKeyPrefix, PreviousSavingsRewardAccrualTimeKeyPrefix, SavingsClaimKeyPrefix, SavingsRewardIndexesKeyPrefix, SwapClaimKeyPrefix, SwapRewardIndexesKeyPrefix, PreviousSwapRewardAccrualTimeKeyPrefix, DelegatorClaimKeyPrefix, DelegatorRewardIndexesKeyPrefix, PreviousDelegatorRewardAccrualTimeKeyPrefix}
	for _, el := range toBeDeleted {
		store := prefix.NewStore(ctx.KVStore(storeKey), el)
		iterator := sdk.KVStorePrefixIterator(store, []byte{})
		if iterator.Valid() {
			panic("should be empty")
		}
		err := iterator.Close()
		if err != nil {
			panic("fail to close the iterator")
		}
	}

	return nil
}
