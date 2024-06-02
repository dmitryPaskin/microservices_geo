package geo

import (
	"GEO_API/proxy/internal/controller/handlers/geo"
	"GEO_API/proxy/internal/middleware/auth"
	_ "GEO_API/proxy/internal/router/geo/docs"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title GEO API
// @version 2.0
// @description This is a sample API for address searching and geocoding using Dadata API.
// @host localhost:8080
// @termsOfService http://localhost:8080/geo/swagger/index.html
// @BasePath /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
type RouterGeo struct {
	handler geo.HandlerGeo
}

func New(handler geo.HandlerGeo) RouterGeo {
	return RouterGeo{
		handler: handler,
	}
}

func (r *RouterGeo) GeoRouter(mux chi.Router) {
	tokenAuth := jwtauth.New("HS256", []byte("mySecret"), nil)
	mux.Get("/geo/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/geo/swagger/doc.json"), httpSwagger.InstanceName("geo")))

	mux.Group(func(router chi.Router) {
		router.Use(jwtauth.Verify(tokenAuth))
		router.Use(auth.Authenticator)

		router.Post("/api/address/search", r.handler.SearchAddressHandler)
		router.Post("/api/address/geocode", r.handler.GeocodeHandler)
	})
}
