package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/viking311/books/internal/logger"
	"github.com/viking311/books/internal/repository"
)

type GetByIdHandler struct {
	Server
}

func (gbi *GetByIdHandler) Handle(c *gin.Context) {
	bookIdStr := c.Param("id")
	bookId, err := strconv.ParseInt(bookIdStr, 10, 64)
	if err != nil {
		logger.Error(err)
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	book, err := gbi.storage.GetByID(bookId)
	if err != nil {
		logger.Error(err)
		c.Writer.WriteHeader(http.StatusNotFound)
		return
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

func NewGetByIdHandler(rep repository.Repository) *GetByIdHandler {
	return &GetByIdHandler{
		Server: Server{
			storage: rep,
		},
	}
}
