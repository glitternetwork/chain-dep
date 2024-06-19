package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/index module sentinel errors
var (
	ErrDatasetExist             = sdkerrors.Register(ModuleName, 1001, "dataset already exist")
	ErrDatasetAbsent            = sdkerrors.Register(ModuleName, 1002, "dataset absent")
	ErrConsumerStakeAbsent      = sdkerrors.Register(ModuleName, 1003, "consumer stake absent")
	ErrDatasetNothingNeedUpdate = sdkerrors.Register(ModuleName, 1004, "dataset nothing need update")
)
