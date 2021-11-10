package main

import (
	"os"
	"swift-playground/server"
)

func main() {
	_, err := server.LoadConfig()
	if err != nil {
		os.Exit(3)
	}
	server.StartServer(":8090")
}
