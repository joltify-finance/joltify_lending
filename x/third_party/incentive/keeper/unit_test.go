package keeper_test

import (
	"fmt"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	tmlog "github.com/tendermint/tendermint/libs/log"
	"time"

	types3 "github.com/joltify-finance/joltify_lending/x/third_party/cdp/types"
	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/keeper"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"
	hardtypes "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"

	"github.com/joltify-finance/joltify_lending/app"
	tmprototypes "github.com/tendermint/tendermint/proto/tendermint/types"
)

// NewTestContext sets up a basic context with an in-memory db
func NewTestContext(requiredStoreKeys ...storetypes.StoreKey) sdk.Context {
	memDB := db.NewMemDB()
	cms := store.NewCommitMultiStore(memDB)

	for _, key := range requiredStoreKeys {
		cms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, nil)
	}

	if err := cms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	return sdk.NewContext(cms, tmprototypes.Header{}, false, log.NewNopLogger())
}

// unitTester is a wrapper around suite.Suite, with common functionality for keeper unit tests.
// It can be embedded in structs the same way as suite.Suite.
type unitTester struct {
	suite.Suite
	keeper keeper.Keeper
	ctx    sdk.Context

	cdc               codec.Codec
	incentiveStoreKey storetypes.StoreKey
}

func (suite *unitTester) SetupSuite() {
	tApp := app.NewTestApp(tmlog.TestingLogger(), suite.T().TempDir())
	suite.cdc = tApp.AppCodec()

	suite.incentiveStoreKey = sdk.NewKVStoreKey(types2.StoreKey)
}

func (suite *unitTester) SetupTest() {
	suite.ctx = NewTestContext(suite.incentiveStoreKey)
	suite.keeper = suite.NewKeeper(&fakeParamSubspace{}, nil, nil, nil, nil)
}

func (suite *unitTester) TearDownTest() {
	suite.keeper = keeper.Keeper{}
	suite.ctx = sdk.Context{}
}

func (suite *unitTester) NewKeeper(paramSubspace types2.ParamSubspace, bk types2.BankKeeper, cdpk types2.CdpKeeper, hk types2.JoltKeeper, ak types2.AccountKeeper) keeper.Keeper {
	return keeper.NewKeeper(suite.cdc, suite.incentiveStoreKey, paramSubspace, bk, cdpk, hk, ak)
}

func (suite *unitTester) storeGlobalBorrowIndexes(indexes types2.MultiRewardIndexes) {
	for _, i := range indexes {
		suite.keeper.SetJoltBorrowRewardIndexes(suite.ctx, i.CollateralType, i.RewardIndexes)
	}
}

func (suite *unitTester) storeGlobalSupplyIndexes(indexes types2.MultiRewardIndexes) {
	for _, i := range indexes {
		suite.keeper.SetJoltSupplyRewardIndexes(suite.ctx, i.CollateralType, i.RewardIndexes)
	}
}

func (suite *unitTester) storeGlobalDelegatorIndexes(multiRewardIndexes types2.MultiRewardIndexes) {
	// Hardcoded to use bond denom
	multiRewardIndex, _ := multiRewardIndexes.GetRewardIndex(types2.BondDenom)
	suite.keeper.SetDelegatorRewardIndexes(suite.ctx, types2.BondDenom, multiRewardIndex.RewardIndexes)
}

func (suite *unitTester) storeGlobalSwapIndexes(indexes types2.MultiRewardIndexes) {
	for _, i := range indexes {
		suite.keeper.SetSwapRewardIndexes(suite.ctx, i.CollateralType, i.RewardIndexes)
	}
}

func (suite *unitTester) storeGlobalSavingsIndexes(indexes types2.MultiRewardIndexes) {
	for _, i := range indexes {
		suite.keeper.SetSavingsRewardIndexes(suite.ctx, i.CollateralType, i.RewardIndexes)
	}
}

func (suite *unitTester) storeJoltClaim(claim types2.JoltLiquidityProviderClaim) {
	suite.keeper.SetJoltLiquidityProviderClaim(suite.ctx, claim)
}

func (suite *unitTester) storeDelegatorClaim(claim types2.DelegatorClaim) {
	suite.keeper.SetDelegatorClaim(suite.ctx, claim)
}

func (suite *unitTester) storeSwapClaim(claim types2.SwapClaim) {
	suite.keeper.SetSwapClaim(suite.ctx, claim)
}

func (suite *unitTester) storeSavingsClaim(claim types2.SavingsClaim) {
	suite.keeper.SetSavingsClaim(suite.ctx, claim)
}

// fakeParamSubspace is a stub paramSpace to simplify keeper unit test setup.
type fakeParamSubspace struct {
	params types2.Params
}

func (subspace *fakeParamSubspace) GetParamSet(_ sdk.Context, ps paramtypes.ParamSet) {
	*(ps.(*types2.Params)) = subspace.params
}

func (subspace *fakeParamSubspace) SetParamSet(_ sdk.Context, ps paramtypes.ParamSet) {
	subspace.params = *(ps.(*types2.Params))
}

func (subspace *fakeParamSubspace) HasKeyTable() bool {
	// return true so the keeper does not try to call WithKeyTable, which does nothing
	return true
}

func (subspace *fakeParamSubspace) WithKeyTable(paramtypes.KeyTable) paramtypes.Subspace {
	// return an non-functional subspace to satisfy the interface
	return paramtypes.Subspace{}
}

// fakeSwapKeeper is a stub swap keeper.
// It can be used to return values to the incentive keeper without having to initialize a full swap keeper.
type fakeSwapKeeper struct {
	poolShares    map[string]sdk.Int
	depositShares map[string](map[string]sdk.Int)
}

var _ types2.SwapKeeper = newFakeSwapKeeper()

func newFakeSwapKeeper() *fakeSwapKeeper {
	return &fakeSwapKeeper{
		poolShares:    map[string]sdk.Int{},
		depositShares: map[string](map[string]sdk.Int){},
	}
}

func (k *fakeSwapKeeper) addPool(id string, shares sdk.Int) *fakeSwapKeeper {
	k.poolShares[id] = shares
	return k
}

func (k *fakeSwapKeeper) addDeposit(poolID string, depositor sdk.AccAddress, shares sdk.Int) *fakeSwapKeeper {
	if k.depositShares[poolID] == nil {
		k.depositShares[poolID] = map[string]sdk.Int{}
	}
	k.depositShares[poolID][depositor.String()] = shares
	return k
}

func (k *fakeSwapKeeper) GetPoolShares(_ sdk.Context, poolID string) (sdk.Int, bool) {
	shares, ok := k.poolShares[poolID]
	return shares, ok
}

func (k *fakeSwapKeeper) GetDepositorSharesAmount(_ sdk.Context, depositor sdk.AccAddress, poolID string) (sdk.Int, bool) {
	shares, found := k.depositShares[poolID][depositor.String()]
	return shares, found
}

// fakeHardKeeper is a stub jolt keeper.
// It can be used to return values to the incentive keeper without having to initialize a full jolt keeper.
type fakeHardKeeper struct {
	borrows  fakeHardState
	deposits fakeHardState
}

type fakeHardState struct {
	total           sdk.Coins
	interestFactors map[string]sdk.Dec
}

func newFakeHardState() fakeHardState {
	return fakeHardState{
		total:           nil,
		interestFactors: map[string]sdk.Dec{}, // initialize map to avoid panics on read
	}
}

var _ types2.JoltKeeper = newFakeHardKeeper()

func newFakeHardKeeper() *fakeHardKeeper {
	return &fakeHardKeeper{
		borrows:  newFakeHardState(),
		deposits: newFakeHardState(),
	}
}

func (k *fakeHardKeeper) addTotalBorrow(coin sdk.Coin, factor sdk.Dec) *fakeHardKeeper {
	k.borrows.total = k.borrows.total.Add(coin)
	k.borrows.interestFactors[coin.Denom] = factor
	return k
}

func (k *fakeHardKeeper) addTotalSupply(coin sdk.Coin, factor sdk.Dec) *fakeHardKeeper {
	k.deposits.total = k.deposits.total.Add(coin)
	k.deposits.interestFactors[coin.Denom] = factor
	return k
}

func (k *fakeHardKeeper) GetBorrowedCoins(_ sdk.Context) (sdk.Coins, bool) {
	if k.borrows.total == nil {
		return nil, false
	}
	return k.borrows.total, true
}

func (k *fakeHardKeeper) GetSuppliedCoins(_ sdk.Context) (sdk.Coins, bool) {
	if k.deposits.total == nil {
		return nil, false
	}
	return k.deposits.total, true
}

func (k *fakeHardKeeper) GetBorrowInterestFactor(_ sdk.Context, denom string) (sdk.Dec, bool) {
	f, ok := k.borrows.interestFactors[denom]
	return f, ok
}

func (k *fakeHardKeeper) GetSupplyInterestFactor(_ sdk.Context, denom string) (sdk.Dec, bool) {
	f, ok := k.deposits.interestFactors[denom]
	return f, ok
}

func (k *fakeHardKeeper) GetBorrow(_ sdk.Context, _ sdk.AccAddress) (hardtypes.Borrow, bool) {
	panic("unimplemented")
}

func (k *fakeHardKeeper) GetDeposit(_ sdk.Context, _ sdk.AccAddress) (hardtypes.Deposit, bool) {
	panic("unimplemented")
}

// fakeStakingKeeper is a stub staking keeper.
// It can be used to return values to the incentive keeper without having to initialize a full staking keeper.
type fakeStakingKeeper struct {
	delegations stakingtypes.Delegations
	validators  stakingtypes.Validators
}

var _ types2.StakingKeeper = newFakeStakingKeeper()

func newFakeStakingKeeper() *fakeStakingKeeper { return &fakeStakingKeeper{} }

func (k *fakeStakingKeeper) addBondedTokens(amount int64) *fakeStakingKeeper {
	if len(k.validators) != 0 {
		panic("cannot set total bonded if keeper already has validators set")
	}
	// add a validator with all the tokens
	k.validators = append(k.validators, stakingtypes.Validator{
		Status: stakingtypes.Bonded,
		Tokens: sdk.NewInt(amount),
	})
	return k
}

func (k *fakeStakingKeeper) TotalBondedTokens(_ sdk.Context) sdk.Int {
	total := sdk.ZeroInt()
	for _, val := range k.validators {
		if val.GetStatus() == stakingtypes.Bonded {
			total = total.Add(val.GetBondedTokens())
		}
	}
	return total
}

func (k *fakeStakingKeeper) GetDelegatorDelegations(_ sdk.Context, delegator sdk.AccAddress, maxRetrieve uint16) []stakingtypes.Delegation {
	return k.delegations
}

func (k *fakeStakingKeeper) GetValidator(_ sdk.Context, addr sdk.ValAddress) (stakingtypes.Validator, bool) {
	for _, val := range k.validators {
		if val.GetOperator().Equals(addr) {
			return val, true
		}
	}
	return stakingtypes.Validator{}, false
}

func (k *fakeStakingKeeper) GetValidatorDelegations(_ sdk.Context, valAddr sdk.ValAddress) []stakingtypes.Delegation {
	var delegations stakingtypes.Delegations
	for _, d := range k.delegations {
		if d.GetValidatorAddr().Equals(valAddr) {
			delegations = append(delegations, d)
		}
	}
	return delegations
}

// fakeCDPKeeper is a stub cdp keeper.
// It can be used to return values to the incentive keeper without having to initialize a full cdp keeper.
type fakeCDPKeeper struct {
	interestFactor *sdk.Dec
	totalPrincipal sdk.Int
}

var _ types2.CdpKeeper = newFakeCDPKeeper()

func newFakeCDPKeeper() *fakeCDPKeeper {
	return &fakeCDPKeeper{
		interestFactor: nil,
		totalPrincipal: sdk.ZeroInt(),
	}
}

func (k *fakeCDPKeeper) addInterestFactor(f sdk.Dec) *fakeCDPKeeper {
	k.interestFactor = &f
	return k
}

func (k *fakeCDPKeeper) addTotalPrincipal(p sdk.Int) *fakeCDPKeeper {
	k.totalPrincipal = p
	return k
}

func (k *fakeCDPKeeper) GetInterestFactor(_ sdk.Context, collateralType string) (sdk.Dec, bool) {
	if k.interestFactor != nil {
		return *k.interestFactor, true
	}
	return sdk.Dec{}, false
}

func (k *fakeCDPKeeper) GetTotalPrincipal(_ sdk.Context, collateralType string, principalDenom string) sdk.Int {
	return k.totalPrincipal
}

func (k *fakeCDPKeeper) GetCdpByOwnerAndCollateralType(_ sdk.Context, owner sdk.AccAddress, collateralType string) (types3.CDP, bool) {
	return types3.CDP{}, false
}

func (k *fakeCDPKeeper) GetCollateral(_ sdk.Context, collateralType string) (types3.CollateralParam, bool) {
	return types3.CollateralParam{}, false
}

// Assorted Testing Data

// note: amino panics when encoding times ≥ the start of year 10000.
var distantFuture = time.Date(9000, 1, 1, 0, 0, 0, 0, time.UTC)

func arbitraryCoins() sdk.Coins {
	return cs(c("btcb", 1))
}

func arbitraryAddress() sdk.AccAddress {
	_, addresses := app.GeneratePrivKeyAddressPairs(1)
	return addresses[0]
}

func arbitraryValidatorAddress() sdk.ValAddress {
	return generateValidatorAddresses(1)[0]
}

func generateValidatorAddresses(n int) []sdk.ValAddress {
	_, addresses := app.GeneratePrivKeyAddressPairs(n)
	var valAddresses []sdk.ValAddress
	for _, a := range addresses {
		valAddresses = append(valAddresses, sdk.ValAddress(a))
	}
	return valAddresses
}

var nonEmptyMultiRewardIndexes = types2.MultiRewardIndexes{
	{
		CollateralType: "bnb",
		RewardIndexes: types2.RewardIndexes{
			{
				CollateralType: "jolt",
				RewardFactor:   d("0.02"),
			},
			{
				CollateralType: "ujolt",
				RewardFactor:   d("0.04"),
			},
		},
	},
	{
		CollateralType: "btcb",
		RewardIndexes: types2.RewardIndexes{
			{
				CollateralType: "jolt",
				RewardFactor:   d("0.2"),
			},
			{
				CollateralType: "ujolt",
				RewardFactor:   d("0.4"),
			},
		},
	},
}

func extractCollateralTypes(indexes types2.MultiRewardIndexes) []string {
	var denoms []string
	for _, ri := range indexes {
		denoms = append(denoms, ri.CollateralType)
	}
	return denoms
}

func increaseAllRewardFactors(indexes types2.MultiRewardIndexes) types2.MultiRewardIndexes {
	increasedIndexes := make(types2.MultiRewardIndexes, len(indexes))
	copy(increasedIndexes, indexes)

	for i := range increasedIndexes {
		increasedIndexes[i].RewardIndexes = increaseRewardFactors(increasedIndexes[i].RewardIndexes)
	}
	return increasedIndexes
}

func increaseRewardFactors(indexes types2.RewardIndexes) types2.RewardIndexes {
	increasedIndexes := make(types2.RewardIndexes, len(indexes))
	copy(increasedIndexes, indexes)

	for i := range increasedIndexes {
		increasedIndexes[i].RewardFactor = increasedIndexes[i].RewardFactor.MulInt64(2)
	}
	return increasedIndexes
}

func appendUniqueMultiRewardIndex(indexes types2.MultiRewardIndexes) types2.MultiRewardIndexes {
	const uniqueDenom = "uniquedenom"

	for _, mri := range indexes {
		if mri.CollateralType == uniqueDenom {
			panic(fmt.Sprintf("tried to add unique multi reward index with denom '%s', but denom already existed", uniqueDenom))
		}
	}

	return append(indexes, types2.NewMultiRewardIndex(
		uniqueDenom,
		types2.RewardIndexes{
			{
				CollateralType: "jolt",
				RewardFactor:   d("0.02"),
			},
			{
				CollateralType: "ujolt",
				RewardFactor:   d("0.04"),
			},
		},
	),
	)
}

func appendUniqueEmptyMultiRewardIndex(indexes types2.MultiRewardIndexes) types2.MultiRewardIndexes {
	const uniqueDenom = "uniquedenom"

	for _, mri := range indexes {
		if mri.CollateralType == uniqueDenom {
			panic(fmt.Sprintf("tried to add unique multi reward index with denom '%s', but denom already existed", uniqueDenom))
		}
	}

	return append(indexes, types2.NewMultiRewardIndex(uniqueDenom, nil))
}

func appendUniqueRewardIndexToFirstItem(indexes types2.MultiRewardIndexes) types2.MultiRewardIndexes {
	newIndexes := make(types2.MultiRewardIndexes, len(indexes))
	copy(newIndexes, indexes)

	newIndexes[0].RewardIndexes = appendUniqueRewardIndex(newIndexes[0].RewardIndexes)
	return newIndexes
}

func appendUniqueRewardIndex(indexes types2.RewardIndexes) types2.RewardIndexes {
	const uniqueDenom = "uniquereward"

	for _, mri := range indexes {
		if mri.CollateralType == uniqueDenom {
			panic(fmt.Sprintf("tried to add unique reward index with denom '%s', but denom already existed", uniqueDenom))
		}
	}

	return append(
		indexes,
		types2.NewRewardIndex(uniqueDenom, d("0.02")),
	)
}
