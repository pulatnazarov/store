package postgres

import (
	"database/sql"
	"fmt"
	"test/api/models"
	"test/storage"
)

type storeRepo struct {
	db   *sql.DB
	repo repo
}

func NewStoreRepo(db *sql.DB) storage.IStore {
	return &storeRepo{db: db}
}

func (s *storeRepo) Sell(product models.ProductSell, user models.UserSell) (models.Ticket, error) {
	if user.Cash < product.Price*product.Quantity {
		fmt.Println("You don't have enough cash")
		return models.Ticket{}, nil
	}

	profit := product.Quantity * (product.Price - product.OriginalPrice)
	query := `update store set profit = $1`
	if _, err := s.db.Exec(query, profit); err != nil {
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
