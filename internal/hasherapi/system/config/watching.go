package config

import (
	"common/config"
	"context"
	"hasherapi/app/log"
)

func Watch(provider *Provider, logger log.Logger, triggers []config.Trigger) (stopWatchingFunc func()) {
	watcher := config.NewWatcher(provider, provider.remoteConfig, triggers, func(e error) {
		logger.LogError(e.Error(), log.Details{
			log.FieldComponent: log.ComponentConfigurator,
		})
	})

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go watcher.Watch(ctx)

	return cancel
}
