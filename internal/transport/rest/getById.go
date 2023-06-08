package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/viking311/books/internal/config"
	"github.com/viking311/books/internal/domain"
	"github.com/viking311/books/internal/logger"
	"github.com/viking311/books/internal/repository"
	"github.com/viking311/cache"
)

type GetByIdHandler struct {
	Server
}

// @Summary      Get book
// @Description  get book by ID
// @Param        id   path      int  true  "Book ID"
// @Success      200 {object} domain.Book
// @Failure      400
// @Failure      500
// @Router       /book/{id} [get]
func (gbi *GetByIdHandler) Handle(c *gin.Context) {
	bookIdStr := c.Param("id")
	bookId, err := strconv.ParseInt(bookIdStr, 10, 64)
	if err != nil {
		logger.Error(err)
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var book *domain.Book
	if b := gbi.cache.Get(fmt.Sprintf(itemCacheKey, bookId)); !(b == nil) {
		bookT := b.(domain.Book)
		book = &bookT
		logger.Debug("read from cache")
	} else {
		bookT, err := gbi.storage.GetByID(bookId)
		if err != nil {
			logger.Error(err)
			c.Writer.WriteHeader(http.StatusNotFound)
			return
		}
		book = &bookT
		gbi.cache.Set(fmt.Sprintf(itemCacheKey, bookId), bookT, config.Cfg.Cache.TTL)
		logger.Debug("read from db")
	}
	body, err := json.Marshal(book)
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

func NewGetByIdHandler(rep repository.Repository, cache *cache.Cache) *GetByIdHandler {
	return &GetByIdHandler{
		Server: Server{
			storage: rep,
			cache:   cache,
		},
	}
}
