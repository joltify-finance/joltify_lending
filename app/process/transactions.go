package process

import (
	"fmt"
	"slices"

	errorsmod "cosmossdk.io/errors"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/lib"
)

const (
	minTxsCount               = 3
	proposedOperationsTxIndex = 0
	// updateMarketPricesTxLenOffset = -1
	addPremiumVotesTxLenOffset    = -1
	acknowledgeBridgesTxLenOffset = -2
	lastOtherTxLenOffset          = acknowledgeBridgesTxLenOffset
	firstOtherTxIndex             = proposedOperationsTxIndex + 1
)

func init() {
	txIndicesAndOffsets := []int{
		proposedOperationsTxIndex,
		acknowledgeBridgesTxLenOffset,
		addPremiumVotesTxLenOffset,
		// updateMarketPricesTxLenOffset,
	}
	if minTxsCount != len(txIndicesAndOffsets) {
		panic("minTxsCount does not match expected count of Txs.")
	}
	if lib.ContainsDuplicates(txIndicesAndOffsets) {
		panic("Duplicate indices/offsets defined for Txs.")
	}
	if slices.Min[[]int](txIndicesAndOffsets) != lastOtherTxLenOffset {
		panic("lastTxLenOffset is not the lowest offset")
	}
	if slices.Max[[]int](txIndicesAndOffsets)+1 != firstOtherTxIndex {
		panic("firstOtherTxIndex is <= the maximum offset")
	}
	txIndicesForMinTxsCount := []int{
		proposedOperationsTxIndex,
		acknowledgeBridgesTxLenOffset + minTxsCount,
		addPremiumVotesTxLenOffset + minTxsCount,
		// updateMarketPricesTxLenOffset + minTxsCount,
	}
	if minTxsCount != len(txIndicesForMinTxsCount) {
		panic("minTxsCount does not match expected count of Txs.")
	}
	if lib.ContainsDuplicates(txIndicesForMinTxsCount) {
		panic("Overlapping indices and offsets defined for Txs.")
	}
	if minTxsCount != firstOtherTxIndex-lastOtherTxLenOffset {
		panic("Unexpected gap between firstOtherTxIndex and lastOtherTxLenOffset which is greater than minTxsCount")
	}
}

// ProcessProposalTxs is used as an intermediary struct to validate a proposed list of txs
// for `ProcessProposal`.
type ProcessProposalTxs struct {
	// Single msg txs.
	ProposedOperationsTx *ProposedOperationsTx
	AcknowledgeBridgesTx *AcknowledgeBridgesTx
	AddPremiumVotesTx    *AddPremiumVotesTx
	// george do not need price
	// UpdateMarketPricesTx *UpdateMarketPricesTx // abstract over MarketPriceUpdates from VEs or default.

	// Multi msgs txs.
	OtherTxs []*OtherMsgsTx
}

// DecodeProcessProposalTxs returns a new `processProposalTxs`.
func DecodeProcessProposalTxs(
	ctx sdk.Context,
	decoder sdk.TxDecoder,
	req *abci.RequestProcessProposal,
	bridgeKeeper ProcessBridgeKeeper,
	pricesTxDecoder UpdateMarketPriceTxDecoder,
) (*ProcessProposalTxs, error) {
	// Check len (accounting for offset from injected vote-extensions if applicable)
	offset := pricesTxDecoder.GetTxOffset(ctx)
	injectedTxCount := minTxsCount + offset
	numTxs := len(req.Txs)
	if numTxs < injectedTxCount {
		return nil, errorsmod.Wrapf(
			ErrUnexpectedNumMsgs,
			"Expected the proposal to contain at least %d txs, but got %d",
			injectedTxCount,
			numTxs,
		)
	}

	// Price updates.
	//updatePricesTx, err := pricesTxDecoder.DecodeUpdateMarketPricesTx(
	//	ctx,
	//	req.Txs,
	//)
	//if err != nil {
	//	return nil, err
	//}

	// Operations.
	// if vote-extensions were injected, offset will be incremented.
	operationsTx, err := DecodeProposedOperationsTx(decoder, req.Txs[proposedOperationsTxIndex+offset])
	if err != nil {
		return nil, err
	}

	// Acknowledge bridges.
	acknowledgeBridgesTx, err := DecodeAcknowledgeBridgesTx(
		ctx,
		bridgeKeeper,
		decoder,
		req.Txs[numTxs+acknowledgeBridgesTxLenOffset],
	)

	if acknowledgeBridgesTx.msg != nil {
	} else {
		fmt.Printf("NOT BBBBBEMPTY!!!!%v\n", acknowledgeBridgesTx.msg.String())
	}

	if err != nil {
		return nil, err
	}

	// Funding samples.
	addPremiumVotesTx, err := DecodeAddPremiumVotesTx(decoder, req.Txs[numTxs+addPremiumVotesTxLenOffset])
	if err != nil {
		return nil, err
	}

	// Other txs.
	// if vote-extensions were injected, offset will be incremented.
	allOtherTxs := make([]*OtherMsgsTx, numTxs-injectedTxCount)
	for i, txBytes := range req.Txs[firstOtherTxIndex+offset : numTxs+lastOtherTxLenOffset] {
		otherTx, err := DecodeOtherMsgsTx(decoder, txBytes)
		if err != nil {
			return nil, err
		}

		allOtherTxs[i] = otherTx
	}

	return &ProcessProposalTxs{
		ProposedOperationsTx: operationsTx,
		AcknowledgeBridgesTx: acknowledgeBridgesTx,
		AddPremiumVotesTx:    addPremiumVotesTx,
		OtherTxs:             allOtherTxs,
	}, nil
}

// Validate performs `ValidateBasic` on the underlying msgs that are part of the txs.
// Returns nil if all are valid. Otherwise, returns error.
//
// Exception: for UpdateMarketPricesTx, we perform "in-memory stateful" validation
// to ensure that the new proposed prices are "valid" in comparison to index prices.
func (ppt *ProcessProposalTxs) Validate() error {
	// Validate single msg txs.
	singleTxs := []SingleMsgTx{
		ppt.ProposedOperationsTx,
		ppt.AddPremiumVotesTx,
		ppt.AcknowledgeBridgesTx,
	}
	for _, smt := range singleTxs {
		if err := smt.Validate(); err != nil {
			return err
		}
	}

	// Validate multi msgs txs.
	for _, mmt := range ppt.OtherTxs {
		if err := mmt.Validate(); err != nil {
			return err
		}
	}

	return nil
}
