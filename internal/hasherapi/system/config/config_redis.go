package config

import "common/config"

type Redis struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
}

func writeDefaultRedisConfig(defaultConfig config.Default) {
	defaultConfig["redis.host"] = "127.0.0.1:6379"
	defaultConfig["redis.password"] = "123"
}

func (p *Provider) Redis() Redis {
	return p.config().Redis
}
