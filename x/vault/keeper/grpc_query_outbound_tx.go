package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) OutboundTxAll(c context.Context, req *types.QueryAllOutboundTxRequest) (*types.QueryAllOutboundTxResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var results []types.QueryGetOutboundTxResponse
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	outboundTxStore := prefix.NewStore(store, types.KeyPrefix(types.OutboundTxKeyPrefix))

	pageRes, err := query.Paginate(outboundTxStore, req.Pagination, func(key []byte, value []byte) error {
		var outboundTx types.OutboundTx
		if err := k.cdc.Unmarshal(value, &outboundTx); err != nil {
			return err
		}

		var views []types.ProposalView
		outboundTxs := outboundTx.GetOutboundTxs()
		for _, el := range outboundTxs {
			proposals, found := k.GetOutboundTxProposal(ctx, outboundTx.Index, el)
			if found {
				view := types.ProposalView{OutboundTx: el, Proposals: proposals}
				views = append(views, view)
			}
		}
		item := types.QueryGetOutboundTxResponse{OutboundTx: outboundTx, View: views}
		results = append(results, item)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllOutboundTxResponse{AllOutbound: results, Pagination: pageRes}, nil
}

func (k Keeper) OutboundTx(c context.Context, req *types.QueryGetOutboundTxRequest) (*types.QueryGetOutboundTxResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetOutboundTx(
		ctx,
		req.RequestID,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}
	var views []types.ProposalView
	outboundTxs := val.GetOutboundTxs()
	for _, el := range outboundTxs {
		proposals, found := k.GetOutboundTxProposal(ctx, req.RequestID, el)
		if found {
			view := types.ProposalView{OutboundTx: el, Proposals: proposals}
			views = append(views, view)
		}
	}

	return &types.QueryGetOutboundTxResponse{OutboundTx: val, View: views}, nil
}
