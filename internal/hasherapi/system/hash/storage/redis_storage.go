package storage

import (
	"common/errors"
	"common/requestid"
	"context"
	"hasherapi/app/log"
	"hasherapi/domain/hash"
	"time"

	// To load resources in go binary
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
	redisClient          *redis.Client
	logger               log.Logger
	saveHashesScriptSHA1 string
}

func New(cfg Config, logger log.Logger) (redisStorage *RedisStorage, closeRedisConnections func(), err error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
	})

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, nil, errors.Errorf(
			"%s: failed to connect redis: %w", log.ComponentHashStorage, err,
		)
	}

	logger.LogInfo("redis connection established", log.Details{
		log.FieldComponent: log.ComponentHashStorage,
	})

	scriptSHA1, err := reloadScript(ctx, rdb)
	if err != nil {
		return nil, nil, errors.Errorf(
			"%s: failed to reload the save hashes script: %w", log.ComponentHashStorage, err,
		)
	}

	redisStorage = &RedisStorage{
		redisClient:          rdb,
		logger:               logger,
		saveHashesScriptSHA1: scriptSHA1,
	}

	closeRedisConnections = func() {
		if err := rdb.Close(); err != nil {
			logger.LogWarn("failed to close redis connections: "+err.Error(), log.Details{
				log.FieldComponent: log.ComponentHashStorage,
			})

			return
		}

		logger.LogInfo("redis connections successfully closed", log.Details{
			log.FieldComponent: log.ComponentHashStorage,
		})
	}

	return redisStorage, closeRedisConnections, nil
}

func reloadScript(ctx context.Context, rdb *redis.Client) (scriptSHA1 string, err error) {
	tx := rdb.TxPipeline()

	tx.ScriptFlush(ctx)
	scriptLoadingResult := tx.ScriptLoad(ctx, saveHashesScript)

	if _, err := tx.Exec(ctx); err != nil {
		return "", errors.Errorf(
			"%s: failed to reload lua script: %w", log.ComponentHashStorage, err,
		)
	}

	return scriptLoadingResult.Val(), nil
}

const (
	idGeneratorKey = "id:generator"
)

func (s *RedisStorage) Save(ctx context.Context, sha3Hashes hash.SHA3Hashes) ([]hash.IdentifiedSHA3Hash, error) {
	saveHashesScriptResult := s.redisClient.EvalSha(ctx, s.saveHashesScriptSHA1,
		[]string{idGeneratorKey}, sha3HashesToEmptyInterfaces(sha3Hashes)...,
	)

	if err := saveHashesScriptResult.Err(); err != nil {
		return []hash.IdentifiedSHA3Hash{}, errors.Errorf(
			"%s: failed to save sha3 hashes in redis storage: %w", log.ComponentHashStorage, err,
		)
	}

	generatedIDs, err := saveHashesScriptResult.Int64Slice()
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
