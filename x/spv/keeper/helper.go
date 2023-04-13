package keeper

import "strings"

func denomConvert(in string) string {
	outs := strings.Split(in, "-")

	if len(outs) != 2 {
		return ""
	}
	return outs[1]
}
