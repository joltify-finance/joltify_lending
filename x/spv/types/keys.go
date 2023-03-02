package types

const (
	// ModuleName defines the module name
	ModuleName = "spv"

	ModuleAccount = "spv"
	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_spv"
)

const (
	ProjectsKeyPrefix        = "ProjectsKeyPrefix-"
	HistoryProjectsKeyPrefix = "HistoryProjectsKeyPrefix-"
	PoolInvestor             = "poolInvestorsPrefix-"
	PoolDepositor            = "depositorPrefix-"
	PoolDeposited            = "depositedPrefix-"
	Pool                     = "pool-"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
