package keeper_test

import (
	"fmt"
	"time"

	nfttypes "cosmossdk.io/x/nft"

	sdkmath "cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"
	types2 "github.com/joltify-finance/joltify_lending/x/spv/types"

	tmlog "github.com/cometbft/cometbft/libs/log"

	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/keeper"
	"github.com/joltify-finance/joltify_lending/x/third_party/incentive/types"
	hardtypes "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"

	"cosmossdk.io/store"
	db "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/suite"

	tmprototypes "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/joltify-finance/joltify_lending/app"
)

// NewTestContext sets up a basic context with an in-memory db
func NewTestContext(requiredStoreKeys ...storetypes.StoreKey) context.Context {
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
	ctx    context.Context

	cdc               codec.Codec
	incentiveStoreKey storetypes.StoreKey
}

func (suite *unitTester) SetupSuite() {
	tApp := app.NewTestApp(tmlog.TestingLogger(), suite.T().TempDir())
	suite.cdc = tApp.AppCodec()

	suite.incentiveStoreKey = sdk.NewKVStoreKey(types.StoreKey)
}

func (suite *unitTester) SetupTest() {
	suite.ctx = NewTestContext(suite.incentiveStoreKey)
	suite.keeper = suite.NewKeeper(&fakeParamSubspace{}, nil, nil, nil, nil, nil, nil)
}

func (suite *unitTester) TearDownTest() {
	suite.keeper = keeper.Keeper{}
	suite.ctx = context.Context{}
}

func (suite *unitTester) NewKeeper(paramSubspace types.ParamSubspace, bk types.BankKeeper, hk types.JoltKeeper, ak types.AccountKeeper, swapKeeper types.SwapKeeper, spvKeeper types.SPVKeeper, nftKeeper types.NFTKeeper) keeper.Keeper {
	return keeper.NewKeeper(suite.cdc, suite.incentiveStoreKey, paramSubspace, bk, hk, ak, swapKeeper, spvKeeper, nftKeeper)
}

func (suite *unitTester) storeJoltClaim(claim types.JoltLiquidityProviderClaim) {
	suite.keeper.SetJoltLiquidityProviderClaim(suite.ctx, claim)
}

func (suite *unitTester) storeGlobalBorrowIndexes(indexes types.MultiRewardIndexes) {
	for _, i := range indexes {
		suite.keeper.SetJoltBorrowRewardIndexes(suite.ctx, i.CollateralType, i.RewardIndexes)
	}
}

func (suite *unitTester) storeGlobalSupplyIndexes(indexes types.MultiRewardIndexes) {
	for _, i := range indexes {
		suite.keeper.SetJoltSupplyRewardIndexes(suite.ctx, i.CollateralType, i.RewardIndexes)
	}
}

// fakeParamSubspace is a stub paramSpace to simplify keeper unit test setup.
type fakeParamSubspace struct {
	params types.Params
}

func (subspace *fakeParamSubspace) GetParamSet(_ context.Context, ps paramtypes.ParamSet) {
	*(ps.(*types.Params)) = subspace.params
}

func (subspace *fakeParamSubspace) SetParamSet(_ context.Context, ps paramtypes.ParamSet) {
	subspace.params = *(ps.(*types.Params))
}

func (subspace *fakeParamSubspace) HasKeyTable() bool {
	// return true so the keeper does not try to call WithKeyTable, which does nothing
	return true
}

func (suite *unitTester) storeSwapClaim(claim types.SwapClaim) {
	suite.keeper.SetSwapClaim(suite.ctx, claim)
}

func (suite *unitTester) storeGlobalSwapIndexes(indexes types.MultiRewardIndexes) {
	for _, i := range indexes {
		suite.keeper.SetSwapRewardIndexes(suite.ctx, i.CollateralType, i.RewardIndexes)
	}
}

func (subspace *fakeParamSubspace) WithKeyTable(paramtypes.KeyTable) paramtypes.Subspace {
	// return an non-functional subspace to satisfy the interface
	return paramtypes.Subspace{}
}

// fakeJoltKeeper is a stub jolt keeper.
// It can be used to return values to the incentive keeper without having to initialize a full jolt keeper.
type fakeJoltKeeper struct {
	borrows  fakeHardState
	deposits fakeHardState
}

type fakeHardState struct {
	total           sdk.Coins
	interestFactors map[string]sdkmath.LegacyDec
}

func newFakeHardState() fakeHardState {
	return fakeHardState{
		total:           nil,
		interestFactors: map[string]sdkmath.LegacyDec{}, // initialize map to avoid panics on read
	}
}

var _ types.JoltKeeper = newFakeHardKeeper()

func newFakeHardKeeper() *fakeJoltKeeper {
	return &fakeJoltKeeper{
		borrows:  newFakeHardState(),
		deposits: newFakeHardState(),
	}
}

func (k *fakeJoltKeeper) addTotalBorrow(coin sdk.Coin, factor sdkmath.LegacyDec) *fakeJoltKeeper {
	k.borrows.total = k.borrows.total.Add(coin)
	k.borrows.interestFactors[coin.Denom] = factor
	return k
}

func (k *fakeJoltKeeper) addTotalSupply(coin sdk.Coin, factor sdkmath.LegacyDec) *fakeJoltKeeper {
	k.deposits.total = k.deposits.total.Add(coin)
	k.deposits.interestFactors[coin.Denom] = factor
	return k
}

func (k *fakeJoltKeeper) GetBorrowedCoins(_ context.Context) (sdk.Coins, bool) {
	if k.borrows.total == nil {
		return nil, false
	}
	return k.borrows.total, true
}

func (k *fakeJoltKeeper) GetSuppliedCoins(_ context.Context) (sdk.Coins, bool) {
	if k.deposits.total == nil {
		return nil, false
	}
	return k.deposits.total, true
}

func (k *fakeJoltKeeper) GetBorrowInterestFactor(_ context.Context, denom string) (sdkmath.LegacyDec, bool) {
	f, ok := k.borrows.interestFactors[denom]
	return f, ok
}

func (k *fakeJoltKeeper) GetSupplyInterestFactor(_ context.Context, denom string) (sdkmath.LegacyDec, bool) {
	f, ok := k.deposits.interestFactors[denom]
	return f, ok
}

func (k *fakeJoltKeeper) GetBorrow(_ context.Context, _ sdk.AccAddress) (hardtypes.Borrow, bool) {
	panic("unimplemented")
}

func (k *fakeJoltKeeper) GetDeposit(_ context.Context, _ sdk.AccAddress) (hardtypes.Deposit, bool) {
	panic("unimplemented")
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

var nonEmptyMultiRewardIndexes = types.MultiRewardIndexes{
	{
		CollateralType: "bnb",
		RewardIndexes: types.RewardIndexes{
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
		RewardIndexes: types.RewardIndexes{
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

func extractCollateralTypes(indexes types.MultiRewardIndexes) []string {
	var denoms []string
	for _, ri := range indexes {
		denoms = append(denoms, ri.CollateralType)
	}
	return denoms
}

func increaseAllRewardFactors(indexes types.MultiRewardIndexes) types.MultiRewardIndexes {
	increasedIndexes := make(types.MultiRewardIndexes, len(indexes))
	copy(increasedIndexes, indexes)

	for i := range increasedIndexes {
		increasedIndexes[i].RewardIndexes = increaseRewardFactors(increasedIndexes[i].RewardIndexes)
	}
	return increasedIndexes
}

func increaseRewardFactors(indexes types.RewardIndexes) types.RewardIndexes {
	increasedIndexes := make(types.RewardIndexes, len(indexes))
	copy(increasedIndexes, indexes)

	for i := range increasedIndexes {
		increasedIndexes[i].RewardFactor = increasedIndexes[i].RewardFactor.MulInt64(2)
	}
	return increasedIndexes
}

func appendUniqueMultiRewardIndex(indexes types.MultiRewardIndexes) types.MultiRewardIndexes {
	const uniqueDenom = "uniquedenom"

	for _, mri := range indexes {
		if mri.CollateralType == uniqueDenom {
			panic(fmt.Sprintf("tried to add unique multi reward index with denom '%s', but denom already existed", uniqueDenom))
		}
	}

	return append(indexes, types.NewMultiRewardIndex(
		uniqueDenom,
		types.RewardIndexes{
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

func appendUniqueEmptyMultiRewardIndex(indexes types.MultiRewardIndexes) types.MultiRewardIndexes {
	const uniqueDenom = "uniquedenom"

	for _, mri := range indexes {
		if mri.CollateralType == uniqueDenom {
			panic(fmt.Sprintf("tried to add unique multi reward index with denom '%s', but denom already existed", uniqueDenom))
		}
	}

	return append(indexes, types.NewMultiRewardIndex(uniqueDenom, nil))
}

func appendUniqueRewardIndexToFirstItem(indexes types.MultiRewardIndexes) types.MultiRewardIndexes {
	newIndexes := make(types.MultiRewardIndexes, len(indexes))
	copy(newIndexes, indexes)

	newIndexes[0].RewardIndexes = appendUniqueRewardIndex(newIndexes[0].RewardIndexes)
	return newIndexes
}

func appendUniqueRewardIndex(indexes types.RewardIndexes) types.RewardIndexes {
	const uniqueDenom = "uniquereward"

	for _, mri := range indexes {
		if mri.CollateralType == uniqueDenom {
			panic(fmt.Sprintf("tried to add unique reward index with denom '%s', but denom already existed", uniqueDenom))
		}
	}

	return append(
		indexes,
		types.NewRewardIndex(uniqueDenom, d("0.02")),
	)
}

// fakeSwapKeeper is a stub swap keeper.
// It can be used to return values to the incentive keeper without having to initialize a full swap keeper.
type fakeSwapKeeper struct {
	poolShares    map[string]sdkmath.Int
	depositShares map[string](map[string]sdkmath.Int)
}

var _ types.SwapKeeper = newFakeSwapKeeper()

func newFakeSwapKeeper() *fakeSwapKeeper {
	return &fakeSwapKeeper{
		poolShares:    map[string]sdkmath.Int{},
		depositShares: map[string](map[string]sdkmath.Int){},
	}
}

func (k *fakeSwapKeeper) addPool(id string, shares sdkmath.Int) *fakeSwapKeeper {
	k.poolShares[id] = shares
	return k
}

func (k *fakeSwapKeeper) addDeposit(poolID string, depositor sdk.AccAddress, shares sdkmath.Int) *fakeSwapKeeper {
	if k.depositShares[poolID] == nil {
		k.depositShares[poolID] = map[string]sdkmath.Int{}
	}
	k.depositShares[poolID][depositor.String()] = shares
	return k
}

func (k *fakeSwapKeeper) GetPoolShares(_ context.Context, poolID string) (sdkmath.Int, bool) {
	shares, ok := k.poolShares[poolID]
	return shares, ok
}

func (k *fakeSwapKeeper) GetDepositorSharesAmount(_ context.Context, depositor sdk.AccAddress, poolID string) (sdkmath.Int, bool) {
	shares, found := k.depositShares[poolID][depositor.String()]
	return shares, found
}

var _ types.SPVKeeper = newFakeSPVKeeper()

// fakeSwapKeeper is a stub swap keeper.
// It can be used to return values to the incentive keeper without having to initialize a full swap keeper.
type fakeSPVKeeper struct {
	poolShares    map[string]sdkmath.Int
	depositShares map[string](map[string]sdkmath.Int)
}

func newFakeSPVKeeper() *fakeSPVKeeper {
	return &fakeSPVKeeper{
		poolShares:    map[string]sdkmath.Int{},
		depositShares: map[string](map[string]sdkmath.Int){},
	}
}

func (k *fakeSPVKeeper) AfterSPVInterestPaid(ctx context.Context, poolID string, interestPaid sdkmath.Int) {
}

func (k *fakeSPVKeeper) GetPools(ctx context.Context, index string) (poolInfo types2.PoolInfo, ok bool) {
	poolInfo = types2.PoolInfo{
		Index:         "test-pool",
		ReserveFactor: sdkmath.LegacyMustNewDecFromStr("0.15"),
		PoolNFTIds:    []string{"c1", "c2", "c3", "c4", "c5", "c6"},
	}
	return poolInfo, true
}

func (k *fakeSPVKeeper) GetDepositor(ctx context.Context, poolIndex string, walletAddress sdk.AccAddress) (depositor types2.DepositorInfo, found bool) {
	return types2.DepositorInfo{}, true
}

type fakeNFTKeeper struct {
	nftClass map[string]*nfttypes.Class
	nft      map[string]*nfttypes.NFT
}

func newFakeNFTKeeper(nftsClass []*nfttypes.Class, nfts []*nfttypes.NFT) *fakeNFTKeeper {
	val := make(map[string]*nfttypes.Class)
	valNFT := make(map[string]*nfttypes.NFT)
	for _, nft := range nftsClass {
		val[nft.Id] = nft
	}

	for _, el := range nfts {
		valNFT[el.Id] = el
	}
	return &fakeNFTKeeper{
		nftClass: val,
		nft:      valNFT,
	}
}

func (fn fakeNFTKeeper) GetClass(ctx context.Context, classID string) (nfttypes.Class, bool) {
	nft, ok := fn.nftClass[classID]
	return *nft, ok
}

func (fn fakeNFTKeeper) GetNFT(ctx context.Context, classID, nftID string) (nfttypes.NFT, bool) {
	nft, ok := fn.nft[nftID]
	return *nft, ok
}

func (fn fakeNFTKeeper) UpdateClass(ctx context.Context, class nfttypes.Class) error {
	fn.nftClass[class.Id] = &class
	return nil
}

func (fn fakeNFTKeeper) Update(ctx context.Context, token nfttypes.NFT) error {
	fn.nft[token.Id] = &token
	return nil
}
