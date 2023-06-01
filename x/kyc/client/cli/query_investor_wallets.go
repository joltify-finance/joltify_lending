package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/joltify-finance/joltify_lending/x/kyc/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdQueryInvestorWallets() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query-investor-wallets [investor-id]",
		Short: "Query query-investor-wallets",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqInvestorId := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryInvestorWalletsRequest{
				InvestorId: reqInvestorId,
			}

			res, err := queryClient.QueryInvestorWallets(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
