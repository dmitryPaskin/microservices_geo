package gRPC

import (
	"GeoAPI/geo/internal/repository"
	"GeoAPI/geo/internal/service"
	"GeoAPI/geo/internal/service/gRPCFunc"
	"GeoAPI/geo/internal/service/proxyService"
	"GeoAPI/geo/pkg/database"
	proto "GeoAPI/geo/pkg/gRPC/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Server struct {
	gRPCFunc.ServiceGRPC
}

func NewServer() Server {
	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.New(db.DB)
	newService := service.NewService("e6b91900da8a4f3c5138bc921a882ee75d42922a", "943062a0ae098458484fa91f7947fd31c3f549df")
	serviceProxy, err := proxyService.NewProxyService(newService, repo, "redis:6379", "")
	if err != nil {
		log.Fatal(err)
	}
	return Server{
		gRPCFunc.ServiceGRPC{
			Service: serviceProxy,
		},
	}
}

func (s *Server) Start() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	reflection.Register(server)
	proto.RegisterGeoServiceServer(server, &s.ServiceGRPC)

	log.Printf("server listening at: %v", ":50051")

	if err = server.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
