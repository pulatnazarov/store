package service

import (
	"context"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
)

type incomeService struct {
	storage storage.IStorage
	log     logger.ILogger
}

func NewIncomeService(storage storage.IStorage, log logger.ILogger) incomeService {
	return incomeService{
		storage: storage,
		log:     log,
	}
}

func (i incomeService) Create(ctx context.Context) (models.Income, error) {
	income, err := i.storage.Income().Create(ctx)
	if err != nil {
		i.log.Error("error while creating income ", logger.Error(err))
		return models.Income{}, err
	}

	return income, nil
}

func (i incomeService) Get(ctx context.Context, key models.PrimaryKey) (models.Income, error) {
	income, err := i.storage.Income().GetByID(ctx, key)
	if err != nil {
		i.log.Error("error in service layer while getting by id", logger.Error(err))
		return models.Income{}, err
	}
	return income, nil
}

func (i incomeService) GetList(ctx context.Context, request models.GetListRequest) (models.IncomesResponse, error) {
	i.log.Info("category create service layer", logger.Any("category", request))

	incomes, err := i.storage.Income().GetList(ctx, request)
	if err != nil {
		i.log.Error("error in service layer while getting list", logger.Error(err))
		return models.IncomesResponse{}, err
	}
	return incomes, nil
}

func (i incomeService) Delete(ctx context.Context, key models.PrimaryKey) error {
	err := i.storage.Income().Delete(ctx, key)
	return err
}
