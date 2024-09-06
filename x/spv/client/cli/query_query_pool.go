package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

const flagHistory = "history"

func CmdQueryPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query-pool [pool-index]",
		Short: "Query query-pool",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			history, _ := cmd.Flags().GetBool(flagHistory)
			reqPoolIndex := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryQueryPoolRequest{
				PoolIndex: reqPoolIndex,
			}
			var res *types.QueryQueryPoolResponse
			if history {
				res, err = queryClient.QueryHistoryPool(cmd.Context(), params)
				if err != nil {
					return err
				}
			} else {
				res, err = queryClient.QueryPool(cmd.Context(), params)
				if err != nil {
					return err
				}
			}
			return clientCtx.PrintProto(res)
		},
	}
	cmd.Flags().String(flagHistory, "", "(optional) showing the pool history")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
