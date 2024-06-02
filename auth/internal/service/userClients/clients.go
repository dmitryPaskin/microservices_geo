package userClients

import (
	protoUser "GEO_API/auth/pkg/gRPC/api/user"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

type Clients struct {
	Address string
}

func (c *Clients) SaveUser(user *protoUser.User) error {
	conn, err := grpc.Dial(c.Address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := protoUser.NewUserServiceClient(conn)
	ctx := context.Background()
	check, err := client.CheckUser(ctx, user)
	if err != nil {
		return err
	} else if check.IsExist {
		return fmt.Errorf("the user is already registered")
	}

	if _, err := client.SaveUser(ctx, user); err != nil {
		return err
	}
	return nil
}

func (c *Clients) GetToken(user *protoUser.User) (string, error) {
	conn, err := grpc.Dial(c.Address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := protoUser.NewUserServiceClient(conn)

	ctx := context.Background()
	if Check, err := client.CheckUser(ctx, user); err != nil {
		return "", err
	} else if !Check.IsExist {
		return "", fmt.Errorf("the user is not exist")
	}
	token, err := client.GetToken(ctx, user)
	if err != nil {
		return "", err
	}

	return token.Token, nil
}
