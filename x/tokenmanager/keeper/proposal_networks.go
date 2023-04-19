package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func (k Keeper) HandleSetNetworkProposal(ctx sdk.Context, proposal *types.SetNetworkProposal) error {
	params := k.GetParams(ctx)
	for i, network := range params.Networks {
		if network.Name == proposal.NetworkParams.Name {
			params.Networks[i] = proposal.NetworkParams
			k.SetParams(ctx, params)
			return nil
		}
	}

	params.Networks = append(params.Networks, proposal.NetworkParams)
	k.SetParams(ctx, params)
	return nil
}