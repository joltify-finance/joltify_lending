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
}
