package ibc_rate_limit

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/ibc-rate-limit/types"
)

type Querier struct {
	K ICS4Wrapper
}

func (q Querier) Params(ctx sdk.Context,
	req *types.QueryParamsRequest,
) (*types.QueryParamsResponse, error) {
	params := q.K.GetParams(ctx)
	return &types.QueryParamsResponse{Params: params}, nil
}
