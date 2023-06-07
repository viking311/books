package server

import (
	"net/http"

	middlewareLogger "github.com/chi-middleware/logrus-logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/viking311/books/internal/config"
	"github.com/viking311/books/internal/logger"
	"github.com/viking311/books/internal/repository"
	"github.com/viking311/books/internal/transport/rest"
	"github.com/viking311/books/pkg/database"
)

func Run() {
	logger.Info("server start")

	defer logger.Info("server stop")

	db, err := database.NewPostgresConnection(config.Cfg.Database)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	storage, err := repository.NewPostgresRepository(db)
	if err != nil {
		logger.Fatal(err)
	}

	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middlewareLogger.Logger("router", logger.Logger))
	r.Use(middleware.Recoverer)

	listHandler := rest.NewGetAllBooksHandler(storage)
	r.Get("/books", listHandler.ServeHTTP)

	getBookHandler := rest.NewGetByIdHandler(storage)
	r.Get("/book/{id}", getBookHandler.ServeHTTP)

	deleteBookHandler := rest.NewDeleteByIdHandler(storage)
	r.Delete("/book/{id}", deleteBookHandler.ServeHTTP)

	updateBookHandler := rest.NewUpdateBookHandler(storage)
	r.Post("/book", updateBookHandler.ServeHTTP)
	r.Put("/book", updateBookHandler.ServeHTTP)
	r.Post("/book/{id}", updateBookHandler.ServeHTTP)
	r.Put("/book/{id}", updateBookHandler.ServeHTTP)

	logger.Fatal(http.ListenAndServe(config.Cfg.Server.GetAddr(), r))

}
