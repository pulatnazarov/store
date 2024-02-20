package postgres

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/net/context"
	"test/api/models"
	"test/pkg/helper"
	"test/pkg/logger"
	"test/storage"
)

type incomeRepo struct {
	db    *pgxpool.Pool
	log   logger.ILogger
	redis storage.IRedisStorage
}

func NewIncomeRepo(db *pgxpool.Pool, log logger.ILogger, redis storage.IRedisStorage) storage.IIncomeStorage {
	return &incomeRepo{
		db:    db,
		log:   log,
		redis: redis,
	}
}

func (i *incomeRepo) Create(ctx context.Context) (models.Income, error) {
	var (
		income = models.Income{}
		extID  string
	)

	query := `select external_id from incomes order by external_id desc`

	if err := i.db.QueryRow(ctx, query).Scan(
		&extID,
	); err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			i.log.Error("error while getting ext id ", logger.Error(err))
			return models.Income{}, err
		}
		extID = "I"
	}

	if extID != "I" {
		extID = helper.GenerateExternalID(extID)
	} else {
		extID = "I-0001"
	}

	fmt.Println("ex", extID)

	query = `insert into incomes values ($1, $2, $3) returning id, external_id`

	fmt.Println("ext id ", extID)
	if err := i.db.QueryRow(ctx, query, uuid.New(), extID, 0).Scan(
		&income.ID,
		&income.ExternalID,
	); err != nil {
		i.log.Error("error while creating income ", logger.Error(err))
		return models.Income{}, err
	}

	return income, nil
}

func (i *incomeRepo) GetByID(ctx context.Context, key models.PrimaryKey) (models.Income, error) {
	income := models.Income{}
	query := `select id, external_id, total_sum from incomes where id = $1 and deleted_at = 0`
	if err := i.db.QueryRow(ctx, query, key.ID).Scan(
		&income.ID,
		&income.ExternalID,
		&income.TotalSum,
	); err != nil {
		i.log.Error("error is while selecting income by id", logger.Error(err))
		return models.Income{}, err
	}
	return income, nil
}

func (i *incomeRepo) GetList(ctx context.Context, request models.GetListRequest) (models.IncomesResponse, error) {
	var (
		page              = request.Page
		offset            = (page - 1) * request.Limit
		incomes           = []models.Income{}
		query, countQuery string
		count             int
		search            = request.Search
	)
	countQuery = `select count(1) from incomes where deleted_at = 0`
	if search != "" {
		countQuery += fmt.Sprintf(` and external_id = '%s'`, search)
	}

	if err := i.db.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		i.log.Error("error is while scanning count", logger.Error(err))
		return models.IncomesResponse{}, err
	}

	query = `select id, external_id, total_sum from incomes where deleted_at = 0`
	if search != "" {
		query += fmt.Sprintf(` and external_id = '%s'`, search)
	}
	query += ` LIMIT $1 OFFSET $2`
	rows, err := i.db.Query(ctx, query, request.Limit, offset)
	if err != nil {
		i.log.Error("error is while selecting all", logger.Error(err))
		return models.IncomesResponse{}, err
	}

	for rows.Next() {
		in := models.Income{}
		if err = rows.Scan(
			&in.ID,
			&in.ExternalID,
			&in.TotalSum,
		); err != nil {
			i.log.Error("error is while scanning all", logger.Error(err))
			return models.IncomesResponse{}, err
		}
		incomes = append(incomes, in)
	}

	return models.IncomesResponse{
		Incomes: incomes,
		Count:   count,
	}, err
}

func (i *incomeRepo) Delete(ctx context.Context, key models.PrimaryKey) error {
	query := `update incomes set deleted_at = extract(epoch from current_timestamp) where id = $1`
	if rowsAffected, err := i.db.Exec(ctx, query, key.ID); err != nil {
		if r := rowsAffected.RowsAffected(); r == 0 {
			i.log.Error("error is while rows affected", logger.Error(err))
			return err
		}
		i.log.Error("error is while delete income", logger.Error(err))
		return err
	}
	return nil
}
