package service

import (
	"project_mongodb-go/entity"
)

type UserServices interface {
	CreateUser(*entity.User) error
	GetUser(*string) (*entity.User, error)
	GetAll() ([]*entity.User, error)
	UpdateUser(*entity.User) error
	DeleteUser(*string) error
}
