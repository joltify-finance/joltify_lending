package legacy

import (
	upgradetypes "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	evmkeeper "github.com/evmos/ethermint/x/evm/keeper"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	evmutilkeeper "github.com/joltify-finance/joltify_lending/x/third_party/evmutil/keeper"
	evmutiltypes "github.com/joltify-finance/joltify_lending/x/third_party/evmutil/types"
)

const V006UpgradeName = "v006_upgrade"

func CreateUpgradeHandlerForV006Upgrade(
	mm *module.Manager,
	evmutilsK evmutilkeeper.Keeper,
	evmK *evmkeeper.Keeper,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		for i := 0; i < 5; i++ {
			ctx.Logger().Info("we update the parameter for v006")
		}

		// params := evmkeeper.GetProposerAddress(ctx)
		evmutilsParams := evmutiltypes.DefaultParams()
		evmutilsK.SetParams(ctx, evmutilsParams)

		evmParams := evmtypes.DefaultParams()
		evmParams.EvmDenom = "ajolt"
		err := evmK.SetParams(ctx, evmParams)
		if err != nil {
			panic("set evm parameters failed")
		}

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
