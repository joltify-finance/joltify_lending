package auction_test

import (
	"sort"
	"testing"
	"time"

	"github.com/joltify-finance/joltify_lending/x/third_party/auction"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/auction/types"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/joltify-finance/joltify_lending/app"
)

var (
	_, testAddrs = app.GeneratePrivKeyAddressPairs(2)
	testTime     = time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC)
	testAuction  = types2.NewCollateralAuction(
		"seller",
		c("lotdenom", 10),
		testTime,
		c("biddenom", 1000),
		types2.WeightedAddresses{Addresses: testAddrs, Weights: []sdk.Int{sdk.OneInt(), sdk.OneInt()}},
		c("debt", 1000),
	).WithID(3).(types2.GenesisAuction)
)

func TestInitGenesis(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		// setup keepers
		tApp := app.NewTestApp()
		ctx := tApp.NewContext(true, tmproto.Header{Height: 1})

		// setup module account
		modBaseAcc := authtypes.NewBaseAccount(authtypes.NewModuleAddress(types2.ModuleName), nil, 0, 0)
		modAcc := authtypes.NewModuleAccount(modBaseAcc, types2.ModuleName, []string{authtypes.Minter, authtypes.Burner}...)
		tApp.GetAccountKeeper().SetModuleAccount(ctx, modAcc)
		tApp.GetBankKeeper().MintCoins(ctx, types2.ModuleName, testAuction.GetModuleAccountCoins())

		// set up auction genesis state with module account
		auctionGS, err := types2.NewGenesisState(
			10,
			types2.DefaultParams(),
			[]types2.GenesisAuction{testAuction},
		)
		require.NoError(t, err)

		// run init
		keeper := tApp.GetAuctionKeeper()
		require.NotPanics(t, func() {
			auction.InitGenesis(ctx, keeper, tApp.GetBankKeeper(), tApp.GetAccountKeeper(), auctionGS)
		})

		// check state is as expected
		actualID, err := keeper.GetNextAuctionID(ctx)
		require.NoError(t, err)
		require.Equal(t, auctionGS.NextAuctionId, actualID)

		require.Equal(t, auctionGS.Params, keeper.GetParams(ctx))

		genesisAuctions, err := types2.UnpackGenesisAuctions(auctionGS.Auctions)
		if err != nil {
			panic(err)
		}

		sort.Slice(genesisAuctions, func(i, j int) bool {
			return genesisAuctions[i].GetID() > genesisAuctions[j].GetID()
		})
		i := 0
		keeper.IterateAuctions(ctx, func(a types2.Auction) bool {
			require.Equal(t, genesisAuctions[i], a)
			i++
			return false
		})
	})
	t.Run("invalid (invalid nextAuctionID)", func(t *testing.T) {
		// setup keepers
		tApp := app.NewTestApp()
		ctx := tApp.NewContext(true, tmproto.Header{Height: 1})

		// setup module account
		modBaseAcc := authtypes.NewBaseAccount(authtypes.NewModuleAddress(types2.ModuleName), nil, 0, 0)
		modAcc := authtypes.NewModuleAccount(modBaseAcc, types2.ModuleName, []string{authtypes.Minter, authtypes.Burner}...)
		tApp.GetAccountKeeper().SetModuleAccount(ctx, modAcc)
		tApp.GetBankKeeper().MintCoins(ctx, types2.ModuleName, testAuction.GetModuleAccountCoins())

		// create invalid genesis
		auctionGS, err := types2.NewGenesisState(
			0, // next id < testAuction ID
			types2.DefaultParams(),
			[]types2.GenesisAuction{testAuction},
		)
		require.NoError(t, err)

		// check init fails
		require.Panics(t, func() {
			auction.InitGenesis(ctx, tApp.GetAuctionKeeper(), tApp.GetBankKeeper(), tApp.GetAccountKeeper(), auctionGS)
		})
	})
	t.Run("invalid (missing mod account coins)", func(t *testing.T) {
		// setup keepers
		tApp := app.NewTestApp()
		ctx := tApp.NewContext(true, tmproto.Header{Height: 1})

		// invalid as there is no module account setup

		// create invalid genesis
		auctionGS, err := types2.NewGenesisState(
			10,
			types2.DefaultParams(),
			[]types2.GenesisAuction{testAuction},
		)
		require.NoError(t, err)

		// check init fails
		require.Panics(t, func() {
			auction.InitGenesis(ctx, tApp.GetAuctionKeeper(), tApp.GetBankKeeper(), tApp.GetAccountKeeper(), auctionGS)
		})
	})
}

func TestExportGenesis(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		// setup state
		tApp := app.NewTestApp()
		ctx := tApp.NewContext(true, tmproto.Header{Height: 1})
		tApp.InitializeFromGenesisStates()

		// export
		gs := auction.ExportGenesis(ctx, tApp.GetAuctionKeeper())

		// check state matches
		defaultGS := types2.DefaultGenesisState()
		require.Equal(t, defaultGS, gs)
	})
	t.Run("one auction", func(t *testing.T) {
		// setup state
		tApp := app.NewTestApp()
		ctx := tApp.NewContext(true, tmproto.Header{Height: 1})
		tApp.InitializeFromGenesisStates()
		tApp.GetAuctionKeeper().SetAuction(ctx, testAuction)

		// export
		gs := auction.ExportGenesis(ctx, tApp.GetAuctionKeeper())

		// check state matches
		expectedGenesisState := types2.DefaultGenesisState()
		packedGenesisAuctions, err := types2.PackGenesisAuctions([]types2.GenesisAuction{testAuction})
		require.NoError(t, err)

		expectedGenesisState.Auctions = append(expectedGenesisState.Auctions, packedGenesisAuctions...)
		require.Equal(t, expectedGenesisState, gs)
	})
}
