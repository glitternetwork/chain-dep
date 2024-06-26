package types

import (
	"github.com/pkg/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgEditTable     = "EditTable"
	TypeMsgCreateDataset = "CreateDataset"
	TypeMsgEditDataset   = "EditDataset"
)

func strSholudNotEmpty(name string, s string) error {
	if len(s) > 0 {
		return nil
	}
	return errors.Errorf("%s should not empty", name)
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
	meta string,
	description string,
	duration int64,

) *CreateDatasetRequest {
	return &CreateDatasetRequest{
		FromAddress:     fromAddress.String(),
		DatasetName:     datasetName,
		WorkStatus:      workStatus,
		Hosts:           hosts,
		ManageAddresses: manageAddresses,
		Meta:            meta,
		Description:     description,
		Duration:        duration,
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
	meta string,
	description string,
) *EditDatasetRequest {
	return &EditDatasetRequest{
		FromAddress:     fromAddress.String(),
		DatasetName:     datasetName,
		WorkStatus:      workStatus,
		Hosts:           hosts,
		ManageAddresses: manageAddresses,
		Meta:            meta,
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
