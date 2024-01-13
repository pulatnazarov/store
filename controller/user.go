package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test/models"
)

func (c Controller) User(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateUser(w, r)
	}
}

func (c Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	createUser := models.CreateUser{}

	if err := json.NewDecoder(r.Body).Decode(&createUser); err != nil {
		fmt.Println("error while reading data from client", err.Error())
		hanldeResponse(w, http.StatusBadRequest, err)
		return
	}

	user, err := c.Storage.User().Create(createUser)
	if err != nil {
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	hanldeResponse(w, http.StatusCreated, user)
}
