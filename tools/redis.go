package tools

import (
	"log"

	"github.com/go-redis/redis"
)

// Redis is instance of redis
var Redis *redis.Client

//ConnectRedis method with init connection with redis instance
func ConnectRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := Redis.Ping().Result()
	if err != nil {
		log.Panic(err)
	}
}
