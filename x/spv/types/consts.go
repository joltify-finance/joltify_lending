package types

const (
	Maxfreq         = 31536000
	Minfreq         = 60
	MaxLiquidattion = 12
	Senior          = "senior"
	Junior          = "junior"
	JOLTPRECISION   = 1e6
)

var RESERVEFACTOR = sdkmath.LegacyNewDecWithPrec(15, 2)
