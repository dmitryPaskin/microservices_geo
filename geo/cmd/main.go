package main

import "GeoAPI/geo/internal/server/gRPC"

func main() {
	server := gRPC.NewServer()
	server.Start()
}
