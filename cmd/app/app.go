package main

import "github.com/Kdaito/go-dev-template/internal/server"

func main() {
	server := server.NewServer()
	server.Start()
}
