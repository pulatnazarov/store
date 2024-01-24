package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test/api/models"
)

func (h Handler) CreateProduct(c *gin.Context) {
	product := models.CreateProduct{}

	if err := c.ShouldBindJSON(&product); err != nil {
		handleResponse(c, "error is while reading body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Product().Create(product)
	if err != nil {
		handleResponse(c, "error is while creating product", http.StatusInternalServerError, err.Error())
		return
	}

	createdProduct, err := h.storage.Product().GetByID(models.PrimaryKey{ID: id})
	if err != nil {
		handleResponse(c, "error is while getting by id product", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusCreated, createdProduct)
}

func (h Handler) GetProduct(c *gin.Context) {
	uid := c.Param("id")

	product, err := h.storage.Product().GetByID(models.PrimaryKey{ID: uid})
	if err != nil {
		handleResponse(c, "error is while getting by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, product)
}

func (h Handler) GetProductList(c *gin.Context) {
	var (
		page, limit int
		search      string
		err         error
	)

	pageStr := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(c, "error is while converting page", http.StatusBadRequest, err.Error())
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		handleResponse(c, "error is while converting limit", http.StatusBadRequest, err.Error())
		return
	}

	search = c.Query("search")

	products, err := h.storage.Product().GetList(models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

	if err != nil {
		handleResponse(c, "error is while getting list", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, products)
}

func (h Handler) UpdateProduct(c *gin.Context) {
	uid := c.Param("id")

	product := models.UpdateProduct{}

	if err := c.ShouldBindJSON(&product); err != nil {
		handleResponse(c, "error is while reading body", http.StatusBadRequest, err.Error())
		return
	}

	product.ID = uid

	id, err := h.storage.Product().Update(product)
	if err != nil {
		handleResponse(c, "error is while updating product", http.StatusInternalServerError, err.Error())
		return
	}

	updatedProduct, err := h.storage.Product().GetByID(models.PrimaryKey{ID: id})
	if err != nil {
		handleResponse(c, "error is while getting by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, updatedProduct)
}

func (h Handler) DeleteProduct(c *gin.Context) {
	uid := c.Param("id")

	if err := h.storage.Product().Delete(models.PrimaryKey{ID: uid}); err != nil {
		handleResponse(c, "error is while delete", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "product deleted")
}
