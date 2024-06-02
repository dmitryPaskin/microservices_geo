package service

import (
	"GEO_API/auth/internal/service/userClients"
	protoAuth "GEO_API/auth/pkg/gRPC/api/auth"
	protoUser "GEO_API/auth/pkg/gRPC/api/user"
	"context"
	"github.com/go-chi/jwtauth"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/lestrrat-go/jwx/jwt"
)

type Service struct {
	userClients.Clients
	protoAuth.UnimplementedAuthServiceServer
}

func New(address string) Service {
	return Service{
		Clients: userClients.Clients{Address: address},
	}
}

func (s *Service) SingUpHandler(ctx context.Context, userAuth *protoAuth.UserAuth) (*protoAuth.Status, error) {
	var status protoAuth.Status

	user := protoUser.User{
		Login:    userAuth.Login,
		Password: userAuth.Password,
	}

	if err := s.Clients.SaveUser(&user); err != nil {
		status.IsSuccessful = false
		return &status, err
	}

	status.IsSuccessful = true
	return &status, nil
}

func (s *Service) SingInHandler(ctx context.Context, userAuth *protoAuth.UserAuth) (*protoAuth.TokenAuth, error) {
	var token protoAuth.TokenAuth

	user := protoUser.User{
		Login:    userAuth.Login,
		Password: userAuth.Password,
	}

	tokenString, err := s.Clients.GetToken(&user)
	if err != nil {
		return nil, err
	}

	token = protoAuth.TokenAuth{
		Token: tokenString,
	}
	return &token, err
}

func (s *Service) CheckToken(ctx context.Context, in *empty.Empty) (*empty.Empty, error) {
	tokenAuth := jwtauth.New("HS256", []byte("mySecret"), nil)
	jwtauth.Verifier(tokenAuth)

	token, _, err := jwtauth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	if token == nil || jwt.Validate(token) != nil {
		return nil, err
	}

	return nil, nil
}
