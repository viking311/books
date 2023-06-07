package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/viking311/books/internal/logger"
	"github.com/viking311/books/internal/repository"
)

type DeleteByIdHandler struct {
	Server
}

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

}

func NewDeleteByIdHandler(resp repository.Repository) *DeleteByIdHandler {
	return &DeleteByIdHandler{
		Server: Server{
			storage: resp,
		},
	}
}
