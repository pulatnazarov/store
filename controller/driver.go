package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"test/models"
	"test/pkg/check"
)

func (c Controller) Driver(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateDriver(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.GetDriver(w, r)
		} else {
			c.GetDriverList(w, r)
		}
	case http.MethodPut:
		// put
	case http.MethodDelete:
		// delete
	}
}

func (c Controller) CreateDriver(w http.ResponseWriter, r *http.Request) {
	driver := models.Driver{}

	if err := json.NewDecoder(r.Body).Decode(&driver); err != nil {
		fmt.Println("error while reading data from client", err.Error())
		hanldeResponse(w, http.StatusBadRequest, err)
		return
	}

	if !check.PhoneNumber(driver.Phone) {
		fmt.Println("the phone number format is not correct!")
		hanldeResponse(w, http.StatusBadRequest, errors.New("phone type is not correct!"))
		return
	}

	// id, err := c.Storage.DriverStorage.Insert(driver)
	// if err != nil {
	// 	fmt.Println("error while inserting driver inside controller err: ", err.Error())
	// 	hanldeResponse(w, http.StatusInternalServerError, err)
	// 	return
	// }

	// resp, err := c.Storage.DriverStorage.GetByID(id)
	// if err != nil {
	// 	//handle error
	// 	return
	// }

	// hanldeResponse(w, http.StatusOK, resp)
}

func (c Controller) GetDriver(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]
	var err error

	// driver, err := c.Storage.DriverStorage.GetByID(id)
	if err != nil {
		fmt.Println("error while getting driver by id")
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	hanldeResponse(w, http.StatusOK, id)
}

func (c Controller) GetDriverList(w http.ResponseWriter, r *http.Request) {
	// cars, err := c.Storage.CarStorage.GetList()
	var err error
	if err != nil {
		fmt.Println("error while getting list of cars:", err.Error())
		hanldeResponse(w, http.StatusInternalServerError, err)
		return
	}

	hanldeResponse(w, http.StatusOK, nil)
}

func (c Controller) UpdateDriver(w http.ResponseWriter, r *http.Request) {
	// var

	// read body

	// send to db for updating body

	// return status ok
}
