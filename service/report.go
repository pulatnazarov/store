package service

import (
	"golang.org/x/net/context"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
)

type reportService struct {
	storage storage.IStorage
	log     logger.ILogger
}

func NewReportService(storage storage.IStorage, log logger.ILogger) reportService {
	return reportService{
		storage: storage,
		log:     log,
	}
}

func (r reportService) ProductReportList(ctx context.Context, request models.ProductRepoRequest) (models.ProductReportList, error) {
	productList, err := r.storage.Report().ProductReportList(ctx, request)
	if err != nil {
		r.log.Error("error is while getting product report list", logger.Error(err))
		return models.ProductReportList{}, err
	}
	return productList, nil
}
