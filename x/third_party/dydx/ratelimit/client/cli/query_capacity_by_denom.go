package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/joltify-finance/joltify_lending/x/third_party/dydx/ratelimit/types"
	"github.com/spf13/cobra"
)

func CmdQueryCapacityByDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "capacity-by-denom",
		Short: "query the list of capacity and its corresponding limiter for each denom",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.CapacityByDenom(cmd.Context(), &types.QueryCapacityByDenomRequest{
				Denom: args[0],
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
