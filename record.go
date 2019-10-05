package ceous

// Valuer ...
type Valuer interface {
	Value(column string) (interface{}, error)
}

type Record interface {
	ColumnAddress(name string) (interface{}, error)
	IsPersisted() bool
	Valuer
	setPersisted()
}

type Model struct {
	persisted bool
}

func (model *Model) IsPersisted() bool {
	return model.persisted
}

func (model *Model) setPersisted() {
	model.persisted = true
}
