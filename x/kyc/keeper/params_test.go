package keeper_test

import (
	"encoding/base64"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/joltify-finance/joltify_lending/x/kyc/types"

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
	require.EqualValues(t, params.ProjectInfo, k.GetParams(ctx).ProjectInfo)
}

func TestGetEach(t *testing.T) {
	config := sdk.GetConfig()
	utils.SetBech32AddressPrefixes(config)
	k, ctx := testkeeper.KycKeeper(t)
	params := newParams()

	k.SetParams(ctx, params)
	projects := k.GetProjects(ctx)

	mb, err := base64.StdEncoding.DecodeString(params.ProjectInfo)
	require.NoError(t, err)
	var decodedProjects types.Projects
	err = proto.Unmarshal(mb, &decodedProjects)
	require.NoError(t, err)

	require.EqualValues(t, projects[0].SPVName, decodedProjects.Items[0].SPVName)
	submitters := k.GetSubmitter(ctx)
	require.EqualValues(t, submitters, params.Submitter)
}
