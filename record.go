package ceous

// Valuer ...
type Valuer interface {
	Value(column string) (interface{}, error)
}

type ColumnAddresser interface {
	ColumnAddress(name string) (interface{}, error)
}

type Record interface {
	GetID() interface{}
	IsPersisted() bool
	setPersisted()
	IsWritable() bool
	setWritable(bool)
	ColumnAddresser
	Valuer
}

type Model struct {
	persisted bool
	writable  bool
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
