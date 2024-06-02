package router

import (
	handlerAuth "GEO_API/proxy/internal/controller/handlers/auth"
	handlerGeo "GEO_API/proxy/internal/controller/handlers/geo"
	handlerUser "GEO_API/proxy/internal/controller/handlers/user"
	"GEO_API/proxy/internal/controller/responder"
	"GEO_API/proxy/internal/router/auth"
	"GEO_API/proxy/internal/router/geo"
	"GEO_API/proxy/internal/router/user"
	"GEO_API/proxy/internal/service/proxyService"
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Router struct {
	chi  *chi.Mux
	geo  geo.RouterGeo
	auth auth.RouterAuth
	user user.RouterUser
}

func New() Router {
	var router Router
	respond := responder.NewRespond(zap.NewExample())
	router.chi = chi.NewRouter()
	{
		serviceUser := proxyService.NewUserProxy()
		handlerUser := handlerUser.New(serviceUser, respond)
		router.user = user.New(handlerUser)
	}
	{
		serviceGeo := proxyService.NewGeoProxy()
		handlerGeo := handlerGeo.New(serviceGeo, respond)
		router.geo = geo.New(handlerGeo)
	}
	{
		serviceAuth := proxyService.NewProxyAuth()
		handlerAuth := handlerAuth.New(serviceAuth, respond)
		router.auth = auth.New(handlerAuth)
	}
	return router
}

func (r *Router) Start() {
	r.chi.Use(middleware.Recoverer)
	r.chi.Use(middleware.Logger)

	r.auth.AuthRouter(r.chi)
	r.geo.GeoRouter(r.chi)
	r.user.UserRouter(r.chi)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r.chi,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	sigChan := make(chan os.Signal, 1)
	defer close(sigChan)
	signal.Notify(sigChan, syscall.SIGINT)

	go func() {
		log.Println("Starting server...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	<-sigChan
	stopCTX, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := server.Shutdown(stopCTX); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
}
