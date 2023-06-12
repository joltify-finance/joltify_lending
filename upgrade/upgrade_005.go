package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	vaultmodulekeeper "github.com/joltify-finance/joltify_lending/x/vault/keeper"
)

const V005UpgradeName = "v005_upgrade"

func CreateUpgradeHandlerForV005Upgrade(
	mm *module.Manager,
	k *vaultmodulekeeper.Keeper,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		for i := 0; i < 5; i++ {
			ctx.Logger().Info("we update the parameter for v005")
		}
		// we need to correct the ajolt to ujolt
		params := k.GetParams(ctx)

		// 30 bnb, 50000 ausdt, 500000 abusd, 100000 aeth, 500000 ajolt
		newQuota := "30000000000000000000abnb,50000000000000000000000ausdt,50000000000000000000000abusd,10000000000000000000aeth,500000000000ujolt"
		coins, err := sdk.ParseCoinsNormalized(newQuota)
		if err != nil {
			panic(err)
		}
		params.TargetQuota = coins
		ctx.Logger().Info("the new quota is", "token", params.TargetQuota)
		k.SetParams(ctx, params)
		ret, err := mm.RunMigrations(ctx, configurator, vm)
		panic("stop here")
		return ret, err
	}
}
