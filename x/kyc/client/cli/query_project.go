package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/joltify-finance/joltify_lending/x/kyc/types"
	"github.com/spf13/cobra"
)

func CmdQueryProject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query-project [project index]",
		Short: "Query query-project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqProject := args[0]
			val, err := strconv.ParseInt(reqProject, 10, 32)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryProjectRequest{
				ProjectId: int32(val),
			}

			res, err := queryClient.QueryProject(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdListProject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-project",
		Short: "Query list-project",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.ListAllProjectsRequest{}
			res, err := queryClient.ListAllProjects(cmd.Context(), params)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
