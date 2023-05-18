package cli

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreatePool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-pool [pool title] [project index] [apy junior] [apy senior] [target junior] [target senior]",
		Short: "Broadcast message create-pool",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			name := args[0]

			projectIndex, err := strconv.ParseInt(args[1], 10, 32)
			if err != nil {
				return err
			}

			argApy := args[2]
			argApy2 := args[3]

			argTarget, err := sdk.ParseCoinNormalized(args[4])
			if err != nil {
				return err
			}
			argTarget2, err := sdk.ParseCoinNormalized(args[5])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreatePool(
				clientCtx.GetFromAddress().String(),
				name,
				int32(projectIndex),
				[]string{argApy, argApy2},
				[]sdk.Coin{argTarget, argTarget2},
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
