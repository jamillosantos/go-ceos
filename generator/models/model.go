package models

type (
	ModelField struct {
		Name string
	}

	Model struct {
		Name   string
		PK     *ModelField
		Fields []*ModelField
	}
)

func NewModel(name string) *Model {
	return &Model{
		Name:   name,
		Fields: make([]*ModelField, 0),
	}
}

func NewModelField(name string) *ModelField {
	return &ModelField{
		Name: name,
	}
}

func (model *Model) AddField(field *ModelField) *ModelField {
	model.Fields = append(model.Fields, field)
	return field
}
