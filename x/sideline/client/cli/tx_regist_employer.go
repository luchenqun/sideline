package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"sideline/x/sideline/types"
)

var _ = strconv.Itoa(0)

func CmdRegistEmployer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "regist-employer [name] [introduce] [email] [avatar]",
		Short: "Broadcast message regist-employer",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argIntroduce := args[1]
			argEmail := args[2]
			argAvatar := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRegistEmployer(
				clientCtx.GetFromAddress().String(),
				argName,
				argIntroduce,
				argEmail,
				argAvatar,
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
