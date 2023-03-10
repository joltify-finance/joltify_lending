package types

const (
	// ModuleName defines the module name
	ModuleName = "vault"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_vault"

	// this line is used by starport scaffolding # ibc/keys/name
)

// this line is used by starport scaffolding # ibc/keys/port

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	CreatePoolKey = "CreatePool-value-"

	LastTwoPoolKey = "LastTwoPool-"

	IssueTokenKey = "IssueToken-value-"

	ValidatorsStoreKey = "ValidatorStore-value-"

	StandbyPwoerStoreKey = "standbyPower-value-"

	FeeStoreKey = "fee_collected"

	QuotaStoreKey = "quota-vaule-"
)
