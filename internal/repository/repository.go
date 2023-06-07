package repository

import "github.com/viking311/books/internal/domain"

type Repository interface {
	Update(value domain.Book) (int64, error)
	Delete(id int64) error
	GetByID(id int64) (domain.Book, error)
	GetAll() ([]domain.Book, error)
}
