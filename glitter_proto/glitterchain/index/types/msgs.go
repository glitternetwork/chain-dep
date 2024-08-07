package types

import (
	"github.com/glitternetwork/chain-dep/utils"
	"github.com/pkg/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgEditTable      = "EditTable"
	TypeMsgCreateDataset  = "CreateDataset"
	TypeMsgEditDataset    = "EditDataset"
	TypeMsgRenewalDataset = "RenewalDataset"
)

var _ sdk.Msg = &EditTableRequest{}

func NewEditTableRequest(fromAddress sdk.AccAddress, datasetName string, tableName string, meta string, description string) *EditTableRequest {
	return &EditTableRequest{
		FromAddress: fromAddress.String(),
		DatasetName: datasetName,
		TableName:   tableName,
		Meta:        meta,
		Description: description,
	}
}
func (msg EditTableRequest) Route() string {
	return RouterKey
}

func (m EditTableRequest) Type() string {
	return TypeMsgEditTable
}

func (msg EditTableRequest) ValidateBasic() error {
	if err := utils.StrSholudNotEmpty("datasetName", msg.GetDatasetName()); err != nil {
		return err
	}
	if err := utils.StrSholudNotEmpty("tableName", msg.GetTableName()); err != nil {
		return err
	}
	if err := utils.StrSholudNotEmpty("meta", msg.GetMeta()); err != nil {
		return err
	}
	if err := utils.StrSholudNotEmpty("description", msg.GetDescription()); err != nil {
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
	if err := utils.StrSholudNotEmpty("fromAddress", msg.GetFromAddress()); err != nil {
		return err
	}
	if err := utils.StrSholudNotEmpty("datasetName", msg.GetDatasetName()); err != nil {
		return err
	}

	isFind := false
	workstatus := msg.GetWorkStatus()
	for k, _ := range ServiceStatus_name {
		if k == int32(workstatus) {
			isFind = true
			break
		}
	}
	if !isFind {
		return errors.New("param work status error")
	}
	if err := utils.StrSholudNotEmpty("hosts", msg.GetHosts()); err != nil {
		return err
	}
	if err := utils.StrSholudNotEmpty("manage address", msg.GetManageAddresses()); err != nil {
		return err
	}
	if err := utils.StrSholudNotEmpty("meta", msg.GetMeta()); err != nil {
		return err
	}
	if err := utils.StrSholudNotEmpty("description", msg.GetDescription()); err != nil {
		return err
	}
	if msg.GetDuration() <= 0 {
		return errors.New("param duration error")
	}
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
	if err := utils.StrSholudNotEmpty("fromAddress", msg.GetFromAddress()); err != nil {
		return err
	}
	if err := utils.StrSholudNotEmpty("datasetName", msg.GetDatasetName()); err != nil {
		return err
	}
	isFind := false
	workstatus := msg.GetWorkStatus()
	for k, _ := range ServiceStatus_name {
		if k == int32(workstatus) {
			isFind = true
			break
		}
	}
	if !isFind {
		return errors.New("param work status error")
	}
	if err := utils.StrSholudNotEmpty("hosts", msg.GetHosts()); err != nil {
		return err
	}
	if err := utils.StrSholudNotEmpty("manage address", msg.GetManageAddresses()); err != nil {
		return err
	}
	if err := utils.StrSholudNotEmpty("meta", msg.GetMeta()); err != nil {
		return err
	}
	if err := utils.StrSholudNotEmpty("description", msg.GetDescription()); err != nil {
		return err
	}
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

// EditDatasetRequest
var _ sdk.Msg = &RenewalDatasetRequest{}

func NewRenewalDatasetRequest(
	fromAddress sdk.AccAddress,
	datasetName string,
	duration int64,
) *RenewalDatasetRequest {
	return &RenewalDatasetRequest{
		FromAddress: fromAddress.String(),
		DatasetName: datasetName,
		Duration:    duration,
	}
}
func (msg RenewalDatasetRequest) Route() string {
	return RouterKey
}
func (m RenewalDatasetRequest) Type() string {
	return TypeMsgRenewalDataset
}
func (msg RenewalDatasetRequest) ValidateBasic() error {
	if err := utils.StrSholudNotEmpty("fromAddress", msg.GetFromAddress()); err != nil {
		return err
	}
	if err := utils.StrSholudNotEmpty("datasetName", msg.GetDatasetName()); err != nil {
		return err
	}
	if msg.GetDuration() <= 0 {
		return errors.New("param duration error")
	}
	return nil
}
func (msg RenewalDatasetRequest) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.GetFromAddress())
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}
func (msg RenewalDatasetRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

/*end*/
