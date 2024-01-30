package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"test/api/models"
	"test/storage"
)

type storeRepo struct {
	db *pgxpool.Pool
}

func NewStoreRepo(db *pgxpool.Pool) storage.IStore {
	return &storeRepo{db: db}
}

func (s *storeRepo) Sell(product models.ProductSell, user models.UserSell) (models.Ticket, error) {
	if user.Cash < product.Price*product.Quantity {
		fmt.Println("You don't have enough cash")
		return models.Ticket{}, nil
	}

	profit := product.Quantity * (product.Price - product.OriginalPrice)
	query := `update store set profit = $1`
	if _, err := s.db.Exec(context.Background(), query, profit); err != nil {
		fmt.Println("error is while updating profit", err.Error())
		return models.Ticket{}, err
	}

	totalPrice := product.Quantity * product.Price
	return models.Ticket{
		Name:         product.Name,
		TotalPrice:   totalPrice,
		SoldQuantity: product.Quantity,
	}, nil
}
