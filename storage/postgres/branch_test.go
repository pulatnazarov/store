package postgres

import (
	"context"
	"test/api/models"
	"test/config"
	"test/pkg/helper"
	"test/pkg/logger"
	"test/storage/redis"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestBranchRepo_Create(t *testing.T) {
	cfg := config.Load()

	log := logger.New(cfg.ServiceName)
	r:=redis.New(cfg)

	pgStore, err := New(context.Background(), cfg,log,r)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}

	createBranch := models.CreateBranch{
		Name:        "Korzinka",
		Address:     "Beruniy",
		PhoneNumber: helper.GeneratePhoneNumber(),
	}
	bramchID, err := pgStore.Branch().Create(context.Background(), createBranch)
	if err != nil {
		t.Errorf("error while creating branch err : %v", err)
	}

	branch, err := pgStore.Branch().GetByID(context.Background(), models.PrimaryKey{ID: bramchID})
	if err != nil {
		t.Errorf("error while gitting branch err : %v", err)
	}

	assert.Equal(t, branch.Name, createBranch.Name)
	assert.Equal(t, branch.Address, createBranch.Address)
	assert.Equal(t, branch.PhoneNumber, createBranch.PhoneNumber)
}

func TestBranchRepo_GetByID(t *testing.T) {
	cfg := config.Load()

	log := logger.New(cfg.ServiceName)
	r:=redis.New(cfg)

	pgStore, err := New(context.Background(), cfg, log,r)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}

	createBranch := models.CreateBranch{
		Name:        "Korzinka",
		Address:     "Beruniy",
		PhoneNumber: helper.GeneratePhoneNumber(),
	}
	branchID, err := pgStore.Branch().Create(context.Background(), createBranch)
	if err != nil {
		t.Errorf("error while creating branch err : %v", err)
	}

	branch, err := pgStore.Branch().GetByID(context.Background(), models.PrimaryKey{ID: branchID})
	if err != nil {
		t.Errorf("error while getting by id err : %v", err)
	}

	if branch.ID != branchID {
		t.Errorf("expected: %q, but got %q", branchID, branch.ID)
	}

	if branch.Name == "" {
		t.Error("expected some full name, but got nothing")
	}

	if branch.PhoneNumber == "" {
		t.Error("expected some full name, but got nothing")
	} else if len(branch.PhoneNumber) >= 14 || len(branch.PhoneNumber) <= 12 {
		t.Errorf("expected phone length: 13, but got %d, user id is %s", len(branch.PhoneNumber), branch.ID)
	}

	if branch.Address == "" {
		t.Error("expected some branch id, but got nothing")
	}

}

func TestBranchRepo_GetList(t *testing.T) {
	cfg := config.Load()

	log := logger.New(cfg.ServiceName)
	r:=redis.New(cfg)

	pgStore, err := New(context.Background(), cfg, log,r)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}

	categoryResp, err := pgStore.Branch().GetList(context.Background(), models.GetListRequest{
		Page:  1,
		Limit: 1000,
	})
	if err != nil {
		t.Errorf("error while getting categoryResp error: %v", err)
	}
	if len(categoryResp.Branches) != 9 {
		t.Errorf("expected 9, but got: %d", len(categoryResp.Branches))
	}

	assert.Equal(t, len(categoryResp.Branches), 9)
}

func TestBranchRepo_Update(t *testing.T) {
	cfg := config.Load()

	log := logger.New(cfg.ServiceName)
	r:=redis.New(cfg)
	pgStore, err := New(context.Background(), cfg, log,r)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}

	createBranch := models.CreateBranch{
		Name:        "korzinka",
		Address:     "Ipakachi",
		PhoneNumber: "+998781401414",
	}

	branchID, err := pgStore.Branch().Create(context.Background(), createBranch)
	if err != nil {
		t.Errorf("error while creating branch error: %v", err)
	}

	updateBranch := models.UpdateBranch{
		ID:          branchID,
		Name:        "Ipakachi",
		PhoneNumber: "+998781401414",
	}

	branchUpdateID, err := pgStore.Branch().Update(context.Background(), updateBranch)
	if err != nil {
		t.Errorf("error while creating branch error: %v", err)
	}

	branch, err := pgStore.Branch().GetByID(context.Background(), models.PrimaryKey{
		ID: branchUpdateID,
	})
	if err != nil {
		t.Errorf("error while getting branch error: %v", err)
	}

	assert.Equal(t, branchID, branch.ID)
	assert.Equal(t, branch.Name, updateBranch.Name)
	assert.Equal(t, branch.PhoneNumber, updateBranch.PhoneNumber)
	assert.Equal(t, branch.Address, updateBranch.Address)
}

func TestBranchRepo_Delete(t *testing.T) {
	cfg := config.Load()

	log := logger.New(cfg.ServiceName)
	r:=redis.New(cfg)
	pgStore, err := New(context.Background(), cfg, log,r)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}

	createBranch := models.CreateBranch{
		Name:        "korzinka",
		Address:     "Qoratosh",
		PhoneNumber: "+998787777777",
	}

	branchID, err := pgStore.Branch().Create(context.Background(), createBranch)
	if err != nil {
		t.Errorf("error while creating banch error: %v", err)
	}

	if err = pgStore.Branch().Delete(context.Background(), models.PrimaryKey{ID: branchID}); err != nil {
		t.Errorf("Error deleting barnch: %v", err)
	}
}
