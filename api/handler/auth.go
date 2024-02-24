package handler

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"test/api/models"
	"test/pkg/jwt"
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

	loginResponse, err := h.services.AuthService().CustomerLogin(ctx, loginRequest)
	if err != nil {
		handleResponse(c, h.log, "incorrect credentials", http.StatusBadRequest, errors.New("password or login incorrect"))
		return
	}

	handleResponse(c, h.log, "success", http.StatusOK, loginResponse)
}

// AdminLogin godoc
// @Router       /auth/admin/login [POST]
// @Summary      Admin login
// @Description  admin login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login body models.AdminLoginRequest true "login"
// @Success      201  {object}  models.AdminLoginResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) AdminLogin(c *gin.Context) {
	adminLogin := models.AdminLoginRequest{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := c.ShouldBindJSON(&adminLogin); err != nil {
		handleResponse(c, h.log, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	loginResponse, err := h.services.AuthService().AdminLogin(ctx, adminLogin)
	if err != nil {
		handleResponse(c, h.log, "error while admin login", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.log, "success", http.StatusOK, loginResponse)
}

func getAuthInfo(c *gin.Context) (models.AuthInfo, error) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		return models.AuthInfo{}, errors.New("unauthorized")
	}

	m, err := jwt.ExtractClaims(accessToken)
	if err != nil {
		return models.AuthInfo{}, err
	}

	return models.AuthInfo{
		UserID:   m["user_id"].(string),
		UserRole: m["user_role"].(string),
	}, nil
}
