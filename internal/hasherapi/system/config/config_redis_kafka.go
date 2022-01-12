package config

import "common/config"

type Kafka struct {
	Host  string `mapstructure:"host"`
	Topic string `mapstructure:"topic"`
}

func writeDefaultKafkaConfig(defaultConfig config.Default) {
	defaultConfig["kafka.host"] = "127.0.0.1:9092"
	defaultConfig["kafka.topic"] = "test"
}

func (p *Provider) Kafka() Kafka {
	return p.config().Kafka
}
