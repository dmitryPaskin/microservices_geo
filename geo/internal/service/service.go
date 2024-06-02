package service

import (
	"GeoAPI/geo/internal/models/interfaceModels"
	"GeoAPI/geo/internal/models/structModels"
	"context"
	"github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/api/model"
	"github.com/ekomobile/dadata/v2/client"
)

type service struct {
	apiKeyValue    string
	secretKeyValue string
}

func NewService(ApiKeyValue, SecretKeyValue string) interfaceModels.Service {
	return &service{
		apiKeyValue:    ApiKeyValue,
		secretKeyValue: SecretKeyValue,
	}
}

func (s *service) SearchService(request structModels.SearchRequest) (*[]*model.Address, error) {
	cleanApi := dadata.NewCleanApi(client.WithCredentialProvider(&client.Credentials{
		ApiKeyValue:    s.apiKeyValue,
		SecretKeyValue: s.secretKeyValue}))

	addresses, err := cleanApi.Address(context.Background(), request.Query)
	if err != nil {
		return nil, err
	}

	return &addresses, nil
}

func (s *service) GeocodeService(request structModels.GeocodeRequest) (*[]*model.Address, error) {
	cleanApi := dadata.NewCleanApi(client.WithCredentialProvider(&client.Credentials{
		ApiKeyValue:    s.apiKeyValue,
		SecretKeyValue: s.secretKeyValue}))
	addresses, err := cleanApi.Address(context.Background(), request.Lat, request.Lon)
	if err != nil {
		return nil, err
	}

	return &addresses, nil
}
