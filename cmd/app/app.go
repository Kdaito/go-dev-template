package main

import "github.com/Kdaito/kinodokuna-be/internal/server"

func main() {
	server := server.NewServer()
	server.Start()
}
