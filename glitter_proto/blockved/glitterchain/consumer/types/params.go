package types

import (
	"fmt"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
	"time"
)

const (
	//DefaultCPDTReleasingTime     time.Duration = time.Hour * 24 * 7 * 3
	DefaultCPDTReleasingTime     time.Duration = time.Minute * 2
	DefaultMaxCPDTNumPerConsumer uint32        = 1000
	DefaultMaxCPDTEntries        uint32        = 7
)

var (
	KeyReleasingCPDTWaitTime = []byte("ReleasingCPDTWaitTime")
	KeyMaxCPDTEntries        = []byte("MaxCPDTEntries")
)
var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(releasingCPDTWaitTime time.Duration, maxCPDTEntries uint32) Params {
	return Params{
		ReleasingCPDTWaitTime: releasingCPDTWaitTime,
		MaxCPDTEntries:        maxCPDTEntries,
	}
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyReleasingCPDTWaitTime, &p.ReleasingCPDTWaitTime, validateReleasingCPDTWaitTime),
		paramtypes.NewParamSetPair(KeyMaxCPDTEntries, &p.MaxCPDTEntries, validateMaxCPDTEntries),
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(DefaultCPDTReleasingTime, DefaultMaxCPDTEntries)
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateReleasingCPDTWaitTime(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("unbonding time must be positive: %d", v)
	}

	return nil
}

func validateMaxCPDTEntries(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("max entries must be positive: %d", v)
	}

	return nil
}
