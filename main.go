package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type redisCache struct {
	redisClient *redis.Client
}

func NewRedisCache(redisClient *redis.Client) *redisCache {
	return &redisCache{redisClient}
}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // 0 means to use default DB
	})

	ExampleStrings(rdb)
	ExampleHashes(rdb)
}

func ExampleStrings(rdb *redis.Client) {

	// set
	err := rdb.Set(ctx, "mykey", "hello", 0).Err()
	if err != nil {
		panic(err)
	}

	// get
	val, err := rdb.Get(ctx, "mykey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("mykey", val)

	// get
	val2, err := rdb.Get(ctx, "yourkey").Result()
	if err == redis.Nil {
		fmt.Println("yourkey does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("yourkey", val2)
	}
}

func ExampleHashes(rdb *redis.Client) {

	// set
	err := rdb.HMSet(ctx, "arash:1001", "username", "tropicanobel", "age", "30").Err()
	if err != nil {
		panic(err)
	}

	// get
	val, err := rdb.HGetAll(ctx, "arash:1001").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("arashkey:1001", val["username"])
	fmt.Println("arashkey:1001", val["age"])
}
