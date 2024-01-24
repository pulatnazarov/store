package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test/api/models"
)

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

func (h Handler) DeleteBasket(c *gin.Context) {
	uid := c.Param("id")

	basketID := models.PrimaryKey{ID: uid}
	if err := h.storage.Basket().Delete(basketID); err != nil {
		handleResponse(c, "error is while deleting basket", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, nil)
}
