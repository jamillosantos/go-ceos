package tests

import (
	"database/sql/driver"
	"errors"
	"strings"
)

type CapitalString string

// Value returns a driver Value.
func (str CapitalString) Value() (driver.Value, error) {
	return strings.ToUpper(string(str)), nil
}

type ErrorString string

// Value returns a driver Value.
func (str ErrorString) Value() (driver.Value, error) {
	return nil, errors.New("valuer with error")
}
