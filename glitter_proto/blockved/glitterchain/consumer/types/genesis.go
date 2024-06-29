package types

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec"
)

func NewGenesisState(params Params, consumers []*Consumer) *GenesisState {
	return &GenesisState{
		Params:    params,
		Consumers: consumers,
	}

}
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
	}
}

func (gs GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}
	return nil
}

func GetGenesisStateFromAppState(cdc codec.JSONCodec, appState map[string]json.RawMessage) *GenesisState {
	var genesisState GenesisState

	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return &genesisState
}
