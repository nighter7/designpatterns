package dao

import (
	"fmt"
	"log"

	r "github.com/nighter7/designpatterns/dao/golang/database"
	"github.com/nighter7/designpatterns/dao/golang/dto"
	rtdb "gopkg.in/gorethink/gorethink.v4"
)

// RethinkDBUserImpl represents a concrete implementation of DAO for rethinkdb
type RethinkDBUserImpl struct {
	rdbc *r.RethinkDBConnection
}

func (b *RethinkDBUserImpl) init() {
	if b.rdbc == nil {
		b.rdbc = new(r.RethinkDBConnection)
		err := b.rdbc.GetError()

		if err != nil {
			log.Fatalln(err.Error())
		}
	}
}

// GetUser returns all the dcouments in the table
func (b *RethinkDBUserImpl) GetUser() map[float64]dto.UserDTO {
	b.init()
	res, err := rtdb.Table("user").Run(b.rdbc.GetSession())
	if err != nil {
		log.Fatalln(err)
	}

	var row map[string]interface{}
	err = res.One(&row)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print(row)

	return map[float64]dto.UserDTO{row["id"].(float64): dto.UserDTO{ID: row["id"].(float64),
		Name: row["name"].(string), Age: row["age"].(float64)}}
}

// DelUser deletes users contained in the slide
func (b *RethinkDBUserImpl) DelUser(u []dto.UserDTO) int {
	for _, v := range u {
		fmt.Printf("Deleting user: %s/n", v.Name)
	}

	return 0
}

// UpdateUser sets information to the users specify in u
func (b *RethinkDBUserImpl) UpdateUser(u []dto.UserDTO) int {
	fmt.Print("UpdateUser")
	return 0
}

// AddUser add the new users specify in u
func (b *RethinkDBUserImpl) AddUser(u []dto.UserDTO) int {
	fmt.Print("AddUser")
	return 0
}
