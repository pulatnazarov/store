package postgres

import (
	"database/sql"
	"fmt"
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

func (d driverRepo) GetByID(id string) (models.Driver, error) {
	driver := models.Driver{}


	fmt.Println("id ", id)
	if err := d.DB.QueryRow(`select id, full_name, phone from drivers where id = $1`, id).Scan(
		&driver.ID,
		&driver.FullName,
		&driver.Phone,
	); err != nil {
		fmt.Println("error in storage", err.Error())
		return models.Driver{}, err	
	}

	return driver, nil
}