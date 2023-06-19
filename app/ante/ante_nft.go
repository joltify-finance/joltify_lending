package ante

import (
	"fmt"

	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
	spvkeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/types"

	errorsmod "cosmossdk.io/errors"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.AnteDecorator = SPVNFTDecorator{}

// SPVNFTDecorator blocks certain msg types from being granted or executed within authz.
type SPVNFTDecorator struct {
	spvKeeper spvkeeper.Keeper
}

// NewSPVNFTDecorator creates a decorator to block spv nft from transferring.
func NewSPVNFTDecorator(keeper spvkeeper.Keeper) SPVNFTDecorator {
	return SPVNFTDecorator{
		spvKeeper: keeper,
	}
}

func (sd SPVNFTDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	err = sd.checkForNFTMsg(ctx, tx.GetMsgs())
	if err != nil {
		return ctx, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "%v", err)
	}
	return next(ctx, tx, simulate)
}

func (sd SPVNFTDecorator) isSPVNFT(ctx sdk.Context, classid string) bool {
	found := false
	sd.spvKeeper.IteratePool(ctx, func(pool types.PoolInfo) (stop bool) {
		for _, el := range pool.PoolNFTIds {
			if el == classid {
				found = true
				return true
			}
		}
		return false
	})
	return found
}

// we disable all spv nft transfer
func (sd SPVNFTDecorator) checkForNFTMsg(ctx sdk.Context, msgs []sdk.Msg) error {
	for _, msg := range msgs {
		if nftMsg, ok := msg.(*nfttypes.MsgSend); ok {
			if sd.isSPVNFT(ctx, nftMsg.ClassId) {
				return fmt.Errorf("found disabled spv nft: %s", nftMsg.String())
			}
		}
	}
	return nil
}
