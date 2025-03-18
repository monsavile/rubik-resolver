package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/monsavile/rubik-resolver/internal/config"
	resolverV1 "github.com/monsavile/rubik-resolver/pkg/resolver_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServerConfig interface {
	Port() int
}

type server struct {
	resolverV1.UnimplementedResolverV1Server
}

func (s *server) Resolve(ctx context.Context, req *resolverV1.ResolveRequest) (*resolverV1.ResolveResponse, error) {
	return &resolverV1.ResolveResponse{
		Cube: "test",
	}, nil
}

func main() {
	if err := config.Load(); err != nil {
		log.Fatalf("failed to load config: %s", err)
	}

	grpcServerConfig, err := config.NewGRPCServerConfig()
	if err != nil {
		log.Fatalf("failed to get grpc server config: %s", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcServerConfig.Port()))
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	grpcServer := grpc.NewServer()
	server := server{}

	reflection.Register(grpcServer)
	resolverV1.RegisterResolverV1Server(grpcServer, &server)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
