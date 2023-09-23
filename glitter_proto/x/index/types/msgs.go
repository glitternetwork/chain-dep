package types

import (
	common "github.com/glitternetwork/chain-dep/glitter_proto/common"
	"github.com/pkg/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgSchema   = "Schema"
	TypeMsgDoc      = "Doc"
	TypeMsgSQLExec  = "SQLExec"
	TypeMsgSQLGrant = "SQLGrant"
)

var _ sdk.Msg = &SQLExecRequest{}

func NewSQLExecRequest(ownerAddress sdk.AccAddress, sql string, args []*common.Argument) *SQLExecRequest {
	return &SQLExecRequest{
		Uid:       ownerAddress.String(),
		Sql:       sql,
		Arguments: args,
	}
}
func (msg SQLExecRequest) Route() string {
	return RouterKey
}

func (m SQLExecRequest) Type() string {
	return TypeMsgSQLExec
}

func (msg SQLExecRequest) ValidateBasic() error {
	return nil
}

func (msg SQLExecRequest) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Uid)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

// GetSignBytes Implements Msg.
func (msg SQLExecRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

var _ sdk.Msg = &SQLGrantRequest{}

func NewSQLGrantRequest(ownerAddress sdk.AccAddress, onDatabase string, onTable string, toUID string, role string) *SQLGrantRequest {
	return &SQLGrantRequest{
		Uid:        ownerAddress.String(),
		OnTable:    onTable,
		ToUID:      toUID,
		Role:       role,
		OnDatabase: onDatabase,
	}
}
func (msg SQLGrantRequest) Route() string {
	return RouterKey
}

func (m SQLGrantRequest) Type() string {
	return TypeMsgSQLGrant
}

func (msg SQLGrantRequest) ValidateBasic() error {
	if err := strSholudNotEmpty("OnDatabase", msg.OnDatabase); err != nil {
		return err
	}
	if err := strSholudNotEmpty("ToUID", msg.ToUID); err != nil {
		return err
	}
	if err := strSholudNotEmpty("Role", msg.Role); err != nil {
		return err
	}
	return nil
}

func (msg SQLGrantRequest) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Uid)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

// GetSignBytes Implements Msg.
func (msg SQLGrantRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func strSholudNotEmpty(name string, s string) error {
	if len(s) > 0 {
		return nil
	}
	return errors.Errorf("%s should not empty", name)
}
