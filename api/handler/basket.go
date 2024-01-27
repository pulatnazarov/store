package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test/api/models"
)

// CreateBasket godoc
// @Router       /basket [POST]
// @Summary      Creates a new basket
// @Description  create a new basket
// @Tags         basket
// @Accept       json
// @Produce      json
// @Param        basket body models.CreateBasket false "basket"
// @Success      201  {object}  models.Basket
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateBasket(c *gin.Context) {
	createBasket := models.CreateBasket{}

	if err := c.ShouldBindJSON(&createBasket); err != nil {
		handleResponse(c, "error is while decoding", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Basket().Create(createBasket)
	if err != nil {
		handleResponse(c, "error is while creating basket", http.StatusInternalServerError, err)
		return
	}

	res, err := h.storage.Basket().GetByID(models.PrimaryKey{ID: id})
	if err != nil {
		handleResponse(c, "error is while getting by id", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "", http.StatusCreated, res)
}

// GetBasket godoc
// @Router       /basket/{id} [GET]
// @Summary      Get basket by id
// @Description  get basket by id
// @Tags         basket
// @Accept       json
// @Produce      json
// @Param        id path string true "basket_id"
// @Success      201  {object}  models.Basket
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetBasket(c *gin.Context) {
	var err error
	fmt.Println("came here")

	uid := c.Param("id")

	basket, err := h.storage.Basket().GetByID(models.PrimaryKey{ID: uid})
	if err != nil {
		handleResponse(c, "error is while getting by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, basket)
}

// GetBasketList godoc
// @Router       /baskets [GET]
// @Summary      Get basket list
// @Description  get basket list
// @Tags         basket
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      201  {object}  models.BasketResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetBasketList(c *gin.Context) {
	var (
		page, limit int
		search      string
		err         error
	)

	pageStr := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(c, "error is while converting pageStr", http.StatusBadRequest, err)
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		handleResponse(c, "error is while converting limitStr", http.StatusBadRequest, err)
		return
	}

	search = c.Query("search")

	baskets, err := h.storage.Basket().GetList(models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		handleResponse(c, "error is while getting list", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, baskets)
}

// UpdateBasket godoc
// @Router       /basket/{id} [PUT]
// @Summary      Update basket
// @Description  update basket
// @Tags         basket
// @Accept       json
// @Produce      json
// @Param        id path string true "basket_id"
// @Param        basket body models.UpdateBasket false "basket"
// @Success      201  {object}  models.Basket
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateBasket(c *gin.Context) {
	updatedBasket := models.UpdateBasket{}

	uid := c.Param("id")
	if err := c.ShouldBindJSON(&updatedBasket); err != nil {
		handleResponse(c, "error is while decoding ", http.StatusBadRequest, err)
		return
	}

	updatedBasket.ID = uid

	if _, err := h.storage.Basket().Update(updatedBasket); err != nil {
		handleResponse(c, "error is while updating basket", http.StatusInternalServerError, err)
		return
	}

	ids := models.PrimaryKey{ID: uid}
	res, err := h.storage.Basket().GetByID(ids)
	if err != nil {
		fmt.Println("error is while getting by id", err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, res)
}

// DeleteBasket godoc
// @Router       /basket/{id} [Delete]
// @Summary      Delete basket
// @Description  delete basket
// @Tags         basket
// @Accept       json
// @Produce      json
// @Param        id path string true "basket_id"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteBasket(c *gin.Context) {
	uid := c.Param("id")

	basketID := models.PrimaryKey{ID: uid}
	if err := h.storage.Basket().Delete(basketID); err != nil {
		handleResponse(c, "error is while deleting basket", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, nil)
}
