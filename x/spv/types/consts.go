package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	Maxfreq = 31536000
	Minfreq = 7257600
)

var (
	RESERVEFACTOR = sdk.NewDecWithPrec(15, 2)
)
