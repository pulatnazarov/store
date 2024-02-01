package service

import (
	"context"
	"fmt"
	"test/api/models"
	"test/storage"
)

type categoryService struct {
	storage storage.IStorage
}

func NewCategoryService(storage storage.IStorage) categoryService {
	return categoryService{
		storage: storage,
	}
}

func (c categoryService) Create(ctx context.Context, createCategory models.CreateCategory) (models.Category, error) {
	pKey, err := c.storage.Category().Create(ctx, createCategory)
	if err != nil {
		fmt.Println("ERROR in service layer while creating category", err.Error())
		return models.Category{}, err
	}

	category, err := c.storage.Category().GetByID(ctx, models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		fmt.Println("ERROR in service layer while getting category", err.Error())
		return models.Category{}, err
	}

	return category, nil
}
