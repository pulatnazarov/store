package service

import (
	"context"
	"fmt"
	"test/api/models"
	"test/storage"
)

type incomeProductService struct {
	storage storage.IStorage
}

func NewIncomeProductService(storage storage.IStorage) incomeProductService {
	return incomeProductService{
		storage: storage,
	}
}

func (i incomeProductService) CreateMultiple(ctx context.Context, request models.CreateIncomeProducts) error {
	if err := i.storage.IncomeProduct().CreateMultiple(ctx, request); err != nil {
		fmt.Println("error while creating multiple income products", err.Error())
		return err
	}

	return nil
}

func (i incomeProductService) GetList(ctx context.Context, request models.GetListRequest) (models.IncomeProductsResponse, error) {
	incomeProducts, err := i.storage.IncomeProduct().GetList(ctx, request)
	if err != nil {
		fmt.Println("error in service layer while getting list", err.Error())
		return models.IncomeProductsResponse{}, err
	}
	return incomeProducts, nil
}

func (i incomeProductService) UpdateMultiple(ctx context.Context, response models.UpdateIncomeProducts) error {
	if err := i.storage.IncomeProduct().UpdateMultiple(ctx, response); err != nil {
		fmt.Println("error in service layer while updating", err.Error())
		return err
	}

	return nil
}

func (i incomeProductService) DeleteMultiple(ctx context.Context, response models.DeleteIncomeProducts) error {
	err := i.storage.IncomeProduct().DeleteMultiple(ctx, response)
	return err
}
