package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type Person struct {
	ID   string
	Name string `json:name`
	Age  int    `json:name`
	Job  string `json:job"`
}

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

	lookupKey := uuid.NewString()
	fmt.Printf("lookup key: %s\n", lookupKey)
	data, err := json.Marshal(Person{
		ID:   lookupKey,
		Age:  21,
		Name: "Savage",
		Job:  "Busy",
	})
	if err != nil {
		fmt.Printf("error converting to JSON: %s\n", err)
		return
	}

	if err = rd.Set(context.Background(), lookupKey, data, 0).Err(); err != nil {
		fmt.Printf("error setting value: %s\n", err)
		return
	}

	val, err := rd.Get(context.Background(), lookupKey).Result()
	if err != nil {
		fmt.Printf("error fetching value %s\n", err)
		return
	}
	fmt.Println(val)
}
