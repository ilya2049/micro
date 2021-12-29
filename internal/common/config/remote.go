package config

import (
	"strings"

	"github.com/spf13/viper"
	// Enable remote viper features
	_ "github.com/spf13/viper/remote"
)

type Default map[string]interface{}

type RemoteConfig interface {
	ReadRemoteConfig() error
	ReadConfigIn(interface{}) error
}

type remoteConfig struct {
	v *viper.Viper
}

func (c *remoteConfig) ReadRemoteConfig() error {
	return c.v.ReadRemoteConfig()
}

func (c *remoteConfig) ReadConfigIn(cfg interface{}) error {
	return c.v.Unmarshal(cfg)
}

func NewRemoteConfig(defaultConfig Default, environmentVariablePrefix string,
) (RemoteConfig, error) {
	v := viper.New()

	readDefaultConfig(v, defaultConfig)
	readEnvironmentConfig(v, environmentVariablePrefix)

	if err := readConsulConfig(v); err != nil {
		return nil, err
	}

	return &remoteConfig{v: v}, nil
}

func readDefaultConfig(v *viper.Viper, defaultValues map[string]interface{}) {
	readDefaultConsulConfig(v)

	for option, defaultValue := range defaultValues {
		v.SetDefault(option, defaultValue)
	}
}

func readEnvironmentConfig(v *viper.Viper, environmentVariablePrefix string) {
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetEnvPrefix(environmentVariablePrefix)
	v.AutomaticEnv()
}

func readConsulConfig(v *viper.Viper) error {
	err := v.AddRemoteProvider("consul",
		v.GetString("consul.host"),
		v.GetString("consul.configKey"),
	)

	if err != nil {
		return err
	}

	v.SetConfigType("json")
	if err := v.ReadRemoteConfig(); err != nil {
		return err
	}

	return nil
}
