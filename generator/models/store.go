package models

type Store struct {
	Name     string
	ModelRef string
}

// NewStore returns a new instance of a `Store` with a given `name`.
func NewStore(name string) *Store {
	return &Store{
		Name: name,
	}
}
