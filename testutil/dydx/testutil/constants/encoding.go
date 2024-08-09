package constants

import "github.com/joltify-finance/joltify_lending/testutil/dydx/testutil/encoding"

var (
	TestEncodingCfg = encoding.GetTestEncodingCfg()
	TestTxBuilder   = TestEncodingCfg.TxConfig.NewTxBuilder()
)
