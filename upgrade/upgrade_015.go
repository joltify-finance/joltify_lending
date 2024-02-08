package v1

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/gogo/protobuf/proto"
	spvmodulekeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"
	spvmoduletypes "github.com/joltify-finance/joltify_lending/x/spv/types"
)

const V015UpgradeName = "v015_upgrade"

func CreateUpgradeHandlerForV015Upgrade(
	mm *module.Manager,
	configurator module.Configurator,
	spvKeeper spvmodulekeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		for i := 0; i < 5; i++ {
			ctx.Logger().Info("we upgrade to v015")
		}

		bigCounter := 0

		var poolIndexs []string
		spvKeeper.IteratePool(ctx, func(poolInfo spvmoduletypes.PoolInfo) bool {
			poolIndexs = append(poolIndexs, poolInfo.Index)
			return false
		})

		for _, poolIndex := range poolIndexs {
			spvKeeper.IterateDepositors(ctx, poolIndex, func(depositor spvmoduletypes.DepositorInfo) bool {
				nfts := depositor.LinkedNFT
				var savelist []string
				for _, el := range nfts {
					ids := strings.Split(el, ":")
					thisNFT, found := spvKeeper.NftKeeper.GetNFT(ctx, ids[0], ids[1])
					if !found {
						panic("should never fail to find the nft")
					}
					var nftinfo spvmoduletypes.NftInfo
					err := proto.Unmarshal(thisNFT.Data.Value, &nftinfo)
					if err != nil {
						panic(err)
					}

					if !nftinfo.Borrowed.Amount.IsPositive() {
						bigCounter++
						fmt.Printf("nft %s is zero and to be removed\n", nftinfo.Receiver)
						err = spvKeeper.NftKeeper.Burn(ctx, ids[0], ids[1])
						if err != nil {
							panic(err)
						}
						continue
					}
					savelist = append(savelist, el)
				}
				depositor.LinkedNFT = savelist
				spvKeeper.SetDepositor(ctx, depositor)
				return false
			})
		}
		return mm.RunMigrations(ctx, configurator, vm)
	}
}
