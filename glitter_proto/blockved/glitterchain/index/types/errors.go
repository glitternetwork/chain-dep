package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/index module sentinel errors
var (
	ErrEsInitClientFirst   = sdkerrors.Register(ModuleName, 1001, "es error init client first")
	ErrEsFailCreateMapping = sdkerrors.Register(ModuleName, 1002, "es error fail create mapping")
	ErrEsSchemaParseError  = sdkerrors.Register(ModuleName, 1003, "es error parse schema error")
	ErrEsInsertDocError    = sdkerrors.Register(ModuleName, 1004, "es error insert doc error")
	ErrEsQueryDocError     = sdkerrors.Register(ModuleName, 1005, "es error query doc error")

	ErrSchemaAlreadyExist   = sdkerrors.Register(ModuleName, 2001, "schema already exist")
	ErrSchemaSchemaNotFound = sdkerrors.Register(ModuleName, 2002, "schema not found")

	ErrParseDocError = sdkerrors.Register(ModuleName, 3001, "parse doc error")

	ErrDatasetExist = sdkerrors.Register(ModuleName, 4001, "dataset exist")

	ErrDatasetAbsent = sdkerrors.Register(ModuleName, 4002, "dataset  absent")

	ErrDatasetNothingNeedUpdate = sdkerrors.Register(ModuleName, 4003, "dataset nothing need update")
)
