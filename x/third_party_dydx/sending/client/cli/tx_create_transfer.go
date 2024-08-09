package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	assettypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/assets/types"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/sending/types"
	satypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/subaccounts/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateTransfer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-transfer sender_owner sender_number recipient_owner recipient_number quantums",
		Short: "Broadcast message CreateTransfer",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSenderOwner := args[0]
			argSenderNumber, err := cast.ToUint32E(args[1])
			if err != nil {
				return err
			}

			argRecipientOwner := args[2]
			argRecipientNumber, err := cast.ToUint32E(args[3])
			if err != nil {
				return err
			}

			argAmount, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}

			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateTransfer(
				&types.Transfer{
					Sender: satypes.SubaccountId{
						Owner:  argSenderOwner,
						Number: argSenderNumber,
					},
					Recipient: satypes.SubaccountId{
						Owner:  argRecipientOwner,
						Number: argRecipientNumber,
					},
					AssetId: assettypes.AssetUsdc.Id,
					Amount:  argAmount,
				},
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
