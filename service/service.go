package service

import (
	"test/pkg/logger"
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
	Income() incomeService
	IncomeProduct() incomeProductService
}

type Service struct {
	userService          userService
	categoryService      categoryService
	basketService        basketService
	basketProductService basketProductService
	productService       productService
	branchService        branchService
	dealerService        dealerService
	incomeService        incomeService
	incomeProductService incomeProductService
}

func New(storage storage.IStorage, log logger.ILogger) Service {
	services := Service{}

	services.userService = NewUserService(storage, log)
	services.categoryService = NewCategoryService(storage, log)
	services.basketService = NewBasketService(storage, log)
	services.basketProductService = NewBasketProductService(storage, log)
	services.productService = NewProductService(storage, log)
	services.branchService = NewBranchService(storage, log)
	services.dealerService = NewDealerService(storage, log)
	services.incomeService = NewIncomeService(storage, log)
	services.incomeProductService = NewIncomeProductService(storage, log)

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

func (s Service) Income() incomeService {
	return s.incomeService
}

func (s Service) IncomeProduct() incomeProductService {
	return s.incomeProductService
}
