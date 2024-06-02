PHONY: generate
generate:
		mkdir -p pkg/gRPC
		protoc --go_out=pkg/gRPC --go_opt=paths=source_relative \
				--go-grpc_out=pkg/gRPC --go-grpc_opt=paths=source_relative \
				api/geo.proto