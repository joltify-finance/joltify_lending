package keeper_test

import (
	"strings"
	"time"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/issuance/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestIncrementCurrentAssetSupply() {
	type args struct {
		assets   []types2.Asset
		supplies []types2.AssetSupply
		coin     sdk.Coin
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
			"valid supply increase",
			args{
				assets: []types2.Asset{
					types2.NewAsset(suite.addrs[0], "usdtoken", []string{suite.addrs[1]}, false, true, types2.NewRateLimit(true, sdk.NewInt(10000000000), time.Hour*24)),
				},
				supplies: []types2.AssetSupply{
					types2.NewAssetSupply(sdk.NewCoin("usdtoken", sdk.ZeroInt()), time.Hour),
				},
				coin: sdk.NewCoin("usdtoken", sdk.NewInt(100000)),
			},
			errArgs{
				expectPass: true,
				contains:   "",
			},
		},
		{
			"over limit increase",
			args{
				assets: []types2.Asset{
					types2.NewAsset(suite.addrs[0], "usdtoken", []string{suite.addrs[1]}, false, true, types2.NewRateLimit(true, sdk.NewInt(10000000000), time.Hour*24)),
				},
				supplies: []types2.AssetSupply{
					types2.NewAssetSupply(sdk.NewCoin("usdtoken", sdk.ZeroInt()), time.Hour),
				},
				coin: sdk.NewCoin("usdtoken", sdk.NewInt(10000000001)),
			},
			errArgs{
				expectPass: false,
				contains:   "asset supply over limit",
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			params := types2.NewParams(tc.args.assets)
			suite.keeper.SetParams(suite.ctx, params)
			for _, supply := range tc.args.supplies {
				suite.keeper.SetAssetSupply(suite.ctx, supply, supply.GetDenom())
			}
			err := suite.keeper.IncrementCurrentAssetSupply(suite.ctx, tc.args.coin)
			if tc.errArgs.expectPass {
				suite.Require().NoError(err, tc.name)
				for _, expectedSupply := range tc.args.supplies {
					expectedSupply.CurrentSupply = expectedSupply.CurrentSupply.Add(tc.args.coin)
					actualSupply, found := suite.keeper.GetAssetSupply(suite.ctx, expectedSupply.GetDenom())
					suite.Require().True(found)
					suite.Require().Equal(expectedSupply, actualSupply, tc.name)
				}
			} else {
				suite.Require().Error(err, tc.name)
				suite.Require().True(strings.Contains(err.Error(), tc.errArgs.contains))
			}
		})
	}
}
