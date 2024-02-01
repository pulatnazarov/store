package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"test/api/models"
	"test/storage"
)

type basketRepo struct {
	db *pgxpool.Pool
}

func NewBasketRepo(db *pgxpool.Pool) storage.IBasketStorage {
	return &basketRepo{db: db}
}

func (b *basketRepo) Create(ctx context.Context, basket models.CreateBasket) (string, error) {
	id := uuid.New()

	if _, err := b.db.Exec(ctx, `insert into baskets(id, customer_id, total_sum)
	values($1, $2, $3)`, id, basket.CustomerID, basket.TotalSum); err != nil {
		fmt.Println("error is while inserting data", err.Error())
		return "", err
	}

	return id.String(), nil
}

func (b *basketRepo) GetByID(ctx context.Context, key models.PrimaryKey) (models.Basket, error) {
	basket := models.Basket{}

	if err := b.db.QueryRow(ctx, `select id, customer_id, total_sum from baskets where id = $1`,
		key.ID).Scan(&basket.ID,
		&basket.CustomerID,
		&basket.TotalSum); err != nil {
		fmt.Println("error is while selecting basket", err.Error())
		return models.Basket{}, err
	}

	return basket, nil
}

func (b *basketRepo) GetList(ctx context.Context, req models.GetListRequest) (models.BasketResponse, error) {
	var (
		baskets           = []models.Basket{}
		count             = 0
		query, countQuery string
		page              = req.Page
		offset            = (page - 1) * req.Limit
		search            = req.Search
	)

	countQuery = `select count(1) from baskets `

	if search != "" {
		countQuery += fmt.Sprintf(` where CAST(total_sum AS TEXT) ilike '%%%s%%'`, search)
	}
	if err := b.db.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		fmt.Println("error is while selecting count", err.Error())
		return models.BasketResponse{}, err
	}

	query = `select id, customer_id, total_sum from baskets `

	if search != "" {
		query += fmt.Sprintf(` where CAST(total_sum AS TEXT) ilike '%%%s%%'`, search)
	}

	query += `LIMIT $1 OFFSET $2`
	rows, err := b.db.Query(ctx, query, req.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting baskets", err.Error())
		return models.BasketResponse{}, err
	}

	for rows.Next() {
		b := models.Basket{}
		if err = rows.Scan(&b.ID, &b.CustomerID, &b.TotalSum); err != nil {
			fmt.Println("error is while scanning data", err.Error())
			return models.BasketResponse{}, err
		}
		baskets = append(baskets, b)

	}

	return models.BasketResponse{
		Baskets: baskets,
		Count:   count,
	}, nil
}

func (b *basketRepo) Update(ctx context.Context, basket models.UpdateBasket) (string, error) {
	bas := models.Basket{}

	if _, err := b.db.Exec(ctx, `update baskets set customer_id = $1, total_sum = $2 where id = $3`, &basket.CustomerID, &basket.TotalSum, &basket.ID); err != nil {
		return "", err
	}

	if err := b.db.QueryRow(ctx, `select id, customer_id, total_sum from baskets where id = $1`, basket.ID).Scan(&bas.ID, &bas.CustomerID, &bas.TotalSum); err != nil {
		fmt.Println("error is while selecting ", err.Error())
		return "", err
	}
	return bas.ID, nil
}

func (b *basketRepo) Delete(ctx context.Context, key models.PrimaryKey) error {
	if _, err := b.db.Exec(ctx, `delete from baskets where id = $1`, key.ID); err != nil {
		return err
	}
	return nil
}
