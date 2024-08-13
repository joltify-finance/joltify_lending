package types

import bridgetypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/bridge/types"

// BridgeQueryClient is an interface that encapsulates the x/bridge `QueryClient` interface.
type BridgeQueryClient interface {
	bridgetypes.QueryClient
}
