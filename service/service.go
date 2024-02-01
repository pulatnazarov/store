package service

import (
	"test/storage"
)

type IServiceManager interface {
	User() userService
	Category() categoryService
}

type Service struct {
	userService     userService
	categoryService categoryService
}

func New(storage storage.IStorage) Service {
	services := Service{}

	services.userService = NewUserService(storage)
	services.categoryService = NewCategoryService(storage)

	return services
}

func (s Service) User() userService {
	return s.userService
}

func (s Service) Category() categoryService {
	return s.categoryService
}
