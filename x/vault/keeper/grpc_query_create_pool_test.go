package keeper_test

import (
	"testing"

	app2 "github.com/joltify-finance/joltify_lending/app"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/joltify-finance/joltify_lending/x/vault/types"
)

func TestCreatePoolQuerySingle(t *testing.T) {
	app2.SetSDKConfig()
	app, _, wctx := setupMsgServer(t)
	k := &app.VaultKeeper
	ctx := sdk.UnwrapSDKContext(wctx)

	var addresses []sdk.AccAddress
	for i := 0; i < 3; i++ {
		sk := ed25519.GenPrivKey()
		addr := sk.PubKey().Address().Bytes()
		addresses = append(addresses, addr)
	}

	msgs := createNCreatePool(k, ctx, 2, addresses)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetCreatePoolRequest
		response *types.QueryGetCreatePoolResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetCreatePoolRequest{Index: msgs[0].BlockHeight},
			response: &types.QueryGetCreatePoolResponse{CreatePool: msgs[0].Proposal[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetCreatePoolRequest{Index: msgs[1].BlockHeight},
			response: &types.QueryGetCreatePoolResponse{CreatePool: msgs[1].Proposal[0]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetCreatePoolRequest{Index: "missing"},
			err:     status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			response, err := k.CreatePool(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestCreatePoolQueryPaginated(t *testing.T) {
	var addresses []sdk.AccAddress
	for i := 0; i < 3; i++ {
		sk := ed25519.GenPrivKey()
		addr := sk.PubKey().Address().Bytes()
		addresses = append(addresses, addr)
	}

	app, _, wctx := setupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(wctx)
	k := &app.VaultKeeper
	msgs := createNCreatePool(k, ctx, 5, addresses)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllCreatePoolRequest {
		return &types.QueryAllCreatePoolRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := k.CreatePoolAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			for j := i; j < len(msgs) && j < i+step; j++ {
				assert.Equal(t, msgs[j].Proposal[0].PoolPubKey, resp.CreatePool[j-i].PoolPubKey)
			}
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := k.CreatePoolAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			for j := i; j < len(msgs) && j < i+step; j++ {
				assert.Equal(t, msgs[j].Proposal[0].PoolPubKey, resp.CreatePool[j-i].PoolPubKey)
			}
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := k.CreatePoolAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := k.CreatePoolAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
