package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdOutstandingInterest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "outstanding-interest [wallet] [pool-index]",
		Short: "Query outstanding-interest",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqWallet := args[0]
			reqPoolIndex := args[1]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryOutstandingInterestRequest{
				Wallet:    reqWallet,
				PoolIndex: reqPoolIndex,
			}

			res, err := queryClient.OutstandingInterest(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
