package kyc_test

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"testing"

	sdkmath "cosmossdk.io/math"
	tmrand "github.com/cometbft/cometbft/libs/rand"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/gogoproto/proto"
	keepertest "github.com/joltify-finance/joltify_lending/testutil/keeper"
	"github.com/joltify-finance/joltify_lending/testutil/nullify"
	"github.com/joltify-finance/joltify_lending/utils"
	"github.com/joltify-finance/joltify_lending/x/kyc"
	"github.com/joltify-finance/joltify_lending/x/kyc/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	config := sdk.GetConfig()
	utils.SetBech32AddressPrefixes(config)
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.KycKeeper(t)
	kyc.InitGenesis(ctx, *k, genesisState)
	got := kyc.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

func TestT2(t *testing.T) {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("jolt", "joltpub")
	acc, err := sdk.AccAddressFromBech32("jolt15qdefkmwswysgg4qxgqpqr35k3m49pkxu8ygkq")
	require.NoError(t, err)
	b := types.BasicInfo{
		"This is the test info for Joshua AAAAA",
		"empty",
		"ABCCCCCCCCCCCCCCCCCCCc",
		"ABC123ddddd",
		[]byte("reserved"),
		"This is the Test Project 11111",
		"example@example.com",
		"example",
		"empty logo url",
		"empty project Brief",
		"empty project description",
	}
	pi := types.ProjectInfo{
		Index:                        int32(1 + 1),
		SPVName:                      strconv.Itoa(101) + ":" + tmrand.NewRand().Str(10),
		ProjectOwner:                 acc,
		BasicInfo:                    &b,
		ProjectLength:                480, // 5 mins
		SeparatePool:                 true,
		BaseApy:                      sdkmath.LegacyNewDecWithPrec(10, 2),
		PayFreq:                      "120",
		PoolLockedSeconds:            100,
		PoolTotalBorrowLimit:         100,
		MarketId:                     "aud:usd",
		WithdrawRequestWindowSeconds: 30,
		MinBorrowAmount:              sdkmath.NewInt(100),
	}

	data, err := proto.Marshal(&pi)
	require.NoError(t, err)

	out := base64.StdEncoding.EncodeToString(data)
	fmt.Printf("%v\n", out)

	aa := authtypes.NewModuleAddress(govtypes.ModuleName)
	fmt.Printf(">>>%v\n", aa.String())
}
