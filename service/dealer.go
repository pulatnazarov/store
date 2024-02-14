package service

import (
	"context"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
)

type dealerService struct {
	storage storage.IStorage
	log     logger.ILogger
}

func NewDealerService(storage storage.IStorage, log logger.ILogger) dealerService {
	return dealerService{storage: storage, log: log}
}

func (d dealerService) Delivery(ctx context.Context, sell models.ProductSell) error {
	var (
		totalSum = 0
	)

	for productID, quantity := range sell.NotEnoughProducts {
		totalSum += quantity * sell.NotEnoughProductPrices[productID]
	}

	budget, err := d.storage.Store().GetStoreBudget(ctx, sell.ProductsBranchID)
	if err != nil {
		d.log.Error("error in service layer while getting store budget", logger.Error(err))
		return err
	}

	if budget < float32(totalSum) {
		d.log.Error("not enough budget", logger.Error(err))
		return err
	}

	if err = d.storage.Product().AddDeliveredProducts(ctx, models.DeliverProducts{
		NotEnoughProducts: sell.NotEnoughProducts,
	}, sell.ProductsBranchID); err != nil {
		d.log.Error("error in service layer while adding delivered products", logger.Error(err))
		return err
	}

	if err = d.storage.Store().WithdrawalDeliveredSum(ctx, float32(totalSum), sell.ProductsBranchID); err != nil {
		d.log.Error("error in service layer while remove delivered sum", logger.Error(err))
		return err
	}

	dealerID := "1cfd84e6-72cb-4135-a802-85d10e4183ea"
	if err = d.storage.Dealer().AddSum(ctx, totalSum, dealerID); err != nil {
		d.log.Error("error in service layer while add sum to dealer", logger.Error(err))
		return err
	}

	return nil
}
