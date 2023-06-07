package rest

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viking311/books/internal/logger"
	"github.com/viking311/books/internal/repository"
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
	list, err := gab.storage.GetAll()
	if err != nil {
		logger.Error(err)
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
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

func NewGetAllBooksHandler(rep repository.Repository) *GetAllBooksHandler {
	return &GetAllBooksHandler{
		Server: Server{
			storage: rep,
		},
	}
}
