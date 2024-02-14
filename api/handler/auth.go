package handler

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"test/api/models"
	"time"
)

// CustomerLogin godoc
// @Router       /auth/customer/login [POST]
// @Summary      Customer login
// @Description  Customer login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login body models.CustomerLoginRequest true "login"
// @Success      201  {object}  models.Basket
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CustomerLogin(c *gin.Context) {
	loginRequest := models.CustomerLoginRequest{}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		handleResponse(c, h.log, "error while binding body", http.StatusBadRequest, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := h.services.AuthService().CustomerLogin(ctx, loginRequest); err != nil {
		handleResponse(c, h.log, "incorrect credentials", http.StatusBadRequest, errors.New("password or login incorrect"))
		return
	}

	handleResponse(c, h.log, "success", http.StatusOK, "")
}
