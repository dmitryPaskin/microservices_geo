package service

import (
	"GEO_API/user/internal/repository"
	proto "GEO_API/user/pkg/gRPC/api"
	"context"
	"github.com/go-chi/jwtauth"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repository.AuthRepository
	proto.UnimplementedUserServiceServer
}

func NewService(r repository.AuthRepository) proto.UserServiceServer {
	return &service{AuthRepository: r}
}

func (s *service) GetToken(ctx context.Context, user *proto.User) (*proto.Token, error) {
	var tokenAuth = jwtauth.New("HS256", []byte("mySecret"), nil)
	userInBD, err := s.AuthRepository.GetUser(user)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userInBD.Password), []byte(user.Password)); err != nil {
		return nil, err
	}

	_, tokenString, err := tokenAuth.Encode(map[string]interface{}{"username": userInBD.Login, "password": userInBD.Password})

	token := proto.Token{
		Token: tokenString,
	}
	return &token, err
}

func (s *service) SaveUser(ctx context.Context, user *proto.User) (*empty.Empty, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	if err := s.AuthRepository.SaveUser(&proto.User{
		Login:    user.Login,
		Password: string(hashedPassword),
	}); err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *service) GetListUsers(context.Context, *empty.Empty) (*proto.Users, error) {
	users, err := s.AuthRepository.GetListUsers()
	if err != nil {
		return nil, err
	}

	return &proto.Users{Users: *users}, nil

}

func (s *service) CheckUser(ctx context.Context, user *proto.User) (*proto.Check, error) {
	isExist, err := s.AuthRepository.CheckUser(user)
	if err != nil {
		return nil, err
	}

	return &proto.Check{IsExist: isExist}, nil
}
