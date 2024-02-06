package service

import (
	"test/storage"
)

type IServiceManager interface {
	User() userService
	Category() categoryService
	Basket() basketService
	BasketProduct() basketProductService
	Product() productService
	Branch() branchService
	Dealer() dealerService
}

type Service struct {
	userService          userService
	categoryService      categoryService
	basketService        basketService
	basketProductService basketProductService
	productService       productService
	branchService        branchService
	dealerService        dealerService
}

func New(storage storage.IStorage) Service {
	services := Service{}

	services.userService = NewUserService(storage)
	services.categoryService = NewCategoryService(storage)
	services.basketService = NewBasketService(storage)
	services.basketProductService = NewBasketProductService(storage)
	services.productService = NewProductService(storage)
	services.branchService = NewBranchService(storage)
	services.dealerService = NewDealerService(storage)

	return services
}

func (s Service) User() userService {
	return s.userService
}

func (s Service) Category() categoryService {
	return s.categoryService
}

func (s Service) Basket() basketService {
	return s.basketService
}

func (s Service) BasketProduct() basketProductService {
	return s.basketProductService
}

func (s Service) Product() productService {
	return s.productService
}

func (s Service) Branch() branchService {
	return s.branchService
}

func (s Service) Dealer() dealerService {
	return s.dealerService
}
