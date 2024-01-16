package api

import (
	"net/http"
	"test/api/handler"
	"test/storage"
)

func New(store storage.IStorage) {
	h := handler.New(store)

	http.HandleFunc("/user", h.User)
}
