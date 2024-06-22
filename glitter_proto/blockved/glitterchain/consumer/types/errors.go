package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInsufficientToken         = sdkerrors.Register(ModuleName, 1001, "insufficient token")
	ErrDatasetAbsent             = sdkerrors.Register(ModuleName, 1002, "dataset absent")
	ErrDatasetAddStakeFail       = sdkerrors.Register(ModuleName, 1003, "dataset add stake fail")
	ErrConsumerAbsent            = sdkerrors.Register(ModuleName, 1004, "consumer absent")
	ErrConsumerStakeAbsent       = sdkerrors.Register(ModuleName, 1005, "consumer stake absent")
	ErrInsufficientUnStakeAmount = sdkerrors.Register(ModuleName, 1006, "insufficient unstake token")
)
