package cli

import (
	"context"
	"fmt"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/issuance/types"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

// GetQueryCmd returns the cli query commands for the issuance module
func GetQueryCmd() *cobra.Command {
	issuanceQueryCmd := &cobra.Command{
		Use:   types2.ModuleName,
		Short: fmt.Sprintf("Querying commands for the %s module", types2.ModuleName),
	}

	cmds := []*cobra.Command{
		GetCmdQueryParams(),
	}

	for _, cmd := range cmds {
		flags.AddQueryFlagsToCmd(cmd)
	}

	issuanceQueryCmd.AddCommand(cmds...)

	return issuanceQueryCmd
}

// GetCmdQueryParams queries the issuance module parameters
func GetCmdQueryParams() *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Short: fmt.Sprintf("get the %s module parameters", types2.ModuleName),
		Long:  "Get the current issuance module parameters.",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types2.NewQueryClient(clientCtx)

			res, err := queryClient.Params(context.Background(), &types2.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(&res.Params)
		},
	}
}
