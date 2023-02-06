package keeper

import (
	"fmt"
	tmlog "github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	jolt "github.com/joltify-finance/joltify_lending/app"
)

type testVaultStaking struct{}

var _ testVaultStaking

func (t testVaultStaking) GetHistoricalInfo(ctx sdk.Context, height int64) (stakingtypes.HistoricalInfo, bool) {
	operatorStr := "joltval1f0atl7egduue8a07j42hyklct0sqa68wyh02h6"
	operator, err := sdk.ValAddressFromBech32(operatorStr)
	if err != nil {
		return stakingtypes.HistoricalInfo{}, false
	}

	sk := ed25519.GenPrivKey()
	desc := stakingtypes.NewDescription("tester", "testId", "www.test.com", "aaa", "aaa")
	testValidator, err := stakingtypes.NewValidator(operator, sk.PubKey(), desc)
	if err != nil {
		return stakingtypes.HistoricalInfo{}, false
	}
	historicalInfo := stakingtypes.HistoricalInfo{
		Valset: []stakingtypes.Validator{testValidator},
	}
	return historicalInfo, true
}

func (t testVaultStaking) IterateLastValidators(context sdk.Context, f func(index int64, validator stakingtypes.ValidatorI) (stop bool)) {
	panic("implement me")
}

func (t testVaultStaking) IterateValidators(context sdk.Context, f func(index int64, validator stakingtypes.ValidatorI) (stop bool)) {
	panic("implement me")
}

func (t testVaultStaking) ValidatorsPowerStoreIterator(ctx sdk.Context) sdk.Iterator {
	panic("implement me")
}

func (t testVaultStaking) UnbondAllMatureValidators(ctx sdk.Context) {
	panic("implement me")
}

func (t testVaultStaking) DequeueAllMatureUBDQueue(ctx sdk.Context, currTime time.Time) (matureUnbonds []stakingtypes.DVPair) {
	panic("implement me")
}

func (t testVaultStaking) CompleteRedelegation(ctx sdk.Context, delAddr sdk.AccAddress, valSrcAddr, valDstAddr sdk.ValAddress) (sdk.Coins, error) {
	panic("implement me")
}

func (t testVaultStaking) GetParams(ctx sdk.Context) stakingtypes.Params {
	return stakingtypes.Params{
		MaxValidators: 6,
	}
}

func (t testVaultStaking) LastValidatorsIterator(ctx sdk.Context) (iterator sdk.Iterator) {
	panic("implement me")
}

func (t testVaultStaking) DeleteValidatorByPowerIndex(ctx sdk.Context, validator stakingtypes.Validator) {
	panic("implement me")
}

func (t testVaultStaking) SetValidator(ctx sdk.Context, validator stakingtypes.Validator) {
	panic("implement me")
}

func (t testVaultStaking) SetValidatorByPowerIndex(ctx sdk.Context, validator stakingtypes.Validator) {
	panic("implement me")
}

func (t testVaultStaking) AfterValidatorBonded(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
	panic("implement me")
}

func (t testVaultStaking) GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator stakingtypes.Validator, found bool) {
	panic("implement me")
}

func (t testVaultStaking) InsertUnbondingValidatorQueue(ctx sdk.Context, val stakingtypes.Validator) {
	panic("implement me")
}

func (t testVaultStaking) DeleteLastValidatorPower(ctx sdk.Context, operator sdk.ValAddress) {
	panic("implement me")
}

func (t testVaultStaking) SetLastTotalPower(ctx sdk.Context, power sdk.Int) {
	panic("implement me")
}

func (t testVaultStaking) BondDenom(ctx sdk.Context) (res string) {
	panic("implement me")
}

// setup the general vault app
func SetupVaultApp(t testing.TB) (*jolt.TestApp, sdk.Context) {
	logger := tmlog.TestingLogger()
	tApp := jolt.NewTestApp(logger, t.TempDir())
	fmt.Printf("3333333333333333333333333#@@@@@@@@@@@@@@@@@@@@s\n")
	tApp.InitializeFromGenesisStates()
	ctx := tApp.App.NewContext(false, tmproto.Header{Height: 100, Time: time.Now().UTC()})
	params := tApp.GetStakingKeeper().GetParams(ctx)
	params.MaxValidators = 3
	tApp.GetStakingKeeper().SetParams(ctx, params)
	return &tApp, ctx

}
