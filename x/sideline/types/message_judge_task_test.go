package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"sideline/testutil/sample"
)

func TestMsgJudgeTask_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgJudgeTask
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgJudgeTask{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgJudgeTask{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
