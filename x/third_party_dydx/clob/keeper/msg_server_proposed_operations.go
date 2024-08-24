package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/lib"
	"github.com/joltify-finance/joltify_lending/lib/log"
	"github.com/joltify-finance/joltify_lending/lib/metrics"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/types"
)

func (k msgServer) ProposedOperations(
	goCtx context.Context,
	msg *types.MsgProposedOperations,
) (resp *types.MsgProposedOperationsResponse, err error) {
	ctx := lib.UnwrapSDKContext(goCtx, types.ModuleName)

	// Attach various logging tags relative to this request. These should be static with no changes.
	ctx = log.AddPersistentTagsToLogger(ctx,
		log.Module, log.Clob,
		log.ProposerConsAddress, sdk.ConsAddress(ctx.BlockHeader().ProposerAddress),
		log.Callback, lib.TxMode(ctx),
		log.BlockHeight, ctx.BlockHeight(),
		log.Handler, log.ProposedOperations,
		// Consider not appending this because it's massive
		// metrics.Msg, msg,
	)

	defer func() {
		metrics.IncrSuccessOrErrorCounter(err, types.ModuleName, metrics.ProposedOperations, metrics.DeliverTx)
		if err != nil {
			log.ErrorLogWithError(ctx, "Error in Proposed Operations", err)
		}
	}()

	fmt.Printf("2222222222222222222222222222222222222222222222222222222222222222222222222222222222222222\n")

	if err := k.Keeper.ProcessProposerOperations(
		ctx,
		msg.GetOperationsQueue(),
	); err != nil {
		return nil, err
	}

	return &types.MsgProposedOperationsResponse{}, nil
}
