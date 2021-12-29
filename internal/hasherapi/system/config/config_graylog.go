package config

import "common/config"

type Graylog struct {
	Host   string `mapstructure:"host"`
	Source string `mapstructure:"source"`
}

func writeDefaultGraylogConfig(defaultConfig config.Default) {
	defaultConfig["graylog.host"] = "127.0.0.1:12201"
	defaultConfig["graylog.source"] = "hasherapi"
}

func (p *Provider) Graylog() Graylog {
	return p.config().Graylog
}
