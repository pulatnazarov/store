package postgres

import (
	"context"
	"fmt"
	"github.com/go-playground/assert/v2"
	"test/api/models"
	"test/config"
	"test/pkg/helper"
	"testing"
)

func TestUserRepo_Create(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}

	createUser := models.CreateUser{
		FullName: "First Person",
		Phone:    helper.PhoneGenerate(),
		Password: "password",
		Cash:     10,
		UserType: "customer",
		BranchID: "aa541fcc-bf74-11ee-ae0b-166244b65504",
	}
	userID, err := pgStore.User().Create(context.Background(), createUser)
	if err != nil {
		t.Errorf("error while creating user error: %v", err)
	}

	fmt.Println("phone", createUser.Phone)
	user, err := pgStore.User().GetByID(context.Background(), models.PrimaryKey{
		ID: userID,
	})
	if err != nil {
		t.Errorf("error while getting user error: %v", err)
	}

	assert.Equal(t, user.FullName, createUser.FullName)
	assert.Equal(t, user.Phone, createUser.Phone)
	assert.Equal(t, user.Cash, createUser.Cash)
}

func TestUserRepo_GetByID(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}

	createUser := models.CreateUser{
		FullName: "First Person",
		Phone:    helper.PhoneGenerate(),
		Password: "password",
		Cash:     10,
		UserType: "customer",
		BranchID: "aa541fcc-bf74-11ee-ae0b-166244b65504",
	}
	userID, err := pgStore.User().Create(context.Background(), createUser)
	if err != nil {
		t.Errorf("error while creating user error: %v", err)
	}

	userData, err := pgStore.User().GetByID(context.Background(), models.PrimaryKey{ID: userID})
	if err != nil {
		t.Errorf("error while getting user error: %v", err)
	}

	assert.Equal(t, userData.FullName, createUser.FullName)
	assert.Equal(t, userData.Phone, createUser.Phone)
	assert.Equal(t, userData.Cash, createUser.Cash)
}

func TestUserRepo_GetList(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}

	usersResp, err := pgStore.User().GetList(context.Background(), models.GetListRequest{
		Page:  1,
		Limit: 1000,
	})
	if err != nil {
		t.Errorf("error while getting usersResp error: %v", err)
	}

	assert.Equal(t, len(usersResp.Users), 46)
}

func TestUserRepo_Update(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}

	user := models.UpdateUser{
		ID:       "74dd7706-6499-44e8-a803-a83f1fb30b94",
		FullName: "Tohirova Farangiz",
		Phone:    "+998669279543",
		Cash:     9000,
	}

	userID, err := pgStore.User().Update(context.Background(), user)
	if err != nil {
		t.Errorf("error while updating user error: %v", err)
	}

	updatedUser, err := pgStore.User().GetByID(context.Background(), models.PrimaryKey{ID: userID})
	if err != nil {
		t.Errorf("error while getting user by id error: %v", err)
	}

	assert.Equal(t, updatedUser.ID, user.ID)
	assert.Equal(t, updatedUser.FullName, user.FullName)
	assert.Equal(t, updatedUser.Phone, user.Phone)
	assert.Equal(t, updatedUser.Cash, user.Cash)
}

func TestUserRepo_Delete(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}

	id := "57504e09-f6c3-49e0-8fd5-58ba687a6eed"

	err = pgStore.User().Delete(context.Background(), models.PrimaryKey{ID: id})
	assert.Equal(t, err, nil)
}

func TestUserRepo_GetPassword(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}

	id := "74dd7706-6499-44e8-a803-a83f1fb30b94"
	password := "password098"

	getPassword, err := pgStore.User().GetPassword(context.Background(), id)
	if err != nil {
		t.Errorf("error while get password error: %v", err)
	}

	assert.Equal(t, getPassword, password)
}

//func TestUserRepo_UpdatePassword(t *testing.T) {
//	cfg := config.Load()
//
//	pgStore, err := New(context.Background(), cfg)
//	if err != nil {
//		t.Errorf("error while connection to db error: %v", err)
//	}
//
//	request := models.UpdateUserPassword{
//		ID:          "620c3532-ba82-44fd-8882-b364ff7b96ba",
//		NewPassword: "password112",
//		OldPassword: "password111",
//	}
//
//	err = pgStore.User().UpdatePassword(context.Background(), request)
//	if err != nil {
//		t.Errorf("error while updating password error: %v", err)
//	}
//	response, err := pgStore.User().GetByID(context.Background(), models.PrimaryKey{ID: request.ID})
//	if err != nil {
//		t.Errorf("error while getting user by id error: %v", err)
//	}
//
//	assert.Equal(t, response.Password, request.NewPassword)
//}

//func TestUserRepo_UpdateCustomerCash(t *testing.T) {
//	cfg := config.Load()
//
//	pgStore, err := New(context.Background(), cfg)
//	if err != nil {
//		t.Errorf("error while connection to db error: %v", err)
//	}
//
//	id := "437e712f-fe9b-4cc7-8a11-737391dcfe7f"
//	sum := 2
//	cash := 2
//
//	err = pgStore.User().UpdateCustomerCash(context.Background(), id, sum)
//	if err != nil {
//		t.Errorf("error while updating customer cash error: %v", err)
//	}
//	response, err := pgStore.User().GetByID(context.Background(), models.PrimaryKey{ID: id})
//	if err != nil {
//		t.Errorf("error while getting user by id error: %v", err)
//	}
//
//	assert.Equal(t, response.Cash, cash)
//}
