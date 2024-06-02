package auth

import (
	proto "GEO_API/proxy/pkg/gRPC/api/auth"
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type serviceAuth struct {
}

func New() proto.AuthServiceClient {
	return &serviceAuth{}
}

func (s *serviceAuth) SingUpHandler(ctx context.Context, in *proto.UserAuth, opts ...grpc.CallOption) (*proto.Status, error) {
	conn, err := grpc.Dial("auth:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err

	}
	defer conn.Close()

	client := proto.NewAuthServiceClient(conn)
	if status, err := client.SingUpHandler(ctx, in); err != nil {
		return nil, err
	} else if !status.IsSuccessful {
		return nil, fmt.Errorf("failed to register user")
	}
	return nil, nil
}

func (s *serviceAuth) SingInHandler(ctx context.Context, in *proto.UserAuth, opts ...grpc.CallOption) (*proto.TokenAuth, error) {
	conn, err := grpc.Dial("auth:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err

	}
	defer conn.Close()

	client := proto.NewAuthServiceClient(conn)
	token, err := client.SingInHandler(ctx, in)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *serviceAuth) CheckToken(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	conn, err := grpc.Dial("auth:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err

	}
	defer conn.Close()

	client := proto.NewAuthServiceClient(conn)
	if _, err := client.CheckToken(ctx, in); err != nil {
		return nil, err
	}
	return nil, nil
}
