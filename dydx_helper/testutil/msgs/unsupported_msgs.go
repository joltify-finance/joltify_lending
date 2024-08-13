package msgs

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	govbeta "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/joltify-finance/joltify_lending/dydx_helper/testutil/constants"
	"github.com/joltify-finance/joltify_lending/dydx_helper/testutil/encoding"
)

func init() {
	testEncodingCfg := encoding.GetTestEncodingCfg()
	testTxBuilder := testEncodingCfg.TxConfig.NewTxBuilder()

	_ = testTxBuilder.SetMsgs(GovBetaMsgSubmitProposal)
	GovBetaMsgSubmitProposalTxBytes, _ = testEncodingCfg.TxConfig.TxEncoder()(testTxBuilder.GetTx())
}

var (
	govbetaContent, _ = govbeta.ContentFromProposalType("test-title", "test-desc", "Text")

	GovBetaMsgSubmitProposal, _ = govbeta.NewMsgSubmitProposal(
		govbetaContent,
		sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 1000)),
		constants.BobAccAddress,
	)
	GovBetaMsgSubmitProposalTxBytes []byte
)
