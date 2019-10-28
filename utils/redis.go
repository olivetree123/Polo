package utils

import (
	"github.com/go-redis/redis"
	. "polo/common"
	"time"
)

var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func SetCache(key string, value interface{}, timeout int) error {
	expire := time.Second * time.Duration(timeout)
	err := RedisClient.Set(key, value, expire).Err()
	if err != nil {
		Logger.Error(err)
		return err
	}
	return nil
}

func GetCache(key string, result interface{}) error {
	err := RedisClient.Get(key).Scan(result)
	if err != nil {
		return nil
	}
	return nil
}
