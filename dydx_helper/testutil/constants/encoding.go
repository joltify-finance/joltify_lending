package constants

import "github.com/joltify-finance/joltify_lending/dydx_helper/testutil/encoding"

var (
	TestEncodingCfg = encoding.GetTestEncodingCfg()
	TestTxBuilder   = TestEncodingCfg.TxConfig.NewTxBuilder()
)
