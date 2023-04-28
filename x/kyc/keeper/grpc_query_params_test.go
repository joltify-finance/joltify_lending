package keeper_test

import (
	"testing"

	"github.com/joltify-finance/joltify_lending/app"
	"github.com/joltify-finance/joltify_lending/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/joltify-finance/joltify_lending/testutil/keeper"
	"github.com/joltify-finance/joltify_lending/x/kyc/types"
	"github.com/stretchr/testify/require"
)

func newParams() types.Params {
	b := types.BasicInfo{
		Description:    "This is the test info",
		ProjectsUrl:    "empty",
		ProjectCountry: "ABC",
		BusinessNumber: "ABC123",
		Reserved:       []byte("reserved"),
	}

	acc, err := sdk.AccAddressFromBech32("jolt1gh6fnh6xt8lzhqy6z8n32lh7esxfrmspey8tp6")
	if err != nil {
		panic(err)
	}
	pi := types.ProjectInfo{
		Index:        1,
		SPVName:      "defaultSPV",
		BasicInfo:    &b,
		ProjectOwner: acc,
	}
	return types.Params{ProjectsInfo: []*types.ProjectInfo{&pi}, Submitter: []sdk.AccAddress{acc}}
}

func TestParamsQuery(t *testing.T) {
	config := app.SetSDKConfig()
	utils.SetBech32AddressPrefixes(config)
	keeper, ctx := testkeeper.KycKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := newParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)

	require.True(t, params.Submitter[0].Equals(response.Params.Submitter[0]))
	require.EqualValues(t, params.ProjectsInfo[0].SPVName, response.GetParams().ProjectsInfo[0].SPVName)
}
