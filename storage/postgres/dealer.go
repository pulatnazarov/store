package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
)

type dealerRepo struct {
	db  *pgxpool.Pool
	log logger.ILogger
}

func NewDealerRepo(db *pgxpool.Pool, log logger.ILogger) storage.IDealerStorage {
	return &dealerRepo{
		db:  db,
		log: log,
	}
}

func (d *dealerRepo) AddSum(ctx context.Context, totalSum int, id string) error {
	//ozini sum: ga qoshish kerak total sum -> update
	query := `update dealer set sum = sum + $1 where id = $2`
	if rowsAffected, err := d.db.Exec(ctx, query, &totalSum, &id); err != nil {
		if r := rowsAffected.RowsAffected(); r == 0 {
			d.log.Error("error is while rows affected", logger.Error(err))
			return err
		}
		d.log.Error("error is while updating dealer sum", logger.Error(err))
		return err
	}
	return nil
}

func (d *dealerRepo) Get(ctx context.Context, key models.PrimaryKey) (models.Dealer, error) {
	dealer := models.Dealer{}
	query := `select id, name, sum from dealer where id = $1`

	if err := d.db.QueryRow(ctx, query, key.ID).Scan(&dealer.ID, &dealer.Name, &dealer.Sum); err != nil {
		d.log.Error("error is while selecting dealer", logger.Error(err))
		return models.Dealer{}, err
	}
	return dealer, nil
}
