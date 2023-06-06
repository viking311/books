package storage

import (
	"database/sql"
	"fmt"
)

type DBRepository struct {
	db *sql.DB
}

func (dbr *DBRepository) Delete(id int64) error {
	_, err := dbr.db.Exec("DELETE FROM books WHERE id=$1", id)
	if err != nil {
		return err
	}

	return nil
}

func (dbr *DBRepository) Update(value Book) (int64, error) {
	if value.Id == 0 {
		var lastinsertedId int64
		err := dbr.db.QueryRow("INSERT INTO books(title, author, publish_year) VALUES($1,$2,$3) RETURNING id", value.Title, value.Author, value.Year).Scan(&lastinsertedId)
		if err != nil {
			return 0, err
		}
		return lastinsertedId, nil
	} else {
		_, err := dbr.db.Exec("UPDATE books SET title=$1, author=$2, publish_year=$3 WHERE id=$4", value.Title, value.Author, value.Year, value.Id)
		if err != nil {
			return 0, err
		}
		return value.Id, nil
	}
}

func (dbr *DBRepository) GetByID(id int64) (Book, error) {
	var book Book
	err := dbr.db.QueryRow("SELECT id, title, author, publish_year FROM books WHERE id=$1", id).Scan(&book.Id, &book.Title, &book.Author, &book.Year)
	if err != nil {
		if err == sql.ErrNoRows {
			return book, fmt.Errorf("not found book with id %d", id)
		}
		return book, err
	}

	return book, nil
}

func (dbr *DBRepository) GetAll() ([]Book, error) {
	rows, err := dbr.db.Query("SELECT id, title, author, publish_year FROM books")
	if err != nil {
		return []Book{}, err
	}
	defer rows.Close()

	var slice []Book = []Book{}
	for rows.Next() {
		var book Book
		err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.Year)
		if err != nil {
			return []Book{}, err
		}
		slice = append(slice, book)
	}
	err = rows.Err()
	if err != nil {
		return []Book{}, err
	}

	return slice, nil
}

func NewDBRepository(db *sql.DB) (*DBRepository, error) {

	if db == nil {
		return nil, fmt.Errorf("db instance is needed")
	}
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS books (id Serial PRIMARY KEY, title TEXT NOT NULL, author TEXT NOT NULL, publish_year smallint)")

	if err != nil {
		return nil, err
	}

	return &DBRepository{
		db: db,
	}, nil
}
