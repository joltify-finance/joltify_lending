package indexer_manager

import "github.com/joltify-finance/joltify_lending/dydx_helper/indexer/msgsender"

func NewIndexerEventManagerNoop() IndexerEventManager {
	return NewIndexerEventManager(
		msgsender.NewIndexerMessageSenderNoop(),
		nil,
		false,
	)
}

func NewIndexerEventManagerNoopEnabled() IndexerEventManager {
	return NewIndexerEventManager(
		msgsender.NewIndexerMessageSenderNoopEnabled(),
		nil,
		false,
	)
}
