package cli

import (
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/joltify-finance/joltify_lending/x/kyc/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdUploadInvestor() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upload-investor [investor-id] [wallet-address]",
		Short: "upload-investor [investor-id] [wallet-address]",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argInvestorId := args[0]
			argWalletAddress := strings.Split(args[1], ",")

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUploadInvestor(
				clientCtx.GetFromAddress().String(),
				argInvestorId,
				argWalletAddress,
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
