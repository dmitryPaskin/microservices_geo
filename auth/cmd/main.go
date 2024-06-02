package main

import "GEO_API/auth/internal/server"

func main() {
	s := server.New()

	s.Start()
}
