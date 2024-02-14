package service

import (
	"context"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
)

type basketService struct {
	storage storage.IStorage
	log     logger.ILogger
}

func NewBasketService(storage storage.IStorage, log logger.ILogger) basketService {
	return basketService{
		storage: storage,
		log:     log,
	}
}

func (b basketService) Create(ctx context.Context, basket models.CreateBasket) (models.Basket, error) {
	b.log.Info("basket create service layer", logger.Any("basket", basket))
	id, err := b.storage.Basket().Create(ctx, basket)
	if err != nil {
		b.log.Error("error in service layer while creating basket", logger.Error(err))
		return models.Basket{}, err
	}

	createdBasket, err := b.storage.Basket().GetByID(context.Background(), models.PrimaryKey{ID: id})
	if err != nil {
		b.log.Error("error is while getting by id", logger.Error(err))
		return models.Basket{}, err
	}

	return createdBasket, err
}

func (b basketService) Get(ctx context.Context, id string) (models.Basket, error) {
	basket, err := b.storage.Basket().GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		b.log.Error("error in service layer while getting by id", logger.Error(err))
		return models.Basket{}, err
	}

	return basket, nil
}

func (b basketService) GetList(ctx context.Context, request models.GetListRequest) (models.BasketResponse, error) {
	b.log.Info("basket get list service layer", logger.Any("basket", request))

	baskets, err := b.storage.Basket().GetList(ctx, request)
	if err != nil {
		b.log.Error("error in service layer  while getting list", logger.Error(err))
		return models.BasketResponse{}, err
	}

	return baskets, nil
}

func (b basketService) Update(ctx context.Context, basket models.UpdateBasket) (models.Basket, error) {
	id, err := b.storage.Basket().Update(ctx, basket)
	if err != nil {
		b.log.Error("error in service layer while updating", logger.Error(err))
		return models.Basket{}, err
	}

	updatedBasket, err := b.storage.Basket().GetByID(context.Background(), models.PrimaryKey{ID: id})
	if err != nil {
		b.log.Error("error in service layer while getting basket by id", logger.Error(err))
		return models.Basket{}, err
	}

	return updatedBasket, nil
}

func (b basketService) Delete(ctx context.Context, key models.PrimaryKey) error {
	err := b.storage.Basket().Delete(ctx, key)

	return err
}
