package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/sideline module sentinel errors
var (
	ErrNotRegistForEmployer  = sdkerrors.Register(ModuleName, 1100, "not regist for a employer")
	ErrNotRegistForDeveloper = sdkerrors.Register(ModuleName, 1101, "not regist for a developer")
	ErrTaskID                = sdkerrors.Register(ModuleName, 1102, "task id is not exist")
	ErrTaskStatus            = sdkerrors.Register(ModuleName, 1103, "task status is forbid do this")
	ErrPermission            = sdkerrors.Register(ModuleName, 1104, "no permission do this")
	ErrTime                  = sdkerrors.Register(ModuleName, 1105, "this operation is not allowed at this moment")
	ErrCoinAmount            = sdkerrors.Register(ModuleName, 1106, "coin amount must more than 0")
)
