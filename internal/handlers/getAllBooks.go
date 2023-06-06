package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/viking311/books/internal/logger"
	"github.com/viking311/books/internal/storage"
)

type GetAllBooksHandler struct {
	Server
}

func (gab *GetAllBooksHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	list, err := gab.storage.GetAll()
	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(list)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(body)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func NewGetAllBooksHandler(rep storage.Repository) *GetAllBooksHandler {
	return &GetAllBooksHandler{
		Server: Server{
			storage: rep,
		},
	}
}
