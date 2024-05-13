package keeper

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

func denomConvertToLocalAndUsd(in string) (string, string) {
	outs := strings.Split(in, "-")

	if len(outs) != 2 {
		return "", ""
	}
	return outs[0], outs[1]
}

func denomConvertToMarketID(in string) string {
	return in + ":usd"
}

func deleteElement(slice []sdk.AccAddress, element sdk.AccAddress) []sdk.AccAddress {
	for i, val := range slice {
		if val.Equals(element) {
			// Found the element, delete it by creating a new slice without it
			return append(slice[:i], slice[i+1:]...)
		}
	}
	// Element not found, return original slice
	return slice
}

func isSupportedTokens(token string) bool {
	supported := strings.Split(types.SupportedToken, ",")
	for _, val := range supported {
		if val == token {
			return true
		}
	}
	return false
}
