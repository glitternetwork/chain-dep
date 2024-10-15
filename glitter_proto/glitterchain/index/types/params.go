package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"strconv"
)

const (
	DefaultGracePeriod int64 = 86400 * 7 * 3
)

var (
	KeyFeePerDatasetPerSecond = []byte("FeePerDatasetPerSecond")
	KeyGracePeriod            = []byte("GracePeriod")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(feePerDatasetPerSecond sdk.Int, gracePeriod int64) Params {
	return Params{
		FeePerDatasetPerSecond: feePerDatasetPerSecond,
		GracePeriod:            gracePeriod,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	DefaultFeePerDatasetPerSecond, _ := sdk.NewIntFromString("1000000000000")
	return NewParams(DefaultFeePerDatasetPerSecond, DefaultGracePeriod)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyFeePerDatasetPerSecond, &p.FeePerDatasetPerSecond, validateFeePerDatasetPerSecond),
		paramtypes.NewParamSetPair(KeyGracePeriod, &p.GracePeriod, validateGracePeriod),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if p.FeePerDatasetPerSecond.LT(sdk.ZeroInt()) {
		return errors.New("param fee_per_dataset_per_second less than 0;fee_per_dataset_per_second=" + p.FeePerDatasetPerSecond.String())
	}
	if p.GracePeriod < 0 {
		return errors.New("param fee_per_dataset_per_second less than 0;grace_period=" + strconv.FormatInt(p.GetGracePeriod(), 10))
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

func validateGracePeriod(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v < 0 {
		return fmt.Errorf("max entries must be positive: %d", v)
	}
	return nil
}
