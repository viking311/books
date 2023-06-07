package rest

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/viking311/books/internal/logger"
	"github.com/viking311/books/internal/repository"
)

type DeleteByIdHandler struct {
	Server
}

func (dih *DeleteByIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bookIdStr := chi.URLParam(r, "id")
	bookId, err := strconv.ParseInt(bookIdStr, 10, 64)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = dih.storage.Delete(bookId)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func NewDeleteByIdHandler(resp repository.Repository) *DeleteByIdHandler {
	return &DeleteByIdHandler{
		Server: Server{
			storage: resp,
		},
	}
}
