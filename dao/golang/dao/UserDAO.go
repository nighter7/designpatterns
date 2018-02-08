package dao

import "github.com/nighter7/designpatterns/dao/golang/dto"

// UserDAO respresents all possible acctions to manipulate users
type UserDAO interface {
	GetUser() map[float64]dto.UserDTO
	DelUser(u []dto.UserDTO) int
	UpdateUser(u []dto.UserDTO) int
	AddUser(u []dto.UserDTO) int
}
