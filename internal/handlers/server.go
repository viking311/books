package handlers

import "github.com/viking311/books/internal/storage"

type Server struct {
	storage storage.Repository
}
