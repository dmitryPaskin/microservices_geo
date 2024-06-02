package user

import (
	"GEO_API/proxy/internal/controller/handlers/user"
	"GEO_API/proxy/internal/middleware/auth"
	_ "GEO_API/proxy/internal/router/user/docs"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title GEO API
// @version 2.0
// @description This is a sample API for address searching and geocoding using Dadata API.
// @host localhost:8080
// @termsOfService http://localhost:8080/user/swagger/index.html
// @BasePath /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
type RouterUser struct {
	handler user.HandlerUser
}

func New(handler user.HandlerUser) RouterUser {
	return RouterUser{
		handler: handler,
	}
}

func (r *RouterUser) UserRouter(mux chi.Router) {
	tokenAuth := jwtauth.New("HS256", []byte("mySecret"), nil)
	mux.Get("/user/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/user/swagger/doc.json"), httpSwagger.InstanceName("user")))

	mux.Group(func(router chi.Router) {
		router.Use(jwtauth.Verify(tokenAuth))
		router.Use(auth.Authenticator)

		mux.Post("/api/user/profile", r.handler.GetCurrentUser)
		mux.Post("/api/user/list", r.handler.GetListUsers)
	})
}
