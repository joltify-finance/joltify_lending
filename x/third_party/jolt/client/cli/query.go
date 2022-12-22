package cli

import (
	"context"
	"fmt"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
)

// flags for cli queries
const (
	flagName   = "name"
	flagDenom  = "denom"
	flagOwner  = "owner"
	flagBorrow = "borrower"
)

// GetQueryCmd returns the cli query commands for the  module
func GetQueryCmd() *cobra.Command {
	joltQueryCmd := &cobra.Command{
		Use:                        types2.ModuleName,
		Short:                      "Querying commands for the jolt module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmds := []*cobra.Command{
		queryParamsCmd(),
		queryAccountsCmd(),
		queryDepositsCmd(),
		queryUnsyncedDepositsCmd(),
		queryTotalDepositedCmd(),
		queryBorrowsCmd(),
		queryUnsyncedBorrowsCmd(),
		queryTotalBorrowedCmd(),
		queryInterestRateCmd(),
		queryReserves(),
		queryInterestFactorsCmd(),
		queryLiquidateCmd(),
	}

	for _, cmd := range cmds {
		flags.AddQueryFlagsToCmd(cmd)
	}

	joltQueryCmd.AddCommand(cmds...)

	return joltQueryCmd
}

func queryParamsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Short: "get the jolt module parameters",
		Long:  "Get the current global jolt module parameters.",
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

func queryAccountsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "accounts",
		Short: "query jolt module accounts",
		Long:  "Query for all jolt module accounts",
		Example: fmt.Sprintf(`%[1]s q %[2]s accounts
%[1]s q %[2]s accounts`, version.AppName, types2.ModuleName),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			req := &types2.QueryAccountsRequest{}

			queryClient := types2.NewQueryClient(clientCtx)

			res, err := queryClient.Accounts(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}

func queryUnsyncedDepositsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unsynced-deposits",
		Short: "query jolt module unsynced deposits with optional filters",
		Long:  "query for all jolt module unsynced deposits or a specific unsynced deposit using flags",
		Example: fmt.Sprintf(`%[1]s q %[2]s unsynced-deposits
%[1]s q %[2]s unsynced-deposits --owner jolt1pm9kvrl64fqgxvym7f7m42dndjk52v9mqnzdnn --denom bnb
%[1]s q %[2]s unsynced-deposits --denom ujolt
%[1]s q %[2]s unsynced-deposits --denom btcb`, version.AppName, types2.ModuleName),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			ownerBech, err := cmd.Flags().GetString(flagOwner)
			if err != nil {
				return err
			}
			denom, err := cmd.Flags().GetString(flagDenom)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			req := &types2.QueryUnsyncedDepositsRequest{
				Denom:      denom,
				Pagination: pageReq,
			}

			if len(ownerBech) != 0 {
				depositOwner, err := sdk.AccAddressFromBech32(ownerBech)
				if err != nil {
					return err
				}
				req.Owner = depositOwner.String()
			}

			queryClient := types2.NewQueryClient(clientCtx)

			res, err := queryClient.UnsyncedDeposits(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "unsynced-deposits")

	cmd.Flags().String(flagOwner, "", "(optional) filter for unsynced deposits by owner address")
	cmd.Flags().String(flagDenom, "", "(optional) filter for unsynced deposits by denom")

	return cmd
}

func queryDepositsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposits",
		Short: "query jolt module deposits with optional filters",
		Long:  "query for all jolt module deposits or a specific deposit using flags",
		Example: fmt.Sprintf(`%[1]s q %[2]s deposits
%[1]s q %[2]s deposits --owner jolt1l0xsq2z7gqd7yly0g40y5836g0appumark77ny --denom bnb
%[1]s q %[2]s deposits --denom ujolt
%[1]s q %[2]s deposits --denom btcb`, version.AppName, types2.ModuleName),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			ownerBech, err := cmd.Flags().GetString(flagOwner)
			if err != nil {
				return err
			}
			denom, err := cmd.Flags().GetString(flagDenom)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			req := &types2.QueryDepositsRequest{
				Denom:      denom,
				Pagination: pageReq,
			}

			if len(ownerBech) != 0 {
				depositOwner, err := sdk.AccAddressFromBech32(ownerBech)
				if err != nil {
					return err
				}
				req.Owner = depositOwner.String()
			}

			queryClient := types2.NewQueryClient(clientCtx)

			res, err := queryClient.Deposits(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "deposits")

	cmd.Flags().String(flagOwner, "", "(optional) filter for deposits by owner address")
	cmd.Flags().String(flagDenom, "", "(optional) filter for deposits by denom")

	return cmd
}

func queryUnsyncedBorrowsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unsynced-borrows",
		Short: "query jolt module unsynced borrows with optional filters",
		Long:  "query for all jolt module unsynced borrows or a specific unsynced borrow using flags",
		Example: fmt.Sprintf(`%[1]s q %[2]s unsynced-borrows
%[1]s q %[2]s unsynced-borrows --owner jolt1pm9kvrl64fqgxvym7f7m42dndjk52v9mqnzdnn
%[1]s q %[2]s unsynced-borrows --denom bnb`, version.AppName, types2.ModuleName),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			ownerBech, err := cmd.Flags().GetString(flagOwner)
			if err != nil {
				return err
			}
			denom, err := cmd.Flags().GetString(flagDenom)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			req := &types2.QueryUnsyncedBorrowsRequest{
				Denom:      denom,
				Pagination: pageReq,
			}

			if len(ownerBech) != 0 {
				borrowOwner, err := sdk.AccAddressFromBech32(ownerBech)
				if err != nil {
					return err
				}
				req.Owner = borrowOwner.String()
			}

			queryClient := types2.NewQueryClient(clientCtx)

			res, err := queryClient.UnsyncedBorrows(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "unsynced borrows")

	cmd.Flags().String(flagOwner, "", "(optional) filter for unsynced borrows by owner address")
	cmd.Flags().String(flagDenom, "", "(optional) filter for unsynced borrows by denom")

	return cmd
}

func queryBorrowsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "borrows",
		Short: "query jolt module borrows with optional filters",
		Long:  "query for all jolt module borrows or a specific borrow using flags",
		Example: fmt.Sprintf(`%[1]s q %[2]s borrows
%[1]s q %[2]s borrows --owner jolt1pm9kvrl64fqgxvym7f7m42dndjk52v9mqnzdnn 
%[1]s q %[2]s borrows --denom bnb`, version.AppName, types2.ModuleName),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			ownerBech, err := cmd.Flags().GetString(flagOwner)
			if err != nil {
				return err
			}
			denom, err := cmd.Flags().GetString(flagDenom)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			req := &types2.QueryBorrowsRequest{
				Denom:      denom,
				Pagination: pageReq,
			}

			if len(ownerBech) != 0 {
				borrowOwner, err := sdk.AccAddressFromBech32(ownerBech)
				if err != nil {
					return err
				}
				req.Owner = borrowOwner.String()
			}

			queryClient := types2.NewQueryClient(clientCtx)
			res, err := queryClient.Borrows(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "borrows")

	cmd.Flags().String(flagOwner, "", "(optional) filter for borrows by owner address")
	cmd.Flags().String(flagDenom, "", "(optional) filter for borrows by denom")

	return cmd
}

func queryTotalBorrowedCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "total-borrowed",
		Short: "get total current borrowed amount",
		Long:  "get the total amount of coins currently borrowed using flags",
		Example: fmt.Sprintf(`%[1]s q %[2]s total-borrowed
%[1]s q %[2]s total-borrowed --denom bnb`, version.AppName, types2.ModuleName),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			denom, err := cmd.Flags().GetString(flagDenom)
			if err != nil {
				return err
			}

			queryClient := types2.NewQueryClient(clientCtx)
			res, err := queryClient.TotalBorrowed(context.Background(), &types2.QueryTotalBorrowedRequest{
				Denom: denom,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().String(flagDenom, "", "(optional) filter total borrowed coins by denom")

	return cmd
}

func queryTotalDepositedCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "total-deposited",
		Short: "get total current deposited amount",
		Long:  "get the total amount of coins currently deposited using flags",
		Example: fmt.Sprintf(`%[1]s q %[2]s total-deposited
%[1]s q %[2]s total-deposited --denom bnb`, version.AppName, types2.ModuleName),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			denom, err := cmd.Flags().GetString(flagDenom)
			if err != nil {
				return err
			}

			queryClient := types2.NewQueryClient(clientCtx)
			res, err := queryClient.TotalDeposited(context.Background(), &types2.QueryTotalDepositedRequest{
				Denom: denom,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().String(flagDenom, "", "(optional) filter total deposited coins by denom")

	return cmd
}

func queryInterestRateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "interest-rate",
		Short: "get current money market interest rates",
		Long:  "get current money market interest rates",
		Example: fmt.Sprintf(`%[1]s q %[2]s interest-rate
%[1]s q %[2]s interest-rate --denom bnb`, version.AppName, types2.ModuleName),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			denom, err := cmd.Flags().GetString(flagDenom)
			if err != nil {
				return err
			}

			queryClient := types2.NewQueryClient(clientCtx)
			res, err := queryClient.InterestRate(context.Background(), &types2.QueryInterestRateRequest{
				Denom: denom,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().String(flagDenom, "", "(optional) filter interest rates by denom")

	return cmd
}

func queryReserves() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reserves",
		Short: "get total current Jolt module reserves",
		Long:  "get the total amount of coins currently held as reserve by the Jolt module",
		Example: fmt.Sprintf(`%[1]s q %[2]s reserves
%[1]s q %[2]s reserves --denom bnb`, version.AppName, types2.ModuleName),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			denom, err := cmd.Flags().GetString(flagDenom)
			if err != nil {
				return err
			}

			queryClient := types2.NewQueryClient(clientCtx)
			res, err := queryClient.Reserves(context.Background(), &types2.QueryReservesRequest{
				Denom: denom,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().String(flagDenom, "", "(optional) filter reserve coins by denom")

	return cmd
}

func queryInterestFactorsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "interest-factors",
		Short: "get current global interest factors",
		Long:  "get current global interest factors",
		Example: fmt.Sprintf(`%[1]s q %[2]s interest-factors
%[1]s q %[2]s interest-factors --denom bnb`, version.AppName, types2.ModuleName),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			denom, err := cmd.Flags().GetString(flagDenom)
			if err != nil {
				return err
			}

			queryClient := types2.NewQueryClient(clientCtx)
			res, err := queryClient.InterestFactors(context.Background(), &types2.QueryInterestFactorsRequest{
				Denom: denom,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().String(flagDenom, "", "(optional) filter interest factors by denom")

	return cmd
}

func queryLiquidateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "liquidate",
		Short: "query jolt module borrowers that close to liquidation with optional filters",
		Long:  "query for all jolt module or a specific borrow using flags that close to liquidation",
		Example: fmt.Sprintf(`%[1]s q %[2]s liquidate 
%[1]s q %[2]s liquidate --owner jolt1l0xsq2z7gqd7yly0g40y5836g0appumark77ny`, version.AppName, types2.ModuleName),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			borrower, err := cmd.Flags().GetString(flagBorrow)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			req := &types2.QueryLiquidateRequest{
				Borrower:   borrower,
				Pagination: pageReq,
			}

			queryClient := types2.NewQueryClient(clientCtx)
			res, err := queryClient.Liquidate(context.Background(), req)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "liquidate")

	cmd.Flags().String(flagBorrow, "", "(optional) filter for qulidtions by borrower address")

	return cmd
}
