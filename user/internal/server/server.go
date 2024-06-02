package server

import (
	"GEO_API/user/internal/repository"
	"GEO_API/user/internal/service"
	"GEO_API/user/pkg/database"
	proto "GEO_API/user/pkg/gRPC/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Server struct {
	proto.UserServiceServer
}

func NewServer() Server {
	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}

	authRepository := repository.New(db.DB)

	newService := service.NewService(authRepository)
	return Server{
		UserServiceServer: newService,
	}
}

func (s *Server) Start() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	reflection.Register(server)
	proto.RegisterUserServiceServer(server, s.UserServiceServer)

	log.Printf("server listening at: %v", ":50051")

	if err = server.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
