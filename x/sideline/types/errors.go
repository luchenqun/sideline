package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/sideline module sentinel errors
var (
	ErrNotRegistForEmployer = sdkerrors.Register(ModuleName, 1100, "not regist for a employer")
)
