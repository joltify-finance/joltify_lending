package types

import (
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/types"
)

type ExecutionData struct {
	Origin string      `json:"origin"`
	Name   string      `json:"name"`
	Args   interface{} `json:"args"`
}

type Slippage struct {
	MaxPenalty   *types.Dec `json:"max_penalty,omitempty"`
	MinIncentive *types.Dec `json:"min_incentive,omitempty"`
}

type VaultSubscribe struct {
	Slippage *Slippage `json:"slippage,omitempty"`
}

type BasicVaultRedeemArgs struct {
	LpTokenBurnAmount math.Int `json:"lp_token_burn_amount"`
	Slippage          Slippage `json:"slippage,omitempty"`
}

type VaultRedeem struct {
	BasicVaultRedeemArgs
	RedemptionType string `json:"redemption_type,omitempty"`
}

type VaultSubscribeRedeem struct {
	Subscribe *VaultSubscribe `json:"subscribe,omitempty"`
	Redeem    interface{}     `json:"redeem,omitempty"`
}

type VaultInput struct {
	VaultSubaccountId  string               `json:"vault_subaccount_id"`
	TraderSubaccountId string               `json:"trader_subaccount_id"`
	Msg                VaultSubscribeRedeem `json:"msg"`
}
