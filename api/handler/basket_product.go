package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test/api/models"
)

// CreateBasketProduct godoc
// @Router       /basketProduct [POST]
// @Summary      Creates a new basketProduct
// @Description  create a new basketProduct
// @Tags         basketProduct
// @Accept       json
// @Produce      json
// @Param        basket body models.CreateBasketProduct false "basket"
// @Success      201  {object}  models.BasketProduct
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateBasketProduct(c *gin.Context) {
	basketProduct := models.CreateBasketProduct{}

	if err := c.ShouldBindJSON(&basketProduct); err != nil {
		handleResponse(c, "error is while reading body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.BasketProduct().Create(context.Background(), basketProduct)
	if err != nil {
		handleResponse(c, "error is while creating basket product", http.StatusInternalServerError, err)
		return
	}

	createdBasketProduct, err := h.storage.BasketProduct().GetByID(context.Background(), models.PrimaryKey{ID: id})
	if err != nil {
		handleResponse(c, "error is while getting by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusCreated, createdBasketProduct)
}

// GetBasketProduct godoc
// @Router       /basketProduct/{id} [GET]
// @Summary      Get basketProduct by id
// @Description  get basketProduct by id
// @Tags         basketProduct
// @Accept       json
// @Produce      json
// @Param        id path string true "basketProduct_id"
// @Success      201  {object}  models.BasketProduct
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetBasketProduct(c *gin.Context) {
	uid := c.Param("id")

	basketProduct, err := h.storage.BasketProduct().GetByID(context.Background(), models.PrimaryKey{ID: uid})
	if err != nil {
		handleResponse(c, "error is while getting by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, basketProduct)
}

// GetBasketProductList godoc
// @Router       /basketProducts [GET]
// @Summary      Get basket list
// @Description  get basket list
// @Tags         basketProduct
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      201  {object}  models.BasketProductResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetBasketProductList(c *gin.Context) {
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
		handleResponse(c, "error is while converting page", http.StatusBadRequest, err.Error())
		return
	}

	search = c.Query("search")

	basketProducts, err := h.storage.BasketProduct().GetList(context.Background(), models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

	handleResponse(c, "", http.StatusOK, basketProducts)
}

// UpdateBasketProduct godoc
// @Router       /basketProduct/{id} [PUT]
// @Summary      Update basketProduct
// @Description  update basketProduct
// @Tags         basketProduct
// @Accept       json
// @Produce      json
// @Param        id path string true "basketProduct_id"
// @Param        basketProduct body models.UpdateBasketProduct false "basketProduct"
// @Success      201  {object}  models.BasketProduct
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateBasketProduct(c *gin.Context) {
	basketProduct := models.UpdateBasketProduct{}
	uid := c.Param("id")

	if err := c.ShouldBindJSON(&basketProduct); err != nil {
		handleResponse(c, "error is while reading from body", http.StatusBadRequest, err.Error())
		return
	}

	basketProduct.ID = uid
	id, err := h.storage.BasketProduct().Update(context.Background(), basketProduct)
	if err != nil {
		handleResponse(c, "error is while updating basket", http.StatusInternalServerError, err.Error())
		return
	}

	updatedBasketProduct, err := h.storage.BasketProduct().GetByID(context.Background(), models.PrimaryKey{ID: id})
	if err != nil {
		handleResponse(c, "error is while getting by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, updatedBasketProduct)
}

// DeleteBasketProduct godoc
// @Router       /basketProduct/{id} [Delete]
// @Summary      Delete basketProduct
// @Description  delete basketProduct
// @Tags         basketProduct
// @Accept       json
// @Produce      json
// @Param        id path string true "basketProduct_id"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteBasketProduct(c *gin.Context) {
	uid := c.Param("id")

	if err := h.storage.BasketProduct().Delete(context.Background(), models.PrimaryKey{ID: uid}); err != nil {
		handleResponse(c, "error is while deleting", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "basket product deleted!")
}
