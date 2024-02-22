package keeper_test

import (
	"encoding/base64"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/cosmos/gogoproto/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/utils"
	"github.com/joltify-finance/joltify_lending/x/kyc/types"
	"github.com/stretchr/testify/require"
)

func TestCreateProject(t *testing.T) {
	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)

	_, k, wctx := setupMsgServer(t)
	ctx := sdk.UnwrapSDKContext(wctx)

	allProjects := types.GenerateTestProjects()

	unauthorised := types.MsgCreateProject{
		Creator: "jolt1p3jl6udk43vw0cvc5hjqrpnncsqmsz56wd32z8",
	}

	mockGovAddr := sdk.AccAddress([]byte("gov"))
	invalid := types.MsgCreateProject{
		Creator:        mockGovAddr.String(),
		EncodedProject: "encodedProject",
	}
	_, err := k.CreateProject(ctx, &unauthorised)
	require.ErrorContains(t, err, "invalid authority, expected")

	_, err = k.CreateProject(ctx, &invalid)
	require.ErrorContains(t, err, "fail to decode the project base64 string")

	var duplicateProject string
	for i, el := range allProjects {
		data, err := proto.Marshal(el)
		require.NoError(t, err)
		encoded := base64.StdEncoding.EncodeToString(data)
		if i == 3 {
			duplicateProject = encoded
		}
		msg := types.MsgCreateProject{
			Creator:        mockGovAddr.String(),
			EncodedProject: encoded,
		}
		_, err = k.CreateProject(ctx, &msg)
		require.NoError(t, err)
	}

	p, found := k.GetProject(ctx, 2)
	require.True(t, found)
	require.Equal(t, p.Index, int32(2))

	// put the duplicate project
	_, err = k.CreateProject(ctx, &types.MsgCreateProject{
		Creator:        mockGovAddr.String(),
		EncodedProject: duplicateProject,
	},
	)
	require.NoError(t, err)
	result, err := k.ListAllProjects(wctx, &types.ListAllProjectsRequest{Pagination: &query.PageRequest{Limit: 10, Offset: 0}})
	require.NoError(t, err)
	require.Equal(t, int32(101), result.TotalNumber)
}
