package types

const (
	// ModuleName defines the module name
	ModuleName = "quota"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_quota"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	QuotaKey = "Quota/value/"
)
