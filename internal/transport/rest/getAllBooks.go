package rest

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viking311/books/internal/config"
	"github.com/viking311/books/internal/domain"
	"github.com/viking311/books/internal/logger"
	"github.com/viking311/books/internal/repository"
	"github.com/viking311/cache"
)

type GetAllBooksHandler struct {
	Server
}

// @Summary      Get all books
// @Description  get all books
// @Success      200 {array} domain.Book
// @Failure      400
// @Failure      500
// @Router       /books [get]

func (gab *GetAllBooksHandler) Handle(c *gin.Context) {
	var list *[]domain.Book
	if t := gab.cache.Get("books"); !(t == nil) {
		listT := t.([]domain.Book)
		list = &listT
		logger.Debug("read from cache")
	} else {
		listT, err := gab.storage.GetAll()
		if err != nil {
			logger.Error(err)
			c.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		list = &listT
		gab.cache.Set(listCacheKey, listT, config.Cfg.Cache.TTL)
		logger.Debug("read from db")
	}

	body, err := json.Marshal(list)
	if err != nil {
		logger.Error(err)
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	c.Writer.Header().Add("Content-Type", "application/json")
	_, err = c.Writer.Write(body)
	if err != nil {
		logger.Error(err)
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func NewGetAllBooksHandler(rep repository.Repository, cache *cache.Cache) *GetAllBooksHandler {
	return &GetAllBooksHandler{
		Server: Server{
			storage: rep,
			cache:   cache,
		},
	}
}
