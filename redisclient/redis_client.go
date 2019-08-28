// redisclient provides the redis connection
package redisclient

import (
	"fmt"
	"github.com/go-redis/redis"
)

// GetClient get redis client
func GetClient() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}

// Ping check redis connection
func Ping() error {
	client := GetClient()
	do, err := client.Ping().Result()

	if err != nil || do == "" {
		fmt.Println("Unable to connect to redis", err)
		return err
	}
	fmt.Println("Response for redis PING => ", do)
	return nil
}
