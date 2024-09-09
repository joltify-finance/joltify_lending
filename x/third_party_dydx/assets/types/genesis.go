package types

import "github.com/joltify-finance/joltify_lending/lib"

const (
	// UusdcDenom is the precomputed denom for IBC Micro USDC.
	UusdcDenom         = "ibc/65D0BEC6DAD96C7F5043D1E54E54B6BB5D5B3AEC3FF6CEBB75B9E059F3580EA3"
	UusdcDenomExponent = -6
)

var AssetUsdc Asset = Asset{
	Id:               0,
	Symbol:           "USDC",
	DenomExponent:    UusdcDenomExponent,
	Denom:            UusdcDenom,
	HasMarket:        false,
	AtomicResolution: lib.QuoteCurrencyAtomicResolution,
}

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Assets: []Asset{
			AssetUsdc,
		},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Genesis state should contain at least one asset.
	if len(gs.Assets) == 0 {
		return ErrNoAssetInGenesis
	}

	// The first asset should always be USDC.
	if gs.Assets[0] != AssetUsdc {
		return ErrUsdcMustBeAssetZero
	}

	// Provided assets should not contain duplicated asset ids, and denoms.
	// Asset ids should be sequential.
	// MarketId should be 0 if HasMarket is false.
	assetIdSet := make(map[uint32]struct{})
	denomSet := make(map[string]struct{})
	expectedId := uint32(0)

	for _, asset := range gs.Assets {
		if _, exists := assetIdSet[asset.Id]; exists {
			return ErrAssetIdAlreadyExists
		}
		if _, exists := denomSet[asset.Denom]; exists {
			return ErrAssetDenomAlreadyExists
		}
		if asset.Id != expectedId {
			return ErrGapFoundInAssetId
		}
		if !asset.HasMarket && asset.MarketId > 0 {
			return ErrInvalidMarketId
		}
		assetIdSet[asset.Id] = struct{}{}
		denomSet[asset.Denom] = struct{}{}
		expectedId++
	}
	return nil
}
