package postgres

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"test/api/models"
	"test/storage"
)

type basketRepo struct {
	db *sql.DB
}

func NewBasketRepo(db *sql.DB) storage.IBasket {
	return basketRepo{db: db}
}

func (b basketRepo) CreateBasket(basket models.CreateBasket) (models.Basket, error) {
	bas := models.Basket{}

	id := uuid.New()

	if _, err := b.db.Exec(`insert into baskets(id, customer_id, total_sum)
	values($1, $2, $3)`, id, basket.CustomerID, basket.TotalSum); err != nil {
		fmt.Println("error is while inserting data", err.Error())
		return models.Basket{}, err
	}

	if err := b.db.QueryRow(`select id, customer_id, total_sum from baskets where id = $1`, id).Scan(&bas.ID, &bas.CustomerID, &bas.TotalSum); err != nil {
		fmt.Println("error is while selecting ", err.Error())
		return models.Basket{}, err
	}
	return bas, nil
}

func (b basketRepo) GetBasketByID(key models.PrimaryKey) (models.Basket, error) {
	basket := models.Basket{}

	if err := b.db.QueryRow(`select id, customer_id, total_sum from baskets where id = $1`, key.ID).Scan(&basket.ID,
		&basket.CustomerID, &basket.TotalSum); err != nil {
		fmt.Println("error is while selecting basket", err.Error())
		return models.Basket{}, err
	}

	return basket, nil
}

func (b basketRepo) GetBasketList() (models.BasketResponse, error) {
	baskets := models.BasketResponse{}

	rows, err := b.db.Query(`select id, customer_id, total_sum, count(*) from baskets group by id, customer_id, total_sum`)
	if err != nil {
		fmt.Println("error is while selecting baskets", err.Error())
		return models.BasketResponse{}, err
	}

	for rows.Next() {
		b := models.Basket{}
		if err = rows.Scan(&b.ID, &b.CustomerID, &b.TotalSum, &baskets.Count); err != nil {
			fmt.Println("error is while scanning data", err.Error())
			return models.BasketResponse{}, err
		}
		baskets.Baskets = append(baskets.Baskets, b)

	}

	return baskets, nil
}

func (b basketRepo) UpdateBasket(basket models.UpdateBasket) (models.Basket, error) {
	bas := models.Basket{}

	if _, err := b.db.Exec(`update baskets set customer_id = $1, total_sum = $2 where id = $3`, &basket.CustomerID, &basket.TotalSum, &basket.ID); err != nil {
		return models.Basket{}, err
	}

	if err := b.db.QueryRow(`select id, customer_id, total_sum from baskets where id = $1`, basket.ID).Scan(&bas.ID, &bas.CustomerID, &bas.TotalSum); err != nil {
		fmt.Println("error is while selecting ", err.Error())
		return models.Basket{}, err
	}
	return bas, nil
}

func (b basketRepo) DeleteBasket(key models.PrimaryKey) error {
	if _, err := b.db.Exec(`delete from baskets where id = $1`, key.ID); err != nil {
		return err
	}
	return nil
}
