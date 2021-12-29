package config

import "common/config"

type Config struct {
	Logger Logger `mapstructure:"logger"`
	GRPC   GRPC   `mapstructure:"grpc"`
}

func defaultConfig() config.Default {
	cfg := config.Default{}

	writeDefaultLoggerConfig(cfg)
	writeDefaultGRPCConfig(cfg)

	return cfg
}
