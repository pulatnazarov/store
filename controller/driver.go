package controller

import (
	"fmt"
	"test/models"
)

func (c Controller) CreateDriver() {
	driver := getDriverInfo()

	if !checkPhoneNumber(driver.Phone) {
		fmt.Println("the phone number format is not correct!")
		return
	}

	id, err := c.Store.DriverStorage.Insert(driver)
	if err != nil {
		fmt.Println("error while inserting driver inside controller err: ", err.Error())
		return
	}

	fmt.Println("your new driver's id is: ", id)
}

func checkPhoneNumber(phone string) bool {
	for _, r := range phone {
		if r > '0' || r < '9' || r != '+' {
			return false
		}
	}

	return true
}

func getDriverInfo() models.Driver {
	var (
		fullName, phone string
	)

	fmt.Print("enter driver's full name: ")
	fmt.Scan(&fullName)

	fmt.Println("enter phone: ")
	fmt.Scan(&phone)

	return models.Driver{
		FullName: fullName,
		Phone:    phone,
	}
}
