package handler

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"test/api/models"
	"time"
)

// CreateUser godoc
// @Router       /user [POST]
// @Summary      Creates a new user
// @Description  create a new user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user body models.CreateUser false "user"
// @Success      201  {object}  models.User
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateUser(c *gin.Context) {
	createUser := models.CreateUser{}

	if err := c.ShouldBindJSON(&createUser); err != nil {
		handleResponse(c, h.log, "error while reading body from client", http.StatusBadRequest, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := h.services.User().Create(ctx, createUser)
	if err != nil {
		handleResponse(c, h.log, "error while creating user", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusCreated, resp)
}

// GetUser godoc
// @Router       /user/{id} [GET]
// @Summary      Gets user
// @Description  get user by ID
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id path string true "user"
// @Success      200  {object}  models.User
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetUser(c *gin.Context) {
	var err error

	uid := c.Param("id")

	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, h.log, "invalid uuid type", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	user, err := h.services.User().GetUser(ctx, models.PrimaryKey{
		ID: id.String(),
	})
	if err != nil {
		handleResponse(c, h.log, "error while getting user by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, user)
}

// GetUserList godoc
// @Router       /users [GET]
// @Summary      Get user list
// @Description  get user list
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param 		 limit query string false "limit"
// @Param 		 search query string false "search"
// @Success      200  {object}  models.UsersResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetUserList(c *gin.Context) {
	var (
		page, limit int
		search      string
		err         error
	)

	pageStr := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(c, h.log, "error while parsing page", http.StatusBadRequest, err.Error())
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		handleResponse(c, h.log, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}

	search = c.Query("search")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := h.services.User().GetUsers(ctx, models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		handleResponse(c, h.log, "error while getting users", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, h.log, "success!", http.StatusOK, resp)
}

// UpdateUser godoc
// @Router       /user/{id} [PUT]
// @Summary      Update user
// @Description  update user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param 		 id path string true "user_id"
// @Param        user body models.UpdateUser true "user"
// @Success      200  {object}  models.User
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateUser(c *gin.Context) {
	updateUser := models.UpdateUser{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c, h.log, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	updateUser.ID = uid

	if err := c.ShouldBindJSON(&updateUser); err != nil {
		handleResponse(c, h.log, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := h.services.User().Update(ctx, updateUser)
	if err != nil {
		handleResponse(c, h.log, "error while updating user", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, resp)
}

// DeleteUser godoc
// @Router       /user/{id} [DELETE]
// @Summary      Delete user
// @Description  delete user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param 		 id path string true "user_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteUser(c *gin.Context) {
	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c, h.log, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err = h.services.User().Delete(ctx, models.PrimaryKey{
		ID: id.String(),
	}); err != nil {
		handleResponse(c, h.log, "error while deleting user by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, "data successfully deleted")
}

// UpdateUserPassword godoc
// @Router       /user/{id} [PATCH]
// @Summary      Update user password
// @Description  update user password
// @Tags         user
// @Accept       json
// @Produce      json
// @Param 		 id path string true "user_id"
// @Param        user body models.UpdateUserPassword true "user"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateUserPassword(c *gin.Context) {
	updateUserPassword := models.UpdateUserPassword{}

	if err := c.ShouldBindJSON(&updateUserPassword); err != nil {
		handleResponse(c, h.log, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	uid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		handleResponse(c, h.log, "error while parsing uuid", http.StatusBadRequest, err.Error())
		return
	}

	updateUserPassword.ID = uid.String()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err = h.services.User().UpdatePassword(ctx, updateUserPassword); err != nil {
		handleResponse(c, h.log, "error while updating user password", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, "password successfully updated")
}
