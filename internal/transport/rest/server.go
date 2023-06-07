package rest

import "github.com/viking311/books/internal/repository"

type Server struct {
	storage repository.Repository
}
