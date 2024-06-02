package proxyService

import (
	"GEO_API/proxy/internal/models"
	"GEO_API/proxy/internal/service/geo"
	proto "GEO_API/proxy/pkg/gRPC/api/geo"
	"context"
	"encoding/json"
)

type Geo interface {
	Address(request models.SearchRequest) ([]*models.AddressSearch, error)
	Geocode(request models.GeocodeRequest) (*models.AddressGeo, error)
}

type geoProxy struct {
	realObject proto.GeoServiceClient
}

func NewGeoProxy() Geo {
	realObject := geo.New()
	return &geoProxy{
		realObject: realObject,
	}
}

func (g *geoProxy) Address(request models.SearchRequest) ([]*models.AddressSearch, error) {
	gRPCRequest := proto.SearchRequest{
		Query: request.Query,
	}
	ctx := context.Background()

	address, err := g.realObject.SearchService(ctx, &gRPCRequest)
	if err != nil {
		return nil, err
	}

	var result []*models.AddressSearch

	if err := json.Unmarshal(address.Address, &result); err != nil {
		return nil, err
	}

	return result, err
}

func (g *geoProxy) Geocode(request models.GeocodeRequest) (*models.AddressGeo, error) {
	gRPCRequest := proto.GeocodeRequest{
		Lat: request.Lat,
		Lon: request.Lon,
	}
	ctx := context.Background()

	address, err := g.realObject.GeocodeService(ctx, &gRPCRequest)
	if err != nil {
		return nil, err
	}

	var result *models.AddressGeo

	if err := json.Unmarshal(address.Address, &result); err != nil {
		return nil, err
	}

	return result, nil
}
