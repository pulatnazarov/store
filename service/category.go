package service

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
)

type categoryService struct {
	storage storage.IStorage
	log     logger.ILogger
}

func NewCategoryService(storage storage.IStorage, log logger.ILogger) categoryService {
	return categoryService{
		storage: storage,
		log:     log,
	}
}

func (c categoryService) Create(ctx context.Context, createCategory models.CreateCategory) (models.Category, error) {
	c.log.Info("category create service layer", logger.Any("category", createCategory))

	pKey, err := c.storage.Category().Create(ctx, createCategory)
	if err != nil {
		c.log.Error("ERROR in service layer while creating category", logger.Error(err))
		return models.Category{}, err
	}

	category, err := c.storage.Category().GetByID(ctx, models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		c.log.Error("ERROR in service layer while getting category", logger.Error(err))
		return models.Category{}, err
	}

	return category, nil
}

func (c categoryService) Get(ctx context.Context, key models.PrimaryKey) (models.Category, error) {
	category, err := c.storage.Category().GetByID(ctx, key)
	if err != nil {
		c.log.Error("error is in service layer while getting by id", logger.Error(err))
		return models.Category{}, err
	}
	return category, nil
}

func (c categoryService) GetList(ctx context.Context, request models.GetListRequest) (models.CategoryResponse, error) {
	c.log.Info("category get list service layer", logger.Any("category", request))

	categories, err := c.storage.Category().GetList(ctx, request)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			c.log.Error("error in service layer while getting list", logger.Error(err))
			return models.CategoryResponse{}, err
		}
	}
	return categories, nil
}

func (c categoryService) Update(ctx context.Context, category models.UpdateCategory) (models.Category, error) {
	id, err := c.storage.Category().Update(ctx, category)
	if err != nil {
		c.log.Error("error in service layer while updating category", logger.Error(err))
		return models.Category{}, err
	}

	updatedCategory, err := c.storage.Category().GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		c.log.Error("error in service layer while getting by id", logger.Error(err))
		return models.Category{}, err
	}

	return updatedCategory, nil
}

func (c categoryService) Delete(ctx context.Context, key models.PrimaryKey) error {
	err := c.storage.Category().Delete(ctx, key)

	return err
}
