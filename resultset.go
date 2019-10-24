package ceous

import (
	"io"
)

type ResultSet interface {
	Next() bool
	Scan(dest ...interface{}) error
	Columns() ([]string, error)
	io.Closer
}

type RecordResultSet struct {
	ResultSet
	lastErr error
}

func NewRecordResultSet(rows ResultSet, err error) *RecordResultSet {
	return &RecordResultSet{
		ResultSet: rows,
		lastErr:   err,
	}
}

func (rs *RecordResultSet) ToModel(model Record) error {
	if rs.lastErr != nil {
		return rs.lastErr
	}

	columns, err := rs.ResultSet.Columns()
	if err != nil {
		return err
	}
	scanColumns := make([]interface{}, len(columns))
	for i, column := range columns {
		sValue, err := model.ColumnAddress(column)
		if err != nil {
			return err
		}
		scanColumns[i] = sValue
	}
	return rs.ResultSet.Scan(scanColumns...)
}

func (rs *RecordResultSet) Scan(columns ...interface{}) error {
	if rs.lastErr != nil {
		return rs.lastErr
	}
	return rs.ResultSet.Scan(columns...)
}
