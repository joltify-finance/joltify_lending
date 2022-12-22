package v0_16

import (
	v015cdp "github.com/joltify-finance/joltify_lending/x/third_party/cdp/legacy/v0_15"
	"github.com/joltify-finance/joltify_lending/x/third_party/cdp/types"
)

func migrateParams(params v015cdp.Params) types.Params {
	// migrate collateral params
	collateralParams := make(types.CollateralParams, len(params.CollateralParams))
	for i, cp := range params.CollateralParams {
		collateralParams[i] = types.CollateralParam{
			Denom:                            cp.Denom,
			Type:                             cp.Type,
			LiquidationRatio:                 cp.LiquidationRatio,
			DebtLimit:                        cp.DebtLimit,
			StabilityFee:                     cp.StabilityFee,
			AuctionSize:                      cp.AuctionSize,
			LiquidationPenalty:               cp.LiquidationPenalty,
			SpotMarketID:                     cp.SpotMarketID,
			LiquidationMarketID:              cp.LiquidationMarketID,
			KeeperRewardPercentage:           cp.KeeperRewardPercentage,
			CheckCollateralizationIndexCount: cp.CheckCollateralizationIndexCount,
			ConversionFactor:                 cp.ConversionFactor,
		}
	}

	return types.Params{
		CollateralParams: collateralParams,
		DebtParam: types.DebtParam{
			Denom:            params.DebtParam.Denom,
			ReferenceAsset:   params.DebtParam.ReferenceAsset,
			ConversionFactor: params.DebtParam.ConversionFactor,
			DebtFloor:        params.DebtParam.DebtFloor,
		},
		GlobalDebtLimit:         params.GlobalDebtLimit,
		SurplusAuctionThreshold: params.SurplusAuctionThreshold,
		SurplusAuctionLot:       params.SurplusAuctionLot,
		DebtAuctionThreshold:    params.DebtAuctionThreshold,
		DebtAuctionLot:          params.DebtAuctionLot,
		CircuitBreaker:          params.CircuitBreaker,
	}
}

func migrateCDPs(oldCDPs v015cdp.CDPs) types.CDPs {
	cdps := make(types.CDPs, len(oldCDPs))
	for i, cdp := range oldCDPs {
		cdps[i] = types.CDP{
			ID:              cdp.ID,
			Owner:           cdp.Owner,
			Type:            cdp.Type,
			Collateral:      cdp.Collateral,
			Principal:       cdp.Principal,
			AccumulatedFees: cdp.AccumulatedFees,
			FeesUpdated:     cdp.FeesUpdated,
			InterestFactor:  cdp.InterestFactor,
		}
	}
	return cdps
}

func migrateDeposits(oldDeposits v015cdp.Deposits) types.Deposits {
	deposits := make(types.Deposits, len(oldDeposits))
	for i, deposit := range oldDeposits {
		deposits[i] = types.Deposit{
			CdpID:     deposit.CdpID,
			Depositor: deposit.Depositor,
			Amount:    deposit.Amount,
		}
	}
	return deposits
}

func migratePrevAccTimes(oldPrevAccTimes v015cdp.GenesisAccumulationTimes) types.GenesisAccumulationTimes {
	prevAccTimes := make(types.GenesisAccumulationTimes, len(oldPrevAccTimes))
	for i, prevAccTime := range oldPrevAccTimes {
		prevAccTimes[i] = types.GenesisAccumulationTime{
			CollateralType:           prevAccTime.CollateralType,
			PreviousAccumulationTime: prevAccTime.PreviousAccumulationTime,
			InterestFactor:           prevAccTime.InterestFactor,
		}
	}
	return prevAccTimes
}

func migrateTotalPrincipals(oldTotalPrincipals v015cdp.GenesisTotalPrincipals) types.GenesisTotalPrincipals {
	totalPrincipals := make(types.GenesisTotalPrincipals, len(oldTotalPrincipals))
	for i, tp := range oldTotalPrincipals {
		totalPrincipals[i] = types.GenesisTotalPrincipal{
			CollateralType: tp.CollateralType,
			TotalPrincipal: tp.TotalPrincipal,
		}
	}
	return totalPrincipals
}

// Migrate converts v0.15 cdp state and returns it in v0.16 format
func Migrate(oldState v015cdp.GenesisState) *types.GenesisState {
	return &types.GenesisState{
		Params:                    migrateParams(oldState.Params),
		CDPs:                      migrateCDPs(oldState.CDPs),
		Deposits:                  migrateDeposits(oldState.Deposits),
		StartingCdpID:             oldState.StartingCdpID,
		DebtDenom:                 oldState.DebtDenom,
		GovDenom:                  oldState.GovDenom,
		PreviousAccumulationTimes: migratePrevAccTimes(oldState.PreviousAccumulationTimes),
		TotalPrincipals:           migrateTotalPrincipals(oldState.TotalPrincipals),
	}
}
