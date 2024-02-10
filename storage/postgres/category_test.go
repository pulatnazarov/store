package postgres

import (
	"context"
	"github.com/go-playground/assert/v2"
	"test/api/models"
	"test/config"
	"testing"
)

func TestCategoryRepo_Create(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}

	category := models.CreateCategory{Name: "vegetable"}

	categoryID, err := pgStore.Category().Create(context.Background(), category)
	if err != nil {
		t.Errorf("error while creating category error: %v", err)
	}

	createdCategory, err := pgStore.Category().GetByID(context.Background(), models.PrimaryKey{ID: categoryID})
	if err != nil {
		t.Errorf("error while getting category error: %v", err)
	}

	assert.Equal(t, createdCategory.Name, category.Name)
}

func TestCategoryRepo_GetByID(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}

	category := models.CreateCategory{Name: "vegetable"}

	categoryID, err := pgStore.Category().Create(context.Background(), category)
	if err != nil {
		t.Errorf("error while creating category error: %v", err)
	}

	//categoryID := "f3b401e7-8414-4820-809c-46c6c88705aa"

	getCategory, err := pgStore.Category().GetByID(context.Background(), models.PrimaryKey{ID: categoryID})
	if err != nil {
		t.Errorf("error while getting category by id error: %v", err)
	}

	assert.Equal(t, getCategory.Name, category.Name)
}

func TestCategoryRepo_GetList(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}

	resp, err := pgStore.Category().GetList(context.Background(), models.GetListRequest{
		Page:  1,
		Limit: 1000,
	})
	if err != nil {
		t.Errorf("error while getting category list error: %v", err)
	}

	assert.Equal(t, len(resp.Category), 16)
}

func TestCategoryRepo_Update(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}

	category := models.UpdateCategory{
		ID:   "2d9bbe1f-1119-432b-98be-05042e80ac34",
		Name: "gadget",
	}

	categoryID, err := pgStore.Category().Update(context.Background(), category)
	if err != nil {
		t.Errorf("error while updating category  error: %v", err)
	}

	updatedCategory, err := pgStore.Category().GetByID(context.Background(), models.PrimaryKey{ID: categoryID})
	if err != nil {
		t.Errorf("error while getting category by categoryID error: %v", err)
	}

	assert.Equal(t, updatedCategory.Name, category.Name)
}

func TestCategoryRepo_Delete(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}

	categoryID := "7e6a876e-ad7d-41f4-b6a3-fef6af51e66d"

	err = pgStore.Category().Delete(context.Background(), models.PrimaryKey{ID: categoryID})
	assert.Equal(t, err, nil)
}
