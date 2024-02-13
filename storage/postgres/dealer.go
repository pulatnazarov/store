package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"test/api/models"
	"test/storage"
)

type dealerRepo struct {
	db *pgxpool.Pool
}

func NewDealerRepo(db *pgxpool.Pool) storage.IDealerStorage {
	return &dealerRepo{
		db: db,
	}
}

func (d *dealerRepo) AddSum(ctx context.Context, totalSum int, id string) error {
	//ozini sum: ga qoshish kerak total sum -> update
	query := `update dealer set sum = sum + $1 where id = $2`
	if rowsAffected, err := d.db.Exec(ctx, query, &totalSum, &id); err != nil {
		if r := rowsAffected.RowsAffected(); r == 0 {
			fmt.Println("error is while rows affected", err.Error())
			return err
		}
		fmt.Println("error is while updating dealer sum", err.Error())
		return err
	}
	return nil
}

func (d *dealerRepo) Get(ctx context.Context, key models.PrimaryKey) (models.Dealer, error) {
	dealer := models.Dealer{}
	query := `select id, name, sum from dealer where id = $1`

	if err := d.db.QueryRow(ctx, query, key.ID).Scan(&dealer.ID, &dealer.Name, &dealer.Sum); err != nil {
		fmt.Println("error is while selecting dealer", err)
		return models.Dealer{}, err
	}
	return dealer, nil
}
