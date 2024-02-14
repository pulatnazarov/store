package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"test/api/models"
	"test/pkg/check"
	"test/pkg/logger"
	"test/pkg/security"
	"test/storage"
)

type userService struct {
	storage storage.IStorage
	log     logger.ILogger
}

func NewUserService(storage storage.IStorage, log logger.ILogger) userService {
	return userService{
		storage: storage,
		log:     log,
	}
}

func (u userService) Create(ctx context.Context, createUser models.CreateUser) (models.User, error) {
	u.log.Info("User create service layer", logger.Any("createUser", createUser))

	password, err := security.HashPassword(createUser.Password)
	if err != nil {
		u.log.Error("error while hashing password", logger.Error(err))
		return models.User{}, err
	}
	createUser.Password = password

	pKey, err := u.storage.User().Create(ctx, createUser)
	if err != nil {
		u.log.Error("error while creating user", logger.Error(err))
		return models.User{}, err
	}

	user, err := u.storage.User().GetByID(ctx, models.PrimaryKey{
		ID: pKey,
	})

	return user, nil
}

func (u userService) GetUser(ctx context.Context, pKey models.PrimaryKey) (models.User, error) {
	user, err := u.storage.User().GetByID(ctx, pKey)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Println("ERROR in service layer while getting user by id", err.Error())
			return models.User{}, err
		}
	}

	return user, nil
}

func (u userService) GetUsers(ctx context.Context, request models.GetListRequest) (models.UsersResponse, error) {
	u.log.Info("Get user list service layer", logger.Any("request", request))
	usersResponse, err := u.storage.User().GetList(ctx, request)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			u.log.Error("error while getting users list", logger.Error(err))
			return models.UsersResponse{}, err
		}
	}

	return usersResponse, err
}

func (u userService) Update(ctx context.Context, updateUser models.UpdateUser) (models.User, error) {
	pKey, err := u.storage.User().Update(ctx, updateUser)
	if err != nil {
		u.log.Error("ERROR in service layer while updating updateUser", logger.Error(err))
		return models.User{}, err
	}

	user, err := u.storage.User().GetByID(ctx, models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		u.log.Error("ERROR in service layer while getting user after update", logger.Error(err))
		return models.User{}, err
	}

	return user, nil
}

func (u userService) Delete(ctx context.Context, key models.PrimaryKey) error {
	err := u.storage.User().Delete(ctx, key)
	return err
}

func (u userService) UpdatePassword(ctx context.Context, request models.UpdateUserPassword) error {
	oldPassword, err := u.storage.User().GetPassword(ctx, request.ID)
	if err != nil {
		u.log.Error("ERROR in service layer while getting user password", logger.Error(err))
		return err
	}

	if oldPassword != request.OldPassword {
		u.log.Error("ERROR in service old password is not correct")
		return errors.New("old password did not match")
	}

	if err = check.ValidatePassword(request.NewPassword); err != nil {
		u.log.Error("ERROR in service layer new password is weak", logger.Error(err))
		return err
	}

	if err = u.storage.User().UpdatePassword(context.Background(), request); err != nil {
		u.log.Error("ERROR in service layer while updating password", logger.Error(err))
		return err
	}

	return nil
}
