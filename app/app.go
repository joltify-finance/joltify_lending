package app

import (
	"cosmossdk.io/log"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/runtime"

	addresscodec "github.com/cosmos/cosmos-sdk/codec/address"

	burnauctionmodule "github.com/joltify-finance/joltify_lending/x/burnauction"
	burnauctionmoduletypes "github.com/joltify-finance/joltify_lending/x/burnauction/types"

	burnauctionmodulekeeper "github.com/joltify-finance/joltify_lending/x/burnauction/keeper"

	v1 "github.com/joltify-finance/joltify_lending/upgrade"

	quotamodule "github.com/joltify-finance/joltify_lending/x/quota"
	quotamodulekeeper "github.com/joltify-finance/joltify_lending/x/quota/keeper"
	quotamoduletypes "github.com/joltify-finance/joltify_lending/x/quota/types"

	ibcratelimit "github.com/joltify-finance/joltify_lending/x/ibc-rate-limit"
	ibcratelimittypes "github.com/joltify-finance/joltify_lending/x/ibc-rate-limit/types"

	ibctm "github.com/cosmos/ibc-go/v8/modules/light-clients/07-tendermint"

	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	reflectionv1 "cosmossdk.io/api/cosmos/reflection/v1"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/spf13/cast"

	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"

	"github.com/cosmos/cosmos-sdk/x/consensus"

	tmos "github.com/cometbft/cometbft/libs/os"
	"github.com/joltify-finance/joltify_lending/x/third_party/swap"

	runtimeservices "github.com/cosmos/cosmos-sdk/runtime/services"

	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	evmante "github.com/evmos/ethermint/app/ante"
	"github.com/evmos/ethermint/x/evm"
	"github.com/gorilla/mux"
	_ "github.com/joltify-finance/joltify_lending/client/docs/statik"

	nodeservice "github.com/cosmos/cosmos-sdk/client/grpc/node"
	"github.com/rakyll/statik/fs"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/server/config"
	ibcporttypes "github.com/cosmos/ibc-go/v8/modules/core/05-port/types"
	"github.com/evmos/ethermint/x/feemarket"
	"github.com/joltify-finance/joltify_lending/x/third_party/auction"
	auctionkeeper "github.com/joltify-finance/joltify_lending/x/third_party/auction/keeper"
	auctiontypes "github.com/joltify-finance/joltify_lending/x/third_party/auction/types"
	"github.com/joltify-finance/joltify_lending/x/third_party/incentive"

	nftmoduletypes "cosmossdk.io/x/nft"
	nftmodulekeeper "cosmossdk.io/x/nft/keeper"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
	kycmodulekeeper "github.com/joltify-finance/joltify_lending/x/kyc/keeper"
	kycmoduletypes "github.com/joltify-finance/joltify_lending/x/kyc/types"
	spvmodulekeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"
	spvmoduletypes "github.com/joltify-finance/joltify_lending/x/spv/types"
	incentivekeeper "github.com/joltify-finance/joltify_lending/x/third_party/incentive/keeper"
	incentivetypes "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"
	"github.com/joltify-finance/joltify_lending/x/third_party/jolt"
	joltkeeper "github.com/joltify-finance/joltify_lending/x/third_party/jolt/keeper"
	jolttypes "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"
	"github.com/joltify-finance/joltify_lending/x/third_party/pricefeed"
	pricefeedkeeper "github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/keeper"
	pricefeedtypes "github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/types"
	swapkeeper "github.com/joltify-finance/joltify_lending/x/third_party/swap/keeper"
	swaptypes "github.com/joltify-finance/joltify_lending/x/third_party/swap/types"

	"github.com/joltify-finance/joltify_lending/x/mint"
	mintkeeper "github.com/joltify-finance/joltify_lending/x/mint/keeper"
	minttypes "github.com/joltify-finance/joltify_lending/x/mint/types"

	"cosmossdk.io/x/evidence"
	evidencekeeper "cosmossdk.io/x/evidence/keeper"
	evidencetypes "cosmossdk.io/x/evidence/types"
	"cosmossdk.io/x/feegrant"
	feegrantkeeper "cosmossdk.io/x/feegrant/keeper"
	feegrantmodule "cosmossdk.io/x/feegrant/module"
	"cosmossdk.io/x/upgrade"
	upgradekeeper "cosmossdk.io/x/upgrade/keeper"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/ibc-go/modules/capability"
	capabilitykeeper "github.com/cosmos/ibc-go/modules/capability/keeper"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	ethermintconfig "github.com/evmos/ethermint/server/config"
	feemarketkeeper "github.com/evmos/ethermint/x/feemarket/keeper"
	feemarkettypes "github.com/evmos/ethermint/x/feemarket/types"

	"github.com/cosmos/ibc-go/v8/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v8/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v8/modules/core"
	ibcclient "github.com/cosmos/ibc-go/v8/modules/core/02-client"
	ibcclienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"

	consensusparamkeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"

	nftmodule "cosmossdk.io/x/nft/module"
	abci "github.com/cometbft/cometbft/abci/types"
	tmjson "github.com/cometbft/cometbft/libs/json"
	"github.com/joltify-finance/joltify_lending/app/ante"
	joltparams "github.com/joltify-finance/joltify_lending/app/params"
	kycmodule "github.com/joltify-finance/joltify_lending/x/kyc"
	spvmodule "github.com/joltify-finance/joltify_lending/x/spv"

	dbm "github.com/cosmos/cosmos-db"
)

const (
	Name = "joltify"
)

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// ModuleBasics manages simple versions of full app modules.
	// It's used for things such as codec registration and genesis file verification.
	ModuleBasics = module.NewBasicManager(
		genutil.AppModuleBasic{},
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(
			[]govclient.ProposalHandler{
				paramsclient.ProposalHandler,
			}),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		ibc.AppModuleBasic{},
		ibctm.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		authzmodule.AppModuleBasic{},
		transfer.AppModuleBasic{},
		vesting.AppModuleBasic{},
		auction.AppModuleBasic{},
		pricefeed.AppModuleBasic{},
		jolt.AppModuleBasic{},
		incentive.AppModuleBasic{},
		kycmodule.AppModuleBasic{},
		ibcratelimit.AppModule{},
		quotamodule.AppModuleBasic{},
		spvmodule.AppModuleBasic{},
		nftmodule.AppModuleBasic{},
		consensus.AppModuleBasic{},
		evm.AppModuleBasic{},
		feemarket.AppModuleBasic{},
		swap.AppModuleBasic{},
		capability.AppModuleBasic{},
		burnauctionmodule.AppModuleBasic{},
	)

	// module account permissions
	// If these are changed, the permissions stored in accounts
	// must also be migrated during a chain upgrade.
	mAccPerms = map[string][]string{
		authtypes.FeeCollectorName:     nil,
		distrtypes.ModuleName:          nil,
		minttypes.ModuleName:           {authtypes.Minter},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:            {authtypes.Burner},
		ibctransfertypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
		auctiontypes.ModuleName:        nil,
		// issuancetypes.ModuleAccountName: {authtypes.Minter, authtypes.Burner},
		// cdptypes.ModuleName:          {authtypes.Minter, authtypes.Burner},
		// cdptypes.LiquidatorMacc:      {authtypes.Minter, authtypes.Burner},
		jolttypes.ModuleName:              {authtypes.Minter, authtypes.Burner},
		incentivetypes.ModuleName:         nil,
		spvmoduletypes.ModuleAccount:      {authtypes.Minter, authtypes.Burner},
		nftmoduletypes.ModuleName:         {authtypes.Minter, authtypes.Burner},
		swaptypes.ModuleName:              nil,
		burnauctionmoduletypes.ModuleName: {authtypes.Burner},
	}
)

// Verify app interface at compile time
// var _ simapp.App = (*App)(nil) // TODO
var _ servertypes.Application = (*App)(nil)

// DefaultOptions is a sensible default Options value.
var DefaultOptions = Options{
	EVMTrace:        ethermintconfig.DefaultEVMTracer,
	EVMMaxGasWanted: ethermintconfig.DefaultMaxTxGasWanted,
}

// Options bundles several configuration params for an App.
type Options struct {
	SkipLoadLatest        bool
	SkipUpgradeHeights    map[int64]bool
	SkipGenesisInvariants bool
	InvariantCheckPeriod  uint
	MempoolEnableAuth     bool
	MempoolAuthAddresses  []sdk.AccAddress
	EVMTrace              string
	EVMMaxGasWanted       uint64
}

// App is the Kava ABCI application.
type App struct {
	*baseapp.BaseApp

	// codec
	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry

	// keys to access the substores
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

	// keepers from all the modules
	accountKeeper    authkeeper.AccountKeeper
	bankKeeper       bankkeeper.Keeper
	capabilityKeeper *capabilitykeeper.Keeper
	stakingKeeper    *stakingkeeper.Keeper
	mintKeeper       mintkeeper.Keeper
	distrKeeper      distrkeeper.Keeper
	feeMarketKeeper  feemarketkeeper.Keeper
	govKeeper        govkeeper.Keeper
	ParamsKeeper     paramskeeper.Keeper
	authzKeeper      authzkeeper.Keeper
	crisisKeeper     *crisiskeeper.Keeper
	slashingKeeper   slashingkeeper.Keeper
	ibcKeeper        *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	upgradeKeeper    *upgradekeeper.Keeper
	evidenceKeeper   *evidencekeeper.Keeper
	transferKeeper   ibctransferkeeper.Keeper
	auctionKeeper    auctionkeeper.Keeper
	// issuanceKeeper   issuancekeeper.Keeper
	pricefeedKeeper pricefeedkeeper.Keeper
	// cdpKeeper        cdpkeeper.Keeper
	joltKeeper              joltkeeper.Keeper
	incentiveKeeper         incentivekeeper.Keeper
	feeGrantKeeper          feegrantkeeper.Keeper
	kycKeeper               kycmodulekeeper.Keeper
	spvKeeper               spvmodulekeeper.Keeper
	burnauctionKeeper       burnauctionmodulekeeper.Keeper
	nftKeeper               nftmodulekeeper.Keeper
	swapKeeper              swapkeeper.Keeper
	ibcQuota                *ibcratelimit.IBCMiddleware
	QuotaKeeper             quotamodulekeeper.Keeper
	RateLimitingICS4Wrapper *ibcratelimit.ICS4Wrapper
	ConsensusParamsKeeper   consensusparamkeeper.Keeper

	// make scoped keepers public for test purposes
	ScopedIBCKeeper      capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper capabilitykeeper.ScopedKeeper

	// the module manager
	mm *module.Manager

	// simulation manager
	sm *module.SimulationManager

	// configurator
	configurator module.Configurator
}

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, ".joltify")
}

// NewApp returns a reference to an initialized App.
func NewApp(
	logger log.Logger,
	db dbm.DB,
	homePath string,
	traceStore io.Writer,
	encodingConfig joltparams.EncodingConfig,
	options Options,
	invCheckPeriod uint,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *App {
	appCodec := encodingConfig.Marshaler
	legacyAmino := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry
	txConfig := encodingConfig.TxConfig

	bApp := baseapp.NewBaseApp(Name, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)
	bApp.SetDisableBlockGasMeter(true)
	bApp.SetTxEncoder(txConfig.TxEncoder())

	keys := storetypes.NewKVStoreKeys(
		authtypes.StoreKey, banktypes.StoreKey, stakingtypes.StoreKey,
		minttypes.StoreKey, distrtypes.StoreKey, slashingtypes.StoreKey,
		govtypes.StoreKey, paramstypes.StoreKey, ibcexported.StoreKey,
		upgradetypes.StoreKey, evidencetypes.StoreKey, ibctransfertypes.StoreKey,
		feegrant.StoreKey, authzkeeper.StoreKey,
		capabilitytypes.StoreKey, auctiontypes.StoreKey,
		crisistypes.StoreKey,
		// issuancetypes.StoreKey,
		pricefeedtypes.StoreKey,
		// cdptypes.StoreKey,
		jolttypes.StoreKey,
		incentivetypes.StoreKey,
		kycmoduletypes.StoreKey,
		spvmoduletypes.StoreKey,
		burnauctionmoduletypes.StoreKey,
		nftmoduletypes.StoreKey,
		quotamoduletypes.StoreKey,
		feemarkettypes.StoreKey,
		swaptypes.StoreKey,
		consensusparamtypes.StoreKey,
	)
	tkeys := storetypes.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := storetypes.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	app := &App{
		BaseApp:           bApp,
		legacyAmino:       legacyAmino,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

	// init params keeper and subspaces
	app.ParamsKeeper = paramskeeper.NewKeeper(
		appCodec,
		legacyAmino,
		keys[paramstypes.StoreKey],
		tkeys[paramstypes.TStoreKey],
	)
	// stakingSubspace := app.ParamsKeeper.Subspace(stakingtypes.ModuleName)
	mintSubspace := app.ParamsKeeper.Subspace(minttypes.ModuleName)
	// distrSubspace := app.ParamsKeeper.Subspace(distrtypes.ModuleName)
	// slashingSubspace := app.ParamsKeeper.Subspace(slashingtypes.ModuleName)
	// crisisSubspace := app.ParamsKeeper.Subspace(crisistypes.ModuleName)
	auctionSubspace := app.ParamsKeeper.Subspace(auctiontypes.ModuleName)
	// issuanceSubspace := app.ParamsKeeper.Subspace(issuancetypes.ModuleName)
	pricefeedSubspace := app.ParamsKeeper.Subspace(pricefeedtypes.ModuleName)
	// cdpSubspace := app.ParamsKeeper.Subspace(cdptypes.ModuleName)
	joltSubspace := app.ParamsKeeper.Subspace(jolttypes.ModuleName)
	incentiveSubspace := app.ParamsKeeper.Subspace(incentivetypes.ModuleName)
	ibcSubspace := app.ParamsKeeper.Subspace(ibcexported.ModuleName)
	ibctransferSubspace := app.ParamsKeeper.Subspace(ibctransfertypes.ModuleName)
	kycSubspace := app.ParamsKeeper.Subspace(kycmoduletypes.ModuleName)
	spvSubspace := app.ParamsKeeper.Subspace(spvmoduletypes.ModuleName)
	ibcQuotaSubspace := app.ParamsKeeper.Subspace(ibcratelimittypes.ModuleName)
	feemarketSubspace := app.ParamsKeeper.Subspace(feemarkettypes.ModuleName)
	swapSubspace := app.ParamsKeeper.Subspace(swaptypes.ModuleName)
	quotaSubspace := app.ParamsKeeper.Subspace(quotamoduletypes.ModuleName)
	burnAuctionSubspace := app.ParamsKeeper.Subspace(burnauctionmoduletypes.ModuleName)
	app.ConsensusParamsKeeper = consensusparamkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[consensusparamtypes.StoreKey]),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		runtime.EventService{},
	)

	// set the BaseApp's parameter store
	bApp.SetParamStore(app.ConsensusParamsKeeper.ParamsStore)

	app.capabilityKeeper = capabilitykeeper.NewKeeper(appCodec, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])
	scopedIBCKeeper := app.capabilityKeeper.ScopeToModule(ibcexported.ModuleName)
	scopedTransferKeeper := app.capabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	app.capabilityKeeper.Seal()

	// add keepers
	app.accountKeeper = authkeeper.NewAccountKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[authtypes.StoreKey]),
		authtypes.ProtoBaseAccount,
		mAccPerms,
		addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
		sdk.Bech32MainPrefix,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	app.feeGrantKeeper = feegrantkeeper.NewKeeper(appCodec, runtime.NewKVStoreService(keys[feegrant.StoreKey]), app.accountKeeper)
	app.bankKeeper = bankkeeper.NewBaseKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[banktypes.StoreKey]),
		app.accountKeeper,
		app.BlockedModuleAccountAddrs(),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		logger,
	)

	app.stakingKeeper = stakingkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[stakingtypes.StoreKey]),
		app.accountKeeper,
		app.bankKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32ValidatorAddrPrefix()),
		addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32ConsensusAddrPrefix()),
	)

	app.authzKeeper = authzkeeper.NewKeeper(
		runtime.NewKVStoreService(keys[authzkeeper.StoreKey]),
		appCodec,
		app.MsgServiceRouter(),
		app.accountKeeper,
	)

	app.distrKeeper = distrkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[distrtypes.StoreKey]),
		app.accountKeeper,
		app.bankKeeper,
		app.stakingKeeper,
		authtypes.FeeCollectorName,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	app.slashingKeeper = slashingkeeper.NewKeeper(
		appCodec,
		encodingConfig.Amino,
		runtime.NewKVStoreService(keys[slashingtypes.StoreKey]),
		app.stakingKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	app.crisisKeeper = crisiskeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[crisistypes.StoreKey]),
		invCheckPeriod,
		app.bankKeeper,
		authtypes.FeeCollectorName,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		app.accountKeeper.AddressCodec(),
	)
	app.upgradeKeeper = upgradekeeper.NewKeeper(
		options.SkipUpgradeHeights,
		runtime.NewKVStoreService(keys[upgradetypes.StoreKey]),
		appCodec,
		homePath,
		app.BaseApp,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	app.evidenceKeeper = evidencekeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[evidencetypes.StoreKey]),
		app.stakingKeeper,
		app.slashingKeeper,
		app.accountKeeper.AddressCodec(),
		runtime.ProvideCometInfoService(),
	)

	app.QuotaKeeper = *quotamodulekeeper.NewKeeper(
		appCodec,
		keys[quotamoduletypes.StoreKey],
		quotaSubspace,
	)

	app.ibcKeeper = ibckeeper.NewKeeper(
		appCodec,
		keys[ibcexported.StoreKey],
		ibcSubspace,
		app.stakingKeeper,
		app.upgradeKeeper,
		scopedIBCKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	rateLimitingICS4Wrapper := ibcratelimit.NewICS4Middleware(
		app.ibcKeeper.ChannelKeeper,
		&app.accountKeeper,
		&app.bankKeeper,
		app.QuotaKeeper,
		ibcQuotaSubspace,
	)
	app.RateLimitingICS4Wrapper = &rateLimitingICS4Wrapper

	app.transferKeeper = ibctransferkeeper.NewKeeper(
		appCodec,
		keys[ibctransfertypes.StoreKey],
		ibctransferSubspace,
		// app.ibcKeeper.ChannelKeeper,
		app.RateLimitingICS4Wrapper,
		app.ibcKeeper.ChannelKeeper,
		app.ibcKeeper.PortKeeper,
		app.accountKeeper,
		app.bankKeeper,
		scopedTransferKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	transferModule := transfer.NewAppModule(app.transferKeeper)
	transferIBCModule := transfer.NewIBCModule(app.transferKeeper)

	app.mintKeeper = *mintkeeper.NewKeeper(
		appCodec,
		keys[minttypes.StoreKey],
		mintSubspace,
		app.accountKeeper,
		app.bankKeeper,
		app.distrKeeper,
		*app.stakingKeeper,
		authtypes.FeeCollectorName,
	)

	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := ibcporttypes.NewRouter()

	transferWithQuota := ibcratelimit.NewIBCMiddleware(transferIBCModule, app.RateLimitingICS4Wrapper)
	app.ibcQuota = &transferWithQuota
	ibcRouter.AddRoute(ibctransfertypes.ModuleName, app.ibcQuota)
	// ibcRouter.AddRoute(ibctransfertypes.ModuleName, &transferIBCModule)
	app.ibcKeeper.SetRouter(ibcRouter)

	app.auctionKeeper = auctionkeeper.NewKeeper(
		appCodec,
		keys[auctiontypes.StoreKey],
		auctionSubspace,
		app.bankKeeper,
		app.accountKeeper,
	)

	app.pricefeedKeeper = pricefeedkeeper.NewKeeper(
		appCodec,
		keys[pricefeedtypes.StoreKey],
		pricefeedSubspace,
	)

	joltKeeper := joltkeeper.NewKeeper(
		appCodec,
		keys[jolttypes.StoreKey],
		joltSubspace,
		app.accountKeeper,
		app.bankKeeper,
		app.pricefeedKeeper,
		app.auctionKeeper,
	)

	mSwapKeeper := swapkeeper.NewKeeper(
		appCodec,
		keys[swaptypes.StoreKey],
		swapSubspace,
		app.accountKeeper,
		app.bankKeeper,
	)

	app.kycKeeper = *kycmodulekeeper.NewKeeper(appCodec, keys[kycmoduletypes.StoreKey], keys[kycmoduletypes.MemStoreKey], kycSubspace, authtypes.NewModuleAddress(govtypes.ModuleName))
	app.nftKeeper = nftmodulekeeper.NewKeeper(runtime.NewKVStoreService(keys[nftmoduletypes.StoreKey]), appCodec, app.accountKeeper, app.bankKeeper)

	mSpvKeeper := spvmodulekeeper.NewKeeper(appCodec, keys[spvmoduletypes.StoreKey], keys[spvmoduletypes.MemStoreKey], spvSubspace, app.kycKeeper, app.bankKeeper, app.accountKeeper, app.nftKeeper, app.pricefeedKeeper, app.auctionKeeper, app.incentiveKeeper)

	app.incentiveKeeper = incentivekeeper.NewKeeper(
		appCodec,
		keys[incentivetypes.StoreKey],
		incentiveSubspace,
		app.bankKeeper,
		// &cdpKeeper,
		&joltKeeper,
		app.accountKeeper,
		mSwapKeeper,
		mSpvKeeper,
		app.nftKeeper,
	)

	app.swapKeeper = *mSwapKeeper.SetHooks(app.incentiveKeeper.Hooks())
	app.spvKeeper = *mSpvKeeper.SetHooks(app.incentiveKeeper.Hooks())
	app.spvKeeper = *mSpvKeeper.SetIncentiveKeeper(app.incentiveKeeper)

	app.burnauctionKeeper = *burnauctionmodulekeeper.NewKeeper(
		appCodec,
		keys[burnauctionmoduletypes.StoreKey],
		keys[burnauctionmoduletypes.MemStoreKey],
		burnAuctionSubspace, app.accountKeeper, app.bankKeeper, app.auctionKeeper)

	// Create Ethermint keepers
	app.feeMarketKeeper = feemarketkeeper.NewKeeper(
		appCodec,
		// Authority
		authtypes.NewModuleAddress(govtypes.ModuleName),
		keys[feemarkettypes.StoreKey],
		feemarketSubspace,
	)

	allKeys := make(map[string]storetypes.StoreKey, len(keys)+len(tkeys)+len(memKeys))
	for k, v := range keys {
		allKeys[k] = v
	}
	for k, v := range tkeys {
		allKeys[k] = v
	}
	for k, v := range memKeys {
		allKeys[k] = v
	}

	govConfig := govtypes.DefaultConfig()
	govKeeper := govkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[govtypes.StoreKey]),
		app.accountKeeper,
		app.bankKeeper,
		app.stakingKeeper,
		app.distrKeeper,
		app.MsgServiceRouter(),
		govConfig,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	// create gov keeper with router
	// NOTE this must be done after any keepers referenced in the gov router (ie committee) are defined
	govRouter := govv1beta1.NewRouter()
	govRouter.
		AddRoute(govtypes.RouterKey, govv1beta1.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.upgradeKeeper)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(app.ibcKeeper.ClientKeeper))

	govKeeper.SetLegacyRouter(govRouter)

	app.govKeeper = *govKeeper.SetHooks(
		govtypes.NewMultiGovHooks(
		// register the governance hooks
		),
	)

	// register the staking hooks
	// NOTE: These keepers are passed by reference above, so they will contain these hooks.
	app.stakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(app.distrKeeper.Hooks(), app.slashingKeeper.Hooks()))

	// app.cdpKeeper = *cdpKeeper.SetHooks(cdptypes.NewMultiCDPHooks(app.incentiveKeeper.Hooks()))
	app.joltKeeper = *joltKeeper.SetHooks(jolttypes.NewMultiJoltHooks(app.incentiveKeeper.Hooks()))

	skipGenesisInvariants := cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))

	// create the module manager (Note: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.)
	authModule := auth.NewAppModule(appCodec, app.accountKeeper, authsims.RandomGenesisAccounts, app.ParamsKeeper.Subspace(authtypes.ModuleName))
	app.mm = module.NewManager(
		genutil.NewAppModule(app.accountKeeper, app.stakingKeeper, app.BaseApp.DeliverTx, encodingConfig.TxConfig),
		authModule,
		bank.NewAppModule(appCodec, app.bankKeeper, app.accountKeeper, app.ParamsKeeper.Subspace(banktypes.ModuleName)),
		capability.NewAppModule(appCodec, *app.capabilityKeeper, false),
		feegrantmodule.NewAppModule(appCodec, app.accountKeeper, app.bankKeeper, app.feeGrantKeeper, app.interfaceRegistry),
		staking.NewAppModule(appCodec, app.stakingKeeper, app.accountKeeper, app.bankKeeper, app.ParamsKeeper.Subspace(stakingtypes.ModuleName)),
		mint.NewAppModule(appCodec, app.mintKeeper, app.accountKeeper, app.bankKeeper),
		distr.NewAppModule(appCodec, app.distrKeeper, app.accountKeeper, app.bankKeeper, app.stakingKeeper, app.ParamsKeeper.Subspace(distrtypes.ModuleName)),
		gov.NewAppModule(appCodec, &app.govKeeper, app.accountKeeper, app.bankKeeper, app.ParamsKeeper.Subspace(govtypes.ModuleName)),
		params.NewAppModule(app.ParamsKeeper),
		slashing.NewAppModule(appCodec, app.slashingKeeper, app.accountKeeper, app.bankKeeper, app.stakingKeeper, app.ParamsKeeper.Subspace(slashingtypes.ModuleName)),
		ibc.NewAppModule(app.ibcKeeper),
		upgrade.NewAppModule(app.upgradeKeeper, addresscodec.NewBech32Codec(sdk.Bech32PrefixAccAddr)),
		evidence.NewAppModule(*app.evidenceKeeper),
		transferModule,
		vesting.NewAppModule(app.accountKeeper, app.bankKeeper),
		authzmodule.NewAppModule(appCodec, app.authzKeeper, app.accountKeeper, app.bankKeeper, app.interfaceRegistry),
		auction.NewAppModule(app.auctionKeeper, app.accountKeeper, app.bankKeeper),
		// issuance.NewAppModule(app.issuanceKeeper, app.accountKeeper, app.bankKeeper),
		pricefeed.NewAppModule(app.pricefeedKeeper, app.accountKeeper),
		// cdp.NewAppModule(app.cdpKeeper, app.accountKeeper, app.pricefeedKeeper, app.bankKeeper),
		jolt.NewAppModule(app.joltKeeper, app.accountKeeper, app.bankKeeper, app.pricefeedKeeper),
		incentive.NewAppModule(app.incentiveKeeper, app.accountKeeper, app.bankKeeper),
		kycmodule.NewAppModule(appCodec, app.kycKeeper, app.accountKeeper, app.bankKeeper),

		spvmodule.NewAppModule(appCodec, app.spvKeeper, app.accountKeeper, app.bankKeeper),
		nftmodule.NewAppModule(appCodec, app.nftKeeper, app.accountKeeper, app.bankKeeper, app.interfaceRegistry),
		ibcratelimit.NewAppModule(*app.RateLimitingICS4Wrapper),
		quotamodule.NewAppModule(appCodec, app.QuotaKeeper),
		feemarket.NewAppModule(app.feeMarketKeeper, feemarketSubspace),
		swap.NewAppModule(app.swapKeeper, app.accountKeeper),
		crisis.NewAppModule(app.crisisKeeper, skipGenesisInvariants, app.ParamsKeeper.Subspace(crisistypes.ModuleName)), // always be last to make sure that it checks for all invariants and not only part of them
		burnauctionmodule.NewAppModule(appCodec, app.burnauctionKeeper, app.accountKeeper, app.bankKeeper),
	)

	// Warning: Some begin blockers must run before others. Ensure the dependencies are understood before modifying this list.
	app.mm.SetOrderBeginBlockers(
		// Upgrade begin blocker runs migrations on the first block after an upgrade. It should run before any other module.
		upgradetypes.ModuleName,
		// Capability begin blocker runs non state changing initialization.
		capabilitytypes.ModuleName,
		// Committee begin blocker changes module params by enacting proposals.
		minttypes.ModuleName,
		distrtypes.ModuleName,
		// During begin block slashing happens after distr.BeginBlocker so that
		// there is nothing left over in the validator fee pool, so as to keep the
		// CanWithdrawInvariant invariant.
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		stakingtypes.ModuleName,
		feemarkettypes.ModuleName,
		feegrant.ModuleName,
		// Auction begin blocker will close out expired auctions and pay debt back to cdp.
		// It should be run before cdp begin blocker which cancels out debt with stable and starts more auctions.
		auctiontypes.ModuleName,
		// cdptypes.ModuleName,
		jolttypes.ModuleName,
		swaptypes.ModuleName,
		// issuancetypes.ModuleName,
		incentivetypes.ModuleName,
		kycmoduletypes.ModuleName,
		nftmoduletypes.ModuleName,
		ibcratelimittypes.ModuleName,
		spvmoduletypes.ModuleName,
		ibcexported.ModuleName,
		// Add all remaining modules with an empty begin blocker below since cosmos 0.45.0 requires it
		vestingtypes.ModuleName,
		pricefeedtypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		govtypes.ModuleName,
		crisistypes.ModuleName,
		genutiltypes.ModuleName,
		quotamoduletypes.ModuleName,
		ibctransfertypes.ModuleName,
		paramstypes.ModuleName,
		authz.ModuleName,
		burnauctionmoduletypes.ModuleName,
	)

	// Warning: Some end blockers must run before others. Ensure the dependencies are understood before modifying this list.
	app.mm.SetOrderEndBlockers(
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		feemarkettypes.ModuleName,
		pricefeedtypes.ModuleName,
		// Add all remaining modules with an empty end blocker below since cosmos 0.45.0 requires it
		capabilitytypes.ModuleName,
		// issuancetypes.ModuleName,
		minttypes.ModuleName,
		slashingtypes.ModuleName,
		distrtypes.ModuleName,
		auctiontypes.ModuleName,
		// cdptypes.ModuleName,
		jolttypes.ModuleName,
		swaptypes.ModuleName,
		kycmoduletypes.ModuleName,
		nftmoduletypes.ModuleName,
		spvmoduletypes.ModuleName,
		incentivetypes.ModuleName,
		quotamoduletypes.ModuleName,
		ibcratelimittypes.ModuleName,
		upgradetypes.ModuleName,
		evidencetypes.ModuleName,
		feegrant.ModuleName,
		vestingtypes.ModuleName,
		ibcexported.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		genutiltypes.ModuleName,
		ibctransfertypes.ModuleName,
		paramstypes.ModuleName,
		authz.ModuleName,
		burnauctionmoduletypes.ModuleName,
	)

	// Warning: Some init genesis methods must run before others. Ensure the dependencies are understood before modifying this list
	app.mm.SetOrderInitGenesis(
		capabilitytypes.ModuleName, // initialize capabilities, run before any module creating or claiming capabilities in InitGenesis
		authtypes.ModuleName,       // loads all accounts, run before any module with a module account
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		slashingtypes.ModuleName, // iterates over validators, run after staking
		govtypes.ModuleName,
		minttypes.ModuleName,
		ibcexported.ModuleName,
		evidencetypes.ModuleName,
		authz.ModuleName,
		ibctransfertypes.ModuleName,
		feemarkettypes.ModuleName,
		feegrant.ModuleName,
		auctiontypes.ModuleName,
		// issuancetypes.ModuleName,
		pricefeedtypes.ModuleName,
		// cdptypes.ModuleName, // reads market prices, so must run after pricefeed genesis
		jolttypes.ModuleName,
		swaptypes.ModuleName,
		kycmoduletypes.ModuleName,
		nftmoduletypes.ModuleName,
		spvmoduletypes.ModuleName,
		quotamoduletypes.ModuleName,
		ibcratelimittypes.ModuleName,
		incentivetypes.ModuleName, // reads cdp params, so must run after cdp genesis
		genutiltypes.ModuleName,   // runs arbitrary txs included in genisis state, so run after modules have been initialized
		crisistypes.ModuleName,    // runs the invariants at genesis, should run after other modules
		// Add all remaining modules with an empty InitGenesis below since cosmos 0.45.0 requires it
		vestingtypes.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		burnauctionmoduletypes.ModuleName,
	)

	app.mm.RegisterInvariants(app.crisisKeeper)
	app.configurator = module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.mm.RegisterServices(app.configurator)

	autocliv1.RegisterQueryServer(app.GRPCQueryRouter(), runtimeservices.NewAutoCLIQueryService(app.mm.Modules))
	reflectionSvc, err := runtimeservices.NewReflectionService()
	if err != nil {
		panic(err)
	}
	reflectionv1.RegisterReflectionServiceServer(app.GRPCQueryRouter(), reflectionSvc)

	// create the simulation manager and define the order of the modules for deterministic simulations
	overrideModules := map[string]module.AppModuleSimulation{
		authtypes.ModuleName: authModule,
	}
	app.sm = module.NewSimulationManagerFromAppModules(app.mm.Modules, overrideModules)
	app.sm.RegisterStoreDecoders()

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: This is not required for apps that don't use the simulator for fuzz testing
	// transactions.
	// TODO
	// app.sm = module.NewSimulationManager(
	// 	auth.NewAppModule(app.accountKeeper),
	// 	bank.NewAppModule(app.bankKeeper, app.accountKeeper),
	// 	gov.NewAppModule(app.govKeeper, app.accountKeeper, app.accountKeeper, app.bankKeeper),
	// 	mint.NewAppModule(app.mintKeeper),
	// 	distr.NewAppModule(app.distrKeeper, app.accountKeeper, app.accountKeeper, app.bankKeeper, app.stakingKeeper),
	//  staking.NewAppModule(app.stakingKeeper, app.accountKeeper, app.accountKeeper, app.bankKeeper),
	// 	slashing.NewAppModule(app.slashingKeeper, app.accountKeeper, app.stakingKeeper),
	// )
	// app.sm.RegisterStoreDecoders()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	// baseAnte := cosante.HandlerOptions{
	//	AccountKeeper:   app.accountKeeper,
	//	BankKeeper:      app.bankKeeper,
	//	SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
	//	FeegrantKeeper:  app.feeGrantKeeper,
	//	SigGasConsumer:  cosante.DefaultSigVerificationGasConsumer,
	//}

	extensionCheck := func(a *types.Any) bool {
		// todo we need to verify here, currently, we allow all the tx to be passed
		return true
	}

	anteOptions := ante.HandlerOptions{
		AccountKeeper:          app.accountKeeper,
		BankKeeper:             app.bankKeeper,
		SignModeHandler:        encodingConfig.TxConfig.SignModeHandler(),
		FeegrantKeeper:         app.feeGrantKeeper,
		SigGasConsumer:         evmante.DefaultSigVerificationGasConsumer,
		SpvKeeper:              app.spvKeeper,
		IBCKeeper:              app.ibcKeeper,
		FeeMarketKeeper:        app.feeMarketKeeper,
		MaxTxGasWanted:         options.EVMMaxGasWanted,
		AddressFetchers:        []ante.AddressFetcher{},
		ExtensionOptionChecker: extensionCheck,
		TxFeeChecker:           nil,
	}

	anteHandler, err := ante.NewAnteHandler(anteOptions, app.ConsensusParamsKeeper)
	if err != nil {
		panic(fmt.Sprintf("failed to create anteHandler: %s", err))
	}

	app.SetAnteHandler(anteHandler)
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)

	app.ScopedIBCKeeper = scopedIBCKeeper
	app.ScopedTransferKeeper = scopedTransferKeeper
	app.setupUpgradeHandlers()

	// load store
	if !options.SkipLoadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}
	}

	return app
}

// Name returns the name of the App
func (app *App) Name() string { return app.BaseApp.Name() }

// BeginBlocker contains app specific logic for the BeginBlock abci call.
func (app *App) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker contains app specific logic for the EndBlock abci call.
func (app *App) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// InitChainer contains app specific logic for the InitChain abci call.
func (app *App) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	if err := tmjson.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}

	// Store current module versions in joltify-10 to setup future in-place upgrades.
	// During in-place migrations, the old module versions in the store will be referenced to determine which migrations to run.
	app.upgradeKeeper.SetModuleVersionMap(ctx, app.mm.GetVersionMap())

	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads the app state for a particular height.
func (app *App) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *App) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range mAccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// InterfaceRegistry returns the app's InterfaceRegistry.
func (app *App) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// SimulationManager implements the SimulationApp interface.
func (app *App) SimulationManager() *module.SimulationManager {
	return app.sm
}

// RegisterAPIRoutes registers all application module routes with the provided API server.
func (app *App) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx

	// Register GRPC Gateway routes
	cmtservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	if apiConfig.Swagger {
		RegisterSwaggerAPI(apiSvr.Router)
	}

	// Swagger API configuration is ignored
	// apiSvr.Router.Handle("/static/openapi.yml", http.FileServer(http.FS(docs.Docs)))
	// apiSvr.Router.HandleFunc("/", openapiconsole.Handler(Name, "/static/openapi.yml"))
}

// RegisterSwaggerAPI registers swagger route with API Server.
func RegisterSwaggerAPI(rtr *mux.Router) {
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	staticServer := http.FileServer(statikFS)
	rtr.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", staticServer))
}

// RegisterTxService implements the Application.RegisterTxService method.
// It registers transaction related endpoints on the app's grpc server.
func (app *App) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
// It registers the standard tendermint grpc endpoints on the app's grpc server.
func (app *App) RegisterTendermintService(clientCtx client.Context) {
	cmtservice.RegisterTendermintService(clientCtx, app.BaseApp.GRPCQueryRouter(), app.interfaceRegistry, app.Query)
}

// loadBlockedMaccAddrs returns a map indicating the blocked status of each module account address
// func (app *App) loadBlockedMaccAddrs() map[string]bool {
//	modAccAddrs := app.ModuleAccountAddrs()
//	joltIncentiveMaccAddr := app.accountKeeper.GetModuleAddress(incentivetypes.ModuleName)
//	for addr := range modAccAddrs {
//		// Set the joltincentives module account address as unblocked
//		if addr == joltIncentiveMaccAddr.String() {
//			modAccAddrs[addr] = false
//		}
//	}
//	return modAccAddrs
// }

func (app *App) setupUpgradeHandlers() {
	upgradeInfo, err := app.upgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	if upgradeInfo.Name == "v013_upgrade" {
		storeUpgrades := storetypes.StoreUpgrades{
			Added: []string{burnauctionmoduletypes.StoreKey},
		}

		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}

	// Set param key table for params module migration
	for _, subspace := range app.ParamsKeeper.GetSubspaces() {
		subspace := subspace
		var keyTable paramstypes.KeyTable
		switch subspace.Name() {
		case authtypes.ModuleName:
			keyTable = authtypes.ParamKeyTable() //nolint:staticcheck
		case banktypes.ModuleName:
			keyTable = banktypes.ParamKeyTable() //nolint:staticcheck
		case stakingtypes.ModuleName:
			keyTable = stakingtypes.ParamKeyTable()
		case minttypes.ModuleName:
			keyTable = minttypes.ParamKeyTable()
		case distrtypes.ModuleName:
			keyTable = distrtypes.ParamKeyTable() //nolint:staticcheck
		case slashingtypes.ModuleName:
			keyTable = slashingtypes.ParamKeyTable() //nolint:staticcheck
		case govtypes.ModuleName:
			keyTable = govv1.ParamKeyTable() //nolint:staticcheck
		case crisistypes.ModuleName:
			keyTable = crisistypes.ParamKeyTable() //nolint:staticcheck

		case feemarkettypes.ModuleName:
			keyTable = feemarkettypes.ParamKeyTable()
		}
		if !subspace.HasKeyTable() {
			subspace.WithKeyTable(keyTable)
		}
	}

	app.upgradeKeeper.SetUpgradeHandler(v1.V011UpgradeName, v1.CreateUpgradeHandlerForV011Upgrade(app.mm, app.configurator, app.kycKeeper, app.spvKeeper, app.QuotaKeeper, app.incentiveKeeper))
	app.upgradeKeeper.SetUpgradeHandler(v1.V012UpgradeName, v1.CreateUpgradeHandlerForV012Upgrade(app.mm, app.configurator))
	app.upgradeKeeper.SetUpgradeHandler(v1.V013UpgradeName, v1.CreateUpgradeHandlerForV013Upgrade(app.mm, app.configurator))
}

// RegisterNodeService implements the Application.RegisterNodeService method.
func (app *App) RegisterNodeService(clientCtx client.Context, cfg config.Config) {
	nodeservice.RegisterNodeService(clientCtx, app.GRPCQueryRouter(), cfg)
}

// BlockedModuleAccountAddrs returns all the app's blocked module account
// addresses.
func (app *App) BlockedModuleAccountAddrs() map[string]bool {
	modAccAddrs := app.ModuleAccountAddrs()
	delete(modAccAddrs, authtypes.NewModuleAddress(govtypes.ModuleName).String())
	delete(modAccAddrs, authtypes.NewModuleAddress(incentivetypes.IncentiveMacc).String())

	return modAccAddrs
}
