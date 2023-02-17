package keeper

import (
	"context"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	kycmoduletypes "github.com/joltify-finance/joltify_lending/x/kyc/types"
	"testing"

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
	//TODO implement me
	panic("implement me")
}

func (m mockKycKeeper) QueryByWallet(goCtx context.Context, req *kycmoduletypes.QueryByWalletRequest) (*kycmoduletypes.QueryByWalletResponse, error) {
	//TODO implement me
	panic("implement me")
}

type mockAccKeeper struct{}

func (m mockAccKeeper) GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI {
	//TODO implement me
	panic("implement me")
}

func (m mockAccKeeper) GetModuleAccount(ctx sdk.Context, name string) authtypes.ModuleAccountI {
	//TODO implement me
	panic("implement me")
}

func (m mockAccKeeper) GetModuleAddress(name string) sdk.AccAddress {
	//TODO implement me
	panic("implement me")
}

type mockNFTKeeper struct{}

func (m mockNFTKeeper) Mint(ctx sdk.Context, nft nft.NFT, receiver sdk.AccAddress) error {
	//TODO implement me
	panic("implement me")
}

func (m mockNFTKeeper) SaveClass(ctx sdk.Context, class nft.Class) error {
	//TODO implement me
	panic("implement me")
}

func (m mockNFTKeeper) GetClass(ctx sdk.Context, classID string) (nft.Class, bool) {
	//TODO implement me
	panic("implement me")
}

type mockbankKeeper struct{}

func (m mockbankKeeper) SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error {
	//TODO implement me
	panic("implement me")
}

func (m mockbankKeeper) SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
	//TODO implement me
	panic("implement me")
}

func (m mockbankKeeper) SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	//TODO implement me
	panic("implement me")
}

func (m mockbankKeeper) GetSupply(ctx sdk.Context, denom string) sdk.Coin {
	//TODO implement me
	panic("implement me")
}

func (m mockbankKeeper) GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	//TODO implement me
	panic("implement me")
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

func SpvKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
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
	nftKeeper := mockNFTKeeper{}
	bankKeeper := mockbankKeeper{}

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

	return k, ctx
}
