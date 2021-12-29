package config

import "common/config"

type Config struct {
	Redis  Redis  `mapstructure:"redis"`
	Hasher Hasher `mapstructure:"hasher"`
	Logger Logger `mapstructure:"logger"`
}

func defaultConfig() config.Default {
	cfg := config.Default{}

	writeDefaultLoggerConfig(cfg)
	writeDefaultRedisConfig(cfg)
	writeDefaultHasherConfig(cfg)

	return cfg
}
