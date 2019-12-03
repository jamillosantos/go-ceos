package models

type (
	Relation struct {
		Query            *Query
		Name             string
		FullName         string
		LocalField       string
		LocalFieldRef    string
		LocalModelType   string
		LocalColumn      string
		ForeignField     string
		ForeignModelType string
		ForeignColumn    string
		ForeignFieldType string
	}
)

func NewRelation(query *Query, name, localField, localFieldRef, localModelType, localColumn, foreignModelType, foreignField, foreignColumn, foreignFieldType string) *Relation {
	return &Relation{
		Query:            query,
		Name:             name,
		FullName:         name + "Relation",
		LocalField:       localField,
		LocalFieldRef:    localFieldRef,
		LocalModelType:   localModelType,
		LocalColumn:      localColumn,
		ForeignModelType: foreignModelType,
		ForeignField:     foreignField,
		ForeignColumn:    foreignColumn,
		ForeignFieldType: foreignFieldType,
	}
}
