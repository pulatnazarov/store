package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"test/api/models"
	"test/storage"
)

type productRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) storage.IProductStorage {
	return &productRepo{db: db}
}

func (p *productRepo) Create(product models.CreateProduct) (string, error) {
	id := uuid.New()
	query := `insert into products(id, name, price, original_price, quantity, category_id) 
						values($1, $2, $3, $4, $5, $6)`

	if _, err := p.db.Exec(context.Background(), query,
		id,
		product.Name,
		product.Price,
		product.OriginalPrice,
		product.Quantity,
		product.CategoryID); err != nil {
		fmt.Println("error is while inserting product", err.Error())
		return "", err
	}

	return id.String(), nil
}

func (p *productRepo) GetByID(key models.PrimaryKey) (models.Product, error) {
	product := models.Product{}
	query := `select id, name, price, original_price, quantity, category_id from products where id = $1 `
	if err := p.db.QueryRow(context.Background(), query, key.ID).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.OriginalPrice,
		&product.Quantity,
		&product.CategoryID); err != nil {
		fmt.Println("error is while selecting product by id", err.Error())
		return models.Product{}, err
	}
	return product, nil
}

func (p *productRepo) GetList(request models.GetListRequest) (models.ProductResponse, error) {
	var (
		products          = []models.Product{}
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
		query, countQuery string
		count             = 0
	)

	countQuery = `select count(1) from products `

	if search != "" {
		countQuery += fmt.Sprintf(` where (name ilike '%%%s%%' or 
			CAST(price AS TEXT) ilike '%%%s%%' or CAST(quantity AS TEXT) ilike '%%%s%%')`, search, search, search)
	}

	if err := p.db.QueryRow(context.Background(), countQuery).Scan(&count); err != nil {
		fmt.Println("error is while scanning count", err.Error())
		return models.ProductResponse{}, err
	}

	query = `select id, name, price, original_price, quantity, category_id from products `

	if search != "" {
		query += fmt.Sprintf(` where (name ilike '%%%s%%' or 
			CAST(price AS TEXT) ilike '%%%s%%' or CAST(quantity AS TEXT) ilike '%%%s%%')`, search, search, search)
	}

	query += ` LIMIT $1 OFFSET $2`

	rows, err := p.db.Query(context.Background(), query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting products", err.Error())
		return models.ProductResponse{}, err
	}

	for rows.Next() {
		product := models.Product{}
		if err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.OriginalPrice,
			&product.Quantity,
			&product.CategoryID); err != nil {
			fmt.Println("error is while scanning products", err.Error())
			return models.ProductResponse{}, err
		}
		products = append(products, product)
	}
	return models.ProductResponse{
		Product: products,
		Count:   count,
	}, err
}

func (p *productRepo) Update(product models.UpdateProduct) (string, error) {
	query := `update products set name = $1, price = $2, original_price = $3, quantity = $4, category_id = $5 where id = $6`

	if _, err := p.db.Exec(context.Background(), query,
		&product.Name,
		&product.Price,
		&product.OriginalPrice,
		&product.Quantity,
		&product.CategoryID,
		&product.ID); err != nil {
		fmt.Println("error is while updating product", err.Error())
		return "", err
	}

	return product.ID, nil
}

func (p *productRepo) Delete(key models.PrimaryKey) error {
	query := `delete from products where id = $1`

	if _, err := p.db.Exec(context.Background(), query, key.ID); err != nil {
		fmt.Println("error is while deleting product", err.Error())
		return err
	}
	return nil
}

func (p *productRepo) Search(customerProductIDs map[string]int) (map[string]int, map[string]int, error) {
	var (
		selectedProducts = models.SellRequest{
			Products: map[string]int{},
		}
		products         = make([]string, len(customerProductIDs))
		productQuantites []int
		productPrices    = make(map[string]int, 0)
	)

	for key, value := range customerProductIDs {
		products = append(products, key)
		productQuantites = append(productQuantites, value)
	}

	query := `
				select id, quantity, price, original_price from products where id::varchar = ANY($1)
	`

	rows, err := p.db.Query(context.Background(), query, pq.Array(products), pq.Array(productQuantites)) // [a, b, c]
	if err != nil {
		fmt.Println("Error while getting products by product ids", err.Error())
		return nil, nil, err
	}

	for rows.Next() {
		var (
			quantity, price, originalPrice int
			productID                      string
		)
		if err = rows.Scan(
			&productID,
			&quantity,
			&price,
			&originalPrice,
		); err != nil {
			fmt.Println("Error while scanning rows one by one", err.Error())
			return nil, nil, err
		}

		if customerProductIDs[productID] <= quantity {
			selectedProducts.Products[productID] = price
			productPrices[productID] = originalPrice
		}
	}

	return selectedProducts.Products, productPrices, nil
}

func (p *productRepo) TakeProducts(products map[string]int) error {
	query := `
		update products set quantity = quantity - $1 where id = $2
`
	for productID, quantity := range products {
		if _, err := p.db.Exec(context.Background(), query, quantity, productID); err != nil {
			fmt.Println("Error while updating product quantity", err.Error())
			return err
		}
	}

	return nil
}
