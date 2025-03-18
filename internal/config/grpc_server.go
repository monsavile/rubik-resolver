package config

import (
	"errors"
	"os"
	"strconv"
)

const portEnvName = "SERVER_PORT"

type grpcServerConfig struct {
	port int
}

func NewGRPCServerConfig() (*grpcServerConfig, error) {
	portStr := os.Getenv(portEnvName)
	if len(portStr) == 0 {
		return nil, errors.New("grpc server port not found")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, errors.New("grpc server port is not a valid number")
	}

	return &grpcServerConfig{
		port: port,
	}, nil
}

func (c *grpcServerConfig) Port() int {
	return c.port
}
