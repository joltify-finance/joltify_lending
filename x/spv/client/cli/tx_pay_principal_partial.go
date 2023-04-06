package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdPayPrincipalPartial() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pay-principal-partial [pool-index] [amount]",
		Short: "Broadcast message pay-principal-partial",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPoolIndex := args[0]

			argToken, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgPayPrincipalPartial(
				clientCtx.GetFromAddress().String(),
				argPoolIndex,
				argToken,
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
