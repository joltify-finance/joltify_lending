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

var SupportedToken = "ibc/65D0BEC6DAD96C7F5043D1E54E54B6BB5D5B3AEC3FF6CEBB75B9E059F3580EA3" //nolint

var RESERVEFACTOR = sdk.NewDecWithPrec(15, 2)
