package keeper

import "strings"

func denomConvert(in string) string {
	outs := strings.Split(in, "-")

	if len(outs) != 2 {
		return ""
	}
	return outs[1]
}

func denomConvertToMarketID(in string) string {
	if len(in) == 0 {
		return ""
	}
	return in[:len(in)-1] + ":usd"
}
