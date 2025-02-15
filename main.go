package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	fmt.Println("ciao from redis tutor")

	rd := redis.NewClient(&redis.Options{
		// ideally these should be read from an .env file.
		Addr:     "localhost:6379",
		Password: "", // default
		DB:       0,  // default
	})

	// check connection to redis
	ping, err := rd.Ping(context.Background()).Result()
	if err != nil {
		fmt.Printf("error pinging redis conn: %s\n", err)
		return
	}
	fmt.Println(ping) // PONG

	lookupKey := "myKey"

	if err = rd.Set(context.Background(), lookupKey, "myValue", 0).Err(); err != nil {
		fmt.Printf("error setting value: %s\n", err)
		return
	}
	// TODO: https://www.youtube.com/watch?v=1C3Ym_JjkMw

}
