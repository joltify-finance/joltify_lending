package types_test

import (
	"strings"
	"testing"
	"time"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/issuance/types"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/app"
)

type GenesisTestSuite struct {
	suite.Suite

	addrs []string
}

func (suite *GenesisTestSuite) SetupTest() {
	_, addrs := app.GeneratePrivKeyAddressPairs(2)
	var strAddrs []string
	for _, addr := range addrs {
		strAddrs = append(strAddrs, addr.String())
	}
	suite.addrs = strAddrs
}

func (suite *GenesisTestSuite) TestValidate() {
	type args struct {
		assets   []types2.Asset
		supplies []types2.AssetSupply
	}
	type errArgs struct {
		expectPass bool
		contains   string
	}
	testCases := []struct {
		name    string
		args    args
		errArgs errArgs
	}{
		{
			"default",
			args{
				assets:   types2.DefaultAssets,
				supplies: types2.DefaultSupplies,
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"with asset",
			args{
				assets: []types2.Asset{
					types2.NewAsset(suite.addrs[0], "usdtoken", []string{suite.addrs[1]}, false, true, types2.NewRateLimit(false, sdk.ZeroInt(), time.Duration(0))),
				},
				supplies: []types2.AssetSupply{types2.NewAssetSupply(sdk.NewCoin("usdtoken", sdk.NewInt(1000000)), time.Hour)},
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"with asset rate limit",
			args{
				assets: []types2.Asset{
					types2.NewAsset(suite.addrs[0], "usdtoken", []string{suite.addrs[1]}, false, true, types2.NewRateLimit(true, sdk.NewInt(1000000000), time.Hour*24)),
				},
				supplies: []types2.AssetSupply{},
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"with multiple assets",
			args{
				assets: []types2.Asset{
					types2.NewAsset(suite.addrs[0], "usdtoken", []string{suite.addrs[1]}, false, true, types2.NewRateLimit(false, sdk.ZeroInt(), time.Duration(0))),
					types2.NewAsset(suite.addrs[0], "pegtoken", []string{suite.addrs[1]}, false, true, types2.NewRateLimit(false, sdk.ZeroInt(), time.Duration(0))),
				},
				supplies: []types2.AssetSupply{},
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"blocked owner",
			args{
				assets: []types2.Asset{
					types2.NewAsset(suite.addrs[0], "usdtoken", []string{suite.addrs[0]}, false, true, types2.NewRateLimit(false, sdk.ZeroInt(), time.Duration(0))),
				},
				supplies: []types2.AssetSupply{},
			},
			errArgs{
				expectPass: false,
				contains:   "asset owner cannot be blocked",
			},
		},
		{
			"empty owner",
			args{
				assets: []types2.Asset{
					types2.NewAsset("", "usdtoken", []string{suite.addrs[0]}, false, true, types2.NewRateLimit(false, sdk.ZeroInt(), time.Duration(0))),
				},
				supplies: []types2.AssetSupply{},
			},
			errArgs{
				expectPass: false,
				contains:   "owner must not be empty",
			},
		},
		{
			"empty blocked address",
			args{
				assets: []types2.Asset{
					types2.NewAsset(suite.addrs[0], "usdtoken", []string{""}, false, true, types2.NewRateLimit(false, sdk.ZeroInt(), time.Duration(0))),
				},
				supplies: []types2.AssetSupply{},
			},
			errArgs{
				expectPass: false,
				contains:   "blocked address must not be empty",
			},
		},
		{
			"invalid denom",
			args{
				assets: []types2.Asset{
					types2.NewAsset(suite.addrs[0], "USD2T ", []string{}, false, true, types2.NewRateLimit(false, sdk.ZeroInt(), time.Duration(0))),
				},
				supplies: []types2.AssetSupply{},
			},
			errArgs{
				expectPass: false,
				contains:   "invalid denom",
			},
		},
		{
			"duplicate denom",
			args{
				assets: []types2.Asset{
					types2.NewAsset(suite.addrs[0], "usdtoken", []string{suite.addrs[1]}, false, true, types2.NewRateLimit(false, sdk.ZeroInt(), time.Duration(0))),
					types2.NewAsset(suite.addrs[1], "usdtoken", []string{}, true, true, types2.NewRateLimit(false, sdk.ZeroInt(), time.Duration(0))),
				},
				supplies: []types2.AssetSupply{},
			},
			errArgs{
				expectPass: false,
				contains:   "duplicate asset denoms",
			},
		},
		{
			"duplicate asset",
			args{
				assets: []types2.Asset{
					types2.NewAsset(suite.addrs[0], "usdtoken", []string{suite.addrs[1]}, false, true, types2.NewRateLimit(false, sdk.ZeroInt(), time.Duration(0))),
					types2.NewAsset(suite.addrs[0], "usdtoken", []string{suite.addrs[1]}, false, true, types2.NewRateLimit(false, sdk.ZeroInt(), time.Duration(0))),
				},
				supplies: []types2.AssetSupply{},
			},
			errArgs{
				expectPass: false,
				contains:   "duplicate asset denoms",
			},
		},
		{
			"invalid block list",
			args{
				assets: []types2.Asset{
					types2.NewAsset(suite.addrs[0], "usdtoken", []string{suite.addrs[1]}, false, false, types2.NewRateLimit(false, sdk.ZeroInt(), time.Duration(0))),
				},
				supplies: []types2.AssetSupply{types2.NewAssetSupply(sdk.NewCoin("usdtoken", sdk.ZeroInt()), time.Hour)},
			},
			errArgs{
				expectPass: false,
				contains:   "blocked-list should be empty",
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			gs := types2.NewGenesisState(types2.NewParams(tc.args.assets), tc.args.supplies)
			err := gs.Validate()
			if tc.errArgs.expectPass {
				suite.Require().NoError(err, tc.name)
			} else {
				suite.Require().Error(err, tc.name)
				suite.Require().True(strings.Contains(err.Error(), tc.errArgs.contains))
			}
		})
	}
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GenesisTestSuite))
}
