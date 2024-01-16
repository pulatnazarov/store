package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"test/api/models"
	"test/pkg/check"
)

func (h Handler) User(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateUser(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			h.GetUser(w, r)
		} else {
			h.GetUserList(w, r)
		}
	case http.MethodPut:
		h.UpdateUser(w, r)
	case http.MethodDelete:
		h.DeleteUser(w, r)
	case http.MethodPatch:
		h.UpdateUserPassword(w, r)
	}
}

func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	createUser := models.CreateUser{}

	if err := json.NewDecoder(r.Body).Decode(&createUser); err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	pKey, err := h.storage.User().Create(createUser)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	user, err := h.storage.User().GetByID(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusCreated, user)
}

func (h Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		handleResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id := values["id"][0]
	var err error

	user, err := h.storage.User().GetByID(models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, user)
}

func (h Handler) GetUserList(w http.ResponseWriter, r *http.Request) {
	var (
		page, limit = 1, 10
		search      string
		err         error
	)
	values := r.URL.Query()

	if len(values["page"]) > 0 {
		page, err = strconv.Atoi(values["page"][0])
		if err != nil {
			page = 1
		}
	}

	if len(values["limit"]) > 0 {
		limit, err = strconv.Atoi(values["limit"][0])
		if err != nil {
			fmt.Println("limit", values["limit"])
			limit = 10
		}
	}

	if len(values["search"]) > 0 {
		search = values["search"][0]
	}

	resp, err := h.storage.User().GetList(models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, resp)
}

func (h Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	updateUser := models.UpdateUser{}

	if err := json.NewDecoder(r.Body).Decode(&updateUser); err != nil {
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	pKey, err := h.storage.User().Update(updateUser)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.storage.User().GetByID(models.PrimaryKey{
		ID: pKey,
	})
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, user)
}

func (h Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		handleResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id := values["id"][0]

	if err := h.storage.User().Delete(models.PrimaryKey{
		ID: id,
	}); err != nil {
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, "data successfully deleted")
}

func (h Handler) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	updateUserPassword := models.UpdateUserPassword{}

	if err := json.NewDecoder(r.Body).Decode(&updateUserPassword); err != nil {
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if len(updateUserPassword.ID) == 0 {
		handleResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	oldPassword, err := h.storage.User().GetPassword(updateUserPassword.ID)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if oldPassword != updateUserPassword.OldPassword {
		handleResponse(w, http.StatusBadRequest, errors.New("old password is not correct"))
		return
	}

	if err = check.ValidatePassword(updateUserPassword.NewPassword); err != nil {
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.storage.User().UpdatePassword(updateUserPassword); err != nil {
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, "password successfully updated")
}
