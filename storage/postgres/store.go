package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"test/pkg/logger"
	"test/storage"
)

type storeRepo struct {
	db    *pgxpool.Pool
	log   logger.ILogger
	redis storage.IRedisStorage
}

func NewStoreRepo(db *pgxpool.Pool, log logger.ILogger, redis storage.IRedisStorage) storage.IStoreStorage {
	return &storeRepo{
		db:    db,
		log:   log,
		redis: redis,
	}
}

func (s *storeRepo) AddProfit(ctx context.Context, profit float32, branchID string) error {
	rowsAffected, err := s.db.Exec(ctx, `update store set profit = profit + $1, updated_at = now() where branch_id = $2`, profit, branchID)
	if err != nil {
		s.log.Error("Error while adding profit to store", logger.Error(err))
		return err
	}

	if n := rowsAffected.RowsAffected(); n == 0 {
		s.log.Error("Error in rows affected", logger.Error(err))
		return err
	}

	return err
}

func (s *storeRepo) GetStoreBudget(ctx context.Context, branchID string) (float32, error) {
	var budget float32
	query := `select budget from store where branch_id = $1`
	if err := s.db.QueryRow(ctx, query, branchID).Scan(&budget); err != nil {
		s.log.Error("error is while getting store budget", logger.Error(err))
		return 0, err
	}

	return budget, nil
}

func (s *storeRepo) WithdrawalDeliveredSum(ctx context.Context, totalSum float32, branchID string) error {
	query := `update store set budget = budget - $1 where branch_id = $2 `
	if rowsAffected, err := s.db.Exec(ctx, query, &totalSum, &branchID); err != nil {
		if r := rowsAffected.RowsAffected(); r == 0 {
			s.log.Error("error is while rows affected", logger.Error(err))
			return err
		}
		s.log.Error("error is while updating budget", logger.Error(err))
		return err
	}

	return nil
}
