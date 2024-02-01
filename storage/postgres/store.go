package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"test/storage"
)

type storeRepo struct {
	db *pgxpool.Pool
}

func NewStoreRepo(db *pgxpool.Pool) storage.IStoreStorage {
	return &storeRepo{
		db: db,
	}
}

func (s *storeRepo) AddProfit(ctx context.Context, profit float32, branchID string) error {
	rowsAffected, err := s.db.Exec(ctx, `update store set profit = profit + $1, updated_at = now() where id = $2`, profit, branchID)
	if err != nil {
		fmt.Println("Error while adding profit to store", err.Error())
		return err
	}

	if n := rowsAffected.RowsAffected(); n == 0 {
		fmt.Println("Error in rows affected", err.Error())
		return err
	}

	return err
}
