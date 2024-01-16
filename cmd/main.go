package main

import (
	"fmt"
	"log"
	"net/http"
	"test/api"
	"test/config"
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

	api.New(store)

	fmt.Println("listening at port :8080")
	if err = http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln("Server has stopped!", err.Error())
	}
}
