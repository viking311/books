package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/viking311/books/internal/config"
	"github.com/viking311/books/internal/domain"
	"github.com/viking311/books/internal/logger"
	"github.com/viking311/books/internal/repository"
	"github.com/viking311/cache"
)

// @Summary      Update/create book
// @Description  Update/create book
// @Param        id   path      int  false  "Book ID"
// @Success      200 {object} domain.Book
// @Failure      400
// @Failure      500
// @Router       /book/{id} [post,put]

type UpdateBookHandler struct {
	Server
}

func (ubh *UpdateBookHandler) Handle(c *gin.Context) {
	contentType := c.Request.Header.Get("Content-Type")

	if contentType != "application/json" {
		logger.Error("incorrect content type")
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	bookIdStr := c.Param("id")
	var bookID int64
	if len(bookIdStr) > 0 {
		bookId, err := strconv.ParseInt(bookIdStr, 10, 64)
		bookID = bookId
		if err != nil {
			logger.Error(err)
			c.Writer.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Error(err)
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	defer c.Request.Body.Close()

	var book domain.Book
	err = json.Unmarshal(body, &book)
	if err != nil {
		logger.Error(err)
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	book.Id = bookID

	bookID, err = ubh.storage.Update(book)
	if err != nil {
		logger.Error(err)
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	newBook, err := ubh.storage.GetByID(bookID)
	if err != nil {
		logger.Error(err)
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	ubh.cache.Set(fmt.Sprintf(itemCacheKey, bookID), newBook, config.Cfg.Cache.TTL)
	ubh.cache.Delete(listCacheKey)
	respBody, err := json.Marshal(newBook)

	if err != nil {
		logger.Error(err)
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Add("Content-Type", "application/json")
	_, err = c.Writer.Write(respBody)
	if err != nil {
		logger.Error(err)
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func NewUpdateBookHandler(resp repository.Repository, cache *cache.Cache) *UpdateBookHandler {
	return &UpdateBookHandler{
		Server: Server{
			storage: resp,
			cache:   cache,
		},
	}
}
