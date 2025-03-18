LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.5
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1

generate:
	make generate-cube_v1
	make generate-resolver_v1

generate-cube_v1:
	mkdir -p pkg/cube_v1
	protoc --proto_path api/cube_v1 \
	--go_out=pkg/cube_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/cube_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/cube_v1/cube.proto

generate-resolver_v1:
	mkdir -p pkg/resolver_v1
	protoc --proto_path api/resolver_v1 \
	--go_out=pkg/resolver_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/resolver_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/resolver_v1/resolver.proto
