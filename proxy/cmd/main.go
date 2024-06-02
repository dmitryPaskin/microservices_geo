package main

import "GEO_API/proxy/internal/router"

func main() {
	r := router.New()
	r.Start()
}
