package config

import "github.com/spf13/viper"

type ConsulConfig struct {
	Host      string `mapstructure:"host"`
	ConfigKey string `mapstructure:"configKey"`
}

func readDefaultConsulConfig(v *viper.Viper) {
	v.SetDefault("consul.host", "127.0.0.1:8500")
	v.SetDefault("consul.configKey", "config")
}
