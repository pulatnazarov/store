package storage

import (
	"test/api/models"
)

type IStorage interface {
	Close()
	User() IUserStorage
	Category() ICategoryStorage
}

type IUserStorage interface {
	Create(models.CreateUser) (string, error)
	GetByID(models.PrimaryKey) (models.User, error)
	GetList(models.GetListRequest) (models.UsersResponse, error)
	Update(models.UpdateUser) (string, error)
	Delete(models.PrimaryKey) error
	GetPassword(id string) (string, error)
	UpdatePassword(password models.UpdateUserPassword) error
}

type ICategoryStorage interface {
}
