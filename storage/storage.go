package storage

import (
	"test/api/models"
)

type IStorage interface {
	Close()
	User() IUserStorage
	Basket() IBasket
	//Category() ICategoryStorage
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

//type ICategoryStorage interface {}

type IBasket interface {
	CreateBasket(models.CreateBasket) (models.Basket, error)
	GetBasketByID(models.PrimaryKey) (models.Basket, error)
	GetBasketList() (models.BasketResponse, error)
	//GetBasketList(models.GetListRequest) (models.BasketResponse, error)
	UpdateBasket(models.UpdateBasket) (models.Basket, error)
	DeleteBasket(key models.PrimaryKey) error
}
