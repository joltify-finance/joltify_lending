package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	Maxfreq = 1310400
	Minfreq = 100800
)

var (
	RESERVEFACTOR = sdk.NewDecWithPrec(15, 2)
)
