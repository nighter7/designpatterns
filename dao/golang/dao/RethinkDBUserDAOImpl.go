package dao

import (
	"fmt"

	"github.com/nighter7/designpatterns/dao/golang/dto"
)

type RethinkDBUserDAOImpl struct {
}

func (c *RethinkDBUserDAOImpl) GetUser() map[int]dto.UserDTO {
	fmt.Print("GetUser")
	return nil
}

func (c *RethinkDBUserDAOImpl) DelUser(u []dto.UserDTO) int {
	for _, v := range u {
		fmt.Printf("Deleting user: %s\n", v.Name)
	}

	return 0
}

func (c *RethinkDBUserDAOImpl) UpdateUser(u []dto.UserDTO) int {
	fmt.Print("UpdateUser")
	return 0
}

func (c *RethinkDBUserDAOImpl) AddUser(u []dto.UserDTO) int {
	fmt.Print("AddUser")
	return 0
}
