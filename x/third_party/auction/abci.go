package auction

import (
	"context"
	"errors"

	"github.com/joltify-finance/joltify_lending/x/third_party/auction/keeper"
	"github.com/joltify-finance/joltify_lending/x/third_party/auction/types"
)

// BeginBlocker closes all expired auctions at the end of each block. It panics if
// there's an error other than ErrAuctionNotFound.
func BeginBlocker(ctx context.Context, k keeper.Keeper) { //nolint:typecheck
	err := k.CloseExpiredAuctions(ctx)
	if err != nil && !errors.Is(err, types.ErrAuctionNotFound) {
		panic(err)
	}
}
