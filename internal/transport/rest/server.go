package rest

import (
	"github.com/viking311/books/internal/repository"
	"github.com/viking311/cache"
)

const itemCacheKey = "book-%d"
const listCacheKey = "books"

type Server struct {
	storage repository.Repository
	cache   *cache.Cache
}
