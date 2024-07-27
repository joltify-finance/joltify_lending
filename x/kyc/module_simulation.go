package kyc

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/joltify-finance/joltify_lending/testutil/sample"
	kycsimulation "github.com/joltify-finance/joltify_lending/x/kyc/simulation"
	"github.com/joltify-finance/joltify_lending/x/kyc/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = kycsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgUploadInvestor = "op_weight_msg_upload_investor" //nolint:gosec
	// TODO: Determine the simulation weight value
	defaultWeightMsgUploadInvestor int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	kycGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&kycGenesis)
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

	var weightMsgUploadInvestor int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUploadInvestor, &weightMsgUploadInvestor, nil,
		func(_ *rand.Rand) {
			weightMsgUploadInvestor = defaultWeightMsgUploadInvestor
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUploadInvestor,
		kycsimulation.SimulateMsgUploadInvestor(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
