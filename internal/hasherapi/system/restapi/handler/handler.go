package handler

import (
	"common/cleanup"
	commonConfig "common/config"
	"common/log/logrus"
	"crypto/tls"
	"hasherapi/app/event"
	apphash "hasherapi/app/hash"
	"hasherapi/app/log"
	"hasherapi/domain/hash"
	"hasherapi/system/config"
	"hasherapi/system/eventstream"
	"hasherapi/system/hash/calculator"
	"hasherapi/system/hash/storage"
	"hasherapi/system/restapi/middlewares"
	"hasherapi/system/restapi/operations"
	stdlog "log"
	"net/http"
	"time"
)

func New() *Handler {
	/* --- Cleanup funcs ----------------------- */
	cleanupFuncs := cleanup.Funcs{}

	/* --- Config ----------------------- */
	configProvider, err := config.NewProvider()
	if err != nil {
		stdlog.Fatalf("%s: failed to create a configurator: %s", log.ComponentAppInitializer, err.Error())
	}

	/* --- Logger ----------------------- */
	logger, updateLogLevel := logrus.NewLogger(logrus.Config{
		GraylogHost:   configProvider.Logger().Graylog.Host,
		GraylogSource: configProvider.Logger().Graylog.Source,
		LogLevel:      log.Level(configProvider.Logger().Level),
	})

	stopConfigWatching := config.Watch(configProvider, logger, []commonConfig.Trigger{
		commonConfig.Trigger(func() {
			updateLogLevel(log.Level(configProvider.Logger().Level))
		}),
	})

	cleanupFuncs = append(cleanupFuncs, stopConfigWatching)

	/* --- Event stream ----------------------- */
	kafkaWriter, disconnectKafka, err := eventstream.NewKafkaWriter(eventstream.KafkaWriterConfig{
		Host:     configProvider.Kafka().Host,
		Username: configProvider.Kafka().Username,
		Password: configProvider.Kafka().Password,
		Topic:    configProvider.Kafka().Topic,
	}, logger)

	if err != nil {
		stdlog.Fatalf("%s: failed to create a kafka writer: %s", log.ComponentAppInitializer, err.Error())
	}

	cleanupFuncs = append(cleanupFuncs, disconnectKafka)

	eventStream, stopStream := eventstream.New(kafkaWriter, logger, 2)
	cleanupFuncs = append(cleanupFuncs, stopStream)

	/* --- Hash calculator ----------------------- */
	var hashCalculator hash.Calculator
	hashCalculator = calculator.NewGRPCCalculator(func() calculator.Config {
		return calculator.Config{
			URL:     configProvider.Hasher().Host,
			Timeout: time.Duration(configProvider.Hasher().TimeoutSec) * time.Second,
		}
	},
		logger,
	)
	hashCalculator = apphash.WrapCalculatorWithLogger(hashCalculator, logger)

	/* --- Hash storage ----------------------- */
	var hashStorage hash.Storage
	hashStorage, closeRedisConnectionsFunc, err := storage.New(storage.Config{
		Address:  configProvider.Redis().Host,
		Password: configProvider.Redis().Password,
	}, logger)

	if err != nil {
		cleanupFuncs.Invoke()

		logger.LogFatal("failed to create a hash storage: "+err.Error(), log.Details{
			log.FieldComponent: log.ComponentAppInitializer,
		})
	}

	cleanupFuncs = append(cleanupFuncs, closeRedisConnectionsFunc)

	hashStorage = apphash.WrapStorageWithLogger(hashStorage, logger)

	/* --- Hash service ----------------------- */
	hashService := hash.NewService(hashCalculator, hashStorage)

	/* --- REST api ----------------------- */
	errorResponderFactory := middlewares.NewResponderFactory(logger)

	return &Handler{
		hashHandler:  newHashHandler(hashService, errorResponderFactory),
		cleanupFuncs: cleanupFuncs,
		logger:       logger,
		eventStream:  eventStream,
	}
}

type Handler struct {
	*hashHandler

	cleanupFuncs cleanup.Funcs

	logger      log.Logger
	eventStream event.Stream
}

func (h *Handler) ConfigureFlags(api *operations.HasherapiAPI) {}

func (h *Handler) ConfigureTLS(tlsConfig *tls.Config) {}

func (h *Handler) ConfigureServer(s *http.Server, scheme, addr string) {}

func (h *Handler) CustomConfigure(api *operations.HasherapiAPI) {
	api.Logger = h.logger.Printf

	api.ServerShutdown = func() {
		h.cleanupFuncs.Invoke()
	}
}

func (h *Handler) SetupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

func (h *Handler) SetupGlobalMiddleware(handler http.Handler) http.Handler {
	handler = middlewares.Logging(handler, h.logger)
	handler = middlewares.TraceRequest(handler)
	handler = middlewares.RegisterCallEvents(handler, h.eventStream)

	return handler
}
