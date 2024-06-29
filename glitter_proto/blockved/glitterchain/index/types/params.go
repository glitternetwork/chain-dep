package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

var (
	KeyFeePerDatasetPerSecond = []byte("FeePerDatasetPerSecond")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(feePerDatasetPerSecond sdk.Int) Params {
	return Params{
		FeePerDatasetPerSecond: feePerDatasetPerSecond,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	DefaultFeePerDatasetPerSecond, _ := sdk.NewIntFromString("1000000000000")
	return NewParams(DefaultFeePerDatasetPerSecond)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyFeePerDatasetPerSecond, &p.FeePerDatasetPerSecond, validateFeePerDatasetPerSecond),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if p.FeePerDatasetPerSecond.LT(sdk.ZeroInt()) {
		return errors.New("param fee_per_dataset_per_second less than 0;fee_per_dataset_per_second=" + p.FeePerDatasetPerSecond.String())
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateFeePerDatasetPerSecond(i interface{}) error {
	v, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.LT(sdk.ZeroInt()) {
		return fmt.Errorf("max entries must be positive: %d", v)
	}

	return nil
}
