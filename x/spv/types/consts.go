package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	Maxfreq = 31536000
	Minfreq = 3600 * 24
)

var (
	RESERVEFACTOR = sdk.NewDecWithPrec(15, 2)
)
