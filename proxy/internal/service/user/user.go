package user

import (
	proto "GEO_API/proxy/pkg/gRPC/api/user"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type serviceUser struct {
}

func New() proto.UserServiceClient {
	return &serviceUser{}
}

func (s *serviceUser) GetToken(ctx context.Context, in *proto.User, opts ...grpc.CallOption) (*proto.Token, error) {
	conn, err := grpc.Dial("user:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err

	}
	defer conn.Close()

	client := proto.NewUserServiceClient(conn)

	token, err := client.GetToken(ctx, in)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *serviceUser) GetListUsers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*proto.Users, error) {
	conn, err := grpc.Dial("user:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err

	}
	defer conn.Close()

	client := proto.NewUserServiceClient(conn)

	users, err := client.GetListUsers(ctx, in)

	return users, err
}

func (s *serviceUser) CheckUser(ctx context.Context, in *proto.User, opts ...grpc.CallOption) (*proto.Check, error) {
	return nil, nil
}

func (s *serviceUser) SaveUser(ctx context.Context, in *proto.User, opts ...grpc.CallOption) (*empty.Empty, error) {
	return nil, nil
}
