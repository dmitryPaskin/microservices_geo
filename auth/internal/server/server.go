package server

import (
	"GEO_API/auth/internal/service"
	proto "GEO_API/auth/pkg/gRPC/api/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Server struct {
	service.Service
}

func New() Server {
	ser := service.New("user:50051")

	return Server{
		ser,
	}
}

func (s *Server) Start() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	reflection.Register(server)
	proto.RegisterAuthServiceServer(server, &s.Service)

	log.Printf("server listening at: %v", ":50051")

	if err = server.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
