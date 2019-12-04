package ceous

// Valuer ...
type (
	Valuer interface {
		Value(column string) (interface{}, error)
	}

	ColumnAddresser interface {
		ColumnAddress(name string) (interface{}, error)
	}

	Record interface {
		GetID() interface{}
		IsPersisted() bool
		setPersisted()
		IsWritable() bool
		setWritable(bool)
		ColumnAddresser
		Valuer
	}

	Model struct {
		persisted bool
		writable  bool
	}
)

func MakeWritable(record Record) {
	record.setWritable(true)
}

func (model *Model) IsPersisted() bool {
	return model.persisted
}

func (model *Model) setPersisted() {
	model.persisted = true
}

func (model *Model) IsWritable() bool {
	return model.writable
}

func (model *Model) setWritable(value bool) {
	model.writable = value
}
