package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (t msgServer) ReportViolation(goCtx context.Context, msg *types.MsgCreateViolationReport) (*types.MsgCreateViolationReportResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := t.checkIsAnActiveParty(ctx, msg.Creator); err != nil {
		return nil, sdkerrors.Wrap(err, "only active party can report the violation")
	}

	if err := t.checkIsAnActiveParty(ctx, msg.Offender); err != nil {
		return nil, sdkerrors.Wrap(err, "only active party can be reported")
	}

	index := types.NewViolationReportIndex(msg.SessionId, msg.Offender, msg.Creator, msg.ViolationType)

	if _, found := t.GetViolationReport(ctx, index); found {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrConflict,
			"session_id: %s, offender: %s, sender: %s violation_type: %s",
			msg.SessionId,
			msg.Offender,
			msg.Creator,
			msg.ViolationType,
		)
	}

	t.SetViolationReport(ctx, types.ViolationReport{
		Index: index,
		Msg:   msg.Msg,
	})

	params := t.GetParams(ctx)
	reports := make(map[string]struct{})

	t.IterateViolationReports(ctx, msg.SessionId, msg.Offender, func(report types.ViolationReport) (stop bool) {
		reports[report.Index.Sender] = struct{}{}
		return false
	})

	party := getPartyByAccount(msg.Offender, params.Parties)

	alreadyIncremented := false
	for _, sessionReported := range party.ReportedSessions {
		alreadyIncremented = alreadyIncremented || (sessionReported == msg.SessionId)
	}

	if !alreadyIncremented {
		if uint64(len(reports)) >= params.Threshold {
			party.ViolationsCount++
			party.ReportedSessions = append(party.ReportedSessions, msg.SessionId)
		}

		if party.ViolationsCount == params.MaxViolationsCount {
			party.Status = types.PartyStatus_Frozen
			party.FreezeEndBlock = uint64(ctx.BlockHeight()) + params.FreezeBlocksPeriod

			ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypePartyFrozen,
				sdk.NewAttribute(types.AttributeKeyPartyAccount, party.Account),
			))

			params.IsUpdateRequired = true
		}

		t.SetParams(ctx, params)
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeNewViolationReport,
		sdk.NewAttribute(types.AttributeKeySessionId, msg.SessionId),
		sdk.NewAttribute(sdk.AttributeKeySender, msg.Offender),
		sdk.NewAttribute(types.AttributeKeyPartyAccount, msg.Offender),
		sdk.NewAttribute(types.AttributeKeyViolationType, msg.ViolationType.String()),
	))

	return &types.MsgCreateViolationReportResponse{}, nil
}

func getPartyByAccount(account string, parties []*types.Party) *types.Party {
	for _, party := range parties {
		if party.Account == account {
			return party
		}
	}

	return nil
}