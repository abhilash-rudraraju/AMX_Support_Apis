package cache

import (
	"errors"
	"sync"

	"github.com/go-redis/redis/v7"
)

var (
	positionRedisClient *redis.Client
	positionMutex       = &sync.Mutex{}
)

func initPositionCache() {
	positionRedisClient = redis.NewClient(&redis.Options{
		Addr:     "redis-5554.amx1-redis.gpx.uat.angelone.in:5554",
		Password: "pass@123",
		DB:       0,
	})
}

func GetPositionCacheClient() (*redis.Client, error) {
	if positionRedisClient == nil {
		positionMutex.Lock()
		defer positionMutex.Unlock()

		if positionRedisClient == nil {
			initPositionCache()
		}
	}
	_, err := positionRedisClient.Ping().Result()
	if err != nil {
		return nil, errors.New("unable to get redis client")
	}
	return positionRedisClient, nil
}