package ceous

import (
	"database/sql"
	"io"
)

type ResultSet interface {
	Next() bool
	io.Closer
}

type RecordResultSet struct {
	*sql.Rows
}

func NewRecordResultSet(rows *sql.Rows) *RecordResultSet {
	return &RecordResultSet{
		Rows: rows,
	}
}

func (rs *RecordResultSet) ToModel(model Model) error {
	columns, err := rs.Rows.Columns()
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
	return rs.Rows.Scan(scanColumns...)
}

func (rs *RecordResultSet) Scan(columns ...interface{}) error {
	return rs.Rows.Scan(columns...)
}
