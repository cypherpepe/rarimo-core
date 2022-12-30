package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
	tokentypes "gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func CmdCreateTransferOp() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-transfer [tx] [event-id] [from-chain] [token-type]",
		Short: "Create a new transfer operation",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexTx := args[0]
			eventId := args[1]

			// Get value arguments
			argFromChain := args[2]
			argTokenType, err := strconv.Atoi(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateTransferOp(
				clientCtx.GetFromAddress().String(),
				indexTx,
				eventId,
				argFromChain,
				tokentypes.Type(argTokenType),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}