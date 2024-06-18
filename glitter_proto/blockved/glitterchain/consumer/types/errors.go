package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInsufficientToken = sdkerrors.Register(ModuleName, 1001, "insufficient token")
)
