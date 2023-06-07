package server

import (
	"github.com/gin-gonic/gin"
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

	router := gin.Default()

	listHandler := rest.NewGetAllBooksHandler(storage)
	router.GET("books", listHandler.Handle)

	getBookHandler := rest.NewGetByIdHandler(storage)
	router.GET("book/:id", getBookHandler.Handle)

	deleteBookHandler := rest.NewDeleteByIdHandler(storage)
	router.DELETE("book/:id", deleteBookHandler.HAndle)

	updateBookHandler := rest.NewUpdateBookHandler(storage)
	router.POST("book", updateBookHandler.Handle)
	router.PUT("book", updateBookHandler.Handle)
	router.POST("book/:id", updateBookHandler.Handle)
	router.PUT("book/:id", updateBookHandler.Handle)

	router.Run(config.Cfg.Server.GetAddr())
}
