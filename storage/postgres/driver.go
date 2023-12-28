package postgres

import (
	"database/sql"
	"test/models"

	"github.com/google/uuid"
)

type driverRepo struct {
	DB *sql.DB
}

func NewDriverRepo(db *sql.DB) driverRepo {
	return driverRepo{
		DB: db,
	}
}

func (d driverRepo) Insert(driver models.Driver) (string, error) {
	id := uuid.New()

	if _, err := d.DB.Exec(`insert into drivers values ($1, $2, $3)`, id, driver.FullName, driver.Phone); err != nil {
		return "", err
	}

	return id.String(), nil
}