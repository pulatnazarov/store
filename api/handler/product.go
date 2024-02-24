package handler

import (
	"context"
	"net/http"
	"strconv"
	"test/api/models"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateProduct godoc
// @Router       /product [POST]
// @Summary      Create a new product
// @Description  create a new product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param 		 product body models.CreateProduct false "product"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateProduct(c *gin.Context) {
	product := models.CreateProduct{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	if err := c.ShouldBindJSON(&product); err != nil {
		handleResponse(c, h.log, "error is while reading body", http.StatusBadRequest, err.Error())
		return
	}

	createdProduct, err := h.services.Product().Create(ctx, product)
	if err != nil {
		handleResponse(c, h.log, "error is while creating product", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusCreated, createdProduct)
}

// GetProduct godoc
// @Router       /product/{id} [GET]
// @Summary      Get product by id
// @Description  get product by id
// @Tags         product
// @Accept       json
// @Produce      json
// @Param 		 id path string true "product_id"
// @Success      200  {object}  models.Product
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetProduct(c *gin.Context) {
	uid := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	product, err := h.services.Product().Get(ctx, models.PrimaryKey{ID: uid})
	if err != nil {
		handleResponse(c, h.log, "error is while getting by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, product)
}

// GetProductList godoc
// @Router       /products [GET]
// @Summary      Get product list
// @Description  get product list
// @Tags         product
// @Accept       json
// @Produce      json
// @Param 		 page query string false "page"
// @Param 		 limit query string false "limit"
// @Param 		 search query string false "search"
// @Success      200  {object}  models.Product
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetProductList(c *gin.Context) {
	var (
		page, limit int
		search      string
		err         error
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	pageStr := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(c, h.log, "error is while converting page", http.StatusBadRequest, err.Error())
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		handleResponse(c, h.log, "error is while converting limit", http.StatusBadRequest, err.Error())
		return
	}

	search = c.Query("search")

	products, err := h.services.Product().GetList(ctx, models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

	if err != nil {
		handleResponse(c, h.log, "error is while getting list", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, products)
}

// UpdateProduct godoc
// @Router       /product/{id} [PUT]
// @Summary      Update product
// @Description  update product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param 		 id path string true "product_id"
// @Param 		 product body models.UpdateProduct false "product"
// @Success      200  {object}  models.Product
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateProduct(c *gin.Context) {
	uid := c.Param("id")

	product := models.UpdateProduct{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := c.ShouldBindJSON(&product); err != nil {
		handleResponse(c, h.log, "error is while reading body", http.StatusBadRequest, err.Error())
		return
	}

	product.ID = uid

	updatedProduct, err := h.services.Product().Update(ctx, product)
	if err != nil {
		handleResponse(c, h.log, "error is while updating product", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, updatedProduct)
}

// DeleteProduct godoc
// @Router       /product/{id} [DELETE]
// @Summary      Delete product
// @Description  delete product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param 		 id path string true "product_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteProduct(c *gin.Context) {
	uid := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*2)
	defer cancel()
	if err := h.services.Product().Delete(ctx, models.PrimaryKey{ID: uid}); err != nil {
		handleResponse(c, h.log, "error is while delete", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, "product deleted")
}

// StartSellNew godoc
// @Router       /sell-new [POST]
// @Summary      Selling products
// @Description  selling products
// @Tags         product
// @Accept       json
// @Produce      json
// @Param 		 sell_request body models.SellRequest false "sell_request"
// @Success      200  {object}  models.Check
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) StartSellNew(c *gin.Context) {
	request := models.SellRequest{}

	if err := c.ShouldBindJSON(&request); err != nil {
		handleResponse(c, h.log, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	productSell, err := h.services.Product().StartSellNew(ctx, request)
	if err != nil {
		handleResponse(c, h.log, "error is while start sell new", http.StatusInternalServerError, err.Error())
		return
	}

	// dealer

	if err = h.services.Dealer().Delivery(ctx, productSell); err != nil {
		handleResponse(c, h.log, "error is while delivery products", http.StatusInternalServerError, err.Error())
		return
	}

	// report

	handleResponse(c, h.log, "successfully finished the purchase", http.StatusOK, productSell.Check)
}
