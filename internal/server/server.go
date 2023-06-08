package server

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/viking311/books/docs"
	"github.com/viking311/books/internal/config"
	"github.com/viking311/books/internal/logger"
	"github.com/viking311/books/internal/repository"
	"github.com/viking311/books/internal/transport/rest"
	"github.com/viking311/books/pkg/database"
	"github.com/viking311/cache"
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

	cache := cache.NewCache()

	router := gin.Default()

	listHandler := rest.NewGetAllBooksHandler(storage, cache)
	router.GET("books", listHandler.Handle)

	getBookHandler := rest.NewGetByIdHandler(storage, cache)
	router.GET("book/:id", getBookHandler.Handle)

	deleteBookHandler := rest.NewDeleteByIdHandler(storage, cache)
	router.DELETE("book/:id", deleteBookHandler.HAndle)

	updateBookHandler := rest.NewUpdateBookHandler(storage, cache)
	router.POST("book", updateBookHandler.Handle)
	router.PUT("book", updateBookHandler.Handle)
	router.POST("book/:id", updateBookHandler.Handle)
	router.PUT("book/:id", updateBookHandler.Handle)

	//programmatically set swagger info
	docs.SwaggerInfo.Title = "Books API"
	docs.SwaggerInfo.Description = "This is a books list server"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(config.Cfg.Server.GetAddr())
}
