package main

import (
	"database/sql"
	"net/http"

	middlewareLogger "github.com/chi-middleware/logrus-logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"github.com/viking311/books/internal/handlers"
	"github.com/viking311/books/internal/logger"
	"github.com/viking311/books/internal/storage"
)

const databseDSN = ""

var db *sql.DB

func main() {

	logger.Info("server start")

	defer logger.Info("server stop")

	err := initDB(databseDSN)

	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	storage, err := storage.NewDBRepository(db)
	if err != nil {
		logger.Fatal(err)
	}

	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middlewareLogger.Logger("router", logger.Logger))
	r.Use(middleware.Recoverer)

	listHandler := handlers.NewGetAllBooksHandler(storage)
	r.Get("/books", listHandler.ServeHTTP)

	getBookHandler := handlers.NewGetByIdHandler(storage)
	r.Get("/book/{id}", getBookHandler.ServeHTTP)

	deleteBookHandler := handlers.NewDeleteByIdHandler(storage)
	r.Delete("/book/{id}", deleteBookHandler.ServeHTTP)

	updateBookHandler := handlers.NewUpdateBookHandler(storage)
	r.Post("/book", updateBookHandler.ServeHTTP)
	r.Put("/book", updateBookHandler.ServeHTTP)
	r.Post("/book/{id}", updateBookHandler.ServeHTTP)
	r.Put("/book/{id}", updateBookHandler.ServeHTTP)

	logger.Fatal(http.ListenAndServe(":8080", r))
}

func initDB(dsn string) error {
	var err error

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	return db.Ping()
}
