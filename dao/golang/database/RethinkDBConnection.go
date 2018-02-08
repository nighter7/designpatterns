package database

import (
	"fmt"
	"log"
	"sync"

	v "github.com/spf13/viper"
	r "gopkg.in/gorethink/gorethink.v4"
)

// RethinkDBConnection struct
type RethinkDBConnection struct {
	session *r.Session
	error   error
}

var rtdbconn *RethinkDBConnection
var once sync.Once

func init() {
	v.AddConfigPath("../config")
	v.SetConfigName("database")

	err := v.ReadInConfig()
	if err != nil {
		fmt.Printf("Config file not found...")
	}
}

// GetInstance returns a reference to RethinkDBConnection instance
func GetInstance() *RethinkDBConnection {
	once.Do(func() {
		rtdbconn = &RethinkDBConnection{}
		rtdbconn.establishConnection()
	})

	return rtdbconn
}

// establishConnection will return a session to RethinkDB database
func (c *RethinkDBConnection) establishConnection() {
	if c.session == nil {
		c.session, c.error = r.Connect(r.ConnectOpts{
			Address:  v.GetString("dao.address"),
			Database: v.GetString("dao.database"),
			Username: v.GetString("dao.username"),
			Password: v.GetString("dao.password"),
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
