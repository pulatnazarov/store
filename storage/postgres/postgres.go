package postgres

import (
	"database/sql"
	"fmt"
	"test/config"

	_ "github.com/lib/pq"
)

type Store struct {
	DB            *sql.DB
	CarStorage    carRepo
	DriverStorage driverRepo
}

func New(cfg config.Config) (Store, error) {
	url := fmt.Sprintf(`host=%s port=%s user=%s password=%s database=%s sslmode=disable`, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return Store{}, err
	}

	carRepo := NewCarRepo(db)
	driverRepo := NewDriverRepo(db)

	return Store{
		DB:         db,
		CarStorage: carRepo,
		DriverStorage: driverRepo,
	}, nil
}
