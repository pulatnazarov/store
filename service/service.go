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
	AuthService() authService
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
	authService          authService
}

func New(storage storage.IStorage, log logger.ILogger, redis storage.IRedisStorage) Service {
	services := Service{}

	services.userService = NewUserService(storage, log, redis)
	services.categoryService = NewCategoryService(storage, log)
	services.basketService = NewBasketService(storage, log)
	services.basketProductService = NewBasketProductService(storage, log)
	services.productService = NewProductService(storage, log, redis)
	services.branchService = NewBranchService(storage, log)
	services.dealerService = NewDealerService(storage, log)
	services.incomeService = NewIncomeService(storage, log)
	services.incomeProductService = NewIncomeProductService(storage, log)
	services.authService = NewAuthService(storage, log)

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

func (s Service) AuthService() authService {
	return s.authService
}
