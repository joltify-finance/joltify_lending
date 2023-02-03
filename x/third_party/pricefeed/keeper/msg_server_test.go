package keeper_test

import (
	"testing"
	"time"

	tmlog "github.com/tendermint/tendermint/libs/log"

	"github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/keeper"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/stretchr/testify/require"
	tmprototypes "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestKeeper_PostPrice(t *testing.T) {
	_, addrs := app.GeneratePrivKeyAddressPairs(4)
	tApp := app.NewTestApp(tmlog.TestingLogger(), t.TempDir())
	ctx := tApp.NewContext(true, tmprototypes.Header{}).
		WithBlockTime(time.Now().UTC())
	k := tApp.GetPriceFeedKeeper()
	msgSrv := keeper.NewMsgServerImpl(k)

	authorizedOracles := addrs[:2]
	unauthorizedAddrs := addrs[2:]

	mp := types2.Params{
		Markets: []types2.Market{
			{MarketID: "tstusd", BaseAsset: "tst", QuoteAsset: "usd", Oracles: authorizedOracles, Active: true},
		},
	}
	k.SetParams(ctx, mp)

	now := time.Now().UTC()

	tests := []struct {
		giveMsg      string
		giveOracle   sdk.AccAddress
		giveMarketId string
		giveExpiry   time.Time
		wantAccepted bool
		errorKind    error
	}{
		{"authorized", authorizedOracles[0], "tstusd", now.Add(time.Hour * 1), true, nil},
		{"expired", authorizedOracles[0], "tstusd", now.Add(-time.Hour * 1), false, types2.ErrExpired},
		{"invalid", authorizedOracles[0], "invalid", now.Add(time.Hour * 1), false, types2.ErrInvalidMarket},
		{"unauthorized", unauthorizedAddrs[0], "tstusd", now.Add(time.Hour * 1), false, types2.ErrInvalidOracle},
	}

	for _, tt := range tests {
		t.Run(tt.giveMsg, func(t *testing.T) {
			// Use MsgServer over keeper methods directly to tests against valid oracles
			msg := types2.NewMsgPostPrice(tt.giveOracle.String(), tt.giveMarketId, sdk.MustNewDecFromStr("0.5"), tt.giveExpiry)
			_, err := msgSrv.PostPrice(sdk.WrapSDKContext(ctx), msg)

			if tt.wantAccepted {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.ErrorIs(t, tt.errorKind, err)
			}
		})
	}
}
