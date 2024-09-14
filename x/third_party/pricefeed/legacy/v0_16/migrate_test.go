package v0_16

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	appconfig "github.com/joltify-finance/joltify_lending/app/config"
	v015pricefeed "github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/legacy/v0_15"
	"github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/joltify-finance/joltify_lending/app"
)

type migrateTestSuite struct {
	suite.Suite

	addresses   []sdk.AccAddress
	v15genstate v015pricefeed.GenesisState
	cdc         codec.Codec
	legacyCdc   *codec.LegacyAmino
}

func (s *migrateTestSuite) SetupTest() {
	appconfig.SetupConfig()

	s.v15genstate = v015pricefeed.GenesisState{
		Params:       v015pricefeed.Params{},
		PostedPrices: v015pricefeed.PostedPrices{},
	}

	config := appconfig.MakeEncodingConfig()
	s.cdc = config.Codec

	legacyCodec := codec.NewLegacyAmino()
	s.legacyCdc = legacyCodec

	_, accAddresses := app.GeneratePrivKeyAddressPairs(10)
	s.addresses = accAddresses
}

func (s *migrateTestSuite) TestMigrate_JSON() {
	// Migrate v15 pricefeed to v16
	v15Params := `{
		"params": {
			"markets": [
				{
					"active": true,
					"base_asset": "bnb",
					"market_id": "bnb:usd",
					"oracles": ["jolt10wlnqzyss4accfqmyxwx5jy5x9nfkwh6kwhjs9"],
					"quote_asset": "usd"
				},
				{
					"active": true,
					"base_asset": "bnb",
					"market_id": "bnb:usd:30",
					"oracles": ["jolt10wlnqzyss4accfqmyxwx5jy5x9nfkwh6kwhjs9"],
					"quote_asset": "usd"
				}
			]
		},
		"posted_prices": [
			{
				"expiry": "2022-07-20T00:00:00Z",
				"market_id": "bnb:usd",
				"oracle_address": "jolt10wlnqzyss4accfqmyxwx5jy5x9nfkwh6kwhjs9",
				"price": "215.962650000000001782"
			},
			{
				"expiry": "2022-07-20T00:00:00Z",
				"market_id": "bnb:usd:30",
				"oracle_address": "jolt10wlnqzyss4accfqmyxwx5jy5x9nfkwh6kwhjs9",
				"price": "217.962650000000001782"
			}
		]
	}`

	expectedV16Params := `{
		"params": {
			"markets": [
				{
					"market_id": "bnb:usd",
					"base_asset": "bnb",
					"quote_asset": "usd",
					"oracles": [
						"jolt10wlnqzyss4accfqmyxwx5jy5x9nfkwh6kwhjs9"
					],
					"active": true
				},
				{
					"market_id": "bnb:usd:30",
					"base_asset": "bnb",
					"quote_asset": "usd",
					"oracles": [
						"jolt10wlnqzyss4accfqmyxwx5jy5x9nfkwh6kwhjs9"
					],
					"active": true
				},
				{
					"market_id": "atom:usd",
					"base_asset": "atom",
					"quote_asset": "usd",
					"oracles": [
						"jolt10wlnqzyss4accfqmyxwx5jy5x9nfkwh6kwhjs9"
					],
					"active": true
				},
				{
					"market_id": "atom:usd:30",
					"base_asset": "atom",
					"quote_asset": "usd",
					"oracles": [
						"jolt10wlnqzyss4accfqmyxwx5jy5x9nfkwh6kwhjs9"
					],
					"active": true
				},
				{
					"market_id": "akt:usd",
					"base_asset": "akt",
					"quote_asset": "usd",
					"oracles": [
						"jolt10wlnqzyss4accfqmyxwx5jy5x9nfkwh6kwhjs9"
					],
					"active": true
				},
				{
					"market_id": "akt:usd:30",
					"base_asset": "akt",
					"quote_asset": "usd",
					"oracles": [
						"jolt10wlnqzyss4accfqmyxwx5jy5x9nfkwh6kwhjs9"
					],
					"active": true
				},
				{
					"market_id": "luna:usd",
					"base_asset": "luna",
					"quote_asset": "usd",
					"oracles": [
						"jolt10wlnqzyss4accfqmyxwx5jy5x9nfkwh6kwhjs9"
					],
					"active": true
				},
				{
					"market_id": "luna:usd:30",
					"base_asset": "luna",
					"quote_asset": "usd",
					"oracles": [
						"jolt10wlnqzyss4accfqmyxwx5jy5x9nfkwh6kwhjs9"
					],
					"active": true
				},
				{
					"market_id": "osmo:usd",
					"base_asset": "osmo",
					"quote_asset": "usd",
					"oracles": [
						"jolt10wlnqzyss4accfqmyxwx5jy5x9nfkwh6kwhjs9"
					],
					"active": true
				},
				{
					"market_id": "osmo:usd:30",
					"base_asset": "osmo",
					"quote_asset": "usd",
					"oracles": [
						"jolt10wlnqzyss4accfqmyxwx5jy5x9nfkwh6kwhjs9"
					],
					"active": true
				},
				{
					"market_id": "ust:usd",
					"base_asset": "ust",
					"quote_asset": "usd",
					"oracles": [
						"jolt10wlnqzyss4accfqmyxwx5jy5x9nfkwh6kwhjs9"
					],
					"active": true
				},
				{
					"market_id": "ust:usd:30",
					"base_asset": "ust",
					"quote_asset": "usd",
					"oracles": [
						"jolt10wlnqzyss4accfqmyxwx5jy5x9nfkwh6kwhjs9"
					],
					"active": true
				}
			]
		},
		"posted_prices": [
			{
				"market_id": "bnb:usd",
				"oracle_address": "jolt10wlnqzyss4accfqmyxwx5jy5x9nfkwh6kwhjs9",
				"price": "215.962650000000001782",
				"expiry": "2022-07-20T00:00:00Z"
			},
			{
				"market_id": "bnb:usd:30",
				"oracle_address": "jolt10wlnqzyss4accfqmyxwx5jy5x9nfkwh6kwhjs9",
				"price": "217.962650000000001782",
				"expiry": "2022-07-20T00:00:00Z"
			}
		]
	}`

	err := s.legacyCdc.UnmarshalJSON([]byte(v15Params), &s.v15genstate)
	s.Require().NoError(err)
	genstate := Migrate(s.v15genstate)

	// v16 pricefeed json should be the same as v15 but with IBC markets added
	actual := s.cdc.MustMarshalJSON(genstate)

	s.Require().NoError(err)
	s.Require().JSONEq(expectedV16Params, string(actual))
}

func (s *migrateTestSuite) TestMigrate_Params() {
	s.v15genstate.Params = v015pricefeed.Params{
		Markets: v015pricefeed.Markets{
			{
				MarketID:   "market-1",
				BaseAsset:  "joltify",
				QuoteAsset: "usd",
				Oracles:    s.addresses,
				Active:     true,
			},
		},
	}
	expectedParams := types.Params{
		Markets: types.Markets{
			{
				MarketID:   "market-1",
				BaseAsset:  "joltify",
				QuoteAsset: "usd",
				Oracles:    s.addresses,
				Active:     true,
			},
			{
				MarketID:   "atom:usd",
				BaseAsset:  "atom",
				QuoteAsset: "usd",
				Oracles:    s.addresses,
				Active:     true,
			},
			{
				MarketID:   "atom:usd:30",
				BaseAsset:  "atom",
				QuoteAsset: "usd",
				Oracles:    s.addresses,
				Active:     true,
			},
			{
				MarketID:   "akt:usd",
				BaseAsset:  "akt",
				QuoteAsset: "usd",
				Oracles:    s.addresses,
				Active:     true,
			},
			{
				MarketID:   "akt:usd:30",
				BaseAsset:  "akt",
				QuoteAsset: "usd",
				Oracles:    s.addresses,
				Active:     true,
			},
			{
				MarketID:   "luna:usd",
				BaseAsset:  "luna",
				QuoteAsset: "usd",
				Oracles:    s.addresses,
				Active:     true,
			},
			{
				MarketID:   "luna:usd:30",
				BaseAsset:  "luna",
				QuoteAsset: "usd",
				Oracles:    s.addresses,
				Active:     true,
			},
			{
				MarketID:   "osmo:usd",
				BaseAsset:  "osmo",
				QuoteAsset: "usd",
				Oracles:    s.addresses,
				Active:     true,
			},
			{
				MarketID:   "osmo:usd:30",
				BaseAsset:  "osmo",
				QuoteAsset: "usd",
				Oracles:    s.addresses,
				Active:     true,
			},
			{
				MarketID:   "ust:usd",
				BaseAsset:  "ust",
				QuoteAsset: "usd",
				Oracles:    s.addresses,
				Active:     true,
			},
			{
				MarketID:   "ust:usd:30",
				BaseAsset:  "ust",
				QuoteAsset: "usd",
				Oracles:    s.addresses,
				Active:     true,
			},
		},
	}
	genState := Migrate(s.v15genstate)
	s.Require().Equal(expectedParams, genState.Params)
}

func (s *migrateTestSuite) TestMigrate_PostedPrices() {
	s.v15genstate.PostedPrices = v015pricefeed.PostedPrices{
		{
			MarketID:      "market-1",
			OracleAddress: s.addresses[0],
			Price:         sdkmath.LegacyMustNewDecFromStr("1.2"),
			Expiry:        time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			MarketID:      "market-2",
			OracleAddress: s.addresses[1],
			Price:         sdkmath.LegacyMustNewDecFromStr("1.899"),
			Expiry:        time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	expected := types.PostedPrices{
		{
			MarketID:      "market-1",
			OracleAddress: s.addresses[0],
			Price:         sdkmath.LegacyMustNewDecFromStr("1.2"),
			Expiry:        time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			MarketID:      "market-2",
			OracleAddress: s.addresses[1],
			Price:         sdkmath.LegacyMustNewDecFromStr("1.899"),
			Expiry:        time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	genState := Migrate(s.v15genstate)
	s.Require().Equal(expected, genState.PostedPrices)
}

func TestPriceFeedMigrateTestSuite(t *testing.T) {
	suite.Run(t, new(migrateTestSuite))
}
