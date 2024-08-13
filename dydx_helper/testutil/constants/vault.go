package constants

import (
	"github.com/joltify-finance/joltify_lending/dydx_helper/dtypes"
	"github.com/joltify-finance/joltify_lending/dydx_helper/x/vault/types"
)

var (
	Vault_Clob_0 = types.VaultId{
		Type:   types.VaultType_VAULT_TYPE_CLOB,
		Number: 0,
	}
	Vault_Clob_1 = types.VaultId{
		Type:   types.VaultType_VAULT_TYPE_CLOB,
		Number: 1,
	}

	MsgDepositToVault_Clob0_Alice0_100 = &types.MsgDepositToVault{
		VaultId:       &Vault_Clob_0,
		SubaccountId:  &Alice_Num0,
		QuoteQuantums: dtypes.NewInt(100),
	}
)
