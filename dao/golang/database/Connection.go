package database

import "database/sql"

// ConnectionType represents the database connection types
type ConnectionType int

// Types representing database engines
const (
	RethinkDB ConnectionType = iota
	PostgresSQL
)

// Connection represents the methods to be implemented by concrete clases
type Connection interface {
	Get(t ConnectionType) (sql.Conn, error)
	Close() error
}
