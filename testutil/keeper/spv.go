package keeper

import (
	"fmt"
	"strconv"
	"testing"

	types2 "github.com/joltify-finance/joltify_lending/x/third_party/pricefeed/types"

	"github.com/cosmos/gogoproto/proto"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	kycmoduletypes "github.com/joltify-finance/joltify_lending/x/kyc/types"

	tmdb "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/joltify-finance/joltify_lending/x/spv/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/stretchr/testify/require"
)

const (
	oneWeek  = 7 * 24 * 3600
	oneMonth = oneWeek * 4
	oneYear  = oneWeek * 52
)

type mockKycKeeper struct{}

var Wallets = []string{
	"jolt1kkujrm0lqeu0e5va5f6mmwk87wva0k8cmam8jq",
	"jolt166yyvsypvn6cwj2rc8sme4dl6v0g62hn3862kl",
	"jolt1z0y0zl0trsnuqmqf5v034pyv9sp39jg3rv6lsm",
	"jolt1fcaa73cc9c2l3l2u57skddgd0zm749ncukx90g",
	"jolt1ut358ywu78ztkt5m90dwmklz79rwau6vs8vhlp",
	"jolt1v9ls99c83dst7x6xwwnsjcyp5zsa3acfhaxq5n",
	"jolt169a92jz2rmxy0ll73kztlmtucswvvft78xeqne",
	"jolt13xxls80rw3p036zyfy8hhtjyvft4ckg5a09agh",
}

func (m mockKycKeeper) GetInvestorWallets(_ sdk.Context, investorID string) (kycmoduletypes.Investor, error) {
	return kycmoduletypes.Investor{InvestorId: investorID, WalletAddress: Wallets}, nil
}

func (m mockKycKeeper) GetProjects(ctx sdk.Context) (projectsInfo []*kycmoduletypes.ProjectInfo) {
	b := kycmoduletypes.BasicInfo{
		Description:    "This is the test info",
		ProjectsUrl:    "empty",
		ProjectCountry: "ABC",
		BusinessNumber: "ABC123",
		Reserved:       []byte("reserved"),
		ProjectName:    "This is the test info",
		Email:          "email address",
		Name:           "example.com",
	}

	acc, _ := sdk.AccAddressFromBech32("jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0")
	pi1 := kycmoduletypes.ProjectInfo{
		Index:           1,
		SPVName:         "defaultSPV",
		ProjectOwner:    acc,
		BasicInfo:       &b,
		ProjectLength:   31536000, // 1 year
		PayFreq:         "15768000",
		BaseApy:         sdk.NewDecWithPrec(12, 2),
		MarketId:        "aud:usd",
		SeparatePool:    false,
		JuniorMinRatio:  sdk.NewDecWithPrec(1, 15),
		MinBorrowAmount: sdk.NewInt(200),
	}

	b2 := kycmoduletypes.BasicInfo{
		Description:    "This is the test info2",
		ProjectsUrl:    "empty2",
		ProjectCountry: "ABC2",
		BusinessNumber: "ABC123-2",
		Reserved:       []byte("reserved"),
		ProjectName:    "This is the test info2",
		Email:          "example.com2",
		Name:           "email address2",
	}

	pi2 := kycmoduletypes.ProjectInfo{
		Index:           2,
		SPVName:         "defaultSPV2",
		ProjectOwner:    acc,
		BasicInfo:       &b2,
		ProjectLength:   31536000, // 1 year
		PayFreq:         "15768000",
		BaseApy:         sdk.NewDecWithPrec(12, 2),
		MarketId:        "aud:usd",
		SeparatePool:    false,
		JuniorMinRatio:  sdk.NewDecWithPrec(1, 15),
		MinBorrowAmount: sdk.NewInt(200),
	}

	pi3 := kycmoduletypes.ProjectInfo{
		Index:           3,
		SPVName:         "defaultSPV3",
		ProjectOwner:    acc,
		BasicInfo:       &b2,
		ProjectLength:   oneYear, // 1 year
		PayFreq:         strconv.Itoa(oneMonth),
		BaseApy:         sdk.NewDecWithPrec(12, 2),
		MarketId:        "aud:usd",
		SeparatePool:    false,
		JuniorMinRatio:  sdk.NewDecWithPrec(1, 15),
		MinBorrowAmount: sdk.NewInt(200),
	}

	pi4 := kycmoduletypes.ProjectInfo{
		Index:           4,
		SPVName:         "defaultSPV3",
		ProjectOwner:    acc,
		BasicInfo:       &b2,
		ProjectLength:   oneYear, // 1 year
		PayFreq:         strconv.Itoa(oneWeek),
		BaseApy:         sdk.NewDecWithPrec(10, 2),
		MarketId:        "aud:usd",
		SeparatePool:    false,
		JuniorMinRatio:  sdk.NewDecWithPrec(1, 15),
		MinBorrowAmount: sdk.NewInt(200),
	}

	return []*kycmoduletypes.ProjectInfo{&pi1, &pi2, &pi3, &pi4}
}

func (m mockKycKeeper) GetByWallet(_ sdk.Context, wallet string) (kycmoduletypes.Investor, error) {
	inv := kycmoduletypes.Investor{
		InvestorId:    "1",
		WalletAddress: []string{wallet},
	}

	inv2 := kycmoduletypes.Investor{InvestorId: "2", WalletAddress: Wallets}

	for _, el := range Wallets {
		if wallet == el {
			return inv2, nil
		}
	}

	return inv, nil
}

type mockAccKeeper struct{}

func (m mockAccKeeper) GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI {
	// TODO implement me
	panic("implement me")
}

func (m mockAccKeeper) GetModuleAccount(ctx sdk.Context, name string) authtypes.ModuleAccountI {
	addr := authtypes.NewModuleAddress(types.ModuleAccount)
	baseAcc := authtypes.NewBaseAccountWithAddress(addr)
	return authtypes.NewModuleAccount(baseAcc, types.ModuleName, "mint")
}

func (m mockAccKeeper) GetModuleAddress(name string) sdk.AccAddress {
	// TODO implement me
	panic("implement me")
}

type mockNFTKeeper struct {
	classes         map[string]*nft.Class
	nfts            map[string]*nft.NFT
	nftsWithClassID map[string]*nft.NFT
}

func (m mockNFTKeeper) Burn(ctx sdk.Context, classID string, nftID string) error {
	key := fmt.Sprintf("%v:%v", classID, nftID)
	delete(m.nftsWithClassID, key)
	return nil
}

func (m mockNFTKeeper) Transfer(ctx sdk.Context, classID string, nftID string, receiver sdk.AccAddress) error {
	panic("implement me")
}

func (m mockNFTKeeper) GetTotalSupply(ctx sdk.Context, classID string) uint64 {
	counter := 0
	for k := range m.nftsWithClassID {
		if classID == k {
			counter++
		}
	}
	return uint64(counter)
}

func (m mockNFTKeeper) GetNFT(ctx sdk.Context, classID, nftID string) (nft.NFT, bool) {
	key := fmt.Sprintf("%v:%v", classID, nftID)
	thisNft, found := m.nftsWithClassID[key]
	if !found {
		return nft.NFT{}, false
	}

	bz, err := proto.Marshal(thisNft)
	if err != nil {
		panic(err)
	}

	var returnNFT nft.NFT
	err = proto.Unmarshal(bz, &returnNFT)
	if err != nil {
		panic(err)
	}
	return returnNFT, true
}

func (m mockNFTKeeper) Update(ctx sdk.Context, nftToken nft.NFT) error {
	key := fmt.Sprintf("%v:%v", nftToken.ClassId, nftToken.Id)
	m.nftsWithClassID[key] = &nftToken

	return nil
}

func (m mockNFTKeeper) Mint(ctx sdk.Context, nft nft.NFT, receiver sdk.AccAddress) error {
	m.nfts[receiver.String()] = &nft
	key := fmt.Sprintf("%v:%v", nft.ClassId, nft.Id)
	m.nftsWithClassID[key] = &nft
	return nil
}

func (m mockNFTKeeper) SaveClass(ctx sdk.Context, class nft.Class) error {
	m.classes[class.Id] = &class
	return nil
}

func (m mockNFTKeeper) UpdateClass(ctx sdk.Context, class nft.Class) error {
	var borrowInterest types.BorrowInterest
	err := proto.Unmarshal(class.Data.Value, &borrowInterest)
	if err != nil {
		panic(err)
	}
	m.classes[class.Id] = &class
	return nil
}

func (m mockNFTKeeper) GetClass(ctx sdk.Context, classID string) (nft.Class, bool) {
	r, ok := m.classes[classID]
	return *r, ok
}

type mockbankKeeper struct {
	BankData map[string]sdk.Coins
}

func (m mockbankKeeper) SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error {
	m.BankData[recipientModule] = m.BankData[recipientModule].Add(amt...)
	return nil
}

func (m mockbankKeeper) SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
	addr := authtypes.NewModuleAddress(recipientModule)
	m.BankData[addr.String()] = m.BankData[addr.String()].Add(amt...)
	return nil
}

func (m mockbankKeeper) SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	// addr := authtypes.NewModuleAddress(senderModule)
	m.BankData[recipientAddr.String()] = m.BankData[recipientAddr.String()].Add(amt...)
	// m.BankData[addr.String()]= m.BankData[addr.String()].Sub(amt...)
	return nil
}

func (m mockbankKeeper) GetSupply(ctx sdk.Context, denom string) sdk.Coin {
	// TODO implement me
	panic("implement me")
}

func (m mockbankKeeper) GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	coins, ok := m.BankData[addr.String()]
	if !ok {
		panic("address cannot be found")
	}

	ok, coin := coins.Find(denom)
	if !ok {
		panic("denom cannot be found")
	}

	return coin
}

func (m mockbankKeeper) GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins {
	// TODO implement me
	panic("implement me")
}

func (m mockbankKeeper) SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins {
	// TODO implement me
	panic("implement me")
}

func (m mockbankKeeper) BurnCoins(ctx sdk.Context, name string, amt sdk.Coins) error {
	// TODO implement me
	panic("implement me")
}

type mockPriceFeedKeeper struct{}

type MockAuctionKeeper struct {
	AuctionAmount []sdk.Coin
	SellerBid     []string
}

func (m MockAuctionKeeper) StartSurplusAuction(ctx sdk.Context, seller string, lot sdk.Coin, bidDenom string) (uint64, error) {
	m.AuctionAmount[0] = lot
	m.SellerBid[0] = seller
	m.SellerBid[1] = bidDenom
	return 1, nil
}

func (m mockPriceFeedKeeper) GetCurrentPrice(ctx sdk.Context, marketID string) (types2.CurrentPrice, error) {
	return types2.CurrentPrice{MarketID: "aud:usd", Price: sdk.MustNewDecFromStr("0.7")}, nil
}

func SpvKeeper(t testing.TB) (*keeper.Keeper, types.NFTKeeper, types.BankKeeper, MockAuctionKeeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"SpvParams",
	)
	kycKeeper := mockKycKeeper{}
	accKeeper := mockAccKeeper{}
	nftKeeper := mockNFTKeeper{
		classes:         make(map[string]*nft.Class),
		nfts:            make(map[string]*nft.NFT),
		nftsWithClassID: make(map[string]*nft.NFT),
	}
	priceFeedKeeper := mockPriceFeedKeeper{}
	bankKeeper := mockbankKeeper{make(map[string]sdk.Coins)}
	auctionKeeper := MockAuctionKeeper{
		AuctionAmount: make([]sdk.Coin, 1),
		SellerBid:     make([]string, 2),
	}

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		paramsSubspace,
		kycKeeper,
		bankKeeper,
		accKeeper,
		nftKeeper,
		priceFeedKeeper,
		auctionKeeper,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, &nftKeeper, &bankKeeper, auctionKeeper, ctx
}
