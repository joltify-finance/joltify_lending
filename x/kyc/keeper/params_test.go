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
	require.EqualValues(t, params, k.GetParams(ctx))
}

func TestGetEach(t *testing.T) {
	config := sdk.GetConfig()
	utils.SetBech32AddressPrefixes(config)
	k, ctx := testkeeper.KycKeeper(t)
	params := newParams()

	k.SetParams(ctx, params)
	projects := k.GetProjects(ctx)
	require.EqualValues(t, projects, params.ProjectsInfo)
	submitters := k.GetSubmitter(ctx)
	require.EqualValues(t, submitters, params.Submitter)
}
