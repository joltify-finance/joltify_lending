package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdDepositor() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "depositor [pool-index] [wallet-address]",
		Short: "Query depositor",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqPoolIndex := args[0]
			reqWalletAddress := args[1]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryDepositorRequest{
				WalletAddress:    reqWalletAddress,
				DepositPoolIndex: reqPoolIndex,
			}

			res, err := queryClient.Depositor(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdDepositorHistory() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "depositor history [pool-index] [wallet-address]",
		Short: "Query depositor",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			spvWalletAddress := args[0]
			poolIndex := args[1]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryDepositorHistoryRequest{
				WalletAddress:    spvWalletAddress,
				DepositPoolIndex: poolIndex,
			}
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			params.Pagination = pageReq

			var res *types.QueryHistoryDepositorResponse

			res, err = queryClient.DepositorHistory(cmd.Context(), params)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
