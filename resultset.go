package ceous

import (
	"io"
)

// ResultSet represents the abstraction of a sql.ResultSet struct.
type ResultSet interface {
	Next() bool
	Scan(dest ...interface{}) error
	Columns() ([]string, error)
	io.Closer
}

// RecordScanner describes a scanner for a record.
type RecordScanner interface {
	ScanRecord(rs ResultSet, model Record) error
}

// RecordScannerColumns describes a column selector.
type RecordScannerColumns interface {
	SelectColumns() []SchemaField
}

// BaseRecordScanner implements a basic and generic `RecordScanner`.
type BaseRecordScanner struct{}

var DefaultRecordScanner BaseRecordScanner

// ScanRecord uses the ColumnAddress to read records from the given `rs`.
func (recordScanner *BaseRecordScanner) ScanRecord(rs ResultSet, model Record) error {
	columns, err := rs.Columns()
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
	err = rs.Scan(scanColumns...)
	if err != nil {
		return err
	}

	// TODO(jota): Check for read only models... /// model.setWritable(!rs.readOnly)
	model.setWritable(true)
	model.setPersisted()

	return nil
}
