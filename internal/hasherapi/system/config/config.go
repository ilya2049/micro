package config

import "common/config"

type Config struct {
	Graylog Graylog `mapstructure:"graylog"`
	Redis   Redis   `mapstructure:"redis"`
	Hasher  Hasher  `mapstructure:"hasher"`
}

func defaultConfig() config.Default {
	cfg := config.Default{}

	writeDefaultGraylogConfig(cfg)
	writeDefaultRedisConfig(cfg)
	writeDefaultHasherConfig(cfg)

	return cfg
}
