package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/viking311/books/internal/logger"
	"github.com/viking311/books/internal/storage"
)

type UpdateBookHandler struct {
	Server
}

func (ubh *UpdateBookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")

	if contentType != "application/json" {
		logger.Error("incorrect content type")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bookIdStr := chi.URLParam(r, "id")
	var bookID int64
	if len(bookIdStr) > 0 {
		bookId, err := strconv.ParseInt(bookIdStr, 10, 64)
		bookID = bookId
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var book storage.Book
	err = json.Unmarshal(body, &book)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	book.Id = bookID

	bookID, err = ubh.storage.Update(book)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	newBook, err := ubh.storage.GetByID(bookID)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respBody, err := json.Marshal(newBook)

	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(respBody)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func NewUpdateBookHandler(resp storage.Repository) *UpdateBookHandler {
	return &UpdateBookHandler{
		Server: Server{
			storage: resp,
		},
	}
}
