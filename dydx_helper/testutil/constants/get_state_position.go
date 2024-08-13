package constants

import (
	"math/big"

	clobtypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/types"
	satypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/subaccounts/types"
)

// Get state position functions.
var GetStatePosition_ZeroPositionSize = func(
	subaccountId satypes.SubaccountId,
	clobPairId clobtypes.ClobPairId,
) (
	statePositionSize *big.Int,
) {
	return big.NewInt(0)
}
