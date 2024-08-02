package upgrade

import (
	"context"
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/ibc-go/v8/modules/core/02-client/types"

	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

const (
	V014UpgradeName = "v014_upgrade"
)

func CreateUpgradeHandlerForV014Upgrade(
	mm *module.Manager,
	configurator module.Configurator,
	keeper *ibckeeper.Keeper,
	subspace paramtypes.Subspace,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, _plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		for i := 0; i < 5; i++ {
			sdk.UnwrapSDKContext(ctx).Logger().Info("we upgrade to v014")
		}

		var params types.Params

		//subspace.GetParamSet(sdk.UnwrapSDKContext(ctx), &params)

		//panow, err := keeper.ClientParams(ctx, nil)
		//if err != nil {
		//	panic(err)
		//}
		//fmt.Printf(">>>%v\n", panow.String())
		subspace.IterateKeys(sdk.UnwrapSDKContext(ctx), func(key []byte) bool {
			fmt.Printf(">>%v---%v\n", string(key), string(subspace.GetRaw(sdk.UnwrapSDKContext(ctx), key)))
			return false
		})

		//pa := keeper.ClientKeeper.GetParams(sdk.UnwrapSDKContext(ctx))

		fmt.Printf(">>>%v\n", params)

		//fmt.Printf(">>>>%v\n", pa)
		return mm.RunMigrations(ctx, configurator, vm)
	}
}
