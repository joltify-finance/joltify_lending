package keeper_test

import (
	"strconv"
	"testing"

	app2 "github.com/joltify-finance/joltify_lending/app"
	keepertest "github.com/joltify-finance/joltify_lending/testutil/keeper"
	"github.com/stretchr/testify/assert"
)

func TestUpdateStakingInfo(t *testing.T) {
	app2.SetSDKConfig()

	app, ctx := keepertest.SetupVaultApp(t)
	k := app.VaultKeeper
	k.UpdateStakingInfo(ctx)
	currentStandbyPower := k.DoGetAllStandbyPower(ctx)
	for _, el := range currentStandbyPower {
		assert.Equal(t, 10000, int(el.Power))
	}

	k.UpdateStakingInfo(ctx)
	currentStandbyPower = k.DoGetAllStandbyPower(ctx)

	for _, el := range currentStandbyPower {
		assert.Equal(t, 9000, int(el.Power))
	}

	k.UpdateStakingInfo(ctx)
	currentStandbyPower = k.DoGetAllStandbyPower(ctx)
	for _, el := range currentStandbyPower {
		assert.Equal(t, 8000, int(el.Power))
	}

	// now we delete one node, and it standby power should be reset
	k.DelStandbyPower(ctx, currentStandbyPower[0].Addr)
	k.UpdateStakingInfo(ctx)
	currentStandbyPower = k.DoGetAllStandbyPower(ctx)
	for i, el := range currentStandbyPower {
		if i == 0 {
			assert.Equal(t, 10000, int(el.Power))
			continue
		}
		assert.Equal(t, 7000, int(el.Power))
	}

	// now we set v1 standby pwoer <0, it should be deleted
	target := currentStandbyPower[1]
	v0Addr := currentStandbyPower[0].Addr
	v1Addr := currentStandbyPower[1].Addr

	target.Power = -1
	k.SetStandbyPower(ctx, currentStandbyPower[1].Addr, target)
	k.UpdateStakingInfo(ctx)
	currentStandbyPower = k.DoGetAllStandbyPower(ctx)
	for _, el := range currentStandbyPower {
		if el.Addr == v0Addr {
			assert.Equal(t, 9000, int(el.Power))
			continue
		}
		assert.Equal(t, 6000, int(el.Power))
	}
	assert.Equal(t, 3, len(currentStandbyPower))
	k.UpdateStakingInfo(ctx)
	currentStandbyPower = k.DoGetAllStandbyPower(ctx)
	for _, el := range currentStandbyPower {
		if el.Addr == v1Addr {
			assert.Equal(t, 10000, int(el.Power))
			continue
		}
		if el.Addr == v0Addr {
			assert.Equal(t, 8000, int(el.Power))
			continue
		}
		assert.Equal(t, 5000, int(el.Power))
	}
}

func TestGetEligibleValidator(t *testing.T) {
	app2.SetSDKConfig()

	app, ctx := keepertest.SetupVaultApp(t)
	k := app.VaultKeeper

	blockHeight := k.GetParams(ctx).BlockChurnInterval
	defaultParam := k.GetParams(ctx)
	defaultParam.Step = 100000
	defaultParam.Power = 100000
	k.SetParams(ctx, defaultParam)
	ctx = ctx.WithBlockHeight(blockHeight)
	k.NewUpdate(ctx)
	result, found := k.GetValidatorsByHeight(ctx, strconv.Itoa(int(ctx.BlockHeight())))
	assert.True(t, found)
	assert.Equal(t, 2, len(result.GetAllValidators()))
	assert.Equal(t, 700000, int(result.GetAllValidators()[0].Power))
	assert.Equal(t, 500000, int(result.GetAllValidators()[1].Power))

	ctx = ctx.WithBlockHeight(blockHeight * 2)
	k.NewUpdate(ctx)
	result, found = k.GetValidatorsByHeight(ctx, strconv.Itoa(int(ctx.BlockHeight())))
	assert.True(t, found)
	assert.Equal(t, 2, len(result.GetAllValidators()))
	assert.Equal(t, 600000, int(result.GetAllValidators()[0].Power))
	assert.Equal(t, 400000, int(result.GetAllValidators()[1].Power))

	ctx = ctx.WithBlockHeight(blockHeight * 3)
	k.NewUpdate(ctx)
	result, found = k.GetValidatorsByHeight(ctx, strconv.Itoa(int(ctx.BlockHeight())))
	assert.True(t, found)
	assert.Equal(t, 2, len(result.GetAllValidators()))
	assert.Equal(t, 500000, int(result.GetAllValidators()[0].Power))
	assert.Equal(t, 320000, int(result.GetAllValidators()[1].Power))

}
