package service

import (
	"context"
	"test/api/models"
	"test/pkg/jwt"
	"test/pkg/logger"
	"test/pkg/security"
	"test/storage"
)

type authService struct {
	storage storage.IStorage
	log     logger.ILogger
	redis   storage.IRedisStorage
}

func NewAuthService(storage storage.IStorage, log logger.ILogger, redis storage.IRedisStorage) authService {
	return authService{
		storage: storage,
		log:     log,
		redis:   redis,
	}
}

func (a authService) CustomerLogin(ctx context.Context, loginRequest models.CustomerLoginRequest) (models.CustomerLoginResponse, error) {
	customer, err := a.storage.User().GetCustomerCredentialsByLogin(ctx, loginRequest.Login)
	if err != nil {
		a.log.Error("error while getting customer credentials by login", logger.Error(err))
		return models.CustomerLoginResponse{}, err
	}

	if err = security.CompareHashAndPassword(customer.Password, loginRequest.Password); err != nil {
		a.log.Error("error while comparing password", logger.Error(err))
		return models.CustomerLoginResponse{}, err
	}

	m := make(map[interface{}]interface{})

	m["user_id"] = customer.ID
	m["user_role"] = "customer"

	accessToken, refreshToken, err := jwt.GenerateJWT(m)
	if err != nil {
		a.log.Error("error while generating tokens for customer login", logger.Error(err))
		return models.CustomerLoginResponse{}, err
	}

	return models.CustomerLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a authService) AdminLogin(ctx context.Context, loginRequest models.AdminLoginRequest) (models.AdminLoginResponse, error) {
	admin, err := a.storage.User().GetAdminCredentialsByLogin(ctx, loginRequest.Login)
	if err != nil {
		a.log.Error("error is while getting admin", logger.Error(err))
		return models.AdminLoginResponse{}, err
	}

	if err := security.CompareHashAndPassword(admin.Password, loginRequest.Password); err != nil {
		a.log.Error("password is incorrect", logger.Error(err))
		return models.AdminLoginResponse{}, err
	}

	m := make(map[interface{}]interface{})
	m["user_id"] = admin.ID
	m["user_role"] = "admin"

	accessToken, refreshToken, err := jwt.GenerateJWT(m)
	if err != nil {
		a.log.Error("error while generate jwt token", logger.Error(err))
		return models.AdminLoginResponse{}, err
	}

	return models.AdminLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
