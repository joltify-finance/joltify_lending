package spv

import (
	"math/rand"

	"github.com/joltify-finance/joltify_lending/testutil/sample"
	spvsimulation "github.com/joltify-finance/joltify_lending/x/spv/simulation"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = spvsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
    opWeightMsgCreatePool = "op_weight_msg_create_pool"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePool int = 100

	opWeightMsgAddInvestors = "op_weight_msg_add_investors"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddInvestors int = 100

	opWeightMsgDeposit = "op_weight_msg_deposit"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeposit int = 100

	opWeightMsgBorrow = "op_weight_msg_borrow"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBorrow int = 100

	opWeightMsgRepayInterest = "op_weight_msg_repay_interest"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRepayInterest int = 100

	opWeightMsgClaimInterest = "op_weight_msg_claim_interest"
	// TODO: Determine the simulation weight value
	defaultWeightMsgClaimInterest int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	spvGenesis := types.GenesisState{
		Params:	types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&spvGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	
	return []simtypes.ParamChange{
	}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreatePool int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreatePool, &weightMsgCreatePool, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePool = defaultWeightMsgCreatePool
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePool,
		spvsimulation.SimulateMsgCreatePool(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddInvestors int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddInvestors, &weightMsgAddInvestors, nil,
		func(_ *rand.Rand) {
			weightMsgAddInvestors = defaultWeightMsgAddInvestors
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddInvestors,
		spvsimulation.SimulateMsgAddInvestors(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeposit int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeposit, &weightMsgDeposit, nil,
		func(_ *rand.Rand) {
			weightMsgDeposit = defaultWeightMsgDeposit
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeposit,
		spvsimulation.SimulateMsgDeposit(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBorrow int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBorrow, &weightMsgBorrow, nil,
		func(_ *rand.Rand) {
			weightMsgBorrow = defaultWeightMsgBorrow
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBorrow,
		spvsimulation.SimulateMsgBorrow(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRepayInterest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRepayInterest, &weightMsgRepayInterest, nil,
		func(_ *rand.Rand) {
			weightMsgRepayInterest = defaultWeightMsgRepayInterest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRepayInterest,
		spvsimulation.SimulateMsgRepayInterest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgClaimInterest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgClaimInterest, &weightMsgClaimInterest, nil,
		func(_ *rand.Rand) {
			weightMsgClaimInterest = defaultWeightMsgClaimInterest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgClaimInterest,
		spvsimulation.SimulateMsgClaimInterest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
