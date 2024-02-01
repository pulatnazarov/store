package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"test/api/models"
	"test/storage"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) storage.IUserStorage {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) Create(ctx context.Context, createUser models.CreateUser) (string, error) {

	uid := uuid.New()

	if _, err := u.db.Exec(ctx, `insert into 
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
		return "", err
	}

	return uid.String(), nil
}

func (u *userRepo) GetByID(ctx context.Context, pKey models.PrimaryKey) (models.User, error) {
	user := models.User{}

	query := `
		select id, full_name, phone, cash from users where id = $1 and user_role = 'customer'
`
	if err := u.db.QueryRow(ctx, query, pKey.ID).Scan(
		&user.ID,
		&user.FullName,
		&user.Phone,
		&user.Cash,
	); err != nil {
		fmt.Println("error while scanning user", err.Error())
		return models.User{}, err
	}

	return user, nil
}

func (u *userRepo) GetList(ctx context.Context, request models.GetListRequest) (models.UsersResponse, error) {
	var (
		users             = []models.User{}
		count             = 0
		countQuery, query string
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
	)

	countQuery = `
		SELECT count(1) from users where user_role = 'customer' `

	if search != "" {
		countQuery += fmt.Sprintf(` and (phone ilike '%%%s%%' or full_name ilike '%%%s%%')`, search, search)
	}

	if err := u.db.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		fmt.Println("error while scanning count of users", err.Error())
		return models.UsersResponse{}, err
	}

	query = `
		SELECT id, full_name, phone, cash
			FROM users
			    WHERE user_role = 'customer'
			    `

	if search != "" {
		query += fmt.Sprintf(` and (phone ilike '%%%s%%' or full_name ilike '%%%s%%') `, search, search)
	}

	query += ` LIMIT $1 OFFSET $2`

	rows, err := u.db.Query(ctx, query, request.Limit, offset)
	if err != nil {
		fmt.Println("error while query rows", err.Error())
		return models.UsersResponse{}, err
	}

	for rows.Next() {
		user := models.User{}

		if err = rows.Scan(
			&user.ID,
			&user.FullName,
			&user.Phone,
			&user.Cash,
		); err != nil {
			fmt.Println("error while scanning row", err.Error())
			return models.UsersResponse{}, err
		}

		users = append(users, user)
	}

	return models.UsersResponse{
		Users: users,
		Count: count,
	}, nil
}

func (u *userRepo) Update(ctx context.Context, request models.UpdateUser) (string, error) {
	query := `
		update users 
			set full_name = $1, phone = $2, cash = $3
				where user_role = 'customer' and id = $4`

	if _, err := u.db.Exec(ctx, query, request.FullName, request.Phone, request.Cash, request.ID); err != nil {
		fmt.Println("error while updating user data", err.Error())
		return "", err
	}

	return request.ID, nil
}

func (u *userRepo) Delete(ctx context.Context, request models.PrimaryKey) error {
	query := `
		delete from users
			where id = $1
`
	if _, err := u.db.Exec(ctx, query, request.ID); err != nil {
		fmt.Println("error while deleting user by id", err.Error())
		return err
	}

	return nil
}

func (u *userRepo) GetPassword(ctx context.Context, id string) (string, error) {
	password := ""

	query := `
		select password from users 
		                where user_role = 'customer' and id = $1`

	if err := u.db.QueryRow(ctx, query, id).Scan(&password); err != nil {
		fmt.Println("Error while scanning password from users", err.Error())
		return "", err
	}

	return password, nil
}

func (u *userRepo) UpdatePassword(ctx context.Context, request models.UpdateUserPassword) error {
	query := `
		update users 
				set password = $1
					where id = $2 and user_role = 'customer'`

	if _, err := u.db.Exec(ctx, query, request.NewPassword, request.ID); err != nil {
		fmt.Println("error while updating password for user", err.Error())
		return err
	}

	return nil
}

func (u *userRepo) UpdateCustomerCash(ctx context.Context, id string, sum int) error {
	query := `update customer set cash = cash - $1 where id = $2`

	if _, err := u.db.Exec(ctx, query, sum, id); err != nil {
		fmt.Println("error while updating customer cash", err.Error())
		return err
	}

	return nil
}
