package types_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	appconfig "github.com/joltify-finance/joltify_lending/app/config"
)

func init() {
	kavaConfig := sdk.GetConfig()
	appconfig.SetupConfig()
	//app.SetBip44CoinType(kavaConfig)
	kavaConfig.Seal()
}
