package constants

import (
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/app"
	satypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/subaccounts/types"
)

func init() {
	// This package does not contain the `app/config` package in its import chain, and therefore needs to call
	// SetAddressPrefixes() explicitly in order to set the `dydx` address prefixes.
	app.SetSDKConfig()
}

var (
	Alice_Num0 = satypes.SubaccountId{
		Owner:  types.MustBech32ifyAddressBytes(app.Bech32MainPrefix, AliceAccAddress),
		Number: 0,
	}
	Alice_Num1 = satypes.SubaccountId{
		Owner:  types.MustBech32ifyAddressBytes(app.Bech32MainPrefix, AliceAccAddress),
		Number: 1,
	}
	Bob_Num0 = satypes.SubaccountId{
		Owner:  types.MustBech32ifyAddressBytes(app.Bech32MainPrefix, BobAccAddress),
		Number: 0,
	}
	Bob_Num1 = satypes.SubaccountId{
		Owner:  types.MustBech32ifyAddressBytes(app.Bech32MainPrefix, BobAccAddress),
		Number: 1,
	}
	Bob_Num2 = satypes.SubaccountId{
		Owner:  types.MustBech32ifyAddressBytes(app.Bech32MainPrefix, BobAccAddress),
		Number: 2,
	}
	Carl_Num0 = satypes.SubaccountId{
		Owner:  types.MustBech32ifyAddressBytes(app.Bech32MainPrefix, CarlAccAddress),
		Number: 0,
	}
	Carl_Num1 = satypes.SubaccountId{
		Owner:  types.MustBech32ifyAddressBytes(app.Bech32MainPrefix, CarlAccAddress),
		Number: 1,
	}
	Dave_Num0 = satypes.SubaccountId{
		Owner:  types.MustBech32ifyAddressBytes(app.Bech32MainPrefix, DaveAccAddress),
		Number: 0,
	}
	Dave_Num1 = satypes.SubaccountId{
		Owner:  types.MustBech32ifyAddressBytes(app.Bech32MainPrefix, DaveAccAddress),
		Number: 1,
	}
)
