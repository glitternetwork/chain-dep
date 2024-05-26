package types

import (
	common "github.com/glitternetwork/chain-dep/glitter_proto/common"
	"github.com/pkg/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgSchema        = "Schema"
	TypeMsgDoc           = "Doc"
	TypeMsgSQLExec       = "SQLExec"
	TypeMsgSQLGrant      = "SQLGrant"
	TypeMsgEditTable     = "EditTable"
	TypeMsgCreateDataset = "CreateDataset"
	TypeMsgEditDataset   = "EditDataset"
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

var _ sdk.Msg = &BindHostRequest{}

func NewBindHostRequest(ownerAddress sdk.AccAddress, database string, url string) *BindHostRequest {
	return &BindHostRequest{
		Uid:      ownerAddress.String(),
		Database: database,
		Url:      url,
	}
}
func (msg BindHostRequest) Route() string {
	return RouterKey
}

func (m BindHostRequest) Type() string {
	return TypeMsgSQLGrant
}

func (msg BindHostRequest) ValidateBasic() error {
	if err := strSholudNotEmpty("uid", msg.GetUid()); err != nil {
		return err
	}
	if err := strSholudNotEmpty("database", msg.GetDatabase()); err != nil {
		return err
	}
	if err := strSholudNotEmpty("host", msg.GetUrl()); err != nil {
		return err
	}
	return nil
}

func (msg BindHostRequest) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Uid)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

// GetSignBytes Implements Msg.
func (msg BindHostRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

var _ sdk.Msg = &EditTableRequest{}

func NewEditTableRequest(fromAddress sdk.AccAddress, datasetName string, tableName string, meta string) *EditTableRequest {
	return &EditTableRequest{
		FromAddress: fromAddress.String(),
		DatasetName: datasetName,
		TableName:   tableName,
		Meta:        meta,
	}
}
func (msg EditTableRequest) Route() string {
	return RouterKey
}

func (m EditTableRequest) Type() string {
	return TypeMsgEditTable
}

func (msg EditTableRequest) ValidateBasic() error {
	if err := strSholudNotEmpty("datasetName", msg.GetDatasetName()); err != nil {
		return err
	}
	if err := strSholudNotEmpty("tableName", msg.GetTableName()); err != nil {
		return err
	}
	if err := strSholudNotEmpty("meta", msg.GetMeta()); err != nil {
		return err
	}
	return nil
}

func (msg EditTableRequest) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.GetFromAddress())
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

func (msg EditTableRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// CreateDatasetRequest
var _ sdk.Msg = &CreateDatasetRequest{}

func NewCreateDatasetRequest(
	fromAddress sdk.AccAddress,
	datasetName string,
	workStatus ServiceStatus,
	hosts string,
	manageAddresses string,
	description string,
) *CreateDatasetRequest {
	return &CreateDatasetRequest{
		FromAddress:     fromAddress.String(),
		DatasetName:     datasetName,
		WorkStatus:      workStatus,
		OwnerAddress:    fromAddress.String(),
		Hosts:           hosts,
		ManageAddresses: manageAddresses,
		Description:     description,
	}
}
func (msg CreateDatasetRequest) Route() string {
	return RouterKey
}
func (m CreateDatasetRequest) Type() string {
	return TypeMsgCreateDataset
}
func (msg CreateDatasetRequest) ValidateBasic() error {
	return nil
}
func (msg CreateDatasetRequest) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.GetFromAddress())
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}
func (msg CreateDatasetRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

/*end*/

// EditDatasetRequest
var _ sdk.Msg = &EditDatasetRequest{}

func NewEditDatasetRequest(
	fromAddress sdk.AccAddress,
	datasetName string,
	workStatus ServiceStatus,
	hosts string,
	manageAddresses string,
	description string,
) *EditDatasetRequest {
	return &EditDatasetRequest{
		FromAddress:     fromAddress.String(),
		DatasetName:     datasetName,
		WorkStatus:      workStatus,
		OwnerAddress:    fromAddress.String(),
		Hosts:           hosts,
		ManageAddresses: manageAddresses,
		Description:     description,
	}
}
func (msg EditDatasetRequest) Route() string {
	return RouterKey
}
func (m EditDatasetRequest) Type() string {
	return TypeMsgEditDataset
}
func (msg EditDatasetRequest) ValidateBasic() error {
	return nil
}
func (msg EditDatasetRequest) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.GetFromAddress())
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}
func (msg EditDatasetRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

/*end*/
