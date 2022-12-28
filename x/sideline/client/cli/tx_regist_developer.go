package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"sideline/x/sideline/types"
	"strings"
)

var _ = strconv.Itoa(0)

func CmdRegistDeveloper() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "regist-developer [name] [introduce] [email] [avatar] [education] [major] [skills]",
		Short: "Broadcast message regist-developer",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argIntroduce := args[1]
			argEmail := args[2]
			argAvatar := args[3]
			argEducation := args[4]
			argMajor := args[5]
			argSkills := strings.Split(args[6], listSeparator)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRegistDeveloper(
				clientCtx.GetFromAddress().String(),
				argName,
				argIntroduce,
				argEmail,
				argAvatar,
				argEducation,
				argMajor,
				argSkills,
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
