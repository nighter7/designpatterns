package dao

import "github.com/nighter7/designpatterns/dao/golang/dto"

type UserDAO interface {
	GetUser() map[int]dto.UserDTO
	DelUser(u []dto.UserDTO) int
	UpdateUser(u []dto.UserDTO) int
	AddUser(u []dto.UserDTO) int
}
