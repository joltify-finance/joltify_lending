package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	Maxfreq         = 31536000
	Minfreq         = 60
	MaxLiquidattion = 12
	Senior          = "senior"
	Junior          = "junior"
)

var RESERVEFACTOR = sdk.NewDecWithPrec(15, 2)
