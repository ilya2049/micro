package config

import (
	"common/config"
	"common/errors"
	"context"
	"hasherapi/app/log"
	"sync"
)

func NewProvider(logger log.Logger) (provider *Provider, stopConfigWatching func(), err error) {
	remoteConfig, err := config.NewRemoteConfig(defaultConfig(), "HASHERAPI")
	if err != nil {
		return nil, nil, err
	}

	provider = &Provider{remoteConfig: remoteConfig}

	if err := provider.UpdateConfig(); err != nil {
		return nil, nil, err
	}

	watcher := config.NewWatcher(provider, remoteConfig, func(e error) {
		logger.LogError(e.Error(), log.Details{
			log.FieldComponent: log.ComponentConfigurator,
		})
	})

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go watcher.Watch(ctx)

	return provider, cancel, nil
}

type Provider struct {
	remoteConfig config.RemoteConfig

	mutex   sync.RWMutex
	aConfig Config
}

func (p *Provider) UpdateConfig() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	var aConfig Config

	if err := p.remoteConfig.ReadConfigIn(&aConfig); err != nil {
		return errors.Errorf("%s: failed to unmarshal app config: %w", log.ComponentConfigurator, err)
	}

	p.aConfig = aConfig

	return nil
}

func (p *Provider) config() Config {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	return p.aConfig
}
