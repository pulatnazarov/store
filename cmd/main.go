package main

import (
	"fmt"
	"log"
	"net/http"
	"test/config"
	"test/controller"
	"test/storage/postgres"
)

func main() {
	cfg := config.Load()

	store, err := postgres.New(cfg)
	if err != nil {
		log.Fatalln("error while connecting to db err:", err.Error())
		return
	}
	defer store.Close()

	con := controller.New(store)

	http.HandleFunc("/user", con.User)
	http.HandleFunc("/car", con.Car)

	fmt.Println("listening at port :8080")
	http.ListenAndServe(":8080", nil)
}
