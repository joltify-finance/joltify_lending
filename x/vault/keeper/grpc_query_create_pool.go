package keeper

import (
	rawcontext "context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) calMinSupportNodes(c rawcontext.Context) int32 {
	ctx := sdk.UnwrapSDKContext(c)

	boundedValidators := k.vaultStaking.GetBondedValidatorsByPower(ctx)

	candidateDec := sdk.NewDecWithPrec(int64(len(boundedValidators)), 0)

	params := k.GetParams(ctx)
	candidateNumDec := candidateDec.Mul(params.CandidateRatio).MulTruncate(sdk.MustNewDecFromStr("0.6667"))
	candidateNum := int32(candidateNumDec.TruncateInt64())
	return candidateNum
}

func (k Keeper) CreatePoolAll(c rawcontext.Context, req *types.QueryAllCreatePoolRequest) (*types.QueryAllCreatePoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var proposals []*types.PoolProposal
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	createPoolStore := prefix.NewStore(store, types.KeyPrefix(types.CreatePoolKey))

	pageRes, err := query.Paginate(createPoolStore, req.Pagination, func(key []byte, value []byte) error {
		var createPool types.CreatePool
		if err := k.cdc.Unmarshal(value, &createPool); err != nil {
			return err
		}
		minSupportNodes := k.calMinSupportNodes(c)
		proposal := getProposal(createPool.Proposal, minSupportNodes)
		proposals = append(proposals, proposal)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCreatePoolResponse{CreatePool: proposals, Pagination: pageRes}, nil
}

func (k Keeper) CreatePool(c rawcontext.Context, req *types.QueryGetCreatePoolRequest) (*types.QueryGetCreatePoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	val, found := k.GetCreatePool(ctx, req.Index)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	minSupportNodes := k.calMinSupportNodes(c)

	proposal := getProposal(val.Proposal, minSupportNodes)
	return &types.QueryGetCreatePoolResponse{CreatePool: proposal}, nil
}

func (k Keeper) GetLastPool(c rawcontext.Context, req *types.QueryLatestPoolRequest) (*types.QueryLastPoolResponse, error) {
	var allProposal []*types.PoolInfo
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	lastTwoPools, found := k.GetLatestTwoPool(ctx)
	if !found {
		return &types.QueryLastPoolResponse{Pools: allProposal}, nil
	}

	minSupportNodes := k.calMinSupportNodes(c)
	for _, el := range lastTwoPools {
		proposalLast := getProposal(el.Proposal, minSupportNodes)
		if proposalLast != nil {
			lastProposal := types.PoolInfo{
				BlockHeight: el.BlockHeight,
				CreatePool:  proposalLast,
			}
			allProposal = append(allProposal, &lastProposal)
		}
	}

	return &types.QueryLastPoolResponse{Pools: allProposal}, nil
}

func getProposal(proposals []*types.PoolProposal, minSupportNodes int32) *types.PoolProposal {
	// since 2/3 nodes are honest, so we will not have 5/5 situation
	maxLength := 0
	proposalIndex := -1
	for index, proposal := range proposals {
		length := len(proposal.Nodes)
		if maxLength > length || int32(length) >= minSupportNodes {
			proposalIndex = index
			maxLength = length
		}
	}
	if proposalIndex == -1 {
		return nil
	}
	return proposals[proposalIndex]
}
