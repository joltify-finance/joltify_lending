package constants

import (
	"math/big"

	satypes "github.com/joltify-finance/joltify_lending/dydx_helper/x/subaccounts/types"
	clobtypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/types"
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
