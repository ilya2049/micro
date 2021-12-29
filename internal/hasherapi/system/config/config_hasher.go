package config

import "common/config"

type Hasher struct {
	Host       string `mapstructure:"host"`
	TimeoutSec int    `mapstructure:"timeoutSec"`
}

func writeDefaultHasherConfig(defaultConfig config.Default) {
	defaultConfig["hasher.host"] = "127.0.0.1:8090"
	defaultConfig["hasher.timeoutSec"] = 1
}

func (p *Provider) Hasher() Hasher {
	return p.config().Hasher
}
