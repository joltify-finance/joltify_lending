package vault

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/joltify-finance/joltify_lending/testutil/sample"
	vaultsimulation "github.com/joltify-finance/joltify_lending/x/vault/simulation"
	"github.com/joltify-finance/joltify_lending/x/vault/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = vaultsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateOutboundTx = "op_weight_msg_create_chain" //nolint:gosec
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateOutboundTx int = 100

	opWeightMsgUpdateOutboundTx = "op_weight_msg_create_chain" //nolint:gosec
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateOutboundTx int = 100

	opWeightMsgDeleteOutboundTx = "op_weight_msg_create_chain" //nolint:gosec
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteOutboundTx int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}

	s1, err := sdk.AccAddressFromBech32(sample.AccAddress())
	if err != nil {
		panic(err)
	}

	s2, err := sdk.AccAddressFromBech32(sample.AccAddress())
	if err != nil {
		panic(err)
	}

	ss1 := make(map[string]types.Proposals)
	mockCoin := sdk.Coin{Denom: "aaa", Amount: sdkmath.NewIntFromUint64(3)}
	a1 := types.Entity{Address: s1, Feecoin: []sdk.Coin{mockCoin}}
	ss1["00"] = types.Proposals{Entry: []*types.Entity{&a1}}

	ss2 := make(map[string]types.Proposals)
	mockCoin2 := sdk.Coin{Denom: "aaa", Amount: sdkmath.NewIntFromUint64(4)}
	a2 := types.Entity{Address: s2, Feecoin: []sdk.Coin{mockCoin2}}
	ss2["11"] = types.Proposals{Entry: []*types.Entity{&a2}}

	vaultGenesis := types.GenesisState{
		OutboundTxList: []types.OutboundTx{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&vaultGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalMsg {
	return nil
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateOutboundTx int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateOutboundTx, &weightMsgCreateOutboundTx, nil,
		func(_ *rand.Rand) {
			weightMsgCreateOutboundTx = defaultWeightMsgCreateOutboundTx
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateOutboundTx,
		vaultsimulation.SimulateMsgCreateOutboundTx(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateOutboundTx int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateOutboundTx, &weightMsgUpdateOutboundTx, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateOutboundTx = defaultWeightMsgUpdateOutboundTx
		},
	)

	var weightMsgDeleteOutboundTx int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteOutboundTx, &weightMsgDeleteOutboundTx, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteOutboundTx = defaultWeightMsgDeleteOutboundTx
		},
	)

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
