package service

import (
	"context"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
)

type incomeProductService struct {
	storage storage.IStorage
	log     logger.ILogger
	redis   storage.IRedisStorage
}

func NewIncomeProductService(storage storage.IStorage, log logger.ILogger, redis storage.IRedisStorage) incomeProductService {
	return incomeProductService{
		storage: storage,
		log:     log,
		redis:   redis,
	}
}

func (i incomeProductService) CreateMultiple(ctx context.Context, createIncomeProducts models.CreateIncomeProducts) error {
	i.log.Info("income products create service layer", logger.Any("income products", createIncomeProducts))

	if err := i.storage.IncomeProduct().CreateMultiple(ctx, createIncomeProducts); err != nil {
		i.log.Error("error while creating multiple income products", logger.Error(err))
		return err
	}

	return nil
}

func (i incomeProductService) GetList(ctx context.Context, request models.GetListRequest) (models.IncomeProductsResponse, error) {
	i.log.Info("income products get list service layer", logger.Any("income products", request))

	incomeProducts, err := i.storage.IncomeProduct().GetList(ctx, request)
	if err != nil {
		i.log.Error("error in service layer while getting list", logger.Error(err))
		return models.IncomeProductsResponse{}, err
	}
	return incomeProducts, nil
}

func (i incomeProductService) UpdateMultiple(ctx context.Context, response models.UpdateIncomeProducts) error {
	if err := i.storage.IncomeProduct().UpdateMultiple(ctx, response); err != nil {
		i.log.Error("error in service layer while updating", logger.Error(err))
		return err
	}

	return nil
}

func (i incomeProductService) DeleteMultiple(ctx context.Context, response models.DeleteIncomeProducts) error {
	err := i.storage.IncomeProduct().DeleteMultiple(ctx, response)
	return err
}
