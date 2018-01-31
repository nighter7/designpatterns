package database

import "database/sql"

type RethinkDBConnection struct {
	conn *sql.Conn
}

func (c *RethinkDBConnection) Get(t ConnectionType) (*sql.Conn, error) {
	return c.conn, nil
}

func (c *RethinkDBConnection) Close() error {
	return nil
}
