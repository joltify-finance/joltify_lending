package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdListInvestors() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-pools",
		Short: "Query list-pools",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			history, _ := cmd.Flags().GetBool(flagHistory)
			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryListPoolsRequest{}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			params.Pagination = pageReq

			var res *types.QueryListPoolsResponse
			if history {
				res, err = queryClient.ListHistoryPools(cmd.Context(), params)
			} else {
				res, err = queryClient.ListPools(cmd.Context(), params)
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
