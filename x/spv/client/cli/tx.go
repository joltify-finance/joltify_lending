package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

// var (
//	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
// )

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreatePool())
	cmd.AddCommand(CmdAddInvestors())
	cmd.AddCommand(CmdDeposit())
	cmd.AddCommand(CmdBorrow())
	cmd.AddCommand(CmdRepayInterest())
	cmd.AddCommand(CmdClaimInterest())
	cmd.AddCommand(CmdUpdatePool())
	cmd.AddCommand(CmdActivePool())
	cmd.AddCommand(CmdPayPrincipal())
	cmd.AddCommand(CmdPayPrincipalPartial())
	cmd.AddCommand(CmdWithdrawPrincipal())
	cmd.AddCommand(CmdSubmitWitdrawProposal())
	cmd.AddCommand(CmdTransferOwnership())
	cmd.AddCommand(CmdLiquidate())
	// this line is used by starport scaffolding # 1

	return cmd
}
