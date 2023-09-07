package cli

import (
	"fmt"
	"strings"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
)

const (
	multiplierFlag      = "multiplier"
	multiplierFlagShort = "m"
)

// GetTxCmd returns the transaction cli commands for the incentive module
func GetTxCmd() *cobra.Command {
	incentiveTxCmd := &cobra.Command{
		Use:   types2.ModuleName,
		Short: "transaction commands for the incentive module",
	}

	cmds := []*cobra.Command{
		// getCmdClaimCdp(),
		getCmdClaimJolt(),
		// getCmdClaimDelegator(),
		// getCmdClaimSwap(),
		// getCmdClaimSavings(),
	}

	for _, cmd := range cmds {
		flags.AddTxFlagsToCmd(cmd)
	}

	incentiveTxCmd.AddCommand(cmds...)

	return incentiveTxCmd
}

func getCmdClaimJolt() *cobra.Command {
	var denomsToClaim map[string]string

	cmd := &cobra.Command{
		Use:   "claim-jolt",
		Short: "claim sender's Jolt module rewards using given multipliers",
		Long:  `Claim sender's outstanding Jolt rewards for deposit/borrow using given multipliers`,
		Example: strings.Join([]string{
			fmt.Sprintf(`  $ %s tx %s claim-jolt --%s jolt=large --%s ujolt=small`, version.AppName, types2.ModuleName, multiplierFlag, multiplierFlag),
			fmt.Sprintf(`  $ %s tx %s claim-jolt --%s jolt=large,ujolt=small`, version.AppName, types2.ModuleName, multiplierFlag),
		}, "\n"),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := cliCtx.GetFromAddress()
			selections := types2.NewSelectionsFromMap(denomsToClaim)

			msg := types2.NewMsgClaimJoltReward(sender.String(), selections)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), &msg)
		},
	}
	cmd.Flags().StringToStringVarP(&denomsToClaim, multiplierFlag, multiplierFlagShort, nil, "specify the denoms to claim, each with a multiplier lockup")
	if err := cmd.MarkFlagRequired(multiplierFlag); err != nil {
		panic(err)
	}
	return cmd
}
