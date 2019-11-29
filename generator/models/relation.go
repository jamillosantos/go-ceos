package models

type (
	Relation struct {
		Query            *Query
		Name             string
		LocalModelType   string
		LocalColumn      string
		ForeignModelType string
		ForeignColumn    string
	}
)

func NewRelation(query *Query, name, localModelType, localColumn, foreignModelType, foreignColumn string) *Relation {
	return &Relation{
		Query:            query,
		Name:             name,
		LocalModelType:   localModelType,
		LocalColumn:      localColumn,
		ForeignModelType: foreignModelType,
		ForeignColumn:    foreignColumn,
	}
}
