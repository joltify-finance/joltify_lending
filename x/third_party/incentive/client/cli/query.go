package cli

import (
	"fmt"
	"strings"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"

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
	typeSavings     = "savings"
)

var rewardTypes = []string{typeDelegator, typeJolt, typeUSDXMinting, typeSwap}

// GetQueryCmd returns the cli query commands for the incentive module
func GetQueryCmd() *cobra.Command {
	incentiveQueryCmd := &cobra.Command{
		Use:   types2.ModuleName,
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
			$ %s query %s rewards
			$ %s query %s rewards --owner jolt16xjuwuy80gffg37ymkjfmafmf6k6e653cey7nn
			$ %s query %s rewards --type jolt
			$ %s query %s rewards --type usdx-minting
			$ %s query %s rewards --type delegator
			$ %s query %s rewards --type swap
			$ %s query %s rewards --type savings
			$ %s query %s rewards --type jolt --ownerjolt16xjuwuy80gffg37ymkjfmafmf6k6e653cey7nn 
			$ %s query %s rewards --type jolt --unsynced
			`,
				version.AppName, types2.ModuleName, version.AppName, types2.ModuleName,
				version.AppName, types2.ModuleName, version.AppName, types2.ModuleName,
				version.AppName, types2.ModuleName, version.AppName, types2.ModuleName,
				version.AppName, types2.ModuleName, version.AppName, types2.ModuleName,
				version.AppName, types2.ModuleName)),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			page, _ := cmd.Flags().GetInt(flags.FlagPage)
			limit, _ := cmd.Flags().GetInt(flags.FlagLimit)
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

			switch strings.ToLower(strType) {
			case typeJolt:
				params := types2.NewQueryRewardsParams(page, limit, owner, boolUnsynced)
				claims, err := executeJoltRewardsQuery(cliCtx, params)
				if err != nil {
					return err
				}
				return cliCtx.PrintObjectLegacy(claims)
			// case typeUSDXMinting:
			//	params := types.NewQueryRewardsParams(page, limit, owner, boolUnsynced)
			//	claims, err := executeUSDXMintingRewardsQuery(cliCtx, params)
			//	if err != nil {
			//		return err
			//	}
			//	return cliCtx.PrintObjectLegacy(claims)
			// case typeDelegator:
			//	params := types.NewQueryRewardsParams(page, limit, owner, boolUnsynced)
			//	claims, err := executeDelegatorRewardsQuery(cliCtx, params)
			//	if err != nil {
			//		return err
			//	}
			//	return cliCtx.PrintObjectLegacy(claims)
			// case typeSwap:
			//	params := types.NewQueryRewardsParams(page, limit, owner, boolUnsynced)
			//	claims, err := executeSwapRewardsQuery(cliCtx, params)
			//	if err != nil {
			//		return err
			//	}
			//	return cliCtx.PrintObjectLegacy(claims)
			// case typeSavings:
			//	params := types.NewQueryRewardsParams(page, limit, owner, boolUnsynced)
			//	claims, err := executeSavingsRewardsQuery(cliCtx, params)
			//	if err != nil {
			//		return err
			//	}
			//	return cliCtx.PrintObjectLegacy(claims)
			default:
				params := types2.NewQueryRewardsParams(page, limit, owner, boolUnsynced)

				joltClaims, err := executeJoltRewardsQuery(cliCtx, params)
				if err != nil {
					return err
				}
				//usdxMintingClaims, err := executeUSDXMintingRewardsQuery(cliCtx, params)
				//if err != nil {
				//	return err
				//}
				//delegatorClaims, err := executeDelegatorRewardsQuery(cliCtx, params)
				//if err != nil {
				//	return err
				//}
				//swapClaims, err := executeSwapRewardsQuery(cliCtx, params)
				//if err != nil {
				//	return err
				//}
				//savingsClaims, err := executeSavingsRewardsQuery(cliCtx, params)
				//if err != nil {
				//	return err
				//}
				if len(joltClaims) > 0 {
					if err := cliCtx.PrintObjectLegacy(joltClaims); err != nil {
						return err
					}
				}
				//if len(usdxMintingClaims) > 0 {
				//	if err := cliCtx.PrintObjectLegacy(usdxMintingClaims); err != nil {
				//		return err
				//	}
				//}
				//if len(delegatorClaims) > 0 {
				//	if err := cliCtx.PrintObjectLegacy(delegatorClaims); err != nil {
				//		return err
				//	}
				//}
				//if len(swapClaims) > 0 {
				//	if err := cliCtx.PrintObjectLegacy(swapClaims); err != nil {
				//		return err
				//	}
				//}
				//if len(savingsClaims) > 0 {
				//	if err := cliCtx.PrintObjectLegacy(savingsClaims); err != nil {
				//		return err
				//	}
				//}
			}
			return nil
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

			// Query
			route := fmt.Sprintf("custom/%s/%s", types2.ModuleName, types2.QueryGetParams)
			res, height, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				return err
			}
			cliCtx = cliCtx.WithHeight(height)

			// Decode and print results
			var params types2.Params
			if err := cliCtx.LegacyAmino.UnmarshalJSON(res, &params); err != nil {
				return fmt.Errorf("failed to unmarshal params: %w", err)
			}
			return cliCtx.PrintObjectLegacy(params)
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

			// Execute query
			route := fmt.Sprintf("custom/%s/%s", types2.ModuleName, types2.QueryGetRewardFactors)
			res, height, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				return err
			}
			cliCtx = cliCtx.WithHeight(height)

			// Decode and print results
			var response types2.QueryGetRewardFactorsResponse
			if err := cliCtx.LegacyAmino.UnmarshalJSON(res, &response); err != nil {
				return fmt.Errorf("failed to unmarshal reward factors: %w", err)
			}
			return cliCtx.PrintObjectLegacy(response)
		},
	}
	cmd.Flags().String(flagDenom, "", "(optional) filter reward factors by denom")
	return cmd
}

func executeJoltRewardsQuery(cliCtx client.Context, params types2.QueryRewardsParams) (types2.JoltLiquidityProviderClaims, error) {
	bz, err := cliCtx.LegacyAmino.MarshalJSON(params)
	if err != nil {
		return types2.JoltLiquidityProviderClaims{}, err
	}

	route := fmt.Sprintf("custom/%s/%s", types2.ModuleName, types2.QueryGetJoltRewards)
	res, height, err := cliCtx.QueryWithData(route, bz)
	if err != nil {
		return types2.JoltLiquidityProviderClaims{}, err
	}

	cliCtx = cliCtx.WithHeight(height)

	var claims types2.JoltLiquidityProviderClaims
	if err := cliCtx.LegacyAmino.UnmarshalJSON(res, &claims); err != nil {
		return types2.JoltLiquidityProviderClaims{}, fmt.Errorf("failed to unmarshal claims: %w", err)
	}

	return claims, nil
}
