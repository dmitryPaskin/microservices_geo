package proxyService

import (
	"GEO_API/proxy/internal/models"
	"GEO_API/proxy/internal/service/user"
	proto "GEO_API/proxy/pkg/gRPC/api/user"
	"context"
	"fmt"
	"github.com/go-chi/jwtauth"
)

type User interface {
	GetCurrentUser(ctx context.Context) (*models.User, error)
	GetListUsers() ([]*models.User, error)
}

type userProxy struct {
	realObject proto.UserServiceClient
}

func NewUserProxy() User {
	realObject := user.New()
	return &userProxy{realObject: realObject}
}

func (u *userProxy) GetCurrentUser(ctx context.Context) (*models.User, error) {
	_, claims, err := jwtauth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	if claims == nil {
		return nil, fmt.Errorf("user is not found")
	}

	var result models.User

	login := fmt.Sprint(claims["username"])
	password := fmt.Sprint(claims["password"])

	result = models.User{
		Login:    login,
		Password: password,
	}

	return &result, nil
}

func (u *userProxy) GetListUsers() ([]*models.User, error) {

	ctx := context.Background()
	users, err := u.realObject.GetListUsers(ctx, nil)
	if err != nil {
		return nil, err
	}

	var result []*models.User

	for _, userProto := range users.Users {
		m := &models.User{
			Login:    userProto.Login,
			Password: userProto.Password,
		}

		result = append(result, m)
	}

	return result, nil
}

