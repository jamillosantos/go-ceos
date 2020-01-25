package models

type Connection struct {
	Name     string
	FullName string
}

// NewConnection retrurns a new instance of `Connection` with the given `name`
// set.
func NewConnection(name string) *Connection {
	return &Connection{
		Name:     name,
		FullName: name + "Connection",
	}
}
