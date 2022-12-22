package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/cosmos/ibc-go/v3/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v3/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
)

const V101UpgradeName = "v101_upgrade"

func CreateUpgradeHandlerForV110Upgrade(
	mm *module.Manager,
	ibcTransfer ibctransferkeeper.Keeper,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		for i := 0; i < 2; i++ {
			ctx.Logger().Info("we update the parameter for v101")
		}
		transferModule := transfer.NewAppModule(ibcTransfer)
		vm[ibctransfertypes.ModuleName] = transferModule.ConsensusVersion()
		return mm.RunMigrations(ctx, configurator, vm)
	}
}
