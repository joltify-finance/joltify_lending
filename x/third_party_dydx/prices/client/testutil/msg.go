package testutil

import (
	pricescli "github.com/joltify-finance/joltify_lending/x/third_party_dydx/prices/client/cli"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
)

// MsgQueryAllMarketParamExec lists all markets params in `Prices`.
func MsgQueryAllMarketParamExec(clientCtx client.Context) (testutil.BufferWriter, error) {
	return clitestutil.ExecTestCLICmd(
		clientCtx,
		pricescli.CmdListMarketParam(),
		[]string{},
	)
}

// MsgQueryAllMarketPriceExec lists all markets prices in `Prices`.
func MsgQueryAllMarketPriceExec(clientCtx client.Context) (testutil.BufferWriter, error) {
	return clitestutil.ExecTestCLICmd(
		clientCtx,
		pricescli.CmdListMarketPrice(),
		[]string{},
	)
}