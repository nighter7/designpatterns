package main

import (
	"github.com/nighter7/designpatterns/dao/golang/dao"
	"github.com/nighter7/designpatterns/dao/golang/dto"
)

var a dao.UserDAO

func main() {
	a = new(dao.RethinkDBUserDAOImpl)
	b := []dto.UserDTO{
		dto.UserDTO{
			1,
			"User 1",
			32,
		},
		{
			2,
			"User 2",
			9,
		},
	}

	a.DelUser(b)
}
