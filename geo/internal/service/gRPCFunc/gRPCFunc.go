package gRPCFunc

import (
	"GeoAPI/geo/internal/models/interfaceModels"
	"GeoAPI/geo/internal/models/structModels"
	proto "GeoAPI/geo/pkg/gRPC/api"
	"context"
	"encoding/json"
)

type ServiceGRPC struct {
	interfaceModels.Service
	proto.UnimplementedGeoServiceServer
}

func (s *ServiceGRPC) SearchService(ctx context.Context, request *proto.SearchRequest) (*proto.AddressResponse, error) {
	req := structModels.SearchRequest{Query: request.Query}

	address, err := s.Service.SearchService(req)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(address)
	if err != nil {
		return nil, err
	}

	result := proto.AddressResponse{Address: data}

	return &result, nil
}

func (s *ServiceGRPC) GeocodeService(ctx context.Context, request *proto.GeocodeRequest) (*proto.AddressResponse, error) {
	req := structModels.GeocodeRequest{
		Lat: request.Lat,
		Lon: request.Lon,
	}

	address, err := s.Service.GeocodeService(req)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(address)
	if err != nil {
		return nil, err
	}

	result := proto.AddressResponse{Address: data}

	return &result, nil
}
