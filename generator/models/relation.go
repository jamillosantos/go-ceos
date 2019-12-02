package models

type (
	Relation struct {
		Query            *Query
		Name             string
		FullName         string
		LocalField       string
		LocalModelType   string
		LocalColumn      string
		ForeignField     string
		ForeignModelType string
		ForeignColumn    string
		ForeignFieldType string
	}
)

func NewRelation(query *Query, name, localField, localModelType, localColumn, foreignModelType, foreignField, foreignColumn, foreignFieldType string) *Relation {
	return &Relation{
		Query:            query,
		Name:             name,
		FullName:         name + "Relation",
		LocalField:       localField,
		LocalModelType:   localModelType,
		LocalColumn:      localColumn,
		ForeignModelType: foreignModelType,
		ForeignField:     foreignField,
		ForeignColumn:    foreignColumn,
		ForeignFieldType: foreignFieldType,
	}
}
