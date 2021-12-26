package storage

import (
	"common/errors"
	"common/requestid"
	"context"
	"hasherapi/app/log"
	"hasherapi/domain/hash"
	"time"

	// Load resources in go binary
	_ "embed"

	"github.com/go-redis/redis/v8"
)

//go:embed redis-save-hashes.lua
var saveHashesScript string

type Config struct {
	Address  string
	Password string
}

type RedisStorage struct {
	redisClient         *redis.Client
	logger              log.Logger
	saveHashesScriptSHA string
}

func New(cfg Config, logger log.Logger) (*RedisStorage, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
	})

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)

	defer cancel()

	if err := rdb.ScriptFlush(ctx).Err(); err != nil {
		logger.LogWarn("failed to flush lua script", log.Details{
			log.FieldComponent: log.ComponentHashStorage,
		})
	}

	scriptLoadingResult := rdb.ScriptLoad(ctx, saveHashesScript)
	if err := scriptLoadingResult.Err(); err != nil {
		return nil, errors.Errorf(
			"%s: failed load lua script: %w", log.ComponentHashStorage, err,
		)
	}

	return &RedisStorage{
		redisClient:         rdb,
		logger:              logger,
		saveHashesScriptSHA: scriptLoadingResult.String(),
	}, nil
}

const (
	idGeneratorKey = "id:generator"
)

func (s *RedisStorage) Save(ctx context.Context, sha3Hashes hash.SHA3Hashes) ([]hash.IdentifiedSHA3Hash, error) {
	evalResult := s.redisClient.Eval(ctx, saveHashesScript,
		[]string{idGeneratorKey}, sha3HashesToEmptyInterfaces(sha3Hashes)...,
	)

	if err := evalResult.Err(); err != nil {
		return []hash.IdentifiedSHA3Hash{}, errors.Errorf(
			"%s: failed to save sha3 hashes in redis storage: %w", log.ComponentHashStorage, err,
		)
	}

	generatedIDs, err := evalResult.Int64Slice()
	if err != nil {
		return []hash.IdentifiedSHA3Hash{}, errors.Errorf(
			"%s: failed to retreive generated ids from redis storage: %w", log.ComponentHashStorage, err,
		)
	}

	if len(sha3Hashes) != len(generatedIDs) {
		requestID := requestid.Get(ctx)

		s.logger.LogWarn("save: quantity of sha3 hashes isn't equal quantity of generated ids", log.Details{
			log.FieldComponent: log.ComponentHashStorage,
			log.FieldRequestID: requestID,
		})
	}

	return sha3HashesToIdentifiedSHA3Hashes(sha3Hashes, generatedIDs), nil
}

func (s *RedisStorage) Get(ctx context.Context, hashIDs []hash.ID) ([]hash.IdentifiedSHA3Hash, error) {
	mgetResult := s.redisClient.MGet(ctx, hashIDsToRedisKeys(hashIDs)...)

	if err := mgetResult.Err(); err != nil {
		return []hash.IdentifiedSHA3Hash{}, errors.Errorf(
			"%s: failed to get sha3 hashes from redis storage: %w", log.ComponentHashStorage, err,
		)
	}

	sha3HashesAsEmptyInterfaces := mgetResult.Val()
	if len(sha3HashesAsEmptyInterfaces) != len(hashIDs) {
		requestID := requestid.Get(ctx)

		s.logger.LogWarn("get: quantity of retrieved sha3 hashes isn't equal quantity of hash ids", log.Details{
			log.FieldComponent: log.ComponentHashStorage,
			log.FieldRequestID: requestID,
		})
	}

	return emptyInterfacesToIdentifiedHashes(sha3HashesAsEmptyInterfaces, hashIDs), nil
}
