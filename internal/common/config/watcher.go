package config

import (
	"common/errors"
	"context"
	"time"
)

type ConfigUpdater interface {
	UpdateConfig() error
}

func NewWatcher(
	configUpdater ConfigUpdater,
	aRemoteConfig RemoteConfig,
	watchingErrorHandler func(error),
) *Watcher {
	return &Watcher{
		configUpdater:       configUpdater,
		aRemoteConfig:       aRemoteConfig,
		handleWatchingError: watchingErrorHandler,

		watchingErrorChan: make(chan error),
	}
}

type Watcher struct {
	configUpdater ConfigUpdater
	aRemoteConfig RemoteConfig

	handleWatchingError func(error)
	watchingErrorChan   chan error
}

func (w *Watcher) Watch(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)

	go w.readWatchingErrors()

	for {
		select {
		case <-ticker.C:
			w.updateConfig()

		case <-ctx.Done():
			close(w.watchingErrorChan)

			break
		}
	}
}

func (w *Watcher) readWatchingErrors() {
	for err := range w.watchingErrorChan {
		w.handleWatchingError(err)
	}
}

const watcherComponentName = "config watcher"

func (w *Watcher) updateConfig() {
	if err := w.aRemoteConfig.ReadRemoteConfig(); err != nil {
		w.watchingErrorChan <- errors.Errorf("%: failed to read remote config: %w", watcherComponentName, err)

		return
	}

	if err := w.configUpdater.UpdateConfig(); err != nil {
		w.watchingErrorChan <- errors.Errorf("%s: failed to update config: %w", watcherComponentName, err)
	}
}
