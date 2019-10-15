package ceous

import (
	"strings"
)

type (
	PrimaryKey interface {
		Columnable
		ColumnAddresser
		Valuer
	}

	BasePrimaryKey struct {
	}

	pkWrapper struct {
		PrimaryKey
		prefix string
	}
)

func WrapPK(prefix string, key PrimaryKey) *pkWrapper {
	return &pkWrapper{
		PrimaryKey: key,
		prefix:     prefix,
	}
}

func (pkw *pkWrapper) ColumnAddress(column string) (interface{}, error) {
	if strings.HasPrefix(column, pkw.prefix) {
		return pkw.PrimaryKey.ColumnAddress(column[len(pkw.prefix):])
	}
	return nil, ErrFieldNotFound
}

func (pkw *pkWrapper) Value(column string) (interface{}, error) {
	if strings.HasPrefix(column, pkw.prefix) {
		return pkw.PrimaryKey.Value(column[len(pkw.prefix):])
	}
	return nil, ErrFieldNotFound
}

func (pkw *pkWrapper) Columns() []SchemaField {
	return pkw.PrimaryKey.Columns()
}
