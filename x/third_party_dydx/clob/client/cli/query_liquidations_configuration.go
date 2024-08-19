package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/types"
	"github.com/spf13/cobra"
)

func CmdGetLiquidationsConfiguration() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-liquidations-config",
		Short: "get the liquidations configuration",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryLiquidationsConfigurationRequest{}

			res, err := queryClient.LiquidationsConfiguration(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}