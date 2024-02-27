package postgres

import (
	"context"
	"fmt"
	"strings"
	"test/api/models"
	"test/config"
	"test/pkg/helper"
	"test/pkg/logger"
	"test/storage/redis"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestProductRepo_Create(t *testing.T) {
	cfg := config.Load()
	log := logger.New(cfg.ServiceName)
	newRedis := redis.New(cfg)
	pgStore, err := New(context.Background(), cfg, log, newRedis)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}

	createproducts := models.CreateProduct{
		Name:          "product 1",
		Price:         12,
		OriginalPrice: 10,
		Quantity:      2,
		CategoryID:    "123e4567-e89b-12d3-a456-426614174001",
		BranchID:      "aa541fcc-bf74-11ee-ae0b-166244b65504",
	}
	id, err := pgStore.Product().Create(context.Background(), createproducts)
	if err != nil {
		t.Error("error while inserting product", err)

	}
	idproduct, err := pgStore.Product().GetByID(context.Background(), id)
	if err != nil {
		t.Error("error", err)
	}
	if id == "" {
		t.Error("error while creating product")
	}
	assert.Equal(t, idproduct.Name, createproducts.Name)
	assert.Equal(t, idproduct.Quantity, createproducts.Quantity)
	assert.Equal(t, idproduct.Price, createproducts.Price)
	assert.Equal(t, idproduct.OriginalPrice, createproducts.OriginalPrice)
	assert.Equal(t, idproduct.CategoryID, createproducts.CategoryID)
	assert.Equal(t, idproduct.BranchID, createproducts.BranchID)

}

func TestProductRepo_GetByID(t *testing.T) {
	cfg := config.Load()
	log := logger.New(cfg.ServiceName)
	newRedis := redis.New(cfg)
	pgStore, err := New(context.Background(), cfg, log, newRedis)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}
	products, err := pgStore.Product().GetList(context.Background(), models.GetListRequest{
		Page:   1,
		Limit:  1,
		Search: "",
	})

	if len(products.Products) == 0 {
		t.Error("error", err)

	}

	expectedproducts := products.Products[0].ID
	t.Run("succes", func(t *testing.T) {
		product, err := pgStore.Product().GetByID(context.Background(), expectedproducts)

		if err != nil {
			t.Error("error while geting by id product", err)
		}
		if product.ID != expectedproducts {
			t.Errorf("expected: %q but got: %q", expectedproducts, product.ID)
		}
		if product.Name == "" {
			t.Error("expected: productname but got : nothing")
		}
		if product.OriginalPrice <= 0 {
			t.Errorf("expected: positive original price but got: %q", product.OriginalPrice)

		}
		if strings.TrimSpace(product.BranchID) == "" {
			t.Errorf("expected: non-empty branch id but got: %q", product.BranchID)
		}
		if strings.TrimSpace(product.CategoryID) == "" {
			t.Errorf("expected: non-empty category id  but got: %q", product.CategoryID)
		}
		if product.Quantity <= 0 {
			t.Errorf("exepcted: more than 0 ,but got: %q", product.Quantity)
		}
		if product.Price <= 0 {

			t.Errorf("expeceted: more than 0 price but got %q", product.Price)
		}

		if product.OriginalPrice <= 0 {
			t.Errorf("expeceted: more than 0 price but got %q", product.OriginalPrice)

		}

	})

	t.Run("fail", func(t *testing.T) {
		productid := ""
		product, err := pgStore.Product().GetByID(context.Background(), expectedproducts)
		if err != nil {
			t.Error("error while getting product id", err)
		}
		if product.ID != productid {
			t.Errorf("expected: %q, but got %q", productid, product.ID)
		}
		if product.Name == "" {
			t.Error("expected: productname but got : nothing")
		}
		if product.OriginalPrice <= 0 {
			t.Errorf("expected: positive original price but got: %q", product.OriginalPrice)

		}
		if strings.TrimSpace(product.BranchID) == "" {
			t.Errorf("expected: non-empty branch id but got: %q", product.BranchID)
		}
		if strings.TrimSpace(product.CategoryID) == "" {
			t.Errorf("expected: non-empty category id  but got: %q", product.CategoryID)
		}
		if product.Quantity <= 0 {
			t.Errorf("exepcted: more than 0 ,but got: %q", product.Quantity)
		}
		if product.Price <= 0 {

			t.Errorf("expeceted: more than 0 price but got %q", product.Price)
		}

		if product.OriginalPrice <= 0 {
			t.Errorf("expeceted: more than 0 price but got %q", product.OriginalPrice)

		}

	})

}

func TestProductRepo_GetList(t *testing.T) {
	cfg := config.Load()
	log := logger.New(cfg.ServiceName)
	newRedis := redis.New(cfg)
	pgStore, err := New(context.Background(), cfg, log, newRedis)
	if err != nil {
		t.Errorf("error while connecting db %q", err)
	}

	products, err := pgStore.Product().GetList(context.Background(), models.GetListRequest{
		Page:  1,
		Limit: 1000,
	})
	if err != nil {
		t.Error("error while getting list of products", err.Error())
	}
	if len(products.Products) != 5 {
		t.Errorf("expected 5 rows , but got %q", len(products.Products))
	}

	assert.Equal(t, len(products.Products), 5)

}

func TestProductRepo_Update(t *testing.T) {
	cfg := config.Load()
	log := logger.New(cfg.ServiceName)
	newRedis := redis.New(cfg)
	pgStore, err := New(context.Background(), cfg, log, newRedis)
	if err != nil {
		t.Error("error while connecting to db ", err)
	}

	createProduct := models.CreateProduct{
		Name:          helper.GenerateProductName(),
		Price:         int(helper.GenerateRandomPrice(10.0, 100.0)),
		OriginalPrice: int(helper.GenerateRandomPrice(2.0, 9.0)),
		Quantity:      10,
		CategoryID:    "123e4567-e89b-12d3-a456-426614174001",
	}

	productID, err := pgStore.Product().Create(context.Background(), createProduct)
	if err != nil {
		t.Error("erro while creating product in tetsing", err)
	}

	if err != nil {
		t.Errorf("error while creating product %v", err)
	}

	updateProduct := models.UpdateProduct{
		ID:            productID,
		Name:          helper.GenerateProductName(),
		Price:         int(helper.GenerateRandomPrice(10.0, 100.0)),
		OriginalPrice: int(helper.GenerateRandomPrice(2.0, 9.0)),
		Quantity:      10,
		CategoryID:    "123e4567-e89b-12d3-a456-426614174001",
	}

	productupdateid, err := pgStore.Product().Update(context.Background(), updateProduct)
	if err != nil {
		t.Error("error updatinf product in testing", err)
	}

	product, err := pgStore.Product().GetByID(context.Background(), productID)
	if err != nil {
		t.Error("error while geting by id in testing product", err)
	}
	if productupdateid == "" {
		t.Error("expected updated product id: but got empty string")
	}

	assert.Equal(t, product.ID, productupdateid)
	assert.Equal(t, product.Name, updateProduct.Name)
	assert.Equal(t, product.Price, updateProduct.Price)
	assert.Equal(t, product.OriginalPrice, updateProduct.OriginalPrice)
	assert.Equal(t, product.CategoryID, updateProduct.CategoryID)
	assert.Equal(t, product.Quantity, updateProduct.Quantity)

}

func TestProductRepo_Delete(t *testing.T) {
	cfg := config.Load()
	log := logger.New(cfg.ServiceName)
	newRedis := redis.New(cfg)
	pgStore, err := New(context.Background(), cfg, log, newRedis)
	if err != nil {
		t.Error("error while connecting to db ", err)
	}

	createproduct := models.CreateProduct{
		Name:          helper.GenerateProductName(),
		Price:         int(helper.GenerateRandomPrice(10.0, 100.0)),
		OriginalPrice: int(helper.GenerateRandomPrice(2.0, 9.0)),
		Quantity:      10,
		CategoryID:    "123e4567-e89b-12d3-a456-426614174001",
	}

	productid, err := pgStore.Product().Create(context.Background(), createproduct)
	if err != nil {
		t.Error("error while creating product in delete tetsing", err)
	}

	if err = pgStore.Product().Delete(context.Background(), models.PrimaryKey{ID: productid}); err != nil {
		t.Error("error while deleteing product in testing", err)
	}
	t.Run("falied", func(t *testing.T) {
		if productid == "" {
			t.Error("expected product id but go nothing", err)
		}
	})

}

func TestProductRepo_Search(t *testing.T) {

}
func TestProductRepo_TakeProducts(t *testing.T) {

	cfg := config.Load()
	log := logger.New(cfg.ServiceName)
	newRedis := redis.New(cfg)
	pgStore, err := New(context.Background(), cfg, log, newRedis)
	if err != nil {
		t.Fatal("error while connecting to db ", err)
	}

	//to cover nil
	err = pgStore.Product().TakeProducts(context.Background(), nil)
	if err == nil {
		fmt.Println("succes")
	}

	products := map[string]int{
		"c2a2fb1c-370c-43cb-8c07-a8035e6b68a2": 5,
		"123e4567-e89b-12d3-a456-426614174011": 10,
	}

	initialQuantities := make(map[string]int)
	for productID := range products {
		initialQuantities[productID] = 100
	}

	err = pgStore.Product().TakeProducts(context.Background(), products)
	if err != nil {
		t.Error("error", err)
	}

	expectedUpdatedQuantities := make(map[string]int)
	for productID, initialQuantity := range initialQuantities {
		expectedUpdatedQuantities[productID] = initialQuantity - products[productID]
	}

	actualUpdatedQuantities := make(map[string]int)
	for productID := range products {
		actualUpdatedQuantities[productID] = 90
	}

	assert.Equal(t, expectedUpdatedQuantities, actualUpdatedQuantities)
}

func TestProductRepo_AddDeliveredpRoducts(t *testing.T) {
	cfg := config.Load()
	log := logger.New(cfg.ServiceName)
	newRedis := redis.New(cfg)
	pgStore, err := New(context.Background(), cfg, log, newRedis)
	if err != nil {
		t.Error("error while connecting to db ", err)
	}

	branchId := "aa541fcc-bf74-11ee-ae0b-166244b65504"
	//test qilinadi agar notenoughproducts nil bo'sa

	if err := pgStore.Product().AddDeliveredProducts(context.Background(), models.DeliverProducts{}, branchId); err != nil {
		t.Error("error:", err)
	}

	deliveredproducts := models.DeliverProducts{
		NotEnoughProducts: map[string]int{
			"123e4567-e89b-12d3-a456-426614174011":  5,
			"c2a2fb1c-370c-43cb-8c07-a8035e6b68a2":  2,
			"f121a0d5-e440-456f-8235-6962328f7ab4 ": 10,
		},
	}

	if err = pgStore.Product().AddDeliveredProducts(context.Background(), deliveredproducts, branchId); err != nil {
		t.Error("err", err)
	}

}
func TestProductRepo_GetListByIDs(t *testing.T) {

	cfg := config.Load()
	log := logger.New(cfg.ServiceName)
	newRedis := redis.New(cfg)
	pgStore, err := New(context.Background(), cfg, log, newRedis)
	if err != nil {
		t.Error("error while connecting to db ", err)
	}

	productids := []string{
		"f121a0d5-e440-456f-8235-6962328f7ab4",
		"123e4567-e89b-12d3-a456-426614174011",
		"f121a0d5-e440-456f-8235-6962328f7ab4",
	}
	productresp, err := pgStore.Product().GetListByIDs(context.Background(), productids)
	if err != nil {

		t.Error("error", err)
	}

	if len(productresp.Products) != 10 {
		t.Errorf("expected: 10 but got %q", len(productresp.Products))
	}

	assert.Equal(t, len(productresp.Products), 10)
}
