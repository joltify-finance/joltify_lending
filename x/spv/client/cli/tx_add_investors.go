package cli

import (
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAddInvestors() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-investors [pool-index] [investor ids]",
		Short: "Broadcast message add-investors",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			investorID := strings.Split(args[1], ",")

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddInvestors(
				clientCtx.GetFromAddress().String(),
				args[0],
				investorID,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
