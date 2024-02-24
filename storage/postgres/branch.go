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

type branchRepo struct {
	db    *pgxpool.Pool
	log   logger.ILogger
	redis storage.IRedisStorage
}

func NewBranchRepo(db *pgxpool.Pool, log logger.ILogger, redis storage.IRedisStorage) storage.IBranchStorage {
	return branchRepo{
		db:    db,
		log:   log,
		redis: redis,
	}
}

func (b branchRepo) Create(ctx context.Context, branch models.CreateBranch) (string, error) {
	branchID := uuid.New()

	query := `insert into branches (id, name, address, phone_number) 
									values($1, $2, $3, $4)`

	if rowsAffected, err := b.db.Exec(ctx, query,
		branchID,
		branch.Name,
		branch.Address,
		branch.PhoneNumber,
	); err != nil {
		if r := rowsAffected.RowsAffected(); r == 0 {
			b.log.Error("error is rows affected", logger.Error(err))
			return "", err
		}
		b.log.Error("error is while inserting branch data", logger.Error(err))
		return "", err
	}

	storeQuery := `insert into store(id, branch_id, profit, budget) values($1, $2, 0, 1000.0)`
	if rowsAffected, err := b.db.Exec(ctx, storeQuery, uuid.New(), branchID); err != nil {
		if r := rowsAffected.RowsAffected(); r == 0 {
			b.log.Error("error is in rows affected", logger.Error(err))
			return "", err
		}
		b.log.Error("error is while inserting store data", logger.Error(err))
		return "", err
	}

	return branchID.String(), nil
}

func (b branchRepo) GetByID(ctx context.Context, key models.PrimaryKey) (models.Branch, error) {
	var createdAt, updatedAt = sql.NullString{}, sql.NullString{}
	branch := models.Branch{}
	query := `select  id, name, address, phone_number, created_at, updated_at 
					from branches where id = $1 and deleted_at = 0
`
	if err := b.db.QueryRow(ctx, query, key.ID).Scan(
		&branch.ID,
		&branch.Name,
		&branch.Address,
		&branch.PhoneNumber,
		&createdAt,
		&updatedAt); err != nil {
		b.log.Error("error is while selecting by id", logger.Error(err))
		return models.Branch{}, err
	}

	if createdAt.Valid {
		branch.CreatedAt = createdAt.String
	}

	if updatedAt.Valid {
		branch.UpdatedAt = updatedAt.String
	}

	return branch, nil
}

func (b branchRepo) GetList(ctx context.Context, request models.GetListRequest) (models.BranchResponse, error) {
	var (
		count                = 0
		branches             = []models.Branch{}
		query, countQuery    string
		page                 = request.Page
		offset               = (page - 1) * request.Limit
		search               = request.Search
		createdAt, updatedAt = sql.NullString{}, sql.NullString{}
	)

	countQuery = `select count(1) from branches where deleted_at = 0 `

	if search != "" {
		countQuery += fmt.Sprintf(` and name ilike '%s'`, search)
	}

	if err := b.db.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		b.log.Error("error is while scanning count", logger.Error(err))
		return models.BranchResponse{}, err
	}

	query = `select id, name, address, phone_number, created_at, updated_at
							from branches where deleted_at = 0 
`
	if search != "" {
		query += fmt.Sprintf(` and name ilike '%s' `, search)
	}

	query += ` order by created_at desc LIMIT $1 OFFSET $2 `
	rows, err := b.db.Query(ctx, query, request.Limit, offset)
	if err != nil {
		b.log.Error("error is while selecting * from branches", logger.Error(err))
		return models.BranchResponse{}, err
	}

	for rows.Next() {
		branch := models.Branch{}
		if err := rows.Scan(
			&branch.ID,
			&branch.Name,
			&branch.Address,
			&branch.PhoneNumber,
			&createdAt,
			&updatedAt); err != nil {
			b.log.Error("error is while scanning branch", logger.Error(err))
			return models.BranchResponse{}, err
		}
		if createdAt.Valid {
			branch.CreatedAt = createdAt.String
		}

		if updatedAt.Valid {
			branch.UpdatedAt = updatedAt.String
		}
		branches = append(branches, branch)
	}

	return models.BranchResponse{
		Branches: branches,
		Count:    count,
	}, err
}
func (b branchRepo) Update(ctx context.Context, branch models.UpdateBranch) (string, error) {
	query := `update branches set name = $1, address = $2, phone_number = $3, updated_at = Now() 
                					where id = $4`

	if rowsAffected, err := b.db.Exec(ctx, query,
		&branch.Name,
		&branch.Address,
		&branch.PhoneNumber,
		&branch.ID); err != nil {
		if r := rowsAffected.RowsAffected(); r == 0 {
			b.log.Error("error is in rows affected", logger.Error(err))
			return "", err
		}
		b.log.Error("error is while updating branch", logger.Error(err))
		return "", err
	}

	return branch.ID, nil
}
func (b branchRepo) Delete(ctx context.Context, key models.PrimaryKey) error {
	query := `update branches set deleted_at = extract(epoch from current_timestamp) where id = $1`

	if rowsAffected, err := b.db.Exec(ctx, query, key.ID); err != nil {
		if r := rowsAffected.RowsAffected(); r == 0 {
			b.log.Error("error is in rows affected", logger.Error(err))
			return err
		}
		b.log.Error("error is while deleting branches", logger.Error(err))
		return err
	}
	return nil
}
