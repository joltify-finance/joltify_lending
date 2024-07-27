package keeper_test

import (
	"testing"
	"time"

	tmlog "github.com/cometbft/cometbft/libs/log"
	"github.com/joltify-finance/joltify_lending/x/third_party/auction/keeper"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/auction/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/stretchr/testify/require"
)

func TestGrpcAuctionsFilter(t *testing.T) {
	// setup
	lg := tmlog.TestingLogger()
	tApp := app.NewTestApp(lg, t.TempDir())
	tApp.InitializeFromGenesisStates(nil, nil)
	auctionsKeeper := tApp.GetAuctionKeeper()
	ctx := tApp.NewContext(true, tmproto.Header{Height: 1})
	_, addrs := app.GeneratePrivKeyAddressPairs(2)

	auctions := []types2.Auction{
		types2.NewSurplusAuction(
			"sellerMod",
			c("swp", 12345678),
			"usdx",
			time.Date(1998, time.January, 1, 0, 0, 0, 0, time.UTC),
		).WithID(0),
		types2.NewDebtAuction(
			"buyerMod",
			c("jolt", 12345678),
			c("usdx", 12345678),
			time.Date(1998, time.January, 1, 0, 0, 0, 0, time.UTC),
			c("debt", 12345678),
		).WithID(1),
		types2.NewCollateralAuction(
			"sellerMod",
			c("ujolt", 12345678),
			time.Date(1998, time.January, 1, 0, 0, 0, 0, time.UTC),
			c("usdx", 12345678),
			types2.WeightedAddresses{
				Addresses: addrs,
				Weights:   []sdkmath.Int{sdkmath.NewInt(100)},
			},
			c("debt", 12345678),
		).WithID(2),
		types2.NewCollateralAuction(
			"sellerMod",
			c("jolt", 12345678),
			time.Date(1998, time.January, 1, 0, 0, 0, 0, time.UTC),
			c("usdx", 12345678),
			types2.WeightedAddresses{
				Addresses: addrs,
				Weights:   []sdkmath.Int{sdkmath.NewInt(100)},
			},
			c("debt", 12345678),
		).WithID(3),
	}
	for _, a := range auctions {
		auctionsKeeper.SetAuction(ctx, a)
	}

	qs := keeper.NewQueryServerImpl(auctionsKeeper)

	tests := []struct {
		giveName     string
		giveRequest  types2.QueryAuctionsRequest
		wantResponse []types2.Auction
	}{
		{
			"empty request",
			types2.QueryAuctionsRequest{},
			auctions,
		},
		{
			"denom query swp",
			types2.QueryAuctionsRequest{
				Denom: "swp",
			},
			auctions[0:1],
		},
		{
			"denom query usdx all",
			types2.QueryAuctionsRequest{
				Denom: "usdx",
			},
			auctions,
		},
		{
			"owner",
			types2.QueryAuctionsRequest{
				Owner: addrs[0].String(),
			},
			auctions[2:4],
		},
		{
			"owner and denom",
			types2.QueryAuctionsRequest{
				Owner: addrs[0].String(),
				Denom: "jolt",
			},
			auctions[3:4],
		},
		{
			"owner, denom, type, phase",
			types2.QueryAuctionsRequest{
				Owner: addrs[0].String(),
				Denom: "jolt",
				Type:  types2.CollateralAuctionType,
				Phase: types2.ForwardAuctionPhase,
			},
			auctions[3:4],
		},
	}

	for _, tc := range tests {
		t.Run(tc.giveName, func(t *testing.T) {
			res, err := qs.Auctions(sdk.WrapSDKContext(ctx), &tc.giveRequest)
			require.NoError(t, err)

			var unpackedAuctions []types2.Auction

			for _, anyAuction := range res.Auctions {
				var auction types2.Auction
				err := tApp.AppCodec().UnpackAny(anyAuction, &auction)
				require.NoError(t, err)

				unpackedAuctions = append(unpackedAuctions, auction)
			}

			require.Equal(t, tc.wantResponse, unpackedAuctions)
		})
	}
}
