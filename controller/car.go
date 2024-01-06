package controller

import (
	"fmt"
	"net/http"
	"test/models"
	"test/pkg/check"
	"time"

	"github.com/google/uuid"
)

func (c Controller) Car(w http.ResponseWriter, r *http.Request)  {
	
}

func (c Controller) CreateCar(w http.ResponseWriter) {
	car := models.Car{}

	if err := check.ValidateCarYear(car.Year); err != nil {
		hanldeResponse(w, http.StatusBadRequest, err)
		return
	}

	id, err := c.Store.CarStorage.Insert(car)
	if err != nil {
		fmt.Println("error while creating data inside controller err: ", err.Error())
		return
	}
	fmt.Println("id: ", id)
}

func (c Controller) GetCarByID() {
	idStr := ""
	fmt.Print("enter id: ")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("id is not uuid err: ", err.Error())
		return
	}

	car, err := c.Store.CarStorage.GetByID(id)
	if err != nil {
		fmt.Print("error while getting car by id err: ", err.Error())
		return
	}
	fmt.Println("your car is: ", car)
}
func (c Controller) GetCarList() {
	cars, err := c.Store.CarStorage.GetList()
	if err != nil {
		fmt.Println("Error while getting list err : ", err.Error())
		return
	}

	fmt.Println(cars)

}
func (c Controller) UpdateCar() {
	car := models.Car{}

	if !checkCarInfo(car) {
		return
	}

	err := c.Store.CarStorage.Update(car)
	if err != nil {
		fmt.Println("error while updating car!", err)
		return
	}
	if car.ID.String() != "" {
		fmt.Println("Successfully updated!")
	} else {
		fmt.Println("Successfully created!")
	}

}

func checkCarInfo(car models.Car) bool {
	if car.Year <= 0 || car.Year > time.Now().Year()+1 {
		fmt.Println("year intput is not correct")
		return false
	}

	return true
}

