package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgStakeRequest   = "StakeRequest"
	TypeMsgUnStakeRequest = "UnStakeRequest"
)

var _ sdk.Msg = &StakeRequest{}

func NewStakeRequest(address sdk.AccAddress, datasetName string, amount sdk.Int) *StakeRequest {
	return &StakeRequest{
		Address:     address.String(),
		DatasetName: datasetName,
		Amount:      amount,
	}
}
func (msg StakeRequest) Route() string {
	return RouterKey
}

func (m StakeRequest) Type() string {
	return TypeMsgStakeRequest
}

func (msg StakeRequest) ValidateBasic() error {
	return nil
}

func (msg StakeRequest) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

// GetSignBytes Implements Msg.
func (msg StakeRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

var _ sdk.Msg = &UnStakeRequest{}

func NeUnStakeRequest(address sdk.AccAddress, amount sdk.Int) *UnStakeRequest {
	return &UnStakeRequest{
		Address: address.String(),
		Amount:  amount,
	}
}
func (msg UnStakeRequest) Route() string {
	return RouterKey
}

func (m UnStakeRequest) Type() string {
	return TypeMsgUnStakeRequest
}

func (msg UnStakeRequest) ValidateBasic() error {
	return nil
}

func (msg UnStakeRequest) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

// GetSignBytes Implements Msg.
func (msg UnStakeRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}
