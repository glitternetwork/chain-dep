package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/index module sentinel errors
var (
	ErrDatasetExist             = sdkerrors.Register(ModuleName, 2001, "dataset already exist")
	ErrDatasetAbsent            = sdkerrors.Register(ModuleName, 2002, "dataset absent")
	ErrConsumerStakeAbsent      = sdkerrors.Register(ModuleName, 2003, "consumer stake absent")
	ErrDatasetNothingNeedUpdate = sdkerrors.Register(ModuleName, 2004, "dataset nothing need update")
)
