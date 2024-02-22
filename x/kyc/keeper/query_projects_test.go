package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/joltify-finance/joltify_lending/x/kyc/types"

	types2 "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/utils"
	"github.com/stretchr/testify/require"
)

func TestQueryProject(t *testing.T) {
	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)
	_, k, wctx := setupMsgServer(t)
	ctx := types2.UnwrapSDKContext(wctx)
	allProjects := types.GenerateTestProjects()

	for _, el := range allProjects {
		_, err := k.SetProject(ctx, el)
		require.NoError(t, err)
	}

	retAllProjects, err := k.QueryProject(wctx, &types.QueryProjectRequest{ProjectId: 1})
	require.NoError(t, err)
	require.Equal(t, int32(1), retAllProjects.Project.Index)

	retAllProjects, err = k.QueryProject(wctx, &types.QueryProjectRequest{ProjectId: 99})
	require.NoError(t, err)
	require.Equal(t, int32(99), retAllProjects.Project.Index)

	retAllProjects, err = k.QueryProject(wctx, &types.QueryProjectRequest{ProjectId: 120})
	require.Errorf(t, err, "rpc error: code = NotFound desc = project not found")
	require.Nil(t, retAllProjects)
}

func TestListProjects(t *testing.T) {
	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)
	_, k, wctx := setupMsgServer(t)
	ctx := types2.UnwrapSDKContext(wctx)
	allProjects := types.GenerateTestProjects()

	for _, el := range allProjects {
		_, err := k.SetProject(ctx, el)
		require.NoError(t, err)
	}

	req := &types.ListAllProjectsRequest{
		Pagination: &query.PageRequest{
			Limit:  10,
			Offset: 0,
		},
	}

	retAllProjects, err := k.ListAllProjects(wctx, req)
	require.NoError(t, err)
	require.Equal(t, 10, len(retAllProjects.Project))
	require.NotEmpty(t, retAllProjects.Project[0].Index)
}
