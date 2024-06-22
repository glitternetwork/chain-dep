package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgPledgeRequest        = "PledgeRequest"
	TypeMsgReleasePledgeRequest = "ReleasePledgeRequest"
)

var _ sdk.Msg = &PledgeRequest{}

func NewPledgeRequest(address sdk.AccAddress, datasetName string, amount sdk.Int) *PledgeRequest {
	return &PledgeRequest{
		Address:     address.String(),
		DatasetName: datasetName,
		Amount:      amount,
	}
}
func (msg PledgeRequest) Route() string {
	return RouterKey
}

func (m PledgeRequest) Type() string {
	return TypeMsgPledgeRequest
}

func (msg PledgeRequest) ValidateBasic() error {
	return nil
}

func (msg PledgeRequest) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

// GetSignBytes Implements Msg.
func (msg PledgeRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

var _ sdk.Msg = &ReleasePledgeRequest{}

func NewReleasePledgeRequest(address sdk.AccAddress, datasetName string, amount sdk.Int) *ReleasePledgeRequest {
	return &ReleasePledgeRequest{
		Address:     address.String(),
		DatasetName: datasetName,
		Amount:      amount,
	}
}
func (msg ReleasePledgeRequest) Route() string {
	return RouterKey
}

func (msg ReleasePledgeRequest) Type() string {
	return TypeMsgReleasePledgeRequest
}

func (msg ReleasePledgeRequest) ValidateBasic() error {
	return nil
}

func (msg ReleasePledgeRequest) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

// GetSignBytes Implements Msg.
func (msg ReleasePledgeRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}
