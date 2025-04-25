package config

import "os"

type Config struct {
    ServerAddress      string
    GRPCServerAddress  string
}

func LoadConfig() (*Config, error) {
    return &Config{
        ServerAddress:      os.Getenv("SERVER_ADDRESS"),
        GRPCServerAddress:  os.Getenv("GRPC_SERVER_ADDRESS"),
    }, nil
}
