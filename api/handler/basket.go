package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test/api/models"
)

func (h Handler) Basket(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateBasket(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		if _, ok := values["id"]; !ok {
			h.GetBasketList(w, r)
		} else {
			h.GetBasketByID(w, r)
		}
	case http.MethodPut:
		h.UpdateBasket(w, r)
	case http.MethodDelete:
		h.DeleteBasket(w, r)
	}
}

func (h Handler) CreateBasket(w http.ResponseWriter, r *http.Request) {
	createBasket := models.CreateBasket{}

	if err := json.NewDecoder(r.Body).Decode(&createBasket); err != nil {
		fmt.Println("error is while decoding", err.Error())
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	basket, err := h.storage.Basket().CreateBasket(createBasket)
	if err != nil {
		fmt.Println("error is while creating basket", err.Error())
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	id := models.PrimaryKey{ID: basket.ID}
	res, err := h.storage.Basket().GetBasketByID(id)
	if err != nil {
		fmt.Println("error is while getting by id", err.Error())
		return
	}
	handleResponse(w, http.StatusCreated, res)
}

func (h Handler) GetBasketByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	idBasket := models.PrimaryKey{ID: id}

	basket, err := h.storage.Basket().GetBasketByID(idBasket)
	if err != nil {
		fmt.Println("error is while getting by id", err.Error())
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, basket)
}

func (h Handler) GetBasketList(w http.ResponseWriter, r *http.Request) {
	baskets, err := h.storage.Basket().GetBasketList()
	if err != nil {
		fmt.Println("error is while getting list", err.Error())
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, baskets)
}

func (h Handler) UpdateBasket(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	updatedBasket := models.UpdateBasket{}

	if err := json.NewDecoder(r.Body).Decode(&updatedBasket); err != nil {
		fmt.Println("error is while decoding ", err.Error())
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	if updatedBasket.ID != id {
		fmt.Println("car ID not mismatch")
		handleResponse(w, http.StatusBadRequest, updatedBasket.ID)
		return
	}

	updatedBasket.ID = id

	if _, err := h.storage.Basket().UpdateBasket(updatedBasket); err != nil {
		fmt.Println("error is while updating basket", err.Error())
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	ids := models.PrimaryKey{ID: id}
	res, err := h.storage.Basket().GetBasketByID(ids)
	if err != nil {
		fmt.Println("error is while getting by id", err.Error())
		return
	}

	handleResponse(w, http.StatusOK, res)
}

func (h Handler) DeleteBasket(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	basketID := models.PrimaryKey{ID: id}
	if err := h.storage.Basket().DeleteBasket(basketID); err != nil {
		fmt.Println("error is while deleting basket", err.Error())
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, nil)
}
