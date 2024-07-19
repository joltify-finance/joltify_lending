package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"

	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/joltify-finance/joltify_lending/testutil/keeper"
	"github.com/joltify-finance/joltify_lending/testutil/nullify"
	"github.com/joltify-finance/joltify_lending/x/quota/types"
)

func TestQuotaQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.QuotaKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs1 := createNQuota("testmodule1", 5)
	keeper.SetQuotaData(ctx, msgs1)

	msgs2 := createNQuota("testmodule2", 5)
	keeper.SetQuotaData(ctx, msgs2)

	msgs3 := createNQuota("testmodule3", 5)
	keeper.SetQuotaData(ctx, msgs3)

	tests := []struct {
		desc     string
		request  *types.QueryGetQuotaRequest
		response *types.QueryGetQuotaResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetQuotaRequest{QuotaModuleName: "testmodule1"},
			response: &types.QueryGetQuotaResponse{Quota: msgs1},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetQuotaRequest{QuotaModuleName: "testmodule2"},
			response: &types.QueryGetQuotaResponse{Quota: msgs2},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetQuotaRequest{QuotaModuleName: "testmodule4"},
			err:     errorsmod.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Quota(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestQuotaQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.QuotaKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)

	msgs1 := createNQuota("testmodule1", 50)
	keeper.SetQuotaData(ctx, msgs1)

	msgs2 := createNQuota("testmodule2", 50)
	keeper.SetQuotaData(ctx, msgs2)

	msgs3 := createNQuota("testmodule3", 50)
	keeper.SetQuotaData(ctx, msgs3)

	msgs4 := createNQuota("testmodule4", 50)
	keeper.SetQuotaData(ctx, msgs4)

	msgs := []types.CoinsQuota{msgs1, msgs2, msgs3, msgs4}

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllQuotaRequest {
		return &types.QueryAllQuotaRequest{
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
			resp, err := keeper.QuotaAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Quota), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Quota),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.QuotaAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Quota), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Quota),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.QuotaAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Quota),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.QuotaAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
