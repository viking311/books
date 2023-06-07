package main

import "github.com/viking311/books/internal/server"

// @title           Books API
// @version         1.0
// @description     This is a books list server.

// @host      localhost:8080
// @BasePath  /

func main() {
	server.Run()
}
