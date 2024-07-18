package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/cosmos/cosmos-sdk/x/group/client/cli"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/joltify-finance/joltify_lending/x/third_party/wasmx/types"
)

// GetQueryCmd returns the parent command for all modules/wasmx CLi query commands.
func GetQueryCmd() *cobra.Command {
	cmd := cli.ModuleRootCommand(types.ModuleName, true)

	cmd.AddCommand(
		GetWasmxParamsCmd(),
		GetContractInfoCmd(),
	)
	return cmd
}

func GetWasmxParamsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Gets wasmx params info.",
		Long:  "Gets wasmx params info.",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryWasmxParamsRequest{}

			res, err := queryClient.WasmxParams(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetContractInfoCmd() *cobra.Command {
	cmd := cli.QueryCmd(
		"contract-info <contract-address>",
		"Gets contract ingo",
		types.NewQueryClient,
		&types.QueryContractRegistrationInfoRequest{}, nil, nil,
	)

	return cmd
}
