package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
	"github.com/spf13/cobra"
)

func CmdCreateOutboundTx() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-outbound-tx [request-id] [outboundtx] [blockheight]",
		Short: "Create a new outboundTx",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexRequestID := args[0]
			outboundTx := args[1]
			blockHeight := args[2]
			chainType := args[3]
			inTxHash := args[4]
			receiverAddr, err := sdk.AccAddressFromBech32(args[5])
			if err != nil {
				return err
			}
			needMint := false
			if chainType == "OPPY" {
				needMint = true
			}
			// Get value arguments

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateOutboundTx(
				clientCtx.GetFromAddress(),
				indexRequestID,
				outboundTx,
				blockHeight,
				chainType,
				needMint,
				inTxHash,
				receiverAddr,
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
