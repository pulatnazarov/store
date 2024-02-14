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

type categoryRepo struct {
	db  *pgxpool.Pool
	log logger.ILogger
}

func NewCategoryRepo(db *pgxpool.Pool, log logger.ILogger) storage.ICategoryStorage {
	return &categoryRepo{
		db:  db,
		log: log,
	}
}
func (c *categoryRepo) Create(ctx context.Context, category models.CreateCategory) (string, error) {
	id := uuid.New()
	query := `insert into categories (id, name) values($1, $2)`

	if rowsAffected, err := c.db.Exec(ctx, query, id, category.Name); err != nil {
		if r := rowsAffected.RowsAffected(); r == 0 {
			c.log.Error("error is in rows affected", logger.Error(err))
			return "", err
		}
		c.log.Error("error is while creating category", logger.Error(err))
		return "", err
	}

	return id.String(), nil
}

func (c *categoryRepo) GetByID(ctx context.Context, key models.PrimaryKey) (models.Category, error) {
	createdAt, updatedAt := sql.NullString{}, sql.NullString{}
	category := models.Category{}

	query := `select id, name, created_at, updated_at from categories where id = $1 and deleted_at = 0`
	if err := c.db.QueryRow(ctx, query, key.ID).Scan(&category.ID, &category.Name, &createdAt, &updatedAt); err != nil {
		c.log.Error("error is while getting by id", logger.Error(err))
		return models.Category{}, err
	}
	if createdAt.Valid {
		category.CreatedAt = createdAt.String
	}

	if updatedAt.Valid {
		category.UpdatedAt = updatedAt.String
	}
	return category, nil
}

func (c *categoryRepo) GetList(ctx context.Context, request models.GetListRequest) (models.CategoryResponse, error) {
	var (
		query, countQuery    string
		count                = 0
		page                 = request.Page
		offset               = (page - 1) * request.Limit
		search               = request.Search
		categories           = []models.Category{}
		createdAt, updatedAt = sql.NullString{}, sql.NullString{}
	)

	countQuery = `select count(1) from categories where deleted_at = 0`

	if search != "" {
		countQuery += fmt.Sprintf(` and name ilike '%%%s%%'`, search)
	}

	if err := c.db.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		c.log.Error("error is while scanning count", logger.Error(err))
		return models.CategoryResponse{}, err
	}

	query = `select id, name, created_at, updated_at from categories where deleted_at = 0`

	if search != "" {
		query += fmt.Sprintf(` and name ilike '%%%s%%' `, search)
	}

	query += ` order by created_at desc LIMIT $1 OFFSET $2`

	rows, err := c.db.Query(ctx, query, request.Limit, offset)
	if err != nil {
		c.log.Error("error is while selecting categories", logger.Error(err))
		return models.CategoryResponse{}, err
	}

	for rows.Next() {
		cat := models.Category{}
		if err = rows.Scan(&cat.ID, &cat.Name, &createdAt, &updatedAt); err != nil {
			c.log.Error("error is while scanning category", logger.Error(err))
			return models.CategoryResponse{}, err
		}
		if createdAt.Valid {
			cat.CreatedAt = createdAt.String
		}
		if updatedAt.Valid {
			cat.UpdatedAt = updatedAt.String
		}
		categories = append(categories, cat)
	}
	return models.CategoryResponse{
		Category: categories,
		Count:    count,
	}, err
}

func (c *categoryRepo) Update(ctx context.Context, category models.UpdateCategory) (string, error) {
	query := `update categories set name = $1, updated_at = now() where id = $2`

	if rowsAffected, err := c.db.Exec(ctx, query, &category.Name, &category.ID); err != nil {
		if r := rowsAffected.RowsAffected(); r == 0 {
			c.log.Error("error is in rows affected", logger.Error(err))
			return "", err
		}
		c.log.Error("error is while updating category", logger.Error(err))
		return "", err
	}
	return category.ID, nil
}

func (c *categoryRepo) Delete(ctx context.Context, key models.PrimaryKey) error {
	query := `update categories set deleted_at = extract(epoch from current_timestamp) where id = $1`

	if rowsAffected, err := c.db.Exec(ctx, query, key.ID); err != nil {
		if r := rowsAffected.RowsAffected(); r == 0 {
			c.log.Error("error is in rows affected", logger.Error(err))
			return err
		}
		c.log.Error("error is while deleting category", logger.Error(err))
		return err
	}
	return nil
}
