package jolt

import (
	"context"
	"encoding/json"

	cli2 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/client/cli"
	"github.com/joltify-finance/joltify_lending/x/third_party/jolt/client/rest"
	keeper2 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/keeper"
	types2 "github.com/joltify-finance/joltify_lending/x/third_party/jolt/types"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	abci "github.com/tendermint/tendermint/abci/types"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
	// _ module.AppModuleSimulation = AppModule{}
)

// AppModuleBasic app module basics object
type AppModuleBasic struct{}

// Name get module name
func (AppModuleBasic) Name() string {
	return types2.ModuleName
}

// RegisterLegacyAminoCodec register module codec
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types2.RegisterLegacyAminoCodec(cdc)
}

// DefaultGenesis returns default genesis state as raw bytes for the jolt
// module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	gs := types2.DefaultGenesisState()
	return cdc.MustMarshalJSON(&gs)
}

// ValidateGenesis performs genesis state validation for the jolt module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	var gs types2.GenesisState
	err := cdc.UnmarshalJSON(bz, &gs)
	if err != nil {
		return err
	}
	return gs.Validate()
}

// RegisterInterfaces implements InterfaceModule.RegisterInterfaces
func (a AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	types2.RegisterInterfaces(registry)
}

// RegisterRESTRoutes registers REST routes for the jolt module.
func (a AppModuleBasic) RegisterRESTRoutes(clientCtx client.Context, rtr *mux.Router) {
	rest.RegisterRoutes(clientCtx, rtr)
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the gov module.
func (a AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	if err := types2.RegisterQueryHandlerClient(context.Background(), mux, types2.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}
}

// LegacyQuerierHandler returns sdk.Querier.
func (am AppModule) LegacyQuerierHandler(legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return keeper2.NewQuerier(am.keeper, legacyQuerierCdc)
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 {
	return 1
}

// GetTxCmd returns the root tx command for the jolt module.
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli2.GetTxCmd()
}

// GetQueryCmd returns no root query command for the jolt module.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli2.GetQueryCmd()
}

//____________________________________________________________________________

// AppModule app module type
type AppModule struct {
	AppModuleBasic

	keeper          keeper2.Keeper
	accountKeeper   types2.AccountKeeper
	bankKeeper      types2.BankKeeper
	pricefeedKeeper types2.PricefeedKeeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(keeper keeper2.Keeper, accountKeeper types2.AccountKeeper, bankKeeper types2.BankKeeper, pricefeedKeeper types2.PricefeedKeeper) AppModule {
	return AppModule{
		AppModuleBasic:  AppModuleBasic{},
		keeper:          keeper,
		accountKeeper:   accountKeeper,
		bankKeeper:      bankKeeper,
		pricefeedKeeper: pricefeedKeeper,
	}
}

// Name module name
func (AppModule) Name() string {
	return types2.ModuleName
}

// RegisterInvariants register module invariants
func (AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// Route returns the message routing key for the jolt module.
func (am AppModule) Route() sdk.Route {
	return sdk.Route{}
}

// QuerierRoute returns the jolt module's querier route name.
func (AppModule) QuerierRoute() string {
	return types2.QuerierRoute
}

// RegisterServices registers module services.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types2.RegisterMsgServer(cfg.MsgServer(), keeper2.NewMsgServerImpl(am.keeper))
	types2.RegisterQueryServer(cfg.QueryServer(), keeper2.NewQueryServerImpl(am.keeper, am.accountKeeper, am.bankKeeper))
}

// InitGenesis performs genesis initialization for the jolt module. It returns
// no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, gs json.RawMessage) []abci.ValidatorUpdate {
	var genState types2.GenesisState
	// Initialize global index to index in genesis state
	cdc.MustUnmarshalJSON(gs, &genState)

	InitGenesis(ctx, am.keeper, am.accountKeeper, genState)
	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the exported genesis state as raw bytes for the jolt
// module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	gs := ExportGenesis(ctx, am.keeper)
	return cdc.MustMarshalJSON(&gs)
}

// BeginBlock module begin-block
func (am AppModule) BeginBlock(ctx sdk.Context, _ abci.RequestBeginBlock) {
	BeginBlocker(ctx, am.keeper)
}

// EndBlock module end-block
func (am AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

//____________________________________________________________________________

// // GenerateGenesisState creates a randomized GenState of the jolt module
// func (AppModuleBasic) GenerateGenesisState(simState *module.SimulationState) {
// 	simulation.RandomizedGenState(simState)
// }

// // ProposalContents doesn't return any content functions for governance proposals.
// func (AppModuleBasic) ProposalContents(_ module.SimulationState) []sim.WeightedProposalContent {
// 	return nil
// }

// // RandomizedParams returns nil because jolt has no params.
// func (AppModuleBasic) RandomizedParams(r *rand.Rand) []sim.ParamChange {
// 	return simulation.ParamChanges(r)
// }

// // RegisterStoreDecoder registers a decoder for jolt module's types
// func (AppModuleBasic) RegisterStoreDecoder(sdr sdk.StoreDecoderRegistry) {
// 	sdr[types.StoreKey] = simulation.DecodeStore
// }

// // WeightedOperations returns the all the jolt module operations with their respective weights.
// func (am AppModule) WeightedOperations(simState module.SimulationState) []sim.WeightedOperation {
// 	return nil
// }
