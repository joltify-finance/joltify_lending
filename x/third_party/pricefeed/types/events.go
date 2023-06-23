package types

// Pricefeed module event types
const (
	EventTypeMarketPriceUpdated = "market_price_updated"
	EventTypeOracleUpdatedPrice = "oracle_updated_price"

	AttributeValueCategory = ModuleName
	AttributeMarketID      = "market_id"
	AttributeMarketPrice   = "market_price"
	AttributeOracle        = "oracle"
	AttributeExpiry        = "expiry"
)
