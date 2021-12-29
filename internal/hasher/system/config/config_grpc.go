package config

import "common/config"

type GRPC struct {
	Host string `mapstructure:"host"`
}

func writeDefaultGRPCConfig(defaultConfig config.Default) {
	defaultConfig["grpc.host"] = "127.0.0.1:8090"
}

func (p *Provider) GRPC() GRPC {
	return p.config().GRPC
}
