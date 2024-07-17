package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/joltify-finance/joltify_lending/x/third_party/exchange/types"
)

type DerivativeOrderStateExpansion struct {
	SubaccountID  common.Hash
	PositionDelta *types.PositionDelta
	Payout        sdk.Dec

	TotalBalanceDelta     sdk.Dec
	AvailableBalanceDelta sdk.Dec

	AuctionFeeReward       sdk.Dec
	TradingRewardPoints    sdk.Dec
	FeeRecipientReward     sdk.Dec
	FeeRecipient           common.Address
	LimitOrderFilledDelta  *types.DerivativeLimitOrderDelta
	MarketOrderFilledDelta *types.DerivativeMarketOrderDelta
	OrderHash              common.Hash
	Cid                    string
}
