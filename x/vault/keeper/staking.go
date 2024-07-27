package keeper

import (
	"context"
	"sort"
	"strconv"

	errorsmod "cosmossdk.io/errors"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	vaulttypes "github.com/joltify-finance/joltify_lending/x/vault/types"
)

func (k Keeper) UpdateStakingInfo(ctx context.Context) {
	stakingKeeper := k.vaultStaking
	params := k.GetParams(ctx)

	allStandbyPowerHistory := k.DoGetAllStandbyPower(ctx)

	existValidators := make(map[string]bool)
	updatedValidators := make(map[string]bool)
	for _, el := range allStandbyPowerHistory {
		existValidators[el.Addr] = true
	}

	stakingKeeper.IterateLastValidators(ctx, func(index int64, validator stakingtypes.ValidatorI) (stop bool) {
		consAddr, err := validator.GetConsAddr()
		if err != nil {
			panic("get cons should never fail")
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
		if current.Power < 0 {
			k.DelStandbyPower(ctx, consAddr.String())
			return false
		}
		current.Power -= params.Step
		k.SetStandbyPower(ctx, consAddr.String(), current)
		return false
	})

	for key := range existValidators {
		_, exist := updatedValidators[key]
		if !exist {
			k.DelStandbyPower(ctx, key)
		}
	}
}

func (k Keeper) getEligibleValidators(ctx context.Context) ([]vaulttypes.ValidatorPowerInfo, error) {
	params := k.GetParams(ctx)

	boundedValidators := k.vaultStaking.GetBondedValidatorsByPower(ctx)
	var candidates []vaulttypes.ValidatorPowerInfo

	candidateDec := sdkmath.LegacyNewDecWithPrec(int64(len(boundedValidators)), 0)
	candidateNumDec := candidateDec.MulTruncate(params.CandidateRatio)

	candidateNum := uint32(candidateNumDec.TruncateInt64())

	for _, validator := range boundedValidators {
		validatorWithPower := vaulttypes.ValidatorPowerInfo{
			Validator: validator,
			Power:     validator.ConsensusPower(sdk.DefaultPowerReduction),
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

func (k Keeper) updateValidators(ctx context.Context) error {
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

func (k Keeper) NewUpdate(ctx context.Context) []abci.ValidatorUpdate {
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
