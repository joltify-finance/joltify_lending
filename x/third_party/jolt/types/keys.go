package types

const (
	// ModuleName name that will be used throughout the module
	ModuleName = "jolt"

	// ModuleAccountName name of module account used to hold deposits
	ModuleAccountName = "jolt"

	// StoreKey Top level store key where all module items will be stored
	StoreKey = ModuleName

	// RouterKey Top level router key
	RouterKey = ModuleName

	// QuerierRoute Top level query string
	QuerierRoute = ModuleName
)

var (
	DepositsKeyPrefix             = []byte{0x01}
	BorrowsKeyPrefix              = []byte{0x02}
	BorrowedCoinsPrefix           = []byte{0x03}
	SuppliedCoinsPrefix           = []byte{0x04}
	MoneyMarketsPrefix            = []byte{0x05}
	PreviousAccrualTimePrefix     = []byte{0x06} // denom -> time
	TotalReservesPrefix           = []byte{0x07} // denom -> sdk.Coin
	BorrowInterestFactorPrefix    = []byte{0x08} // denom -> sdkmath.LegacyDec
	SupplyInterestFactorPrefix    = []byte{0x09} // denom -> sdkmath.LegacyDec
	DelegatorInterestFactorPrefix = []byte{0x10} // denom -> sdkmath.LegacyDec
)
