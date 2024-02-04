package service

import (
	"context"
	"fmt"
	"math/rand"
	"test/api/models"
	"test/storage"
)

type dealerService struct {
	storage storage.IStorage
}

func NewDealerService(storage storage.IStorage) dealerService {
	return dealerService{storage}
}

func (d dealerService) Delivery(ctx context.Context, sell models.ProductSell) error {
	var (
		sum, newSum, totalSum = 0, 0, 0
		newProductPrices      = make(map[string]int)
	)

	for productID, quantity := range sell.NotEnoughProducts {
		sum += quantity * sell.Prices[productID]
	}

	for productID, quantity := range sell.NewProducts {
		originalPrice := rand.Intn(20000) + 1000
		newSum += quantity * originalPrice
		newProductPrices[productID] = originalPrice
	}

	totalSum = sum + newSum

	budget, err := d.storage.Store().GetStoreBudget(ctx, sell.ProductsBranchID)
	if err != nil {
		fmt.Println("error in service layer while getting store budget", err.Error())
		return err
	}

	if budget < float32(totalSum) {
		fmt.Println("not enough budget", err.Error())
		return err
	}

	if err = d.storage.Product().AddDeliveredProducts(ctx, models.DeliverProducts{
		NotEnoughProducts: sell.NotEnoughProducts,
		NewProducts:       sell.NewProducts,
		NewProductPrices:  newProductPrices,
	}, sell.ProductsBranchID); err != nil {
		fmt.Println("error in service layer while adding delivered products", err.Error())
		return err
	}

	if err = d.storage.Store().RemoveDeliveredSum(ctx, float32(totalSum), sell.ProductsBranchID); err != nil {
		fmt.Println("error in service layer while remove delivered sum", err.Error())
		return err
	}

	if err = d.storage.Dealer().AddSum(ctx, totalSum); err != nil {
		fmt.Println("error in service layer while add sum to dealer", err.Error())
		return err
	}

	return nil
}
