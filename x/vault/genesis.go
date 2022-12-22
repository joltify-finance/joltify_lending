package vault

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/vault/keeper"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the outboundTx
	for _, elem := range genState.OutboundTxList {
		k.SetOutboundTx(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the issueToken
	for _, elem := range genState.IssueTokenList {
		k.SetIssueToken(ctx, *elem)
	}

	// Set all the createPool
	for _, elem := range genState.CreatePoolList {
		k.SetCreatePool(ctx, *elem)
	}

	k.SetParams(ctx, genState.Params)

	// FIXME we assume that the first and the second pool info is for the latest two pool
	// set the latest two pool info

	if len(genState.LatestTwoPool) == 2 {
		k.GenSetLastTwoPool(ctx, genState.LatestTwoPool)
	}
	// set all the validator info
	for _, elem := range genState.ValidatorinfoList {
		k.SetValidators(ctx, strconv.FormatInt(elem.Height, 10), *elem)
	}

	// set all standbyPower info
	for _, elem := range genState.StandbypowerList {
		k.SetStandbyPower(ctx, elem.Addr, *elem)
	}

	k.SetStoreFeeAmount(ctx, genState.FeeCollectedList)

	k.SetQuotaData(ctx, genState.CoinsQuota)

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.OutboundTxList = k.GetAllOutboundTx(ctx)
	// this line is used by starport scaffolding # genesis/module/export
	// Get all issueToken
	issueTokenList := k.GetAllIssueToken(ctx)
	for _, elem := range issueTokenList {
		elem := elem
		genesis.IssueTokenList = append(genesis.IssueTokenList, &elem)
	}

	// Get all createPool
	createPoolList := k.GetAllCreatePool(ctx)
	for _, elem := range createPoolList {
		elem := elem
		genesis.CreatePoolList = append(genesis.CreatePoolList, &elem)
	}

	// Get all validator info
	validators := k.DoGetAllValidators(ctx)
	for _, elem := range validators {
		elem := elem
		genesis.ValidatorinfoList = append(genesis.ValidatorinfoList, &elem)
	}

	// set all standbypower info

	allStandbyPowerInfo := k.DoGetAllStandbyPower(ctx)
	for _, elem := range allStandbyPowerInfo {
		elem := elem
		genesis.StandbypowerList = append(genesis.StandbypowerList, &elem)
	}

	// get all fees
	fees := k.GetAllFeeAmount(ctx)
	for _, elem := range fees {
		elem := elem
		genesis.FeeCollectedList = append(genesis.FeeCollectedList, elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/export

	coinsQuota, found := k.GetQuotaData(ctx)
	if found {
		genesis.CoinsQuota = coinsQuota
	}

	lastTwoPools, found := k.GetLatestTwoPool(ctx)
	if found {
		genesis.LatestTwoPool = lastTwoPools
	}

	return genesis
}
