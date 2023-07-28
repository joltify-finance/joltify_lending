package v1

import (
	"fmt"

	"github.com/evmos/ethermint/x/evm"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authKeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
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
	accKeeper authKeeper.AccountKeeper,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		//
		//dbprovider := node.DefaultDBProvider
		//stateDB, err := dbprovider(&node.DBContext{ID: "state", Config: })
		//if err != nil {
		//	panic(err)
		//}

		fmt.Printf(">>>>>>>>>>>>>>hello>>>>>>>>>>>>>>>\n")

		ctx = ctx.WithChainID("joltify_1729-1")
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

		genStatus := evmtypes.DefaultGenesisState()
		genStatus.Params = evmParams
		fmt.Printf(">>>>>>>>>>>>>>>>>>>>we apply!!!!")
		evm.InitGenesis(ctx, evmK, accKeeper, *genStatus)

		vm[evmtypes.ModuleName] = evm.AppModule{}.ConsensusVersion()
		return mm.RunMigrations(ctx, configurator, vm)
	}
}
