package keeper_test

import (
	"bytes"
	"sort"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/bech32"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	vaulttypes "github.com/joltify-finance/joltify_lending/x/vault/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"

	app2 "github.com/joltify-finance/joltify_lending/app"
	keepertest "github.com/joltify-finance/joltify_lending/testutil/keeper"
	"github.com/stretchr/testify/assert"
)

func packPubkey(b []byte) crypto.PubKey {
	switch len(b) {
	case ed25519.PubKeySize:
		pk := make(ed25519.PubKey, ed25519.PubKeySize)
		copy(pk, b)
		return pk
	case secp256k1.PubKeySize:
		pk := make(ed25519.PubKey, ed25519.PubKeySize)
		copy(pk, b)
		return pk

	}
	return nil
}

func TestUpdateStakingInfo(t *testing.T) {
	app2.SetSDKConfig()

	app, ctx := keepertest.SetupVaultApp(t)
	k := app.VaultKeeper

	pa := k.GetParams(ctx)

	ctx = ctx.WithBlockHeight(pa.BlockChurnInterval)
	ctxFirst := ctx
	k.NewUpdate(ctx)
	currentStandbyPower := k.DoGetAllStandbyPower(ctx)
	for _, el := range currentStandbyPower {
		assert.Equal(t, 10000, int(el.Power))
	}

	ctx = ctx.WithBlockHeight(pa.BlockChurnInterval * 2)

	k.UpdateStakingInfo(ctx)
	currentStandbyPower = k.DoGetAllStandbyPower(ctx)

	latestBridgeValidators, found := k.GetValidatorsByHeight(ctx, strconv.FormatUint(uint64(ctxFirst.BlockHeight()), 10))
	if !found {
		latestBridgeValidators = vaulttypes.Validators{}
	}

	var validatorAddr [][]byte
	for _, el := range latestBridgeValidators.GetAllValidators() {
		ppk := packPubkey(el.Pubkey)
		if ppk == nil {
			panic("unrecognized pubkey type")
		}
		validatorAddr = append(validatorAddr, ppk.Address().Bytes())
	}

	// as no one is the validator, so all have the default power
	for _, el := range currentStandbyPower {
		validator := false
		for _, addr := range validatorAddr {
			_, b, err := bech32.DecodeAndConvert(el.GetAddr())
			assert.NoError(t, err)

			if bytes.Equal(b, addr) {
				assert.Equal(t, 9000, int(el.Power))
				validator = true
				break
			}
		}

		if validator {
			continue
		}
		assert.Equal(t, 10000, int(el.Power))
	}

	latestBridgeValidators, found = k.GetValidatorsByHeight(ctx, strconv.FormatUint(uint64(ctxFirst.BlockHeight()), 10))
	if !found {
		latestBridgeValidators = vaulttypes.Validators{}
	}

	validatorAddr = make([][]byte, 0)
	for _, el := range latestBridgeValidators.GetAllValidators() {
		ppk := packPubkey(el.Pubkey)
		if ppk == nil {
			panic("unrecognized pubkey type")
		}
		validatorAddr = append(validatorAddr, ppk.Address().Bytes())
	}

	assert.Equal(t, len(validatorAddr), 2)

	k.UpdateStakingInfo(ctx)
	currentStandbyPower = k.DoGetAllStandbyPower(ctx)

	for _, el := range currentStandbyPower {
		validator := false
		for _, addr := range validatorAddr {
			_, b, err := bech32.DecodeAndConvert(el.GetAddr())
			assert.NoError(t, err)
			if bytes.Equal(b, addr) {
				assert.Equal(t, 8000, int(el.Power))
				validator = true
				break
			}
		}

		if validator {
			continue
		}

		assert.Equal(t, 10000, int(el.Power))
	}

	var tryToDelete string
	for _, el := range currentStandbyPower {
		if el.Power == 10000 {
			tryToDelete = el.Addr
			break
		}
	}

	// now we delete one node, and it standby power should be reset
	k.DelStandbyPower(ctx, tryToDelete)
	k.UpdateStakingInfo(ctx)
	currentStandbyPower = k.DoGetAllStandbyPower(ctx)
	for _, el := range currentStandbyPower {
		validator := false
		for _, addr := range validatorAddr {
			_, b, err := bech32.DecodeAndConvert(el.GetAddr())
			assert.NoError(t, err)

			if bytes.Equal(b, addr) {
				assert.Equal(t, 7000, int(el.Power))
				validator = true
				break
			}
		}

		if validator {
			continue
		}

		assert.Equal(t, 10000, int(el.Power))
	}

	currentStandbyPower = k.DoGetAllStandbyPower(ctx)
	var lastValidators []string
	for _, el := range currentStandbyPower {

		validator := false
		for _, addr := range validatorAddr {
			_, b, err := bech32.DecodeAndConvert(el.GetAddr())
			assert.NoError(t, err)
			if bytes.Equal(b, addr) {
				assert.Equal(t, 7000, int(el.Power))
				validator = true
				break
			}
		}
		if validator {
			lastValidators = append(lastValidators, el.GetAddr())
			continue
		}

		assert.Equal(t, 10000, int(el.Power))
	}

	var v0AddrIndex, v1AddrIndex int
	for index, el := range currentStandbyPower {
		if el.GetAddr() == lastValidators[0] {
			v0AddrIndex = index
			continue
		}
		if el.GetAddr() == lastValidators[1] {
			v1AddrIndex = index
			continue
		}
	}

	// now we set v1 standby pwoer <0, it should be deleted
	target := currentStandbyPower[v0AddrIndex]
	v0Addr := currentStandbyPower[v0AddrIndex].Addr
	v1Addr := currentStandbyPower[v1AddrIndex].Addr

	target.Power = -1
	k.SetStandbyPower(ctx, currentStandbyPower[v0AddrIndex].Addr, target)
	k.UpdateStakingInfo(ctx)
	currentStandbyPower = k.DoGetAllStandbyPower(ctx)

	for _, el := range currentStandbyPower {
		assert.NotEqual(t, v0Addr, el.GetAddr())
		if el.Addr == v1Addr {
			assert.Equal(t, 6000, int(el.Power))
			continue
		}
		assert.Equal(t, 10000, int(el.Power))
	}
	assert.Equal(t, 3, len(currentStandbyPower))
	k.UpdateStakingInfo(ctx)
	currentStandbyPower = k.DoGetAllStandbyPower(ctx)
	for _, el := range currentStandbyPower {
		if el.Addr == v1Addr {
			assert.Equal(t, 5000, int(el.Power))
			continue
		}
		if el.Addr == v0Addr {
			assert.Equal(t, 10000, int(el.Power))
			continue
		}
		assert.Equal(t, 10000, int(el.Power))
	}
}

func TestGetEligibleValidator(t *testing.T) {
	app2.SetSDKConfig()

	// the validator default voting power
	//	 600000,400000,320000,300000

	// first run should be
	// 700000,500000,420000,400000
	// second run should be
	// 600000,420000,400000,400000
	// third run should be
	// 500000,400000,300000,300000

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

	validators1 := result.GetAllValidators()
	sort.Slice(validators1, func(i, j int) bool {
		return validators1[i].Power > validators1[j].Power
	})

	// the first time all validator has the validator voting power + standby voting power
	assert.Equal(t, 2, len(result.GetAllValidators()))
	assert.Equal(t, 700000, int(result.GetAllValidators()[0].Power))
	assert.Equal(t, 500000, int(result.GetAllValidators()[1].Power))

	ctx = ctx.WithBlockHeight(blockHeight * 2)

	k.NewUpdate(ctx)

	allStandByNodes := k.DoGetAllStandbyPower(ctx)
	sort.Slice(allStandByNodes, func(i, j int) bool {
		return allStandByNodes[i].Power > allStandByNodes[j].Power
	})

	standbyPower := make(map[string]int)
	for _, el := range allStandByNodes {
		standbyPower[el.Addr] = int(el.Power)
	}

	result, found = k.GetValidatorsByHeight(ctx, strconv.Itoa(int(ctx.BlockHeight())))
	assert.True(t, found)
	assert.Equal(t, 2, len(result.GetAllValidators()))

	validators2 := result.GetAllValidators()
	sort.Slice(validators2, func(i, j int) bool {
		return validators2[i].Power > validators2[j].Power
	})

	assert.Equal(t, 600000, int(result.GetAllValidators()[0].Power))
	assert.Equal(t, 420000, int(result.GetAllValidators()[1].Power))

	ctx = ctx.WithBlockHeight(blockHeight * 3)
	k.NewUpdate(ctx)
	result, found = k.GetValidatorsByHeight(ctx, strconv.Itoa(int(ctx.BlockHeight())))
	assert.True(t, found)

	validators3 := result.GetAllValidators()
	sort.Slice(validators3, func(i, j int) bool {
		return validators3[i].Power > validators3[j].Power
	})

	assert.Equal(t, 2, len(result.GetAllValidators()))
	assert.Equal(t, 500000, int(result.GetAllValidators()[0].Power))
	assert.Equal(t, 400000, int(result.GetAllValidators()[1].Power))

	assert.True(t, bytes.Equal(validators1[0].GetPubkey(), validators2[0].GetPubkey()))
	assert.True(t, bytes.Equal(validators3[0].GetPubkey(), validators2[0].GetPubkey()))

	assert.False(t, bytes.Equal(validators1[1].GetPubkey(), validators2[1].GetPubkey()))
	assert.False(t, bytes.Equal(validators2[1].GetPubkey(), validators3[1].GetPubkey()))
}
