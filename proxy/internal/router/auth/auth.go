package auth

import (
	"GEO_API/proxy/internal/controller/handlers/auth"
	_ "GEO_API/proxy/internal/router/auth/docs"
	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title GEO API
// @version 2.0
// @description This is a sample API for address searching and geocoding using Dadata API.
// @host localhost:8080
// @termsOfService http://localhost:8080/auth/swagger/index.html
// @BasePath /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
type RouterAuth struct {
	handler auth.HandlerAuth
}

func New(handler auth.HandlerAuth) RouterAuth {
	return RouterAuth{
		handler: handler,
	}
}

func (r *RouterAuth) AuthRouter(mux chi.Router) {
	mux.Get("/auth/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/auth/swagger/doc.json"), httpSwagger.InstanceName("auth")))

	mux.Post("/api/auth/register", r.handler.SingUpHandler)
	mux.Post("/api/auth/login", r.handler.SingInHandler)

}
