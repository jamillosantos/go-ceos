package ceous

import "github.com/pkg/errors"

var (
	ErrFieldNotFound              = errors.New("field not found")
	ErrInvalidRecordType          = errors.New("invalid record type")
	ErrInconsistentRelationResult = errors.New("inconsistent relation result")

	ErrULIDInvalidStringFormat = errors.New("invalid ulid string format")
)
