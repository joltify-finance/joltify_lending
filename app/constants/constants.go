package constants

const (
	AppName       = "joltify"
	AppDaemonName = AppName
	ServiceName   = "validator"
)

// Slinky Constants

const (
	// OracleInfoIndex is the index at which slinky will inject VE data
	OracleInfoIndex = 0
	// OracleVEInjectedTxs is the number of transactions Slinky injects into the block (for VE data)
	OracleVEInjectedTxs = 1
)
