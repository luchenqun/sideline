package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"sideline/x/sideline/types"
)

var _ = strconv.Itoa(0)

func CmdCreateTask() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-task [title] [description] [remuneration] [deposit] [collateral] [deadline]",
		Short: "Broadcast message create-task",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTitle := args[0]
			argDescription := args[1]
			argRemuneration := args[2]
			argDeposit := args[3]
			argCollateral := args[4]
			argDeadline, err := cast.ToUint64E(args[5])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateTask(
				clientCtx.GetFromAddress().String(),
				argTitle,
				argDescription,
				argRemuneration,
				argDeposit,
				argCollateral,
				argDeadline,
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
