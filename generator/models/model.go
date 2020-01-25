package models

type (
	ModelField struct {
		Name string
	}

	Model struct {
		Name      string
		PK        *ModelField
		Fields    []*ModelField
		Relations []*Relation
	}
)

func NewModel(name string) *Model {
	return &Model{
		Name:      name,
		Fields:    make([]*ModelField, 0),
		Relations: make([]*Relation, 0),
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

func (model *Model) AddRelation(relation *Relation) *Relation {
	model.Relations = append(model.Relations, relation)
	return relation
}
