package types

const (
	// ModuleName defines the module name
	ModuleName = "mint"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName
)

const (
	FirstDistTime = "first-dist-time-"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
