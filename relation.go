package ceous

type (
	Relation interface {
		Aggregate(model Record) error
		Realize() error
	}
)
