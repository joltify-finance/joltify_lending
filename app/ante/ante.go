package ante

import (
	"fmt"
	"runtime/debug"

	"cosmossdk.io/log"
	accountkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	txsigning "cosmossdk.io/x/tx/signing"

	consensusparamkeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	ibcante "github.com/cosmos/ibc-go/v8/modules/core/ante"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"
	spvkeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"
)

// cosmosHandlerOptions extends HandlerOptions to provide some Cosmos specific configurations
type cosmosHandlerOptions struct {
	HandlerOptions
	isEIP712 bool
}

// HandlerOptions extend the SDK's AnteHandler options by requiring the IBC
// channel keeper, EVM Keeper and Fee Market Keeper.
type HandlerOptions struct {
	AccountKeeper          *accountkeeper.AccountKeeper
	BankKeeper             bankkeeper.Keeper
	IBCKeeper              *ibckeeper.Keeper
	SpvKeeper              spvkeeper.Keeper
	FeegrantKeeper         authante.FeegrantKeeper
	SignModeHandler        *txsigning.HandlerMap
	SigGasConsumer         authante.SignatureVerificationGasConsumer
	MaxTxGasWanted         uint64
	AddressFetchers        []AddressFetcher
	ExtensionOptionChecker authante.ExtensionOptionChecker
	TxFeeChecker           authante.TxFeeChecker
}

func normalCosmosTxAnte(ctx sdk.Context, tx sdk.Tx, sim bool, options HandlerOptions) (sdk.Context, error) {
	var anteHandler sdk.AnteHandler
	switch tx.(type) {
	case sdk.Tx:
		anteHandler = newCosmosAnteHandler(
			cosmosHandlerOptions{
				HandlerOptions: options,
				isEIP712:       false,
			})
	default:
		return ctx, errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "invalid transaction type: %T", tx)
	}
	return anteHandler(ctx, tx, sim)
}

// NewAnteHandler returns an 'AnteHandler' that will run actions before a tx is sent to a module's handler.
func NewAnteHandler(options HandlerOptions, consensusKeeper consensusparamkeeper.Keeper) (sdk.AnteHandler, error) {
	if options.AccountKeeper == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "account keeper is required for ante builder")
	}

	if options.BankKeeper == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "bank keeper is required for ante builder")
	}

	if options.SignModeHandler == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "sign mode handler is required for ante builder")
	}

	if options.SigGasConsumer == nil {
		options.SigGasConsumer = authante.DefaultSigVerificationGasConsumer
	}

	return func(
		ctx sdk.Context, tx sdk.Tx, sim bool,
	) (newCtx sdk.Context, err error) {

		defer Recover(ctx.Logger(), &err)

		ctxR := sdk.UnwrapSDKContext(ctx)
		p, err := consensusKeeper.Params(ctxR, nil)
		if err != nil {
			panic("fail to load the consensus")
		}
		ctx = ctx.WithConsensusParams(*p.GetParams())

		txWithExtensions, ok := tx.(authante.HasExtensionOptionsTx)
		if ok {
			opts := txWithExtensions.GetExtensionOptions()
			if len(opts) > 1 {
				return ctx, errorsmod.Wrap(
					sdkerrors.ErrInvalidRequest,
					"rejecting tx with more than 1 extension option",
				)
			}

			if len(opts) == 1 {
				typeURL := opts[0].GetTypeUrl()
				switch typeURL {

				case "/cosmos.authz.v1beta1.MsgRevoke", "/cosmos.authz.v1beta1.MsgExec", "/cosmos.authz.v1beta1.MsgGrant":
					return normalCosmosTxAnte(ctx, tx, sim, options)
				default:
					return ctx, errorsmod.Wrapf(
						sdkerrors.ErrUnknownExtensionOptions,
						"rejecting tx with unsupported extension option: %s", typeURL,
					)
				}

			}
		}

		return normalCosmosTxAnte(ctx, tx, sim, options)
	}, nil
}

func newCosmosAnteHandler(options cosmosHandlerOptions) sdk.AnteHandler {
	var decorators []sdk.AnteDecorator

	decorators = append(decorators,
		authante.NewSetUpContextDecorator(), // second decorator. SetUpContext must be called before other decorators
	)

	if !options.isEIP712 {
		decorators = append(decorators, authante.NewExtensionOptionsDecorator(options.ExtensionOptionChecker))
	}

	if len(options.AddressFetchers) > 0 {
		decorators = append(decorators, NewAuthenticatedMempoolDecorator(options.AddressFetchers...))
	}

	var sigVerification sdk.AnteDecorator = authante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler)
	//if options.isEIP712 {
	// fixme ignore the deprecated warning
	//	sigVerification = evmante.NewLegacyEip712SigVerificationDecorator(options.AccountKeeper, options.SignModeHandler) //nolint
	//}

	decorators = append(decorators,
		NewAuthzLimiterDecorator(
		//sdk.MsgTypeURL(&evmtypes.MsgEthereumTx{}),
		// sdk.MsgTypeURL(&vesting.MsgCreateVestingAccount{}),
		// sdk.MsgTypeURL(&vesting.MsgCreatePermanentLockedAccount{}),
		// sdk.MsgTypeURL(&vesting.MsgCreatePeriodicVestingAccount{}),
		),
		authante.NewValidateBasicDecorator(),
		authante.NewTxTimeoutHeightDecorator(),
		// If ethermint x/feemarket is enabled, align Cosmos min fee with the EVM
		// evmante.NewMinGasPriceDecorator(options.FeeMarketKeeper, options.EvmKeeper),
		authante.NewValidateMemoDecorator(options.AccountKeeper),
		authante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		authante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper, options.TxFeeChecker),
		authante.NewSetPubKeyDecorator(options.AccountKeeper), // SetPubKeyDecorator must be called before all signature verification decorators
		authante.NewValidateSigCountDecorator(options.AccountKeeper),
		authante.NewSigGasConsumeDecorator(options.AccountKeeper, options.SigGasConsumer),
		sigVerification,
		authante.NewIncrementSequenceDecorator(options.AccountKeeper), // innermost AnteDecorator
		ibcante.NewRedundantRelayDecorator(options.IBCKeeper),
		NewSPVNFTDecorator(options.SpvKeeper),
	)
	return sdk.ChainAnteDecorators(decorators...)
}

func Recover(logger log.Logger, err *error) {
	if r := recover(); r != nil {
		*err = errorsmod.Wrapf(sdkerrors.ErrPanic, "%v", r)

		if e, ok := r.(error); ok {
			logger.Error(
				"ante handler panicked",
				"error", e,
				"stack trace", string(debug.Stack()),
			)
		} else {
			logger.Error(
				"ante handler panicked",
				"recover", fmt.Sprintf("%v", r),
			)
		}
	}
}
