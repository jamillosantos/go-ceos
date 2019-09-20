package ceous

type Model interface {
	ColumnAddress(name string) (interface{}, error)
}

type Record interface {
	Model
}
