package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/multisig/types"
)

func (k Keeper) EndBlocker(ctx sdk.Context) {
	param := k.GetParams(ctx)
	currentBlock := uint64(ctx.BlockHeight())

	k.IterateProposals(ctx, func(proposal types.Proposal) (stop bool) {
		if proposal.VotingEndBlock == currentBlock {
			proposal.FinalTallyResult = k.Tally(ctx, proposal.Id)
			proposal.Status = types.ProposalStatus_REJECTED

			group, _ := k.GetGroup(ctx, proposal.Group)

			if proposal.FinalTallyResult.YesCount >= group.Threshold {
				proposal.Status = types.ProposalStatus_ACCEPTED
			}

			k.SetProposal(ctx, proposal)

			ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeProposalStatusChanged,
				sdk.NewAttribute(types.AttributeKeyProposal, fmt.Sprintf("%d", proposal.Id)),
				sdk.NewAttribute(types.AttributeKeyProposalStatus, proposal.Status.String()),
			))
		}

		if proposal.Status == types.ProposalStatus_ACCEPTED {
			var logs string
			proposal.Status, logs = k.ExecuteProposal(ctx, proposal)

			k.SetProposal(ctx, proposal)

			ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeProposalExecuted,
				sdk.NewAttribute(types.AttributeKeyProposal, fmt.Sprintf("%d", proposal.Id)),
				sdk.NewAttribute(types.AttributeKeyProposalStatus, proposal.Status.String()),
				sdk.NewAttribute(types.AttributeKeyProposalExecutionLogs, logs),
			))
		}

		if proposal.VotingEndBlock+param.PrunePeriod == currentBlock {
			votes := k.GetAllVoteByProposalId(ctx, proposal.Id)

			for _, vote := range votes {
				k.RemoveVote(ctx, vote.ProposalId, vote.Voter)
			}

			k.RemoveProposal(ctx, proposal.Id)

			ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeProposalPruned,
				sdk.NewAttribute(types.AttributeKeyProposal, fmt.Sprintf("%d", proposal.Id)),
			))
		}

		return false
	})
}
