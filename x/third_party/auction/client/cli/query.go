package cli

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/auction/types"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
)

// GetQueryCmd returns the cli query commands for the auction module
func GetQueryCmd() *cobra.Command {
	auctionQueryCmd := &cobra.Command{
		Use:   types2.ModuleName,
		Short: fmt.Sprintf("Querying commands for the %s module", types2.ModuleName),
	}

	cmds := []*cobra.Command{
		GetCmdQueryParams(),
		GetCmdQueryAuction(),
		GetCmdQueryAuctions(),
	}

	for _, cmd := range cmds {
		flags.AddQueryFlagsToCmd(cmd)
	}

	auctionQueryCmd.AddCommand(cmds...)

	return auctionQueryCmd
}

// GetCmdQueryParams queries the issuance module parameters
func GetCmdQueryParams() *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Short: fmt.Sprintf("get the %s module parameters", types2.ModuleName),
		Long:  "Get the current auction module parameters.",
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

// GetCmdQueryAuction queries one auction in the store
func GetCmdQueryAuction() *cobra.Command {
	return &cobra.Command{
		Use:   "auction [auction-id]",
		Short: "get info about an auction",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types2.NewQueryClient(clientCtx)

			auctionID, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			params := types2.QueryAuctionRequest{
				AuctionId: uint64(auctionID),
			}

			res, err := queryClient.Auction(context.Background(), &params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
}

// Query auction flags
const (
	flagType  = "type"
	flagDenom = "denom"
	flagPhase = "phase"
	flagOwner = "owner"
)

// GetCmdQueryAuctions queries the auctions in the store
func GetCmdQueryAuctions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auctions",
		Short: "query auctions with optional filters",
		Long:  "Query for all paginated auctions that match optional filters.",
		Example: strings.Join([]string{
			fmt.Sprintf("  $ %s q %s auctions --type=(collateral|surplus|debt)", version.AppName, types2.ModuleName),
			fmt.Sprintf("  $ %s q %s auctions --owner=jolt1pm9kvrl64fqgxvym7f7m42dndjk52v9mqnzdnn", version.AppName, types2.ModuleName),
			fmt.Sprintf("  $ %s q %s auctions --denom=bnb", version.AppName, types2.ModuleName),
			fmt.Sprintf("  $ %s q %s auctions --phase=(forward|reverse)", version.AppName, types2.ModuleName),
			fmt.Sprintf("  $ %s q %s auctions --page=2 --limit=100", version.AppName, types2.ModuleName),
		}, "\n"),
		RunE: func(cmd *cobra.Command, args []string) error {
			auctionType, err := cmd.Flags().GetString(flagType)
			if err != nil {
				return err
			}
			owner, err := cmd.Flags().GetString(flagOwner)
			if err != nil {
				return err
			}
			denom, err := cmd.Flags().GetString(flagDenom)
			if err != nil {
				return err
			}
			phase, err := cmd.Flags().GetString(flagPhase)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			if len(auctionType) != 0 {
				auctionType = strings.ToLower(auctionType)

				if auctionType != types2.CollateralAuctionType &&
					auctionType != types2.SurplusAuctionType &&
					auctionType != types2.DebtAuctionType {
					return fmt.Errorf("invalid auction type %s", auctionType)
				}
			}

			if len(owner) != 0 {
				if auctionType != types2.CollateralAuctionType {
					return fmt.Errorf("cannot apply owner flag to non-collateral auction type")
				}
				_, err := sdk.AccAddressFromBech32(owner)
				if err != nil {
					return fmt.Errorf("cannot parse address from auction owner %s", owner)
				}
			}

			if len(denom) != 0 {
				err := sdk.ValidateDenom(denom)
				if err != nil {
					return err
				}
			}

			if len(phase) != 0 {
				phase = strings.ToLower(phase)

				if len(auctionType) > 0 && auctionType != types2.CollateralAuctionType {
					return fmt.Errorf("cannot apply phase flag to non-collateral auction type")
				}
				if phase != types2.ForwardAuctionPhase && phase != types2.ReverseAuctionPhase {
					return fmt.Errorf("invalid auction phase %s", phase)
				}
			}

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types2.NewQueryClient(clientCtx)
			request := types2.QueryAuctionsRequest{
				Type:       auctionType,
				Owner:      owner,
				Denom:      denom,
				Phase:      phase,
				Pagination: pageReq,
			}

			res, err := queryClient.Auctions(context.Background(), &request)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "auctions")

	cmd.Flags().String(flagType, "", "(optional) filter by auction type, type: collateral, debt, surplus")
	cmd.Flags().String(flagOwner, "", "(optional) filter by collateral auction owner")
	cmd.Flags().String(flagDenom, "", "(optional) filter by auction denom")
	cmd.Flags().String(flagPhase, "", "(optional) filter by collateral auction phase, phase: forward/reverse")

	return cmd
}
