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
	err = rs.ResultSet.Scan(scanColumns...)
	if err != nil {
		return err
	}

	// TODO(jota): Check for read only models... /// model.setWritable(!rs.readOnly)
	model.setWritable(true)
	model.setPersisted()

	return nil
}

func (rs *RecordResultSet) Next() bool {
	if rs.lastErr != nil {
		return false
	}
	return rs.ResultSet.Next()
}

func (rs *RecordResultSet) Close() error {
	if rs.lastErr != nil {
		return rs.lastErr
	}
	return rs.ResultSet.Close()
}

func (rs *RecordResultSet) Scan(columns ...interface{}) error {
	if rs.lastErr != nil {
		return rs.lastErr
	}
	return rs.ResultSet.Scan(columns...)
}
