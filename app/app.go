package app

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	types2 "github.com/joltify-finance/joltify_lending/app/ante/types"
	"github.com/joltify-finance/joltify_lending/lib/metrics"

	"github.com/joltify-finance/joltify_lending/app/middleware"

	"github.com/joltify-finance/joltify_lending/dydx_helper/mempool"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	dydxante "github.com/joltify-finance/joltify_lending/app/ante/dydx_ante"

	"github.com/cosmos/cosmos-sdk/types/msgservice"
	"github.com/cosmos/gogoproto/proto"

	config2 "github.com/joltify-finance/joltify_lending/app/config"
	appFlag "github.com/joltify-finance/joltify_lending/app/flags"
	daemonflags "github.com/joltify-finance/joltify_lending/daemons/flags"
	liquidationtypes "github.com/joltify-finance/joltify_lending/daemons/server/types/liquidations"
	daemontypes "github.com/joltify-finance/joltify_lending/daemons/types"
	"github.com/joltify-finance/joltify_lending/dydx_helper/indexer"
	"github.com/joltify-finance/joltify_lending/dydx_helper/indexer/indexer_manager"
	"github.com/joltify-finance/joltify_lending/dydx_helper/indexer/msgsender"
	streaming "github.com/joltify-finance/joltify_lending/dydx_helper/streaming/grpc"
	streamingtypes "github.com/joltify-finance/joltify_lending/dydx_helper/streaming/grpc/types"
	"github.com/joltify-finance/joltify_lending/lib"
	timelib "github.com/joltify-finance/joltify_lending/lib/time"
	assetsmodule "github.com/joltify-finance/joltify_lending/x/third_party_dydx/assets"

	bridgedaemontypes "github.com/joltify-finance/joltify_lending/daemons/server/types/bridge"
	bridgemodule "github.com/joltify-finance/joltify_lending/x/third_party_dydx/bridge"
	bridgemoduletypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/bridge/types"
	clobmodule "github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob"
	clobflags "github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/flags"
	clobmodulememclob "github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/memclob"
	"github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/rate_limit"
	delaymsgmodule "github.com/joltify-finance/joltify_lending/x/third_party_dydx/delaymsg"
	delaymsgmoduletypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/delaymsg/types"
	epochsmodule "github.com/joltify-finance/joltify_lending/x/third_party_dydx/epochs"
	feetiersmodule "github.com/joltify-finance/joltify_lending/x/third_party_dydx/feetiers"
	perpetualsmodule "github.com/joltify-finance/joltify_lending/x/third_party_dydx/perpetuals"
	perpetualsmoduletypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/perpetuals/types"
	dydxpricesmodule "github.com/joltify-finance/joltify_lending/x/third_party_dydx/prices"
	dydxpricesmoduletypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/prices/types"
	rewardsmodule "github.com/joltify-finance/joltify_lending/x/third_party_dydx/rewards"
	statsmodule "github.com/joltify-finance/joltify_lending/x/third_party_dydx/stats"
	subaccountsmodule "github.com/joltify-finance/joltify_lending/x/third_party_dydx/subaccounts"
	satypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/subaccounts/types"
	// vestmodule "github.com/joltify-finance/joltify_lending/x/third_party_dydx/vest"
	"google.golang.org/grpc"

	"cosmossdk.io/client/v2/autocli"
	"cosmossdk.io/core/appmodule"
	daemonservertypes "github.com/joltify-finance/joltify_lending/daemons/server/types"

	tmos "github.com/cometbft/cometbft/libs/os"
	authcodec "github.com/cosmos/cosmos-sdk/x/auth/codec"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	ibcconnectiontypes "github.com/cosmos/ibc-go/v8/modules/core/03-connection/types"
	solomachine "github.com/cosmos/ibc-go/v8/modules/light-clients/06-solomachine"
	clobmoduletypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/types"
	epochsmoduletypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/epochs/types"
	statsmodulekeeper "github.com/joltify-finance/joltify_lending/x/third_party_dydx/stats/keeper"
	statsmoduletypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/stats/types"

	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/cosmos/cosmos-sdk/server"

	"cosmossdk.io/log"

	"github.com/cosmos/cosmos-sdk/runtime"

	addresscodec "github.com/cosmos/cosmos-sdk/codec/address"

	burnauctionmodule "github.com/joltify-finance/joltify_lending/x/burnauction"
	burnauctionmoduletypes "github.com/joltify-finance/joltify_lending/x/burnauction/types"

	burnauctionmodulekeeper "github.com/joltify-finance/joltify_lending/x/burnauction/keeper"

	v1 "github.com/joltify-finance/joltify_lending/upgrade"

	quotamodule "github.com/joltify-finance/joltify_lending/x/quota"
	quotamodulekeeper "github.com/joltify-finance/joltify_lending/x/quota/keeper"
	quotamoduletypes "github.com/joltify-finance/joltify_lending/x/quota/types"

	clobmodulekeeper "github.com/joltify-finance/joltify_lending/x/third_party_dydx/clob/keeper"

	ibcratelimit "github.com/joltify-finance/joltify_lending/x/ibc-rate-limit"
	ibcratelimittypes "github.com/joltify-finance/joltify_lending/x/ibc-rate-limit/types"

	ibctm "github.com/cosmos/ibc-go/v8/modules/light-clients/07-tendermint"

	subaccountsmodulekeeper "github.com/joltify-finance/joltify_lending/x/third_party_dydx/subaccounts/keeper"

	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	reflectionv1 "cosmossdk.io/api/cosmos/reflection/v1"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/spf13/cast"

	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"

	"github.com/cosmos/cosmos-sdk/x/consensus"

	"github.com/joltify-finance/joltify_lending/x/third_party/swap"

	runtimeservices "github.com/cosmos/cosmos-sdk/runtime/services"

	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	"github.com/gorilla/mux"
	_ "github.com/joltify-finance/joltify_lending/client/docs/statik"

	nodeservice "github.com/cosmos/cosmos-sdk/client/grpc/node"
	"github.com/rakyll/statik/fs"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/server/config"
	ibcporttypes "github.com/cosmos/ibc-go/v8/modules/core/05-port/types"
	"github.com/joltify-finance/joltify_lending/x/third_party/auction"
	auctionkeeper "github.com/joltify-finance/joltify_lending/x/third_party/auction/keeper"
	auctiontypes "github.com/joltify-finance/joltify_lending/x/third_party/auction/types"
	"github.com/joltify-finance/joltify_lending/x/third_party/incentive"

	nftmoduletypes "cosmossdk.io/x/nft"
	nftmodulekeeper "cosmossdk.io/x/nft/keeper"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
	appconfig "github.com/joltify-finance/joltify_lending/app/config"
	dydxPricefeedtypes "github.com/joltify-finance/joltify_lending/daemons/server/types/pricefeed"
	kycmodulekeeper "github.com/joltify-finance/joltify_lending/x/kyc/keeper"
	kycmoduletypes "github.com/joltify-finance/joltify_lending/x/kyc/types"
	spvmodulekeeper "github.com/joltify-finance/joltify_lending/x/spv/keeper"
	spvmoduletypes "github.com/joltify-finance/joltify_lending/x/spv/types"
	incentivekeeper "github.com/joltify-finance/joltify_lending/x/third_party/incentive/keeper"
	incentivetypes "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"
	"github.com/joltify-finance/joltify_lending/x/third_party/jolt"
	joltkeeper "github.com/joltify-finance/joltify_lending/x/third_party/jolt/keeper"
	jolttypes "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"
	kavapricefeed "github.com/joltify-finance/joltify_lending/x/third_party/pricefeed"
	kavapricefeedkeeper "github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/keeper"
	kavapricefeedtypes "github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/types"
	swapkeeper "github.com/joltify-finance/joltify_lending/x/third_party/swap/keeper"
	swaptypes "github.com/joltify-finance/joltify_lending/x/third_party/swap/types"
	blocktimemodulekeeper "github.com/joltify-finance/joltify_lending/x/third_party_dydx/blocktime/keeper"

	daemonserver "github.com/joltify-finance/joltify_lending/daemons/server"

	assetsmodulekeeper "github.com/joltify-finance/joltify_lending/x/third_party_dydx/assets/keeper"
	assetsmoduletypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/assets/types"
	blocktimemoduletypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/blocktime/types"
	epochsmodulekeeper "github.com/joltify-finance/joltify_lending/x/third_party_dydx/epochs/keeper"
	feetiersmodulekeeper "github.com/joltify-finance/joltify_lending/x/third_party_dydx/feetiers/keeper"
	feetiersmoduletypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/feetiers/types"
	govplusmoduletypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/govplus/types"
	rewardsmodulekeeper "github.com/joltify-finance/joltify_lending/x/third_party_dydx/rewards/keeper"
	rewardsmoduletypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/rewards/types"
	sendingmoduletypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/sending/types"
	vaultmoduletypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/vault/types"
	vestmoduletypes "github.com/joltify-finance/joltify_lending/x/third_party_dydx/vest/types"

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
	bridgemodulekeeper "github.com/joltify-finance/joltify_lending/x/third_party_dydx/bridge/keeper"
	delaymsgmodulekeeper "github.com/joltify-finance/joltify_lending/x/third_party_dydx/delaymsg/keeper"
	dydxpricesmodulekeeper "github.com/joltify-finance/joltify_lending/x/third_party_dydx/prices/keeper"
	sendingmodulekeeper "github.com/joltify-finance/joltify_lending/x/third_party_dydx/sending/keeper"
	vaultmodulekeeper "github.com/joltify-finance/joltify_lending/x/third_party_dydx/vault/keeper"
	vestmodulekeeper "github.com/joltify-finance/joltify_lending/x/third_party_dydx/vest/keeper"

	"github.com/cosmos/ibc-go/v8/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v8/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v8/modules/core"
	ibcclienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"

	consensusparamkeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"
	perpetualsmodulekeeper "github.com/joltify-finance/joltify_lending/x/third_party_dydx/perpetuals/keeper"

	nftmodule "cosmossdk.io/x/nft/module"
	abci "github.com/cometbft/cometbft/abci/types"
	tmjson "github.com/cometbft/cometbft/libs/json"
	jante "github.com/joltify-finance/joltify_lending/app/ante"
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
		kavapricefeed.AppModuleBasic{},
		jolt.AppModuleBasic{},
		incentive.AppModuleBasic{},
		kycmodule.AppModuleBasic{},
		ibcratelimit.AppModule{},
		quotamodule.AppModuleBasic{},
		spvmodule.AppModuleBasic{},
		nftmodule.AppModuleBasic{},
		consensus.AppModuleBasic{},
		swap.AppModuleBasic{},
		capability.AppModuleBasic{},
		burnauctionmodule.AppModuleBasic{},

		// dydx

	)

	preBlockers = []string{
		upgradetypes.ModuleName,
		clobmoduletypes.ModuleName,
		dydxpricesmoduletypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/preBlockers
	}

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
var DefaultOptions = Options{}

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
	txConfig          client.TxConfig

	// keys to access the substores
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

	// keepers from all the modules
	AccountKeeper    authkeeper.AccountKeeper
	BankKeeper       bankkeeper.Keeper
	capabilityKeeper *capabilitykeeper.Keeper
	stakingKeeper    *stakingkeeper.Keeper
	mintKeeper       mintkeeper.Keeper
	distrKeeper      distrkeeper.Keeper
	GovKeeper        govkeeper.Keeper
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
	kavaPricefeedKeeper kavapricefeedkeeper.Keeper
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

	// dydx keepers
	ClobKeeper        *clobmodulekeeper.Keeper
	PerpetualsKeeper  *perpetualsmodulekeeper.Keeper
	DydxPricesKeeper  dydxpricesmodulekeeper.Keeper
	StatsKeeper       statsmodulekeeper.Keeper
	SubaccountsKeeper subaccountsmodulekeeper.Keeper
	BlockTimeKeeper   blocktimemodulekeeper.Keeper

	AssetsKeeper   assetsmodulekeeper.Keeper
	FeeTiersKeeper *feetiersmodulekeeper.Keeper
	RewardsKeeper  rewardsmodulekeeper.Keeper
	EpochsKeeper   epochsmodulekeeper.Keeper

	BridgeKeeper bridgemodulekeeper.Keeper

	DelayMsgKeeper delaymsgmodulekeeper.Keeper

	VestKeeper vestmodulekeeper.Keeper

	SendingKeeper sendingmodulekeeper.Keeper

	VaultKeeper vaultmodulekeeper.Keeper

	// the module manager
	ModuleManager *module.Manager

	BasicModuleManager module.BasicManager

	// simulation manager
	sm *module.SimulationManager

	// configurator
	configurator module.Configurator

	// dydx related

	IndexerEventManager  indexer_manager.IndexerEventManager
	GrpcStreamingManager streamingtypes.GrpcStreamingManager
	DaemonHealthMonitor  *daemonservertypes.HealthMonitor
	Server               *daemonserver.Server
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
	traceStore io.Writer,
	loadLatest bool,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *App {
	encodingConfig := appconfig.MakeEncodingConfig()
	appCodec := encodingConfig.Marshaler
	legacyAmino := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry
	txConfig := encodingConfig.TxConfig

	invCheckPeriod := cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod))

	// get skipUpgradeHeights from the app options
	skipUpgradeHeights := map[int64]bool{}
	for _, h := range cast.ToIntSlice(appOpts.Get(server.FlagUnsafeSkipUpgrades)) {
		skipUpgradeHeights[int64(h)] = true
	}

	homePath := cast.ToString(appOpts.Get(flags.FlagHome))
	if homePath == "" {
		homePath = DefaultNodeHome
	}

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
		kavapricefeedtypes.StoreKey,
		// cdptypes.StoreKey,
		jolttypes.StoreKey,
		incentivetypes.StoreKey,
		kycmoduletypes.StoreKey,
		spvmoduletypes.StoreKey,
		burnauctionmoduletypes.StoreKey,
		nftmoduletypes.StoreKey,
		quotamoduletypes.StoreKey,
		swaptypes.StoreKey,
		consensusparamtypes.StoreKey,

		// dydx
		dydxpricesmoduletypes.StoreKey,
		assetsmoduletypes.StoreKey,
		blocktimemoduletypes.StoreKey,
		bridgemoduletypes.StoreKey,
		feetiersmoduletypes.StoreKey,
		perpetualsmoduletypes.StoreKey,
		satypes.StoreKey,
		statsmoduletypes.StoreKey,
		vestmoduletypes.StoreKey,
		rewardsmoduletypes.StoreKey,
		clobmoduletypes.StoreKey,
		sendingmoduletypes.StoreKey,
		delaymsgmoduletypes.StoreKey,
		epochsmoduletypes.StoreKey,
		govplusmoduletypes.StoreKey,
		vaultmoduletypes.StoreKey,
	)

	keys[authtypes.StoreKey] = keys[authtypes.StoreKey].WithLocking()

	tkeys := storetypes.NewTransientStoreKeys(
		paramstypes.TStoreKey,
		clobmoduletypes.TransientStoreKey,
		statsmoduletypes.TransientStoreKey,
		rewardsmoduletypes.TransientStoreKey,
		indexer_manager.TransientStoreKey,
		perpetualsmoduletypes.TransientStoreKey,
	)

	memKeys := storetypes.NewMemoryStoreKeys(capabilitytypes.MemStoreKey, clobmoduletypes.MemStoreKey)

	// load state streaming if enabled
	if err := bApp.RegisterStreamingServices(appOpts, keys); err != nil {
		fmt.Printf("failed to load state streaming: %s", err)
		os.Exit(1)
	}

	app := &App{
		BaseApp:           bApp,
		legacyAmino:       legacyAmino,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
		txConfig:          txConfig,
	}

	timeProvider := &timelib.TimeProviderImpl{}

	appFlagsInstance := appFlag.GetFlagValuesFromOptions(appOpts)
	app.GrpcStreamingManager = getGrpcStreamingManagerFromOptions(appFlagsInstance, logger)

	msgSender, indexerFlags := getIndexerFromOptions(appOpts, logger)
	app.IndexerEventManager = indexer_manager.NewIndexerEventManager(
		msgSender,
		tkeys[indexer_manager.TransientStoreKey],
		indexerFlags.SendOffchainData,
	)

	// Get Daemon Flags.
	daemonFlags := daemonflags.GetDaemonFlagValuesFromOptions(appOpts)
	logger.Info("Parsed Daemon flags", "Flags", daemonFlags)

	// Create server that will ingest gRPC messages from daemon clients.
	// Note that gRPC clients will block on new gRPC connection until the gRPC server is ready to
	// accept new connections.
	app.Server = daemonserver.NewServer(
		logger,
		grpc.NewServer(),
		&daemontypes.FileHandlerImpl{},
		daemonFlags.Shared.SocketAddress,
	)
	// Setup server for pricefeed messages. The server will wait for gRPC messages containing price
	// updates and then encode them into an in-memory cache shared by the prices module.
	// The in-memory data structure is shared by the x/prices module and PriceFeed daemon.
	indexPriceCache := dydxPricefeedtypes.NewMarketToExchangePrices(dydxPricefeedtypes.MaxPriceAge)
	app.Server.WithPriceFeedMarketToExchangePrices(indexPriceCache)

	// Setup server for liquidation messages. The server will wait for gRPC messages containing
	// potentially liquidatable subaccounts and then encode them into an in-memory slice shared by
	// the liquidations module.
	// The in-memory data structure is shared by the x/clob module and liquidations daemon.
	daemonLiquidationInfo := liquidationtypes.NewDaemonLiquidationInfo()
	app.Server.WithDaemonLiquidationInfo(daemonLiquidationInfo)

	// Setup server for bridge messages.
	// The in-memory data structure is shared by the x/bridge module and bridge daemon.
	bridgeEventManager := bridgedaemontypes.NewBridgeEventManager(timeProvider)
	app.Server.WithBridgeEventManager(bridgeEventManager)

	// init params keeper and subspaces
	app.ParamsKeeper = paramskeeper.NewKeeper(
		appCodec,
		legacyAmino,
		keys[paramstypes.StoreKey],
		tkeys[paramstypes.TStoreKey],
	)

	keyTable := ibcclienttypes.ParamKeyTable()
	keyTable.RegisterParamSet(&ibcconnectiontypes.Params{})
	keyTable.RegisterParamSet(&ibctransfertypes.Params{})

	// stakingSubspace := app.ParamsKeeper.Subspace(stakingtypes.ModuleName)
	mintSubspace := app.ParamsKeeper.Subspace(minttypes.ModuleName)
	// distrSubspace := app.ParamsKeeper.Subspace(distrtypes.ModuleName)
	// slashingSubspace := app.ParamsKeeper.Subspace(slashingtypes.ModuleName)
	// crisisSubspace := app.ParamsKeeper.Subspace(crisistypes.ModuleName)
	auctionSubspace := app.ParamsKeeper.Subspace(auctiontypes.ModuleName)
	// issuanceSubspace := app.ParamsKeeper.Subspace(issuancetypes.ModuleName)
	kavaPricefeedSubspace := app.ParamsKeeper.Subspace(kavapricefeedtypes.ModuleName)
	// cdpSubspace := app.ParamsKeeper.Subspace(cdptypes.ModuleName)
	joltSubspace := app.ParamsKeeper.Subspace(jolttypes.ModuleName)
	incentiveSubspace := app.ParamsKeeper.Subspace(incentivetypes.ModuleName)
	ibcSubspace := app.ParamsKeeper.Subspace(ibcexported.ModuleName).WithKeyTable(keyTable)
	ibctransferSubspace := app.ParamsKeeper.Subspace(ibctransfertypes.ModuleName).WithKeyTable(keyTable)
	kycSubspace := app.ParamsKeeper.Subspace(kycmoduletypes.ModuleName)
	spvSubspace := app.ParamsKeeper.Subspace(spvmoduletypes.ModuleName)
	ibcQuotaSubspace := app.ParamsKeeper.Subspace(ibcratelimittypes.ModuleName)
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
	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[authtypes.StoreKey]),
		authtypes.ProtoBaseAccount,
		mAccPerms,
		addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
		sdk.Bech32MainPrefix,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	app.feeGrantKeeper = feegrantkeeper.NewKeeper(appCodec, runtime.NewKVStoreService(keys[feegrant.StoreKey]), app.AccountKeeper)
	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[banktypes.StoreKey]),
		app.AccountKeeper,
		app.BlockedModuleAccountAddrs(),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		logger,
	)

	app.stakingKeeper = stakingkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[stakingtypes.StoreKey]),
		app.AccountKeeper,
		app.BankKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32ValidatorAddrPrefix()),
		addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32ConsensusAddrPrefix()),
	)

	app.authzKeeper = authzkeeper.NewKeeper(
		runtime.NewKVStoreService(keys[authzkeeper.StoreKey]),
		appCodec,
		app.MsgServiceRouter(),
		app.AccountKeeper,
	)

	app.distrKeeper = distrkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[distrtypes.StoreKey]),
		app.AccountKeeper,
		app.BankKeeper,
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
		app.BankKeeper,
		authtypes.FeeCollectorName,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		app.AccountKeeper.AddressCodec(),
	)
	app.upgradeKeeper = upgradekeeper.NewKeeper(
		skipUpgradeHeights,
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
		app.AccountKeeper.AddressCodec(),
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
		&app.AccountKeeper,
		&app.BankKeeper,
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
		app.AccountKeeper,
		app.BankKeeper,
		scopedTransferKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	transferModule := transfer.NewAppModule(app.transferKeeper)
	transferIBCModule := transfer.NewIBCModule(app.transferKeeper)

	app.mintKeeper = *mintkeeper.NewKeeper(
		appCodec,
		keys[minttypes.StoreKey],
		mintSubspace,
		app.AccountKeeper,
		app.BankKeeper,
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
		app.BankKeeper,
		app.AccountKeeper,
	)

	app.kavaPricefeedKeeper = kavapricefeedkeeper.NewKeeper(
		appCodec,
		keys[kavapricefeedtypes.StoreKey],
		kavaPricefeedSubspace,
	)

	joltKeeper := joltkeeper.NewKeeper(
		appCodec,
		keys[jolttypes.StoreKey],
		joltSubspace,
		app.AccountKeeper,
		app.BankKeeper,
		app.kavaPricefeedKeeper,
		app.auctionKeeper,
	)

	mSwapKeeper := swapkeeper.NewKeeper(
		appCodec,
		keys[swaptypes.StoreKey],
		swapSubspace,
		app.AccountKeeper,
		app.BankKeeper,
	)

	app.kycKeeper = *kycmodulekeeper.NewKeeper(appCodec, keys[kycmoduletypes.StoreKey], keys[kycmoduletypes.MemStoreKey], kycSubspace, authtypes.NewModuleAddress(govtypes.ModuleName))
	app.nftKeeper = nftmodulekeeper.NewKeeper(runtime.NewKVStoreService(keys[nftmoduletypes.StoreKey]), appCodec, app.AccountKeeper, app.BankKeeper)

	mSpvKeeper := spvmodulekeeper.NewKeeper(appCodec, keys[spvmoduletypes.StoreKey], keys[spvmoduletypes.MemStoreKey], spvSubspace, app.kycKeeper, app.BankKeeper, app.AccountKeeper, app.nftKeeper, app.kavaPricefeedKeeper, app.auctionKeeper, app.incentiveKeeper)

	app.incentiveKeeper = incentivekeeper.NewKeeper(
		appCodec,
		keys[incentivetypes.StoreKey],
		incentiveSubspace,
		app.BankKeeper,
		// &cdpKeeper,
		&joltKeeper,
		app.AccountKeeper,
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
		burnAuctionSubspace, app.AccountKeeper, app.BankKeeper, app.auctionKeeper)

	// dydx keepers

	app.EpochsKeeper = *epochsmodulekeeper.NewKeeper(
		appCodec,
		keys[epochsmoduletypes.StoreKey],
	)
	epochsModule := epochsmodule.NewAppModule(appCodec, app.EpochsKeeper)

	app.DydxPricesKeeper = *dydxpricesmodulekeeper.NewKeeper(
		appCodec,
		keys[dydxpricesmoduletypes.StoreKey],
		indexPriceCache,
		timeProvider,
		app.IndexerEventManager,
		// set the governance and delaymsg module accounts as the authority for conducting upgrades
		[]string{
			lib.GovModuleAddress.String(),
			delaymsgmoduletypes.ModuleAddress.String(),
		},
	)
	pricesModule := dydxpricesmodule.NewAppModule(appCodec, app.DydxPricesKeeper, app.AccountKeeper, app.BankKeeper)

	app.AssetsKeeper = *assetsmodulekeeper.NewKeeper(
		appCodec,
		keys[assetsmoduletypes.StoreKey],
		app.DydxPricesKeeper,
		app.IndexerEventManager,
	)
	assetsModule := assetsmodule.NewAppModule(appCodec, app.AssetsKeeper)

	app.DelayMsgKeeper = *delaymsgmodulekeeper.NewKeeper(
		appCodec,
		keys[delaymsgmoduletypes.StoreKey],
		bApp.MsgServiceRouter(),
		// Permit delayed messages to be signed by the following modules.
		[]string{
			lib.GovModuleAddress.String(),
		},
	)
	delayMsgModule := delaymsgmodule.NewAppModule(appCodec, app.DelayMsgKeeper)

	app.BridgeKeeper = *bridgemodulekeeper.NewKeeper(
		appCodec,
		keys[bridgemoduletypes.StoreKey],
		bridgeEventManager,
		app.BankKeeper,
		app.DelayMsgKeeper,
		// gov module and delayMsg module accounts are allowed to send messages to the bridge module.
		[]string{
			lib.GovModuleAddress.String(),
			delaymsgmoduletypes.ModuleAddress.String(),
		},
	)
	bridgeModule := bridgemodule.NewAppModule(appCodec, app.BridgeKeeper)

	app.PerpetualsKeeper = perpetualsmodulekeeper.NewKeeper(
		appCodec,
		keys[perpetualsmoduletypes.StoreKey],
		app.DydxPricesKeeper,
		app.EpochsKeeper,
		app.IndexerEventManager,
		// gov module and delayMsg module accounts are allowed to send messages to the bridge module.
		[]string{
			lib.GovModuleAddress.String(),
			delaymsgmoduletypes.ModuleAddress.String(),
		},
		tkeys[perpetualsmoduletypes.TransientStoreKey],
	)
	perpetualsModule := perpetualsmodule.NewAppModule(appCodec, app.PerpetualsKeeper)

	app.StatsKeeper = *statsmodulekeeper.NewKeeper(
		appCodec,
		app.EpochsKeeper,
		keys[statsmoduletypes.StoreKey],
		tkeys[statsmoduletypes.TransientStoreKey],
		// set the governance and delaymsg module accounts as the authority for conducting upgrades
		[]string{
			lib.GovModuleAddress.String(),
			delaymsgmoduletypes.ModuleAddress.String(),
		},
	)
	statsModule := statsmodule.NewAppModule(appCodec, app.StatsKeeper)

	app.FeeTiersKeeper = feetiersmodulekeeper.NewKeeper(
		appCodec,
		app.StatsKeeper,
		keys[feetiersmoduletypes.StoreKey],
		// set the governance and delaymsg module accounts as the authority for conducting upgrades
		[]string{
			lib.GovModuleAddress.String(),
			delaymsgmoduletypes.ModuleAddress.String(),
		},
	)
	feeTiersModule := feetiersmodule.NewAppModule(appCodec, app.FeeTiersKeeper)

	//app.VestKeeper = *vestmodulekeeper.NewKeeper(
	//	appCodec,
	//	keys[vestmoduletypes.StoreKey],
	//	app.BankKeeper,
	//	app.BlockTimeKeeper,
	//	// set the governance and delaymsg module accounts as the authority for conducting upgrades
	//	[]string{
	//		lib.GovModuleAddress.String(),
	//		delaymsgmoduletypes.ModuleAddress.String(),
	//	},
	//)
	//vestModule := vestmodule.NewAppModule(appCodec, app.VestKeeper)

	app.RewardsKeeper = *rewardsmodulekeeper.NewKeeper(
		appCodec,
		keys[rewardsmoduletypes.StoreKey],
		tkeys[rewardsmoduletypes.TransientStoreKey],
		app.AssetsKeeper,
		app.BankKeeper,
		app.FeeTiersKeeper,
		app.DydxPricesKeeper,
		app.IndexerEventManager,
		// set the governance and delaymsg module accounts as the authority for conducting upgrades
		[]string{
			lib.GovModuleAddress.String(),
			delaymsgmoduletypes.ModuleAddress.String(),
		},
	)
	rewardsModule := rewardsmodule.NewAppModule(appCodec, app.RewardsKeeper)

	app.SubaccountsKeeper = *subaccountsmodulekeeper.NewKeeper(
		appCodec,
		keys[satypes.StoreKey],
		app.AssetsKeeper,
		app.BankKeeper,
		app.PerpetualsKeeper,
		app.BlockTimeKeeper,
		app.IndexerEventManager,
	)
	subaccountsModule := subaccountsmodule.NewAppModule(
		appCodec,
		app.SubaccountsKeeper,
	)

	clobFlags := clobflags.GetClobFlagValuesFromOptions(appOpts)
	logger.Info("Parsed CLOB flags", "Flags", clobFlags)

	memClob := clobmodulememclob.NewMemClobPriceTimePriority(app.IndexerEventManager.Enabled())
	memClob.SetGenerateOrderbookUpdates(app.GrpcStreamingManager.Enabled())

	app.ClobKeeper = clobmodulekeeper.NewKeeper(
		appCodec,
		keys[clobmoduletypes.StoreKey],
		memKeys[clobmoduletypes.MemStoreKey],
		tkeys[clobmoduletypes.TransientStoreKey],
		// set the governance and delaymsg module accounts as the authority for conducting upgrades
		[]string{
			lib.GovModuleAddress.String(),
			delaymsgmoduletypes.ModuleAddress.String(),
		},
		memClob,
		app.SubaccountsKeeper,
		app.AssetsKeeper,
		app.BlockTimeKeeper,
		app.BankKeeper,
		app.FeeTiersKeeper,
		app.PerpetualsKeeper,
		app.DydxPricesKeeper,
		app.StatsKeeper,
		app.RewardsKeeper,
		app.IndexerEventManager,
		app.GrpcStreamingManager,
		txConfig.TxDecoder(),
		clobFlags,
		rate_limit.NewPanicRateLimiter[sdk.Msg](),
		daemonLiquidationInfo,
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
		app.AccountKeeper,
		app.BankKeeper,
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
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper))
	//	AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(app.ibcKeeper.ClientKeeper))

	govKeeper.SetLegacyRouter(govRouter)

	app.GovKeeper = *govKeeper.SetHooks(
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
	authModule := auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts, app.ParamsKeeper.Subspace(authtypes.ModuleName))
	app.ModuleManager = module.NewManager(
		genutil.NewAppModule(app.AccountKeeper, app.stakingKeeper, app, encodingConfig.TxConfig),
		authModule,
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper, app.ParamsKeeper.Subspace(banktypes.ModuleName)),
		capability.NewAppModule(appCodec, *app.capabilityKeeper, false),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.feeGrantKeeper, app.interfaceRegistry),
		staking.NewAppModule(appCodec, app.stakingKeeper, app.AccountKeeper, app.BankKeeper, app.ParamsKeeper.Subspace(stakingtypes.ModuleName)),
		mint.NewAppModule(appCodec, app.mintKeeper, app.AccountKeeper, app.BankKeeper),
		distr.NewAppModule(appCodec, app.distrKeeper, app.AccountKeeper, app.BankKeeper, app.stakingKeeper, app.ParamsKeeper.Subspace(distrtypes.ModuleName)),
		gov.NewAppModule(appCodec, &app.GovKeeper, app.AccountKeeper, app.BankKeeper, app.ParamsKeeper.Subspace(govtypes.ModuleName)),
		params.NewAppModule(app.ParamsKeeper),
		slashing.NewAppModule(appCodec, app.slashingKeeper, app.AccountKeeper, app.BankKeeper, app.stakingKeeper, app.ParamsKeeper.Subspace(slashingtypes.ModuleName), app.interfaceRegistry),
		ibc.NewAppModule(app.ibcKeeper),
		upgrade.NewAppModule(app.upgradeKeeper, addresscodec.NewBech32Codec(sdk.Bech32PrefixAccAddr)),
		evidence.NewAppModule(*app.evidenceKeeper),
		transferModule,
		vesting.NewAppModule(app.AccountKeeper, app.BankKeeper),
		authzmodule.NewAppModule(appCodec, app.authzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		auction.NewAppModule(app.auctionKeeper, app.AccountKeeper, app.BankKeeper),
		// issuance.NewAppModule(app.issuanceKeeper, app.AccountKeeper, app.bankKeeper),
		kavapricefeed.NewAppModule(app.kavaPricefeedKeeper, app.AccountKeeper),
		// cdp.NewAppModule(app.cdpKeeper, app.AccountKeeper, app.pricefeedKeeper, app.bankKeeper),
		jolt.NewAppModule(app.joltKeeper, app.AccountKeeper, app.BankKeeper, app.kavaPricefeedKeeper),
		incentive.NewAppModule(app.incentiveKeeper, app.AccountKeeper, app.BankKeeper),
		kycmodule.NewAppModule(appCodec, app.kycKeeper, app.AccountKeeper, app.BankKeeper),

		spvmodule.NewAppModule(appCodec, app.spvKeeper, app.AccountKeeper, app.BankKeeper),
		nftmodule.NewAppModule(appCodec, app.nftKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		ibcratelimit.NewAppModule(*app.RateLimitingICS4Wrapper),
		quotamodule.NewAppModule(appCodec, app.QuotaKeeper),
		swap.NewAppModule(app.swapKeeper, app.AccountKeeper),
		crisis.NewAppModule(app.crisisKeeper, skipGenesisInvariants, app.ParamsKeeper.Subspace(crisistypes.ModuleName)), // always be last to make sure that it checks for all invariants and not only part of them
		burnauctionmodule.NewAppModule(appCodec, app.burnauctionKeeper, app.AccountKeeper, app.BankKeeper),

		// dydx
		clobmodule.NewAppModule(appCodec, app.ClobKeeper, app.AccountKeeper, app.BankKeeper, app.SubaccountsKeeper),
		rewardsModule,
		subaccountsModule,
		assetsModule,
		bridgeModule,
		perpetualsModule,
		statsModule,
		feeTiersModule,
		// vestModule,
		delayMsgModule,
		epochsModule,
		pricesModule,
	)

	// BasicModuleManager defines the module BasicManager which is in charge of setting up basic,
	// non-dependant module elements, such as codec registration and genesis verification.
	// By default, it is composed of all the modules from the module manager.
	// Additionally, app module basics can be overwritten by passing them as an argument.
	app.BasicModuleManager = module.NewBasicManagerFromManager(
		app.ModuleManager,
		map[string]module.AppModuleBasic{
			genutiltypes.ModuleName: genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
			govtypes.ModuleName: gov.NewAppModuleBasic(
				[]govclient.ProposalHandler{
					paramsclient.ProposalHandler,
				},
			),
		},
	)
	// app.BasicModuleManager.RegisterLegacyAminoCodec(encodingConfig.Amino)
	app.BasicModuleManager.RegisterInterfaces(interfaceRegistry)

	app.ModuleManager.SetOrderPreBlockers(preBlockers...)

	// Warning: Some begin blockers must run before others. Ensure the dependencies are understood before modifying this list.
	app.ModuleManager.SetOrderBeginBlockers(

		blocktimemoduletypes.ModuleName, // Must be first
		authz.ModuleName,
		epochsmoduletypes.ModuleName,
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
		kavapricefeedtypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		govtypes.ModuleName,
		crisistypes.ModuleName,
		genutiltypes.ModuleName,
		quotamoduletypes.ModuleName,
		ibctransfertypes.ModuleName,
		paramstypes.ModuleName,
		burnauctionmoduletypes.ModuleName,
		// dydx

		dydxpricesmoduletypes.ModuleName,
		assetsmoduletypes.ModuleName,
		bridgemoduletypes.ModuleName,
		feetiersmoduletypes.ModuleName,
		perpetualsmoduletypes.ModuleName,
		statsmoduletypes.ModuleName,
		satypes.ModuleName,
		clobmoduletypes.ModuleName,
		// vestmoduletypes.ModuleName,
		rewardsmoduletypes.ModuleName,
		sendingmoduletypes.ModuleName,
		govplusmoduletypes.ModuleName,
		delaymsgmoduletypes.ModuleName,
		vaultmoduletypes.ModuleName,
	)

	app.ModuleManager.SetOrderPrepareCheckStaters(
		clobmoduletypes.ModuleName,
	)

	// Warning: Some end blockers must run before others. Ensure the dependencies are understood before modifying this list.
	app.ModuleManager.SetOrderEndBlockers(
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		kavapricefeedtypes.ModuleName,
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
		burnauctionmoduletypes.ModuleName,

		// dydx

		dydxpricesmoduletypes.ModuleName,
		assetsmoduletypes.ModuleName,
		bridgemoduletypes.ModuleName,
		feetiersmoduletypes.ModuleName,
		perpetualsmoduletypes.ModuleName,
		statsmoduletypes.ModuleName,
		satypes.ModuleName,
		clobmoduletypes.ModuleName,
		sendingmoduletypes.ModuleName,
		// vestmoduletypes.ModuleName,
		rewardsmoduletypes.ModuleName,
		epochsmoduletypes.ModuleName,
		govplusmoduletypes.ModuleName,
		delaymsgmoduletypes.ModuleName,
		// Vault endblocker should be after clob endblocker. Otherwise,
		// the block after the one where vault orders expire won't have
		// any vault orders.
		vaultmoduletypes.ModuleName,
		authz.ModuleName,                // No-op.
		blocktimemoduletypes.ModuleName, // Must be last

	)

	// Warning: Some init genesis methods must run before others. Ensure the dependencies are understood before modifying this list
	app.ModuleManager.SetOrderInitGenesis(
		epochsmoduletypes.ModuleName,
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
		ibctransfertypes.ModuleName,
		feegrant.ModuleName,
		auctiontypes.ModuleName,
		// issuancetypes.ModuleName,
		kavapricefeedtypes.ModuleName,
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

		// dydx
		dydxpricesmoduletypes.ModuleName,
		assetsmoduletypes.ModuleName,
		blocktimemoduletypes.ModuleName,
		bridgemoduletypes.ModuleName,
		feetiersmoduletypes.ModuleName,
		perpetualsmoduletypes.ModuleName,
		statsmoduletypes.ModuleName,
		satypes.ModuleName,
		clobmoduletypes.ModuleName,
		// vestmoduletypes.ModuleName,
		rewardsmoduletypes.ModuleName,
		sendingmoduletypes.ModuleName,
		govplusmoduletypes.ModuleName,
		delaymsgmoduletypes.ModuleName,
		vaultmoduletypes.ModuleName,
		authz.ModuleName,
	)

	app.ModuleManager.RegisterInvariants(app.crisisKeeper)
	app.configurator = module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	err := app.ModuleManager.RegisterServices(app.configurator)
	if err != nil {
		panic(err)
	}

	autocliv1.RegisterQueryServer(app.GRPCQueryRouter(), runtimeservices.NewAutoCLIQueryService(app.ModuleManager.Modules))
	reflectionSvc, err := runtimeservices.NewReflectionService()
	if err != nil {
		panic(err)
	}
	reflectionv1.RegisterReflectionServiceServer(app.GRPCQueryRouter(), reflectionSvc)

	// create the simulation manager and define the order of the modules for deterministic simulations
	overrideModules := map[string]module.AppModuleSimulation{
		authtypes.ModuleName: authModule,
	}
	app.sm = module.NewSimulationManagerFromAppModules(app.ModuleManager.Modules, overrideModules)
	app.sm.RegisterStoreDecoders()

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: This is not required for apps that don't use the simulator for fuzz testing
	// transactions.
	// TODO
	// app.sm = module.NewSimulationManager(
	// 	auth.NewAppModule(app.AccountKeeper),
	// 	bank.NewAppModule(app.bankKeeper, app.AccountKeeper),
	// 	gov.NewAppModule(app.GovKeeper, app.AccountKeeper, app.AccountKeeper, app.bankKeeper),
	// 	mint.NewAppModule(app.mintKeeper),
	// 	distr.NewAppModule(app.distrKeeper, app.AccountKeeper, app.AccountKeeper, app.bankKeeper, app.stakingKeeper),
	//  staking.NewAppModule(app.stakingKeeper, app.AccountKeeper, app.AccountKeeper, app.bankKeeper),
	// 	slashing.NewAppModule(app.slashingKeeper, app.AccountKeeper, app.stakingKeeper),
	// )
	// app.sm.RegisterStoreDecoders()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	// baseAnte := cosante.HandlerOptions{
	//	AccountKeeper:   app.AccountKeeper,
	//	BankKeeper:      app.bankKeeper,
	//	SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
	//	FeegrantKeeper:  app.feeGrantKeeper,
	//	SigGasConsumer:  cosante.DefaultSigVerificationGasConsumer,
	//}

	extensionCheck := func(a *types.Any) bool {
		// todo we need to verify here, currently, we allow all the tx to be passed
		return true
	}
	//
	anteOptions := jante.HandlerOptions{
		AccountKeeper:          &app.AccountKeeper,
		BankKeeper:             app.BankKeeper,
		SignModeHandler:        encodingConfig.TxConfig.SignModeHandler(),
		FeegrantKeeper:         app.feeGrantKeeper,
		SpvKeeper:              app.spvKeeper,
		IBCKeeper:              app.ibcKeeper,
		AddressFetchers:        []jante.AddressFetcher{},
		ExtensionOptionChecker: extensionCheck,
		TxFeeChecker:           nil,
	}

	anteHandlerold, err := jante.NewAnteHandler(anteOptions, app.ConsensusParamsKeeper)
	if err != nil {
		panic(fmt.Sprintf("failed to create anteHandler: %s", err))
	}
	_ = anteHandlerold

	// app.SetAnteHandler(anteHandler)
	app.setupUpgradeHandlers()

	app.SetInitChainer(app.InitChainer)
	app.SetMempool(mempool.NewNoOpMempool())
	app.setAnteHandler(encodingConfig.TxConfig)

	app.SetPreBlocker(app.PreBlocker)

	app.SetBeginBlocker(app.BeginBlockerDydx)
	app.SetEndBlocker(app.EndBlockerDydx)
	app.SetPrecommiter(app.PrecommitterDydx)
	app.SetPrepareCheckStater(app.PrepareCheckStaterDydx)

	// fixme we disable it here
	// ProposalHandler setup.
	// prepareProposalHandler, processProposalHandler := app.createProposalHandlers(appFlags, txConfig, appOpts)
	// app.SetPrepareProposal(prepareProposalHandler)
	// app.SetProcessProposal(processProposalHandler)

	// app.SetBeginBlocker(app.BeginBlocker)
	// app.SetEndBlocker(app.EndBlocker)
	// app.SetPreBlocker(app.PreBlocker)

	app.ScopedIBCKeeper = scopedIBCKeeper
	app.ScopedTransferKeeper = scopedTransferKeeper

	// At startup, after all modules have been registered, check that all prot
	// annotations are correct.
	// fixme we ignore the dydx proto
	protoFiles, err := proto.MergedRegistry()
	if err != nil {
		if strings.Contains(err.Error(), "dydxprotocol/bridge") || strings.Contains(err.Error(), "subaccounts") {
		} else {
			panic(err)
		}
	}
	err = msgservice.ValidateProtoAnnotations(protoFiles)
	if err != nil {
		// Once we switch to using protoreflect-based antehandlers, we might
		// want to panic here instead of logging a warning.
		fmt.Fprintln(os.Stderr, err.Error())
	}

	// load store
	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}
	}

	app.initializeRateLimiters()

	return app
}

// Name returns the name of the App
func (app *App) Name() string { return app.BaseApp.Name() }

// BeginBlocker contains app specific logic for the BeginBlock abci call.
func (app *App) BeginBlocker(ctx sdk.Context) (sdk.BeginBlock, error) {
	return app.ModuleManager.BeginBlock(ctx)
}

// EndBlocker contains app specific logic for the EndBlock abci call.
func (app *App) EndBlocker(ctx sdk.Context) (sdk.EndBlock, error) {
	ret, err := app.ModuleManager.EndBlock(ctx)

	block := app.IndexerEventManager.ProduceBlock(ctx)
	app.IndexerEventManager.SendOnchainData(block)
	return ret, err
}

// BeginBlockerDydx application updates every begin block
func (app *App) BeginBlockerDydx(ctx sdk.Context) (sdk.BeginBlock, error) {
	ctx = ctx.WithExecMode(lib.ExecModeBeginBlock)

	// Update the proposer address in the logger for the panic logging middleware.
	proposerAddr := sdk.ConsAddress(ctx.BlockHeader().ProposerAddress)
	middleware.Logger = ctx.Logger().With("proposer_cons_addr", proposerAddr.String())

	// app.scheduleForkUpgrade(ctx)
	return app.ModuleManager.BeginBlock(ctx)
}

// EndBlockerDydx application updates every end block
func (app *App) EndBlockerDydx(ctx sdk.Context) (sdk.EndBlock, error) {
	// Measure the lag between current timestamp and the end blocker time stamp
	// as an indicator of whether the node is lagging behind.
	metrics.ModuleMeasureSince(metrics.EndBlocker, metrics.EndBlockerLag, ctx.BlockTime())

	ctx = ctx.WithExecMode(lib.ExecModeEndBlock)

	// Reset the logger for middleware.
	// Note that the middleware is only used by `CheckTx` and `DeliverTx`, and not `EndBlocker`.
	// Panics from `EndBlocker` will not be logged by the middleware and will lead to consensus failures.
	middleware.Logger = app.Logger()

	response, err := app.ModuleManager.EndBlock(ctx)
	if err != nil {
		return response, err
	}
	block := app.IndexerEventManager.ProduceBlock(ctx)
	app.IndexerEventManager.SendOnchainData(block)
	return response, err
}

// PrecommitterDydx application updates before the commital of a block after all transactions have been delivered.
func (app *App) PrecommitterDydx(ctx sdk.Context) {
	if err := app.ModuleManager.Precommit(ctx); err != nil {
		panic(err)
	}
}

// PrepareCheckStaterDydx application updates after commit and before any check state is invoked.
func (app *App) PrepareCheckStaterDydx(ctx sdk.Context) {
	ctx = ctx.WithExecMode(lib.ExecModePrepareCheckState)

	if err := app.ModuleManager.PrepareCheckState(ctx); err != nil {
		panic(err)
	}
}

// PreBlocker application updates every pre block
func (app *App) PreBlocker(ctx sdk.Context, _ *abci.RequestFinalizeBlock) (*sdk.ResponsePreBlock, error) {
	ctx = ctx.WithGasMeter(types2.NewFreeInfiniteGasMeter())
	return app.ModuleManager.PreBlock(ctx)
}

// InitChainer contains app specific logic for the InitChain abci call.
func (app *App) InitChainer(ctx sdk.Context, req *abci.RequestInitChain) (*abci.ResponseInitChain, error) {
	var genesisState GenesisState
	if err := tmjson.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}

	// Store current module versions in joltify-10 to setup future in-place upgrades.
	// During in-place migrations, the old module versions in the store will be referenced to determine which migrations to run.
	app.upgradeKeeper.SetModuleVersionMap(ctx, app.ModuleManager.GetVersionMap())

	return app.ModuleManager.InitGenesis(ctx, app.appCodec, genesisState)
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

	app.BasicModuleManager.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

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
//	joltIncentiveMaccAddr := app.AccountKeeper.GetModuleAddress(incentivetypes.ModuleName)
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

	modules := map[string]appmodule.AppModule{
		ibcexported.ModuleName:     ibc.AppModule{},
		capabilitytypes.ModuleName: capability.AppModule{},
		ibctm.ModuleName:           ibctm.AppModule{},
		solomachine.ModuleName:     solomachine.AppModule{},
	}

	for name, m := range modules {
		module.CoreAppModuleBasicAdaptor(name, m).RegisterInterfaces(app.interfaceRegistry)
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
			keyTable = stakingtypes.ParamKeyTable() //nolint:staticcheck
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
		case ibctransfertypes.ModuleName:
			keyTable = ibctransfertypes.ParamKeyTable()

		}
		if !subspace.HasKeyTable() {
			subspace.WithKeyTable(keyTable)
		}
	}

	app.upgradeKeeper.SetUpgradeHandler(v1.V011UpgradeName, v1.CreateUpgradeHandlerForV011Upgrade(app.ModuleManager, app.configurator, app.kycKeeper, app.spvKeeper, app.QuotaKeeper, app.incentiveKeeper))
	app.upgradeKeeper.SetUpgradeHandler(v1.V012UpgradeName, v1.CreateUpgradeHandlerForV012Upgrade(app.ModuleManager, app.configurator))
	app.upgradeKeeper.SetUpgradeHandler(v1.V013UpgradeName, v1.CreateUpgradeHandlerForV013Upgrade(app.ModuleManager, app.configurator))
	app.upgradeKeeper.SetUpgradeHandler(v1.V014UpgradeName, v1.CreateUpgradeHandlerForV014Upgrade(app.ModuleManager, app.configurator))
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

func (app *App) EncodingConfig() config2.EncodingConfig {
	return config2.EncodingConfig{
		InterfaceRegistry: app.InterfaceRegistry(),
		Marshaler:         app.AppCodec(),
		TxConfig:          app.TxConfig(),
		Amino:             app.LegacyAmino(),
	}
}

// AutoCliOpts returns the autocli options for the app.
func (app *App) AutoCliOpts() autocli.AppOptions {
	modules := make(map[string]appmodule.AppModule, 0)
	for _, m := range app.ModuleManager.Modules {
		if moduleWithName, ok := m.(module.HasName); ok {
			moduleName := moduleWithName.Name()
			if appModule, ok := moduleWithName.(appmodule.AppModule); ok {
				modules[moduleName] = appModule
			}
		}
	}
	return autocli.AppOptions{
		Modules:               modules,
		ModuleOptions:         runtimeservices.ExtractAutoCLIOptions(app.ModuleManager.Modules),
		AddressCodec:          authcodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
		ValidatorAddressCodec: authcodec.NewBech32Codec(sdk.GetConfig().GetBech32ValidatorAddrPrefix()),
		ConsensusAddressCodec: authcodec.NewBech32Codec(sdk.GetConfig().GetBech32ConsensusAddrPrefix()),
	}
}

// getIndexerFromOptions returns an instance of a msgsender.IndexerMessageSender from the specified options.
// This function will default to try to use any instance that is configured for test execution followed by loading
// an instance from command line flags and finally returning a no-op instance.
func getIndexerFromOptions(
	appOpts servertypes.AppOptions,
	logger log.Logger,
) (msgsender.IndexerMessageSender, indexer.IndexerFlags) {
	v, ok := appOpts.Get(indexer.MsgSenderInstanceForTest).(msgsender.IndexerMessageSender)
	if ok {
		return v, indexer.IndexerFlags{
			SendOffchainData: true,
		}
	}

	indexerFlags := indexer.GetIndexerFlagValuesFromOptions(appOpts)
	logger.Info(
		"Parsed Indexer flags",
		"Flags", indexerFlags,
	)

	var indexerMessageSender msgsender.IndexerMessageSender
	if len(indexerFlags.KafkaAddrs) == 0 {
		indexerMessageSender = msgsender.NewIndexerMessageSenderNoop()
	} else {
		var err error
		indexerMessageSender, err = msgsender.NewIndexerMessageSenderKafka(
			indexerFlags,
			nil,
			logger,
		)
		if err != nil {
			panic(err)
		}
	}
	return indexerMessageSender, indexerFlags
}

// getGrpcStreamingManagerFromOptions returns an instance of a streamingtypes.GrpcStreamingManager from the specified
// options. This function will default to returning a no-op instance.
func getGrpcStreamingManagerFromOptions(
	appFlags appFlag.Flags,
	logger log.Logger,
) (manager streamingtypes.GrpcStreamingManager) {
	if appFlags.GrpcStreamingEnabled {
		logger.Info("GRPC streaming is enabled")
		return streaming.NewGrpcStreamingManager(
			logger,
			appFlags.GrpcStreamingFlushIntervalMs,
			appFlags.GrpcStreamingMaxBatchSize,
			appFlags.GrpcStreamingMaxChannelBufferSize,
		)
	}
	return streaming.NewNoopGrpcStreamingManager()
}

// initializeRateLimiters initializes the rate limiters from state if the application is
// not started from genesis.
func (app *App) initializeRateLimiters() {
	// Create an `uncachedCtx` where the underlying MultiStore is the `rootMultiStore`.
	// We use this to hydrate the `orderRateLimiter` with values from the underlying `rootMultiStore`.
	uncachedCtx := app.BaseApp.NewUncachedContext(true, tmproto.Header{})
	app.ClobKeeper.InitalizeBlockRateLimitFromStateIfExists(uncachedCtx)
}

// buildAnteHandler builds an AnteHandler object configured for the app.
func (app *App) buildAnteHandler(txConfig client.TxConfig) sdk.AnteHandler {
	anteHandler, err := dydxante.NewAnteHandler(
		dydxante.HandlerOptions{
			HandlerOptions: ante.HandlerOptions{
				AccountKeeper:   app.AccountKeeper,
				BankKeeper:      app.BankKeeper,
				SignModeHandler: txConfig.SignModeHandler(),
				FeegrantKeeper:  app.feeGrantKeeper,
				SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
			},
			ClobKeeper:   app.ClobKeeper,
			Codec:        app.appCodec,
			AuthStoreKey: app.keys[authtypes.StoreKey],
		},
	)
	if err != nil {
		panic(err)
	}

	return anteHandler
}

// setAnteHandler creates a new AnteHandler and sets it on the base app and clob keeper.
func (app *App) setAnteHandler(txConfig client.TxConfig) {
	anteHandler := app.buildAnteHandler(txConfig)
	// Prevent a cycle between when we create the clob keeper and the ante handler.
	app.ClobKeeper.SetAnteHandler(anteHandler)
	app.SetAnteHandler(anteHandler)
}
