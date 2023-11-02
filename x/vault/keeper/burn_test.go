package keeper_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	joltapp "github.com/joltify-finance/joltify_lending/app"
	keepertest "github.com/joltify-finance/joltify_lending/testutil/keeper"
	"github.com/joltify-finance/joltify_lending/testutil/simapp"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
	"github.com/stretchr/testify/assert"
)

func prepare(t *testing.T) (*joltapp.TestApp, sdk.Context, []sdk.AccAddress) {
	app, ctx := keepertest.SetupVaultApp(t)

	testValidators, creators := keepertest.GenerateNValidators(t, 4)

	for _, el := range creators {
		err := simapp.FundAccount(app.GetBankKeeper(), ctx, el, sdk.Coins{sdk.Coin{Denom: "mock", Amount: sdk.NewInt(1000)}})
		assert.NoError(t, err)
	}

	f1Before := app.GetBankKeeper().GetAllBalances(ctx, creators[0])
	assert.Equal(t, f1Before.IsEqual(sdk.Coins{sdk.Coin{Denom: "mock", Amount: sdk.NewInt(1000)}}), true)
	f2Before := app.GetBankKeeper().GetAllBalances(ctx, creators[1])
	assert.Equal(t, f2Before.IsEqual(sdk.Coins{sdk.Coin{Denom: "mock", Amount: sdk.NewInt(1000)}}), true)

	p1 := types.PoolProposal{PoolAddr: creators[0], Nodes: []sdk.AccAddress{creators[0], creators[1], creators[2]}}
	p2 := types.PoolProposal{PoolAddr: creators[1], Nodes: []sdk.AccAddress{creators[0], creators[1], creators[2]}}
	params := app.VaultKeeper.GetParams(ctx)
	params.BlockChurnInterval = 20
	app.VaultKeeper.SetParams(ctx, params)

	position1 := ctx.BlockHeight() - params.BlockChurnInterval + 1
	position2 := ctx.BlockHeight() - params.BlockChurnInterval*2 + 1

	createPool := types.CreatePool{
		BlockHeight: strconv.FormatInt(position1, 10),
		Validators:  testValidators,
		Proposal:    []*types.PoolProposal{&p1, &p1, &p1},
	}
	app.VaultKeeper.SetCreatePool(ctx, createPool)
	app.VaultKeeper.UpdateLastTwoPool(ctx, createPool)

	createPool = types.CreatePool{
		BlockHeight: strconv.FormatInt(position2, 10),
		Validators:  testValidators,
		Proposal:    []*types.PoolProposal{&p2, &p2, &p2},
	}
	app.VaultKeeper.SetCreatePool(ctx, createPool)
	app.VaultKeeper.UpdateLastTwoPool(ctx, createPool)

	items := make([]types.CreatePool, 2)
	for i := range items {
		sk := ed25519.GenPrivKey()
		poolProposal := types.PoolProposal{
			PoolAddr: creators[i].Bytes(),
		}
		err := simapp.FundAccount(app.GetBankKeeper(), ctx, sk.PubKey().Address().Bytes(), sdk.Coins{sdk.Coin{Denom: "mock", Amount: sdk.NewInt(1000)}})
		assert.NoError(t, err)
		items[i].Proposal = []*types.PoolProposal{&poolProposal}
		items[i].BlockHeight = fmt.Sprintf("%d", i)
		app.VaultKeeper.SetCreatePool(ctx, items[i])
		app.VaultKeeper.UpdateLastTwoPool(ctx, items[i])
	}
	return app, ctx, creators
}

func TestProcessAccountLeft(t *testing.T) {
	app, ctx, _ := prepare(t)
	err := app.GetBankKeeper().MintCoins(ctx, types.ModuleName, sdk.Coins{sdk.Coin{Denom: "mock", Amount: sdk.NewInt(1000)}})
	assert.NoError(t, err)
	app.VaultKeeper.SetStoreFeeAmount(ctx, sdk.NewCoins(sdk.NewCoin("mock", sdk.NewInt(12))))
	app.VaultKeeper.ProcessAccountLeft(ctx)
	fee, ok := app.VaultKeeper.GetFeeAmount(ctx, "mock")
	assert.True(t, ok)
	assert.Equal(t, fee.Amount, sdk.NewInt(0))

	balance := app.GetModuleAccountBalance(ctx, types.ModuleName, "mock")
	assert.True(t, balance.IsZero())
}

func TestProcessAccountLeftWithAccountLessThanFee(t *testing.T) {
	app, ctx, _ := prepare(t)
	err := app.GetBankKeeper().MintCoins(ctx, types.ModuleName, sdk.Coins{sdk.Coin{Denom: "mock", Amount: sdk.NewInt(1000)}})
	assert.NoError(t, err)
	app.VaultKeeper.SetStoreFeeAmount(ctx, sdk.NewCoins(sdk.NewCoin("mock", sdk.NewInt(120000))))
	app.VaultKeeper.ProcessAccountLeft(ctx)
	fee, ok := app.VaultKeeper.GetFeeAmount(ctx, "mock")
	assert.True(t, ok)
	assert.Equal(t, fee.Amount, sdk.NewInt(120000))
	allCoins := app.GetBankKeeper().GetAllBalances(ctx, authtypes.NewModuleAddress(types.ModuleName))
	assert.Equal(t, len(allCoins), 0)
}
