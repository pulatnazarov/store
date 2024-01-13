package postgres

import (
	"database/sql"
	"fmt"
	"test/models"
	"test/storage"

	"github.com/google/uuid"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) storage.IUserStorage {
	return userRepo{
		db: db,
	}
}

func (u userRepo) Create(createUser models.CreateUser) (models.User, error) {

	uid := uuid.New()

	if _, err := u.db.Exec(`insert into 
			users values ($1, $2, $3, $4, $5, $6)
			`,
		uid,
		createUser.FullName,
		createUser.Phone,
		createUser.Password,
		createUser.UserType,
		createUser.Cash,
	); err != nil {
		fmt.Println("error while inserting data", err.Error())
		return models.User{}, err
	}

	return models.User{}, nil
}

func (u userRepo) GetByID(models.PrimaryKey) (models.User, error) {
	return models.User{}, nil
}

func (u userRepo) GetList(models.GetListRequest) (models.UsersResponse, error) {
	return models.UsersResponse{}, nil
}

func (u userRepo) Update(models.UpdateUser) (models.User, error) {

	return models.User{}, nil
}

func (u userRepo) Delete(models.PrimaryKey) error {

	return nil
}
