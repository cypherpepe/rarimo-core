package tokenmanager

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"gitlab.com/rarify-protocol/rarimo-core/testutil/sample"
	tokenmanagersimulation "gitlab.com/rarify-protocol/rarimo-core/x/tokenmanager/simulation"
	"gitlab.com/rarify-protocol/rarimo-core/x/tokenmanager/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = tokenmanagersimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateInfo = "op_weight_msg_info"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateInfo int = 100

	opWeightMsgUpdateInfo = "op_weight_msg_info"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateInfo int = 100

	opWeightMsgDeleteInfo = "op_weight_msg_info"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteInfo int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	tokenmanagerGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		InfoList: []types.Info{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&tokenmanagerGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateInfo int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateInfo, &weightMsgCreateInfo, nil,
		func(_ *rand.Rand) {
			weightMsgCreateInfo = defaultWeightMsgCreateInfo
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateInfo,
		tokenmanagersimulation.SimulateMsgCreateInfo(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateInfo int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateInfo, &weightMsgUpdateInfo, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateInfo = defaultWeightMsgUpdateInfo
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateInfo,
		tokenmanagersimulation.SimulateMsgUpdateInfo(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteInfo int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteInfo, &weightMsgDeleteInfo, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteInfo = defaultWeightMsgDeleteInfo
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteInfo,
		tokenmanagersimulation.SimulateMsgDeleteInfo(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
