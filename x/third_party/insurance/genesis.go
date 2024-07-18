package insurance

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/joltify-finance/joltify_lending/x/third_party/insurance/keeper"
	"github.com/joltify-finance/joltify_lending/x/third_party/insurance/types"
)

// InitGenesis init state of module
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) {
	k.SetParams(ctx, data.Params)
	for i := range data.InsuranceFunds {
		k.SetInsuranceFund(ctx, &data.InsuranceFunds[i])
	}
	for _, schedule := range data.RedemptionSchedule {
		k.SetRedemptionSchedule(ctx, schedule)
	}
	k.SetNextShareDenomId(ctx, data.NextShareDenomId)
	k.SetNextRedemptionScheduleId(ctx, data.NextRedemptionScheduleId)
}

// ExportGenesis export the state of module
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Params:                   k.GetParams(ctx),
		InsuranceFunds:           k.GetAllInsuranceFunds(ctx),
		RedemptionSchedule:       k.GetAllInsuranceFundRedemptions(ctx),
		NextShareDenomId:         k.ExportNextShareDenomId(ctx),
		NextRedemptionScheduleId: k.ExportNextRedemptionScheduleId(ctx),
	}
}
