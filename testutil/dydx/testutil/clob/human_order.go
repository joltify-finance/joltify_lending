package clob

import clobtypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/types"

// TestHumanOrder is a test order with human readable price and size.
type TestHumanOrder struct {
	Order      clobtypes.Order
	HumanPrice string
	HumanSize  string
}
