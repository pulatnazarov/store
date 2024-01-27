package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test/api/models"
)

// StartSell godoc
// @Router       /sell [PUT]
// @Summary      sale products
// @Description  sale products
// @Tags         sale
// @Accept       json
// @Produce      json
// @Param        product query string true "product"
// @Param        quantity query string true "quantity"
// @Param        user body models.UserSell false "user"
// @Success      201  {object}  models.User
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) StartSell(c *gin.Context) {
	productName := c.Query("product")
	quantityStr := c.Query("quantity")
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		handleResponse(c, "error is while converting quantity", http.StatusBadRequest, err.Error())
		return
	}

	searchedProduct, err := h.storage.Repo().Search(productName, uint(quantity))
	if err != nil {
		handleResponse(c, "error is while searching product: we do not have enough product", http.StatusInternalServerError, err.Error())
		return
	}

	user := models.UserSell{}
	if err = c.ShouldBindJSON(&user); err != nil {
		handleResponse(c, "error is while reading from query", http.StatusBadRequest, err.Error())
		return
	}
	ticket, err := h.storage.Store().Sell(searchedProduct, user)
	if err != nil {
		handleResponse(c, "error is while sell", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, ticket)
}
