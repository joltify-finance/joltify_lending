package cli

import (
	"fmt"
	"strings"

	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
)

const (
	flagOwner    = "owner"
	flagType     = "type"
	flagUnsynced = "unsynced"
	flagDenom    = "denom"

	typeDelegator   = "delegator"
	typeJolt        = "jolt"
	typeUSDXMinting = "usdx-minting"
	typeSwap        = "swap"
)

var rewardTypes = []string{typeDelegator, typeJolt, typeUSDXMinting, typeSwap}

// GetQueryCmd returns the cli query commands for the incentive module
func GetQueryCmd() *cobra.Command {
	incentiveQueryCmd := &cobra.Command{
		Use:   types.ModuleName,
		Short: "Querying commands for the incentive module",
	}

	cmds := []*cobra.Command{
		queryParamsCmd(),
		queryRewardsCmd(),
		queryRewardFactorsCmd(),
	}

	for _, cmd := range cmds {
		flags.AddQueryFlagsToCmd(cmd)
	}

	incentiveQueryCmd.AddCommand(cmds...)

	return incentiveQueryCmd
}

func queryRewardsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rewards",
		Short: "query claimable rewards",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query rewards with optional flags for owner and type
			Example:
			$ %s query %s rewards --type jolt --owner jolt16xjuwuy80gffg37ymkjfmafmf6k6e653cey7nn --unsynced
			`,
				version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			strOwner, _ := cmd.Flags().GetString(flagOwner)
			strType, _ := cmd.Flags().GetString(flagType)
			boolUnsynced, _ := cmd.Flags().GetBool(flagUnsynced)

			// Prepare params for querier
			var owner sdk.AccAddress
			if strOwner != "" {
				if owner, err = sdk.AccAddressFromBech32(strOwner); err != nil {
					return err
				}
			}

			queryClient := types.NewQueryClient(cliCtx)
			resp, err := queryClient.Rewards(cmd.Context(), &types.QueryRewardsRequest{
				Owner:          owner.String(),
				Unsynchronized: boolUnsynced,
				RewardType:     strType,
			})
			if err != nil {
				return err
			}
			return cliCtx.PrintProto(resp)
		},
	}
	cmd.Flags().String(flagOwner, "", "(optional) filter by owner address")
	cmd.Flags().String(flagType, "", fmt.Sprintf("(optional) filter by a reward type: %s", strings.Join(rewardTypes, "|")))
	cmd.Flags().Bool(flagUnsynced, false, "(optional) get unsynced claims")
	cmd.Flags().Int(flags.FlagPage, 1, "pagination page rewards of to to query for")
	cmd.Flags().Int(flags.FlagLimit, 100, "pagination limit of rewards to query for")
	return cmd
}

func queryParamsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Short: "get the incentive module parameters",
		Long:  "Get the current global incentive module parameters.",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			req := types.QueryParamsRequest{}
			queryClient := types.NewQueryClient(cliCtx)
			res, err := queryClient.Params(cmd.Context(), &req)
			if err != nil {
				return err
			}

			return cliCtx.PrintProto(res)
		},
	}
}

func queryRewardFactorsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reward-factors",
		Short: "get current global reward factors",
		Long:  `Get current global reward factors for all reward types.`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			req := types.QueryRewardFactorsRequest{}
			queryClient := types.NewQueryClient(cliCtx)
			res, err := queryClient.RewardFactors(cmd.Context(), &req)
			if err != nil {
				return err
			}

			return cliCtx.PrintProto(res)
		},
	}
	cmd.Flags().String(flagDenom, "", "(optional) filter reward factors by denom")
	return cmd
}
