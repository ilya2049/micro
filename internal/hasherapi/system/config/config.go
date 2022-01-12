package config

import "common/config"

type Config struct {
	Redis  Redis  `mapstructure:"redis"`
	Kafka  Kafka  `mapstructure:"kafka"`
	Hasher Hasher `mapstructure:"hasher"`
	Logger Logger `mapstructure:"logger"`
}

func defaultConfig() config.Default {
	cfg := config.Default{}

	writeDefaultLoggerConfig(cfg)
	writeDefaultRedisConfig(cfg)
	writeDefaultHasherConfig(cfg)
	writeDefaultKafkaConfig(cfg)

	return cfg
}
