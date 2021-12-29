package config

import "common/config"

type Logger struct {
	Graylog Graylog `mapstructure:"graylog"`
	Level   string  `mapstructure:"level"`
}

type Graylog struct {
	Host   string `mapstructure:"host"`
	Source string `mapstructure:"source"`
}

func writeDefaultLoggerConfig(defaultConfig config.Default) {
	defaultConfig["logger.graylog.host"] = "127.0.0.1:12201"
	defaultConfig["logger.graylog.source"] = "hasher"
	defaultConfig["logger.level"] = "info"
}

func (p *Provider) Logger() Logger {
	return p.config().Logger
}
