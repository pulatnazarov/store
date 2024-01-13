package storage

import "test/models"

type IStorage interface {
	Close()
	User() IUserStorage
}

type IUserStorage interface {
	Create(models.CreateUser) (models.User, error)
	GetByID(models.PrimaryKey) (models.User, error)
	GetList(models.GetListRequest) (models.UsersResponse, error)
	Update(models.UpdateUser) (models.User, error)
	Delete(models.PrimaryKey) error
}
