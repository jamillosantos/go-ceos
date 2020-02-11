package models

import (
	"github.com/jamillosantos/go-ceous/generator/naming"
)

type Store struct {
	Name     string
	FullName string
	ModelRef string
}

// NewStore returns a new instance of a `Store` with a given `name`.
func NewStore(name string) *Store {
	return &Store{
		Name:     name,
		FullName: naming.PascalCase.Do(name) + "Store",
	}
}
