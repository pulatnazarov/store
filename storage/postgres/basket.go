package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
)

type basketRepo struct {
	db    *pgxpool.Pool
	log   logger.ILogger
	redis storage.IRedisStorage
}

func NewBasketRepo(db *pgxpool.Pool, log logger.ILogger, redis storage.IRedisStorage) storage.IBasketStorage {
	return &basketRepo{
		db:    db,
		log:   log,
		redis: redis,
	}
}

func (b *basketRepo) Create(ctx context.Context, basket models.CreateBasket) (string, error) {
	id := uuid.New()

	query := `insert into baskets(id, customer_id, total_sum) values($1, $2, $3)`
	if rowsAffected, err := b.db.Exec(ctx, query, id, basket.CustomerID, basket.TotalSum); err != nil {
		if r := rowsAffected.RowsAffected(); r == 0 {
			b.log.Error("error is in rows affected", logger.Error(err))
			return "", err
		}
		b.log.Error("error is while inserting basket data", logger.Error(err))
		return "", err
	}

	return id.String(), nil
}

func (b *basketRepo) GetByID(ctx context.Context, key models.PrimaryKey) (models.Basket, error) {
	var createdAt, updatedAt = sql.NullString{}, sql.NullString{}
	basket := models.Basket{}

	if err := b.db.QueryRow(ctx, `select id, customer_id, total_sum, created_at, updated_at from baskets where id = $1 and deleted_at = 0 `,
		key.ID).Scan(&basket.ID,
		&basket.CustomerID,
		&basket.TotalSum,
		&createdAt,
		&updatedAt,
	); err != nil {
		b.log.Error("error is while selecting basket", logger.Error(err))
		return models.Basket{}, err
	}

	if createdAt.Valid {
		basket.CreatedAt = createdAt.String
	}

	if updatedAt.Valid {
		basket.UpdatedAt = updatedAt.String
	}

	return basket, nil
}

func (b *basketRepo) GetList(ctx context.Context, req models.GetListRequest) (models.BasketResponse, error) {
	var (
		baskets              = []models.Basket{}
		count                = 0
		query, countQuery    string
		filter               string
		page                 = req.Page
		offset               = (page - 1) * req.Limit
		search               = req.Search
		createdAt, updatedAt = sql.NullString{}, sql.NullString{}
	)

	if req.UserID != "" {
		filter += fmt.Sprintf(" and customer_id = '%s'", req.UserID)
	}

	countQuery = `select count(1) from baskets where deleted_at = 0 ` + filter

	if search != "" {
		countQuery += fmt.Sprintf(` and CAST(total_sum AS TEXT) ilike '%%%s%%'`, search)
	}
	if err := b.db.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		b.log.Error("error is while selecting count", logger.Error(err))
		return models.BasketResponse{}, err
	}

	query = `select id, customer_id, total_sum, created_at, updated_at from baskets where deleted_at = 0 ` + filter

	if search != "" {
		query += fmt.Sprintf(` and CAST(total_sum AS TEXT) ilike '%%%s%%'`, search)
	}

	query += ` order by created_at desc LIMIT $1 OFFSET $2`
	rows, err := b.db.Query(ctx, query, req.Limit, offset)
	if err != nil {
		b.log.Error("error is while selecting baskets", logger.Error(err))
		return models.BasketResponse{}, err
	}

	for rows.Next() {
		basket := models.Basket{}
		if err = rows.Scan(&basket.ID, &basket.CustomerID, &basket.TotalSum, &createdAt, &updatedAt); err != nil {
			b.log.Error("error is while scanning data", logger.Error(err))
			return models.BasketResponse{}, err
		}

		if createdAt.Valid {
			basket.CreatedAt = createdAt.String
		}

		if updatedAt.Valid {
			basket.UpdatedAt = updatedAt.String
		}

		baskets = append(baskets, basket)

	}

	return models.BasketResponse{
		Baskets: baskets,
		Count:   count,
	}, nil
}

func (b *basketRepo) Update(ctx context.Context, basket models.UpdateBasket) (string, error) {
	bas := models.Basket{}

	if rowsAffected, err := b.db.Exec(ctx, `update baskets set customer_id = $1, total_sum = $2, updated_at = now() where id = $3 `,
		&basket.CustomerID,
		&basket.TotalSum,
		&basket.ID,
	); err != nil {
		if r := rowsAffected.RowsAffected(); r == 0 {
			b.log.Error("error is in rows affected", logger.Error(err))
			return "", err
		}
		return bas.ID, err
	}

	if err := b.db.QueryRow(ctx, `select id, customer_id, total_sum from baskets where id = $1`,
		basket.ID).Scan(&bas.ID, &bas.CustomerID, &bas.TotalSum); err != nil {
		b.log.Error("error is while selecting ", logger.Error(err))
		return "", err
	}
	return bas.ID, nil
}

func (b *basketRepo) Delete(ctx context.Context, key models.PrimaryKey) error {
	query := `update baskets set deleted_at = extract(epoch from current_timestamp) where id = $1`
	if rowsAffected, err := b.db.Exec(ctx, query, key.ID); err != nil {
		if r := rowsAffected.RowsAffected(); r == 0 {
			b.log.Error("error is while deleting basket", logger.Error(err))
			return err
		}
		return err
	}
	return nil
}
