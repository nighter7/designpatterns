package main

import (
	"fmt"

	"github.com/nighter7/designpatterns/dao/golang/dao"
	"github.com/nighter7/designpatterns/dao/golang/dto"
)

func main() {
	a := &dao.RethinkDBUserImpl{}
	b := []dto.UserDTO{
		dto.UserDTO{
			ID:   1,
			Name: "User 1",
			Age:  32,
		},
		{
			ID:   2,
			Name: "User 2",
			Age:  9,
		},
	}

	users := a.GetUser()

	for _, u := range users {
		fmt.Print(u)
	}

	a.DelUser(b)
}
