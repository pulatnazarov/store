package postgres

import (
	"context"
	"fmt"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
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

func (r *report) Report(ctx context.Context, request models.ProductReportListRequest) (models.ProductReportList, error) {

	var (
		page              = request.Page
		offset            = (page - 1) * request.Limit
		reportproducts    = []models.ProductReport{}
		from              = request.From
		to                = request.To
		branchid          = request.BranchId
		query, countQuery string
		overal            string
		overalPrice       int
		count             = 0
	)
	fromtime, err := time.Parse("15:04", from)
	if err != nil {
		fmt.Println("error while parsing time ", err.Error())
		r.log.Error("error",logger.Error(err))
		return models.ProductReportList{}, err
	}

	toTime, err := time.Parse("15:04", to)
	if err != nil {
		fmt.Println("error while parsing time ", err.Error())
		r.log.Error("error",logger.Error(err))
		return models.ProductReportList{}, err
	}

	countQuery = `SELECT COUNT(1) FROM products WHERE deleted_at =0`
	if from != "" && to != "" {
		countQuery += fmt.Sprintf(" AND created_at BETWEEN '%s' AND '%s'", fromtime.Format("2006-01-02 15:04:05"), toTime.Format("2006-01-02 15:04:05"))
	} else if from != "" {
		countQuery += fmt.Sprintf(" AND created_at >= '%s'", fromtime.Format("2006-01-02 15:04:05"))
	} else if to != "" {
		countQuery += fmt.Sprintf(" and created_at <='%s'", toTime.Format("2006-01-02 15:04:05"))
	}
	if branchid != "" {
		countQuery += fmt.Sprintf(" and branch_id ='%s'", request.BranchId)
	}

	query = `SELECT name, quantity, price , quantity*price as total_price FROM products WHERE deleted_at =0` + countQuery

	overal = ` select sum(quantity*price) as overall price from products where` + countQuery

	query += ` LIMIT $1 OFFSET $2`

	if err := r.db.QueryRow(ctx, overal).Scan(&overalPrice); err != nil {

		fmt.Println("error in postgres ", err.Error())
		r.log.Error("error",logger.Error(err))
		return models.ProductReportList{}, err
	}

	rows, err := r.db.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error")
		r.log.Error("error",logger.Error(err))
		return models.ProductReportList{}, err
	}
	for rows.Next() {
		reportproduct := models.ProductReport{}
		if err = rows.Scan(&reportproduct.ProductName, &reportproduct.Quantity, &reportproduct.Price, &reportproduct.TotalPrice); err != nil {
			fmt.Println("error")
			r.log.Error("error",logger.Error(err))
			return models.ProductReportList{}, err
		}

		reportproducts = append(reportproducts, reportproduct)
	}
	return models.ProductReportList{
		Products:     reportproducts,
		OverallPrice: overalPrice,
		Count:        count,
	}, nil

}

func (r *report) IncomeReport(ctx context.Context, request models.IncomeProductReportListRequest) (models.IncomeProductReportList, error) {

	var (
		page                      = request.Page
		offset                    = (page - 1) * request.Limit
		incomereports             = []models.IncomeProductReport{}
		from                      = request.From
		to                        = request.To
		branchID                  = request.BranchID
		query, countQuery, overal string
		count, overallPrice       int
	)
	fromtime, err := time.Parse("15:04", from)
	if err != nil {
		fmt.Println("error while parsing time ", err.Error())
		r.log.Error("error",logger.Error(err))
		return models.IncomeProductReportList{}, err
	}

	toTime, err := time.Parse("15:04", to)
	if err != nil {
		fmt.Println("error while parsing time ", err.Error())
		r.log.Error("error",logger.Error(err))
		return models.IncomeProductReportList{}, err
	}
	countQuery = ` select count(1) from income_products where deleted_at=0 `
	if from != "" && to != "" {
		countQuery += fmt.Sprintf(" and created_at between '%s' and '%s'", fromtime.Format("2006-01-02 15:04:05"), toTime.Format("2006-01-02 15:04:05"))
	} else if from != "" {
		countQuery += fmt.Sprintf(" and created_at <='%s'", fromtime.Format("2006-01-02 15:04:05"))
	} else if to != "" {
		countQuery += fmt.Sprintf(" and created_at <='%s'", toTime.Format("2006-01-02 15:04:05"))
	}
	if branchID != "" {
		countQuery += fmt.Sprintf(" and branch_id='%s'", request.BranchID)
	}

	query = ` select p.name, p.quantity, p.price, p.quantity*p.price as total_price from products as p left join income_products as ip on ip.id=p.id` + countQuery
	overal = `
    SELECT SUM(p.quantity * p.price) AS overall_price
    FROM products AS p
    JOIN income_products AS ip ON p.id = ip.productID
    WHERE your_condition_here
` + countQuery

	query += ` LIMIT $1 OFFSET $2`
	if err := r.db.QueryRow(ctx, overal).Scan(&overallPrice); err != nil {
		fmt.Println("error in postgres while scanning overall price of income products")
		r.log.Error("error",logger.Error(err))
		return models.IncomeProductReportList{}, err
	}
	rows, err := r.db.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error")
		r.log.Error("error",logger.Error(err))
		return models.IncomeProductReportList{}, err
	}
	for rows.Next() {
		incomeproduct := models.IncomeProductReport{}
		if err = rows.Scan(&incomeproduct.IncomeProdutName, &incomeproduct.Quantity, &incomeproduct.Price, &incomeproduct.TotalPrice); err != nil {
			fmt.Println("error", err.Error())
			r.log.Error("error",logger.Error(err))
			return models.IncomeProductReportList{}, err
		}
		incomereports = append(incomereports, incomeproduct)
	}

	return models.IncomeProductReportList{
		IncomeProducts: incomereports,
		OverallPrice:   overallPrice,
		Count:          count,
	}, nil

}
