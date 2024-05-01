package keeper

import (
	"fmt"

	coserrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

// ArchieveDepositor archives the depositor
func (k Keeper) ArchieveDepositor(ctx sdk.Context, depositor types.DepositorInfo) {
	k.SetDepositorHistory(ctx, depositor)
	k.DelDepositor(ctx, depositor)
}

// ArchiveNFT archives the NFT
func (k Keeper) ArchiveNFT(ctx sdk.Context, classID, nftID string) error {
	thisNFT, ok := k.NftKeeper.GetNFT(ctx, classID, nftID)
	if !ok {
		return coserrors.Wrap(types.ErrNFTNotFound, fmt.Sprintf("classID: %v, nftID: %v", classID, nftID))
	}
	owner := k.NftKeeper.GetOwner(ctx, classID, nftID)

	archiveClass, err := k.GetOrCreate(ctx, classID)
	if err != nil {
		return fmt.Errorf("failed to get or create the class: %v", err)
	}

	thisNFT.Id = fmt.Sprintf("%v%v-%v", types.ArchivePrefix, thisNFT.Id, ctx.BlockTime().Unix())
	thisNFT.ClassId = archiveClass.Id
	err = k.NftKeeper.Mint(ctx, thisNFT, owner)
	if err != nil {
		return fmt.Errorf("failed to update the nft: %v", err)
	}

	if err := k.NftKeeper.Burn(ctx, classID, nftID); err != nil {
		return fmt.Errorf("failed to burn the nft: %v", err)
	}
	return nil
}

func (k Keeper) GetOrCreate(ctx sdk.Context, classID string) (nfttypes.Class, error) {
	archiveClassID := fmt.Sprintf("%v%v", types.ArchivePrefix, classID)

	var ok bool
	var thisClass nfttypes.Class
	thisClass, ok = k.NftKeeper.GetClass(ctx, archiveClassID)
	if !ok {
		thisClass, ok = k.NftKeeper.GetClass(ctx, classID)
		if !ok {
			panic("should never failed to get the class")
		}
		thisClass.Id = archiveClassID
		err := k.NftKeeper.SaveClass(ctx, thisClass)
		if err != nil {
			panic("should never failed to save the class" + err.Error())
		}
	}
	return thisClass, nil
}

// ArchiveClass archives the class
func (k Keeper) ArchiveClass(ctx sdk.Context, classID string) {
	_, err := k.GetOrCreate(ctx, classID)
	if err != nil {
		panic("should never failed to get or create the class" + err.Error())
	}
	thisClass, ok := k.NftKeeper.GetClass(ctx, classID)
	if !ok {
		panic("fail to get the class")
	}
	err = k.NftKeeper.UpdateClass(ctx, thisClass)
	if err != nil {
		panic("should never failed to update the class" + err.Error())
	}
}

// ArchivePool archives the pool
func (k Keeper) ArchivePool(ctx sdk.Context, poolInfo types.PoolInfo) {
	k.DelPool(ctx, poolInfo.Index)
	poolInfo.Index = fmt.Sprintf("%v%v", types.ArchivePrefix, poolInfo.Index)
	for i, el := range poolInfo.PoolNFTIds {
		poolInfo.PoolNFTIds[i] = fmt.Sprintf("%v%v", types.ArchivePrefix, el)
		k.ArchiveClass(ctx, el)
	}
	k.SetHistoryPool(ctx, poolInfo)
}
