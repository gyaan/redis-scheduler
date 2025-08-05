// redisclient provides the redis connection
package redisclient

import (
	"fmt"
	"github.com/go-redis/redis"
	"redis-scheduler/config"
)

// IRedisClient defines the interface for a Redis client.
type IRedisClient interface {
	RPush(key string, args ...interface{}) *redis.IntCmd
	LRange(key string, start, stop int64) *redis.StringSliceCmd
	Del(keys ...string) *redis.IntCmd
	LPop(key string) *redis.StringCmd
	Ping() *redis.StatusCmd
	Close() error
}

var (
	client IRedisClient
)

// Init initializes the redis client
func Init(conf *config.Config) {
	client = redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
	})
}

// GetClient get redis client
func GetClient() IRedisClient {
	return client
}

// SetClient sets the redis client (for testing)
func SetClient(c IRedisClient) {
	client = c
}

// Ping check redis connection
func Ping() error {
	do, err := GetClient().Ping().Result()

	if err != nil || do == "" {
		fmt.Println("Unable to connect to redis", err)
		return err
	}
	fmt.Println("Response for redis PING => ", do)
	return nil
}
