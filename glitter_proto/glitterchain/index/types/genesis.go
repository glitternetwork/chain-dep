package types

func NewGenesisState(params Params, datasets []Dataset, cpdts []CPDT) *GenesisState {
	return &GenesisState{
		Params:   params,
		Datasets: datasets,
		CPDTs:    cpdts,
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
