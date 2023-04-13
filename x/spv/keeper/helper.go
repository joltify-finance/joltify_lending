package keeper

import "strings"

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
