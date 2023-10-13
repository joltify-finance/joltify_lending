package keeper

import (
	"bytes"
	"math/big"
	"sort"
	"strconv"

	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/crypto/secp256k1"

	"github.com/tendermint/tendermint/crypto"

	errorsmod "cosmossdk.io/errors"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	abci "github.com/tendermint/tendermint/abci/types"

	vaulttypes "github.com/joltify-finance/joltify_lending/x/vault/types"
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

func (k Keeper) UpdateStakingInfo(ctx sdk.Context) {
	stakingKeeper := k.vaultStaking
	params := k.GetParams(ctx)

	allStandbyPowerHistory := k.DoGetAllStandbyPower(ctx)

	allStandbyValidators := make(map[string]bool)
	updatedValidators := make(map[string]bool)
	for _, el := range allStandbyPowerHistory {
		allStandbyValidators[el.Addr] = true
	}

	lastBlockHeight := ctx.BlockHeight() - params.BlockChurnInterval
	latestBridgeValidators, found := k.GetValidatorsByHeight(ctx, strconv.FormatUint(uint64(lastBlockHeight), 10))
	if !found {
		latestBridgeValidators = vaulttypes.Validators{}
	}

	stakingKeeper.IterateLastValidators(ctx, func(index int64, validator stakingtypes.ValidatorI) (stop bool) {
		consAddr, err := validator.GetConsAddr()
		if err != nil {
			panic("get cons should never fail")
		}

		isLastValidator := false
		// if this node is not in the last bridge validator, we should skip the standby power deduction
		for _, el := range latestBridgeValidators.AllValidators {

			ppk := packPubkey(el.Pubkey)
			if ppk == nil {
				panic("unrecognized pubkey type")
			}

			if bytes.Equal(ppk.Address(), consAddr.Bytes()) {
				isLastValidator = true
				break
			}
		}

		current, found := k.GetStandbyPower(ctx, consAddr.String())
		if !found {
			item := vaulttypes.StandbyPower{
				Addr:  consAddr.String(),
				Power: params.Power,
			}
			k.SetStandbyPower(ctx, consAddr.String(), item)
			return false
		}
		updatedValidators[consAddr.String()] = true

		// if logPower.Int64()+current.Power < consensusPower/2 {
		if current.Power < 0 {
			k.DelStandbyPower(ctx, consAddr.String())
			return false
		}
		if !isLastValidator {
			return false
		}

		current.Power -= params.Step
		k.SetStandbyPower(ctx, consAddr.String(), current)
		return false
	})
}

func (k Keeper) getEligibleValidators(ctx sdk.Context) ([]vaulttypes.ValidatorPowerInfo, error) {
	params := k.GetParams(ctx)

	boundedValidators := k.vaultStaking.GetBondedValidatorsByPower(ctx)
	var candidates []vaulttypes.ValidatorPowerInfo

	candidateDec := sdk.NewDecWithPrec(int64(len(boundedValidators)), 0)
	candidateNumDec := candidateDec.MulTruncate(params.CandidateRatio)

	candidateNum := uint32(candidateNumDec.TruncateInt64())

	for _, validator := range boundedValidators {
		logPower := new(big.Int).Sqrt(big.NewInt(validator.ConsensusPower(sdk.DefaultPowerReduction)))

		validatorWithPower := vaulttypes.ValidatorPowerInfo{
			Validator: validator,
			Power:     logPower.Int64(),
		}
		candidates = append(candidates, validatorWithPower)
	}
	// we get rotate the nodes
	for i := 0; i < len(candidates); i++ {
		consAddr, err := candidates[i].Validator.GetConsAddr()
		if err != nil {
			panic("it should never fail to get the cons addr")
		}
		standbyPower, found := k.GetStandbyPower(ctx, consAddr.String())
		if !found {
			item := vaulttypes.StandbyPower{
				Addr:  consAddr.String(),
				Power: params.Power,
			}
			k.SetStandbyPower(ctx, consAddr.String(), item)
			standbyPower = item
		}

		candidates[i].Power += standbyPower.GetPower()
	}
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].Power > candidates[j].Power
	})

	return candidates[:candidateNum], nil
}

func (k Keeper) updateValidators(ctx sdk.Context) error {
	vs, err := k.getEligibleValidators(ctx)
	if err != nil {
		return errorsmod.Wrap(vaulttypes.ErrFormat, "fail to convert the format")
	}

	stakingValidators := make([]*vaulttypes.Validator, len(vs))

	for i, el := range vs {
		key, err := el.Validator.ConsPubKey()
		if err != nil {
			return errorsmod.Wrap(vaulttypes.ErrFormat, "fail to convert the format")
		}
		v := vaulttypes.Validator{
			Pubkey: key.Bytes(),
			Power:  el.Power,
		}
		stakingValidators[i] = &v
	}

	v := vaulttypes.Validators{
		AllValidators: stakingValidators,
		Height:        ctx.BlockHeight(),
	}
	k.SetValidators(ctx, strconv.FormatInt(ctx.BlockHeight(), 10), v)
	return nil
}

func (k Keeper) NewUpdate(ctx sdk.Context) []abci.ValidatorUpdate {
	defer telemetry.ModuleMeasureSince(vaulttypes.ModuleName, ctx.BlockTime(), telemetry.MetricKeyEndBlocker)

	blockHeight := k.GetParams(ctx).BlockChurnInterval
	if ctx.BlockHeight()%blockHeight == 0 {
		k.UpdateStakingInfo(ctx)
		ctx.EventManager().EmitEvents(sdk.Events{
			sdk.NewEvent(
				vaulttypes.EventTypeCompleteChurn,
				sdk.NewAttribute(vaulttypes.AttributeValidators, "oppy_churn"),
			),
		})
		err := k.updateValidators(ctx)
		if err != nil {
			ctx.Logger().Error("error in update the validator with err %v", err)
		}
	}
	return []abci.ValidatorUpdate{}
}
