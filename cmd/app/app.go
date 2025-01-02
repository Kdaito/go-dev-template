package main

import (
	"go-dev-sample/internal/application/server"
)

func main() {
	server := server.NewServer()
	server.Start()
}
