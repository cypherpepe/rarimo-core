package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/multisig/types"
)

func (k msgServer) ChangeGroup(goCtx context.Context, msg *types.MsgChangeGroup) (*types.MsgChangeGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: check is was called by the module
	group, found := k.GetGroup(ctx, msg.Group)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "group (%s) not found", msg.Group)
	}

	isCreatorMember := false
	for _, member := range group.Members {
		if member == msg.Creator {
			isCreatorMember = true
			break
		}
	}

	if !isCreatorMember {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "creator (%s) is not a member of group (%s)", msg.Creator, msg.Group)
	}

	group.Members = msg.Members
	group.Threshold = msg.Threshold

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeChangeGroup,
			sdk.NewAttribute(types.AttributeKeyGroup, group.Account),
		),
	)

	return &types.MsgChangeGroupResponse{}, nil
}
