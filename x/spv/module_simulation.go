package spv

import (
	"math/rand"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/joltify-finance/joltify_lending/testutil/sample"
	spvsimulation "github.com/joltify-finance/joltify_lending/x/spv/simulation"
	"github.com/joltify-finance/joltify_lending/x/spv/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = spvsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreatePool = "op_weight_msg_create_pool" //nolint:gosec
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePool int = 100

	opWeightMsgAddInvestors = "op_weight_msg_add_investors" //nolint:gosec
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddInvestors int = 100

	opWeightMsgDeposit = "op_weight_msg_deposit" //nolint:gosec
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeposit int = 100

	opWeightMsgBorrow = "op_weight_msg_borrow" //nolint:gosec
	// TODO: Determine the simulation weight value
	defaultWeightMsgBorrow int = 100

	opWeightMsgRepayInterest = "op_weight_msg_repay_interest" //nolint:gosec
	// TODO: Determine the simulation weight value
	defaultWeightMsgRepayInterest int = 100

	opWeightMsgClaimInterest = "op_weight_msg_claim_interest" //nolint:gosec
	// TODO: Determine the simulation weight value
	defaultWeightMsgClaimInterest int = 100

	opWeightMsgUpdatePool = "op_weight_msg_update_pool" //nolint:gosec
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdatePool int = 100

	opWeightMsgActivePool = "op_weight_msg_active_pool" //nolint:gosec
	// TODO: Determine the simulation weight value
	defaultWeightMsgActivePool int = 100

	opWeightMsgPayPrincipal = "op_weight_msg_pay_principal" //nolint:gosec
	// TODO: Determine the simulation weight value
	defaultWeightMsgPayPrincipal int = 100

	opWeightMsgWithdrawPrincipal = "op_weight_msg_withdraw_principal" //nolint:gosec
	// TODO: Determine the simulation weight value
	defaultWeightMsgWithdrawPrincipal int = 100

	opWeightMsgSubmitWitdrawProposal = "op_weight_msg_submit_witdraw_proposal" //nolint:gosec
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitWitdrawProposal int = 100

	opWeightMsgTransferOwnership = "op_weight_msg_transfer_ownership" //nolint:gosec
	// TODO: Determine the simulation weight value
	defaultWeightMsgTransferOwnership int = 100

	opWeightMsgLiquidate = "op_weight_msg_liquidate" //nolint:gosec
	// TODO: Determine the simulation weight value
	defaultWeightMsgLiquidate int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	spvGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&spvGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalMsg {
	return nil
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

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

	var weightMsgUpdatePool int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdatePool, &weightMsgUpdatePool, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePool = defaultWeightMsgUpdatePool
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePool,
		spvsimulation.SimulateMsgUpdatePool(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgActivePool int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgActivePool, &weightMsgActivePool, nil,
		func(_ *rand.Rand) {
			weightMsgActivePool = defaultWeightMsgActivePool
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgActivePool,
		spvsimulation.SimulateMsgActivePool(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgPayPrincipal int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgPayPrincipal, &weightMsgPayPrincipal, nil,
		func(_ *rand.Rand) {
			weightMsgPayPrincipal = defaultWeightMsgPayPrincipal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPayPrincipal,
		spvsimulation.SimulateMsgPayPrincipal(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgWithdrawPrincipal int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgWithdrawPrincipal, &weightMsgWithdrawPrincipal, nil,
		func(_ *rand.Rand) {
			weightMsgWithdrawPrincipal = defaultWeightMsgWithdrawPrincipal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWithdrawPrincipal,
		spvsimulation.SimulateMsgWithdrawPrincipal(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSubmitWitdrawProposal int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSubmitWitdrawProposal, &weightMsgSubmitWitdrawProposal, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitWitdrawProposal = defaultWeightMsgSubmitWitdrawProposal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitWitdrawProposal,
		spvsimulation.SimulateMsgSubmitWitdrawProposal(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgTransferOwnership int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgTransferOwnership, &weightMsgTransferOwnership, nil,
		func(_ *rand.Rand) {
			weightMsgTransferOwnership = defaultWeightMsgTransferOwnership
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTransferOwnership,
		spvsimulation.SimulateMsgTransferOwnership(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgLiquidate int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgLiquidate, &weightMsgLiquidate, nil,
		func(_ *rand.Rand) {
			weightMsgLiquidate = defaultWeightMsgLiquidate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgLiquidate,
		spvsimulation.SimulateMsgLiquidate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
