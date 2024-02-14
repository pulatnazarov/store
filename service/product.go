package service

import (
	"context"
	"encoding/json"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
)

type productService struct {
	storage storage.IStorage
	log     logger.ILogger
}

func NewProductService(storage storage.IStorage, log logger.ILogger) productService {
	return productService{storage: storage, log: log}
}

func (p productService) Create(ctx context.Context, product models.CreateProduct) (models.Product, error) {
	p.log.Info("product create service layer", logger.Any("product", product))

	id, err := p.storage.Product().Create(ctx, product)
	if err != nil {
		p.log.Error("error in service layer while creating product", logger.Error(err))
		return models.Product{}, err
	}

	createdProduct, err := p.storage.Product().GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		p.log.Error("error in service layer while getting by id", logger.Error(err))
		return models.Product{}, err
	}

	return createdProduct, nil
}

func (p productService) Get(ctx context.Context, key models.PrimaryKey) (models.Product, error) {
	product, err := p.storage.Product().GetByID(ctx, key)
	if err != nil {
		p.log.Error("error in service layer while getting by id", logger.Error(err))
		return models.Product{}, err
	}

	return product, nil
}

func (p productService) GetList(ctx context.Context, request models.GetListRequest) (models.ProductResponse, error) {
	p.log.Info("product get list service layer", logger.Any("product", request))

	products, err := p.storage.Product().GetList(ctx, request)
	if err != nil {
		p.log.Error("error in service layer while getting list", logger.Error(err))
		return models.ProductResponse{}, err
	}

	return products, nil
}

func (p productService) Update(ctx context.Context, product models.UpdateProduct) (models.Product, error) {
	id, err := p.storage.Product().Update(ctx, product)
	if err != nil {
		p.log.Error("error in service layer while update", logger.Error(err))
		return models.Product{}, err
	}

	updatedProduct, err := p.storage.Product().GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		p.log.Error("error in service layer while getting by id", logger.Error(err))
		return models.Product{}, err
	}

	return updatedProduct, nil
}

func (p productService) Delete(ctx context.Context, key models.PrimaryKey) error {
	err := p.storage.Product().Delete(ctx, key)

	return err
}

func (p productService) StartSellNew(ctx context.Context, request models.SellRequest) (models.ProductSell, error) {
	check := models.Check{
		Products: make([]models.Product, 0),
		TotalSum: 0,
	}

	productSell, err := p.storage.Product().Search(ctx, request.Products)
	if err != nil {
		p.log.Error("error in service layer while searching product", logger.Error(err))
		return models.ProductSell{}, err
	}

	basket, err := p.storage.Basket().GetByID(ctx, models.PrimaryKey{ID: request.BasketID})
	if err != nil {
		p.log.Error("error in service layer while getting basket by id", logger.Error(err))
		return models.ProductSell{}, err
	}

	customer, err := p.storage.User().GetByID(ctx, models.PrimaryKey{ID: basket.CustomerID})
	if err != nil {
		p.log.Error("error in service layer while getting user by id", logger.Error(err))
		return models.ProductSell{}, err
	}

	totalSum, profit := 0, float32(0.0)
	basketProducts := map[string]int{}

	for productID, price := range productSell.SelectedProducts.Products {
		customerQuantity := request.Products[productID]
		totalSum += price * customerQuantity

		//profit logic
		profit += float32(customerQuantity*price - productSell.ProductPrices[productID])
		basketProducts[productID] = customerQuantity
	}

	if customer.Cash < uint(totalSum) {
		p.log.Error("error in service layer while not enough customer cash", logger.Error(err))
		return models.ProductSell{}, err
	}

	if err = p.storage.User().UpdateCustomerCash(ctx, customer.ID, totalSum); err != nil {
		p.log.Error("error in service layer while updating customer cash", logger.Error(err))
		return models.ProductSell{}, err
	}

	if err = p.storage.Product().TakeProducts(ctx, basketProducts); err != nil {
		p.log.Error("error in service layer while taking product", logger.Error(err))
		return models.ProductSell{}, err
	}

	if err = p.storage.BasketProduct().AddProducts(ctx, basket.ID, basketProducts); err != nil {
		p.log.Error("error in service later while adding products to basket", logger.Error(err))
		return models.ProductSell{}, err
	}

	if err = p.storage.Store().AddProfit(ctx, profit, customer.BranchID); err != nil {
		p.log.Error("error in service layer while adding amount of profit", logger.Error(err))
		return models.ProductSell{}, err
	}

	// dealer

	//check
	productIDs := []string{}
	for productID := range productSell.SelectedProducts.Products {
		productIDs = append(productIDs, productID)
	}

	productsResp, err := p.storage.Product().GetListByIDs(ctx, productIDs)
	if err != nil {
		p.log.Error("error in service layer while getting products by ids", logger.Error(err))
		return models.ProductSell{}, err
	}

	js, _ := json.Marshal(productsResp.Products)

	json.Unmarshal(js, &check.Products)

	totalSum = 0
	for i, checkProduct := range check.Products {
		quantity := request.Products[checkProduct.ID]

		check.Products[i].Quantity = quantity

		totalSum += quantity * checkProduct.Price
	}

	check.TotalSum = totalSum

	productSell.Check = check

	//report

	return productSell, nil
}
