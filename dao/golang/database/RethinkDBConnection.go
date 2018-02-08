package database

import (
	"log"

	r "gopkg.in/gorethink/gorethink.v4"
)

// RethinkDBConnection struct
type RethinkDBConnection struct {
	session *r.Session
	error   error
}

// establishConnection will return a session to RethinkDB database
func (c *RethinkDBConnection) establishConnection() {
	if c.session == nil {
		c.session, c.error = r.Connect(r.ConnectOpts{
			Address:  "192.168.99.100:8081",
			Database: "test",
			Username: "admin",
			Password: "",
		})

		if c.error != nil {
			log.Fatalln(c.error.Error())
		}
	}
}

// GetSession will return the session generated
func (c *RethinkDBConnection) GetSession() *r.Session {
	c.establishConnection()
	return c.session
}

// GetError will return any error generated when establishing the connection
func (c *RethinkDBConnection) GetError() error {
	return c.error
}
