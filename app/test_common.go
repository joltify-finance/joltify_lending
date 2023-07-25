package app

import (
	sdkmath "cosmossdk.io/math"
	"encoding/json"
	feemarketkeeper "github.com/evmos/ethermint/x/feemarket/keeper"
	evmutilkeeper "github.com/joltify-finance/joltify_lending/x/third_party/evmutil/keeper"
	"math/rand"
	"testing"
	"time"

	evmkeeper "github.com/evmos/ethermint/x/evm/keeper"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	pruningtypes "github.com/cosmos/cosmos-sdk/pruning/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	tmjson "github.com/tendermint/tendermint/libs/json"
	tmlog "github.com/tendermint/tendermint/libs/log"

	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	distkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	mintkeeper "github.com/joltify-finance/joltify_lending/x/mint/keeper"
	minttypes "github.com/joltify-finance/joltify_lending/x/mint/types"
	auctionkeeper "github.com/joltify-finance/joltify_lending/x/third_party/auction/keeper"
	incentivekeeper "github.com/joltify-finance/joltify_lending/x/third_party/incentive/keeper"
	joltkeeper "github.com/joltify-finance/joltify_lending/x/third_party/jolt/keeper"
	pricefeedkeeper "github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/keeper"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
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
func NewTestApp(logger tmlog.Logger, rootDir string) TestApp {
	SetSDKConfig()
	return NewTestAppFromSealed(logger, rootDir)
}

func genesisStateWithValSet(
	app *App, genesisState GenesisState,
	valSet *tmtypes.ValidatorSet, genAccs []authtypes.GenesisAccount,
	balances ...banktypes.Balance,
) GenesisState {
	// set genesis accounts
	// authGenesis := authtypes.NewGenesisState(authtypes.DefaultParams(), genAccs)
	// genesisState[authtypes.ModuleName] = app.AppCodec().MustMarshalJSON(authGenesis)

	validators := make([]stakingtypes.Validator, 0, len(valSet.Validators))
	delegations := make([]stakingtypes.Delegation, 0, len(valSet.Validators))

	bondAmt := sdk.DefaultPowerReduction

	for _, val := range valSet.Validators {
		pk, err := cryptocodec.FromTmPubKeyInterface(val.PubKey)
		if err != nil {
			panic(err)
		}
		pkAny, err := codectypes.NewAnyWithValue(pk)
		if err != nil {
			panic(err)
		}
		validator := stakingtypes.Validator{
			OperatorAddress:   sdk.ValAddress(val.Address).String(),
			ConsensusPubkey:   pkAny,
			Jailed:            false,
			Status:            stakingtypes.Bonded,
			Tokens:            bondAmt,
			DelegatorShares:   sdk.OneDec(),
			Description:       stakingtypes.Description{},
			UnbondingHeight:   int64(0),
			UnbondingTime:     time.Unix(0, 0).UTC(),
			Commission:        stakingtypes.NewCommission(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec()),
			MinSelfDelegation: sdk.ZeroInt(),
		}
		validators = append(validators, validator)
		delegations = append(delegations, stakingtypes.NewDelegation(genAccs[0].GetAddress(), val.Address.Bytes(), sdk.OneDec()))

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
	bankGenesis := banktypes.NewGenesisState(banktypes.DefaultGenesisState().Params, balances, totalSupply, []banktypes.Metadata{})
	genesisState[banktypes.ModuleName] = app.AppCodec().MustMarshalJSON(bankGenesis)

	return genesisState
}

var DefaultConsensusParams = &abci.ConsensusParams{
	Block: &abci.BlockParams{
		MaxBytes: 200000,
		MaxGas:   2000000,
	},
	Evidence: &tmproto.EvidenceParams{
		MaxAgeNumBlocks: 302400,
		MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
		MaxBytes:        10000,
	},
	Validator: &tmproto.ValidatorParams{
		PubKeyTypes: []string{
			tmtypes.ABCIPubKeyTypeEd25519,
		},
	},
}

// NewTestAppFromSealed creates a TestApp without first setting sdk config.
func NewTestAppFromSealed(logger tmlog.Logger, rootDir string) TestApp {
	encCfg := MakeEncodingConfig()
	app := NewApp(
		logger, tmdb.NewMemDB(), rootDir, nil, encCfg,
		Options{},
		baseapp.SetPruning(pruningtypes.NewPruningOptionsFromString(pruningtypes.PruningOptionDefault)),
		baseapp.SetMinGasPrices("0stake"),
	)

	privVal := ed25519.GenPrivKey()
	pubKey, err := cryptocodec.ToTmPubKeyInterface(privVal.PubKey())
	if err != nil {
		panic(err)
	}

	// create validator set with single validator
	validator := tmtypes.NewValidator(pubKey, 1)
	valSet := tmtypes.NewValidatorSet([]*tmtypes.Validator{validator})

	// generate genesis account
	senderPrivKey := secp256k1.GenPrivKey()
	acc := authtypes.NewBaseAccount(senderPrivKey.PubKey().Address().Bytes(), senderPrivKey.PubKey(), 0, 0)
	balance := banktypes.Balance{
		Address: acc.GetAddress().String(),
		Coins:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100000000000000))),
	}

	genesisState := NewDefaultGenesisState(encCfg.Marshaler)
	genesisState = genesisStateWithValSet(app, genesisState, valSet, []authtypes.GenesisAccount{acc}, balance)

	stateBytes, err := tmjson.MarshalIndent(genesisState, "", " ")
	if err != nil {
		panic(err)
	}

	// Initialize the chain
	app.InitChain(
		abci.RequestInitChain{
			Validators:      []abci.ValidatorUpdate{},
			ConsensusParams: DefaultConsensusParams,
			AppStateBytes:   stateBytes,
			ChainId:         "joltifychain_888-1",
		},
	)

	header := tmproto.Header{Height: 1, ChainID: "joltifychain_888-1", Time: time.Now().UTC()}

	ctx := app.BaseApp.NewContext(false, header)
	return TestApp{App: *app, Ctx: ctx}
}

func (tApp TestApp) GetAccountKeeper() authkeeper.AccountKeeper { return tApp.accountKeeper }
func (tApp TestApp) GetBankKeeper() bankkeeper.Keeper           { return tApp.bankKeeper }
func (tApp TestApp) GetStakingKeeper() stakingkeeper.Keeper     { return tApp.stakingKeeper }
func (tApp TestApp) GetSlashingKeeper() slashingkeeper.Keeper   { return tApp.slashingKeeper }
func (tApp TestApp) GetMintKeeper() mintkeeper.Keeper           { return tApp.mintKeeper }
func (tApp TestApp) GetDistrKeeper() distkeeper.Keeper          { return tApp.distrKeeper }
func (tApp TestApp) GetGovKeeper() govkeeper.Keeper             { return tApp.govKeeper }
func (tApp TestApp) GetCrisisKeeper() crisiskeeper.Keeper       { return tApp.crisisKeeper }
func (tApp TestApp) GetParamsKeeper() paramskeeper.Keeper       { return tApp.paramsKeeper }

func (tApp TestApp) GetAuctionKeeper() auctionkeeper.Keeper     { return tApp.auctionKeeper }
func (tApp TestApp) GetPriceFeedKeeper() pricefeedkeeper.Keeper { return tApp.pricefeedKeeper }
func (tApp TestApp) GetJoltKeeper() joltkeeper.Keeper           { return tApp.joltKeeper }
func (tApp TestApp) GetIncentiveKeeper() incentivekeeper.Keeper { return tApp.incentiveKeeper }
func (tApp TestApp) GetEVMKeeper() evmkeeper.Keeper             { return *tApp.evmKeeper }
func (tApp TestApp) GetEVMUtilKeeper() evmutilkeeper.Keeper     { return tApp.evmutilKeeper }
func (tApp TestApp) GetFeeMarketKeeper() feemarketkeeper.Keeper { return tApp.feeMarketKeeper }

// LegacyAmino returns the app's amino codec.
func (app *App) LegacyAmino() *codec.LegacyAmino {
	return app.legacyAmino
}

// AppCodec returns the app's app codec.
func (app *App) AppCodec() codec.Codec {
	return app.appCodec
}

// InitializeFromGenesisStates calls InitChain on the app using the provided genesis states.
// If any module genesis states are missing, defaults are used.
func (tApp TestApp) InitializeFromGenesisStates(genAccs []authtypes.GenesisAccount, coins sdk.Coins, genesisStates ...GenesisState) TestApp {
	return tApp.InitializeFromGenesisStatesWithTimeAndChainIDAndHeight(emptyTime, testChainID, defaultInitialHeight, genAccs, coins, genesisStates...)
}

// InitializeFromGenesisStatesWithTime calls InitChain on the app using the provided genesis states and time.
// If any module genesis states are missing, defaults are used.
func (tApp TestApp) InitializeFromGenesisStatesWithTime(genTime time.Time, genAccs []authtypes.GenesisAccount, coins sdk.Coins, genesisStates ...GenesisState) TestApp {
	return tApp.InitializeFromGenesisStatesWithTimeAndChainIDAndHeight(genTime, testChainID, defaultInitialHeight, genAccs, coins, genesisStates...)
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

	privVal := ed25519.GenPrivKey()
	pubKey, err := cryptocodec.ToTmPubKeyInterface(privVal.PubKey())
	if err != nil {
		panic(err)
	}

	// create validator set with single validator
	validator := tmtypes.NewValidator(pubKey, 1)
	valSet := tmtypes.NewValidatorSet([]*tmtypes.Validator{validator})
	defaultCoins := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100000000000000))
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

	genesisState = genesisStateWithValSet(&tApp.App, genesisState, valSet, genAccs, balances...)
	// Initialize the chain
	stateBytes, err := json.Marshal(genesisState)
	if err != nil {
		panic(err)
	}
	tApp.InitChain(
		abci.RequestInitChain{
			Time:          genTime,
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: stateBytes,
			ChainId:       chainID,
			// Set consensus params, which is needed by x/feemarket
			ConsensusParams: &abci.ConsensusParams{
				Block: &abci.BlockParams{
					MaxBytes: 200000,
					MaxGas:   20000000,
				},
			},
			InitialHeight: initialHeight,
		},
	)
	tApp.Commit()
	tApp.BeginBlock(abci.RequestBeginBlock{
		Header: tmproto.Header{
			Height: tApp.LastBlockHeight() + 1, Time: genTime, ChainID: chainID,
		},
	})
	return tApp
}

// RandomAddress non-deterministically generates a new address, discarding the private key.
func RandomAddress() sdk.AccAddress {
	secret := make([]byte, 32)
	_, err := rand.Read(secret)
	if err != nil {
		panic("Could not read randomness")
	}
	key := secp256k1.GenPrivKeyFromSecret(secret)
	return sdk.AccAddress(key.PubKey().Address())
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
