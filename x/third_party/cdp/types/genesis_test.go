package types_test

import (
	"testing"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/cdp/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenesis_Default(t *testing.T) {
	defaultGenesis := types2.DefaultGenesisState()

	require.NoError(t, defaultGenesis.Validate())

	defaultParams := types2.DefaultParams()
	assert.Equal(t, defaultParams, defaultGenesis.Params)
}

func TestGenesisTotalPrincipal(t *testing.T) {
	tests := []struct {
		giveName           string
		giveCollateralType string
		givePrincipal      sdk.Int
		wantIsError        bool
		wantError          string
	}{
		{"valid", "usdx", sdk.NewInt(10), false, ""},
		{"zero principal", "usdx", sdk.NewInt(0), false, ""},
		{"invalid empty collateral type", "", sdk.NewInt(10), true, "collateral type cannot be empty"},
		{"invalid negative principal", "usdx", sdk.NewInt(-10), true, "total principal should be positive"},
		{"both invalid", "", sdk.NewInt(-10), true, "collateral type cannot be empty"},
	}

	for _, tt := range tests {
		t.Run(tt.giveName, func(t *testing.T) {
			tp := types2.NewGenesisTotalPrincipal(tt.giveCollateralType, tt.givePrincipal)

			err := tp.Validate()
			if tt.wantIsError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
