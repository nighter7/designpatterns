package database

import "database/sql"

// ConnectionType represents the database connection types
type ConnectionType int

const (
	// RethinkDB represents itself
	RethinkDB = iota
)

// Connection represents the methods to be implemented by concrete clases
type Connection interface {
	Get(t ConnectionType) sql.Conn
	Close() error
}
