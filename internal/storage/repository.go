package storage

type Repository interface {
	Update(value Book) (int64, error)
	Delete(id int64) error
	GetByID(id int64) (Book, error)
	GetAll() ([]Book, error)
}
