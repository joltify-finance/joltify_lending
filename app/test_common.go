package app

import (
	crand "crypto/rand"
	"encoding/json"
	"math/rand"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	distkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	dbm "github.com/cosmos/cosmos-db"

	pruningtypes "cosmossdk.io/store/pruning/types"

	storetypes "cosmossdk.io/store/types"

	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"

	swapkeeper "github.com/joltify-finance/joltify_lending/x/third_party/swap/keeper"

	sdkmath "cosmossdk.io/math"

	"cosmossdk.io/log"
	tmjson "github.com/cometbft/cometbft/libs/json"

	abci "github.com/cometbft/cometbft/abci/types"
	mintkeeper "github.com/joltify-finance/joltify_lending/x/mint/keeper"
	minttypes "github.com/joltify-finance/joltify_lending/x/mint/types"
	auctionkeeper "github.com/joltify-finance/joltify_lending/x/third_party/auction/keeper"
	incentivekeeper "github.com/joltify-finance/joltify_lending/x/third_party/incentive/keeper"
	joltkeeper "github.com/joltify-finance/joltify_lending/x/third_party/jolt/keeper"
	pricefeedkeeper "github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/keeper"
)

var (
	emptyTime            time.Time
	testChainID                = "joltifytest_888-1"
	defaultInitialHeight int64 = 1
)

// TestApp is a simple wrapper around an App. It exposes internal keepers for use in integration tests.
// This file also contains test helpers. Ideally they would be in separate package.
// Basic Usage:
// Create a test app with NewTestApp, then all keepers and their methods can be accessed for test setup and execution.
// Advanced Usage:
// Some tests call for an app to be initialized with some state. This can be achieved through keeper method calls (ie keeper.SetParams(...)).
// However this leads to a lot of duplicated logic similar to InitGenesis methods.
// So TestApp.InitializeFromGenesisStates() will call InitGenesis with the default genesis state.
// and TestApp.InitializeFromGenesisStates(authState, cdpState) will do the same but overwrite the auth and cdp sections of the default genesis state
// Creating the genesis states can be combersome, but helper methods can make it easier such as NewAuthGenStateFromAccounts below.
type TestApp struct {
	App
	Ctx sdk.Context
}

// NewTestApp creates a new TestApp
//
// Note, it also sets the sdk config with the app's address prefix, coin type, etc.
func NewTestApp(logger log.Logger, rootDir string) TestApp {
	SetSDKConfig()
	return NewTestAppFromSealed(logger, rootDir)
}

func genesisStateWithValSet(
	app *App, genesisState GenesisState,
	valSet []*stakingtypes.Validator, genAccs []authtypes.GenesisAccount,
	balances ...banktypes.Balance,
) GenesisState {
	// set genesis accounts
	// authGenesis := authtypes.NewGenesisState(authtypes.DefaultParams(), genAccs)
	// genesisState[authtypes.ModuleName] = app.AppCodec().MustMarshalJSON(authGenesis)

	validators := make([]stakingtypes.Validator, 0, len(valSet))
	delegations := make([]stakingtypes.Delegation, 0, len(valSet))

	bondAmt := sdk.DefaultPowerReduction.Mul(sdkmath.NewInt(1000000))

	for _, val := range valSet {
		//pk, err := cryptocodec.FromTmPubKeyInterface(val.PubKey)
		//if err != nil {
		//	panic(err)
		//}
		//pkAny, err := codectypes.NewAnyWithValue(pk)
		//if err != nil {
		//	panic(err)
		//}
		//validator := stakingtypes.Validator{
		//	OperatorAddress:   sdk.ValAddress(val.Address).String(),
		//	ConsensusPubkey:   pkAny,
		//	Jailed:            false,
		//	Status:            stakingtypes.Bonded,
		//	Tokens:            bondAmt,
		//	DelegatorShares:   sdkmath.LegacyOneDec(),
		//	Description:       stakingtypes.Description{},
		//	UnbondingHeight:   int64(0),
		//	UnbondingTime:     time.Unix(0, 0).UTC(),
		//	Commission:        stakingtypes.NewCommission(sdkmath.LegacyZeroDec(), sdkmath.LegacyZeroDec(), sdkmath.LegacyZeroDec()),
		//	MinSelfDelegation: sdkmath.ZeroInt(),
		//}

		val.DelegatorShares = sdkmath.LegacyOneDec()
		val.Tokens = bondAmt
		val.Status = stakingtypes.Bonded
		validators = append(validators, *val)
		delegations = append(delegations, stakingtypes.NewDelegation(genAccs[0].GetAddress().String(), val.GetOperator(), sdkmath.LegacyOneDec()))

	}
	// set validators and delegations
	stakingGenesis := stakingtypes.NewGenesisState(stakingtypes.DefaultParams(), validators, delegations)
	genesisState[stakingtypes.ModuleName] = app.AppCodec().MustMarshalJSON(stakingGenesis)

	totalSupply := sdk.NewCoins()
	for _, b := range balances {
		// add genesis acc tokens to total supply
		totalSupply = totalSupply.Add(b.Coins...)
	}

	for range delegations {
		// add delegated tokens to total supply
		totalSupply = totalSupply.Add(sdk.NewCoin(sdk.DefaultBondDenom, bondAmt))
	}

	// add bonded amount to bonded pool module account
	balances = append(balances, banktypes.Balance{
		Address: authtypes.NewModuleAddress(stakingtypes.BondedPoolName).String(),
		Coins:   sdk.Coins{sdk.NewCoin(sdk.DefaultBondDenom, bondAmt)},
	})

	// update total supply
	bankGenesis := banktypes.NewGenesisState(banktypes.DefaultGenesisState().Params, balances, totalSupply, []banktypes.Metadata{}, []banktypes.SendEnabled{})
	genesisState[banktypes.ModuleName] = app.AppCodec().MustMarshalJSON(bankGenesis)

	return genesisState
}

var DefaultConsensusParams = simtestutil.DefaultConsensusParams

// NewTestAppFromSealed creates a TestApp without first setting sdk config.
func NewTestAppFromSealed(logger log.Logger, rootDir string) TestApp {
	app := NewApp(
		logger, dbm.NewMemDB(), nil,
		true,
		simtestutil.NewAppOptionsWithFlagHome(rootDir),
		baseapp.SetPruning(pruningtypes.NewPruningOptionsFromString(pruningtypes.PruningOptionDefault)),
		baseapp.SetMinGasPrices("0stake"),
		baseapp.SetChainID("joltifytest_888-1"),
	)

	encCfg := app.EncodingConfig()

	_, pubKey, addr := testdata.KeyTestPubAddr()
	valAddr := sdk.ValAddress(addr)
	val, err := stakingtypes.NewValidator(valAddr.String(), pubKey, stakingtypes.Description{Moniker: "test"})
	if err != nil {
		panic(err)
	}

	// generate genesis account
	senderPrivKey := secp256k1.GenPrivKey()
	acc := authtypes.NewBaseAccount(senderPrivKey.PubKey().Address().Bytes(), senderPrivKey.PubKey(), 0, 0)
	balance := banktypes.Balance{
		Address: acc.GetAddress().String(),
		Coins:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdkmath.NewInt(12300000000000000))),
	}

	genesisState := NewDefaultGenesisState(encCfg.Marshaler)
	genesisState = genesisStateWithValSet(app, genesisState, []*stakingtypes.Validator{&val}, []authtypes.GenesisAccount{acc}, balance)

	stateBytes, err := tmjson.MarshalIndent(genesisState, "", " ")
	if err != nil {
		panic(err)
	}

	// Initialize the chain
	app.InitChain(
		&abci.RequestInitChain{
			Validators:      []abci.ValidatorUpdate{},
			ConsensusParams: DefaultConsensusParams,
			AppStateBytes:   stateBytes,
			ChainId:         "joltifytest_888-1",
		},
	)

	ctx := app.BaseApp.NewContext(false)
	ctx = ctx.WithBlockGasMeter(storetypes.NewGasMeter(1000000000000000000))
	return TestApp{App: *app, Ctx: ctx}
}

func (tApp TestApp) GetAccountKeeper() authkeeper.AccountKeeper { return tApp.accountKeeper }
func (tApp TestApp) GetBankKeeper() bankkeeper.Keeper           { return tApp.bankKeeper }
func (tApp TestApp) GetStakingKeeper() stakingkeeper.Keeper     { return *tApp.stakingKeeper }
func (tApp TestApp) GetSlashingKeeper() slashingkeeper.Keeper   { return tApp.slashingKeeper }
func (tApp TestApp) GetMintKeeper() mintkeeper.Keeper           { return tApp.mintKeeper }
func (tApp TestApp) GetDistrKeeper() distkeeper.Keeper          { return tApp.distrKeeper }
func (tApp TestApp) GetGovKeeper() govkeeper.Keeper             { return tApp.govKeeper }
func (tApp TestApp) GetCrisisKeeper() crisiskeeper.Keeper       { return *tApp.crisisKeeper }
func (tApp TestApp) GetParamsKeeper() paramskeeper.Keeper       { return tApp.ParamsKeeper }

func (tApp TestApp) GetAuctionKeeper() auctionkeeper.Keeper     { return tApp.auctionKeeper }
func (tApp TestApp) GetPriceFeedKeeper() pricefeedkeeper.Keeper { return tApp.pricefeedKeeper }
func (tApp TestApp) GetJoltKeeper() joltkeeper.Keeper           { return tApp.joltKeeper }
func (tApp TestApp) GetIncentiveKeeper() incentivekeeper.Keeper { return tApp.incentiveKeeper }
func (tApp TestApp) GetSwapKeeper() swapkeeper.Keeper           { return tApp.swapKeeper }

// LegacyAmino returns the app's amino codec.
func (app *App) LegacyAmino() *codec.LegacyAmino {
	return app.legacyAmino
}

// AppCodec returns the app's app codec.
func (app *App) AppCodec() codec.Codec {
	return app.appCodec
}

func (app *App) TxConfig() client.TxConfig {
	return app.txConfig
}

// InitializeFromGenesisStates calls InitChain on the app using the provided genesis states.
// If any module genesis states are missing, defaults are used.
func (tApp TestApp) InitializeFromGenesisStates(genAccs []authtypes.GenesisAccount, coins sdk.Coins, genesisStates ...GenesisState) TestApp {
	return tApp.InitializeFromGenesisStatesWithTimeAndChainIDAndHeight(emptyTime, testChainID, defaultInitialHeight, genAccs, coins, genesisStates...)
}

// InitializeFromGenesisStatesWithTime calls InitChain on the app using the provided genesis states and time.
// If any module genesis states are missing, defaults are used.
func (tApp TestApp) InitializeFromGenesisStatesWithTime(genTime time.Time, genAccs []authtypes.GenesisAccount, coins sdk.Coins, genesisStates ...GenesisState) TestApp {
	t := tApp.InitializeFromGenesisStatesWithTimeAndChainIDAndHeight(genTime, testChainID, defaultInitialHeight, genAccs, coins, genesisStates...)
	return t
}

// InitializeFromGenesisStatesWithTimeAndChainID calls InitChain on the app using the provided genesis states, time, and chain id.
// If any module genesis states are missing, defaults are used.
func (tApp TestApp) InitializeFromGenesisStatesWithTimeAndChainID(genTime time.Time, chainID string,
	genAccs []authtypes.GenesisAccount, coins sdk.Coins, genesisStates ...GenesisState,
) TestApp {
	return tApp.InitializeFromGenesisStatesWithTimeAndChainIDAndHeight(genTime, chainID, defaultInitialHeight, genAccs, coins, genesisStates...)
}

// InitializeFromGenesisStatesWithTimeAndChainIDAndHeight calls InitChain on the app using the provided genesis states and other parameters.
// If any module genesis states are missing, defaults are used.
func (tApp TestApp) InitializeFromGenesisStatesWithTimeAndChainIDAndHeight(genTime time.Time, chainID string, initialHeight int64, genAccs []authtypes.GenesisAccount, coins sdk.Coins, genesisStates ...GenesisState) TestApp {
	// Create a default genesis state and overwrite with provided values
	encoding := MakeEncodingConfig()
	genesisState := NewDefaultGenesisState(encoding.Marshaler)
	for _, state := range genesisStates {
		for k, v := range state {
			genesisState[k] = v
		}
	}

	_, pubKey, addr := testdata.KeyTestPubAddr()
	valAddr := sdk.ValAddress(addr)
	val, err := stakingtypes.NewValidator(valAddr.String(), pubKey, stakingtypes.Description{Moniker: "test"})
	if err != nil {
		panic(err)
	}

	// create validator set with single validator
	defaultCoins := sdk.NewCoin(sdk.DefaultBondDenom, sdkmath.NewInt(100000000000000))
	coins = coins.Add(defaultCoins)
	var balances []banktypes.Balance
	if len(genAccs) == 0 {
		senderPrivKey := secp256k1.GenPrivKey()
		acc := authtypes.NewBaseAccount(senderPrivKey.PubKey().Address().Bytes(), senderPrivKey.PubKey(), 0, 0)
		balanceItem := banktypes.Balance{
			Address: acc.GetAddress().String(),
			Coins:   coins,
		}
		genAccs = []authtypes.GenesisAccount{acc}
		balances = []banktypes.Balance{balanceItem}
	} else {
		for _, el := range genAccs {
			balanceItem := banktypes.Balance{
				Address: el.GetAddress().String(),
				Coins:   coins,
			}
			balances = append(balances, balanceItem)
		}
	}

	genesisState = genesisStateWithValSet(&tApp.App, genesisState, []*stakingtypes.Validator{&val}, genAccs, balances...)
	// Initialize the chain
	stateBytes, err := json.Marshal(genesisState)
	if err != nil {
		panic(err)
	}

	_, err = tApp.InitChain(
		&abci.RequestInitChain{
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: stateBytes,
			ChainId:       chainID,
			// Set consensus params, which is needed by x/feemarket
			ConsensusParams: simtestutil.DefaultConsensusParams,
		},
	)
	if err != nil {
		panic(err)
	}

	_, err = tApp.Commit()
	if err != nil {
		panic(err)
	}

	return tApp
}

// RandomAddress non-deterministically generates a new address, discarding the private key.
func RandomAddress() (sdk.AccAddress, string) {
	secret := make([]byte, 32)
	_, err := crand.Read(secret)
	if err != nil {
		panic("Could not read randomness")
	}
	key := secp256k1.GenPrivKeyFromSecret(secret)
	return sdk.AccAddress(key.PubKey().Address()), common.Bytes2Hex(key.PubKey().Bytes())
}

// CheckBalance requires the account address has the expected amount of coins.
func (tApp TestApp) CheckBalance(t *testing.T, ctx sdk.Context, owner sdk.AccAddress, expectedCoins sdk.Coins) {
	coins := tApp.GetBankKeeper().GetAllBalances(ctx, owner)
	require.Equal(t, expectedCoins, coins)
}

// GetModuleAccountBalance gets the current balance of the denom for a module account
func (tApp TestApp) GetModuleAccountBalance(ctx sdk.Context, moduleName string, denom string) sdkmath.Int {
	moduleAcc := tApp.accountKeeper.GetModuleAccount(ctx, moduleName)
	balance := tApp.bankKeeper.GetBalance(ctx, moduleAcc.GetAddress(), denom)
	return balance.Amount
}

// FundAccount is a utility function that funds an account by minting and sending the coins to the address.
func (tApp TestApp) FundAccount(ctx sdk.Context, addr sdk.AccAddress, amounts sdk.Coins) error {
	if err := tApp.bankKeeper.MintCoins(ctx, minttypes.ModuleName, amounts); err != nil {
		return err
	}

	return tApp.bankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, addr, amounts)
}

// NewQueryServerTestHelper creates a new QueryServiceTestHelper that wraps the provided sdk.Context.
func (tApp TestApp) NewQueryServerTestHelper(ctx sdk.Context) *baseapp.QueryServiceTestHelper {
	return baseapp.NewQueryServerTestHelper(ctx, tApp.interfaceRegistry)
}

// FundModuleAccount is a utility function that funds a module account by minting and sending the coins to the address.
func (tApp TestApp) FundModuleAccount(ctx sdk.Context, recipientMod string, amounts sdk.Coins) error {
	if err := tApp.bankKeeper.MintCoins(ctx, minttypes.ModuleName, amounts); err != nil {
		return err
	}

	return tApp.bankKeeper.SendCoinsFromModuleToModule(ctx, minttypes.ModuleName, recipientMod, amounts)
}

// GeneratePrivKeyAddressPairs generates (deterministically) a total of n private keys and addresses.
func GeneratePrivKeyAddressPairs(n int) (keys []cryptotypes.PrivKey, addrs []sdk.AccAddress) {
	r := rand.New(rand.NewSource(12345)) // make the generation deterministic

	keys = make([]cryptotypes.PrivKey, n)
	addrs = make([]sdk.AccAddress, n)
	for i := 0; i < n; i++ {
		secret := make([]byte, 32)
		_, err := r.Read(secret)
		if err != nil {
			panic("Could not read randomness")
		}
		keys[i] = secp256k1.GenPrivKeyFromSecret(secret)
		addrs[i] = sdk.AccAddress(keys[i].PubKey().Address())
	}
	return
}

// NewFundedGenStateWithSameCoins creates a (auth and bank) genesis state populated with accounts from the given addresses and balance.
func NewFundedGenStateWithSameCoins(cdc codec.JSONCodec, balance sdk.Coins, addresses []sdk.AccAddress) GenesisState {
	builder := NewAuthBankGenesisBuilder()
	for _, address := range addresses {
		builder.WithSimpleAccount(address, balance)
	}
	return builder.BuildMarshalled(cdc)
}

// NewFundedGenStateWithCoins creates a (auth and bank) genesis state populated with accounts from the given addresses and coins.
func NewFundedGenStateWithCoins(cdc codec.JSONCodec, coins []sdk.Coins, addresses []sdk.AccAddress) GenesisState {
	builder := NewAuthBankGenesisBuilder()
	for i, address := range addresses {
		builder.WithSimpleAccount(address, coins[i])
	}
	return builder.BuildMarshalled(cdc)
}

// NewFundedGenStateWithSameCoinsWithModuleAccount creates a (auth and bank) genesis state populated with accounts from the given addresses and balance along with an empty module account
func NewFundedGenStateWithSameCoinsWithModuleAccount(cdc codec.JSONCodec, coins sdk.Coins, addresses []sdk.AccAddress, modAcc *authtypes.ModuleAccount) GenesisState {
	builder := NewAuthBankGenesisBuilder()

	for _, address := range addresses {
		builder.WithSimpleAccount(address, coins)
	}

	builder.WithSimpleModuleAccount(modAcc.Address, nil)

	return builder.BuildMarshalled(cdc)
}
