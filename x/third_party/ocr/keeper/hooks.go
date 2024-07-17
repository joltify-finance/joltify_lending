package keeper

import (
	"github.com/joltify-finance/joltify_lending/x/third_party/ocr/types"
)

type OcrHooks interface {
	SetHooks(h types.OcrHooks)
}

// Set the hooks
func (k *Keeper) SetHooks(h types.OcrHooks) {
	if k.hooks != nil {
		panic("cannot set hooks twice")
	}

	k.hooks = h
}
