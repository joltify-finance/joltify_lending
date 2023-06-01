package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdwithdrawalPrincipal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdrawal-principal [pool-index] [wallet-address]",
		Short: "Query withdrawal-principal",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqPoolIndex := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QuerywithdrawalPrincipalRequest{
				PoolIndex:     reqPoolIndex,
				WalletAddress: args[1],
			}

			res, err := queryClient.WithdrawalPrincipal(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
