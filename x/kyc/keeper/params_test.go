package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/utils"

	testkeeper "github.com/joltify-finance/joltify_lending/testutil/keeper"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	config := sdk.GetConfig()
	utils.SetBech32AddressPrefixes(config)
	k, ctx := testkeeper.KycKeeper(t)
	params := newParams()
	k.SetParams(ctx, params)
	require.EqualValues(t, params.Submitter, k.GetParams(ctx).Submitter)
	require.EqualValues(t, params.ProjectsInfo[0].SPVName, k.GetParams(ctx).ProjectsInfo[0].SPVName)
}

func TestGetEach(t *testing.T) {
	config := sdk.GetConfig()
	utils.SetBech32AddressPrefixes(config)
	k, ctx := testkeeper.KycKeeper(t)
	params := newParams()

	k.SetParams(ctx, params)
	projects := k.GetProjects(ctx)
	require.EqualValues(t, projects[0].SPVName, params.ProjectsInfo[0].SPVName)
	submitters := k.GetSubmitter(ctx)
	require.EqualValues(t, submitters, params.Submitter)
}
