package types

const (
	// ModuleName The name that will be used throughout the module
	ModuleName = "incentive"

	// StoreKey Top level store key where all module items will be stored
	StoreKey = ModuleName

	// RouterKey Top level router key
	RouterKey = ModuleName

	// QuerierRoute route used for abci queries
	QuerierRoute = ModuleName
)

// Key Prefixes
var (
	USDXMintingClaimKeyPrefix                     = []byte{0x01} // prefix for keys that store USDX minting claims
	USDXMintingRewardFactorKeyPrefix              = []byte{0x02} // prefix for key that stores USDX minting reward factors
	PreviousUSDXMintingRewardAccrualTimeKeyPrefix = []byte{0x03} // prefix for key that stores the blocktime
	JoltLiquidityClaimKeyPrefix                   = []byte{0x04} // prefix for keys that store Jolt liquidity claims
	JoltSupplyRewardIndexesKeyPrefix              = []byte{0x05} // prefix for key that stores Jolt supply reward indexes
	PreviousJoltSupplyRewardAccrualTimeKeyPrefix  = []byte{0x06} // prefix for key that stores the previous time Jolt supply rewards accrued
	JoltBorrowRewardIndexesKeyPrefix              = []byte{0x07} // prefix for key that stores Jolt borrow reward indexes
	PreviousJoltBorrowRewardAccrualTimeKeyPrefix  = []byte{0x08} // prefix for key that stores the previous time Jolt borrow rewards accrued
	DelegatorClaimKeyPrefix                       = []byte{0x09} // prefix for keys that store delegator claims
	DelegatorRewardIndexesKeyPrefix               = []byte{0x10} // prefix for key that stores delegator reward indexes
	PreviousDelegatorRewardAccrualTimeKeyPrefix   = []byte{0x11} // prefix for key that stores the previous time delegator rewards accrued
	SwapClaimKeyPrefix                            = []byte{0x12} // prefix for keys that store swap claims
	SwapRewardIndexesKeyPrefix                    = []byte{0x13} // prefix for key that stores swap reward indexes
	PreviousSwapRewardAccrualTimeKeyPrefix        = []byte{0x14} // prefix for key that stores the previous time swap rewards accrued
	SavingsClaimKeyPrefix                         = []byte{0x15} // prefix for keys that store savings claims
	SavingsRewardIndexesKeyPrefix                 = []byte{0x16} // prefix for key that stores savings reward indexes
	PreviousSavingsRewardAccrualTimeKeyPrefix     = []byte{0x17} // prefix for key that stores the previous time savings rewards accrued
	SPVClaimKeyPrefix                             = []byte{0x18} // prefix for keys that store spv claims
	SPVRewardIndexesKeyPrefix                     = []byte{0x19} // prefix for key that stores svp reward indexes
	PreviousSPVRewardAccrualTimeKeyPrefix         = []byte{0x20} // prefix for key that stores the previous time spv rewards accrued
)
