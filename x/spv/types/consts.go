package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	Maxfreq = 24
	Minfreq = 1
)

var (
	RESERVEFACTOR = sdk.NewDecWithPrec(15, 2)
)
