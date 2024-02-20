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
	RedisService() redisService
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
	redisService         redisService
}

func New(storage storage.IStorage, log logger.ILogger, redis storage.IRedisStorage) Service {
	services := Service{}

	services.userService = NewUserService(storage, log, redis)
	services.categoryService = NewCategoryService(storage, log, redis)
	services.basketService = NewBasketService(storage, log, redis)
	services.basketProductService = NewBasketProductService(storage, log, redis)
	services.productService = NewProductService(storage, log, redis)
	services.branchService = NewBranchService(storage, log, redis)
	services.dealerService = NewDealerService(storage, log, redis)
	services.incomeService = NewIncomeService(storage, log, redis)
	services.incomeProductService = NewIncomeProductService(storage, log, redis)
	services.authService = NewAuthService(storage, log, redis)
	services.redisService = NewRedisService(storage, log, redis)

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

func (s Service) RedisService() redisService {
	return s.redisService
}
