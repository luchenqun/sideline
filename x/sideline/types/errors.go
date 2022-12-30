package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/sideline module sentinel errors
var (
	ErrNotRegistForEmployer = sdkerrors.Register(ModuleName, 1100, "not regist for a employer")
	ErrTaskID               = sdkerrors.Register(ModuleName, 1101, "task id is not exist")
	ErrTaskStatus           = sdkerrors.Register(ModuleName, 1102, "task status is forbid do this")
	ErrPermission           = sdkerrors.Register(ModuleName, 1103, "no permission do this")
	ErrTime                 = sdkerrors.Register(ModuleName, 1104, "this operation is not allowed at this moment")
)
