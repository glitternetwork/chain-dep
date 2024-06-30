package utils

import (
	"github.com/pkg/errors"
)

const (
	DefaultStringFieldUnsetData = "unset-data"
	DefaultInt64FieldUnsetdata  = -999999
)

func StrSholudNotEmpty(name string, s string) error {
	if len(s) > 0 {
		return nil
	}
	return errors.Errorf("%s should not empty", name)
}
