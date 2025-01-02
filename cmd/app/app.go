package main

import (
	"go-dev-sample/internal/server"
)

func main() {
	server := server.NewServer()
	server.Start()
}
