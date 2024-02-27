package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
)

type report struct {
	db  *pgxpool.Pool
	log logger.ILogger
}

func NewReport(db *pgxpool.Pool, log logger.ILogger) storage.IReportStorage {
	return &report{
		db:  db,
		log: log,
	}
}

func (r *report) ProductReportList(ctx context.Context, request models.ReportRequest) (models.ProductReportList, error) {
	var (
		page                                         = request.Page
		offset                                       = (page - 1) * request.Limit
		overallPriceQuery, pagination, query, filter string
		overallPrice                                 int
		products                                     []models.ProductReport
	)

	pagination = ` limit $1 offset $2`

	if request.From != "" && request.To != "" {
		filter += fmt.Sprintf(` and created_at::text between '%s' and '%s' `, request.From, request.To)
	} else if request.From != "" {
		filter += fmt.Sprintf(` and created_at::text <= '%s' `, request.From)
	} else if request.To != "" {
		filter += fmt.Sprintf(` and created_at::text >= '%s' `, request.To)
	}

	if request.BranchID != "" {
		filter += fmt.Sprintf(` and branch_id = '%s'`, request.BranchID)
	}

	query = `select name, quantity, price, quantity*price as total_price from products where deleted_at = 0 ` + filter + pagination

	overallPriceQuery = `select sum(price*quantity) as overall_price from products where deleted_at = 0 ` + filter

	if err := r.db.QueryRow(ctx, overallPriceQuery).Scan(&overallPrice); err != nil {
		r.log.Error("error is while scanning overall price", logger.Error(err), logger.Any("overallQuery", overallPriceQuery))
		return models.ProductReportList{}, err
	}

	rows, err := r.db.Query(ctx, query, request.Limit, offset)
	if err != nil {
		r.log.Error("error is while all selecting products", logger.Error(err), logger.Any("main query", query))
		return models.ProductReportList{}, err
	}

	for rows.Next() {
		product := models.ProductReport{}

		if err := rows.Scan(
			&product.ProductName,
			&product.Quantity,
			&product.Price,
			&product.TotalPrice,
		); err != nil {
			return models.ProductReportList{}, err
		}

		products = append(products, product)
	}

	return models.ProductReportList{
		Products:     products,
		OverallPrice: overallPrice,
	}, nil
}

func (r *report) IncomeProductList(ctx context.Context, request models.ReportRequest) (models.IncomeProductReportList, error) {
	var (
		page                      = request.Page
		offset                    = (page - 1) * request.Limit
		pagination, query, filter string
		overallPrice              int
		incomeProducts            []models.IncomeProductReport
	)
	pagination = ` limit $1 offset $2`

	if request.From != "" {
		filter += fmt.Sprintf(` and i.created_at::text <= '%s'`, request.From)
	}

	if request.To != "" {
		filter += fmt.Sprintf(` and i.created_at::text >= '%s'`, request.To)
	}

	if request.BranchID != "" {
		filter += fmt.Sprintf(` and i.branch_id = '%s'`, request.BranchID)
	}

	query = `select p.name, i.quantity, i.price, i.quantity*i.price as total_price from income_products as i 
                INNER JOIN products as p ON i.product_id = p.id where i.deleted_at = 0 ` + filter + pagination

	rows, err := r.db.Query(ctx, query, request.Limit, offset)
	if err != nil {
		r.log.Error("error is while selecting all from income products", logger.Error(err))
		return models.IncomeProductReportList{}, err
	}
	for rows.Next() {
		incomeProduct := models.IncomeProductReport{}
		if err := rows.Scan(
			&incomeProduct.ProductName,
			&incomeProduct.Quantity,
			&incomeProduct.Price,
			&incomeProduct.TotalPrice,
		); err != nil {
			r.log.Error("error is while scanning all from income products", logger.Error(err))
			return models.IncomeProductReportList{}, err
		}
		overallPrice += incomeProduct.TotalPrice
		incomeProducts = append(incomeProducts, incomeProduct)
	}
	return models.IncomeProductReportList{
		IncomeProducts: incomeProducts,
		OverallPrice:   overallPrice,
	}, err
}
