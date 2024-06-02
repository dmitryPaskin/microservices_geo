package geo

import (
	proto "GEO_API/proxy/pkg/gRPC/api/geo"
	"context"
	"google.golang.org/grpc"
)

type serviceGeo struct {
}

func New() proto.GeoServiceClient {
	return &serviceGeo{}
}

func (s *serviceGeo) SearchService(ctx context.Context, in *proto.SearchRequest, opts ...grpc.CallOption) (*proto.AddressResponse, error) {
	conn, err := grpc.Dial("geo:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err

	}
	defer conn.Close()

	client := proto.NewGeoServiceClient(conn)

	address, err := client.SearchService(ctx, in)
	if err != nil {
		return nil, err
	}

	return address, nil
}

func (s *serviceGeo) GeocodeService(ctx context.Context, in *proto.GeocodeRequest, opts ...grpc.CallOption) (*proto.AddressResponse, error) {
	conn, err := grpc.Dial("geo:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err

	}
	defer conn.Close()

	client := proto.NewGeoServiceClient(conn)

	address, err := client.GeocodeService(ctx, in)
	if err != nil {
		return nil, err
	}

	return address, nil
}