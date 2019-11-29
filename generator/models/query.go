package models

type (
	Query struct {
		Name   string
		Fields []*QueryField
	}

	QueryField struct {
		Name      string
		FieldPath []string
		Type      string
	}
)

// NewQuery returns a new instance of `Query` with the given `name` set.
func NewQuery(name string) *Query {
	return &Query{
		Name:   name,
		Fields: make([]*QueryField, 0),
	}
}

// NewQueryField returns a new instance of `QueryField` with the given params
// set.
func NewQueryField(name string, fieldPath []string) *QueryField {
	return &QueryField{
		Name:      name,
		FieldPath: fieldPath,
	}
}

// AddField appends the given `field` to the fields list.
func (q *Query) AddField(field *QueryField) *QueryField {
	q.Fields = append(q.Fields, field)
	return field
}
