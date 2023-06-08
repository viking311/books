package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/viking311/books/internal/logger"
	"github.com/viking311/books/internal/repository"
	"github.com/viking311/cache"
)

type DeleteByIdHandler struct {
	Server
}

// @Summary      Delete book
// @Description  delete book by ID
// @Param        id   path      int  true  "Book ID"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /book/{id} [delete]
func (dih *DeleteByIdHandler) HAndle(c *gin.Context) {
	bookIdStr := c.Param("id")
	bookId, err := strconv.ParseInt(bookIdStr, 10, 64)
	if err != nil {
		logger.Error(err)
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = dih.storage.Delete(bookId)
	if err != nil {
		logger.Error(err)
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	dih.cache.Delete(fmt.Sprintf(itemCacheKey, bookId))
	dih.cache.Delete(listCacheKey)
}

func NewDeleteByIdHandler(resp repository.Repository, cache *cache.Cache) *DeleteByIdHandler {
	return &DeleteByIdHandler{
		Server: Server{
			storage: resp,
			cache:   cache,
		},
	}
}
