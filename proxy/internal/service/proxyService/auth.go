package proxyService

import (
	"GEO_API/proxy/internal/models"
	"GEO_API/proxy/internal/service/auth"
	proto "GEO_API/proxy/pkg/gRPC/api/auth"
	"context"
)

type Auth interface {
	SingUpHandler(user models.User) error
	SingInHandler(user models.User) (string, error)
}

type proxyAuth struct {
	realObject proto.AuthServiceClient
}

func NewProxyAuth() Auth {
	realObject := auth.New()
	return &proxyAuth{realObject}
}

func (p *proxyAuth) SingUpHandler(user models.User) error {
	userAuth := proto.UserAuth{
		Login:    user.Login,
		Password: user.Password,
	}
	ctx := context.Background()

	if _, err := p.realObject.SingUpHandler(ctx, &userAuth); err != nil {
		return err
	}
	return nil
}

func (p *proxyAuth) SingInHandler(user models.User) (string, error) {
	userAuth := proto.UserAuth{
		Login:    user.Login,
		Password: user.Password,
	}
	ctx := context.Background()

	token, err := p.realObject.SingInHandler(ctx, &userAuth)

	return token.Token, err
}