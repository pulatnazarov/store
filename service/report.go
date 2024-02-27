package service

import (
	"context"
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

func (r reportService) ReportProduct(ctx context.Context, request models.ProductReportListRequest) (models.ProductReportList, error) {
	r.log.Info("report of product service layer", logger.Any("report of income product", request))
	list, err := r.storage.Report().Report(ctx, request)
	if err != nil {
		r.log.Error("error in service layer", logger.Error(err))
	}
	return list, nil

}

func (r reportService) ReportIncome(ctx context.Context, request models.IncomeProductReportListRequest) (models.IncomeProductReportList, error) {
	r.log.Info("report of income product report", logger.Any("report", request))
	list, err := r.storage.Report().IncomeReport(ctx, request)
	if err != nil {
		r.log.Error("error", logger.Error(err))
	}
	return list, nil
}
