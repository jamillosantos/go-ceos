package models

type Connection struct {
	Name string
}

func NewConnection(name string) *Connection {
	return &Connection{
		Name: name,
	}
}

func (conn *Connection) ConnectionName() string {
	return conn.Name + "Connection"
}
