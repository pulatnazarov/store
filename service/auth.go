package service

import (
	"context"
	"test/api/models"
	"test/pkg/logger"
	"test/pkg/security"
	"test/storage"
)

type authService struct {
	storage storage.IStorage
	log     logger.ILogger
}

func NewAuthService(storage storage.IStorage, log logger.ILogger) authService {
	return authService{
		storage: storage,
		log:     log,
	}
}

func (a authService) CustomerLogin(ctx context.Context, loginRequest models.CustomerLoginRequest) error {
	customerPassword, err := a.storage.User().GetCustomerCredentialsByLogin(ctx, loginRequest.Login)
	if err != nil {
		return err
	}

	return security.CompareHashAndPassword(customerPassword, loginRequest.Password)
}
