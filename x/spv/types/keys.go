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
	ProjectsKeyPrefix    = "ProjectsKeyPrefix-"
	PoolInvestor         = "poolInvestorsPrefix-"
	PoolDepositor        = "depositorPrefix-"
	PoolDepositorHistory = "depositorhistoryPrefix-"
	Pool                 = "pool-"
	HistoryPool          = "historyPool-"
	ArchivePrefix        = "archive-"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
