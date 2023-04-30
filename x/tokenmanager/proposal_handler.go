package tokenmanager

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/keeper"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func NewProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.AddNetworkProposal:
			return k.HandleAddNetworkProposal(ctx, c)
		case *types.RemoveNetworkProposal:
			return k.HandleRemoveNetworkProposal(ctx, c)
		case *types.UpdateContractAddressProposal:
			return k.HandleUpdateContractAddressProposal(ctx, c)
		case *types.AddFeeTokenProposal:
			return k.HandleAddFeeTokenProposal(ctx, c)
		case *types.RemoveFeeTokenProposal:
			return k.HandleRemoveFeeTokenProposal(ctx, c)
		case *types.UpdateFeeTokenProposal:
			return k.HandleUpdateFeeTokenProposal(ctx, c)
		case *types.WithdrawFeeProposal:
			return k.HandleWithdrawFeeProposal(ctx, c)
		case *types.UpdateTokenItemProposal:
			return k.HandleUpdateTokenItemProposal(ctx, c)
		case *types.RemoveTokenItemProposal:
			return k.HandleRemoveTokenItemProposal(ctx, c)
		case *types.CreateCollectionProposal:
			return k.HandleCreateCollectionProposal(ctx, c)
		case *types.UpdateCollectionDataProposal:
			return k.HandleUpdateCollectionDataProposal(ctx, c)
		case *types.AddCollectionDataProposal:
			return k.HandleAddCollectionDataProposal(ctx, c)
		case *types.RemoveCollectionDataProposal:
			return k.HandleRemoveCollectionDataProposal(ctx, c)
		case *types.RemoveCollectionProposal:
			return k.HandleRemoveCollectionProposal(ctx, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized proposal content type: %T", c)
		}
	}
}
