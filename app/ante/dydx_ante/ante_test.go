package dydxante_test

import (
	"testing"

	"cosmossdk.io/store/rootmulti"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	dydxante "github.com/joltify-finance/joltify_lending/app/ante/dydx_ante"
	appconfig "github.com/joltify-finance/joltify_lending/app/config"

	testApp "github.com/joltify-finance/joltify_lending/testutil/dydx/testutil/app"
	"github.com/stretchr/testify/require"
)

func newHandlerOptions() dydxante.HandlerOptions {
	encodingConfig := appconfig.MakeEncodingConfig()
	dydxApp := testApp.DefaultTestApp(nil)
	return dydxante.HandlerOptions{
		HandlerOptions: ante.HandlerOptions{
			AccountKeeper:   dydxApp.AccountKeeper,
			BankKeeper:      dydxApp.BankKeeper,
			SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
			FeegrantKeeper:  dydxApp.FeeGrantKeeper,
			SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
		},
		ClobKeeper:   dydxApp.ClobKeeper,
		Codec:        encodingConfig.Codec,
		AuthStoreKey: dydxApp.CommitMultiStore().(*rootmulti.Store).StoreKeysByName()[authtypes.StoreKey],
	}
}

func TestNewAnteHandler(t *testing.T) {
	handlerOptions := newHandlerOptions()
	anteHandler, err := dydxante.NewAnteHandler(handlerOptions)
	require.NoError(t, err, "NewAnteHandler call failed")
	require.NotNil(t, anteHandler, "expected non-nil AnteHandler function")
}

func TestNewAnteHandler_Error(t *testing.T) {
	tests := map[string]struct {
		handlerMutation func(*dydxante.HandlerOptions)
		errorMsg        string
	}{
		"nil handlerOptions.AccountKeeper": {
			handlerMutation: func(options *dydxante.HandlerOptions) { options.AccountKeeper = nil },
			errorMsg:        "account keeper is required for ante builder",
		},
		"nil handlerOptions.BankKeeper": {
			handlerMutation: func(options *dydxante.HandlerOptions) { options.BankKeeper = nil },
			errorMsg:        "bank keeper is required for ante builder",
		},
		"nil handlerOptions.SignModeHandler": {
			handlerMutation: func(options *dydxante.HandlerOptions) { options.SignModeHandler = nil },
			errorMsg:        "sign mode handler is required for ante builder",
		},
		"nil ClobKeeper": {
			handlerMutation: func(options *dydxante.HandlerOptions) { options.ClobKeeper = nil },
			errorMsg:        "clob keeper is required for ante builder",
		},
		"nil Codec": {
			handlerMutation: func(options *dydxante.HandlerOptions) { options.Codec = nil },
			errorMsg:        "codec is required for ante builder",
		},
		"nil AuthStoreKey": {
			handlerMutation: func(options *dydxante.HandlerOptions) { options.AuthStoreKey = nil },
			errorMsg:        "auth store key is required for ante builder",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			handlerOptions := newHandlerOptions()
			tc.handlerMutation(&handlerOptions)

			anteHandler, err := dydxante.NewAnteHandler(handlerOptions)
			require.Nil(t, anteHandler, "Expected Ante Handler creation to error")
			require.Errorf(t, err, tc.errorMsg)
		})
	}
}
