package keeper

import (
	"context"
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	kycmoduletypes "github.com/joltify-finance/joltify_lending/x/kyc/types"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/joltify-finance/joltify_lending/x/spv/keeper"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

type mockKycKeeper struct{}

func (m mockKycKeeper) GetProjects(ctx sdk.Context) (projectsInfo []*kycmoduletypes.ProjectInfo) {

	b := kycmoduletypes.BasicInfo{
		"This is the test info",
		"empty",
		"ABC",
		"ABC123",
		[]byte("reserved"),
		"This is the test info",
	}

	acc, _ := sdk.AccAddressFromBech32("jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0")
	pi1 := kycmoduletypes.ProjectInfo{
		Index:               1,
		SPVName:             "defaultSPV",
		ProjectOwner:        acc,
		BasicInfo:           &b,
		ProjectLength:       31536000, //1 year
		PayFreq:             "15768000",
		BaseApy:             sdk.NewDecWithPrec(12, 2),
		ProjectTargetAmount: sdk.NewCoin("ausdc", sdk.NewIntFromUint64(100000000)),
	}

	b2 := kycmoduletypes.BasicInfo{
		"This is the test info2",
		"empty2",
		"ABC2",
		"ABC123-2",
		[]byte("reserved"),
		"This is the test info2",
	}

	pi2 := kycmoduletypes.ProjectInfo{
		Index:               2,
		SPVName:             "defaultSPV2",
		ProjectOwner:        acc,
		BasicInfo:           &b2,
		ProjectLength:       31536000, //1 year
		PayFreq:             "15768000",
		BaseApy:             sdk.NewDecWithPrec(12, 2),
		ProjectTargetAmount: sdk.NewCoin("ausdc", sdk.NewIntFromUint64(1e12)),
	}

	pi3 := kycmoduletypes.ProjectInfo{
		Index:               3,
		SPVName:             "defaultSPV3",
		ProjectOwner:        acc,
		BasicInfo:           &b2,
		ProjectLength:       31536000, //1 year
		PayFreq:             "2592000",
		BaseApy:             sdk.NewDecWithPrec(12, 2),
		ProjectTargetAmount: sdk.NewCoin("ausdc", sdk.NewIntFromUint64(1e10)),
	}

	return []*kycmoduletypes.ProjectInfo{&pi1, &pi2, &pi3}
}

func (m mockKycKeeper) QueryByWallet(goCtx context.Context, req *kycmoduletypes.QueryByWalletRequest) (*kycmoduletypes.QueryByWalletResponse, error) {

	inv := kycmoduletypes.Investor{
		InvestorId:    "1",
		WalletAddress: []string{req.Wallet},
	}

	inv2 := kycmoduletypes.Investor{InvestorId: "2", WalletAddress: []string{"jolt1kkujrm0lqeu0e5va5f6mmwk87wva0k8cmam8jq", "jolt166yyvsypvn6cwj2rc8sme4dl6v0g62hn3862kl"}}

	if req.Wallet == "jolt166yyvsypvn6cwj2rc8sme4dl6v0g62hn3862kl" || req.Wallet == "jolt1kkujrm0lqeu0e5va5f6mmwk87wva0k8cmam8jq" {
		return &kycmoduletypes.QueryByWalletResponse{Investor: &inv2}, nil
	}

	return &kycmoduletypes.QueryByWalletResponse{
		Investor: &inv,
	}, nil
}

type mockAccKeeper struct{}

func (m mockAccKeeper) GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI {
	//TODO implement me
	panic("implement me")
}

func (m mockAccKeeper) GetModuleAccount(ctx sdk.Context, name string) authtypes.ModuleAccountI {
	addr := authtypes.NewModuleAddress(types.ModuleAccount)
	baseAcc := authtypes.NewBaseAccountWithAddress(addr)
	return authtypes.NewModuleAccount(baseAcc, types.ModuleName, "mint")
}

func (m mockAccKeeper) GetModuleAddress(name string) sdk.AccAddress {
	//TODO implement me
	panic("implement me")
}

type mockNFTKeeper struct {
	classes         map[string]*nft.Class
	nfts            map[string]*nft.NFT
	nftsWithClassID map[string]*nft.NFT
}

func (m mockNFTKeeper) Burn(ctx sdk.Context, classID string, nftID string) error {
	//TODO implement me
	panic("implement me")
}

func (m mockNFTKeeper) Transfer(ctx sdk.Context, classID string, nftID string, receiver sdk.AccAddress) error {
	panic("implement me")
}

func (m mockNFTKeeper) GetTotalSupply(ctx sdk.Context, classID string) uint64 {
	//TODO implement me
	panic("implement me")
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

func (m mockNFTKeeper) GetClass(ctx sdk.Context, classID string) (nft.Class, bool) {
	r, ok := m.classes[classID]
	return *r, ok
}

type mockbankKeeper struct {
	bankData map[string]sdk.Coins
}

func (m mockbankKeeper) SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error {
	m.bankData[recipientModule] = m.bankData[recipientModule].Add(amt...)
	return nil
}

func (m mockbankKeeper) SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
	m.bankData[recipientModule] = m.bankData[recipientModule].Add(amt...)
	return nil
}

func (m mockbankKeeper) SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	m.bankData[recipientAddr.String()] = m.bankData[recipientAddr.String()].Add(amt...)
	return nil
}

func (m mockbankKeeper) GetSupply(ctx sdk.Context, denom string) sdk.Coin {
	//TODO implement me
	panic("implement me")
}

func (m mockbankKeeper) GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	return sdk.NewCoin("usdc", sdk.NewInt(0))
}

func (m mockbankKeeper) GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins {
	//TODO implement me
	panic("implement me")
}

func (m mockbankKeeper) SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins {
	//TODO implement me
	panic("implement me")
}

func (m mockbankKeeper) BurnCoins(ctx sdk.Context, name string, amt sdk.Coins) error {
	//TODO implement me
	panic("implement me")
}

func SpvKeeper(t testing.TB) (*keeper.Keeper, types.NFTKeeper, sdk.Context) {
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
	bankKeeper := mockbankKeeper{make(map[string]sdk.Coins)}

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		paramsSubspace,
		kycKeeper,
		bankKeeper,
		accKeeper,
		nftKeeper,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, &nftKeeper, ctx
}
