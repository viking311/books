package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/viking311/books/internal/logger"
	"github.com/viking311/books/internal/storage"
)

type GetByIdHandler struct {
	Server
}

func (gbi *GetByIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bookIdStr := chi.URLParam(r, "id")
	bookId, err := strconv.ParseInt(bookIdStr, 10, 64)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	book, err := gbi.storage.GetByID(bookId)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	body, err := json.Marshal(book)
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

func NewGetByIdHandler(rep storage.Repository) *GetByIdHandler {
	return &GetByIdHandler{
		Server: Server{
			storage: rep,
		},
	}
}
