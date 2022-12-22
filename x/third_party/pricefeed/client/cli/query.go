package cli

import (
	"context"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/types"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	// Group nameservice queries under a subcommand
	pricefeedQueryCmd := &cobra.Command{
		Use:                        types2.ModuleName,
		Short:                      "Querying commands for the pricefeed module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmds := []*cobra.Command{
		GetCmdPrice(),
		GetCmdQueryPrices(),
		GetCmdRawPrices(),
		GetCmdOracles(),
		GetCmdMarkets(),
		GetCmdQueryParams(),
	}

	for _, cmd := range cmds {
		flags.AddQueryFlagsToCmd(cmd)
	}

	pricefeedQueryCmd.AddCommand(cmds...)

	return pricefeedQueryCmd
}

// GetCmdOracles queries the oracle set of an asset
func GetCmdOracles() *cobra.Command {
	return &cobra.Command{
		Use:   "oracles [marketID]",
		Short: "get the oracle set for a market",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types2.NewQueryClient(clientCtx)

			marketID := args[0]

			params := types2.QueryOraclesRequest{
				MarketId: marketID,
			}

			res, err := queryClient.Oracles(context.Background(), &params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
}

// GetCmdPrice queries the current price of an asset
func GetCmdPrice() *cobra.Command {
	return &cobra.Command{
		Use:   "price [marketID]",
		Short: "get the current price for the input market",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types2.NewQueryClient(clientCtx)

			marketID := args[0]

			params := types2.QueryPriceRequest{
				MarketId: marketID,
			}

			res, err := queryClient.Price(context.Background(), &params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
}

// GetCmdQueryPrices queries the pricefeed module for current prices
func GetCmdQueryPrices() *cobra.Command {
	return &cobra.Command{
		Use:   "prices",
		Short: "get the current price of each market",
		Long:  "Get the current prices of each market in the pricefeed module.",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types2.NewQueryClient(clientCtx)

			res, err := queryClient.Prices(context.Background(), &types2.QueryPricesRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
}

// GetCmdRawPrices queries the current price of an asset
func GetCmdRawPrices() *cobra.Command {
	return &cobra.Command{
		Use:   "rawprices [marketID]",
		Short: "get the raw oracle prices for the input market",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types2.NewQueryClient(clientCtx)

			marketID := args[0]

			params := types2.QueryRawPricesRequest{
				MarketId: marketID,
			}

			res, err := queryClient.RawPrices(context.Background(), &params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
}

// GetCmdMarkets queries list of markets in the pricefeed
func GetCmdMarkets() *cobra.Command {
	return &cobra.Command{
		Use:   "markets",
		Short: "get the markets in the pricefeed",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types2.NewQueryClient(clientCtx)

			res, err := queryClient.Markets(context.Background(), &types2.QueryMarketsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
}

// GetCmdQueryParams queries the pricefeed module parameters
func GetCmdQueryParams() *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Short: "get the pricefeed module parameters",
		Long:  "Get the current global pricefeed module parameters.",
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
