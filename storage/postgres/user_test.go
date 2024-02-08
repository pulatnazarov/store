package postgres

import (
	"context"
	"github.com/go-playground/assert/v2"
	"test/api/models"
	"test/config"
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
		Phone:    "+998112221155",
		Password: "password",
		Cash:     10,
		UserType: "customer",
		BranchID: "aa541fcc-bf74-11ee-ae0b-166244b65504",
	}

	userID, err := pgStore.User().Create(context.Background(), createUser)
	if err != nil {
		t.Errorf("error while creating user error: %v", err)
	}

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

	if len(usersResp.Users) != 16 {
		t.Errorf("expected 16, but got: %d", len(usersResp.Users))
	}

	assert.Equal(t, len(usersResp.Users), 16)

}
