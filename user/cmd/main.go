package main

import "GEO_API/user/internal/server"

func main() {
	newServer := server.NewServer()
	newServer.Start()
}
